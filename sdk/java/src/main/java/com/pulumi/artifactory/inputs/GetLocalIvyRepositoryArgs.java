// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetLocalIvyRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLocalIvyRepositoryArgs Empty = new GetLocalIvyRepositoryArgs();

    @Import(name="archiveBrowsingEnabled")
    private @Nullable Output<Boolean> archiveBrowsingEnabled;

    public Optional<Output<Boolean>> archiveBrowsingEnabled() {
        return Optional.ofNullable(this.archiveBrowsingEnabled);
    }

    @Import(name="blackedOut")
    private @Nullable Output<Boolean> blackedOut;

    public Optional<Output<Boolean>> blackedOut() {
        return Optional.ofNullable(this.blackedOut);
    }

    @Import(name="cdnRedirect")
    private @Nullable Output<Boolean> cdnRedirect;

    public Optional<Output<Boolean>> cdnRedirect() {
        return Optional.ofNullable(this.cdnRedirect);
    }

    @Import(name="checksumPolicyType")
    private @Nullable Output<String> checksumPolicyType;

    public Optional<Output<String>> checksumPolicyType() {
        return Optional.ofNullable(this.checksumPolicyType);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="downloadDirect")
    private @Nullable Output<Boolean> downloadDirect;

    public Optional<Output<Boolean>> downloadDirect() {
        return Optional.ofNullable(this.downloadDirect);
    }

    @Import(name="excludesPattern")
    private @Nullable Output<String> excludesPattern;

    public Optional<Output<String>> excludesPattern() {
        return Optional.ofNullable(this.excludesPattern);
    }

    @Import(name="handleReleases")
    private @Nullable Output<Boolean> handleReleases;

    public Optional<Output<Boolean>> handleReleases() {
        return Optional.ofNullable(this.handleReleases);
    }

    @Import(name="handleSnapshots")
    private @Nullable Output<Boolean> handleSnapshots;

    public Optional<Output<Boolean>> handleSnapshots() {
        return Optional.ofNullable(this.handleSnapshots);
    }

    @Import(name="includesPattern")
    private @Nullable Output<String> includesPattern;

    public Optional<Output<String>> includesPattern() {
        return Optional.ofNullable(this.includesPattern);
    }

    @Import(name="key", required=true)
    private Output<String> key;

    public Output<String> key() {
        return this.key;
    }

    @Import(name="maxUniqueSnapshots")
    private @Nullable Output<Integer> maxUniqueSnapshots;

    public Optional<Output<Integer>> maxUniqueSnapshots() {
        return Optional.ofNullable(this.maxUniqueSnapshots);
    }

    @Import(name="notes")
    private @Nullable Output<String> notes;

    public Optional<Output<String>> notes() {
        return Optional.ofNullable(this.notes);
    }

    @Import(name="priorityResolution")
    private @Nullable Output<Boolean> priorityResolution;

    public Optional<Output<Boolean>> priorityResolution() {
        return Optional.ofNullable(this.priorityResolution);
    }

    @Import(name="projectEnvironments")
    private @Nullable Output<List<String>> projectEnvironments;

    public Optional<Output<List<String>>> projectEnvironments() {
        return Optional.ofNullable(this.projectEnvironments);
    }

    @Import(name="projectKey")
    private @Nullable Output<String> projectKey;

    public Optional<Output<String>> projectKey() {
        return Optional.ofNullable(this.projectKey);
    }

    @Import(name="propertySets")
    private @Nullable Output<List<String>> propertySets;

    public Optional<Output<List<String>>> propertySets() {
        return Optional.ofNullable(this.propertySets);
    }

    @Import(name="repoLayoutRef")
    private @Nullable Output<String> repoLayoutRef;

    public Optional<Output<String>> repoLayoutRef() {
        return Optional.ofNullable(this.repoLayoutRef);
    }

    @Import(name="snapshotVersionBehavior")
    private @Nullable Output<String> snapshotVersionBehavior;

    public Optional<Output<String>> snapshotVersionBehavior() {
        return Optional.ofNullable(this.snapshotVersionBehavior);
    }

    @Import(name="suppressPomConsistencyChecks")
    private @Nullable Output<Boolean> suppressPomConsistencyChecks;

    public Optional<Output<Boolean>> suppressPomConsistencyChecks() {
        return Optional.ofNullable(this.suppressPomConsistencyChecks);
    }

    @Import(name="xrayIndex")
    private @Nullable Output<Boolean> xrayIndex;

    public Optional<Output<Boolean>> xrayIndex() {
        return Optional.ofNullable(this.xrayIndex);
    }

    private GetLocalIvyRepositoryArgs() {}

    private GetLocalIvyRepositoryArgs(GetLocalIvyRepositoryArgs $) {
        this.archiveBrowsingEnabled = $.archiveBrowsingEnabled;
        this.blackedOut = $.blackedOut;
        this.cdnRedirect = $.cdnRedirect;
        this.checksumPolicyType = $.checksumPolicyType;
        this.description = $.description;
        this.downloadDirect = $.downloadDirect;
        this.excludesPattern = $.excludesPattern;
        this.handleReleases = $.handleReleases;
        this.handleSnapshots = $.handleSnapshots;
        this.includesPattern = $.includesPattern;
        this.key = $.key;
        this.maxUniqueSnapshots = $.maxUniqueSnapshots;
        this.notes = $.notes;
        this.priorityResolution = $.priorityResolution;
        this.projectEnvironments = $.projectEnvironments;
        this.projectKey = $.projectKey;
        this.propertySets = $.propertySets;
        this.repoLayoutRef = $.repoLayoutRef;
        this.snapshotVersionBehavior = $.snapshotVersionBehavior;
        this.suppressPomConsistencyChecks = $.suppressPomConsistencyChecks;
        this.xrayIndex = $.xrayIndex;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLocalIvyRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLocalIvyRepositoryArgs $;

        public Builder() {
            $ = new GetLocalIvyRepositoryArgs();
        }

        public Builder(GetLocalIvyRepositoryArgs defaults) {
            $ = new GetLocalIvyRepositoryArgs(Objects.requireNonNull(defaults));
        }

        public Builder archiveBrowsingEnabled(@Nullable Output<Boolean> archiveBrowsingEnabled) {
            $.archiveBrowsingEnabled = archiveBrowsingEnabled;
            return this;
        }

        public Builder archiveBrowsingEnabled(Boolean archiveBrowsingEnabled) {
            return archiveBrowsingEnabled(Output.of(archiveBrowsingEnabled));
        }

        public Builder blackedOut(@Nullable Output<Boolean> blackedOut) {
            $.blackedOut = blackedOut;
            return this;
        }

        public Builder blackedOut(Boolean blackedOut) {
            return blackedOut(Output.of(blackedOut));
        }

        public Builder cdnRedirect(@Nullable Output<Boolean> cdnRedirect) {
            $.cdnRedirect = cdnRedirect;
            return this;
        }

        public Builder cdnRedirect(Boolean cdnRedirect) {
            return cdnRedirect(Output.of(cdnRedirect));
        }

        public Builder checksumPolicyType(@Nullable Output<String> checksumPolicyType) {
            $.checksumPolicyType = checksumPolicyType;
            return this;
        }

        public Builder checksumPolicyType(String checksumPolicyType) {
            return checksumPolicyType(Output.of(checksumPolicyType));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder downloadDirect(@Nullable Output<Boolean> downloadDirect) {
            $.downloadDirect = downloadDirect;
            return this;
        }

        public Builder downloadDirect(Boolean downloadDirect) {
            return downloadDirect(Output.of(downloadDirect));
        }

        public Builder excludesPattern(@Nullable Output<String> excludesPattern) {
            $.excludesPattern = excludesPattern;
            return this;
        }

        public Builder excludesPattern(String excludesPattern) {
            return excludesPattern(Output.of(excludesPattern));
        }

        public Builder handleReleases(@Nullable Output<Boolean> handleReleases) {
            $.handleReleases = handleReleases;
            return this;
        }

        public Builder handleReleases(Boolean handleReleases) {
            return handleReleases(Output.of(handleReleases));
        }

        public Builder handleSnapshots(@Nullable Output<Boolean> handleSnapshots) {
            $.handleSnapshots = handleSnapshots;
            return this;
        }

        public Builder handleSnapshots(Boolean handleSnapshots) {
            return handleSnapshots(Output.of(handleSnapshots));
        }

        public Builder includesPattern(@Nullable Output<String> includesPattern) {
            $.includesPattern = includesPattern;
            return this;
        }

        public Builder includesPattern(String includesPattern) {
            return includesPattern(Output.of(includesPattern));
        }

        public Builder key(Output<String> key) {
            $.key = key;
            return this;
        }

        public Builder key(String key) {
            return key(Output.of(key));
        }

        public Builder maxUniqueSnapshots(@Nullable Output<Integer> maxUniqueSnapshots) {
            $.maxUniqueSnapshots = maxUniqueSnapshots;
            return this;
        }

        public Builder maxUniqueSnapshots(Integer maxUniqueSnapshots) {
            return maxUniqueSnapshots(Output.of(maxUniqueSnapshots));
        }

        public Builder notes(@Nullable Output<String> notes) {
            $.notes = notes;
            return this;
        }

        public Builder notes(String notes) {
            return notes(Output.of(notes));
        }

        public Builder priorityResolution(@Nullable Output<Boolean> priorityResolution) {
            $.priorityResolution = priorityResolution;
            return this;
        }

        public Builder priorityResolution(Boolean priorityResolution) {
            return priorityResolution(Output.of(priorityResolution));
        }

        public Builder projectEnvironments(@Nullable Output<List<String>> projectEnvironments) {
            $.projectEnvironments = projectEnvironments;
            return this;
        }

        public Builder projectEnvironments(List<String> projectEnvironments) {
            return projectEnvironments(Output.of(projectEnvironments));
        }

        public Builder projectEnvironments(String... projectEnvironments) {
            return projectEnvironments(List.of(projectEnvironments));
        }

        public Builder projectKey(@Nullable Output<String> projectKey) {
            $.projectKey = projectKey;
            return this;
        }

        public Builder projectKey(String projectKey) {
            return projectKey(Output.of(projectKey));
        }

        public Builder propertySets(@Nullable Output<List<String>> propertySets) {
            $.propertySets = propertySets;
            return this;
        }

        public Builder propertySets(List<String> propertySets) {
            return propertySets(Output.of(propertySets));
        }

        public Builder propertySets(String... propertySets) {
            return propertySets(List.of(propertySets));
        }

        public Builder repoLayoutRef(@Nullable Output<String> repoLayoutRef) {
            $.repoLayoutRef = repoLayoutRef;
            return this;
        }

        public Builder repoLayoutRef(String repoLayoutRef) {
            return repoLayoutRef(Output.of(repoLayoutRef));
        }

        public Builder snapshotVersionBehavior(@Nullable Output<String> snapshotVersionBehavior) {
            $.snapshotVersionBehavior = snapshotVersionBehavior;
            return this;
        }

        public Builder snapshotVersionBehavior(String snapshotVersionBehavior) {
            return snapshotVersionBehavior(Output.of(snapshotVersionBehavior));
        }

        public Builder suppressPomConsistencyChecks(@Nullable Output<Boolean> suppressPomConsistencyChecks) {
            $.suppressPomConsistencyChecks = suppressPomConsistencyChecks;
            return this;
        }

        public Builder suppressPomConsistencyChecks(Boolean suppressPomConsistencyChecks) {
            return suppressPomConsistencyChecks(Output.of(suppressPomConsistencyChecks));
        }

        public Builder xrayIndex(@Nullable Output<Boolean> xrayIndex) {
            $.xrayIndex = xrayIndex;
            return this;
        }

        public Builder xrayIndex(Boolean xrayIndex) {
            return xrayIndex(Output.of(xrayIndex));
        }

        public GetLocalIvyRepositoryArgs build() {
            $.key = Objects.requireNonNull($.key, "expected parameter 'key' to be non-null");
            return $;
        }
    }

}
