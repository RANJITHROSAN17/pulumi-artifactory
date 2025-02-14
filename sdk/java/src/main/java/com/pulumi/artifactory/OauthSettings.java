// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.OauthSettingsArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.OauthSettingsState;
import com.pulumi.artifactory.outputs.OauthSettingsOauthProvider;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can be used to manage Artifactory&#39;s OAuth SSO settings.
 * 
 * Only a single `artifactory.OauthSettings` resource is meant to be defined.
 * 
 * ~&gt;The `artifactory.OauthSettings` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.OauthSettings;
 * import com.pulumi.artifactory.OauthSettingsArgs;
 * import com.pulumi.artifactory.inputs.OauthSettingsOauthProviderArgs;
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
 *         var oauth = new OauthSettings(&#34;oauth&#34;, OauthSettingsArgs.builder()        
 *             .allowUserToAccessProfile(true)
 *             .enable(true)
 *             .oauthProviders(OauthSettingsOauthProviderArgs.builder()
 *                 .apiUrl(&#34;https://organization.okta.com/oauth2/v1/userinfo&#34;)
 *                 .authUrl(&#34;https://organization.okta.com/oauth2/v1/authorize&#34;)
 *                 .clientId(&#34;foo&#34;)
 *                 .clientSecret(&#34;bar&#34;)
 *                 .enabled(false)
 *                 .name(&#34;okta&#34;)
 *                 .tokenUrl(&#34;https://organization.okta.com/oauth2/v1/token&#34;)
 *                 .type(&#34;openId&#34;)
 *                 .build())
 *             .persistUsers(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Current OAuth SSO settings can be imported using `oauth_settings` as the `ID`. If the resource is being imported, there will be a state drift, because `client_secret` can&#39;t be known. There are two options on how to approach this:
 * 
 * 1) Don&#39;t set `client_secret` initially, import, then update the config with actual secret; 2) Accept that there is a drift initially and run `terraform apply` twice;
 * 
 * ```sh
 *  $ pulumi import artifactory:index/oauthSettings:OauthSettings oauth oauth_settings
 * ```
 * 
 */
@ResourceType(type="artifactory:index/oauthSettings:OauthSettings")
public class OauthSettings extends com.pulumi.resources.CustomResource {
    /**
     * Allow persisted users to access their profile.  Default value is `false`.
     * 
     */
    @Export(name="allowUserToAccessProfile", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> allowUserToAccessProfile;

    /**
     * @return Allow persisted users to access their profile.  Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> allowUserToAccessProfile() {
        return Codegen.optional(this.allowUserToAccessProfile);
    }
    /**
     * Enable OAuth SSO.  Default value is `true`.
     * 
     */
    @Export(name="enable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enable;

    /**
     * @return Enable OAuth SSO.  Default value is `true`.
     * 
     */
    public Output<Optional<Boolean>> enable() {
        return Codegen.optional(this.enable);
    }
    /**
     * OAuth provider settings block. Multiple blocks can be defined, at least one is required.
     * 
     */
    @Export(name="oauthProviders", type=List.class, parameters={OauthSettingsOauthProvider.class})
    private Output<List<OauthSettingsOauthProvider>> oauthProviders;

    /**
     * @return OAuth provider settings block. Multiple blocks can be defined, at least one is required.
     * 
     */
    public Output<List<OauthSettingsOauthProvider>> oauthProviders() {
        return this.oauthProviders;
    }
    /**
     * Enable the creation of local Artifactory users.  Default value is `false`.
     * 
     */
    @Export(name="persistUsers", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> persistUsers;

    /**
     * @return Enable the creation of local Artifactory users.  Default value is `false`.
     * 
     */
    public Output<Optional<Boolean>> persistUsers() {
        return Codegen.optional(this.persistUsers);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OauthSettings(String name) {
        this(name, OauthSettingsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OauthSettings(String name, OauthSettingsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OauthSettings(String name, OauthSettingsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/oauthSettings:OauthSettings", name, args == null ? OauthSettingsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OauthSettings(String name, Output<String> id, @Nullable OauthSettingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/oauthSettings:OauthSettings", name, state, makeResourceOptions(options, id));
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
    public static OauthSettings get(String name, Output<String> id, @Nullable OauthSettingsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OauthSettings(name, id, state, options);
    }
}
