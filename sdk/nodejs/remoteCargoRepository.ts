// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * ## # Artifactory Remote Cargo Repository Resource
 *
 * Provides an Artifactory remote `cargo` repository resource. This provides cargo specific fields and is the only way to get them
 * Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Cargo+Registry)
 *
 * ## Example Usage
 *
 * Create a new Artifactory remote cargo repository called my-remote-cargo
 * for brevity sake, only cargo specific fields are included; for other fields see documentation for
 * generic repo.
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const my_remote_cargo = new artifactory.RemoteCargoRepository("my-remote-cargo", {
 *     anonymousAccess: true,
 *     gitRegistryUrl: "https://github.com/rust-lang/foo.index",
 *     key: "my-remote-cargo",
 * });
 * ```
 * ## Note
 *
 * If you get a 400 error: `"Custom Base URL should be defined prior to creating a Cargo repository"`,
 * you must set the base url at: `http://${host}/ui/admin/configuration/general`
 */
export class RemoteCargoRepository extends pulumi.CustomResource {
    /**
     * Get an existing RemoteCargoRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RemoteCargoRepositoryState, opts?: pulumi.CustomResourceOptions): RemoteCargoRepository {
        return new RemoteCargoRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/remoteCargoRepository:RemoteCargoRepository';

    /**
     * Returns true if the given object is an instance of RemoteCargoRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RemoteCargoRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RemoteCargoRepository.__pulumiType;
    }

    /**
     * Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to any other host.
     */
    public readonly allowAnyHostAuth!: pulumi.Output<boolean>;
    /**
     * - Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option.
     */
    public readonly anonymousAccess!: pulumi.Output<boolean | undefined>;
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time, an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed offline. Default to 300.
     */
    public readonly assumedOfflinePeriodSecs!: pulumi.Output<number | undefined>;
    /**
     * (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact resolution.
     */
    public readonly blackedOut!: pulumi.Output<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources, HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked, Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    public readonly blockMismatchingMimeTypes!: pulumi.Output<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources, HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked, Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    public readonly bypassHeadRequests!: pulumi.Output<boolean>;
    public readonly clientTlsCertificate!: pulumi.Output<string>;
    /**
     * Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
     */
    public readonly contentSynchronisation!: pulumi.Output<outputs.RemoteCargoRepositoryContentSynchronisation>;
    public readonly description!: pulumi.Output<string>;
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     */
    public readonly enableCookieManagement!: pulumi.Output<boolean>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*. By default no artifacts are excluded.
     */
    public readonly excludesPattern!: pulumi.Output<string>;
    /**
     * This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
     *
     * @deprecated This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
     */
    public /*out*/ readonly failedRetrievalCachePeriodSecs!: pulumi.Output<number>;
    /**
     * - This is the index url, expected to be a git repository. for remote artifactory use "arturl/git/repokey.git"
     */
    public readonly gitRegistryUrl!: pulumi.Output<string>;
    /**
     * When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to communicate with this repository.
     */
    public readonly hardFail!: pulumi.Output<boolean>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    public readonly includesPattern!: pulumi.Output<string>;
    /**
     * The repository identifier. Must be unique system-wide
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * - Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of the 'Retrieval Cache Period'. This field exists in the API but not in the UI.
     */
    public readonly listRemoteFolderItems!: pulumi.Output<boolean | undefined>;
    /**
     * The local address to be used when creating connections. Useful for specifying the interface to use on systems with multiple network interfaces.
     */
    public readonly localAddress!: pulumi.Output<string | undefined>;
    /**
     * The number of seconds to cache artifact retrieval misses (artifact not found). A value of 0 indicates no caching.
     */
    public readonly missedCachePeriodSeconds!: pulumi.Output<number>;
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     */
    public readonly offline!: pulumi.Output<boolean>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    public readonly priorityResolution!: pulumi.Output<boolean>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    public readonly projectEnvironments!: pulumi.Output<string[] | undefined>;
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    /**
     * When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     */
    public readonly propagateQueryParams!: pulumi.Output<boolean | undefined>;
    /**
     * List of property set name
     */
    public readonly propertySets!: pulumi.Output<string[] | undefined>;
    public readonly proxy!: pulumi.Output<string>;
    /**
     * Repository layout key for the remote layout mapping
     */
    public readonly remoteRepoLayoutRef!: pulumi.Output<string>;
    /**
     * Repository layout key for the remote repository
     */
    public readonly repoLayoutRef!: pulumi.Output<string>;
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     */
    public readonly retrievalCachePeriodSeconds!: pulumi.Output<number>;
    public readonly shareConfiguration!: pulumi.Output<boolean>;
    /**
     * Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network operation is considered a retrieval failure.
     */
    public readonly socketTimeoutMillis!: pulumi.Output<number>;
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory servers.
     */
    public readonly storeArtifactsLocally!: pulumi.Output<boolean>;
    /**
     * When set, remote artifacts are fetched along with their properties.
     */
    public readonly synchronizeProperties!: pulumi.Output<boolean>;
    public readonly unusedArtifactsCleanupPeriodEnabled!: pulumi.Output<boolean>;
    /**
     * The number of hours to wait before an artifact is deemed "unused" and eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     */
    public readonly unusedArtifactsCleanupPeriodHours!: pulumi.Output<number>;
    /**
     * the remote repo URL. You kinda don't have a remote repo without it
     */
    public readonly url!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string | undefined>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
     */
    public readonly xrayIndex!: pulumi.Output<boolean | undefined>;

