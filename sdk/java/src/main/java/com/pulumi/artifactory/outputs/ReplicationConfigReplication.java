// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ReplicationConfigReplication {
    private @Nullable Boolean enabled;
    /**
     * @return Requires password encryption to be turned off `POST /api/system/decrypt`.
     * 
     */
    private @Nullable String password;
    private @Nullable String pathPrefix;
    /**
     * @return Proxy key from Artifactory Proxies setting
     * 
     */
    private @Nullable String proxy;
    private @Nullable Integer socketTimeoutMillis;
    private @Nullable Boolean syncDeletes;
    private @Nullable Boolean syncProperties;
    private @Nullable Boolean syncStatistics;
    private @Nullable String url;
    private @Nullable String username;

    private ReplicationConfigReplication() {}
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return Requires password encryption to be turned off `POST /api/system/decrypt`.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    public Optional<String> pathPrefix() {
        return Optional.ofNullable(this.pathPrefix);
    }
    /**
     * @return Proxy key from Artifactory Proxies setting
     * 
     */
    public Optional<String> proxy() {
        return Optional.ofNullable(this.proxy);
    }
    public Optional<Integer> socketTimeoutMillis() {
        return Optional.ofNullable(this.socketTimeoutMillis);
    }
    public Optional<Boolean> syncDeletes() {
        return Optional.ofNullable(this.syncDeletes);
    }
    public Optional<Boolean> syncProperties() {
        return Optional.ofNullable(this.syncProperties);
    }
    public Optional<Boolean> syncStatistics() {
        return Optional.ofNullable(this.syncStatistics);
    }
    public Optional<String> url() {
        return Optional.ofNullable(this.url);
    }
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ReplicationConfigReplication defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean enabled;
        private @Nullable String password;
        private @Nullable String pathPrefix;
        private @Nullable String proxy;
        private @Nullable Integer socketTimeoutMillis;
        private @Nullable Boolean syncDeletes;
        private @Nullable Boolean syncProperties;
        private @Nullable Boolean syncStatistics;
        private @Nullable String url;
        private @Nullable String username;
        public Builder() {}
        public Builder(ReplicationConfigReplication defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.enabled = defaults.enabled;
    	      this.password = defaults.password;
    	      this.pathPrefix = defaults.pathPrefix;
    	      this.proxy = defaults.proxy;
    	      this.socketTimeoutMillis = defaults.socketTimeoutMillis;
    	      this.syncDeletes = defaults.syncDeletes;
    	      this.syncProperties = defaults.syncProperties;
    	      this.syncStatistics = defaults.syncStatistics;
    	      this.url = defaults.url;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder pathPrefix(@Nullable String pathPrefix) {
            this.pathPrefix = pathPrefix;
            return this;
        }
        @CustomType.Setter
        public Builder proxy(@Nullable String proxy) {
            this.proxy = proxy;
            return this;
        }
        @CustomType.Setter
        public Builder socketTimeoutMillis(@Nullable Integer socketTimeoutMillis) {
            this.socketTimeoutMillis = socketTimeoutMillis;
            return this;
        }
        @CustomType.Setter
        public Builder syncDeletes(@Nullable Boolean syncDeletes) {
            this.syncDeletes = syncDeletes;
            return this;
        }
        @CustomType.Setter
        public Builder syncProperties(@Nullable Boolean syncProperties) {
            this.syncProperties = syncProperties;
            return this;
        }
        @CustomType.Setter
        public Builder syncStatistics(@Nullable Boolean syncStatistics) {
            this.syncStatistics = syncStatistics;
            return this;
        }
        @CustomType.Setter
        public Builder url(@Nullable String url) {
            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {
            this.username = username;
            return this;
        }
        public ReplicationConfigReplication build() {
            final var o = new ReplicationConfigReplication();
            o.enabled = enabled;
            o.password = password;
            o.pathPrefix = pathPrefix;
            o.proxy = proxy;
            o.socketTimeoutMillis = socketTimeoutMillis;
            o.syncDeletes = syncDeletes;
            o.syncProperties = syncProperties;
            o.syncStatistics = syncStatistics;
            o.url = url;
            o.username = username;
            return o;
        }
    }
}
