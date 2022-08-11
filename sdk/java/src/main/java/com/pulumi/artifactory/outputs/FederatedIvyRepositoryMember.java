// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class FederatedIvyRepositoryMember {
    /**
     * @return Represents the active state of the federated member. It is supported to change the enabled
     * status of my own member. The config will be updated on the other federated members automatically.
     * 
     */
    private final Boolean enabled;
    /**
     * @return Full URL to ending with the repository name.
     * 
     */
    private final String url;

    @CustomType.Constructor
    private FederatedIvyRepositoryMember(
        @CustomType.Parameter("enabled") Boolean enabled,
        @CustomType.Parameter("url") String url) {
        this.enabled = enabled;
        this.url = url;
    }

    /**
     * @return Represents the active state of the federated member. It is supported to change the enabled
     * status of my own member. The config will be updated on the other federated members automatically.
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return Full URL to ending with the repository name.
     * 
     */
    public String url() {
        return this.url;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FederatedIvyRepositoryMember defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Boolean enabled;
        private String url;

        public Builder() {
    	      // Empty
        }

        public Builder(FederatedIvyRepositoryMember defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.url = defaults.url;
        }

        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }        public FederatedIvyRepositoryMember build() {
            return new FederatedIvyRepositoryMember(enabled, url);
        }
    }
}
