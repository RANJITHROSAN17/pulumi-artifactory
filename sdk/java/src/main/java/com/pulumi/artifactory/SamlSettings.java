// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.SamlSettingsArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.SamlSettingsState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can be used to manage Artifactory&#39;s SAML SSO settings.
 * 
 * Only a single `artifactory.SamlSettings` resource is meant to be defined.
 * 
 * ~&gt;The `artifactory.SamlSettings` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.SamlSettings;
 * import com.pulumi.artifactory.SamlSettingsArgs;
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
 *         var saml = new SamlSettings(&#34;saml&#34;, SamlSettingsArgs.builder()        
 *             .allowUserToAccessProfile(true)
 *             .autoRedirect(true)
 *             .certificate(&#34;test-certificate&#34;)
 *             .emailAttribute(&#34;email&#34;)
 *             .enable(true)
 *             .groupAttribute(&#34;groups&#34;)
 *             .loginUrl(&#34;test-login-url&#34;)
 *             .logoutUrl(&#34;test-logout-url&#34;)
 *             .noAutoUserCreation(false)
 *             .serviceProviderName(&#34;okta&#34;)
 *             .syncGroups(true)
 *             .useEncryptedAssertion(false)
 *             .verifyAudienceRestriction(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Current SAML SSO settings can be imported using `saml_settings` as the `ID`, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/samlSettings:SamlSettings saml saml_settings
 * ```
 * 
 */
@ResourceType(type="artifactory:index/samlSettings:SamlSettings")
public class SamlSettings extends com.pulumi.resources.CustomResource {
    /**
     * Allow persisted users to access their profile.  Default value is `true`.
     * 
     */
    @Export(name="allowUserToAccessProfile", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowUserToAccessProfile;

    /**
     * @return Allow persisted users to access their profile.  Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> allowUserToAccessProfile() {
        return Codegen.optional(this.allowUserToAccessProfile);
    }
    /**
     * Auto redirect to login through the IdP when clicking on Artifactory&#39;s login link.  Default value is `false`.
     * 
     */
    @Export(name="autoRedirect", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoRedirect;

    /**
     * @return Auto redirect to login through the IdP when clicking on Artifactory&#39;s login link.  Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> autoRedirect() {
        return Codegen.optional(this.autoRedirect);
    }
    /**
     * SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests. Default value is ``.
     * 
     */
    @Export(name="certificate", type=String.class, parameters={})
    private Output</* @Nullable */ String> certificate;

    /**
     * @return SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests. Default value is ``.
     * 
     */
    public Output<Optional<String>> certificate() {
        return Codegen.optional(this.certificate);
    }
    /**
     * Name of the attribute in the SAML response from the IdP that contains the user&#39;s email. Default value is ``.
     * 
     */
    @Export(name="emailAttribute", type=String.class, parameters={})
    private Output</* @Nullable */ String> emailAttribute;

    /**
     * @return Name of the attribute in the SAML response from the IdP that contains the user&#39;s email. Default value is ``.
     * 
     */
    public Output<Optional<String>> emailAttribute() {
        return Codegen.optional(this.emailAttribute);
    }
    /**
     * Enable SAML SSO.  Default value is `true`.
     * 
     */
    @Export(name="enable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enable;

    /**
     * @return Enable SAML SSO.  Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> enable() {
        return Codegen.optional(this.enable);
    }
    /**
     * Name of the attribute in the SAML response from the IdP that contains the user&#39;s group memberships. Default value is ``.
     * 
     */
    @Export(name="groupAttribute", type=String.class, parameters={})
    private Output</* @Nullable */ String> groupAttribute;

    /**
     * @return Name of the attribute in the SAML response from the IdP that contains the user&#39;s group memberships. Default value is ``.
     * 
     */
    public Output<Optional<String>> groupAttribute() {
        return Codegen.optional(this.groupAttribute);
    }
    /**
     * Service provider login url configured on the IdP.
     * 
     */
    @Export(name="loginUrl", type=String.class, parameters={})
    private Output<String> loginUrl;

    /**
     * @return Service provider login url configured on the IdP.
     * 
     */
    public Output<String> loginUrl() {
        return this.loginUrl;
    }
    /**
     * Service provider logout url, or where to redirect after user logs out.
     * 
     */
    @Export(name="logoutUrl", type=String.class, parameters={})
    private Output<String> logoutUrl;

    /**
     * @return Service provider logout url, or where to redirect after user logs out.
     * 
     */
    public Output<String> logoutUrl() {
        return this.logoutUrl;
    }
    /**
     * When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
     * 
     */
    @Export(name="noAutoUserCreation", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> noAutoUserCreation;

    /**
     * @return When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> noAutoUserCreation() {
        return Codegen.optional(this.noAutoUserCreation);
    }
    /**
     * The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or entity identity.
     * 
     */
    @Export(name="serviceProviderName", type=String.class, parameters={})
    private Output<String> serviceProviderName;

    /**
     * @return The SAML service provider name. This should be a URI that is also known as the entityID, providerID, or entity identity.
     * 
     */
    public Output<String> serviceProviderName() {
        return this.serviceProviderName;
    }
    /**
     * Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
     * 
     */
    @Export(name="syncGroups", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> syncGroups;

    /**
     * @return Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> syncGroups() {
        return Codegen.optional(this.syncGroups);
    }
    /**
     * When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it to your IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your SAML response. Default value is `false`.
     * 
     */
    @Export(name="useEncryptedAssertion", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> useEncryptedAssertion;

    /**
     * @return When set, an X.509 public certificate will be created by Artifactory. Download this certificate and upload it to your IDP and choose your own encryption algorithm. This process will let you encrypt the assertion section in your SAML response. Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> useEncryptedAssertion() {
        return Codegen.optional(this.useEncryptedAssertion);
    }
    /**
     * Enable &#34;audience&#34;, or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
     * 
     */
    @Export(name="verifyAudienceRestriction", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> verifyAudienceRestriction;

    /**
     * @return Enable &#34;audience&#34;, or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> verifyAudienceRestriction() {
        return Codegen.optional(this.verifyAudienceRestriction);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SamlSettings(String name) {
        this(name, SamlSettingsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SamlSettings(String name, SamlSettingsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SamlSettings(String name, SamlSettingsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/samlSettings:SamlSettings", name, args == null ? SamlSettingsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SamlSettings(String name, Output<String> id, @Nullable SamlSettingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/samlSettings:SamlSettings", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static SamlSettings get(String name, Output<String> id, @Nullable SamlSettingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SamlSettings(name, id, state, options);
    }
}
