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


public final class UnmanagedUserArgs extends com.pulumi.resources.ResourceArgs {

    public static final UnmanagedUserArgs Empty = new UnmanagedUserArgs();

    /**
     * When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
     * 
     */
    @Import(name="admin")
    private @Nullable Output<Boolean> admin;

    /**
     * @return When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
     * 
     */
    public Optional<Output<Boolean>> admin() {
        return Optional.ofNullable(this.admin);
    }

    /**
     * When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
     * 
     */
    @Import(name="disableUiAccess")
    private @Nullable Output<Boolean> disableUiAccess;

    /**
     * @return When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> disableUiAccess() {
        return Optional.ofNullable(this.disableUiAccess);
    }

    /**
     * Email for user.
     * 
     */
    @Import(name="email", required=true)
    private Output<String> email;

    /**
     * @return Email for user.
     * 
     */
    public Output<String> email() {
        return this.email;
    }

    /**
     * List of groups this user is a part of.
     * 
     */
    @Import(name="groups")
    private @Nullable Output<List<String>> groups;

    /**
     * @return List of groups this user is a part of.
     * 
     */
    public Optional<Output<List<String>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
     * 
     */
    @Import(name="internalPasswordDisabled")
    private @Nullable Output<Boolean> internalPasswordDisabled;

    /**
     * @return When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
     * 
     */
    public Optional<Output<Boolean>> internalPasswordDisabled() {
        return Optional.ofNullable(this.internalPasswordDisabled);
    }

    /**
     * Username for user.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Username for user.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
     * 
     */
    @Import(name="profileUpdatable")
    private @Nullable Output<Boolean> profileUpdatable;

    /**
     * @return When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> profileUpdatable() {
        return Optional.ofNullable(this.profileUpdatable);
    }

    private UnmanagedUserArgs() {}

    private UnmanagedUserArgs(UnmanagedUserArgs $) {
        this.admin = $.admin;
        this.disableUiAccess = $.disableUiAccess;
        this.email = $.email;
        this.groups = $.groups;
        this.internalPasswordDisabled = $.internalPasswordDisabled;
        this.name = $.name;
        this.password = $.password;
        this.profileUpdatable = $.profileUpdatable;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UnmanagedUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UnmanagedUserArgs $;

        public Builder() {
            $ = new UnmanagedUserArgs();
        }

        public Builder(UnmanagedUserArgs defaults) {
            $ = new UnmanagedUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param admin When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder admin(@Nullable Output<Boolean> admin) {
            $.admin = admin;
            return this;
        }

        /**
         * @param admin When enabled, this user is an administrator with all the ensuing privileges. Default value is `false`.
         * 
         * @return builder
         * 
         */
        public Builder admin(Boolean admin) {
            return admin(Output.of(admin));
        }

        /**
         * @param disableUiAccess When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder disableUiAccess(@Nullable Output<Boolean> disableUiAccess) {
            $.disableUiAccess = disableUiAccess;
            return this;
        }

        /**
         * @param disableUiAccess When set, this user can only access Artifactory through the REST API. This option cannot be set if the user has Admin privileges. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder disableUiAccess(Boolean disableUiAccess) {
            return disableUiAccess(Output.of(disableUiAccess));
        }

        /**
         * @param email Email for user.
         * 
         * @return builder
         * 
         */
        public Builder email(Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email Email for user.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param groups List of groups this user is a part of.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable Output<List<String>> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups List of groups this user is a part of.
         * 
         * @return builder
         * 
         */
        public Builder groups(List<String> groups) {
            return groups(Output.of(groups));
        }

        /**
         * @param groups List of groups this user is a part of.
         * 
         * @return builder
         * 
         */
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param internalPasswordDisabled When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
         * 
         * @return builder
         * 
         */
        public Builder internalPasswordDisabled(@Nullable Output<Boolean> internalPasswordDisabled) {
            $.internalPasswordDisabled = internalPasswordDisabled;
            return this;
        }

        /**
         * @param internalPasswordDisabled When set, disables the fallback of using an internal password when external authentication (such as LDAP) is enabled.
         * 
         * @return builder
         * 
         */
        public Builder internalPasswordDisabled(Boolean internalPasswordDisabled) {
            return internalPasswordDisabled(Output.of(internalPasswordDisabled));
        }

        /**
         * @param name Username for user.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Username for user.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password for the user. When omitted, a random password is generated using the following password policy: 12 characters with 1 digit, 1 symbol, with upper and lower case letters.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param profileUpdatable When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder profileUpdatable(@Nullable Output<Boolean> profileUpdatable) {
            $.profileUpdatable = profileUpdatable;
            return this;
        }

        /**
         * @param profileUpdatable When set, this user can update his profile details (except for the password. Only an administrator can update the password). Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder profileUpdatable(Boolean profileUpdatable) {
            return profileUpdatable(Output.of(profileUpdatable));
        }

        public UnmanagedUserArgs build() {
            $.email = Objects.requireNonNull($.email, "expected parameter 'email' to be non-null");
            return $;
        }
    }

}
