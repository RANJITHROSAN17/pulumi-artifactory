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
public final class PushReplicationReplication {
    /**
     * @return When true, enables distributed checksum storage. For more information, see
     * [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    private @Nullable Boolean checkBinaryExistenceInFilestore;
    /**
     * @return When set, this replication will be enabled when saved.
     * 
     */
    private @Nullable Boolean enabled;
    /**
     * @return Required for local repository, but not needed for remote repository.
     * 
     */
    private String password;
    /**
     * @return Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
     * 
     */
    private @Nullable String pathPrefix;
    /**
     * @return Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
     * 
     */
    private @Nullable String proxy;
    /**
     * @return The network timeout in milliseconds to use for remote operations.
     * 
     */
    private @Nullable Integer socketTimeoutMillis;
    /**
     * @return When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
     * Note that enabling this option, will delete artifacts on the target that do not exist in the source repository.
     * 
     */
    private @Nullable Boolean syncDeletes;
    /**
     * @return When set, the task also synchronizes the properties of replicated artifacts.
     * 
     */
    private @Nullable Boolean syncProperties;
    /**
     * @return When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
     * 
     */
    private @Nullable Boolean syncStatistics;
    /**
     * @return The URL of the target local repository on a remote Artifactory server. Required for local repository, but not needed for remote repository.
     * 
     */
    private String url;
    /**
     * @return Required for local repository, but not needed for remote repository.
     * 
     */
    private String username;

    private PushReplicationReplication() {}
    /**
     * @return When true, enables distributed checksum storage. For more information, see
     * [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    public Optional<Boolean> checkBinaryExistenceInFilestore() {
        return Optional.ofNullable(this.checkBinaryExistenceInFilestore);
    }
    /**
     * @return When set, this replication will be enabled when saved.
     * 
     */
    public Optional<Boolean> enabled() {
        return Optional.ofNullable(this.enabled);
    }
    /**
     * @return Required for local repository, but not needed for remote repository.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return Only artifacts that located in path that matches the subpath within the remote repository will be replicated.
     * 
     */
    public Optional<String> pathPrefix() {
        return Optional.ofNullable(this.pathPrefix);
    }
    /**
     * @return Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
     * 
     */
    public Optional<String> proxy() {
        return Optional.ofNullable(this.proxy);
    }
    /**
     * @return The network timeout in milliseconds to use for remote operations.
     * 
     */
    public Optional<Integer> socketTimeoutMillis() {
        return Optional.ofNullable(this.socketTimeoutMillis);
    }
    /**
     * @return When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata).
     * Note that enabling this option, will delete artifacts on the target that do not exist in the source repository.
     * 
     */
    public Optional<Boolean> syncDeletes() {
        return Optional.ofNullable(this.syncDeletes);
    }
    /**
     * @return When set, the task also synchronizes the properties of replicated artifacts.
     * 
     */
    public Optional<Boolean> syncProperties() {
        return Optional.ofNullable(this.syncProperties);
    }
    /**
     * @return When set, artifact download statistics will also be replicated. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery.
     * 
     */
    public Optional<Boolean> syncStatistics() {
        return Optional.ofNullable(this.syncStatistics);
    }
    /**
     * @return The URL of the target local repository on a remote Artifactory server. Required for local repository, but not needed for remote repository.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return Required for local repository, but not needed for remote repository.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PushReplicationReplication defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean checkBinaryExistenceInFilestore;
        private @Nullable Boolean enabled;
        private String password;
        private @Nullable String pathPrefix;
        private @Nullable String proxy;
        private @Nullable Integer socketTimeoutMillis;
        private @Nullable Boolean syncDeletes;
        private @Nullable Boolean syncProperties;
        private @Nullable Boolean syncStatistics;
        private String url;
        private String username;
        public Builder() {}
        public Builder(PushReplicationReplication defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.checkBinaryExistenceInFilestore = defaults.checkBinaryExistenceInFilestore;
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
        public Builder checkBinaryExistenceInFilestore(@Nullable Boolean checkBinaryExistenceInFilestore) {
            this.checkBinaryExistenceInFilestore = checkBinaryExistenceInFilestore;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(@Nullable Boolean enabled) {
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
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
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public PushReplicationReplication build() {
            final var o = new PushReplicationReplication();
            o.checkBinaryExistenceInFilestore = checkBinaryExistenceInFilestore;
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
