// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.LocalRpmRepositoryArgs;
import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.inputs.LocalRpmRepositoryState;
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
 * Creates a local RPM repository.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.artifactory.Keypair;
 * import com.pulumi.artifactory.KeypairArgs;
 * import com.pulumi.artifactory.LocalRpmRepository;
 * import com.pulumi.artifactory.LocalRpmRepositoryArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var some_keypair_gpg_1 = new Keypair(&#34;some-keypair-gpg-1&#34;, KeypairArgs.builder()        
 *             .pairName(String.format(&#34;some-keypair%s&#34;, random_id.randid().id()))
 *             .pairType(&#34;GPG&#34;)
 *             .alias(&#34;foo-alias1&#34;)
 *             .privateKey(Files.readString(Paths.get(&#34;samples/gpg.priv&#34;)))
 *             .publicKey(Files.readString(Paths.get(&#34;samples/gpg.pub&#34;)))
 *             .build());
 * 
 *         var some_keypair_gpg_2 = new Keypair(&#34;some-keypair-gpg-2&#34;, KeypairArgs.builder()        
 *             .pairName(String.format(&#34;some-keypair%s&#34;, random_id.randid().id()))
 *             .pairType(&#34;GPG&#34;)
 *             .alias(&#34;foo-alias2&#34;)
 *             .privateKey(Files.readString(Paths.get(&#34;samples/gpg.priv&#34;)))
 *             .publicKey(Files.readString(Paths.get(&#34;samples/gpg.pub&#34;)))
 *             .build());
 * 
 *         var terraform_local_test_rpm_repo_basic = new LocalRpmRepository(&#34;terraform-local-test-rpm-repo-basic&#34;, LocalRpmRepositoryArgs.builder()        
 *             .key(&#34;terraform-local-test-rpm-repo-basic&#34;)
 *             .yumRootDepth(5)
 *             .calculateYumMetadata(true)
 *             .enableFileListsIndexing(true)
 *             .yumGroupFileNames(&#34;file-1.xml,file-2.xml&#34;)
 *             .primaryKeypairRef(artifactory_keypair.some-keypairGPG1().pair_name())
 *             .secondaryKeypairRef(artifactory_keypair.some-keypairGPG2().pair_name())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     some_keypair_gpg_1,
 *                     some_keypair_gpg_2)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Local repositories can be imported using their name, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/localRpmRepository:LocalRpmRepository terraform-local-test-rpm-repo-basic terraform-local-test-rpm-repo-basic
 * ```
 * 
 */
@ResourceType(type="artifactory:index/localRpmRepository:LocalRpmRepository")
public class LocalRpmRepository extends com.pulumi.resources.CustomResource {
    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    @Export(name="archiveBrowsingEnabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> archiveBrowsingEnabled;

    /**
     * @return When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    public Output<Optional<Boolean>> archiveBrowsingEnabled() {
        return Codegen.optional(this.archiveBrowsingEnabled);
    }
    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    @Export(name="blackedOut", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> blackedOut;

    /**
     * @return When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    public Output<Optional<Boolean>> blackedOut() {
        return Codegen.optional(this.blackedOut);
    }
    /**
     * Default: false.
     * 
     */
    @Export(name="calculateYumMetadata", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> calculateYumMetadata;

    /**
     * @return Default: false.
     * 
     */
    public Output<Optional<Boolean>> calculateYumMetadata() {
        return Codegen.optional(this.calculateYumMetadata);
    }
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    @Export(name="downloadDirect", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> downloadDirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    public Output<Optional<Boolean>> downloadDirect() {
        return Codegen.optional(this.downloadDirect);
    }
    /**
     * Default: false.
     * 
     */
    @Export(name="enableFileListsIndexing", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enableFileListsIndexing;

    /**
     * @return Default: false.
     * 
     */
    public Output<Optional<Boolean>> enableFileListsIndexing() {
        return Codegen.optional(this.enableFileListsIndexing);
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", type=String.class, parameters={})
    private Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    public Output<String> excludesPattern() {
        return this.excludesPattern;
    }
    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Export(name="includesPattern", type=String.class, parameters={})
    private Output<String> includesPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Output<String> includesPattern() {
        return this.includesPattern;
    }
    /**
     * the identity key of the repo.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    @Export(name="notes", type=String.class, parameters={})
    private Output</* @Nullable */ String> notes;

    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    @Export(name="packageType", type=String.class, parameters={})
    private Output<String> packageType;

    public Output<String> packageType() {
        return this.packageType;
    }
    /**
     * The primary GPG key to be used to sign packages.
     * 
     */
    @Export(name="primaryKeypairRef", type=String.class, parameters={})
    private Output</* @Nullable */ String> primaryKeypairRef;

    /**
     * @return The primary GPG key to be used to sign packages.
     * 
     */
    public Output<Optional<String>> primaryKeypairRef() {
        return Codegen.optional(this.primaryKeypairRef);
    }
    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    @Export(name="priorityResolution", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> priorityResolution;

    /**
     * @return Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    public Output<Optional<Boolean>> priorityResolution() {
        return Codegen.optional(this.priorityResolution);
    }
    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    @Export(name="projectEnvironments", type=List.class, parameters={String.class})
    private Output<List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;
     * 
     */
    public Output<List<String>> projectEnvironments() {
        return this.projectEnvironments;
    }
    /**
     * Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
     * with project key, separated by a dash.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
    }
    /**
     * List of property set name
     * 
     */
    @Export(name="propertySets", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> propertySets;

    /**
     * @return List of property set name
     * 
     */
    public Output<Optional<List<String>>> propertySets() {
        return Codegen.optional(this.propertySets);
    }
    /**
     * Repository layout key for the local repository
     * 
     */
    @Export(name="repoLayoutRef", type=String.class, parameters={})
    private Output</* @Nullable */ String> repoLayoutRef;

    /**
     * @return Repository layout key for the local repository
     * 
     */
    public Output<Optional<String>> repoLayoutRef() {
        return Codegen.optional(this.repoLayoutRef);
    }
    /**
     * The secondary GPG key to be used to sign packages.
     * 
     */
    @Export(name="secondaryKeypairRef", type=String.class, parameters={})
    private Output</* @Nullable */ String> secondaryKeypairRef;

    /**
     * @return The secondary GPG key to be used to sign packages.
     * 
     */
    public Output<Optional<String>> secondaryKeypairRef() {
        return Codegen.optional(this.secondaryKeypairRef);
    }
    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    @Export(name="xrayIndex", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> xrayIndex;

    /**
     * @return Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    public Output<Optional<Boolean>> xrayIndex() {
        return Codegen.optional(this.xrayIndex);
    }
    /**
     * A comma separated list of XML file names containing RPM group component definitions.
     * Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
     * generating a gzipped version of the group files, if required. Default is empty string.
     * 
     */
    @Export(name="yumGroupFileNames", type=String.class, parameters={})
    private Output</* @Nullable */ String> yumGroupFileNames;

    /**
     * @return A comma separated list of XML file names containing RPM group component definitions.
     * Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
     * generating a gzipped version of the group files, if required. Default is empty string.
     * 
     */
    public Output<Optional<String>> yumGroupFileNames() {
        return Codegen.optional(this.yumGroupFileNames);
    }
    /**
     * The depth, relative to the repository&#39;s root folder, where RPM metadata is created.
     * This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
     * your RPMs are stored under &#39;fedora/linux/$releasever/$basearch&#39;, specify a depth of 4. Once the number of snapshots
     * exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
     * snapshots are not cleaned up.
     * 
     */
    @Export(name="yumRootDepth", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> yumRootDepth;

    /**
     * @return The depth, relative to the repository&#39;s root folder, where RPM metadata is created.
     * This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
     * your RPMs are stored under &#39;fedora/linux/$releasever/$basearch&#39;, specify a depth of 4. Once the number of snapshots
     * exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
     * snapshots are not cleaned up.
     * 
     */
    public Output<Optional<Integer>> yumRootDepth() {
        return Codegen.optional(this.yumRootDepth);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LocalRpmRepository(String name) {
        this(name, LocalRpmRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LocalRpmRepository(String name, LocalRpmRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LocalRpmRepository(String name, LocalRpmRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/localRpmRepository:LocalRpmRepository", name, args == null ? LocalRpmRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LocalRpmRepository(String name, Output<String> id, @Nullable LocalRpmRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/localRpmRepository:LocalRpmRepository", name, state, makeResourceOptions(options, id));
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
    public static LocalRpmRepository get(String name, Output<String> id, @Nullable LocalRpmRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LocalRpmRepository(name, id, state, options);
    }
}
