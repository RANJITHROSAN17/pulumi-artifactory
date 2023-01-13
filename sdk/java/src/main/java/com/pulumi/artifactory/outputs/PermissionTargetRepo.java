// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.PermissionTargetRepoActions;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PermissionTargetRepo {
    private @Nullable PermissionTargetRepoActions actions;
    /**
     * @return Pattern of artifacts to exclude.
     * 
     */
    private @Nullable List<String> excludesPatterns;
    /**
     * @return Pattern of artifacts to include.
     * 
     */
    private @Nullable List<String> includesPatterns;
    /**
     * @return List of repositories this permission target is applicable for. You can specify the name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
     * 
     */
    private List<String> repositories;

    private PermissionTargetRepo() {}
    public Optional<PermissionTargetRepoActions> actions() {
        return Optional.ofNullable(this.actions);
    }
    /**
     * @return Pattern of artifacts to exclude.
     * 
     */
    public List<String> excludesPatterns() {
        return this.excludesPatterns == null ? List.of() : this.excludesPatterns;
    }
    /**
     * @return Pattern of artifacts to include.
     * 
     */
    public List<String> includesPatterns() {
        return this.includesPatterns == null ? List.of() : this.includesPatterns;
    }
    /**
     * @return List of repositories this permission target is applicable for. You can specify the name `ANY` in the repositories section in order to apply to all repositories, `ANY REMOTE` for all remote repositories and `ANY LOCAL` for all local repositories. The default value will be `[]` if nothing is specified.
     * 
     */
    public List<String> repositories() {
        return this.repositories;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PermissionTargetRepo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable PermissionTargetRepoActions actions;
        private @Nullable List<String> excludesPatterns;
        private @Nullable List<String> includesPatterns;
        private List<String> repositories;
        public Builder() {}
        public Builder(PermissionTargetRepo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actions = defaults.actions;
    	      this.excludesPatterns = defaults.excludesPatterns;
    	      this.includesPatterns = defaults.includesPatterns;
    	      this.repositories = defaults.repositories;
        }

        @CustomType.Setter
        public Builder actions(@Nullable PermissionTargetRepoActions actions) {
            this.actions = actions;
            return this;
        }
        @CustomType.Setter
        public Builder excludesPatterns(@Nullable List<String> excludesPatterns) {
            this.excludesPatterns = excludesPatterns;
            return this;
        }
        public Builder excludesPatterns(String... excludesPatterns) {
            return excludesPatterns(List.of(excludesPatterns));
        }
        @CustomType.Setter
        public Builder includesPatterns(@Nullable List<String> includesPatterns) {
            this.includesPatterns = includesPatterns;
            return this;
        }
        public Builder includesPatterns(String... includesPatterns) {
            return includesPatterns(List.of(includesPatterns));
        }
        @CustomType.Setter
        public Builder repositories(List<String> repositories) {
            this.repositories = Objects.requireNonNull(repositories);
            return this;
        }
        public Builder repositories(String... repositories) {
            return repositories(List.of(repositories));
        }
        public PermissionTargetRepo build() {
            final var o = new PermissionTargetRepo();
            o.actions = actions;
            o.excludesPatterns = excludesPatterns;
            o.includesPatterns = includesPatterns;
            o.repositories = repositories;
            return o;
        }
    }
}
