import { CampaignState } from '../../shared/src/graphql/schema'
import { OrganizationInvitationResponseType } from '../../shared/src/graphql/schema'
import { EventSource } from '../../shared/src/graphql/schema'
import { RepositoryPermission } from '../../shared/src/graphql/schema'
import { ChangesetState } from '../../shared/src/graphql/schema'
import { ChangesetExternalState } from '../../shared/src/graphql/schema'
import { ChangesetReviewState } from '../../shared/src/graphql/schema'
import { ChangesetCheckState } from '../../shared/src/graphql/schema'
import { SymbolKind } from '../../shared/src/graphql/schema'
import { DiagnosticSeverity } from '../../shared/src/graphql/schema'
import { ExternalServiceKind } from '../../shared/src/graphql/schema'
import { GitRefType } from '../../shared/src/graphql/schema'
import { GitObjectType } from '../../shared/src/graphql/schema'
import { GitRefOrder } from '../../shared/src/graphql/schema'
import { DiffHunkLineType } from '../../shared/src/graphql/schema'
import { LSIFUploadState } from '../../shared/src/graphql/schema'
import { LSIFIndexState } from '../../shared/src/graphql/schema'
import { RepositoryOrderBy } from '../../shared/src/graphql/schema'
import { UserActivePeriod } from '../../shared/src/graphql/schema'
import { SearchVersion } from '../../shared/src/graphql/schema'
import { SearchPatternType } from '../../shared/src/graphql/schema'
import { AlertType } from '../../shared/src/graphql/schema'
import { UserEvent } from '../../shared/src/graphql/schema'
import { ChangesetSpecType } from '../../shared/src/graphql/schema'
export type Maybe<T> = T | null
export type Exact<T extends { [key: string]: any }> = { [K in keyof T]: T[K] }
/* This is an autogenerated file. Do not edit this file directly! */

export interface BrowserGraphQlOperations {
    /** browser/src/shared/backend/diffs.tsx */
    RepositoryComparisonDiff: (variables: RepositoryComparisonDiffVariables) => RepositoryComparisonDiffResult

    /** browser/src/shared/backend/search.tsx */
    SearchSuggestions: (variables: SearchSuggestionsVariables) => SearchSuggestionsResult

    /** browser/src/shared/backend/server.ts */
    SiteProductVersion: (variables: SiteProductVersionVariables) => SiteProductVersionResult

    /** browser/src/shared/backend/userEvents.tsx */
    logUserEvent: (variables: logUserEventVariables) => logUserEventResult

    /** browser/src/shared/backend/userEvents.tsx */
    logEvent: (variables: logEventVariables) => logEventResult

    /** browser/src/shared/code-hosts/phabricator/backend.tsx */
    addPhabricatorRepo: (variables: addPhabricatorRepoVariables) => addPhabricatorRepoResult

    /** browser/src/shared/code-hosts/phabricator/backend.tsx */
    ResolveStagingRev: (variables: ResolveStagingRevVariables) => ResolveStagingRevResult

    /** browser/src/shared/platform/context.test.tsx */
    ResolveRepoTest: (variables: ResolveRepoTestVariables) => ResolveRepoTestResult

    /** browser/src/shared/repo/backend.tsx */
    ResolveRepo: (variables: ResolveRepoVariables) => ResolveRepoResult

    /** browser/src/shared/repo/backend.tsx */
    ResolveRev: (variables: ResolveRevVariables) => ResolveRevResult

    /** browser/src/shared/repo/backend.tsx */
    BlobContent: (variables: BlobContentVariables) => BlobContentResult
}
/** All built-in and custom scalars, mapped to their actual values */
export interface Scalars {
    ID: string
    String: string
    Boolean: boolean
    Int: number
    Float: number
    DateTime: string
    JSONCString: string
    JSONValue: unknown
    GitObjectID: string
}

export interface AddExternalServiceInput {
    kind: ExternalServiceKind
    displayName: Scalars['String']
    config: Scalars['String']
}

export { AlertType }

export { CampaignState }

export { ChangesetCheckState }

export { ChangesetExternalState }

export { ChangesetReviewState }

export { ChangesetSpecType }

export { ChangesetState }

export interface ConfigurationEdit {
    keyPath: Array<KeyPathSegment>
    value?: Maybe<Scalars['JSONValue']>
    valueIsJSONCEncodedString?: Maybe<Scalars['Boolean']>
}

export { DiagnosticSeverity }

export { DiffHunkLineType }

export { EventSource }

export { ExternalServiceKind }

export { GitObjectType }

export { GitRefOrder }

export { GitRefType }

export interface KeyPathSegment {
    property?: Maybe<Scalars['String']>
    index?: Maybe<Scalars['Int']>
}

export { LSIFIndexState }

export { LSIFUploadState }

export interface MarkdownOptions {
    alwaysNil?: Maybe<Scalars['String']>
}

export { OrganizationInvitationResponseType }

