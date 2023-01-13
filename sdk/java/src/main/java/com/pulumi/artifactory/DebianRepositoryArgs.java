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


public final class DebianRepositoryArgs extends com.pulumi.resources.ResourceArgs {

    public static final DebianRepositoryArgs Empty = new DebianRepositoryArgs();

    /**
     * When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    @Import(name="archiveBrowsingEnabled")
    private @Nullable Output<Boolean> archiveBrowsingEnabled;

    /**
     * @return When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
     * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
     * security (e.g., cross-site scripting attacks).
     * 
     */
    public Optional<Output<Boolean>> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }

    /**
     * When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    @Import(name="blackedOut")
    private @Nullable Output<Boolean> blackedOut;

    /**
     * @return When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
     * 
     */
    public Optional<Output<Boolean>> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    @Import(name="downloadDirect")
    private @Nullable Output<Boolean> downloadDirect;

    /**
     * @return When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
     * storage provider. Available in Enterprise+ and Edge licenses only.
     * 
     */
    public Optional<Output<Boolean>> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }

    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
     * artifacts are excluded.
     * 
     */
    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    @Import(name="includesPattern")
    private @Nullable Output<String> includesPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
     * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
     * 
     */
    public Optional<Output<String>> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    /**
     * The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
     * and XZ (.xz extension).
     * 
     */
    @Import(name="indexCompressionFormats")
    private @Nullable Output<List<String>> indexCompressionFormats;

    /**
     * @return The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
     * and XZ (.xz extension).
     * 
     */
    public Optional<Output<List<String>>> indexCompressionFormats() {
        return Optional.ofNullable(this.indexCompressionFormats);
    }

    /**
     * the identity key of the repo.
     * 
     */
    @Import(name="key", required=true)
    private Output<String> key;

    /**
     * @return the identity key of the repo.
     * 
     */
    public Output<String> key() {
        return this.key;
    }

    @Import(name="notes")
    private @Nullable Output<String> notes;

    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    /**
     * The primary RSA key to be used to sign packages.
     * 
     */
    @Import(name="primaryKeypairRef")
    private @Nullable Output<String> primaryKeypairRef;

    /**
     * @return The primary RSA key to be used to sign packages.
     * 
     */
    public Optional<Output<String>> primaryKeypairRef() {
        return Optional.ofNullable(this.primaryKeypairRef);
    }

    /**
     * Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    @Import(name="priorityResolution")
    private @Nullable Output<Boolean> priorityResolution;

    /**
     * @return Setting repositories with priority will cause metadata to be merged only from repositories set with this field
     * 
     */
    public Optional<Output<Boolean>> priorityResolution() {
        return Optional.ofNullable(this.priorityResolution);
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
     * List of property set name
     * 
     */
    @Import(name="propertySets")
    private @Nullable Output<List<String>> propertySets;

    /**
     * @return List of property set name
     * 
     */
    public Optional<Output<List<String>>> propertySets() {
        return Optional.ofNullable(this.propertySets);
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
     * The secondary RSA key to be used to sign packages.
     * 
     */
    @Import(name="secondaryKeypairRef")
    private @Nullable Output<String> secondaryKeypairRef;

    /**
     * @return The secondary RSA key to be used to sign packages.
     * 
     */
    public Optional<Output<String>> secondaryKeypairRef() {
        return Optional.ofNullable(this.secondaryKeypairRef);
    }

    /**
     * When set, the repository will use the deprecated trivial layout.
     * 
     * @deprecated
     * You shouldn&#39;t be using this
     * 
     */
    @Deprecated /* You shouldn't be using this */
    @Import(name="trivialLayout")
    private @Nullable Output<Boolean> trivialLayout;

    /**
     * @return When set, the repository will use the deprecated trivial layout.
     * 
     * @deprecated
     * You shouldn&#39;t be using this
     * 
     */
    @Deprecated /* You shouldn't be using this */
    public Optional<Output<Boolean>> trivialLayout() {
        return Optional.ofNullable(this.trivialLayout);
    }

    /**
     * Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    @Import(name="xrayIndex")
    private @Nullable Output<Boolean> xrayIndex;

    /**
     * @return Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
     * Xray settings.
     * 
     */
    public Optional<Output<Boolean>> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    private DebianRepositoryArgs() {}

    private DebianRepositoryArgs(DebianRepositoryArgs $) {
        this.archiveBrowsingEnabled = $.archiveBrowsingEnabled;
        this.blackedOut = $.blackedOut;
        this.description = $.description;
        this.downloadDirect = $.downloadDirect;
        this.excludesPattern = $.excludesPattern;
        this.includesPattern = $.includesPattern;
        this.indexCompressionFormats = $.indexCompressionFormats;
        this.key = $.key;
        this.notes = $.notes;
        this.primaryKeypairRef = $.primaryKeypairRef;
        this.priorityResolution = $.priorityResolution;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.propertySets = $.propertySets;
        this.repoLayoutRef = $.repoLayoutRef;
        this.secondaryKeypairRef = $.secondaryKeypairRef;
        this.trivialLayout = $.trivialLayout;
        this.xrayIndex = $.xrayIndex;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DebianRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DebianRepositoryArgs $;

        public Builder() {
            $ = new DebianRepositoryArgs();
        }

        public Builder(DebianRepositoryArgs defaults) {
            $ = new DebianRepositoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param archiveBrowsingEnabled When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
         * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
         * security (e.g., cross-site scripting attacks).
         * 
         * @return builder
         * 
         */
        public Builder archiveBrowsingEnabled(@Nullable Output<Boolean> archiveBrowsingEnabled) {
            $.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }

        /**
         * @param archiveBrowsingEnabled When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
         * therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
         * security (e.g., cross-site scripting attacks).
         * 
         * @return builder
         * 
         */
        public Builder archiveBrowsingEnabled(Boolean archiveBrowsingEnabled) {
            return archiveBrowsingEnabled(Output.of(archiveBrowsingEnabled));
        }

        /**
         * @param blackedOut When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
         * 
         * @return builder
         * 
         */
        public Builder blackedOut(@Nullable Output<Boolean> blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        /**
         * @param blackedOut When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
         * 
         * @return builder
         * 
         */
        public Builder blackedOut(Boolean blackedOut) {
            return blackedOut(Output.of(blackedOut));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param downloadDirect When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
         * storage provider. Available in Enterprise+ and Edge licenses only.
         * 
         * @return builder
         * 
         */
        public Builder downloadDirect(@Nullable Output<Boolean> downloadDirect) {
            $.downloadDirect = downloadDirect;
            return this;
        }

        /**
         * @param downloadDirect When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
         * storage provider. Available in Enterprise+ and Edge licenses only.
         * 
         * @return builder
         * 
         */
        public Builder downloadDirect(Boolean downloadDirect) {
            return downloadDirect(Output.of(downloadDirect));
        }

        /**
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
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
         * @param excludesPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**{@literal /}z/*. By default no
         * artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
        }

        /**
         * @param includesPattern List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
         * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
         * 
         * @return builder
         * 
         */
        public Builder includesPattern(@Nullable Output<String> includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        /**
         * @param includesPattern List of artifact patterns to include when evaluating artifact requests in the form of x/y/**{@literal /}z/*. When used, only
         * artifacts matching one of the include patterns are served. By default, all artifacts are included (**{@literal /}*).
         * 
         * @return builder
         * 
         */
        public Builder includesPattern(String includesPattern) {
            return includesPattern(Output.of(includesPattern));
        }

        /**
         * @param indexCompressionFormats The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
         * and XZ (.xz extension).
         * 
         * @return builder
         * 
         */
        public Builder indexCompressionFormats(@Nullable Output<List<String>> indexCompressionFormats) {
            $.indexCompressionFormats = indexCompressionFormats;
            return this;
        }

        /**
         * @param indexCompressionFormats The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
         * and XZ (.xz extension).
         * 
         * @return builder
         * 
         */
        public Builder indexCompressionFormats(List<String> indexCompressionFormats) {
            return indexCompressionFormats(Output.of(indexCompressionFormats));
        }

        /**
         * @param indexCompressionFormats The options are Bzip2 (.bz2 extension) (default), LZMA (.lzma extension)
         * and XZ (.xz extension).
         * 
         * @return builder
         * 
         */
        public Builder indexCompressionFormats(String... indexCompressionFormats) {
            return indexCompressionFormats(List.of(indexCompressionFormats));
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key the identity key of the repo.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        /**
         * @param primaryKeypairRef The primary RSA key to be used to sign packages.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeypairRef(@Nullable Output<String> primaryKeypairRef) {
            $.primaryKeypairRef = primaryKeypairRef;
            return this;
        }

        /**
         * @param primaryKeypairRef The primary RSA key to be used to sign packages.
         * 
         * @return builder
         * 
         */
        public Builder primaryKeypairRef(String primaryKeypairRef) {
            return primaryKeypairRef(Output.of(primaryKeypairRef));
        }

        /**
         * @param priorityResolution Setting repositories with priority will cause metadata to be merged only from repositories set with this field
         * 
         * @return builder
         * 
         */
        public Builder priorityResolution(@Nullable Output<Boolean> priorityResolution) {
            $.priorityResolution = priorityResolution;
            return this;
        }

        /**
         * @param priorityResolution Setting repositories with priority will cause metadata to be merged only from repositories set with this field
         * 
         * @return builder
         * 
         */
        public Builder priorityResolution(Boolean priorityResolution) {
            return priorityResolution(Output.of(priorityResolution));
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
         * @param propertySets List of property set name
         * 
         * @return builder
         * 
         */
        public Builder propertySets(@Nullable Output<List<String>> propertySets) {
            $.propertySets = propertySets;
            return this;
        }

        /**
         * @param propertySets List of property set name
         * 
         * @return builder
         * 
         */
        public Builder propertySets(List<String> propertySets) {
            return propertySets(Output.of(propertySets));
        }

        /**
         * @param propertySets List of property set name
         * 
         * @return builder
         * 
         */
        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
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
         * @param secondaryKeypairRef The secondary RSA key to be used to sign packages.
         * 
         * @return builder
         * 
         */
        public Builder secondaryKeypairRef(@Nullable Output<String> secondaryKeypairRef) {
            $.secondaryKeypairRef = secondaryKeypairRef;
            return this;
        }

        /**
         * @param secondaryKeypairRef The secondary RSA key to be used to sign packages.
         * 
         * @return builder
         * 
         */
        public Builder secondaryKeypairRef(String secondaryKeypairRef) {
            return secondaryKeypairRef(Output.of(secondaryKeypairRef));
        }

        /**
         * @param trivialLayout When set, the repository will use the deprecated trivial layout.
         * 
         * @return builder
         * 
         * @deprecated
         * You shouldn&#39;t be using this
         * 
         */
        @Deprecated /* You shouldn't be using this */
        public Builder trivialLayout(@Nullable Output<Boolean> trivialLayout) {
            $.trivialLayout = trivialLayout;
            return this;
        }

        /**
         * @param trivialLayout When set, the repository will use the deprecated trivial layout.
         * 
         * @return builder
         * 
         * @deprecated
         * You shouldn&#39;t be using this
         * 
         */
        @Deprecated /* You shouldn't be using this */
        public Builder trivialLayout(Boolean trivialLayout) {
            return trivialLayout(Output.of(trivialLayout));
        }

        /**
         * @param xrayIndex Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
         * Xray settings.
         * 
         * @return builder
         * 
         */
        public Builder xrayIndex(@Nullable Output<Boolean> xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        /**
         * @param xrayIndex Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
         * Xray settings.
         * 
         * @return builder
         * 
         */
        public Builder xrayIndex(Boolean xrayIndex) {
            return xrayIndex(Output.of(xrayIndex));
        }

        public DebianRepositoryArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}
