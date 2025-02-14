// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.artifactory.outputs;

import com.pulumi.artifactory.outputs.GetPermissionTargetBuildActionsGroup;
import com.pulumi.artifactory.outputs.GetPermissionTargetBuildActionsUser;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetPermissionTargetBuildActions {
    /**
     * @return Groups this permission applies for.
     * 
     */
    private @Nullable List<GetPermissionTargetBuildActionsGroup> groups;
    /**
     * @return Users this permission target applies for.
     * 
     */
    private @Nullable List<GetPermissionTargetBuildActionsUser> users;

    private GetPermissionTargetBuildActions() {}
    /**
     * @return Groups this permission applies for.
     * 
     */
    public List<GetPermissionTargetBuildActionsGroup> groups() {
        return this.groups == null ? List.of() : this.groups;
    }
    /**
     * @return Users this permission target applies for.
     * 
     */
    public List<GetPermissionTargetBuildActionsUser> users() {
        return this.users == null ? List.of() : this.users;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPermissionTargetBuildActions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<GetPermissionTargetBuildActionsGroup> groups;
        private @Nullable List<GetPermissionTargetBuildActionsUser> users;
        public Builder() {}
        public Builder(GetPermissionTargetBuildActions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groups = defaults.groups;
    	      this.users = defaults.users;
        }

        @CustomType.Setter
        public Builder groups(@Nullable List<GetPermissionTargetBuildActionsGroup> groups) {
            this.groups = groups;
            return this;
        }
        public Builder groups(GetPermissionTargetBuildActionsGroup... groups) {
            return groups(List.of(groups));
        }
        @CustomType.Setter
        public Builder users(@Nullable List<GetPermissionTargetBuildActionsUser> users) {
            this.users = users;
            return this;
        }
        public Builder users(GetPermissionTargetBuildActionsUser... users) {
            return users(List.of(users));
        }
        public GetPermissionTargetBuildActions build() {
            final var o = new GetPermissionTargetBuildActions();
            o.groups = groups;
            o.users = users;
            return o;
        }
    }
}
