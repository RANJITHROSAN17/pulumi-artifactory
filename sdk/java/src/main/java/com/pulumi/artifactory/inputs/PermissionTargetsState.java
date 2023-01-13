// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.artifactory.inputs.PermissionTargetsBuildArgs;
import com.pulumi.artifactory.inputs.PermissionTargetsReleaseBundleArgs;
import com.pulumi.artifactory.inputs.PermissionTargetsRepoArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PermissionTargetsState extends com.pulumi.resources.ResourceArgs {

    public static final PermissionTargetsState Empty = new PermissionTargetsState();

    /**
     * As for repo but for artifactory-build-info permissions.
     * 
     */
    @Import(name="build")
    private @Nullable Output<PermissionTargetsBuildArgs> build;

    /**
     * @return As for repo but for artifactory-build-info permissions.
     * 
     */
    public Optional<Output<PermissionTargetsBuildArgs>> build() {
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
    private @Nullable Output<PermissionTargetsReleaseBundleArgs> releaseBundle;

    /**
     * @return As for repo for for release-bundles permissions.
     * 
     */
    public Optional<Output<PermissionTargetsReleaseBundleArgs>> releaseBundle() {
        return Optional.ofNullable(this.releaseBundle);
    }

    /**
     * Repository permission configuration.
     * 
     */
    @Import(name="repo")
    private @Nullable Output<PermissionTargetsRepoArgs> repo;

    /**
     * @return Repository permission configuration.
     * 
     */
    public Optional<Output<PermissionTargetsRepoArgs>> repo() {
        return Optional.ofNullable(this.repo);
    }

    private PermissionTargetsState() {}

    private PermissionTargetsState(PermissionTargetsState $) {
        this.build = $.build;
        this.name = $.name;
        this.releaseBundle = $.releaseBundle;
        this.repo = $.repo;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PermissionTargetsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PermissionTargetsState $;

        public Builder() {
            $ = new PermissionTargetsState();
        }

        public Builder(PermissionTargetsState defaults) {
            $ = new PermissionTargetsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param build As for repo but for artifactory-build-info permissions.
         * 
         * @return builder
         * 
         */
        public Builder build(@Nullable Output<PermissionTargetsBuildArgs> build) {
            $.build = build;
            return this;
        }

        /**
         * @param build As for repo but for artifactory-build-info permissions.
         * 
         * @return builder
         * 
         */
        public Builder build(PermissionTargetsBuildArgs build) {
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
        public Builder releaseBundle(@Nullable Output<PermissionTargetsReleaseBundleArgs> releaseBundle) {
            $.releaseBundle = releaseBundle;
            return this;
        }

        /**
         * @param releaseBundle As for repo for for release-bundles permissions.
         * 
         * @return builder
         * 
         */
        public Builder releaseBundle(PermissionTargetsReleaseBundleArgs releaseBundle) {
            return releaseBundle(Output.of(releaseBundle));
        }

        /**
         * @param repo Repository permission configuration.
         * 
         * @return builder
         * 
         */
        public Builder repo(@Nullable Output<PermissionTargetsRepoArgs> repo) {
            $.repo = repo;
            return this;
        }

        /**
         * @param repo Repository permission configuration.
         * 
         * @return builder
         * 
         */
        public Builder repo(PermissionTargetsRepoArgs repo) {
            return repo(Output.of(repo));
        }

        public PermissionTargetsState build() {
            return $;
        }
    }

}
