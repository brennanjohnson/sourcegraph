package correlation

import (
	"errors"
	"fmt"
	"path"

	protocol "github.com/sourcegraph/lsif-protocol"
	"github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/stores/lsifstore"
)

func gatherSymbols(state *State) ([]*lsifstore.Symbol, error) {
	symbolByID := map[string]*lsifstore.Symbol{}

	for docID, resultID := range state.TextDocumentDocumentSymbol {
		docURI, ok := state.DocumentData[docID]
		if !ok {
			return nil, fmt.Errorf("Document ID %v doesn't exist in state.DocumentData", docID)
		}
		docSymbols, ok := state.DocumentSymbols[resultID]
		if !ok {
			return nil, fmt.Errorf("DocumentSymbolResult %v doesn't exist in state.DocumentSymbols", resultID)
		}
		for _, docSymbol := range docSymbols {
			symbol, err := convertSymbol(state, docURI, "", docSymbol)
			if err != nil {
				return nil, err
			}
			existingSymbol, ok := symbolByID[symbol.Identifier]
			if ok {
				// TODO(beyang): other rectification?
				existingSymbol.Children = append(existingSymbol.Children, symbol.Children...)
			} else {
				symbolByID[symbol.Identifier] = symbol
			}
		}
	}

	// TODO: Associate monikers based on range

	symbols := make([]*lsifstore.Symbol, 0, len(symbolByID))
	for _, symbol := range symbolByID {
		symbols = append(symbols, symbol)
	}

	return symbols, nil
}

func convertSymbol(state *State, filename string, parentID string, ds *protocol.RangeBasedDocumentSymbol) (*lsifstore.Symbol, error) {
	rng := state.RangeData[int(ds.ID)] // TODO(beyang): uint to int conversion
	if rng.Tag == nil {
		return nil, errors.New("RangeBasedDocumentSymbol range has no tag")
	}

	symbolID := path.Join(parentID, rng.Tag.Text)
	var children []*lsifstore.Symbol
	if ds.Children != nil {
		children = make([]*lsifstore.Symbol, len(ds.Children))
		for i, child := range ds.Children {
			symbolChild, err := convertSymbol(state, filename, symbolID, child)
			if err != nil {
				return nil, err
			}
			children[i] = symbolChild
		}
	}

	var symbolMonikers []lsifstore.MonikerData
	monikers := state.Monikers.Get(int(ds.ID)) // TODO(beyang): uint to int conversion
	if monikers != nil && monikers.Len() > 0 {
		symbolMonikers = make([]lsifstore.MonikerData, 0, monikers.Len())
		monikers.Each(func(m int) {
			moniker := state.MonikerData[m]
			symbolMonikers = append(symbolMonikers, lsifstore.MonikerData{
				Kind:       moniker.Kind,
				Scheme:     moniker.Scheme,
				Identifier: moniker.Identifier,
			})
		})
	}

	return &lsifstore.Symbol{
		Identifier: symbolID,
		Text:       rng.Tag.Text,
		Detail:     rng.Tag.Detail,
		Kind:       rng.Tag.Kind,
		Tags:       rng.Tag.Tags,
		Locations: []lsifstore.SymbolLocation{
			{
				Path: filename,
				Range: lsifstore.Range{
					Start: lsifstore.Position{Line: rng.Start.Line, Character: rng.Start.Character},
					End:   lsifstore.Position{Line: rng.End.Line, Character: rng.End.Character},
				},
			},
		},
		Monikers: symbolMonikers,
		Children: children,
	}, nil
}