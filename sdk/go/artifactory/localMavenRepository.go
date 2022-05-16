// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a local Maven repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewLocalMavenRepository(ctx, "terraform-local-test-maven-repo-basic", &artifactory.LocalMavenRepositoryArgs{
// 			ChecksumPolicyType:           pulumi.String("client-checksums"),
// 			HandleReleases:               pulumi.Bool(true),
// 			HandleSnapshots:              pulumi.Bool(true),
// 			Key:                          pulumi.String("terraform-local-test-maven-repo-basic"),
// 			MaxUniqueSnapshots:           pulumi.Int(10),
// 			SnapshotVersionBehavior:      pulumi.String("unique"),
// 			SuppressPomConsistencyChecks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Local repositories can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/localMavenRepository:LocalMavenRepository terraform-local-test-maven-repo-basic terraform-local-test-maven-repo-basic
// ```
type LocalMavenRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details,
	// please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
	ChecksumPolicyType pulumi.StringPtrOutput `pulumi:"checksumPolicyType"`
	Description        pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
	HandleReleases pulumi.BoolPtrOutput `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
	HandleSnapshots pulumi.BoolPtrOutput `pulumi:"handleSnapshots"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store.
	// Once the number of snapshots exceeds this setting, older versions are removed.
	// A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrOutput    `pulumi:"maxUniqueSnapshots"`
	Notes              pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType        pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// * `unique`: Version number is based on a time-stamp (default)
	// * `non-unique`: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// * `deployer`: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior pulumi.StringPtrOutput `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
	SuppressPomConsistencyChecks pulumi.BoolPtrOutput `pulumi:"suppressPomConsistencyChecks"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewLocalMavenRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalMavenRepository(ctx *pulumi.Context,
	name string, args *LocalMavenRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalMavenRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalMavenRepository
	err := ctx.RegisterResource("artifactory:index/localMavenRepository:LocalMavenRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalMavenRepository gets an existing LocalMavenRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalMavenRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalMavenRepositoryState, opts ...pulumi.ResourceOption) (*LocalMavenRepository, error) {
	var resource LocalMavenRepository
	err := ctx.ReadResource("artifactory:index/localMavenRepository:LocalMavenRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalMavenRepository resources.
type localMavenRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details,
	// please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	Description        *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
	HandleReleases *bool `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
	HandleSnapshots *bool `pulumi:"handleSnapshots"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store.
	// Once the number of snapshots exceeds this setting, older versions are removed.
	// A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int    `pulumi:"maxUniqueSnapshots"`
	Notes              *string `pulumi:"notes"`
	PackageType        *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// * `unique`: Version number is based on a time-stamp (default)
	// * `non-unique`: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// * `deployer`: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior *string `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
	SuppressPomConsistencyChecks *bool `pulumi:"suppressPomConsistencyChecks"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalMavenRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details,
	// please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
	ChecksumPolicyType pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
	HandleReleases pulumi.BoolPtrInput
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
	HandleSnapshots pulumi.BoolPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// The maximum number of unique snapshots of a single artifact to store.
	// Once the number of snapshots exceeds this setting, older versions are removed.
	// A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	Notes              pulumi.StringPtrInput
	PackageType        pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// * `unique`: Version number is based on a time-stamp (default)
	// * `non-unique`: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// * `deployer`: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior pulumi.StringPtrInput
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
	SuppressPomConsistencyChecks pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalMavenRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localMavenRepositoryState)(nil)).Elem()
}

type localMavenRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details,
	// please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
	ChecksumPolicyType *string `pulumi:"checksumPolicyType"`
	Description        *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
	HandleReleases *bool `pulumi:"handleReleases"`
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
	HandleSnapshots *bool `pulumi:"handleSnapshots"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The maximum number of unique snapshots of a single artifact to store.
	// Once the number of snapshots exceeds this setting, older versions are removed.
	// A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int    `pulumi:"maxUniqueSnapshots"`
	Notes              *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// * `unique`: Version number is based on a time-stamp (default)
	// * `non-unique`: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// * `deployer`: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior *string `pulumi:"snapshotVersionBehavior"`
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
	SuppressPomConsistencyChecks *bool `pulumi:"suppressPomConsistencyChecks"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalMavenRepository resource.
type LocalMavenRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
	// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
	// `client-checksums` and `generated-checksums`. For more details,
	// please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
	ChecksumPolicyType pulumi.StringPtrInput
	Description        pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
	HandleReleases pulumi.BoolPtrInput
	// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
	HandleSnapshots pulumi.BoolPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// The maximum number of unique snapshots of a single artifact to store.
	// Once the number of snapshots exceeds this setting, older versions are removed.
	// A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	Notes              pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
	// with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Specifies the naming convention for Maven SNAPSHOT versions.
	// The options are -
	// * `unique`: Version number is based on a time-stamp (default)
	// * `non-unique`: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
	// * `deployer`: Respects the settings in the Maven client that is deploying the artifact.
	SnapshotVersionBehavior pulumi.StringPtrInput
	// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
	// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
	// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
	SuppressPomConsistencyChecks pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalMavenRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localMavenRepositoryArgs)(nil)).Elem()
}

type LocalMavenRepositoryInput interface {
	pulumi.Input

	ToLocalMavenRepositoryOutput() LocalMavenRepositoryOutput
	ToLocalMavenRepositoryOutputWithContext(ctx context.Context) LocalMavenRepositoryOutput
}

func (*LocalMavenRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalMavenRepository)(nil)).Elem()
}

func (i *LocalMavenRepository) ToLocalMavenRepositoryOutput() LocalMavenRepositoryOutput {
	return i.ToLocalMavenRepositoryOutputWithContext(context.Background())
}

func (i *LocalMavenRepository) ToLocalMavenRepositoryOutputWithContext(ctx context.Context) LocalMavenRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalMavenRepositoryOutput)
}

// LocalMavenRepositoryArrayInput is an input type that accepts LocalMavenRepositoryArray and LocalMavenRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalMavenRepositoryArrayInput` via:
//
//          LocalMavenRepositoryArray{ LocalMavenRepositoryArgs{...} }
type LocalMavenRepositoryArrayInput interface {
	pulumi.Input

	ToLocalMavenRepositoryArrayOutput() LocalMavenRepositoryArrayOutput
	ToLocalMavenRepositoryArrayOutputWithContext(context.Context) LocalMavenRepositoryArrayOutput
}

