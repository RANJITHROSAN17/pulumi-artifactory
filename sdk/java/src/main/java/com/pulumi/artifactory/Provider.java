// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.ProviderArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The provider type for the artifactory package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 * 
 */
@ResourceType(type="pulumi:providers:artifactory")
public class Provider extends com.pulumi.resources.ProviderResource {
    /**
     * This is a access token that can be given to you by your admin under `Identity and Access`. If not set, the &#39;api_key&#39;
     * attribute value will be used.
     * 
     */
    @Export(name="accessToken", type=String.class, parameters={})
    private Output</* @Nullable */ String> accessToken;

    /**
     * @return This is a access token that can be given to you by your admin under `Identity and Access`. If not set, the &#39;api_key&#39;
     * attribute value will be used.
     * 
     */
    public Output<Optional<String>> accessToken() {
        return Codegen.optional(this.accessToken);
    }
    /**
     * API token. Projects functionality will not work with any auth method other than access tokens
     * 
     * @deprecated
     * An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
     * In September 2022, the option to block the usage/creation of API Keys will be enabled by default, with the option for admins to change it back to enable API Keys.
     * In January 2023, API Keys will be deprecated all together and the option to use them will no longer be available.
     * 
     */
    @Deprecated /* An upcoming version will support the option to block the usage/creation of API Keys (for admins to set on their platform).
In September 2022, the option to block the usage/creation of API Keys will be enabled by default, with the option for admins to change it back to enable API Keys.
In January 2023, API Keys will be deprecated all together and the option to use them will no longer be available. */
    @Export(name="apiKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> apiKey;

    /**
     * @return API token. Projects functionality will not work with any auth method other than access tokens
     * 
     */
    public Output<Optional<String>> apiKey() {
        return Codegen.optional(this.apiKey);
    }
    @Export(name="url", type=String.class, parameters={})
    private Output</* @Nullable */ String> url;

    public Output<Optional<String>> url() {
        return Codegen.optional(this.url);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Provider(String name) {
        this(name, ProviderArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Provider(String name, @Nullable ProviderArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Provider(String name, @Nullable ProviderArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory", name, args == null ? ProviderArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

}
