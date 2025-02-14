// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVirtualBowerRepositoryPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVirtualBowerRepositoryPlainArgs Empty = new GetVirtualBowerRepositoryPlainArgs();

    @Import(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    private @Nullable Boolean artifactoryRequestsCanRetrieveRemoteArtifacts;

    public Optional<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Optional.ofNullable(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }

    @Import(name="defaultDeploymentRepo")
    private @Nullable String defaultDeploymentRepo;

    public Optional<String> defaultDeploymentRepo() {
        return Optional.ofNullable(this.defaultDeploymentRepo);
    }

    @Import(name="description")
    private @Nullable String description;

    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="excludesPattern")
    private @Nullable String excludesPattern;

    public Optional<String> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    /**
     * (Optional) When set, external dependencies are rewritten. Default value is false.
     * 
     */
    @Import(name="externalDependenciesEnabled")
    private @Nullable Boolean externalDependenciesEnabled;

    /**
     * @return (Optional) When set, external dependencies are rewritten. Default value is false.
     * 
     */
    public Optional<Boolean> externalDependenciesEnabled() {
        return Optional.ofNullable(this.externalDependenciesEnabled);
    }

    /**
     * (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
     * 
     */
    @Import(name="externalDependenciesPatterns")
    private @Nullable List<String> externalDependenciesPatterns;

    /**
     * @return (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
     * 
     */
    public Optional<List<String>> externalDependenciesPatterns() {
        return Optional.ofNullable(this.externalDependenciesPatterns);
    }

    /**
     * (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
     * 
     */
    @Import(name="externalDependenciesRemoteRepo")
    private @Nullable String externalDependenciesRemoteRepo;

    /**
     * @return (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
     * 
     */
    public Optional<String> externalDependenciesRemoteRepo() {
        return Optional.ofNullable(this.externalDependenciesRemoteRepo);
    }

    @Import(name="includesPattern")
    private @Nullable String includesPattern;

    public Optional<String> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    /**
     * the identity key of the repo.
     * 
     */
    @Import(name="key", required=true)
    private String key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public String key() {
        return this.key;
    }

    @Import(name="notes")
    private @Nullable String notes;

    public Optional<String> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="projectEnvironments")
    private @Nullable List<String> projectEnvironments;

    public Optional<List<String>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    @Import(name="projectKey")
    private @Nullable String projectKey;

    public Optional<String> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    @Import(name="repoLayoutRef")
    private @Nullable String repoLayoutRef;

    public Optional<String> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    @Import(name="repositories")
    private @Nullable List<String> repositories;

    public Optional<List<String>> repositories() {
        return Optional.ofNullable(this.repositories);
    }

    private GetVirtualBowerRepositoryPlainArgs() {}

    private GetVirtualBowerRepositoryPlainArgs(GetVirtualBowerRepositoryPlainArgs $) {
        this.artifactoryRequestsCanRetrieveRemoteArtifacts = $.artifactoryRequestsCanRetrieveRemoteArtifacts;
        this.defaultDeploymentRepo = $.defaultDeploymentRepo;
        this.description = $.description;
        this.excludesPattern = $.excludesPattern;
        this.externalDependenciesEnabled = $.externalDependenciesEnabled;
        this.externalDependenciesPatterns = $.externalDependenciesPatterns;
        this.externalDependenciesRemoteRepo = $.externalDependenciesRemoteRepo;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.notes = $.notes;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.repoLayoutRef = $.repoLayoutRef;
        this.repositories = $.repositories;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVirtualBowerRepositoryPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVirtualBowerRepositoryPlainArgs $;

        public Builder() {
            $ = new GetVirtualBowerRepositoryPlainArgs();
        }

        public Builder(GetVirtualBowerRepositoryPlainArgs defaults) {
            $ = new GetVirtualBowerRepositoryPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(@Nullable Boolean artifactoryRequestsCanRetrieveRemoteArtifacts) {
            $.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            return this;
        }

        public Builder defaultDeploymentRepo(@Nullable String defaultDeploymentRepo) {
            $.defaultDeploymentRepo = defaultDeploymentRepo;
            return this;
        }

        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        public Builder excludesPattern(@Nullable String excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        /**
         * @param externalDependenciesEnabled (Optional) When set, external dependencies are rewritten. Default value is false.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesEnabled(@Nullable Boolean externalDependenciesEnabled) {
            $.externalDependenciesEnabled = externalDependenciesEnabled;
            return this;
        }

        /**
         * @param externalDependenciesPatterns (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesPatterns(@Nullable List<String> externalDependenciesPatterns) {
            $.externalDependenciesPatterns = externalDependenciesPatterns;
            return this;
        }

        /**
         * @param externalDependenciesPatterns (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesPatterns(String... externalDependenciesPatterns) {
            return externalDependenciesPatterns(List.of(externalDependenciesPatterns));
        }

        /**
         * @param externalDependenciesRemoteRepo (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
         * 
         * @return builder
         * 
         */
        public Builder externalDependenciesRemoteRepo(@Nullable String externalDependenciesRemoteRepo) {
            $.externalDependenciesRemoteRepo = externalDependenciesRemoteRepo;
            return this;
        }

        public Builder includesPattern(@Nullable String includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            $.key = key;
            return this;
        }

        public Builder notes(@Nullable String notes) {
            $.notes = notes;
            return this;
        }

        public Builder projectEnvironments(@Nullable List<String> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        public Builder projectKey(@Nullable String projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        public Builder repoLayoutRef(@Nullable String repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        public Builder repositories(@Nullable List<String> repositories) {
            $.repositories = repositories;
            return this;
        }

        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }

        public GetVirtualBowerRepositoryPlainArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}
