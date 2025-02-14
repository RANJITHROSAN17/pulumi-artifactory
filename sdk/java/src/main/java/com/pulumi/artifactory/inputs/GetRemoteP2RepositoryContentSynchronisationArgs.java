// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRemoteP2RepositoryContentSynchronisationArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetRemoteP2RepositoryContentSynchronisationArgs Empty = new GetRemoteP2RepositoryContentSynchronisationArgs();

    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    @Import(name="propertiesEnabled")
    private @Nullable Output<Boolean> propertiesEnabled;

    public Optional<Output<Boolean>> propertiesEnabled() {
        return Optional.ofNullable(this.propertiesEnabled);
    }

    @Import(name="sourceOriginAbsenceDetection")
    private @Nullable Output<Boolean> sourceOriginAbsenceDetection;

    public Optional<Output<Boolean>> sourceOriginAbsenceDetection() {
        return Optional.ofNullable(this.sourceOriginAbsenceDetection);
    }

    @Import(name="statisticsEnabled")
    private @Nullable Output<Boolean> statisticsEnabled;

    public Optional<Output<Boolean>> statisticsEnabled() {
        return Optional.ofNullable(this.statisticsEnabled);
    }

    private GetRemoteP2RepositoryContentSynchronisationArgs() {}

    private GetRemoteP2RepositoryContentSynchronisationArgs(GetRemoteP2RepositoryContentSynchronisationArgs $) {
        this.enabled = $.enabled;
        this.propertiesEnabled = $.propertiesEnabled;
        this.sourceOriginAbsenceDetection = $.sourceOriginAbsenceDetection;
        this.statisticsEnabled = $.statisticsEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRemoteP2RepositoryContentSynchronisationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRemoteP2RepositoryContentSynchronisationArgs $;

        public Builder() {
            $ = new GetRemoteP2RepositoryContentSynchronisationArgs();
        }

        public Builder(GetRemoteP2RepositoryContentSynchronisationArgs defaults) {
            $ = new GetRemoteP2RepositoryContentSynchronisationArgs(Objects.requireNonNull(defaults));
        }

        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public Builder propertiesEnabled(@Nullable Output<Boolean> propertiesEnabled) {
            $.propertiesEnabled = propertiesEnabled;
            return this;
        }

        public Builder propertiesEnabled(Boolean propertiesEnabled) {
            return propertiesEnabled(Output.of(propertiesEnabled));
        }

        public Builder sourceOriginAbsenceDetection(@Nullable Output<Boolean> sourceOriginAbsenceDetection) {
            $.sourceOriginAbsenceDetection = sourceOriginAbsenceDetection;
            return this;
        }

        public Builder sourceOriginAbsenceDetection(Boolean sourceOriginAbsenceDetection) {
            return sourceOriginAbsenceDetection(Output.of(sourceOriginAbsenceDetection));
        }

        public Builder statisticsEnabled(@Nullable Output<Boolean> statisticsEnabled) {
            $.statisticsEnabled = statisticsEnabled;
            return this;
        }

        public Builder statisticsEnabled(Boolean statisticsEnabled) {
            return statisticsEnabled(Output.of(statisticsEnabled));
        }

        public GetRemoteP2RepositoryContentSynchronisationArgs build() {
            return $;
        }
    }

}
