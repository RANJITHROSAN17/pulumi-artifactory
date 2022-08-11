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


public final class ScopedTokenState extends com.pulumi.resources.ResourceArgs {

    public static final ScopedTokenState Empty = new ScopedTokenState();

    /**
     * Returns the access token to authenticate to Artifactory
     * 
     */
    @Import(name="accessToken")
    private @Nullable Output<String> accessToken;

    /**
     * @return Returns the access token to authenticate to Artifactory
     * 
     */
    public Optional<Output<String>> accessToken() {
        return Optional.ofNullable(this.accessToken);
    }

    /**
     * (Optional) A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to &#39;*@*&#39; if not set. Service ID must begin with &#39;jfrt@&#39;. For instructions to retrieve the Artifactory Service ID see this [documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GetServiceID).
     * 
     */
    @Import(name="audiences")
    private @Nullable Output<List<String>> audiences;

    /**
     * @return (Optional) A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to &#39;*@*&#39; if not set. Service ID must begin with &#39;jfrt@&#39;. For instructions to retrieve the Artifactory Service ID see this [documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GetServiceID).
     * 
     */
    public Optional<Output<List<String>>> audiences() {
        return Optional.ofNullable(this.audiences);
    }

    /**
     * (Optional) Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return (Optional) Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * (Optional) The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in `access.config.yaml`. See [API documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RevokeTokenbyIDrevoketokenbyid) for details.
     * 
     */
    @Import(name="expiresIn")
    private @Nullable Output<Integer> expiresIn;

    /**
     * @return (Optional) The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in `access.config.yaml`. See [API documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RevokeTokenbyIDrevoketokenbyid) for details.
     * 
     */
    public Optional<Output<Integer>> expiresIn() {
        return Optional.ofNullable(this.expiresIn);
    }

    /**
     * Returns the token expiry
     * 
     */
    @Import(name="expiry")
    private @Nullable Output<Integer> expiry;

    /**
     * @return Returns the token expiry
     * 
     */
    public Optional<Output<Integer>> expiry() {
        return Optional.ofNullable(this.expiry);
    }

    /**
     * Returns the token issued at date/time
     * 
     */
    @Import(name="issuedAt")
    private @Nullable Output<Integer> issuedAt;

    /**
     * @return Returns the token issued at date/time
     * 
     */
    public Optional<Output<Integer>> issuedAt() {
        return Optional.ofNullable(this.issuedAt);
    }

    /**
     * Returns the token issuer
     * 
     */
    @Import(name="issuer")
    private @Nullable Output<String> issuer;

    /**
     * @return Returns the token issuer
     * 
     */
    public Optional<Output<String>> issuer() {
        return Optional.ofNullable(this.issuer);
    }

    /**
     * (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    @Import(name="refreshable")
    private @Nullable Output<Boolean> refreshable;

    /**
     * @return (Optional) Is this token refreshable? Defaults to `false`
     * 
     */
    public Optional<Output<Boolean>> refreshable() {
        return Optional.ofNullable(this.refreshable);
    }

    /**
     * (Optional) The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
     * 
     */
    @Import(name="scopes")
    private @Nullable Output<List<String>> scopes;

    /**
     * @return (Optional) The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
     * 
     */
    public Optional<Output<List<String>>> scopes() {
        return Optional.ofNullable(this.scopes);
    }

    /**
     * Returns the token type
     * 
     */
    @Import(name="subject")
    private @Nullable Output<String> subject;

    /**
     * @return Returns the token type
     * 
     */
    public Optional<Output<String>> subject() {
        return Optional.ofNullable(this.subject);
    }

    /**
     * Returns the token type
     * 
     */
    @Import(name="tokenType")
    private @Nullable Output<String> tokenType;

    /**
     * @return Returns the token type
     * 
     */
    public Optional<Output<String>> tokenType() {
        return Optional.ofNullable(this.tokenType);
    }

    /**
     * (Optional) The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: `&lt;service-id&gt;/users/&lt;username&gt;`. Limited to 255 characters.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return (Optional) The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: `&lt;service-id&gt;/users/&lt;username&gt;`. Limited to 255 characters.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private ScopedTokenState() {}

    private ScopedTokenState(ScopedTokenState $) {
        this.accessToken = $.accessToken;
        this.audiences = $.audiences;
        this.description = $.description;
        this.expiresIn = $.expiresIn;
        this.expiry = $.expiry;
        this.issuedAt = $.issuedAt;
        this.issuer = $.issuer;
        this.refreshable = $.refreshable;
        this.scopes = $.scopes;
        this.subject = $.subject;
        this.tokenType = $.tokenType;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ScopedTokenState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ScopedTokenState $;

        public Builder() {
            $ = new ScopedTokenState();
        }

        public Builder(ScopedTokenState defaults) {
            $ = new ScopedTokenState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessToken Returns the access token to authenticate to Artifactory
         * 
         * @return builder
         * 
         */
        public Builder accessToken(@Nullable Output<String> accessToken) {
            $.accessToken = accessToken;
            return this;
        }

        /**
         * @param accessToken Returns the access token to authenticate to Artifactory
         * 
         * @return builder
         * 
         */
        public Builder accessToken(String accessToken) {
            return accessToken(Output.of(accessToken));
        }

        /**
         * @param audiences (Optional) A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to &#39;*@*&#39; if not set. Service ID must begin with &#39;jfrt@&#39;. For instructions to retrieve the Artifactory Service ID see this [documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GetServiceID).
         * 
         * @return builder
         * 
         */
        public Builder audiences(@Nullable Output<List<String>> audiences) {
            $.audiences = audiences;
            return this;
        }

