// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PermissionTargetsRepoActionsGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("groups")]
        private InputList<Inputs.PermissionTargetsRepoActionsGroupGetArgs>? _groups;

        /// <summary>
        /// Groups this permission applies for.
        /// </summary>
        public InputList<Inputs.PermissionTargetsRepoActionsGroupGetArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.PermissionTargetsRepoActionsGroupGetArgs>());
            set => _groups = value;
        }

        [Input("users")]
        private InputList<Inputs.PermissionTargetsRepoActionsUserGetArgs>? _users;

        /// <summary>
        /// Users this permission target applies for.
        /// </summary>
        public InputList<Inputs.PermissionTargetsRepoActionsUserGetArgs> Users
        {
            get => _users ?? (_users = new InputList<Inputs.PermissionTargetsRepoActionsUserGetArgs>());
            set => _users = value;
        }

        public PermissionTargetsRepoActionsGetArgs()
        {
        }
        public static new PermissionTargetsRepoActionsGetArgs Empty => new PermissionTargetsRepoActionsGetArgs();
    }
}
