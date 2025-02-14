// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GeneralSecurityArgs extends com.pulumi.resources.ResourceArgs {

    public static final GeneralSecurityArgs Empty = new GeneralSecurityArgs();

    @Import(name="enableAnonymousAccess")
    private @Nullable Output<Boolean> enableAnonymousAccess;

    public Optional<Output<Boolean>> enableAnonymousAccess() {
        return Optional.ofNullable(this.enableAnonymousAccess);
    }

    private GeneralSecurityArgs() {}

    private GeneralSecurityArgs(GeneralSecurityArgs $) {
        this.enableAnonymousAccess = $.enableAnonymousAccess;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GeneralSecurityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GeneralSecurityArgs $;

        public Builder() {
            $ = new GeneralSecurityArgs();
        }

        public Builder(GeneralSecurityArgs defaults) {
            $ = new GeneralSecurityArgs(Objects.requireNonNull(defaults));
        }

        public Builder enableAnonymousAccess(@Nullable Output<Boolean> enableAnonymousAccess) {
            $.enableAnonymousAccess = enableAnonymousAccess;
            return this;
        }

        public Builder enableAnonymousAccess(Boolean enableAnonymousAccess) {
            return enableAnonymousAccess(Output.of(enableAnonymousAccess));
        }

        public GeneralSecurityArgs build() {
            return $;
        }
    }

}
