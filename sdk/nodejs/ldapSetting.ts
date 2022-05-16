// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can be used to manage Artifactory's LDAP settings for user authentication.
 *
 * When specified LDAP setting is active, Artifactory first attempts to authenticate the user against the LDAP server.
 * If LDAP authentication fails, it then tries to authenticate via its internal database.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as artifactory from "@pulumi/artifactory";
 *
 * // Configure Artifactory LDAP setting
 * const ldapName = new artifactory.LdapSetting("ldap_name", {
 *     allowUserToAccessProfile: false,
 *     autoCreateUser: true,
 *     emailAttribute: "mail",
 *     enabled: true,
 *     key: "ldap_name",
 *     ldapPoisoningProtection: true,
 *     ldapUrl: "ldap://ldap_server_url",
 *     managerDn: "mgr_dn",
 *     managerPassword: "mgr_passwd_random",
 *     pagingSupportEnabled: false,
 *     searchBase: "ou=users",
 *     searchFilter: "(uid={0})",
 *     searchSubTree: true,
 *     userDnPattern: "uid={0},ou=People",
 * });
 * ```
 * Note: `Key` argument has to match to the resource name.\
 * Reference Link: [JFrog LDAP](https://www.jfrog.com/confluence/display/JFROG/LDAP)
 *
 * ## Import
 *
 * LDAP setting can be imported using the key, e.g.
 *
 * ```sh
 *  $ pulumi import artifactory:index/ldapSetting:LdapSetting ldap_name ldap_name
 * ```
 */
export class LdapSetting extends pulumi.CustomResource {
    /**
     * Get an existing LdapSetting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LdapSettingState, opts?: pulumi.CustomResourceOptions): LdapSetting {
        return new LdapSetting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'artifactory:index/ldapSetting:LdapSetting';

    /**
     * Returns true if the given object is an instance of LdapSetting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LdapSetting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LdapSetting.__pulumiType;
    }

    /**
     * When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
     */
    public readonly allowUserToAccessProfile!: pulumi.Output<boolean | undefined>;
    /**
     * When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
     */
    public readonly autoCreateUser!: pulumi.Output<boolean | undefined>;
    /**
     * An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is `mail`.
     * - Note: If blank/empty string input was set for email_attribute, Default value "mail" takes effect. This is to match with Artifactory behavior.
     */
    public readonly emailAttribute!: pulumi.Output<string | undefined>;
    /**
     * When set, these settings are enabled. Default value is `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The unique ID of the LDAP setting.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
     */
    public readonly ldapPoisoningProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
     */
    public readonly ldapUrl!: pulumi.Output<string>;
    /**
     * The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
     */
    public readonly managerDn!: pulumi.Output<string | undefined>;
    /**
     * The password of the user binding to the LDAP server when using "search" authentication.
     */
    public readonly managerPassword!: pulumi.Output<string>;
    /**
     * When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
     */
    public readonly pagingSupportEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
     */
    public readonly searchBase!: pulumi.Output<string | undefined>;
    /**
     * A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, denoted by '{0}'. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty. 
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
     */
    public readonly searchFilter!: pulumi.Output<string | undefined>;
    /**
     * When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
     */
    public readonly searchSubTree!: pulumi.Output<boolean | undefined>;
    /**
     * A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for "direct" user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
     */
    public readonly userDnPattern!: pulumi.Output<string | undefined>;

    /**
     * Create a LdapSetting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LdapSettingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LdapSettingArgs | LdapSettingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LdapSettingState | undefined;
            resourceInputs["allowUserToAccessProfile"] = state ? state.allowUserToAccessProfile : undefined;
            resourceInputs["autoCreateUser"] = state ? state.autoCreateUser : undefined;
            resourceInputs["emailAttribute"] = state ? state.emailAttribute : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["ldapPoisoningProtection"] = state ? state.ldapPoisoningProtection : undefined;
            resourceInputs["ldapUrl"] = state ? state.ldapUrl : undefined;
            resourceInputs["managerDn"] = state ? state.managerDn : undefined;
            resourceInputs["managerPassword"] = state ? state.managerPassword : undefined;
            resourceInputs["pagingSupportEnabled"] = state ? state.pagingSupportEnabled : undefined;
            resourceInputs["searchBase"] = state ? state.searchBase : undefined;
            resourceInputs["searchFilter"] = state ? state.searchFilter : undefined;
            resourceInputs["searchSubTree"] = state ? state.searchSubTree : undefined;
            resourceInputs["userDnPattern"] = state ? state.userDnPattern : undefined;
        } else {
            const args = argsOrState as LdapSettingArgs | undefined;
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            if ((!args || args.ldapUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapUrl'");
            }
            resourceInputs["allowUserToAccessProfile"] = args ? args.allowUserToAccessProfile : undefined;
            resourceInputs["autoCreateUser"] = args ? args.autoCreateUser : undefined;
            resourceInputs["emailAttribute"] = args ? args.emailAttribute : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["ldapPoisoningProtection"] = args ? args.ldapPoisoningProtection : undefined;
            resourceInputs["ldapUrl"] = args ? args.ldapUrl : undefined;
            resourceInputs["managerDn"] = args ? args.managerDn : undefined;
            resourceInputs["managerPassword"] = args ? args.managerPassword : undefined;
            resourceInputs["pagingSupportEnabled"] = args ? args.pagingSupportEnabled : undefined;
            resourceInputs["searchBase"] = args ? args.searchBase : undefined;
            resourceInputs["searchFilter"] = args ? args.searchFilter : undefined;
            resourceInputs["searchSubTree"] = args ? args.searchSubTree : undefined;
            resourceInputs["userDnPattern"] = args ? args.userDnPattern : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LdapSetting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LdapSetting resources.
 */
