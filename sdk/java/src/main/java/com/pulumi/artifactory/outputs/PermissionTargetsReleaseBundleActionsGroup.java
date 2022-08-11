// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class PermissionTargetsReleaseBundleActionsGroup {
    /**
     * @return Name of permission.
     * 
     */
    private final String name;
    private final List<String> permissions;

    @CustomType.Constructor
    private PermissionTargetsReleaseBundleActionsGroup(
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("permissions") List<String> permissions) {
        this.name = name;
        this.permissions = permissions;
    }

    /**
     * @return Name of permission.
     * 
     */
    public String name() {
        return this.name;
    }
    public List<String> permissions() {
        return this.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PermissionTargetsReleaseBundleActionsGroup defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String name;
        private List<String> permissions;

        public Builder() {
    	      // Empty
        }

        public Builder(PermissionTargetsReleaseBundleActionsGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.permissions = defaults.permissions;
        }

        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder permissions(List<String> permissions) {
            this.permissions = Objects.requireNonNull(permissions);
            return this;
        }
        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }        public PermissionTargetsReleaseBundleActionsGroup build() {
            return new PermissionTargetsReleaseBundleActionsGroup(name, permissions);
        }
    }
}
