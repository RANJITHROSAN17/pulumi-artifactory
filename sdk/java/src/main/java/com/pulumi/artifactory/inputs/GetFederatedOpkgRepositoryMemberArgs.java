// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;


public final class GetFederatedOpkgRepositoryMemberArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetFederatedOpkgRepositoryMemberArgs Empty = new GetFederatedOpkgRepositoryMemberArgs();

    /**
     * Represents the active state of the federated member. It is supported to change the enabled
     * status of my own member. The config will be updated on the other federated members automatically.
     * 
     */
    @Import(name="enabled", required=true)
    private Output<Boolean> enabled;

    /**
     * @return Represents the active state of the federated member. It is supported to change the enabled
     * status of my own member. The config will be updated on the other federated members automatically.
     * 
     */
    public Output<Boolean> enabled() {
        return this.enabled;
    }

    /**
     * Full URL to ending with the repository name.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return Full URL to ending with the repository name.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    private GetFederatedOpkgRepositoryMemberArgs() {}

    private GetFederatedOpkgRepositoryMemberArgs(GetFederatedOpkgRepositoryMemberArgs $) {
        this.enabled = $.enabled;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFederatedOpkgRepositoryMemberArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFederatedOpkgRepositoryMemberArgs $;

        public Builder() {
            $ = new GetFederatedOpkgRepositoryMemberArgs();
        }

        public Builder(GetFederatedOpkgRepositoryMemberArgs defaults) {
            $ = new GetFederatedOpkgRepositoryMemberArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Represents the active state of the federated member. It is supported to change the enabled
         * status of my own member. The config will be updated on the other federated members automatically.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Represents the active state of the federated member. It is supported to change the enabled
         * status of my own member. The config will be updated on the other federated members automatically.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param url Full URL to ending with the repository name.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Full URL to ending with the repository name.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public GetFederatedOpkgRepositoryMemberArgs build() {
            $.enabled = Objects.requireNonNull($.enabled, "expected parameter 'enabled' to be non-null");
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
