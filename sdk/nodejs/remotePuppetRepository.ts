// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Creates a remote Puppet repository.
 * Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Puppet+Repositories).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * const my_remote_puppet = new artifactory.RemotePuppetRepository("my-remote-puppet", {
 *     key: "my-remote-puppet",
 *     url: "https://forgeapi.puppetlabs.com/",
 * });
 * ```
 *
 * ## Import
 *
 * Remote repositories can be imported using their name, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/remotePuppetRepository:RemotePuppetRepository my-remote-puppet my-remote-puppet
 * ```
 */
export class RemotePuppetRepository extends pulumi.CustomResource {
    /**
     * Get an existing RemotePuppetRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RemotePuppetRepositoryState, opts?: pulumi.CustomResourceOptions): RemotePuppetRepository {
        return new RemotePuppetRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/remotePuppetRepository:RemotePuppetRepository';

    /**
     * Returns true if the given object is an instance of RemotePuppetRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RemotePuppetRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RemotePuppetRepository.__pulumiType;
    }

    /**
     * 'Lenient Host Authentication' in the UI. Allow credentials of this repository to be used on requests redirected to any
     * other host.
     */
    public readonly allowAnyHostAuth!: pulumi.Output<boolean | undefined>;
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline.
     */
    public readonly assumedOfflinePeriodSecs!: pulumi.Output<number | undefined>;
    /**
     * (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     */
    public readonly blackedOut!: pulumi.Output<boolean | undefined>;
    /**
     * If set, artifacts will fail to download if a mismatch is detected between requested and received mimetype, according to
     * the list specified in the system properties file under blockedMismatchingMimeTypes. You can override by adding mimetypes
     * to the override list 'mismatching_mime_types_override_list'.
     */
    public readonly blockMismatchingMimeTypes!: pulumi.Output<boolean | undefined>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    public readonly bypassHeadRequests!: pulumi.Output<boolean | undefined>;
    /**
     * Client TLS certificate name.
     */
    public readonly clientTlsCertificate!: pulumi.Output<string>;
    public readonly contentSynchronisation!: pulumi.Output<outputs.RemotePuppetRepositoryContentSynchronisation>;
    /**
     * Public description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only. Default value is 'false'.
     */
    public readonly downloadDirect!: pulumi.Output<boolean | undefined>;
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     */
    public readonly enableCookieManagement!: pulumi.Output<boolean | undefined>;
    /**
     * List of comma-separated artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*. By
     * default no artifacts are excluded.
     */
    public readonly excludesPattern!: pulumi.Output<string | undefined>;
    /**
     * When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
     * communicate with this repository.
     */
    public readonly hardFail!: pulumi.Output<boolean | undefined>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    public readonly includesPattern!: pulumi.Output<string | undefined>;
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
     * the 'Retrieval Cache Period'. Default value is 'true'.
     */
    public readonly listRemoteFolderItems!: pulumi.Output<boolean | undefined>;
    /**
     * The local address to be used when creating connections. Useful for specifying the interface to use on systems with
     * multiple network interfaces.
     */
    public readonly localAddress!: pulumi.Output<string | undefined>;
    /**
     * Metadata Retrieval Cache Timeout (Sec) in the UI.This value refers to the number of seconds to wait for retrieval from
     * the remote before serving locally cached artifact or fail the request.
     */
    public readonly metadataRetrievalTimeoutSecs!: pulumi.Output<number | undefined>;
    /**
     * The set of mime types that should override the block_mismatching_mime_types setting. Eg:
     * 'application/json,application/xml'. Default value is empty.
     */
    public readonly mismatchingMimeTypesOverrideList!: pulumi.Output<string | undefined>;
    /**
     * Missed Retrieval Cache Period (Sec) in the UI. The number of seconds to cache artifact retrieval misses (artifact not
     * found). A value of 0 indicates no caching.
     */
    public readonly missedCachePeriodSeconds!: pulumi.Output<number | undefined>;
    /**
     * Internal description.
     */
    public readonly notes!: pulumi.Output<string | undefined>;
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     */
    public readonly offline!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Setting Priority Resolution takes precedence over the resolution order when resolving virtual repositories. Setting
     * repositories with priority will cause metadata to be merged only from repositories set with a priority. If a package is
     * not found in those repositories, Artifactory will merge from repositories marked as non-priority.
     */
    public readonly priorityResolution!: pulumi.Output<boolean | undefined>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
     * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
     * will remain in the Terraform state, which will create state drift during the update.
     */
    public readonly projectEnvironments!: pulumi.Output<string[]>;
    /**
     * Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    /**
     * When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     */
    public readonly propagateQueryParams!: pulumi.Output<boolean | undefined>;
    /**
     * List of property set names
     */
    public readonly propertySets!: pulumi.Output<string[] | undefined>;
    /**
     * Proxy key from Artifactory Proxies settings
     */
    public readonly proxy!: pulumi.Output<string | undefined>;
    /**
     * Custom HTTP query parameters that will be automatically included in all remote resource requests. For example:
     * `param1=val1&param2=val2&param3=val3`
     */
    public readonly queryParams!: pulumi.Output<string | undefined>;
    /**
     * Repository layout key for the remote layout mapping.
     */
    public readonly remoteRepoLayoutRef!: pulumi.Output<string | undefined>;
    /**
     * Repository layout key for the local repository
     */
    public readonly repoLayoutRef!: pulumi.Output<string | undefined>;
    /**
     * Metadata Retrieval Cache Period (Sec) in the UI. This value refers to the number of seconds to cache metadata files
     * before checking for newer versions on remote server. A value of 0 indicates no caching.
     */
    public readonly retrievalCachePeriodSeconds!: pulumi.Output<number | undefined>;
    public readonly shareConfiguration!: pulumi.Output<boolean>;
    /**
     * Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
     * operation is considered a retrieval failure.
     */
    public readonly socketTimeoutMillis!: pulumi.Output<number | undefined>;
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     */
    public readonly storeArtifactsLocally!: pulumi.Output<boolean | undefined>;
    /**
     * When set, remote artifacts are fetched along with their properties.
     */
    public readonly synchronizeProperties!: pulumi.Output<boolean | undefined>;
    /**
     * Unused Artifacts Cleanup Period (Hr) in the UI. The number of hours to wait before an artifact is deemed 'unused' and
     * eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     */
    public readonly unusedArtifactsCleanupPeriodHours!: pulumi.Output<number | undefined>;
    /**
     * The remote repo URL.
     */
    public readonly url!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string | undefined>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     */
    public readonly xrayIndex!: pulumi.Output<boolean | undefined>;

