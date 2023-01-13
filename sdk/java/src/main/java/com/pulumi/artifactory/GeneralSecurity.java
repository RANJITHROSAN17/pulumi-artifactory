// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.GeneralSecurityArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.GeneralSecurityState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can be used to manage Artifactory&#39;s general security settings.
 * 
 * Only a single `artifactory.GeneralSecurity` resource is meant to be defined.
 * 
 * ~&gt;The `artifactory.GeneralSecurity` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.GeneralSecurity;
 * import com.pulumi.artifactory.GeneralSecurityArgs;
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
 *         var security = new GeneralSecurity(&#34;security&#34;, GeneralSecurityArgs.builder()        
 *             .enableAnonymousAccess(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Current general security settings can be imported using `security` as the `ID`, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/generalSecurity:GeneralSecurity security security
 * ```
 * 
 *  ~&gt;The `artifactory_general_security` resource uses endpoints that are undocumented and may not work with SaaS environments, or may change without notice.
 * 
 */
@ResourceType(type="artifactory:index/generalSecurity:GeneralSecurity")
public class GeneralSecurity extends com.pulumi.resources.CustomResource {
    @Export(name="enableAnonymousAccess", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enableAnonymousAccess;

    public Output<Optional<Boolean>> enableAnonymousAccess() {
        return Codegen.optional(this.enableAnonymousAccess);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GeneralSecurity(String name) {
        this(name, GeneralSecurityArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GeneralSecurity(String name, @Nullable GeneralSecurityArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GeneralSecurity(String name, @Nullable GeneralSecurityArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/generalSecurity:GeneralSecurity", name, args == null ? GeneralSecurityArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GeneralSecurity(String name, Output<String> id, @Nullable GeneralSecurityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/generalSecurity:GeneralSecurity", name, state, makeResourceOptions(options, id));
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
    public static GeneralSecurity get(String name, Output<String> id, @Nullable GeneralSecurityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GeneralSecurity(name, id, state, options);
    }
}
