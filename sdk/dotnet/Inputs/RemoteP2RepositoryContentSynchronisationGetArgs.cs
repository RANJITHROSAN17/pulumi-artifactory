// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory.Inputs
{

    public sealed class RemoteP2RepositoryContentSynchronisationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("propertiesEnabled")]
        public Input<bool>? PropertiesEnabled { get; set; }

        [Input("sourceOriginAbsenceDetection")]
        public Input<bool>? SourceOriginAbsenceDetection { get; set; }

        [Input("statisticsEnabled")]
        public Input<bool>? StatisticsEnabled { get; set; }

        public RemoteP2RepositoryContentSynchronisationGetArgs()
        {
        }
        public static new RemoteP2RepositoryContentSynchronisationGetArgs Empty => new RemoteP2RepositoryContentSynchronisationGetArgs();
    }
}
