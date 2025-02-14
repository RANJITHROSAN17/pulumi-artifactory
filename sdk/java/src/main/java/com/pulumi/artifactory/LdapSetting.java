// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.LdapSettingArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.LdapSettingState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can be used to manage Artifactory&#39;s LDAP settings for user authentication.
 * 
 * When specified LDAP setting is active, Artifactory first attempts to authenticate the user against the LDAP server.
 * If LDAP authentication fails, it then tries to authenticate via its internal database.
 * 
 * ~&gt;The `artifactory.LdapSetting` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.LdapSetting;
 * import com.pulumi.artifactory.LdapSettingArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var ldapName = new LdapSetting(&#34;ldapName&#34;, LdapSettingArgs.builder()        
 *             .allowUserToAccessProfile(false)
 *             .autoCreateUser(true)
 *             .emailAttribute(&#34;mail&#34;)
 *             .enabled(true)
 *             .key(&#34;ldap_name&#34;)
 *             .ldapPoisoningProtection(true)
 *             .ldapUrl(&#34;ldap://ldap_server_url&#34;)
 *             .managerDn(&#34;mgr_dn&#34;)
 *             .managerPassword(&#34;mgr_passwd_random&#34;)
 *             .pagingSupportEnabled(false)
 *             .searchBase(&#34;ou=users&#34;)
 *             .searchFilter(&#34;(uid={0})&#34;)
 *             .searchSubTree(true)
 *             .userDnPattern(&#34;uid={0},ou=People&#34;)
 *             .build());
 * 
 *     }
 * }
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
 * 
 */
@ResourceType(type="artifactory:index/ldapSetting:LdapSetting")
public class LdapSetting extends com.pulumi.resources.CustomResource {
    /**
     * When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
     * 
     */
    @Export(name="allowUserToAccessProfile", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowUserToAccessProfile;

    /**
     * @return When set, users created after logging in using LDAP will be able to access their profile page.  Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> allowUserToAccessProfile() {
        return Codegen.optional(this.allowUserToAccessProfile);
    }
    /**
     * When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
     * 
     */
    @Export(name="autoCreateUser", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoCreateUser;

    /**
     * @return When set, the system will automatically create new users for those who have logged in using LDAP, and assign them to the default groups.  Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> autoCreateUser() {
        return Codegen.optional(this.autoCreateUser);
    }
    /**
     * An attribute that can be used to map a user&#39;s email address to a user created automatically in Artifactory. Default value is `mail`.
     * - Note: If blank/empty string input was set for email_attribute, Default value `mail` takes effect. This is to match with Artifactory behavior.
     * 
     */
    @Export(name="emailAttribute", type=String.class, parameters={})
    private Output</* @Nullable */ String> emailAttribute;

    /**
     * @return An attribute that can be used to map a user&#39;s email address to a user created automatically in Artifactory. Default value is `mail`.
     * - Note: If blank/empty string input was set for email_attribute, Default value `mail` takes effect. This is to match with Artifactory behavior.
     * 
     */
    public Output<Optional<String>> emailAttribute() {
        return Codegen.optional(this.emailAttribute);
    }
    /**
     * When set, these settings are enabled. Default value is `true`.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return When set, these settings are enabled. Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The unique ID of the LDAP setting.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return The unique ID of the LDAP setting.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
     * 
     */
    @Export(name="ldapPoisoningProtection", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> ldapPoisoningProtection;

    /**
     * @return Protects against LDAP poisoning by filtering out users exposed to vulnerabilities.  Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> ldapPoisoningProtection() {
        return Codegen.optional(this.ldapPoisoningProtection);
    }
    /**
     * Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
     * 
     */
    @Export(name="ldapUrl", type=String.class, parameters={})
    private Output<String> ldapUrl;

    /**
     * @return Location of the LDAP server in the following format: ldap://myserver:myport/dc=sampledomain,dc=com. The URL should include the base DN used to search for and/or authenticate users.
     * 
     */
    public Output<String> ldapUrl() {
        return this.ldapUrl;
    }
    /**
     * The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
     * 
     */
    @Export(name="managerDn", type=String.class, parameters={})
    private Output</* @Nullable */ String> managerDn;

    /**
     * @return The full DN of a user with permissions that allow querying the LDAP server. When working with LDAP Groups, the user should have permissions for any extra group attributes such as memberOf.
     * 
     */
    public Output<Optional<String>> managerDn() {
        return Codegen.optional(this.managerDn);
    }
    /**
     * The password of the user binding to the LDAP server when using &#34;search&#34; authentication.
     * 
     */
    @Export(name="managerPassword", type=String.class, parameters={})
    private Output<String> managerPassword;

    /**
     * @return The password of the user binding to the LDAP server when using &#34;search&#34; authentication.
     * 
     */
    public Output<String> managerPassword() {
        return this.managerPassword;
    }
    /**
     * When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
     * 
     */
    @Export(name="pagingSupportEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> pagingSupportEnabled;

    /**
     * @return When set, supports paging results for the LDAP server. This feature requires that the LDAP Server supports a PagedResultsControl configuration.  Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> pagingSupportEnabled() {
        return Codegen.optional(this.pagingSupportEnabled);
    }
    /**
     * The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
     * 
     */
    @Export(name="searchBase", type=String.class, parameters={})
    private Output</* @Nullable */ String> searchBase;

    /**
     * @return The Context name in which to search relative to the base DN in the LDAP URL. Multiple search bases may be specified separated by a pipe ( | ).
     * 
     */
    public Output<Optional<String>> searchBase() {
        return Codegen.optional(this.searchBase);
    }
    /**
     * A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in &#39;RFC 2254&#39;) with optional arguments. In this case, the username is the only argument, denoted by &#39;{0}&#39;. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty.
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
     * 
     */
    @Export(name="searchFilter", type=String.class, parameters={})
    private Output</* @Nullable */ String> searchFilter;

    /**
     * @return A filter expression used to search for the user DN that is used in LDAP authentication. This is an LDAP search filter (as defined in &#39;RFC 2254&#39;) with optional arguments. In this case, the username is the only argument, denoted by &#39;{0}&#39;. Possible examples are: uid={0}) - this would search for a username match on the uid attribute. Authentication using LDAP is performed from the DN found if successful. Default value is blank/empty.
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both)
     * 
     */
    public Output<Optional<String>> searchFilter() {
        return Codegen.optional(this.searchFilter);
    }
    /**
     * When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
     * 
     */
    @Export(name="searchSubTree", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> searchSubTree;

    /**
     * @return When set, enables deep search through the sub-tree of the LDAP URL + Search Base.  Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> searchSubTree() {
        return Codegen.optional(this.searchSubTree);
    }
    /**
     * A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for &#34;direct&#34; user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
     * 
     */
    @Export(name="userDnPattern", type=String.class, parameters={})
    private Output</* @Nullable */ String> userDnPattern;

    /**
     * @return A DN pattern used to log users directly in to the LDAP database. This pattern is used to create a DN string for &#34;direct&#34; user authentication, and is relative to the base DN in the LDAP URL. The pattern argument {0} is replaced with the username at runtime. This only works if anonymous binding is allowed and a direct user DN can be used (which is not the default case for Active Directory). For example: uid={0},ou=People. Default value is blank/empty.
     * - Note: LDAP settings should provide a userDnPattern or a searchFilter (or both).
     * 
     */
    public Output<Optional<String>> userDnPattern() {
        return Codegen.optional(this.userDnPattern);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LdapSetting(String name) {
        this(name, LdapSettingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LdapSetting(String name, LdapSettingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LdapSetting(String name, LdapSettingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/ldapSetting:LdapSetting", name, args == null ? LdapSettingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LdapSetting(String name, Output<String> id, @Nullable LdapSettingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/ldapSetting:LdapSetting", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "managerPassword"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static LdapSetting get(String name, Output<String> id, @Nullable LdapSettingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LdapSetting(name, id, state, options);
    }
}
