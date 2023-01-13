// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.Utilities;
import com.pulumi.artifactory.VirtualGitlfsRepositoryArgs;
import com.pulumi.artifactory.inputs.VirtualGitlfsRepositoryState;
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
 * Creates a virtual Git LFS repository.
 * Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Git+LFS+Repositories#GitLFSRepositories-VirtualRepositories).
 * 
 * ## Import
 * 
 * Virtual repositories can be imported using their name, e.g.
 * 
 * ```sh
 *  $ pulumi import artifactory:index/virtualGitlfsRepository:VirtualGitlfsRepository foo-gitlfs foo-gitlfs
 * ```
 * 
 */
@ResourceType(type="artifactory:index/virtualGitlfsRepository:VirtualGitlfsRepository")
public class VirtualGitlfsRepository extends com.pulumi.resources.CustomResource {
    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    @Export(name="artifactoryRequestsCanRetrieveRemoteArtifacts", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts;

    /**
     * @return Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    public Output<Optional<Boolean>> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Codegen.optional(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }
    /**
     * Default repository to deploy artifacts.
     * 
     */
    @Export(name="defaultDeploymentRepo", type=String.class, parameters={})
    private Output</* @Nullable */ String> defaultDeploymentRepo;

    /**
     * @return Default repository to deploy artifacts.
     * 
     */
    public Output<Optional<String>> defaultDeploymentRepo() {
        return Codegen.optional(this.defaultDeploymentRepo);
    }
    /**
     * A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
     * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
     * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    @Export(name="excludesPattern", type=String.class, parameters={})
    private Output</* @Nullable */ String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    public Output<Optional<String>> excludesPattern() {
        return Codegen.optional(this.excludesPattern);
    }
    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Export(name="includesPattern", type=String.class, parameters={})
    private Output</* @Nullable */ String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Output<Optional<String>> includesPattern() {
        return Codegen.optional(this.includesPattern);
    }
    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    @Export(name="key", type=String.class, parameters={})
    private Output<String> key;

    /**
     * @return A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    public Output<String> key() {
        return this.key;
    }
    /**
     * A free text field to add additional notes about the repository. These are only visible to the administrator.
     * 
     */
    @Export(name="notes", type=String.class, parameters={})
    private Output</* @Nullable */ String> notes;

    /**
     * @return A free text field to add additional notes about the repository. These are only visible to the administrator.
     * 
     */
    public Output<Optional<String>> notes() {
        return Codegen.optional(this.notes);
    }
    /**
     * The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
     * 
     */
    @Export(name="packageType", type=String.class, parameters={})
    private Output<String> packageType;

    /**
     * @return The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
     * 
     */
    public Output<String> packageType() {
        return this.packageType;
    }
    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;. The attribute should only be used
     * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
     * will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    @Export(name="projectEnvironments", type=List.class, parameters={String.class})
    private Output<List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;. The attribute should only be used
     * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
     * will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    public Output<List<String>> projectEnvironments() {
        return this.projectEnvironments;
    }
    /**
     * Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Export(name="projectKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Output<Optional<String>> projectKey() {
        return Codegen.optional(this.projectKey);
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
     * The effective list of actual repositories included in this virtual repository.
     * 
     */
    @Export(name="repositories", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> repositories;

    /**
     * @return The effective list of actual repositories included in this virtual repository.
     * 
     */
    public Output<Optional<List<String>>> repositories() {
        return Codegen.optional(this.repositories);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VirtualGitlfsRepository(String name) {
        this(name, VirtualGitlfsRepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualGitlfsRepository(String name, VirtualGitlfsRepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualGitlfsRepository(String name, VirtualGitlfsRepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/virtualGitlfsRepository:VirtualGitlfsRepository", name, args == null ? VirtualGitlfsRepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VirtualGitlfsRepository(String name, Output<String> id, @Nullable VirtualGitlfsRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("artifactory:index/virtualGitlfsRepository:VirtualGitlfsRepository", name, state, makeResourceOptions(options, id));
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
    public static VirtualGitlfsRepository get(String name, Output<String> id, @Nullable VirtualGitlfsRepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualGitlfsRepository(name, id, state, options);
    }
}