        /**
         * @param audiences (Optional) A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to &#39;*@*&#39; if not set. Service ID must begin with &#39;jfrt@&#39;. For instructions to retrieve the Artifactory Service ID see this [documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GetServiceID).
         * 
         * @return builder
         * 
         */
        public Builder audiences(List<String> audiences) {
            return audiences(Output.of(audiences));
        }

        /**
         * @param audiences (Optional) A list of the other instances or services that should accept this token identified by their Service-IDs. Limited to total 255 characters. Default to &#39;*@*&#39; if not set. Service ID must begin with &#39;jfrt@&#39;. For instructions to retrieve the Artifactory Service ID see this [documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GetServiceID).
         * 
         * @return builder
         * 
         */
        public Builder audiences(String... audiences) {
            return audiences(List.of(audiences));
        }

        /**
         * @param description (Optional) Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description (Optional) Free text token description. Useful for filtering and managing tokens. Limited to 1024 characters.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param expiresIn (Optional) The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in `access.config.yaml`. See [API documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RevokeTokenbyIDrevoketokenbyid) for details.
         * 
         * @return builder
         * 
         */
        public Builder expiresIn(@Nullable Output<Integer> expiresIn) {
            $.expiresIn = expiresIn;
            return this;
        }

        /**
         * @param expiresIn (Optional) The amount of time, in seconds, it would take for the token to expire. An admin shall be able to set whether expiry is mandatory, what is the default expiry, and what is the maximum expiry allowed. Must be non-negative. Default value is based on configuration in `access.config.yaml`. See [API documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RevokeTokenbyIDrevoketokenbyid) for details.
         * 
         * @return builder
         * 
         */
        public Builder expiresIn(Integer expiresIn) {
            return expiresIn(Output.of(expiresIn));
        }

        /**
         * @param expiry Returns the token expiry
         * 
         * @return builder
         * 
         */
        public Builder expiry(@Nullable Output<Integer> expiry) {
            $.expiry = expiry;
            return this;
        }

        /**
         * @param expiry Returns the token expiry
         * 
         * @return builder
         * 
         */
        public Builder expiry(Integer expiry) {
            return expiry(Output.of(expiry));
        }

        /**
         * @param issuedAt Returns the token issued at date/time
         * 
         * @return builder
         * 
         */
        public Builder issuedAt(@Nullable Output<Integer> issuedAt) {
            $.issuedAt = issuedAt;
            return this;
        }

        /**
         * @param issuedAt Returns the token issued at date/time
         * 
         * @return builder
         * 
         */
        public Builder issuedAt(Integer issuedAt) {
            return issuedAt(Output.of(issuedAt));
        }

        /**
         * @param issuer Returns the token issuer
         * 
         * @return builder
         * 
         */
        public Builder issuer(@Nullable Output<String> issuer) {
            $.issuer = issuer;
            return this;
        }

        /**
         * @param issuer Returns the token issuer
         * 
         * @return builder
         * 
         */
        public Builder issuer(String issuer) {
            return issuer(Output.of(issuer));
        }

        /**
         * @param refreshable (Optional) Is this token refreshable? Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder refreshable(@Nullable Output<Boolean> refreshable) {
            $.refreshable = refreshable;
            return this;
        }

        /**
         * @param refreshable (Optional) Is this token refreshable? Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder refreshable(Boolean refreshable) {
            return refreshable(Output.of(refreshable));
        }

        /**
         * @param scopes (Optional) The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
         * 
         * @return builder
         * 
         */
        public Builder scopes(@Nullable Output<List<String>> scopes) {
            $.scopes = scopes;
            return this;
        }

        /**
         * @param scopes (Optional) The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
         * 
         * @return builder
         * 
         */
        public Builder scopes(List<String> scopes) {
            return scopes(Output.of(scopes));
        }

        /**
         * @param scopes (Optional) The scope of access that the token provides. Access to the REST API is always provided by default. Administrators can set any scope, while non-admin users can only set the scope to a subset of the groups to which they belong.
         * 
         * @return builder
         * 
         */
        public Builder scopes(String... scopes) {
            return scopes(List.of(scopes));
        }

        /**
         * @param subject Returns the token type
         * 
         * @return builder
         * 
         */
        public Builder subject(@Nullable Output<String> subject) {
            $.subject = subject;
            return this;
        }

        /**
         * @param subject Returns the token type
         * 
         * @return builder
         * 
         */
        public Builder subject(String subject) {
            return subject(Output.of(subject));
        }

        /**
         * @param tokenType Returns the token type
         * 
         * @return builder
         * 
         */
        public Builder tokenType(@Nullable Output<String> tokenType) {
            $.tokenType = tokenType;
            return this;
        }

        /**
         * @param tokenType Returns the token type
         * 
         * @return builder
         * 
         */
        public Builder tokenType(String tokenType) {
            return tokenType(Output.of(tokenType));
        }

        /**
         * @param username (Optional) The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: `&lt;service-id&gt;/users/&lt;username&gt;`. Limited to 255 characters.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username (Optional) The user name for which this token is created. The username is based on the authenticated user - either from the user of the authenticated token or based on the username (if basic auth was used). The username is then used to set the subject of the token: `&lt;service-id&gt;/users/&lt;username&gt;`. Limited to 255 characters.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ScopedTokenState build() {
            return $;
        }
    }

}