    /**
     * Create a RemotePuppetRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemotePuppetRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RemotePuppetRepositoryArgs | RemotePuppetRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RemotePuppetRepositoryState | undefined;
            resourceInputs["allowAnyHostAuth"] = state ? state.allowAnyHostAuth : undefined;
            resourceInputs["assumedOfflinePeriodSecs"] = state ? state.assumedOfflinePeriodSecs : undefined;
            resourceInputs["blackedOut"] = state ? state.blackedOut : undefined;
            resourceInputs["blockMismatchingMimeTypes"] = state ? state.blockMismatchingMimeTypes : undefined;
            resourceInputs["bypassHeadRequests"] = state ? state.bypassHeadRequests : undefined;
            resourceInputs["clientTlsCertificate"] = state ? state.clientTlsCertificate : undefined;
            resourceInputs["contentSynchronisation"] = state ? state.contentSynchronisation : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["downloadDirect"] = state ? state.downloadDirect : undefined;
            resourceInputs["enableCookieManagement"] = state ? state.enableCookieManagement : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["hardFail"] = state ? state.hardFail : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["listRemoteFolderItems"] = state ? state.listRemoteFolderItems : undefined;
            resourceInputs["localAddress"] = state ? state.localAddress : undefined;
            resourceInputs["metadataRetrievalTimeoutSecs"] = state ? state.metadataRetrievalTimeoutSecs : undefined;
            resourceInputs["mismatchingMimeTypesOverrideList"] = state ? state.mismatchingMimeTypesOverrideList : undefined;
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
            resourceInputs["queryParams"] = state ? state.queryParams : undefined;
            resourceInputs["remoteRepoLayoutRef"] = state ? state.remoteRepoLayoutRef : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = state ? state.retrievalCachePeriodSeconds : undefined;
            resourceInputs["shareConfiguration"] = state ? state.shareConfiguration : undefined;
            resourceInputs["socketTimeoutMillis"] = state ? state.socketTimeoutMillis : undefined;
            resourceInputs["storeArtifactsLocally"] = state ? state.storeArtifactsLocally : undefined;
            resourceInputs["synchronizeProperties"] = state ? state.synchronizeProperties : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodHours"] = state ? state.unusedArtifactsCleanupPeriodHours : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
            resourceInputs["xrayIndex"] = state ? state.xrayIndex : undefined;
        } else {
            const args = argsOrState as RemotePuppetRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["allowAnyHostAuth"] = args ? args.allowAnyHostAuth : undefined;
            resourceInputs["assumedOfflinePeriodSecs"] = args ? args.assumedOfflinePeriodSecs : undefined;
            resourceInputs["blackedOut"] = args ? args.blackedOut : undefined;
            resourceInputs["blockMismatchingMimeTypes"] = args ? args.blockMismatchingMimeTypes : undefined;
            resourceInputs["bypassHeadRequests"] = args ? args.bypassHeadRequests : undefined;
            resourceInputs["clientTlsCertificate"] = args ? args.clientTlsCertificate : undefined;
            resourceInputs["contentSynchronisation"] = args ? args.contentSynchronisation : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["downloadDirect"] = args ? args.downloadDirect : undefined;
            resourceInputs["enableCookieManagement"] = args ? args.enableCookieManagement : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["hardFail"] = args ? args.hardFail : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["listRemoteFolderItems"] = args ? args.listRemoteFolderItems : undefined;
            resourceInputs["localAddress"] = args ? args.localAddress : undefined;
            resourceInputs["metadataRetrievalTimeoutSecs"] = args ? args.metadataRetrievalTimeoutSecs : undefined;
            resourceInputs["mismatchingMimeTypesOverrideList"] = args ? args.mismatchingMimeTypesOverrideList : undefined;
            resourceInputs["missedCachePeriodSeconds"] = args ? args.missedCachePeriodSeconds : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["offline"] = args ? args.offline : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["priorityResolution"] = args ? args.priorityResolution : undefined;
            resourceInputs["projectEnvironments"] = args ? args.projectEnvironments : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["propagateQueryParams"] = args ? args.propagateQueryParams : undefined;
            resourceInputs["propertySets"] = args ? args.propertySets : undefined;
            resourceInputs["proxy"] = args ? args.proxy : undefined;
            resourceInputs["queryParams"] = args ? args.queryParams : undefined;
            resourceInputs["remoteRepoLayoutRef"] = args ? args.remoteRepoLayoutRef : undefined;
            resourceInputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            resourceInputs["retrievalCachePeriodSeconds"] = args ? args.retrievalCachePeriodSeconds : undefined;
            resourceInputs["shareConfiguration"] = args ? args.shareConfiguration : undefined;
            resourceInputs["socketTimeoutMillis"] = args ? args.socketTimeoutMillis : undefined;
            resourceInputs["storeArtifactsLocally"] = args ? args.storeArtifactsLocally : undefined;
            resourceInputs["synchronizeProperties"] = args ? args.synchronizeProperties : undefined;
            resourceInputs["unusedArtifactsCleanupPeriodHours"] = args ? args.unusedArtifactsCleanupPeriodHours : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["xrayIndex"] = args ? args.xrayIndex : undefined;
            resourceInputs["packageType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(RemotePuppetRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RemotePuppetRepository resources.
 */
export interface RemotePuppetRepositoryState {
    /**
     * 'Lenient Host Authentication' in the UI. Allow credentials of this repository to be used on requests redirected to any
     * other host.
     */
    allowAnyHostAuth?: pulumi.Input<boolean>;
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline.
     */
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    /**
     * (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     */
    blackedOut?: pulumi.Input<boolean>;
    /**
     * If set, artifacts will fail to download if a mismatch is detected between requested and received mimetype, according to
     * the list specified in the system properties file under blockedMismatchingMimeTypes. You can override by adding mimetypes
     * to the override list 'mismatching_mime_types_override_list'.
     */
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    bypassHeadRequests?: pulumi.Input<boolean>;
    /**
     * Client TLS certificate name.
     */
    clientTlsCertificate?: pulumi.Input<string>;
    contentSynchronisation?: pulumi.Input<inputs.RemotePuppetRepositoryContentSynchronisation>;
    /**
     * Public description.
     */
    description?: pulumi.Input<string>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only. Default value is 'false'.
     */
    downloadDirect?: pulumi.Input<boolean>;
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     */
    enableCookieManagement?: pulumi.Input<boolean>;
    /**
     * List of comma-separated artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*. By
     * default no artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
     * communicate with this repository.
     */
    hardFail?: pulumi.Input<boolean>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     */
    key?: pulumi.Input<string>;
    /**
     * Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
     * the 'Retrieval Cache Period'. Default value is 'true'.
     */
    listRemoteFolderItems?: pulumi.Input<boolean>;
    /**
     * The local address to be used when creating connections. Useful for specifying the interface to use on systems with
     * multiple network interfaces.
     */
    localAddress?: pulumi.Input<string>;
    /**
     * Metadata Retrieval Cache Timeout (Sec) in the UI.This value refers to the number of seconds to wait for retrieval from
     * the remote before serving locally cached artifact or fail the request.
     */
    metadataRetrievalTimeoutSecs?: pulumi.Input<number>;
    /**
     * The set of mime types that should override the block_mismatching_mime_types setting. Eg:
     * 'application/json,application/xml'. Default value is empty.
     */
    mismatchingMimeTypesOverrideList?: pulumi.Input<string>;
    /**
     * Missed Retrieval Cache Period (Sec) in the UI. The number of seconds to cache artifact retrieval misses (artifact not
     * found). A value of 0 indicates no caching.
     */
    missedCachePeriodSeconds?: pulumi.Input<number>;
    /**
     * Internal description.
     */
    notes?: pulumi.Input<string>;
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     */
    offline?: pulumi.Input<boolean>;
    packageType?: pulumi.Input<string>;
    password?: pulumi.Input<string>;
    /**
     * Setting Priority Resolution takes precedence over the resolution order when resolving virtual repositories. Setting
     * repositories with priority will cause metadata to be merged only from repositories set with a priority. If a package is
     * not found in those repositories, Artifactory will merge from repositories marked as non-priority.
     */
    priorityResolution?: pulumi.Input<boolean>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
     * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
     * will remain in the Terraform state, which will create state drift during the update.
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     */
    propagateQueryParams?: pulumi.Input<boolean>;
    /**
     * List of property set names
     */
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Proxy key from Artifactory Proxies settings
     */
    proxy?: pulumi.Input<string>;
    /**
     * Custom HTTP query parameters that will be automatically included in all remote resource requests. For example:
     * `param1=val1&param2=val2&param3=val3`
     */
    queryParams?: pulumi.Input<string>;
    /**
     * Repository layout key for the remote layout mapping.
     */
    remoteRepoLayoutRef?: pulumi.Input<string>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * Metadata Retrieval Cache Period (Sec) in the UI. This value refers to the number of seconds to cache metadata files
     * before checking for newer versions on remote server. A value of 0 indicates no caching.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    /**
     * Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
     * operation is considered a retrieval failure.
     */
    socketTimeoutMillis?: pulumi.Input<number>;
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     */
    storeArtifactsLocally?: pulumi.Input<boolean>;
    /**
     * When set, remote artifacts are fetched along with their properties.
     */
    synchronizeProperties?: pulumi.Input<boolean>;
    /**
     * Unused Artifacts Cleanup Period (Hr) in the UI. The number of hours to wait before an artifact is deemed 'unused' and
     * eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     */
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    /**
     * The remote repo URL.
     */
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     */
    xrayIndex?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a RemotePuppetRepository resource.
 */
export interface RemotePuppetRepositoryArgs {
    /**
     * 'Lenient Host Authentication' in the UI. Allow credentials of this repository to be used on requests redirected to any
     * other host.
     */
    allowAnyHostAuth?: pulumi.Input<boolean>;
    /**
     * The number of seconds the repository stays in assumed offline state after a connection error. At the end of this time,
     * an online check is attempted in order to reset the offline status. A value of 0 means the repository is never assumed
     * offline.
     */
    assumedOfflinePeriodSecs?: pulumi.Input<number>;
    /**
     * (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
     * resolution.
     */
    blackedOut?: pulumi.Input<boolean>;
    /**
     * If set, artifacts will fail to download if a mismatch is detected between requested and received mimetype, according to
     * the list specified in the system properties file under blockedMismatchingMimeTypes. You can override by adding mimetypes
     * to the override list 'mismatching_mime_types_override_list'.
     */
    blockMismatchingMimeTypes?: pulumi.Input<boolean>;
    /**
     * Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
     * HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
     * Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
     */
    bypassHeadRequests?: pulumi.Input<boolean>;
    /**
     * Client TLS certificate name.
     */
    clientTlsCertificate?: pulumi.Input<string>;
    contentSynchronisation?: pulumi.Input<inputs.RemotePuppetRepositoryContentSynchronisation>;
    /**
     * Public description.
     */
    description?: pulumi.Input<string>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only. Default value is 'false'.
     */
    downloadDirect?: pulumi.Input<boolean>;
    /**
     * Enables cookie management if the remote repository uses cookies to manage client state.
     */
    enableCookieManagement?: pulumi.Input<boolean>;
    /**
     * List of comma-separated artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*. By
     * default no artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * When set, Artifactory will return an error to the client that causes the build to fail if there is a failure to
     * communicate with this repository.
     */
    hardFail?: pulumi.Input<boolean>;
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     */
    key: pulumi.Input<string>;
    /**
     * Lists the items of remote folders in simple and list browsing. The remote content is cached according to the value of
     * the 'Retrieval Cache Period'. Default value is 'true'.
     */
    listRemoteFolderItems?: pulumi.Input<boolean>;
    /**
     * The local address to be used when creating connections. Useful for specifying the interface to use on systems with
     * multiple network interfaces.
     */
    localAddress?: pulumi.Input<string>;
    /**
     * Metadata Retrieval Cache Timeout (Sec) in the UI.This value refers to the number of seconds to wait for retrieval from
     * the remote before serving locally cached artifact or fail the request.
     */
    metadataRetrievalTimeoutSecs?: pulumi.Input<number>;
    /**
     * The set of mime types that should override the block_mismatching_mime_types setting. Eg:
     * 'application/json,application/xml'. Default value is empty.
     */
    mismatchingMimeTypesOverrideList?: pulumi.Input<string>;
    /**
     * Missed Retrieval Cache Period (Sec) in the UI. The number of seconds to cache artifact retrieval misses (artifact not
     * found). A value of 0 indicates no caching.
     */
    missedCachePeriodSeconds?: pulumi.Input<number>;
    /**
     * Internal description.
     */
    notes?: pulumi.Input<string>;
    /**
     * If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
     */
    offline?: pulumi.Input<boolean>;
    password?: pulumi.Input<string>;
    /**
     * Setting Priority Resolution takes precedence over the resolution order when resolving virtual repositories. Setting
     * repositories with priority will cause metadata to be merged only from repositories set with a priority. If a package is
     * not found in those repositories, Artifactory will merge from repositories marked as non-priority.
     */
    priorityResolution?: pulumi.Input<boolean>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
     * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
     * will remain in the Terraform state, which will create state drift during the update.
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * When set, if query params are included in the request to Artifactory, they will be passed on to the remote repository.
     */
    propagateQueryParams?: pulumi.Input<boolean>;
    /**
     * List of property set names
     */
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Proxy key from Artifactory Proxies settings
     */
    proxy?: pulumi.Input<string>;
    /**
     * Custom HTTP query parameters that will be automatically included in all remote resource requests. For example:
     * `param1=val1&param2=val2&param3=val3`
     */
    queryParams?: pulumi.Input<string>;
    /**
     * Repository layout key for the remote layout mapping.
     */
    remoteRepoLayoutRef?: pulumi.Input<string>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * Metadata Retrieval Cache Period (Sec) in the UI. This value refers to the number of seconds to cache metadata files
     * before checking for newer versions on remote server. A value of 0 indicates no caching.
     */
    retrievalCachePeriodSeconds?: pulumi.Input<number>;
    shareConfiguration?: pulumi.Input<boolean>;
    /**
     * Network timeout (in ms) to use when establishing a connection and for unanswered requests. Timing out on a network
     * operation is considered a retrieval failure.
     */
    socketTimeoutMillis?: pulumi.Input<number>;
    /**
     * When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
     * direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
     * one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
     * servers.
     */
    storeArtifactsLocally?: pulumi.Input<boolean>;
    /**
     * When set, remote artifacts are fetched along with their properties.
     */
    synchronizeProperties?: pulumi.Input<boolean>;
    /**
     * Unused Artifacts Cleanup Period (Hr) in the UI. The number of hours to wait before an artifact is deemed 'unused' and
     * eligible for cleanup from the repository. A value of 0 means automatic cleanup of cached artifacts is disabled.
     */
    unusedArtifactsCleanupPeriodHours?: pulumi.Input<number>;
    /**
     * The remote repo URL.
     */
    url: pulumi.Input<string>;
    username?: pulumi.Input<string>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     */
    xrayIndex?: pulumi.Input<boolean>;
}