export interface LdapSettingState {
    /**
     * When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
     */
    allowUserToAccessProfile?: pulumi.Input<boolean>;
    /**
     * When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
     */
    autoCreateUser?: pulumi.Input<boolean>;
    /**
     * An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is `mail`.
     * - Note: If blank/empty string input was set for email_attribute, Default value "mail" takes effect. This is to match with Artifactory behavior.
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * When set, these settings are enabled. Default value is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique ID of the LDAP setting.
     */
    key?: pulumi.Input<string>;
    /**
     * Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
     */
    ldapPoisoningProtection?: pulumi.Input<boolean>;
    /**
     * Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
     */
    ldapUrl?: pulumi.Input<string>;
    /**
     * The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
     */
    managerDn?: pulumi.Input<string>;
    /**
     * The password of the user binding to the LDAP server when using "search" authentication.
     */
    managerPassword?: pulumi.Input<string>;
    /**
     * When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
     */
    pagingSupportEnabled?: pulumi.Input<boolean>;
    /**
     * The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
     */
    searchBase?: pulumi.Input<string>;
    /**
     * A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, denoted by '{0}'. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty. 
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
     */
    searchFilter?: pulumi.Input<string>;
    /**
     * When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
     */
    searchSubTree?: pulumi.Input<boolean>;
    /**
     * A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for "direct" user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
     */
    userDnPattern?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LdapSetting resource.
 */
export interface LdapSettingArgs {
    /**
     * When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
     */
    allowUserToAccessProfile?: pulumi.Input<boolean>;
    /**
     * When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
     */
    autoCreateUser?: pulumi.Input<boolean>;
    /**
     * An attribute that can be used to map a user's email address to a user created automatically in Artifactory. Default value is `mail`.
     * - Note: If blank/empty string input was set for email_attribute, Default value "mail" takes effect. This is to match with Artifactory behavior.
     */
    emailAttribute?: pulumi.Input<string>;
    /**
     * When set, these settings are enabled. Default value is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The unique ID of the LDAP setting.
     */
    key: pulumi.Input<string>;
    /**
     * Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
     */
    ldapPoisoningProtection?: pulumi.Input<boolean>;
    /**
     * Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
     */
    ldapUrl: pulumi.Input<string>;
    /**
     * The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
     */
    managerDn?: pulumi.Input<string>;
    /**
     * The password of the user binding to the LDAP server when using "search" authentication.
     */
    managerPassword?: pulumi.Input<string>;
    /**
     * When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
     */
    pagingSupportEnabled?: pulumi.Input<boolean>;
    /**
     * The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
     */
    searchBase?: pulumi.Input<string>;
    /**
     * A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in 'RFC 2254') with optional arguments. In this case, the username is the only argument, denoted by '{0}'. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty. 
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
     */
    searchFilter?: pulumi.Input<string>;
    /**
     * When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
     */
    searchSubTree?: pulumi.Input<boolean>;
    /**
     * A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for "direct" user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
     */
    userDnPattern?: pulumi.Input<string>;
}