type LocalMavenRepositoryArray []LocalMavenRepositoryInput

func (LocalMavenRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalMavenRepository)(nil)).Elem()
}

func (i LocalMavenRepositoryArray) ToLocalMavenRepositoryArrayOutput() LocalMavenRepositoryArrayOutput {
	return i.ToLocalMavenRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalMavenRepositoryArray) ToLocalMavenRepositoryArrayOutputWithContext(ctx context.Context) LocalMavenRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalMavenRepositoryArrayOutput)
}

// LocalMavenRepositoryMapInput is an input type that accepts LocalMavenRepositoryMap and LocalMavenRepositoryMapOutput values.
// You can construct a concrete instance of `LocalMavenRepositoryMapInput` via:
//
//          LocalMavenRepositoryMap{ "key": LocalMavenRepositoryArgs{...} }
type LocalMavenRepositoryMapInput interface {
	pulumi.Input

	ToLocalMavenRepositoryMapOutput() LocalMavenRepositoryMapOutput
	ToLocalMavenRepositoryMapOutputWithContext(context.Context) LocalMavenRepositoryMapOutput
}

type LocalMavenRepositoryMap map[string]LocalMavenRepositoryInput

func (LocalMavenRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalMavenRepository)(nil)).Elem()
}

func (i LocalMavenRepositoryMap) ToLocalMavenRepositoryMapOutput() LocalMavenRepositoryMapOutput {
	return i.ToLocalMavenRepositoryMapOutputWithContext(context.Background())
}

func (i LocalMavenRepositoryMap) ToLocalMavenRepositoryMapOutputWithContext(ctx context.Context) LocalMavenRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalMavenRepositoryMapOutput)
}

type LocalMavenRepositoryOutput struct{ *pulumi.OutputState }

