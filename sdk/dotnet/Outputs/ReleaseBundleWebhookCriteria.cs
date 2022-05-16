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
    public sealed class ReleaseBundleWebhookCriteria
    {
        /// <summary>
        /// Trigger on any release bundle.
        /// </summary>
        public readonly bool AnyReleaseBundle;
        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: "org/apache/**".
        /// </summary>
        public readonly ImmutableArray<string> ExcludePatterns;
        /// <summary>
        /// Simple comma separated wildcard patterns for repository artifact paths (with no leading slash).\n Ant-style path expressions are supported (*, *\*, ?).\nFor example: "org/apache/**".
        /// </summary>
        public readonly ImmutableArray<string> IncludePatterns;
        /// <summary>
        /// Trigger on this list of release bundle names.
        /// </summary>
        public readonly ImmutableArray<string> RegisteredReleaseBundleNames;

        [OutputConstructor]
        private ReleaseBundleWebhookCriteria(
            bool anyReleaseBundle,

            ImmutableArray<string> excludePatterns,

            ImmutableArray<string> includePatterns,

            ImmutableArray<string> registeredReleaseBundleNames)
        {
            AnyReleaseBundle = anyReleaseBundle;
            ExcludePatterns = excludePatterns;
            IncludePatterns = includePatterns;
            RegisteredReleaseBundleNames = registeredReleaseBundleNames;
        }
    }
}
