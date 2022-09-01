// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RepositoryLayoutState extends com.pulumi.resources.ResourceArgs {

    public static final RepositoryLayoutState Empty = new RepositoryLayoutState();

    /**
     * Please refer to: [Path
     * Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts)
     * in the Artifactory Wiki documentation.
     * 
     */
    @Import(name="artifactPathPattern")
    private @Nullable Output<String> artifactPathPattern;

    /**
     * @return Please refer to: [Path
     * Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts)
     * in the Artifactory Wiki documentation.
     * 
     */
    public Optional<Output<String>> artifactPathPattern() {
        return Optional.ofNullable(this.artifactPathPattern);
    }

    /**
     * Please refer to: [Descriptor Path
     * Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in
     * the Artifactory Wiki documentation.
     * 
     */
    @Import(name="descriptorPathPattern")
    private @Nullable Output<String> descriptorPathPattern;

    /**
     * @return Please refer to: [Descriptor Path
     * Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in
     * the Artifactory Wiki documentation.
     * 
     */
    public Optional<Output<String>> descriptorPathPattern() {
        return Optional.ofNullable(this.descriptorPathPattern);
    }

    /**
     * When set, &#39;descriptor_path_pattern&#39; will be used. Default to &#39;false&#39;.
     * 
     */
    @Import(name="distinctiveDescriptorPathPattern")
    private @Nullable Output<Boolean> distinctiveDescriptorPathPattern;

    /**
     * @return When set, &#39;descriptor_path_pattern&#39; will be used. Default to &#39;false&#39;.
     * 
     */
    public Optional<Output<Boolean>> distinctiveDescriptorPathPattern() {
        return Optional.ofNullable(this.distinctiveDescriptorPathPattern);
    }

    /**
     * A regular expression matching the integration revision string appearing in a file name as part of the artifact&#39;s path.
     * For example, &#39;SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))&#39;, in Maven. Note! Take care not to introduce any regexp
     * capturing groups within this expression. If not applicable use &#39;.*&#39;
     * 
     */
    @Import(name="fileIntegrationRevisionRegexp")
    private @Nullable Output<String> fileIntegrationRevisionRegexp;

    /**
     * @return A regular expression matching the integration revision string appearing in a file name as part of the artifact&#39;s path.
     * For example, &#39;SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))&#39;, in Maven. Note! Take care not to introduce any regexp
     * capturing groups within this expression. If not applicable use &#39;.*&#39;
     * 
     */
    public Optional<Output<String>> fileIntegrationRevisionRegexp() {
        return Optional.ofNullable(this.fileIntegrationRevisionRegexp);
    }

    /**
     * A regular expression matching the integration revision string appearing in a folder name as part of the artifact&#39;s path.
     * For example, &#39;SNAPSHOT&#39;, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression.
     * If not applicable use &#39;.*&#39;
     * 
     */
    @Import(name="folderIntegrationRevisionRegexp")
    private @Nullable Output<String> folderIntegrationRevisionRegexp;

    /**
     * @return A regular expression matching the integration revision string appearing in a folder name as part of the artifact&#39;s path.
     * For example, &#39;SNAPSHOT&#39;, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression.
     * If not applicable use &#39;.*&#39;
     * 
     */
    public Optional<Output<String>> folderIntegrationRevisionRegexp() {
        return Optional.ofNullable(this.folderIntegrationRevisionRegexp);
    }

    /**
     * Layout name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Layout name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private RepositoryLayoutState() {}

    private RepositoryLayoutState(RepositoryLayoutState $) {
        this.artifactPathPattern = $.artifactPathPattern;
        this.descriptorPathPattern = $.descriptorPathPattern;
        this.distinctiveDescriptorPathPattern = $.distinctiveDescriptorPathPattern;
        this.fileIntegrationRevisionRegexp = $.fileIntegrationRevisionRegexp;
        this.folderIntegrationRevisionRegexp = $.folderIntegrationRevisionRegexp;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RepositoryLayoutState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RepositoryLayoutState $;

        public Builder() {
            $ = new RepositoryLayoutState();
        }

        public Builder(RepositoryLayoutState defaults) {
            $ = new RepositoryLayoutState(Objects.requireNonNull(defaults));
        }

        /**
         * @param artifactPathPattern Please refer to: [Path
         * Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts)
         * in the Artifactory Wiki documentation.
         * 
         * @return builder
         * 
         */
        public Builder artifactPathPattern(@Nullable Output<String> artifactPathPattern) {
            $.artifactPathPattern = artifactPathPattern;
            return this;
        }

        /**
         * @param artifactPathPattern Please refer to: [Path
         * Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-ModulesandPathPatternsusedbyRepositoryLayouts)
         * in the Artifactory Wiki documentation.
         * 
         * @return builder
         * 
         */
        public Builder artifactPathPattern(String artifactPathPattern) {
            return artifactPathPattern(Output.of(artifactPathPattern));
        }

        /**
         * @param descriptorPathPattern Please refer to: [Descriptor Path
         * Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in
         * the Artifactory Wiki documentation.
         * 
         * @return builder
         * 
         */
        public Builder descriptorPathPattern(@Nullable Output<String> descriptorPathPattern) {
            $.descriptorPathPattern = descriptorPathPattern;
            return this;
        }

        /**
         * @param descriptorPathPattern Please refer to: [Descriptor Path
         * Patterns](https://www.jfrog.com/confluence/display/JFROG/Repository+Layouts#RepositoryLayouts-DescriptorPathPatterns) in
         * the Artifactory Wiki documentation.
         * 
         * @return builder
         * 
         */
        public Builder descriptorPathPattern(String descriptorPathPattern) {
            return descriptorPathPattern(Output.of(descriptorPathPattern));
        }

        /**
         * @param distinctiveDescriptorPathPattern When set, &#39;descriptor_path_pattern&#39; will be used. Default to &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder distinctiveDescriptorPathPattern(@Nullable Output<Boolean> distinctiveDescriptorPathPattern) {
            $.distinctiveDescriptorPathPattern = distinctiveDescriptorPathPattern;
            return this;
        }

        /**
         * @param distinctiveDescriptorPathPattern When set, &#39;descriptor_path_pattern&#39; will be used. Default to &#39;false&#39;.
         * 
         * @return builder
         * 
         */
        public Builder distinctiveDescriptorPathPattern(Boolean distinctiveDescriptorPathPattern) {
            return distinctiveDescriptorPathPattern(Output.of(distinctiveDescriptorPathPattern));
        }

        /**
         * @param fileIntegrationRevisionRegexp A regular expression matching the integration revision string appearing in a file name as part of the artifact&#39;s path.
         * For example, &#39;SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))&#39;, in Maven. Note! Take care not to introduce any regexp
         * capturing groups within this expression. If not applicable use &#39;.*&#39;
         * 
         * @return builder
         * 
         */
        public Builder fileIntegrationRevisionRegexp(@Nullable Output<String> fileIntegrationRevisionRegexp) {
            $.fileIntegrationRevisionRegexp = fileIntegrationRevisionRegexp;
            return this;
        }

        /**
         * @param fileIntegrationRevisionRegexp A regular expression matching the integration revision string appearing in a file name as part of the artifact&#39;s path.
         * For example, &#39;SNAPSHOT|(?:(?:[0-9]{8}.[0-9]{6})-(?:[0-9]+))&#39;, in Maven. Note! Take care not to introduce any regexp
         * capturing groups within this expression. If not applicable use &#39;.*&#39;
         * 
         * @return builder
         * 
         */
        public Builder fileIntegrationRevisionRegexp(String fileIntegrationRevisionRegexp) {
            return fileIntegrationRevisionRegexp(Output.of(fileIntegrationRevisionRegexp));
        }

        /**
         * @param folderIntegrationRevisionRegexp A regular expression matching the integration revision string appearing in a folder name as part of the artifact&#39;s path.
         * For example, &#39;SNAPSHOT&#39;, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression.
         * If not applicable use &#39;.*&#39;
         * 
         * @return builder
         * 
         */
        public Builder folderIntegrationRevisionRegexp(@Nullable Output<String> folderIntegrationRevisionRegexp) {
            $.folderIntegrationRevisionRegexp = folderIntegrationRevisionRegexp;
            return this;
        }

        /**
         * @param folderIntegrationRevisionRegexp A regular expression matching the integration revision string appearing in a folder name as part of the artifact&#39;s path.
         * For example, &#39;SNAPSHOT&#39;, in Maven. Note! Take care not to introduce any regexp capturing groups within this expression.
         * If not applicable use &#39;.*&#39;
         * 
         * @return builder
         * 
         */
        public Builder folderIntegrationRevisionRegexp(String folderIntegrationRevisionRegexp) {
            return folderIntegrationRevisionRegexp(Output.of(folderIntegrationRevisionRegexp));
        }

        /**
         * @param name Layout name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Layout name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public RepositoryLayoutState build() {
            return $;
        }
    }

}
