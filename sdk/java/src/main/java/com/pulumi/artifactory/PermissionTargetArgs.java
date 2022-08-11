// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.artifactory.inputs.PermissionTargetBuildArgs;
import com.pulumi.artifactory.inputs.PermissionTargetReleaseBundleArgs;
import com.pulumi.artifactory.inputs.PermissionTargetRepoArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PermissionTargetArgs extends com.pulumi.resources.ResourceArgs {

    public static final PermissionTargetArgs Empty = new PermissionTargetArgs();

    /**
     * As for repo but for artifactory-build-info permssions.
     * 
     */
    @Import(name="build")
    private @Nullable Output<PermissionTargetBuildArgs> build;

    /**
     * @return As for repo but for artifactory-build-info permssions.
     * 
     */
    public Optional<Output<PermissionTargetBuildArgs>> build() {
        return Optional.ofNullable(this.build);
    }

    /**
     * Name of permission.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of permission.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * As for repo for for release-bundles permissions.
     * 
     */
    @Import(name="releaseBundle")
    private @Nullable Output<PermissionTargetReleaseBundleArgs> releaseBundle;

    /**
     * @return As for repo for for release-bundles permissions.
     * 
     */
    public Optional<Output<PermissionTargetReleaseBundleArgs>> releaseBundle() {
        return Optional.ofNullable(this.releaseBundle);
    }

    /**
     * Repository permission configuration.
     * 
     */
    @Import(name="repo")
    private @Nullable Output<PermissionTargetRepoArgs> repo;

    /**
     * @return Repository permission configuration.
     * 
     */
    public Optional<Output<PermissionTargetRepoArgs>> repo() {
        return Optional.ofNullable(this.repo);
    }

    private PermissionTargetArgs() {}

    private PermissionTargetArgs(PermissionTargetArgs $) {
        this.build = $.build;
        this.name = $.name;
        this.releaseBundle = $.releaseBundle;
        this.repo = $.repo;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PermissionTargetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PermissionTargetArgs $;

        public Builder() {
            $ = new PermissionTargetArgs();
        }

        public Builder(PermissionTargetArgs defaults) {
            $ = new PermissionTargetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param build As for repo but for artifactory-build-info permssions.
         * 
         * @return builder
         * 
         */
        public Builder build(@Nullable Output<PermissionTargetBuildArgs> build) {
            $.build = build;
            return this;
        }

        /**
         * @param build As for repo but for artifactory-build-info permssions.
         * 
         * @return builder
         * 
         */
        public Builder build(PermissionTargetBuildArgs build) {
            return build(Output.of(build));
        }

        /**
         * @param name Name of permission.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of permission.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param releaseBundle As for repo for for release-bundles permissions.
         * 
         * @return builder
         * 
         */
        public Builder releaseBundle(@Nullable Output<PermissionTargetReleaseBundleArgs> releaseBundle) {
            $.releaseBundle = releaseBundle;
            return this;
        }

        /**
         * @param releaseBundle As for repo for for release-bundles permissions.
         * 
         * @return builder
         * 
         */
        public Builder releaseBundle(PermissionTargetReleaseBundleArgs releaseBundle) {
            return releaseBundle(Output.of(releaseBundle));
        }

        /**
         * @param repo Repository permission configuration.
         * 
         * @return builder
         * 
         */
        public Builder repo(@Nullable Output<PermissionTargetRepoArgs> repo) {
            $.repo = repo;
            return this;
        }

        /**
         * @param repo Repository permission configuration.
         * 
         * @return builder
         * 
         */
        public Builder repo(PermissionTargetRepoArgs repo) {
            return repo(Output.of(repo));
        }

        public PermissionTargetArgs build() {
            return $;
        }
    }

}
