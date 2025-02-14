// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.DistributionWebhookArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.DistributionWebhookState;
import com.pulumi.artifactory.outputs.DistributionWebhookCriteria;
import com.pulumi.artifactory.outputs.DistributionWebhookHandler;
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
 * Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
 * 
 * ## Example Usage
 * 
 * .
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.DistributionWebhook;
 * import com.pulumi.artifactory.DistributionWebhookArgs;
 * import com.pulumi.artifactory.inputs.DistributionWebhookCriteriaArgs;
 * import com.pulumi.artifactory.inputs.DistributionWebhookHandlerArgs;
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
 *         var distribution_webhook = new DistributionWebhook(&#34;distribution-webhook&#34;, DistributionWebhookArgs.builder()        
 *             .criteria(DistributionWebhookCriteriaArgs.builder()
 *                 .anyReleaseBundle(false)
 *                 .excludePatterns(&#34;bar/**&#34;)
 *                 .includePatterns(&#34;foo/**&#34;)
 *                 .registeredReleaseBundleNames(&#34;bundle-name&#34;)
 *                 .build())
 *             .eventTypes(            
 *                 &#34;distribute_started&#34;,
 *                 &#34;distribute_completed&#34;,
 *                 &#34;distribute_aborted&#34;,
 *                 &#34;distribute_failed&#34;,
 *                 &#34;delete_started&#34;,
 *                 &#34;delete_completed&#34;,
 *                 &#34;delete_failed&#34;)
 *             .handlers(DistributionWebhookHandlerArgs.builder()
 *                 .customHttpHeaders(Map.ofEntries(
 *                     Map.entry(&#34;header-1&#34;, &#34;value-1&#34;),
 *                     Map.entry(&#34;header-2&#34;, &#34;value-2&#34;)
 *                 ))
 *                 .proxy(&#34;proxy-key&#34;)
 *                 .secret(&#34;some-secret&#34;)
 *                 .url(&#34;http://tempurl.org/webhook&#34;)
 *                 .build())
 *             .key(&#34;distribution-webhook&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="artifactory:index/distributionWebhook:DistributionWebhook")
public class DistributionWebhook extends com.pulumi.resources.CustomResource {
    /**
     * Specifies where the webhook will be applied on which repositories.
     * 
     */
    @Export(name="criteria", type=DistributionWebhookCriteria.class, parameters={})
    private Output<DistributionWebhookCriteria> criteria;

    /**
     * @return Specifies where the webhook will be applied on which repositories.
     * 
     */
    public Output<DistributionWebhookCriteria> criteria() {
        return this.criteria;
    }
    /**
     * Webhook description. Max length 1000 characters.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Webhook description. Max length 1000 characters.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Status of webhook. Default to `true`.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Status of webhook. Default to `true`.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distribute_started`, `distribute_completed`, `distribute_aborted`, ` distribute_failed,  `delete_started` ,  `delete_completed` ,  `delete_failed`
     * 
     */
    @Export(name="eventTypes", type=List.class, parameters={String.class})
    private Output<List<String>> eventTypes;

    /**
     * @return List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `distribute_started`, `distribute_completed`, `distribute_aborted`, ` distribute_failed,  `delete_started` ,  `delete_completed` ,  `delete_failed`
     * 
     */
    public Output<List<String>> eventTypes() {
        return this.eventTypes;
    }
    /**
     * At least one is required.
     * 
     */
    @Export(name="handlers", type=List.class, parameters={DistributionWebhookHandler.class})
    private Output<List<DistributionWebhookHandler>> handlers;

    /**
     * @return At least one is required.
     * 
     */
    public Output<List<DistributionWebhookHandler>> handlers() {
        return this.handlers;
    }
    /**
     * The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DistributionWebhook(String name) {
        this(name, DistributionWebhookArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DistributionWebhook(String name, DistributionWebhookArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DistributionWebhook(String name, DistributionWebhookArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/distributionWebhook:DistributionWebhook", name, args == null ? DistributionWebhookArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DistributionWebhook(String name, Output<String> id, @Nullable DistributionWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/distributionWebhook:DistributionWebhook", name, state, makeResourceOptions(options, id));
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
    public static DistributionWebhook get(String name, Output<String> id, @Nullable DistributionWebhookState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DistributionWebhook(name, id, state, options);
    }
}
