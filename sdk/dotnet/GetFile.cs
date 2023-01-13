// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Artifactory
{
    public static class GetFile
    {
        /// <summary>
        /// ## # Artifactory File Data Source
        /// 
        /// Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_file = Artifactory.GetFile.Invoke(new()
        ///     {
        ///         OutputPath = "tmp/artifact.zip",
        ///         Path = "/path/to/the/artifact.zip",
        ///         Repository = "repo-key",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFileResult> InvokeAsync(GetFileArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFileResult>("artifactory:index/getFile:getFile", args ?? new GetFileArgs(), options.WithDefaults());

        /// <summary>
        /// ## # Artifactory File Data Source
        /// 
        /// Provides an Artifactory file datasource. This can be used to download a file from a given Artifactory repository.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Artifactory = Pulumi.Artifactory;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_file = Artifactory.GetFile.Invoke(new()
        ///     {
        ///         OutputPath = "tmp/artifact.zip",
        ///         Path = "/path/to/the/artifact.zip",
        ///         Repository = "repo-key",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFileResult> Invoke(GetFileInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFileResult>("artifactory:index/getFile:getFile", args ?? new GetFileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If set to true, an existing file in the output_path will be overwritten. Default: false
        /// </summary>
        [Input("forceOverwrite")]
        public bool? ForceOverwrite { get; set; }

        /// <summary>
        /// The local path the file should be downloaded to.
        /// </summary>
        [Input("outputPath", required: true)]
        public string OutputPath { get; set; } = null!;

        /// <summary>
        /// The path to the file within the repository.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        /// <summary>
        /// If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory
        /// if the file exists. More details in the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveLatestArtifact)
        /// </summary>
        [Input("pathIsAliased")]
        public bool? PathIsAliased { get; set; }

        /// <summary>
        /// Name of the repository where the file is stored.
        /// </summary>
        [Input("repository", required: true)]
        public string Repository { get; set; } = null!;

        public GetFileArgs()
        {
        }
        public static new GetFileArgs Empty => new GetFileArgs();
    }

    public sealed class GetFileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// If set to true, an existing file in the output_path will be overwritten. Default: false
        /// </summary>
        [Input("forceOverwrite")]
        public Input<bool>? ForceOverwrite { get; set; }

        /// <summary>
        /// The local path the file should be downloaded to.
        /// </summary>
        [Input("outputPath", required: true)]
        public Input<string> OutputPath { get; set; } = null!;

        /// <summary>
        /// The path to the file within the repository.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// If set to `true`, the provider will get the artifact directly from Artifactory without attempting to resolve it or verify it and will delegate this to artifactory
        /// if the file exists. More details in the [official documentation](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-RetrieveLatestArtifact)
        /// </summary>
        [Input("pathIsAliased")]
        public Input<bool>? PathIsAliased { get; set; }

        /// <summary>
        /// Name of the repository where the file is stored.
        /// </summary>
        [Input("repository", required: true)]
        public Input<string> Repository { get; set; } = null!;

        public GetFileInvokeArgs()
        {
        }
        public static new GetFileInvokeArgs Empty => new GetFileInvokeArgs();
    }


    [OutputType]
    public sealed class GetFileResult
    {
        /// <summary>
        /// The time &amp; date when the file was created.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// The user who created the file.
        /// </summary>
        public readonly string CreatedBy;
        /// <summary>
        /// The URI that can be used to download the file.
        /// </summary>
        public readonly string DownloadUri;
        public readonly bool? ForceOverwrite;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The time &amp; date when the file was last modified.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// The time &amp; date when the file was last updated.
        /// </summary>
        public readonly string LastUpdated;
        /// <summary>
        /// MD5 checksum of the file.
        /// </summary>
        public readonly string Md5;
        /// <summary>
        /// The mimetype of the file.
        /// </summary>
        public readonly string Mimetype;
        /// <summary>
        /// The user who last modified the file.
        /// </summary>
        public readonly string ModifiedBy;
        public readonly string OutputPath;
        public readonly string Path;
        public readonly bool? PathIsAliased;
        public readonly string Repository;
        /// <summary>
        /// SHA1 checksum of the file.
        /// </summary>
        public readonly string Sha1;
        /// <summary>
        /// SHA256 checksum of the file.
        /// </summary>
        public readonly string Sha256;
        /// <summary>
        /// The size of the file.
        /// </summary>
        public readonly int Size;

        [OutputConstructor]
        private GetFileResult(
            string created,

            string createdBy,

            string downloadUri,

            bool? forceOverwrite,

            string id,

            string lastModified,

            string lastUpdated,

            string md5,

            string mimetype,

            string modifiedBy,

            string outputPath,

            string path,

            bool? pathIsAliased,

            string repository,

            string sha1,

            string sha256,

            int size)
        {
            Created = created;
            CreatedBy = createdBy;
            DownloadUri = downloadUri;
            ForceOverwrite = forceOverwrite;
            Id = id;
            LastModified = lastModified;
            LastUpdated = lastUpdated;
            Md5 = md5;
            Mimetype = mimetype;
            ModifiedBy = modifiedBy;
            OutputPath = outputPath;
            Path = path;
            PathIsAliased = pathIsAliased;
            Repository = repository;
            Sha1 = sha1;
            Sha256 = sha256;
            Size = size;
        }
    }
}