func (LocalMavenRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalMavenRepository)(nil)).Elem()
}

func (o LocalMavenRepositoryOutput) ToLocalMavenRepositoryOutput() LocalMavenRepositoryOutput {
	return o
}

func (o LocalMavenRepositoryOutput) ToLocalMavenRepositoryOutputWithContext(ctx context.Context) LocalMavenRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalMavenRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalMavenRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// Checksum policy determines how Artifactory behaves when a client checksum for a deployed
// resource is missing or conflicts with the locally calculated checksum (bad checksum). The options are
// `client-checksums` and `generated-checksums`. For more details,
// please refer to [Checksum Policy](https://www.jfrog.com/confluence/display/JFROG/Local+Repositories#LocalRepositories-ChecksumPolicy).
func (o LocalMavenRepositoryOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringPtrOutput { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LocalMavenRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalMavenRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o LocalMavenRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// If set, Artifactory allows you to deploy release artifacts into this repository. Default is `true`.
func (o LocalMavenRepositoryOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.BoolPtrOutput { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

// If set, Artifactory allows you to deploy snapshot artifacts into this repository. Default is `true`.
func (o LocalMavenRepositoryOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.BoolPtrOutput { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalMavenRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalMavenRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique snapshots of a single artifact to store.
// Once the number of snapshots exceeds this setting, older versions are removed.
// A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
func (o LocalMavenRepositoryOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.IntPtrOutput { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

func (o LocalMavenRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalMavenRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalMavenRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
func (o LocalMavenRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
// with project key, separated by a dash.
func (o LocalMavenRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalMavenRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalMavenRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Specifies the naming convention for Maven SNAPSHOT versions.
// The options are -
// * `unique`: Version number is based on a time-stamp (default)
// * `non-unique`: Version number uses a self-overriding naming pattern of artifactId-version-SNAPSHOT.type
// * `deployer`: Respects the settings in the Maven client that is deploying the artifact.
func (o LocalMavenRepositoryOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.StringPtrOutput { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

// By default, Artifactory keeps your repositories healthy by refusing POMs with incorrect coordinates (path).
// If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error.
// You can disable this behavior by setting the Suppress POM Consistency Checks checkbox. False by default for Maven repository.
func (o LocalMavenRepositoryOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.BoolPtrOutput { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalMavenRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalMavenRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalMavenRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalMavenRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalMavenRepository)(nil)).Elem()
}

func (o LocalMavenRepositoryArrayOutput) ToLocalMavenRepositoryArrayOutput() LocalMavenRepositoryArrayOutput {
	return o
}

func (o LocalMavenRepositoryArrayOutput) ToLocalMavenRepositoryArrayOutputWithContext(ctx context.Context) LocalMavenRepositoryArrayOutput {
	return o
}

func (o LocalMavenRepositoryArrayOutput) Index(i pulumi.IntInput) LocalMavenRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalMavenRepository {
		return vs[0].([]*LocalMavenRepository)[vs[1].(int)]
	}).(LocalMavenRepositoryOutput)
}

type LocalMavenRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalMavenRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalMavenRepository)(nil)).Elem()
}

func (o LocalMavenRepositoryMapOutput) ToLocalMavenRepositoryMapOutput() LocalMavenRepositoryMapOutput {
	return o
}

func (o LocalMavenRepositoryMapOutput) ToLocalMavenRepositoryMapOutputWithContext(ctx context.Context) LocalMavenRepositoryMapOutput {
	return o
}

func (o LocalMavenRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalMavenRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalMavenRepository {
		return vs[0].(map[string]*LocalMavenRepository)[vs[1].(string)]
	}).(LocalMavenRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalMavenRepositoryInput)(nil)).Elem(), &LocalMavenRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalMavenRepositoryArrayInput)(nil)).Elem(), LocalMavenRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalMavenRepositoryMapInput)(nil)).Elem(), LocalMavenRepositoryMap{})
	pulumi.RegisterOutputType(LocalMavenRepositoryOutput{})
	pulumi.RegisterOutputType(LocalMavenRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalMavenRepositoryMapOutput{})
}