    /**
     * Create a RemoteCargoRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemoteCargoRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RemoteCargoRepositoryArgs | RemoteCargoRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RemoteCargoRepositoryState | undefined;
            resourceInputs["allowAnyHostAuth"] = state ? state.allowAnyHostAuth : undefined;
            resourceInputs["anonymousAccess"] = state ? state.anonymousAccess : undefined;
            resourceInputs["assumedOfflinePeriodSecs"] = state ? state.assumedOfflinePeriodSecs : undefined;
            resourceInputs["blackedOut"] = state ? state.blackedOut : undefined;
            resourceInputs["blockMismatchingMimeTypes"] = state ? state.blockMismatchingMimeTypes : undefined;
            resourceInputs["bypassHeadRequests"] = state ? state.bypassHeadRequests : undefined;
            resourceInputs["clientTlsCertificate"] = state ? state.clientTlsCertificate : undefined;
            resourceInputs["contentSynchronisation"] = state ? state.contentSynchronisation : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enableCookieManagement"] = state ? state.enableCookieManagement : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["failedRetrievalCachePeriodSecs"] = state ? state.failedRetrievalCachePeriodSecs : undefined;
            resourceInputs["gitRegistryUrl"] = state ? state.gitRegistryUrl : undefined;
            resourceInputs["hardFail"] = state ? state.hardFail : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["listRemoteFolderItems"] = state ? state.listRemoteFolderItems : undefined;
            resourceInputs["localAddress"] = state ? state.localAddress : undefined;
            resourceInputs["missedCachePeriodSeconds"] = state ? state.missedCachePeriodSeconds : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["offline"] = state ? state.offline : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["priorityResolution"] = state ? state.priorityResolution : undefined;
            resourceInputs["projectEnvironments"] = state ? state.projectEnvironments : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["propagateQueryParams"] = state ? state.propagateQueryParams : undefined;
            resourceInputs["propertySets"] = state ? state.propertySets : undefined;
            resourceInputs["proxy"] = state ? state.proxy : undefined;
            resourceInputs["remoteRepoLayoutRef"] = state ? state.remoteRepoLayoutRef : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = state ? state.retrievalCachePeriodSeconds : undefined;
            resourceInputs["shareConfiguration"] = state ? state.shareConfiguration : undefined;
            resourceInputs["socketTimeoutMillis"] = state ? state.socketTimeoutMillis : undefined;
            resourceInputs["storeArtifactsLocally"] = state ? state.storeArtifactsLocally : undefined;
            resourceInputs["synchronizeProperties"] = state ? state.synchronizeProperties : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodEnabled"] = state ? state.unusedArtifactsCleanupPeriodEnabled : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodHours"] = state ? state.unusedArtifactsCleanupPeriodHours : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["xrayIndex"] = state ? state.xrayIndex : undefined;
        } else {
            const args = argsOrState as RemoteCargoRepositoryArgs | undefined;
            if ((!args || args.gitRegistryUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gitRegistryUrl'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["allowAnyHostAuth"] = args ? args.allowAnyHostAuth : undefined;
            resourceInputs["anonymousAccess"] = args ? args.anonymousAccess : undefined;
            resourceInputs["assumedOfflinePeriodSecs"] = args ? args.assumedOfflinePeriodSecs : undefined;
            resourceInputs["blackedOut"] = args ? args.blackedOut : undefined;
            resourceInputs["blockMismatchingMimeTypes"] = args ? args.blockMismatchingMimeTypes : undefined;
            resourceInputs["bypassHeadRequests"] = args ? args.bypassHeadRequests : undefined;
            resourceInputs["clientTlsCertificate"] = args ? args.clientTlsCertificate : undefined;
            resourceInputs["contentSynchronisation"] = args ? args.contentSynchronisation : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableCookieManagement"] = args ? args.enableCookieManagement : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["gitRegistryUrl"] = args ? args.gitRegistryUrl : undefined;
            resourceInputs["hardFail"] = args ? args.hardFail : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["listRemoteFolderItems"] = args ? args.listRemoteFolderItems : undefined;
            resourceInputs["localAddress"] = args ? args.localAddress : undefined;
            resourceInputs["missedCachePeriodSeconds"] = args ? args.missedCachePeriodSeconds : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["offline"] = args ? args.offline : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["priorityResolution"] = args ? args.priorityResolution : undefined;
            resourceInputs["projectEnvironments"] = args ? args.projectEnvironments : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["propagateQueryParams"] = args ? args.propagateQueryParams : undefined;
            resourceInputs["propertySets"] = args ? args.propertySets : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["remoteRepoLayoutRef"] = args ? args.remoteRepoLayoutRef : undefined;
            resourceInputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = args ? args.retrievalCachePeriodSeconds : undefined;
            resourceInputs["shareConfiguration"] = args ? args.shareConfiguration : undefined;
            resourceInputs["socketTimeoutMillis"] = args ? args.socketTimeoutMillis : undefined;
            resourceInputs["storeArtifactsLocally"] = args ? args.storeArtifactsLocally : undefined;
            resourceInputs["synchronizeProperties"] = args ? args.synchronizeProperties : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodEnabled"] = args ? args.unusedArtifactsCleanupPeriodEnabled : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodHours"] = args ? args.unusedArtifactsCleanupPeriodHours : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["xrayIndex"] = args ? args.xrayIndex : undefined;
            resourceInputs["failedRetrievalCachePeriodSecs"] = undefined /*out*/;
            resourceInputs["packageType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RemoteCargoRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RemoteCargoRepository resources.
 */
export interface RemoteCargoRepositoryState {
    /**
     * Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to any other host.
     */
    allowAnyHostAuth?: pulumi.Input<boolean>;
    /**
     * - Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option.
     */
    anonymousAccess?: pulumi.Input<boolean>;
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time, an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed offline. Default to 300.
     */
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    /**
     * (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact resolution.
     */
    blackedOut?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources, HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked, Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources, HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked, Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    bypassHeadRequests?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    /**
     * Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
     */
    contentSynchronisation?: pulumi.Input<inputs.RemoteCargoRepositoryContentSynchronisation>;
    description?: pulumi.Input<string>;
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     */
    enableCookieManagement?: pulumi.Input<boolean>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*. By default no artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
     *
     * @deprecated This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
     */
    failedRetrievalCachePeriodSecs?: pulumi.Input<number>;
    /**
     * - This is the index url, expected to be a git repository. for remote artifactory use "arturl/git/repokey.git"
     */
    gitRegistryUrl?: pulumi.Input<string>;
    /**
     * When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to communicate with this repository.
     */
    hardFail?: pulumi.Input<boolean>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * The repository identifier. Must be unique system-wide
     */
    key?: pulumi.Input<string>;
    /**
     * - Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of the 'Retrieval Cache Period'. This field exists in the API but not in the UI.
     */
    listRemoteFolderItems?: pulumi.Input<boolean>;
    /**
     * The local address to be used when creating connections. Useful for specifying the interface to use on systems with multiple network interfaces.
     */
    localAddress?: pulumi.Input<string>;
    /**
     * The number of seconds to cache artifact retrieval misses (artifact not found). A value of 0 indicates no caching.
     */
    missedCachePeriodSeconds?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     */
    offline?: pulumi.Input<boolean>;
    packageType?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     */
    propagateQueryParams?: pulumi.Input<boolean>;
    /**
     * List of property set name
     */
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    proxy?: pulumi.Input<string>;
    /**
     * Repository layout key for the remote layout mapping
     */
    remoteRepoLayoutRef?: pulumi.Input<string>;
    /**
     * Repository layout key for the remote repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    /**
     * Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network operation is considered a retrieval failure.
     */
    socketTimeoutMillis?: pulumi.Input<number>;
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory servers.
     */
    storeArtifactsLocally?: pulumi.Input<boolean>;
    /**
     * When set, remote artifacts are fetched along with their properties.
     */
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodEnabled?: pulumi.Input<boolean>;
    /**
     * The number of hours to wait before an artifact is deemed "unused" and eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     */
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    /**
     * the remote repo URL. You kinda don't have a remote repo without it
     */
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
     */
    xrayIndex?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RemoteCargoRepository resource.
 */
export interface RemoteCargoRepositoryArgs {
    /**
     * Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to any other host.
     */
    allowAnyHostAuth?: pulumi.Input<boolean>;
    /**
     * - Cargo client does not send credentials when performing download and search for crates. Enable this to allow anonymous access to these resources (only), note that this will override the security anonymous access option.
     */
    anonymousAccess?: pulumi.Input<boolean>;
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time, an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed offline. Default to 300.
     */
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    /**
     * (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact resolution.
     */
    blackedOut?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources, HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked, Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources, HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked, Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    bypassHeadRequests?: pulumi.Input<boolean>;
    clientTlsCertificate?: pulumi.Input<string>;
    /**
     * Reference [JFROG Smart Remote Repositories](https://www.jfrog.com/confluence/display/JFROG/Smart+Remote+Repositories)
     */
    contentSynchronisation?: pulumi.Input<inputs.RemoteCargoRepositoryContentSynchronisation>;
    description?: pulumi.Input<string>;
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     */
    enableCookieManagement?: pulumi.Input<boolean>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*. By default no artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * - This is the index url, expected to be a git repository. for remote artifactory use "arturl/git/repokey.git"
     */
    gitRegistryUrl: pulumi.Input<string>;
    /**
     * When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to communicate with this repository.
     */
    hardFail?: pulumi.Input<boolean>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * The repository identifier. Must be unique system-wide
     */
    key: pulumi.Input<string>;
    /**
     * - Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of the 'Retrieval Cache Period'. This field exists in the API but not in the UI.
     */
    listRemoteFolderItems?: pulumi.Input<boolean>;
    /**
     * The local address to be used when creating connections. Useful for specifying the interface to use on systems with multiple network interfaces.
     */
    localAddress?: pulumi.Input<string>;
    /**
     * The number of seconds to cache artifact retrieval misses (artifact not found). A value of 0 indicates no caching.
     */
    missedCachePeriodSeconds?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     */
    offline?: pulumi.Input<boolean>;
    password?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     */
    propagateQueryParams?: pulumi.Input<boolean>;
    /**
     * List of property set name
     */
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    proxy?: pulumi.Input<string>;
    /**
     * Repository layout key for the remote layout mapping
     */
    remoteRepoLayoutRef?: pulumi.Input<string>;
    /**
     * Repository layout key for the remote repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    /**
     * Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network operation is considered a retrieval failure.
     */
    socketTimeoutMillis?: pulumi.Input<number>;
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory servers.
     */
    storeArtifactsLocally?: pulumi.Input<boolean>;
    /**
     * When set, remote artifacts are fetched along with their properties.
     */
    synchronizeProperties?: pulumi.Input<boolean>;
    unusedArtifactsCleanupPeriodEnabled?: pulumi.Input<boolean>;
    /**
     * The number of hours to wait before an artifact is deemed "unused" and eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     */
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    /**
     * the remote repo URL. You kinda don't have a remote repo without it
     */
    url: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via Xray settings.
     */
    xrayIndex?: pulumi.Input<boolean>;
}
