// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Outputs
{

    [OutputType]
    public sealed class PermissionTargetsBuildActionsUser
    {
        /// <summary>
        /// Name of permission.
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<string> Permissions;

        [OutputConstructor]
        private PermissionTargetsBuildActionsUser(
            string name,

            ImmutableArray<string> permissions)
        {
            Name = name;
            Permissions = permissions;
        }
    }
}
