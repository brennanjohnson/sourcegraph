import Dialog from '@reach/dialog'
import React from 'react'
import { CampaignsCodeHostFields, CampaignsCredentialFields } from '../../../graphql-operations'
import { CodeHostSshPublicKey } from './CodeHostSshPublicKey'
import { ModalHeader } from './ModalHeader'

interface ViewCredentialModalProps {
    codeHost: CampaignsCodeHostFields
    credential: CampaignsCredentialFields

    onClose: () => void
}

export const ViewCredentialModal: React.FunctionComponent<ViewCredentialModalProps> = ({
    credential,
    codeHost,
    onClose,
}) => {
    const labelId = 'viewCredential'
    return (
        <Dialog
            className="modal-body modal-body--top-third p-4 rounded border"
            onDismiss={onClose}
            aria-labelledby={labelId}
        >
            <div className="test-remove-credential-modal">
                <ModalHeader
                    id={labelId}
                    externalServiceKind={codeHost.externalServiceKind}
                    externalServiceURL={codeHost.externalServiceURL}
                />

                <h4>Personal access token</h4>
                <p>
                    <i>PATs cannot be viewed after entering.</i>
                </p>

                <hr className="mb-3" />

                <CodeHostSshPublicKey
                    externalServiceKind={codeHost.externalServiceKind}
                    sshPublicKey={credential.sshPublicKey!}
                />

                <div className="d-flex justify-content-end">
                    <button type="button" className="btn btn-outline-secondary" onClick={onClose}>
                        Close
                    </button>
                </div>
            </div>
        </Dialog>
    )
}
