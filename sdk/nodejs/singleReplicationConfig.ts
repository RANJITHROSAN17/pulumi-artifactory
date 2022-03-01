// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * // Create a replication between two artifactory local repositories
 * const providerTestSource = new artifactory.LocalRepository("provider_test_source", {
 *     key: "provider_test_source",
 *     packageType: "maven",
 * });
 * const providerTestDest = new artifactory.LocalRepository("provider_test_dest", {
 *     key: "provider_test_dest",
 *     packageType: "maven",
 * });
 * const foo_rep = new artifactory.SingleReplicationConfig("foo-rep", {
 *     cronExp: "0 0 * * * ?",
 *     enableEventReplication: true,
 *     password: var_artifactory_password,
 *     repoKey: providerTestSource.key,
 *     url: var_artifactory_url,
 *     username: var_artifactory_username,
 * });
 * ```
 *
 * ## Import
 *
 * Replication configs can be imported using their repo key, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/singleReplicationConfig:SingleReplicationConfig foo-rep repository-key
 * ```
 */
export class SingleReplicationConfig extends pulumi.CustomResource {
    /**
     * Get an existing SingleReplicationConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SingleReplicationConfigState, opts?: pulumi.CustomResourceOptions): SingleReplicationConfig {
        return new SingleReplicationConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/singleReplicationConfig:SingleReplicationConfig';

    /**
     * Returns true if the given object is an instance of SingleReplicationConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SingleReplicationConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SingleReplicationConfig.__pulumiType;
    }

    public readonly cronExp!: pulumi.Output<string>;
    public readonly enableEventReplication!: pulumi.Output<boolean>;
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    public /*out*/ readonly password!: pulumi.Output<string>;
    public readonly pathPrefix!: pulumi.Output<string | undefined>;
    public readonly repoKey!: pulumi.Output<string>;
    public readonly socketTimeoutMillis!: pulumi.Output<number>;
    public readonly syncDeletes!: pulumi.Output<boolean>;
    public readonly syncProperties!: pulumi.Output<boolean>;
    public readonly syncStatistics!: pulumi.Output<boolean>;
    public readonly url!: pulumi.Output<string | undefined>;
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a SingleReplicationConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SingleReplicationConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SingleReplicationConfigArgs | SingleReplicationConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SingleReplicationConfigState | undefined;
            resourceInputs["cronExp"] = state ? state.cronExp : undefined;
            resourceInputs["enableEventReplication"] = state ? state.enableEventReplication : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["pathPrefix"] = state ? state.pathPrefix : undefined;
            resourceInputs["repoKey"] = state ? state.repoKey : undefined;
            resourceInputs["socketTimeoutMillis"] = state ? state.socketTimeoutMillis : undefined;
            resourceInputs["syncDeletes"] = state ? state.syncDeletes : undefined;
            resourceInputs["syncProperties"] = state ? state.syncProperties : undefined;
            resourceInputs["syncStatistics"] = state ? state.syncStatistics : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as SingleReplicationConfigArgs | undefined;
            if ((!args || args.cronExp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cronExp'");
            }
            if ((!args || args.repoKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'repoKey'");
            }
            resourceInputs["cronExp"] = args ? args.cronExp : undefined;
            resourceInputs["enableEventReplication"] = args ? args.enableEventReplication : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["pathPrefix"] = args ? args.pathPrefix : undefined;
            resourceInputs["repoKey"] = args ? args.repoKey : undefined;
            resourceInputs["socketTimeoutMillis"] = args ? args.socketTimeoutMillis : undefined;
            resourceInputs["syncDeletes"] = args ? args.syncDeletes : undefined;
            resourceInputs["syncProperties"] = args ? args.syncProperties : undefined;
            resourceInputs["syncStatistics"] = args ? args.syncStatistics : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
            resourceInputs["password"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SingleReplicationConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SingleReplicationConfig resources.
 */
export interface SingleReplicationConfigState {
    cronExp?: pulumi.Input<string>;
    enableEventReplication?: pulumi.Input<boolean>;
    enabled?: pulumi.Input<boolean>;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    password?: pulumi.Input<string>;
    pathPrefix?: pulumi.Input<string>;
    repoKey?: pulumi.Input<string>;
    socketTimeoutMillis?: pulumi.Input<number>;
    syncDeletes?: pulumi.Input<boolean>;
    syncProperties?: pulumi.Input<boolean>;
    syncStatistics?: pulumi.Input<boolean>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SingleReplicationConfig resource.
 */
export interface SingleReplicationConfigArgs {
    cronExp: pulumi.Input<string>;
    enableEventReplication?: pulumi.Input<boolean>;
    enabled?: pulumi.Input<boolean>;
    pathPrefix?: pulumi.Input<string>;
    repoKey: pulumi.Input<string>;
    socketTimeoutMillis?: pulumi.Input<number>;
    syncDeletes?: pulumi.Input<boolean>;
    syncProperties?: pulumi.Input<boolean>;
    syncStatistics?: pulumi.Input<boolean>;
    url?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}
