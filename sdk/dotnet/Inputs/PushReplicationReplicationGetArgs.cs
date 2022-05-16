// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class PushReplicationReplicationGetArgs : Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        [Input("pathPrefix")]
        public Input<string>? PathPrefix { get; set; }

        /// <summary>
        /// Proxy key from Artifactory Proxies setting.
        /// </summary>
        [Input("proxy")]
        public Input<string>? Proxy { get; set; }

        [Input("socketTimeoutMillis")]
        public Input<int>? SocketTimeoutMillis { get; set; }

        [Input("syncDeletes")]
        public Input<bool>? SyncDeletes { get; set; }

        [Input("syncProperties")]
        public Input<bool>? SyncProperties { get; set; }

        [Input("syncStatistics")]
        public Input<bool>? SyncStatistics { get; set; }

        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public PushReplicationReplicationGetArgs()
        {
        }
    }
}
