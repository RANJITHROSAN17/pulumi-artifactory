// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Creates a local Debian repository and allows for the creation of a GPG key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 * import * from "fs";
 *
 * const some_keypairGPG1 = new artifactory.Keypair("some-keypairGPG1", {
 *     pairName: `some-keypair${random_id.randid.id}`,
 *     pairType: "GPG",
 *     alias: "foo-alias1",
 *     privateKey: fs.readFileSync("samples/gpg.priv"),
 *     publicKey: fs.readFileSync("samples/gpg.pub"),
 * });
 * const some_keypairGPG2 = new artifactory.Keypair("some-keypairGPG2", {
 *     pairName: `some-keypair4${random_id.randid.id}`,
 *     pairType: "GPG",
 *     alias: "foo-alias2",
 *     privateKey: fs.readFileSync("samples/gpg.priv"),
 *     publicKey: fs.readFileSync("samples/gpg.pub"),
 * });
 * const my_debian_repo = new artifactory.DebianRepository("my-debian-repo", {
 *     key: "my-debian-repo",
 *     primaryKeypairRef: some_keypairGPG1.pairName,
 *     secondaryKeypairRef: some_keypairGPG2.pairName,
 *     indexCompressionFormats: [
 *         "bz2",
 *         "lzma",
 *         "xz",
 *     ],
 *     trivialLayout: true,
 * }, {
 *     dependsOn: [
 *         some_keypairGPG1,
 *         some_keypairGPG2,
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Local repositories can be imported using their name, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/debianRepository:DebianRepository my-debian-repo my-debian-repo
 * ```
 */
export class DebianRepository extends pulumi.CustomResource {
    /**
     * Get an existing DebianRepository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DebianRepositoryState, opts?: pulumi.CustomResourceOptions): DebianRepository {
        return new DebianRepository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/debianRepository:DebianRepository';

    /**
     * Returns true if the given object is an instance of DebianRepository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DebianRepository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DebianRepository.__pulumiType;
    }

    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    public readonly archiveBrowsingEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     */
    public readonly blackedOut!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     */
    public readonly downloadDirect!: pulumi.Output<boolean | undefined>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*. By default no
     * artifacts are excluded.
     */
    public readonly excludesPattern!: pulumi.Output<string>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    public readonly includesPattern!: pulumi.Output<string>;
    /**
     * The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
     * and XZ (.xz extension).
     */
    public readonly indexCompressionFormats!: pulumi.Output<string[] | undefined>;
    /**
     * the identity key of the repo.
     */
    public readonly key!: pulumi.Output<string>;
    public readonly notes!: pulumi.Output<string | undefined>;
    public /*out*/ readonly packageType!: pulumi.Output<string>;
    /**
     * The primary RSA key to be used to sign packages.
     */
    public readonly primaryKeypairRef!: pulumi.Output<string | undefined>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    public readonly priorityResolution!: pulumi.Output<boolean | undefined>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    public readonly projectEnvironments!: pulumi.Output<string[] | undefined>;
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     */
    public readonly projectKey!: pulumi.Output<string | undefined>;
    /**
     * List of property set name
     */
    public readonly propertySets!: pulumi.Output<string[] | undefined>;
    /**
     * Repository layout key for the local repository
     */
    public readonly repoLayoutRef!: pulumi.Output<string | undefined>;
    /**
     * The secondary RSA key to be used to sign packages.
     */
    public readonly secondaryKeypairRef!: pulumi.Output<string | undefined>;
    /**
     * When set, the repository will use the deprecated trivial layout.
     *
     * @deprecated You shouldn't be using this
     */
    public readonly trivialLayout!: pulumi.Output<boolean | undefined>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     */
    public readonly xrayIndex!: pulumi.Output<boolean | undefined>;

    /**
     * Create a DebianRepository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DebianRepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DebianRepositoryArgs | DebianRepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DebianRepositoryState | undefined;
            resourceInputs["archiveBrowsingEnabled"] = state ? state.archiveBrowsingEnabled : undefined;
            resourceInputs["blackedOut"] = state ? state.blackedOut : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["downloadDirect"] = state ? state.downloadDirect : undefined;
            resourceInputs["excludesPattern"] = state ? state.excludesPattern : undefined;
            resourceInputs["includesPattern"] = state ? state.includesPattern : undefined;
            resourceInputs["indexCompressionFormats"] = state ? state.indexCompressionFormats : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["notes"] = state ? state.notes : undefined;
            resourceInputs["packageType"] = state ? state.packageType : undefined;
            resourceInputs["primaryKeypairRef"] = state ? state.primaryKeypairRef : undefined;
            resourceInputs["priorityResolution"] = state ? state.priorityResolution : undefined;
            resourceInputs["projectEnvironments"] = state ? state.projectEnvironments : undefined;
            resourceInputs["projectKey"] = state ? state.projectKey : undefined;
            resourceInputs["propertySets"] = state ? state.propertySets : undefined;
            resourceInputs["repoLayoutRef"] = state ? state.repoLayoutRef : undefined;
            resourceInputs["secondaryKeypairRef"] = state ? state.secondaryKeypairRef : undefined;
            resourceInputs["trivialLayout"] = state ? state.trivialLayout : undefined;
            resourceInputs["xrayIndex"] = state ? state.xrayIndex : undefined;
        } else {
            const args = argsOrState as DebianRepositoryArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["archiveBrowsingEnabled"] = args ? args.archiveBrowsingEnabled : undefined;
            resourceInputs["blackedOut"] = args ? args.blackedOut : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["downloadDirect"] = args ? args.downloadDirect : undefined;
            resourceInputs["excludesPattern"] = args ? args.excludesPattern : undefined;
            resourceInputs["includesPattern"] = args ? args.includesPattern : undefined;
            resourceInputs["indexCompressionFormats"] = args ? args.indexCompressionFormats : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["notes"] = args ? args.notes : undefined;
            resourceInputs["primaryKeypairRef"] = args ? args.primaryKeypairRef : undefined;
            resourceInputs["priorityResolution"] = args ? args.priorityResolution : undefined;
            resourceInputs["projectEnvironments"] = args ? args.projectEnvironments : undefined;
            resourceInputs["projectKey"] = args ? args.projectKey : undefined;
            resourceInputs["propertySets"] = args ? args.propertySets : undefined;
            resourceInputs["repoLayoutRef"] = args ? args.repoLayoutRef : undefined;
            resourceInputs["secondaryKeypairRef"] = args ? args.secondaryKeypairRef : undefined;
            resourceInputs["trivialLayout"] = args ? args.trivialLayout : undefined;
            resourceInputs["xrayIndex"] = args ? args.xrayIndex : undefined;
            resourceInputs["packageType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DebianRepository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DebianRepository resources.
 */
export interface DebianRepositoryState {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     */
    blackedOut?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     */
    downloadDirect?: pulumi.Input<boolean>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*. By default no
     * artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
     * and XZ (.xz extension).
     */
    indexCompressionFormats?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * the identity key of the repo.
     */
    key?: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    packageType?: pulumi.Input<string>;
    /**
     * The primary RSA key to be used to sign packages.
     */
    primaryKeypairRef?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * List of property set name
     */
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The secondary RSA key to be used to sign packages.
     */
    secondaryKeypairRef?: pulumi.Input<string>;
    /**
     * When set, the repository will use the deprecated trivial layout.
     *
     * @deprecated You shouldn't be using this
     */
    trivialLayout?: pulumi.Input<boolean>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     */
    xrayIndex?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a DebianRepository resource.
 */
export interface DebianRepositoryArgs {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     */
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     */
    blackedOut?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     */
    downloadDirect?: pulumi.Input<boolean>;
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**&#47;z/*. By default no
     * artifacts are excluded.
     */
    excludesPattern?: pulumi.Input<string>;
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**&#47;z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**&#47;*).
     */
    includesPattern?: pulumi.Input<string>;
    /**
     * The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
     * and XZ (.xz extension).
     */
    indexCompressionFormats?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * the identity key of the repo.
     */
    key: pulumi.Input<string>;
    notes?: pulumi.Input<string>;
    /**
     * The primary RSA key to be used to sign packages.
     */
    primaryKeypairRef?: pulumi.Input<string>;
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     */
    priorityResolution?: pulumi.Input<boolean>;
    /**
     * Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
     */
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     */
    projectKey?: pulumi.Input<string>;
    /**
     * List of property set name
     */
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Repository layout key for the local repository
     */
    repoLayoutRef?: pulumi.Input<string>;
    /**
     * The secondary RSA key to be used to sign packages.
     */
    secondaryKeypairRef?: pulumi.Input<string>;
    /**
     * When set, the repository will use the deprecated trivial layout.
     *
     * @deprecated You shouldn't be using this
     */
    trivialLayout?: pulumi.Input<boolean>;
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     */
    xrayIndex?: pulumi.Input<boolean>;
}
