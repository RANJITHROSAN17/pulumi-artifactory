// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LocalRepositoryMultiReplicationReplicationArgs extends com.pulumi.resources.ResourceArgs {

    public static final LocalRepositoryMultiReplicationReplicationArgs Empty = new LocalRepositoryMultiReplicationReplicationArgs();

    /**
     * Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    @Import(name="checkBinaryExistenceInFilestore")
    private @Nullable Output<Boolean> checkBinaryExistenceInFilestore;

    /**
     * @return Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
     * 
     */
    public Optional<Output<Boolean>> checkBinaryExistenceInFilestore() {
        return Optional.ofNullable(this.checkBinaryExistenceInFilestore);
    }

    /**
     * When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**{@literal /}z/*`. By default, no artifacts are excluded.
     * 
     */
    @Import(name="excludePathPrefixPattern")
    private @Nullable Output<String> excludePathPrefixPattern;

    /**
     * @return List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**{@literal /}z/*`. By default, no artifacts are excluded.
     * 
     */
    public Optional<Output<String>> excludePathPrefixPattern() {
        return Optional.ofNullable(this.excludePathPrefixPattern);
    }

    /**
     * List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**{@literal /}z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**{@literal /}*)`.
     * 
     */
    @Import(name="includePathPrefixPattern")
    private @Nullable Output<String> includePathPrefixPattern;

    /**
     * @return List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**{@literal /}z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**{@literal /}*)`.
     * 
     */
    public Optional<Output<String>> includePathPrefixPattern() {
        return Optional.ofNullable(this.includePathPrefixPattern);
    }

    /**
     * Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
     * 
     */
    @Import(name="proxy")
    private @Nullable Output<String> proxy;

    /**
     * @return Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
     * 
     */
    public Optional<Output<String>> proxy() {
        return Optional.ofNullable(this.proxy);
    }

    /**
     * Replication ID, the value is unknown until the resource is created. Can&#39;t be set or updated.
     * 
     */
    @Import(name="replicationKey")
    private @Nullable Output<String> replicationKey;

    /**
     * @return Replication ID, the value is unknown until the resource is created. Can&#39;t be set or updated.
     * 
     */
    public Optional<Output<String>> replicationKey() {
        return Optional.ofNullable(this.replicationKey);
    }

    /**
     * The network timeout in milliseconds to use for remote operations. Default value is `15000`.
     * 
     */
    @Import(name="socketTimeoutMillis")
    private @Nullable Output<Integer> socketTimeoutMillis;

    /**
     * @return The network timeout in milliseconds to use for remote operations. Default value is `15000`.
     * 
     */
    public Optional<Output<Integer>> socketTimeoutMillis() {
        return Optional.ofNullable(this.socketTimeoutMillis);
    }

    /**
     * When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
     * 
     */
    @Import(name="syncDeletes")
    private @Nullable Output<Boolean> syncDeletes;

    /**
     * @return When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> syncDeletes() {
        return Optional.ofNullable(this.syncDeletes);
    }

    /**
     * When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
     * 
     */
    @Import(name="syncProperties")
    private @Nullable Output<Boolean> syncProperties;

    /**
     * @return When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> syncProperties() {
        return Optional.ofNullable(this.syncProperties);
    }

    /**
     * When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
     * 
     */
    @Import(name="syncStatistics")
    private @Nullable Output<Boolean> syncStatistics;

    /**
     * @return When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
     * 
     */
    public Optional<Output<Boolean>> syncStatistics() {
        return Optional.ofNullable(this.syncStatistics);
    }

    /**
     * The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * Username on the remote Artifactory instance.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return Username on the remote Artifactory instance.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private LocalRepositoryMultiReplicationReplicationArgs() {}

    private LocalRepositoryMultiReplicationReplicationArgs(LocalRepositoryMultiReplicationReplicationArgs $) {
        this.checkBinaryExistenceInFilestore = $.checkBinaryExistenceInFilestore;
        this.enabled = $.enabled;
        this.excludePathPrefixPattern = $.excludePathPrefixPattern;
        this.includePathPrefixPattern = $.includePathPrefixPattern;
        this.password = $.password;
        this.proxy = $.proxy;
        this.replicationKey = $.replicationKey;
        this.socketTimeoutMillis = $.socketTimeoutMillis;
        this.syncDeletes = $.syncDeletes;
        this.syncProperties = $.syncProperties;
        this.syncStatistics = $.syncStatistics;
        this.url = $.url;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LocalRepositoryMultiReplicationReplicationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LocalRepositoryMultiReplicationReplicationArgs $;

        public Builder() {
            $ = new LocalRepositoryMultiReplicationReplicationArgs();
        }

        public Builder(LocalRepositoryMultiReplicationReplicationArgs defaults) {
            $ = new LocalRepositoryMultiReplicationReplicationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param checkBinaryExistenceInFilestore Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
         * 
         * @return builder
         * 
         */
        public Builder checkBinaryExistenceInFilestore(@Nullable Output<Boolean> checkBinaryExistenceInFilestore) {
            $.checkBinaryExistenceInFilestore = checkBinaryExistenceInFilestore;
            return this;
        }

        /**
         * @param checkBinaryExistenceInFilestore Enabling the `check_binary_existence_in_filestore` flag requires an Enterprise Plus license. When true, enables distributed checksum storage. For more information, see [Optimizing Repository Replication with Checksum-Based Storage](https://www.jfrog.com/confluence/display/JFROG/Repository+Replication#RepositoryReplication-OptimizingRepositoryReplicationUsingStorageLevelSynchronizationOptions).
         * 
         * @return builder
         * 
         */
        public Builder checkBinaryExistenceInFilestore(Boolean checkBinaryExistenceInFilestore) {
            return checkBinaryExistenceInFilestore(Output.of(checkBinaryExistenceInFilestore));
        }

        /**
         * @param enabled When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled When set, enables replication of this repository to the target specified in `url` attribute. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param excludePathPrefixPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**{@literal /}z/*`. By default, no artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludePathPrefixPattern(@Nullable Output<String> excludePathPrefixPattern) {
            $.excludePathPrefixPattern = excludePathPrefixPattern;
            return this;
        }

        /**
         * @param excludePathPrefixPattern List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**{@literal /}z/*`. By default, no artifacts are excluded.
         * 
         * @return builder
         * 
         */
        public Builder excludePathPrefixPattern(String excludePathPrefixPattern) {
            return excludePathPrefixPattern(Output.of(excludePathPrefixPattern));
        }

        /**
         * @param includePathPrefixPattern List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**{@literal /}z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**{@literal /}*)`.
         * 
         * @return builder
         * 
         */
        public Builder includePathPrefixPattern(@Nullable Output<String> includePathPrefixPattern) {
            $.includePathPrefixPattern = includePathPrefixPattern;
            return this;
        }

        /**
         * @param includePathPrefixPattern List of artifact patterns to include when evaluating artifact requests in the form of `x/y/**{@literal /}z/*`. When used, only artifacts matching one of the include patterns are served. By default, all artifacts are included `(**{@literal /}*)`.
         * 
         * @return builder
         * 
         */
        public Builder includePathPrefixPattern(String includePathPrefixPattern) {
            return includePathPrefixPattern(Output.of(includePathPrefixPattern));
        }

        /**
         * @param password Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Use either the HTTP authentication password or [identity token](https://www.jfrog.com/confluence/display/JFROG/User+Profile#UserProfile-IdentityTokenidentitytoken).
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param proxy Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
         * 
         * @return builder
         * 
         */
        public Builder proxy(@Nullable Output<String> proxy) {
            $.proxy = proxy;
            return this;
        }

        /**
         * @param proxy Proxy key from Artifactory Proxies settings. The proxy configuration will be used when communicating with the remote instance.
         * 
         * @return builder
         * 
         */
        public Builder proxy(String proxy) {
            return proxy(Output.of(proxy));
        }

        /**
         * @param replicationKey Replication ID, the value is unknown until the resource is created. Can&#39;t be set or updated.
         * 
         * @return builder
         * 
         */
        public Builder replicationKey(@Nullable Output<String> replicationKey) {
            $.replicationKey = replicationKey;
            return this;
        }

        /**
         * @param replicationKey Replication ID, the value is unknown until the resource is created. Can&#39;t be set or updated.
         * 
         * @return builder
         * 
         */
        public Builder replicationKey(String replicationKey) {
            return replicationKey(Output.of(replicationKey));
        }

        /**
         * @param socketTimeoutMillis The network timeout in milliseconds to use for remote operations. Default value is `15000`.
         * 
         * @return builder
         * 
         */
        public Builder socketTimeoutMillis(@Nullable Output<Integer> socketTimeoutMillis) {
            $.socketTimeoutMillis = socketTimeoutMillis;
            return this;
        }

        /**
         * @param socketTimeoutMillis The network timeout in milliseconds to use for remote operations. Default value is `15000`.
         * 
         * @return builder
         * 
         */
        public Builder socketTimeoutMillis(Integer socketTimeoutMillis) {
            return socketTimeoutMillis(Output.of(socketTimeoutMillis));
        }

        /**
         * @param syncDeletes When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder syncDeletes(@Nullable Output<Boolean> syncDeletes) {
            $.syncDeletes = syncDeletes;
            return this;
        }

        /**
         * @param syncDeletes When set, items that were deleted locally should also be deleted remotely (also applies to properties metadata). Note that enabling this option, will delete artifacts on the target that do not exist in the source repository. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder syncDeletes(Boolean syncDeletes) {
            return syncDeletes(Output.of(syncDeletes));
        }

        /**
         * @param syncProperties When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder syncProperties(@Nullable Output<Boolean> syncProperties) {
            $.syncProperties = syncProperties;
            return this;
        }

        /**
         * @param syncProperties When set, the task also synchronizes the properties of replicated artifacts. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder syncProperties(Boolean syncProperties) {
            return syncProperties(Output.of(syncProperties));
        }

        /**
         * @param syncStatistics When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
         * 
         * @return builder
         * 
         */
        public Builder syncStatistics(@Nullable Output<Boolean> syncStatistics) {
            $.syncStatistics = syncStatistics;
            return this;
        }

        /**
         * @param syncStatistics When set, the task also synchronizes artifact download statistics. Set to avoid inadvertent cleanup at the target instance when setting up replication for disaster recovery. Default value is `false`
         * 
         * @return builder
         * 
         */
        public Builder syncStatistics(Boolean syncStatistics) {
            return syncStatistics(Output.of(syncStatistics));
        }

        /**
         * @param url The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL of the target local repository on a remote Artifactory server. Use the format `https://&lt;artifactory_url&gt;/artifactory/&lt;repository_name&gt;`.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param username Username on the remote Artifactory instance.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Username on the remote Artifactory instance.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public LocalRepositoryMultiReplicationReplicationArgs build() {
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}
