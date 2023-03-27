// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class GetFederatedHelmRepositoryMemberArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Represents the active state of the federated member. It is supported to change the enabled
        /// status of my own member. The config will be updated on the other federated members automatically.
        /// </summary>
        [Input("enabled", required: true)]
        public bool Enabled { get; set; }

        /// <summary>
        /// Full URL to ending with the repository name.
        /// </summary>
        [Input("url", required: true)]
        public string Url { get; set; } = null!;

        public GetFederatedHelmRepositoryMemberArgs()
        {
        }
        public static new GetFederatedHelmRepositoryMemberArgs Empty => new GetFederatedHelmRepositoryMemberArgs();
    }
}
