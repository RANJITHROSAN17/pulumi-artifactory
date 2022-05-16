// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Groups can be imported using their name, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/group:Group terraform-group mygroup
 * ```
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * Any users added to this group will automatically be assigned with admin privileges in the system.
     */
    public readonly adminPrivileges!: pulumi.Output<boolean>;
    /**
     * When this parameter is set, any new users defined in the system are automatically assigned to this group.
     */
    public readonly autoJoin!: pulumi.Output<boolean>;
    /**
     * A description for the group
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * When this override is set, an empty or missing usernames array will detach all users from the group
     */
    public readonly detachAllUsers!: pulumi.Output<boolean | undefined>;
    /**
     * New external group ID used to configure the corresponding group in Azure AD.
     */
    public readonly externalId!: pulumi.Output<string | undefined>;
    /**
     * Name of the group
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * When this override is set, User in the group can set Xray security and compliance policies. Default value is 'false'.
     */
    public readonly policyManager!: pulumi.Output<boolean | undefined>;
    /**
     * The realm for the group.
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * The realm attributes for the group.
     */
    public readonly realmAttributes!: pulumi.Output<string | undefined>;
    /**
     * When this override is set, User in the group can manage Xray Reports on any resource type. Default value is 'false'.
     */
    public readonly reportsManager!: pulumi.Output<boolean | undefined>;
    /**
     * List of users assigned to the group. If missing or empty, tf will not manage group membership
     */
    public readonly usersNames!: pulumi.Output<string[] | undefined>;
    /**
     * When this override is set, User in the group can manage Xray Watches on any resource type. Default value is 'false'.
     */
    public readonly watchManager!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["adminPrivileges"] = state ? state.adminPrivileges : undefined;
            resourceInputs["autoJoin"] = state ? state.autoJoin : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["detachAllUsers"] = state ? state.detachAllUsers : undefined;
            resourceInputs["externalId"] = state ? state.externalId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyManager"] = state ? state.policyManager : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["realmAttributes"] = state ? state.realmAttributes : undefined;
            resourceInputs["reportsManager"] = state ? state.reportsManager : undefined;
            resourceInputs["usersNames"] = state ? state.usersNames : undefined;
            resourceInputs["watchManager"] = state ? state.watchManager : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            resourceInputs["adminPrivileges"] = args ? args.adminPrivileges : undefined;
            resourceInputs["autoJoin"] = args ? args.autoJoin : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["detachAllUsers"] = args ? args.detachAllUsers : undefined;
            resourceInputs["externalId"] = args ? args.externalId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policyManager"] = args ? args.policyManager : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["realmAttributes"] = args ? args.realmAttributes : undefined;
            resourceInputs["reportsManager"] = args ? args.reportsManager : undefined;
            resourceInputs["usersNames"] = args ? args.usersNames : undefined;
            resourceInputs["watchManager"] = args ? args.watchManager : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * Any users added to this group will automatically be assigned with admin privileges in the system.
     */
    adminPrivileges?: pulumi.Input<boolean>;
    /**
     * When this parameter is set, any new users defined in the system are automatically assigned to this group.
     */
    autoJoin?: pulumi.Input<boolean>;
    /**
     * A description for the group
     */
    description?: pulumi.Input<string>;
    /**
     * When this override is set, an empty or missing usernames array will detach all users from the group
     */
    detachAllUsers?: pulumi.Input<boolean>;
    /**
     * New external group ID used to configure the corresponding group in Azure AD.
     */
    externalId?: pulumi.Input<string>;
    /**
     * Name of the group
     */
    name?: pulumi.Input<string>;
    /**
     * When this override is set, User in the group can set Xray security and compliance policies. Default value is 'false'.
     */
    policyManager?: pulumi.Input<boolean>;
    /**
     * The realm for the group.
     */
    realm?: pulumi.Input<string>;
    /**
     * The realm attributes for the group.
     */
    realmAttributes?: pulumi.Input<string>;
    /**
     * When this override is set, User in the group can manage Xray Reports on any resource type. Default value is 'false'.
     */
    reportsManager?: pulumi.Input<boolean>;
    /**
     * List of users assigned to the group. If missing or empty, tf will not manage group membership
     */
    usersNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When this override is set, User in the group can manage Xray Watches on any resource type. Default value is 'false'.
     */
    watchManager?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * Any users added to this group will automatically be assigned with admin privileges in the system.
     */
    adminPrivileges?: pulumi.Input<boolean>;
    /**
     * When this parameter is set, any new users defined in the system are automatically assigned to this group.
     */
    autoJoin?: pulumi.Input<boolean>;
    /**
     * A description for the group
     */
    description?: pulumi.Input<string>;
    /**
     * When this override is set, an empty or missing usernames array will detach all users from the group
     */
    detachAllUsers?: pulumi.Input<boolean>;
    /**
     * New external group ID used to configure the corresponding group in Azure AD.
     */
    externalId?: pulumi.Input<string>;
    /**
     * Name of the group
     */
    name?: pulumi.Input<string>;
    /**
     * When this override is set, User in the group can set Xray security and compliance policies. Default value is 'false'.
     */
    policyManager?: pulumi.Input<boolean>;
    /**
     * The realm for the group.
     */
    realm?: pulumi.Input<string>;
    /**
     * The realm attributes for the group.
     */
    realmAttributes?: pulumi.Input<string>;
    /**
     * When this override is set, User in the group can manage Xray Reports on any resource type. Default value is 'false'.
     */
    reportsManager?: pulumi.Input<boolean>;
    /**
     * List of users assigned to the group. If missing or empty, tf will not manage group membership
     */
    usersNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * When this override is set, User in the group can manage Xray Watches on any resource type. Default value is 'false'.
     */
    watchManager?: pulumi.Input<boolean>;
}
