// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * RSA key pairs are used to sign and verify the Alpine Linux index files in JFrog Artifactory, while GPG key pairs are
 * used to sign and validate packages integrity in JFrog Distribution. The JFrog Platform enables you to manage multiple
 * RSA and GPG signing keys through the Keys Management UI and REST API. The JFrog Platform supports managing multiple
 * pairs of GPG signing keys to sign packages for authentication of several package types such as Debian, Opkg, and RPM
 * through the Keys Management UI and REST API.
 * Passphrases are not currently supported, though they exist in the API.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 * import * as fs from "fs";
 *
 * const some_keypair6543461672124900137 = new artifactory.Keypair("some-keypair6543461672124900137", {
 *     pairName: "some-keypair6543461672124900137",
 *     pairType: "RSA",
 *     alias: "foo-alias6543461672124900137",
 *     privateKey: fs.readFileSync("samples/rsa.priv"),
 *     publicKey: fs.readFileSync("samples/rsa.pub"),
 * });
 * ```
 *
 * ## Import
 *
 * Keypair can be imported using their name, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/keypair:Keypair my-keypair my-keypair
 * ```
 */
export class Keypair extends pulumi.CustomResource {
    /**
     * Get an existing Keypair resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: KeypairState, opts?: pulumi.CustomResourceOptions): Keypair {
        return new Keypair(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/keypair:Keypair';

    /**
     * Returns true if the given object is an instance of Keypair.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Keypair {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Keypair.__pulumiType;
    }

    /**
     * Will be used as a filename when retrieving the public key via REST API.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * A unique identifier for the Key Pair record.
     */
    public readonly pairName!: pulumi.Output<string>;
    /**
     * Key Pair type. Supported types - GPG and RSA.
     */
    public readonly pairType!: pulumi.Output<string>;
    /**
     * Passphrase will be used to decrypt the private key. Validated server side.
     */
    public readonly passphrase!: pulumi.Output<string | undefined>;
    /**
     * - Private key. PEM format will be validated.
     */
    public readonly privateKey!: pulumi.Output<string>;
    /**
     * Public key. PEM format will be validated.
     */
    public readonly publicKey!: pulumi.Output<string>;
    /**
     * Unknown usage. Returned in the json payload and cannot be set.
     */
    public /*out*/ readonly unavailable!: pulumi.Output<boolean>;

    /**
     * Create a Keypair resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KeypairArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: KeypairArgs | KeypairState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as KeypairState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["pairName"] = state ? state.pairName : undefined;
            resourceInputs["pairType"] = state ? state.pairType : undefined;
            resourceInputs["passphrase"] = state ? state.passphrase : undefined;
            resourceInputs["privateKey"] = state ? state.privateKey : undefined;
            resourceInputs["publicKey"] = state ? state.publicKey : undefined;
            resourceInputs["unavailable"] = state ? state.unavailable : undefined;
        } else {
            const args = argsOrState as KeypairArgs | undefined;
            if ((!args || args.alias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alias'");
            }
            if ((!args || args.pairName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pairName'");
            }
            if ((!args || args.pairType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'pairType'");
            }
            if ((!args || args.privateKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'privateKey'");
            }
            if ((!args || args.publicKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'publicKey'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["pairName"] = args ? args.pairName : undefined;
            resourceInputs["pairType"] = args ? args.pairType : undefined;
            resourceInputs["passphrase"] = args ? args.passphrase : undefined;
            resourceInputs["privateKey"] = args ? args.privateKey : undefined;
            resourceInputs["publicKey"] = args ? args.publicKey : undefined;
            resourceInputs["unavailable"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Keypair.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Keypair resources.
 */
export interface KeypairState {
    /**
     * Will be used as a filename when retrieving the public key via REST API.
     */
    alias?: pulumi.Input<string>;
    /**
     * A unique identifier for the Key Pair record.
     */
    pairName?: pulumi.Input<string>;
    /**
     * Key Pair type. Supported types - GPG and RSA.
     */
    pairType?: pulumi.Input<string>;
    /**
     * Passphrase will be used to decrypt the private key. Validated server side.
     */
    passphrase?: pulumi.Input<string>;
    /**
     * - Private key. PEM format will be validated.
     */
    privateKey?: pulumi.Input<string>;
    /**
     * Public key. PEM format will be validated.
     */
    publicKey?: pulumi.Input<string>;
    /**
     * Unknown usage. Returned in the json payload and cannot be set.
     */
    unavailable?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Keypair resource.
 */
export interface KeypairArgs {
    /**
     * Will be used as a filename when retrieving the public key via REST API.
     */
    alias: pulumi.Input<string>;
    /**
     * A unique identifier for the Key Pair record.
     */
    pairName: pulumi.Input<string>;
    /**
     * Key Pair type. Supported types - GPG and RSA.
     */
    pairType: pulumi.Input<string>;
    /**
     * Passphrase will be used to decrypt the private key. Validated server side.
     */
    passphrase?: pulumi.Input<string>;
    /**
     * - Private key. PEM format will be validated.
     */
    privateKey: pulumi.Input<string>;
    /**
     * Public key. PEM format will be validated.
     */
    publicKey: pulumi.Input<string>;
}
