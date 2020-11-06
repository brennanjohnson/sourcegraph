// Code generated by github.com/efritz/go-mockgen 0.1.0; DO NOT EDIT.

package mocks

import (
	"context"
	graphqlbackend "github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
	resolvers "github.com/sourcegraph/sourcegraph/enterprise/cmd/frontend/internal/codeintel/resolvers"
	dbstore "github.com/sourcegraph/sourcegraph/enterprise/internal/codeintel/stores/dbstore"
	"sync"
)

// MockResolver is a mock implementation of the Resolver interface (from the
// package
// github.com/sourcegraph/sourcegraph/enterprise/cmd/frontend/internal/codeintel/resolvers)
// used for unit testing.
type MockResolver struct {
	// DeleteIndexByIDFunc is an instance of a mock function object
	// controlling the behavior of the method DeleteIndexByID.
	DeleteIndexByIDFunc *ResolverDeleteIndexByIDFunc
	// DeleteUploadByIDFunc is an instance of a mock function object
	// controlling the behavior of the method DeleteUploadByID.
	DeleteUploadByIDFunc *ResolverDeleteUploadByIDFunc
	// GetIndexByIDFunc is an instance of a mock function object controlling
	// the behavior of the method GetIndexByID.
	GetIndexByIDFunc *ResolverGetIndexByIDFunc
	// GetUploadByIDFunc is an instance of a mock function object
	// controlling the behavior of the method GetUploadByID.
	GetUploadByIDFunc *ResolverGetUploadByIDFunc
	// IndexConnectionResolverFunc is an instance of a mock function object
	// controlling the behavior of the method IndexConnectionResolver.
	IndexConnectionResolverFunc *ResolverIndexConnectionResolverFunc
	// QueryResolverFunc is an instance of a mock function object
	// controlling the behavior of the method QueryResolver.
	QueryResolverFunc *ResolverQueryResolverFunc
	// UploadConnectionResolverFunc is an instance of a mock function object
	// controlling the behavior of the method UploadConnectionResolver.
	UploadConnectionResolverFunc *ResolverUploadConnectionResolverFunc
}

// NewMockResolver creates a new mock of the Resolver interface. All methods
// return zero values for all results, unless overwritten.
func NewMockResolver() *MockResolver {
	return &MockResolver{
		DeleteIndexByIDFunc: &ResolverDeleteIndexByIDFunc{
			defaultHook: func(context.Context, int) error {
				return nil
			},
		},
		DeleteUploadByIDFunc: &ResolverDeleteUploadByIDFunc{
			defaultHook: func(context.Context, int) error {
				return nil
			},
		},
		GetIndexByIDFunc: &ResolverGetIndexByIDFunc{
			defaultHook: func(context.Context, int) (dbstore.Index, bool, error) {
				return dbstore.Index{}, false, nil
			},
		},
		GetUploadByIDFunc: &ResolverGetUploadByIDFunc{
			defaultHook: func(context.Context, int) (dbstore.Upload, bool, error) {
				return dbstore.Upload{}, false, nil
			},
		},
		IndexConnectionResolverFunc: &ResolverIndexConnectionResolverFunc{
			defaultHook: func(dbstore.GetIndexesOptions) *resolvers.IndexesResolver {
				return nil
			},
		},
		QueryResolverFunc: &ResolverQueryResolverFunc{
			defaultHook: func(context.Context, *graphqlbackend.GitBlobLSIFDataArgs) (resolvers.QueryResolver, error) {
				return nil, nil
			},
		},
		UploadConnectionResolverFunc: &ResolverUploadConnectionResolverFunc{
			defaultHook: func(dbstore.GetUploadsOptions) *resolvers.UploadsResolver {
				return nil
			},
		},
	}
}

// NewMockResolverFrom creates a new mock of the MockResolver interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockResolverFrom(i resolvers.Resolver) *MockResolver {
	return &MockResolver{
		DeleteIndexByIDFunc: &ResolverDeleteIndexByIDFunc{
			defaultHook: i.DeleteIndexByID,
		},
		DeleteUploadByIDFunc: &ResolverDeleteUploadByIDFunc{
			defaultHook: i.DeleteUploadByID,
		},
		GetIndexByIDFunc: &ResolverGetIndexByIDFunc{
			defaultHook: i.GetIndexByID,
		},
		GetUploadByIDFunc: &ResolverGetUploadByIDFunc{
			defaultHook: i.GetUploadByID,
		},
		IndexConnectionResolverFunc: &ResolverIndexConnectionResolverFunc{
			defaultHook: i.IndexConnectionResolver,
		},
		QueryResolverFunc: &ResolverQueryResolverFunc{
			defaultHook: i.QueryResolver,
		},
		UploadConnectionResolverFunc: &ResolverUploadConnectionResolverFunc{
			defaultHook: i.UploadConnectionResolver,
		},
	}
}

