// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReleaseBundleWebhookCriteriaArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReleaseBundleWebhookCriteriaArgs Empty = new ReleaseBundleWebhookCriteriaArgs();

    /**
     * Trigger on any release bundle.
     * 
     */
    @Import(name="anyReleaseBundle", required=true)
    private Output<Boolean> anyReleaseBundle;

    /**
     * @return Trigger on any release bundle.
     * 
     */
    public Output<Boolean> anyReleaseBundle() {
        return this.anyReleaseBundle;
    }

    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
     * 
     */
    @Import(name="excludePatterns")
    private @Nullable Output<List<String>> excludePatterns;

    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
     * 
     */
    public Optional<Output<List<String>>> excludePatterns() {
        return Optional.ofNullable(this.excludePatterns);
    }

    /**
     * Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
     * 
     */
    @Import(name="includePatterns")
    private @Nullable Output<List<String>> includePatterns;

    /**
     * @return Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
     * 
     */
    public Optional<Output<List<String>>> includePatterns() {
        return Optional.ofNullable(this.includePatterns);
    }

    /**
     * Trigger on this list of release bundle names.
     * 
     */
    @Import(name="registeredReleaseBundleNames", required=true)
    private Output<List<String>> registeredReleaseBundleNames;

    /**
     * @return Trigger on this list of release bundle names.
     * 
     */
    public Output<List<String>> registeredReleaseBundleNames() {
        return this.registeredReleaseBundleNames;
    }

    private ReleaseBundleWebhookCriteriaArgs() {}

    private ReleaseBundleWebhookCriteriaArgs(ReleaseBundleWebhookCriteriaArgs $) {
        this.anyReleaseBundle = $.anyReleaseBundle;
        this.excludePatterns = $.excludePatterns;
        this.includePatterns = $.includePatterns;
        this.registeredReleaseBundleNames = $.registeredReleaseBundleNames;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReleaseBundleWebhookCriteriaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReleaseBundleWebhookCriteriaArgs $;

        public Builder() {
            $ = new ReleaseBundleWebhookCriteriaArgs();
        }

        public Builder(ReleaseBundleWebhookCriteriaArgs defaults) {
            $ = new ReleaseBundleWebhookCriteriaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param anyReleaseBundle Trigger on any release bundle.
         * 
         * @return builder
         * 
         */
        public Builder anyReleaseBundle(Output<Boolean> anyReleaseBundle) {
            $.anyReleaseBundle = anyReleaseBundle;
            return this;
        }

        /**
         * @param anyReleaseBundle Trigger on any release bundle.
         * 
         * @return builder
         * 
         */
        public Builder anyReleaseBundle(Boolean anyReleaseBundle) {
            return anyReleaseBundle(Output.of(anyReleaseBundle));
        }

        /**
         * @param excludePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
         * 
         * @return builder
         * 
         */
        public Builder excludePatterns(@Nullable Output<List<String>> excludePatterns) {
            $.excludePatterns = excludePatterns;
            return this;
        }

        /**
         * @param excludePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
         * 
         * @return builder
         * 
         */
        public Builder excludePatterns(List<String> excludePatterns) {
            return excludePatterns(Output.of(excludePatterns));
        }

        /**
         * @param excludePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
         * 
         * @return builder
         * 
         */
        public Builder excludePatterns(String... excludePatterns) {
            return excludePatterns(List.of(excludePatterns));
        }

        /**
         * @param includePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
         * 
         * @return builder
         * 
         */
        public Builder includePatterns(@Nullable Output<List<String>> includePatterns) {
            $.includePatterns = includePatterns;
            return this;
        }

        /**
         * @param includePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
         * 
         * @return builder
         * 
         */
        public Builder includePatterns(List<String> includePatterns) {
            return includePatterns(Output.of(includePatterns));
        }

        /**
         * @param includePatterns Simple comma separated wildcard patterns for repository artifact paths (with no leading slash). Ant-style path expressions are supported (*, *\*, ?). For example: &#34;org/apache/**&#34;.
         * 
         * @return builder
         * 
         */
        public Builder includePatterns(String... includePatterns) {
            return includePatterns(List.of(includePatterns));
        }

        /**
         * @param registeredReleaseBundleNames Trigger on this list of release bundle names.
         * 
         * @return builder
         * 
         */
        public Builder registeredReleaseBundleNames(Output<List<String>> registeredReleaseBundleNames) {
            $.registeredReleaseBundleNames = registeredReleaseBundleNames;
            return this;
        }

        /**
         * @param registeredReleaseBundleNames Trigger on this list of release bundle names.
         * 
         * @return builder
         * 
         */
        public Builder registeredReleaseBundleNames(List<String> registeredReleaseBundleNames) {
            return registeredReleaseBundleNames(Output.of(registeredReleaseBundleNames));
        }

        /**
         * @param registeredReleaseBundleNames Trigger on this list of release bundle names.
         * 
         * @return builder
         * 
         */
        public Builder registeredReleaseBundleNames(String... registeredReleaseBundleNames) {
            return registeredReleaseBundleNames(List.of(registeredReleaseBundleNames));
        }

        public ReleaseBundleWebhookCriteriaArgs build() {
            $.anyReleaseBundle = Objects.requireNonNull($.anyReleaseBundle, "expected parameter 'anyReleaseBundle' to be non-null");
            $.registeredReleaseBundleNames = Objects.requireNonNull($.registeredReleaseBundleNames, "expected parameter 'registeredReleaseBundleNames' to be non-null");
            return $;
        }
    }

}