export interface ProductLicenseInput {
    tags: Array<Scalars['String']>
    userCount: Scalars['Int']
    expiresAt: Scalars['Int']
}

export interface ProductSubscriptionInput {
    billingPlanID: Scalars['String']
    userCount: Scalars['Int']
}

export { RepositoryOrderBy }

export { RepositoryPermission }

export { SearchPatternType }

export { SearchVersion }

export interface SettingsEdit {
    keyPath: Array<KeyPathSegment>
    value?: Maybe<Scalars['JSONValue']>
    valueIsJSONCEncodedString?: Maybe<Scalars['Boolean']>
}

export interface SettingsMutationGroupInput {
    subject: Scalars['ID']
    lastID?: Maybe<Scalars['Int']>
}

export interface SurveySubmissionInput {
    email?: Maybe<Scalars['String']>
    score: Scalars['Int']
    reason?: Maybe<Scalars['String']>
    better?: Maybe<Scalars['String']>
}

export { SymbolKind }

export interface UpdateExternalServiceInput {
    id: Scalars['ID']
    displayName?: Maybe<Scalars['String']>
    config?: Maybe<Scalars['String']>
}

export { UserActivePeriod }

export { UserEvent }

export interface UserPermission {
    bindID: Scalars['String']
    permission?: Maybe<RepositoryPermission>
}

export type RepositoryComparisonDiffVariables = Exact<{
    repo: Scalars['String']
    base: Maybe<Scalars['String']>
    head: Maybe<Scalars['String']>
    first: Maybe<Scalars['Int']>
}>

export type RepositoryComparisonDiffResult = {
    repository: Maybe<{ comparison: { fileDiffs: { totalCount: Maybe<number>; nodes: Array<FileDiffFields> } } }>
}

export type FileDiffFields = { oldPath: Maybe<string>; newPath: Maybe<string>; internalID: string }

export type SymbolFields = {
    __typename: 'Symbol'
    name: string
    containerName: Maybe<string>
    url: string
    kind: SymbolKind
    location: { url: string; resource: { path: string; repository: { name: string } } }
}

export type SearchSuggestionsVariables = Exact<{
    query: Scalars['String']
    first: Scalars['Int']
}>

export type SearchSuggestionsResult = {
    search: Maybe<{
        suggestions: Array<
            | { __typename: 'Repository'; name: string; url: string }
            | {
                  __typename: 'File'
                  path: string
                  name: string
                  isDirectory: boolean
                  url: string
                  repository: { name: string }
              }
            | SymbolFields
        >
    }>
}

export type SiteProductVersionVariables = Exact<{ [key: string]: never }>

export type SiteProductVersionResult = {
    site: { productVersion: string; buildVersion: string; hasCodeIntelligence: boolean }
}

export type logUserEventVariables = Exact<{
    event: UserEvent
    userCookieID: Scalars['String']
}>

export type logUserEventResult = { logUserEvent: Maybe<{ alwaysNil: Maybe<string> }> }

export type logEventVariables = Exact<{
    name: Scalars['String']
    userCookieID: Scalars['String']
    url: Scalars['String']
    source: EventSource
    argument: Maybe<Scalars['String']>
}>

export type logEventResult = { logEvent: Maybe<{ alwaysNil: Maybe<string> }> }

export type addPhabricatorRepoVariables = Exact<{
    callsign: Scalars['String']
    repoName: Scalars['String']
    phabricatorURL: Scalars['String']
}>

export type addPhabricatorRepoResult = { addPhabricatorRepo: Maybe<{ alwaysNil: Maybe<string> }> }

export type ResolveStagingRevVariables = Exact<{
    repoName: Scalars['String']
    diffID: Scalars['ID']
    baseRev: Scalars['String']
    patch: Maybe<Scalars['String']>
    date: Maybe<Scalars['String']>
    authorName: Maybe<Scalars['String']>
    authorEmail: Maybe<Scalars['String']>
    description: Maybe<Scalars['String']>
}>

export type ResolveStagingRevResult = { resolvePhabricatorDiff: Maybe<{ oid: string }> }

export type ResolveRepoTestVariables = Exact<{
    repoName: Scalars['String']
}>

export type ResolveRepoTestResult = { repository: Maybe<{ url: string }> }

export type ResolveRepoVariables = Exact<{
    rawRepoName: Scalars['String']
}>

export type ResolveRepoResult = { repository: Maybe<{ name: string }> }

export type ResolveRevVariables = Exact<{
    repoName: Scalars['String']
    revision: Scalars['String']
}>

export type ResolveRevResult = {
    repository: Maybe<{ mirrorInfo: { cloned: boolean }; commit: Maybe<{ oid: string }> }>
}

export type BlobContentVariables = Exact<{
    repoName: Scalars['String']
    commitID: Scalars['String']
    filePath: Scalars['String']
}>

export type BlobContentResult = {
    repository: Maybe<{ commit: Maybe<{ file: Maybe<{ content: string } | { content: string }> }> }>
}