// ResolverDeleteIndexByIDFunc describes the behavior when the
// DeleteIndexByID method of the parent MockResolver instance is invoked.
type ResolverDeleteIndexByIDFunc struct {
	defaultHook func(context.Context, int) error
	hooks       []func(context.Context, int) error
	history     []ResolverDeleteIndexByIDFuncCall
	mutex       sync.Mutex
}

// DeleteIndexByID delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockResolver) DeleteIndexByID(v0 context.Context, v1 int) error {
	r0 := m.DeleteIndexByIDFunc.nextHook()(v0, v1)
	m.DeleteIndexByIDFunc.appendCall(ResolverDeleteIndexByIDFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the DeleteIndexByID
// method of the parent MockResolver instance is invoked and the hook queue
// is empty.
func (f *ResolverDeleteIndexByIDFunc) SetDefaultHook(hook func(context.Context, int) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DeleteIndexByID method of the parent MockResolver instance inovkes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *ResolverDeleteIndexByIDFunc) PushHook(hook func(context.Context, int) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ResolverDeleteIndexByIDFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, int) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ResolverDeleteIndexByIDFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, int) error {
		return r0
	})
}

func (f *ResolverDeleteIndexByIDFunc) nextHook() func(context.Context, int) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ResolverDeleteIndexByIDFunc) appendCall(r0 ResolverDeleteIndexByIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ResolverDeleteIndexByIDFuncCall objects
// describing the invocations of this function.
func (f *ResolverDeleteIndexByIDFunc) History() []ResolverDeleteIndexByIDFuncCall {
	f.mutex.Lock()
	history := make([]ResolverDeleteIndexByIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ResolverDeleteIndexByIDFuncCall is an object that describes an invocation
// of method DeleteIndexByID on an instance of MockResolver.
type ResolverDeleteIndexByIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ResolverDeleteIndexByIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ResolverDeleteIndexByIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ResolverDeleteUploadByIDFunc describes the behavior when the
// DeleteUploadByID method of the parent MockResolver instance is invoked.
type ResolverDeleteUploadByIDFunc struct {
	defaultHook func(context.Context, int) error
	hooks       []func(context.Context, int) error
	history     []ResolverDeleteUploadByIDFuncCall
	mutex       sync.Mutex
}

// DeleteUploadByID delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockResolver) DeleteUploadByID(v0 context.Context, v1 int) error {
	r0 := m.DeleteUploadByIDFunc.nextHook()(v0, v1)
	m.DeleteUploadByIDFunc.appendCall(ResolverDeleteUploadByIDFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the DeleteUploadByID
// method of the parent MockResolver instance is invoked and the hook queue
// is empty.
func (f *ResolverDeleteUploadByIDFunc) SetDefaultHook(hook func(context.Context, int) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// DeleteUploadByID method of the parent MockResolver instance inovkes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *ResolverDeleteUploadByIDFunc) PushHook(hook func(context.Context, int) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ResolverDeleteUploadByIDFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, int) error {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ResolverDeleteUploadByIDFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, int) error {
		return r0
	})
}

func (f *ResolverDeleteUploadByIDFunc) nextHook() func(context.Context, int) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ResolverDeleteUploadByIDFunc) appendCall(r0 ResolverDeleteUploadByIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ResolverDeleteUploadByIDFuncCall objects
// describing the invocations of this function.
func (f *ResolverDeleteUploadByIDFunc) History() []ResolverDeleteUploadByIDFuncCall {
	f.mutex.Lock()
	history := make([]ResolverDeleteUploadByIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ResolverDeleteUploadByIDFuncCall is an object that describes an
// invocation of method DeleteUploadByID on an instance of MockResolver.
type ResolverDeleteUploadByIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ResolverDeleteUploadByIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ResolverDeleteUploadByIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ResolverGetIndexByIDFunc describes the behavior when the GetIndexByID
// method of the parent MockResolver instance is invoked.
type ResolverGetIndexByIDFunc struct {
	defaultHook func(context.Context, int) (dbstore.Index, bool, error)
	hooks       []func(context.Context, int) (dbstore.Index, bool, error)
	history     []ResolverGetIndexByIDFuncCall
	mutex       sync.Mutex
}

// GetIndexByID delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockResolver) GetIndexByID(v0 context.Context, v1 int) (dbstore.Index, bool, error) {
	r0, r1, r2 := m.GetIndexByIDFunc.nextHook()(v0, v1)
	m.GetIndexByIDFunc.appendCall(ResolverGetIndexByIDFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the GetIndexByID method
// of the parent MockResolver instance is invoked and the hook queue is
// empty.
func (f *ResolverGetIndexByIDFunc) SetDefaultHook(hook func(context.Context, int) (dbstore.Index, bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetIndexByID method of the parent MockResolver instance inovkes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ResolverGetIndexByIDFunc) PushHook(hook func(context.Context, int) (dbstore.Index, bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ResolverGetIndexByIDFunc) SetDefaultReturn(r0 dbstore.Index, r1 bool, r2 error) {
	f.SetDefaultHook(func(context.Context, int) (dbstore.Index, bool, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ResolverGetIndexByIDFunc) PushReturn(r0 dbstore.Index, r1 bool, r2 error) {
	f.PushHook(func(context.Context, int) (dbstore.Index, bool, error) {
		return r0, r1, r2
	})
}

func (f *ResolverGetIndexByIDFunc) nextHook() func(context.Context, int) (dbstore.Index, bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ResolverGetIndexByIDFunc) appendCall(r0 ResolverGetIndexByIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ResolverGetIndexByIDFuncCall objects
// describing the invocations of this function.
func (f *ResolverGetIndexByIDFunc) History() []ResolverGetIndexByIDFuncCall {
	f.mutex.Lock()
	history := make([]ResolverGetIndexByIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ResolverGetIndexByIDFuncCall is an object that describes an invocation of
// method GetIndexByID on an instance of MockResolver.
type ResolverGetIndexByIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 dbstore.Index
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ResolverGetIndexByIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ResolverGetIndexByIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// ResolverGetUploadByIDFunc describes the behavior when the GetUploadByID
// method of the parent MockResolver instance is invoked.
type ResolverGetUploadByIDFunc struct {
	defaultHook func(context.Context, int) (dbstore.Upload, bool, error)
	hooks       []func(context.Context, int) (dbstore.Upload, bool, error)
	history     []ResolverGetUploadByIDFuncCall
	mutex       sync.Mutex
}

// GetUploadByID delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockResolver) GetUploadByID(v0 context.Context, v1 int) (dbstore.Upload, bool, error) {
	r0, r1, r2 := m.GetUploadByIDFunc.nextHook()(v0, v1)
	m.GetUploadByIDFunc.appendCall(ResolverGetUploadByIDFuncCall{v0, v1, r0, r1, r2})
	return r0, r1, r2
}

// SetDefaultHook sets function that is called when the GetUploadByID method
// of the parent MockResolver instance is invoked and the hook queue is
// empty.
func (f *ResolverGetUploadByIDFunc) SetDefaultHook(hook func(context.Context, int) (dbstore.Upload, bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetUploadByID method of the parent MockResolver instance inovkes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ResolverGetUploadByIDFunc) PushHook(hook func(context.Context, int) (dbstore.Upload, bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ResolverGetUploadByIDFunc) SetDefaultReturn(r0 dbstore.Upload, r1 bool, r2 error) {
	f.SetDefaultHook(func(context.Context, int) (dbstore.Upload, bool, error) {
		return r0, r1, r2
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ResolverGetUploadByIDFunc) PushReturn(r0 dbstore.Upload, r1 bool, r2 error) {
	f.PushHook(func(context.Context, int) (dbstore.Upload, bool, error) {
		return r0, r1, r2
	})
}

func (f *ResolverGetUploadByIDFunc) nextHook() func(context.Context, int) (dbstore.Upload, bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ResolverGetUploadByIDFunc) appendCall(r0 ResolverGetUploadByIDFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ResolverGetUploadByIDFuncCall objects
// describing the invocations of this function.
func (f *ResolverGetUploadByIDFunc) History() []ResolverGetUploadByIDFuncCall {
	f.mutex.Lock()
	history := make([]ResolverGetUploadByIDFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ResolverGetUploadByIDFuncCall is an object that describes an invocation
// of method GetUploadByID on an instance of MockResolver.
type ResolverGetUploadByIDFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 dbstore.Upload
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 bool
	// Result2 is the value of the 3rd result returned from this method
	// invocation.
	Result2 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ResolverGetUploadByIDFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ResolverGetUploadByIDFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1, c.Result2}
}

// ResolverIndexConnectionResolverFunc describes the behavior when the
// IndexConnectionResolver method of the parent MockResolver instance is
// invoked.
type ResolverIndexConnectionResolverFunc struct {
	defaultHook func(dbstore.GetIndexesOptions) *resolvers.IndexesResolver
	hooks       []func(dbstore.GetIndexesOptions) *resolvers.IndexesResolver
	history     []ResolverIndexConnectionResolverFuncCall
	mutex       sync.Mutex
}

// IndexConnectionResolver delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockResolver) IndexConnectionResolver(v0 dbstore.GetIndexesOptions) *resolvers.IndexesResolver {
	r0 := m.IndexConnectionResolverFunc.nextHook()(v0)
	m.IndexConnectionResolverFunc.appendCall(ResolverIndexConnectionResolverFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the
// IndexConnectionResolver method of the parent MockResolver instance is
// invoked and the hook queue is empty.
func (f *ResolverIndexConnectionResolverFunc) SetDefaultHook(hook func(dbstore.GetIndexesOptions) *resolvers.IndexesResolver) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// IndexConnectionResolver method of the parent MockResolver instance
// inovkes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *ResolverIndexConnectionResolverFunc) PushHook(hook func(dbstore.GetIndexesOptions) *resolvers.IndexesResolver) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ResolverIndexConnectionResolverFunc) SetDefaultReturn(r0 *resolvers.IndexesResolver) {
	f.SetDefaultHook(func(dbstore.GetIndexesOptions) *resolvers.IndexesResolver {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ResolverIndexConnectionResolverFunc) PushReturn(r0 *resolvers.IndexesResolver) {
	f.PushHook(func(dbstore.GetIndexesOptions) *resolvers.IndexesResolver {
		return r0
	})
}

func (f *ResolverIndexConnectionResolverFunc) nextHook() func(dbstore.GetIndexesOptions) *resolvers.IndexesResolver {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ResolverIndexConnectionResolverFunc) appendCall(r0 ResolverIndexConnectionResolverFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ResolverIndexConnectionResolverFuncCall
// objects describing the invocations of this function.
func (f *ResolverIndexConnectionResolverFunc) History() []ResolverIndexConnectionResolverFuncCall {
	f.mutex.Lock()
	history := make([]ResolverIndexConnectionResolverFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ResolverIndexConnectionResolverFuncCall is an object that describes an
// invocation of method IndexConnectionResolver on an instance of
// MockResolver.
type ResolverIndexConnectionResolverFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 dbstore.GetIndexesOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *resolvers.IndexesResolver
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ResolverIndexConnectionResolverFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ResolverIndexConnectionResolverFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// ResolverQueryResolverFunc describes the behavior when the QueryResolver
// method of the parent MockResolver instance is invoked.
type ResolverQueryResolverFunc struct {
	defaultHook func(context.Context, *graphqlbackend.GitBlobLSIFDataArgs) (resolvers.QueryResolver, error)
	hooks       []func(context.Context, *graphqlbackend.GitBlobLSIFDataArgs) (resolvers.QueryResolver, error)
	history     []ResolverQueryResolverFuncCall
	mutex       sync.Mutex
}

// QueryResolver delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockResolver) QueryResolver(v0 context.Context, v1 *graphqlbackend.GitBlobLSIFDataArgs) (resolvers.QueryResolver, error) {
	r0, r1 := m.QueryResolverFunc.nextHook()(v0, v1)
	m.QueryResolverFunc.appendCall(ResolverQueryResolverFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the QueryResolver method
// of the parent MockResolver instance is invoked and the hook queue is
// empty.
func (f *ResolverQueryResolverFunc) SetDefaultHook(hook func(context.Context, *graphqlbackend.GitBlobLSIFDataArgs) (resolvers.QueryResolver, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// QueryResolver method of the parent MockResolver instance inovkes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *ResolverQueryResolverFunc) PushHook(hook func(context.Context, *graphqlbackend.GitBlobLSIFDataArgs) (resolvers.QueryResolver, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ResolverQueryResolverFunc) SetDefaultReturn(r0 resolvers.QueryResolver, r1 error) {
	f.SetDefaultHook(func(context.Context, *graphqlbackend.GitBlobLSIFDataArgs) (resolvers.QueryResolver, error) {
		return r0, r1
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ResolverQueryResolverFunc) PushReturn(r0 resolvers.QueryResolver, r1 error) {
	f.PushHook(func(context.Context, *graphqlbackend.GitBlobLSIFDataArgs) (resolvers.QueryResolver, error) {
		return r0, r1
	})
}

func (f *ResolverQueryResolverFunc) nextHook() func(context.Context, *graphqlbackend.GitBlobLSIFDataArgs) (resolvers.QueryResolver, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ResolverQueryResolverFunc) appendCall(r0 ResolverQueryResolverFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ResolverQueryResolverFuncCall objects
// describing the invocations of this function.
func (f *ResolverQueryResolverFunc) History() []ResolverQueryResolverFuncCall {
	f.mutex.Lock()
	history := make([]ResolverQueryResolverFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ResolverQueryResolverFuncCall is an object that describes an invocation
// of method QueryResolver on an instance of MockResolver.
type ResolverQueryResolverFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 *graphqlbackend.GitBlobLSIFDataArgs
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 resolvers.QueryResolver
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ResolverQueryResolverFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ResolverQueryResolverFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// ResolverUploadConnectionResolverFunc describes the behavior when the
// UploadConnectionResolver method of the parent MockResolver instance is
// invoked.
type ResolverUploadConnectionResolverFunc struct {
	defaultHook func(dbstore.GetUploadsOptions) *resolvers.UploadsResolver
	hooks       []func(dbstore.GetUploadsOptions) *resolvers.UploadsResolver
	history     []ResolverUploadConnectionResolverFuncCall
	mutex       sync.Mutex
}

// UploadConnectionResolver delegates to the next hook function in the queue
// and stores the parameter and result values of this invocation.
func (m *MockResolver) UploadConnectionResolver(v0 dbstore.GetUploadsOptions) *resolvers.UploadsResolver {
	r0 := m.UploadConnectionResolverFunc.nextHook()(v0)
	m.UploadConnectionResolverFunc.appendCall(ResolverUploadConnectionResolverFuncCall{v0, r0})
	return r0
}

// SetDefaultHook sets function that is called when the
// UploadConnectionResolver method of the parent MockResolver instance is
// invoked and the hook queue is empty.
func (f *ResolverUploadConnectionResolverFunc) SetDefaultHook(hook func(dbstore.GetUploadsOptions) *resolvers.UploadsResolver) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// UploadConnectionResolver method of the parent MockResolver instance
// inovkes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *ResolverUploadConnectionResolverFunc) PushHook(hook func(dbstore.GetUploadsOptions) *resolvers.UploadsResolver) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultDefaultHook with a function that returns
// the given values.
func (f *ResolverUploadConnectionResolverFunc) SetDefaultReturn(r0 *resolvers.UploadsResolver) {
	f.SetDefaultHook(func(dbstore.GetUploadsOptions) *resolvers.UploadsResolver {
		return r0
	})
}

// PushReturn calls PushDefaultHook with a function that returns the given
// values.
func (f *ResolverUploadConnectionResolverFunc) PushReturn(r0 *resolvers.UploadsResolver) {
	f.PushHook(func(dbstore.GetUploadsOptions) *resolvers.UploadsResolver {
		return r0
	})
}

func (f *ResolverUploadConnectionResolverFunc) nextHook() func(dbstore.GetUploadsOptions) *resolvers.UploadsResolver {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *ResolverUploadConnectionResolverFunc) appendCall(r0 ResolverUploadConnectionResolverFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of ResolverUploadConnectionResolverFuncCall
// objects describing the invocations of this function.
func (f *ResolverUploadConnectionResolverFunc) History() []ResolverUploadConnectionResolverFuncCall {
	f.mutex.Lock()
	history := make([]ResolverUploadConnectionResolverFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// ResolverUploadConnectionResolverFuncCall is an object that describes an
// invocation of method UploadConnectionResolver on an instance of
// MockResolver.
type ResolverUploadConnectionResolverFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 dbstore.GetUploadsOptions
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 *resolvers.UploadsResolver
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c ResolverUploadConnectionResolverFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c ResolverUploadConnectionResolverFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
