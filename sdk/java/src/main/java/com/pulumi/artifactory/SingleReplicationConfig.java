// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.SingleReplicationConfigArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.SingleReplicationConfigState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Note: this resource is deprecated in favor of `artifactory.PullReplication` resource.
 * 
 * Provides an Artifactory single replication config resource. This can be used to create and manage a single Artifactory
 * replication. Primarily used when pull replication is needed.
 * 
 * **WARNING: This should not be used on a repository with `artifactory.ReplicationConfig`. Using both together will cause
 * unexpected behaviour and will almost certainly cause your replications to break.**
 * 
 * ## Example Usage
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.LocalMavenRepository;
 * import com.pulumi.artifactory.LocalMavenRepositoryArgs;
 * import com.pulumi.artifactory.SingleReplicationConfig;
 * import com.pulumi.artifactory.SingleReplicationConfigArgs;
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
 *         var providerTestSource = new LocalMavenRepository(&#34;providerTestSource&#34;, LocalMavenRepositoryArgs.builder()        
 *             .key(&#34;provider_test_source&#34;)
 *             .build());
 * 
 *         var providerTestDest = new LocalMavenRepository(&#34;providerTestDest&#34;, LocalMavenRepositoryArgs.builder()        
 *             .key(&#34;provider_test_dest&#34;)
 *             .build());
 * 
 *         var foo_rep = new SingleReplicationConfig(&#34;foo-rep&#34;, SingleReplicationConfigArgs.builder()        
 *             .cronExp(&#34;0 0 * * * ?&#34;)
 *             .enableEventReplication(true)
 *             .password(var_.artifactory_password())
 *             .repoKey(providerTestSource.key())
 *             .url(var_.artifactory_url())
 *             .username(var_.artifactory_username())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Replication configs can be imported using their repo key, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/singleReplicationConfig:SingleReplicationConfig foo-rep repository-key
 * ```
 * 
 */
@ResourceType(type="artifactory:index/singleReplicationConfig:SingleReplicationConfig")
public class SingleReplicationConfig extends com.pulumi.resources.CustomResource {
    /**
     * Cron expression to control the operation frequency.
     * 
     */
    @Export(name="cronExp", type=String.class, parameters={})
    private Output<String> cronExp;

    /**
     * @return Cron expression to control the operation frequency.
     * 
     */
    public Output<String> cronExp() {
        return this.cronExp;
    }
    @Export(name="enableEventReplication", type=Boolean.class, parameters={})
    private Output<Boolean> enableEventReplication;

    public Output<Boolean> enableEventReplication() {
        return this.enableEventReplication;
    }
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output<Boolean> enabled;

    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output<String> password;

    /**
     * @return Requires password encryption to be turned off `POST /api/system/decrypt`.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    @Export(name="pathPrefix", type=String.class, parameters={})
    private Output</* @Nullable */ String> pathPrefix;

    public Output<Optional<String>> pathPrefix() {
        return Codegen.optional(this.pathPrefix);
    }
    /**
     * Proxy key from Artifactory Proxies setting.
     * 
     */
    @Export(name="proxy", type=String.class, parameters={})
    private Output</* @Nullable */ String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies setting.
     * 
     */
    public Output<Optional<String>> proxy() {
        return Codegen.optional(this.proxy);
    }
    @Export(name="repoKey", type=String.class, parameters={})
    private Output<String> repoKey;

    public Output<String> repoKey() {
        return this.repoKey;
    }
    @Export(name="socketTimeoutMillis", type=Integer.class, parameters={})
    private Output<Integer> socketTimeoutMillis;

    public Output<Integer> socketTimeoutMillis() {
        return this.socketTimeoutMillis;
    }
    @Export(name="syncDeletes", type=Boolean.class, parameters={})
    private Output<Boolean> syncDeletes;

    public Output<Boolean> syncDeletes() {
        return this.syncDeletes;
    }
    @Export(name="syncProperties", type=Boolean.class, parameters={})
    private Output<Boolean> syncProperties;

    public Output<Boolean> syncProperties() {
        return this.syncProperties;
    }
    @Export(name="syncStatistics", type=Boolean.class, parameters={})
    private Output<Boolean> syncStatistics;

    public Output<Boolean> syncStatistics() {
        return this.syncStatistics;
    }
    @Export(name="url", type=String.class, parameters={})
    private Output</* @Nullable */ String> url;

    public Output<Optional<String>> url() {
        return Codegen.optional(this.url);
    }
    @Export(name="username", type=String.class, parameters={})
    private Output</* @Nullable */ String> username;

    public Output<Optional<String>> username() {
        return Codegen.optional(this.username);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SingleReplicationConfig(String name) {
        this(name, SingleReplicationConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SingleReplicationConfig(String name, SingleReplicationConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SingleReplicationConfig(String name, SingleReplicationConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/singleReplicationConfig:SingleReplicationConfig", name, args == null ? SingleReplicationConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SingleReplicationConfig(String name, Output<String> id, @Nullable SingleReplicationConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/singleReplicationConfig:SingleReplicationConfig", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
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
    public static SingleReplicationConfig get(String name, Output<String> id, @Nullable SingleReplicationConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SingleReplicationConfig(name, id, state, options);
    }
}
