// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VirtualSbtRepositoryArgs extends com.pulumi.resources.ResourceArgs {

    public static final VirtualSbtRepositoryArgs Empty = new VirtualSbtRepositoryArgs();

    /**
     * Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    @Import(name="artifactoryRequestsCanRetrieveRemoteArtifacts")
    private @Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts;

    /**
     * @return Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
     * another Artifactory instance.
     * 
     */
    public Optional<Output<Boolean>> artifactoryRequestsCanRetrieveRemoteArtifacts() {
        return Optional.ofNullable(this.artifactoryRequestsCanRetrieveRemoteArtifacts);
    }

    /**
     * Default repository to deploy artifacts.
     * 
     */
    @Import(name="defaultDeploymentRepo")
    private @Nullable Output<String> defaultDeploymentRepo;

    /**
     * @return Default repository to deploy artifacts.
     * 
     */
    public Optional<Output<String>> defaultDeploymentRepo() {
        return Optional.ofNullable(this.defaultDeploymentRepo);
    }

    /**
     * A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
     * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
     * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
     * artifacts are excluded.
     * 
     */
    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    /**
     * User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
     * is also enforced when aggregated repositories support anonymous requests.
     * 
     */
    @Import(name="forceMavenAuthentication")
    private @Nullable Output<Boolean> forceMavenAuthentication;

    /**
     * @return User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
     * is also enforced when aggregated repositories support anonymous requests.
     * 
     */
    public Optional<Output<Boolean>> forceMavenAuthentication() {
        return Optional.ofNullable(this.forceMavenAuthentication);
    }

    /**
     * List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Import(name="includesPattern")
    private @Nullable Output<String> includesPattern;

    /**
     * @return List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
     * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Optional<Output<String>> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    /**
     * A mandatory identifier for the repository that must be unique. It cannot begin with a number or
     * contain spaces or special characters.
     * 
     */
    @Import(name="key", required=true)
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
     * The keypair used to sign artifacts.
     * 
     */
    @Import(name="keyPair")
    private @Nullable Output<String> keyPair;

    /**
     * @return The keypair used to sign artifacts.
     * 
     */
    public Optional<Output<String>> keyPair() {
        return Optional.ofNullable(this.keyPair);
    }

    /**
     * A free text field to add additional notes about the repository. These are only visible to the administrator.
     * 
     */
    @Import(name="notes")
    private @Nullable Output<String> notes;

    /**
     * @return A free text field to add additional notes about the repository. These are only visible to the administrator.
     * 
     */
    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    /**
     * - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
     * - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
     * - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
     * 
     */
    @Import(name="pomRepositoryReferencesCleanupPolicy")
    private @Nullable Output<String> pomRepositoryReferencesCleanupPolicy;

    /**
     * @return - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
     * - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
     * - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
     * 
     */
    public Optional<Output<String>> pomRepositoryReferencesCleanupPolicy() {
        return Optional.ofNullable(this.pomRepositoryReferencesCleanupPolicy);
    }

    /**
     * Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;. The attribute should only be used
     * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
     * will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    @Import(name="projectEnvironments")
    private @Nullable Output<List<String>> projectEnvironments;

    /**
     * @return Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;. The attribute should only be used
     * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
     * will remain in the Terraform state, which will create state drift during the update.
     * 
     */
    public Optional<Output<List<String>>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    /**
     * Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    /**
     * @return Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
     * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
     * 
     */
    public Optional<Output<String>> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    /**
     * Repository layout key for the local repository
     * 
     */
    @Import(name="repoLayoutRef")
    private @Nullable Output<String> repoLayoutRef;

    /**
     * @return Repository layout key for the local repository
     * 
     */
    public Optional<Output<String>> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    /**
     * The effective list of actual repositories included in this virtual repository.
     * 
     */
    @Import(name="repositories")
    private @Nullable Output<List<String>> repositories;

    /**
     * @return The effective list of actual repositories included in this virtual repository.
     * 
     */
    public Optional<Output<List<String>>> repositories() {
        return Optional.ofNullable(this.repositories);
    }

    private VirtualSbtRepositoryArgs() {}

    private VirtualSbtRepositoryArgs(VirtualSbtRepositoryArgs $) {
        this.artifactoryRequestsCanRetrieveRemoteArtifacts = $.artifactoryRequestsCanRetrieveRemoteArtifacts;
        this.defaultDeploymentRepo = $.defaultDeploymentRepo;
        this.description = $.description;
        this.excludesPattern = $.excludesPattern;
        this.forceMavenAuthentication = $.forceMavenAuthentication;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.keyPair = $.keyPair;
        this.notes = $.notes;
        this.pomRepositoryReferencesCleanupPolicy = $.pomRepositoryReferencesCleanupPolicy;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.repoLayoutRef = $.repoLayoutRef;
        this.repositories = $.repositories;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VirtualSbtRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VirtualSbtRepositoryArgs $;

        public Builder() {
            $ = new VirtualSbtRepositoryArgs();
        }

        public Builder(VirtualSbtRepositoryArgs defaults) {
            $ = new VirtualSbtRepositoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param artifactoryRequestsCanRetrieveRemoteArtifacts Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
         * another Artifactory instance.
         * 
         * @return builder
         * 
         */
        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(@Nullable Output<Boolean> artifactoryRequestsCanRetrieveRemoteArtifacts) {
            $.artifactoryRequestsCanRetrieveRemoteArtifacts = artifactoryRequestsCanRetrieveRemoteArtifacts;
            return this;
        }

        /**
         * @param artifactoryRequestsCanRetrieveRemoteArtifacts Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
         * another Artifactory instance.
         * 
         * @return builder
         * 
         */
        public Builder artifactoryRequestsCanRetrieveRemoteArtifacts(Boolean artifactoryRequestsCanRetrieveRemoteArtifacts) {
            return artifactoryRequestsCanRetrieveRemoteArtifacts(Output.of(artifactoryRequestsCanRetrieveRemoteArtifacts));
        }

        /**
         * @param defaultDeploymentRepo Default repository to deploy artifacts.
         * 
         * @return builder
         * 
         */
        public Builder defaultDeploymentRepo(@Nullable Output<String> defaultDeploymentRepo) {
            $.defaultDeploymentRepo = defaultDeploymentRepo;
            return this;
        }

        /**
         * @param defaultDeploymentRepo Default repository to deploy artifacts.
         * 
         * @return builder
         * 
         */
        public Builder defaultDeploymentRepo(String defaultDeploymentRepo) {
            return defaultDeploymentRepo(Output.of(defaultDeploymentRepo));
        }

        /**
         * @param description A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
         * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
         * field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
         * artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludesPattern(@Nullable Output<String> excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        /**
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*.By default no
         * artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
        }

        /**
         * @param forceMavenAuthentication User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
         * is also enforced when aggregated repositories support anonymous requests.
         * 
         * @return builder
         * 
         */
        public Builder forceMavenAuthentication(@Nullable Output<Boolean> forceMavenAuthentication) {
            $.forceMavenAuthentication = forceMavenAuthentication;
            return this;
        }

        /**
         * @param forceMavenAuthentication User authentication is required when accessing the repository. An anonymous request will display an HTTP 401 error. This
         * is also enforced when aggregated repositories support anonymous requests.
         * 
         * @return builder
         * 
         */
        public Builder forceMavenAuthentication(Boolean forceMavenAuthentication) {
            return forceMavenAuthentication(Output.of(forceMavenAuthentication));
        }

        /**
         * @param includesPattern List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
         * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
         * 
         * @return builder
         * 
         */
        public Builder includesPattern(@Nullable Output<String> includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        /**
         * @param includesPattern List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When
         * used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
         * 
         * @return builder
         * 
         */
        public Builder includesPattern(String includesPattern) {
            return includesPattern(Output.of(includesPattern));
        }

        /**
         * @param key A mandatory identifier for the repository that must be unique. It cannot begin with a number or
         * contain spaces or special characters.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key A mandatory identifier for the repository that must be unique. It cannot begin with a number or
         * contain spaces or special characters.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param keyPair The keypair used to sign artifacts.
         * 
         * @return builder
         * 
         */
        public Builder keyPair(@Nullable Output<String> keyPair) {
            $.keyPair = keyPair;
            return this;
        }

        /**
         * @param keyPair The keypair used to sign artifacts.
         * 
         * @return builder
         * 
         */
        public Builder keyPair(String keyPair) {
            return keyPair(Output.of(keyPair));
        }

        /**
         * @param notes A free text field to add additional notes about the repository. These are only visible to the administrator.
         * 
         * @return builder
         * 
         */
        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        /**
         * @param notes A free text field to add additional notes about the repository. These are only visible to the administrator.
         * 
         * @return builder
         * 
         */
        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        /**
         * @param pomRepositoryReferencesCleanupPolicy - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
         * - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
         * - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
         * 
         * @return builder
         * 
         */
        public Builder pomRepositoryReferencesCleanupPolicy(@Nullable Output<String> pomRepositoryReferencesCleanupPolicy) {
            $.pomRepositoryReferencesCleanupPolicy = pomRepositoryReferencesCleanupPolicy;
            return this;
        }

        /**
         * @param pomRepositoryReferencesCleanupPolicy - (1: discard_active_reference) Discard Active References - Removes repository elements that are declared directly under project or under a profile in the same POM that is activeByDefault.
         * - (2: discard_any_reference) Discard Any References - Removes all repository elements regardless of whether they are included in an active profile or not.
         * - (3: nothing) Nothing - Does not remove any repository elements declared in the POM.
         * 
         * @return builder
         * 
         */
        public Builder pomRepositoryReferencesCleanupPolicy(String pomRepositoryReferencesCleanupPolicy) {
            return pomRepositoryReferencesCleanupPolicy(Output.of(pomRepositoryReferencesCleanupPolicy));
        }

        /**
         * @param projectEnvironments Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;. The attribute should only be used
         * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
         * will remain in the Terraform state, which will create state drift during the update.
         * 
         * @return builder
         * 
         */
        public Builder projectEnvironments(@Nullable Output<List<String>> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        /**
         * @param projectEnvironments Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;. The attribute should only be used
         * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
         * will remain in the Terraform state, which will create state drift during the update.
         * 
         * @return builder
         * 
         */
        public Builder projectEnvironments(List<String> projectEnvironments) {
            return projectEnvironments(Output.of(projectEnvironments));
        }

        /**
         * @param projectEnvironments Project environment for assigning this repository to. Allow values: &#34;DEV&#34; or &#34;PROD&#34;. The attribute should only be used
         * if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
         * will remain in the Terraform state, which will create state drift during the update.
         * 
         * @return builder
         * 
         */
        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        /**
         * @param projectKey Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
         * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(@Nullable Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        /**
         * @param projectKey Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
         * assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
         * 
         * @return builder
         * 
         */
        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
        }

        /**
         * @param repoLayoutRef Repository layout key for the local repository
         * 
         * @return builder
         * 
         */
        public Builder repoLayoutRef(@Nullable Output<String> repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        /**
         * @param repoLayoutRef Repository layout key for the local repository
         * 
         * @return builder
         * 
         */
        public Builder repoLayoutRef(String repoLayoutRef) {
            return repoLayoutRef(Output.of(repoLayoutRef));
        }

        /**
         * @param repositories The effective list of actual repositories included in this virtual repository.
         * 
         * @return builder
         * 
         */
        public Builder repositories(@Nullable Output<List<String>> repositories) {
            $.repositories = repositories;
            return this;
        }

        /**
         * @param repositories The effective list of actual repositories included in this virtual repository.
         * 
         * @return builder
         * 
         */
        public Builder repositories(List<String> repositories) {
            return repositories(Output.of(repositories));
        }

        /**
         * @param repositories The effective list of actual repositories included in this virtual repository.
         * 
         * @return builder
         * 
         */
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }

        public VirtualSbtRepositoryArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}
