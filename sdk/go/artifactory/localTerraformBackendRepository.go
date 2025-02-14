// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewLocalTerraformBackendRepository(ctx, "terraform-local-test-terraformbackend-repo", &artifactory.LocalTerraformBackendRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-terraformbackend-repo"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Local repositories can be imported using their name, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/localTerraformBackendRepository:LocalTerraformBackendRepository terraform-local-test-terraformbackend-repo terraform-local-test-terraformbackend-repo
//
// ```
type LocalTerraformBackendRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrOutput `pulumi:"cdnRedirect"`
	// Public description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// Internal description.
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewLocalTerraformBackendRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalTerraformBackendRepository(ctx *pulumi.Context,
	name string, args *LocalTerraformBackendRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalTerraformBackendRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalTerraformBackendRepository
	err := ctx.RegisterResource("artifactory:index/localTerraformBackendRepository:LocalTerraformBackendRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalTerraformBackendRepository gets an existing LocalTerraformBackendRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalTerraformBackendRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalTerraformBackendRepositoryState, opts ...pulumi.ResourceOption) (*LocalTerraformBackendRepository, error) {
	var resource LocalTerraformBackendRepository
	err := ctx.ReadResource("artifactory:index/localTerraformBackendRepository:LocalTerraformBackendRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalTerraformBackendRepository resources.
type localTerraformBackendRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// Internal description.
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalTerraformBackendRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// Internal description.
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalTerraformBackendRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localTerraformBackendRepositoryState)(nil)).Elem()
}

type localTerraformBackendRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// Internal description.
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalTerraformBackendRepository resource.
type LocalTerraformBackendRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// Internal description.
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
	// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
	// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
	// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalTerraformBackendRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localTerraformBackendRepositoryArgs)(nil)).Elem()
}

type LocalTerraformBackendRepositoryInput interface {
	pulumi.Input

	ToLocalTerraformBackendRepositoryOutput() LocalTerraformBackendRepositoryOutput
	ToLocalTerraformBackendRepositoryOutputWithContext(ctx context.Context) LocalTerraformBackendRepositoryOutput
}

func (*LocalTerraformBackendRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTerraformBackendRepository)(nil)).Elem()
}

func (i *LocalTerraformBackendRepository) ToLocalTerraformBackendRepositoryOutput() LocalTerraformBackendRepositoryOutput {
	return i.ToLocalTerraformBackendRepositoryOutputWithContext(context.Background())
}

func (i *LocalTerraformBackendRepository) ToLocalTerraformBackendRepositoryOutputWithContext(ctx context.Context) LocalTerraformBackendRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTerraformBackendRepositoryOutput)
}

// LocalTerraformBackendRepositoryArrayInput is an input type that accepts LocalTerraformBackendRepositoryArray and LocalTerraformBackendRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalTerraformBackendRepositoryArrayInput` via:
//
//	LocalTerraformBackendRepositoryArray{ LocalTerraformBackendRepositoryArgs{...} }
type LocalTerraformBackendRepositoryArrayInput interface {
	pulumi.Input

	ToLocalTerraformBackendRepositoryArrayOutput() LocalTerraformBackendRepositoryArrayOutput
	ToLocalTerraformBackendRepositoryArrayOutputWithContext(context.Context) LocalTerraformBackendRepositoryArrayOutput
}

type LocalTerraformBackendRepositoryArray []LocalTerraformBackendRepositoryInput

func (LocalTerraformBackendRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalTerraformBackendRepository)(nil)).Elem()
}

func (i LocalTerraformBackendRepositoryArray) ToLocalTerraformBackendRepositoryArrayOutput() LocalTerraformBackendRepositoryArrayOutput {
	return i.ToLocalTerraformBackendRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalTerraformBackendRepositoryArray) ToLocalTerraformBackendRepositoryArrayOutputWithContext(ctx context.Context) LocalTerraformBackendRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTerraformBackendRepositoryArrayOutput)
}

// LocalTerraformBackendRepositoryMapInput is an input type that accepts LocalTerraformBackendRepositoryMap and LocalTerraformBackendRepositoryMapOutput values.
// You can construct a concrete instance of `LocalTerraformBackendRepositoryMapInput` via:
//
//	LocalTerraformBackendRepositoryMap{ "key": LocalTerraformBackendRepositoryArgs{...} }
type LocalTerraformBackendRepositoryMapInput interface {
	pulumi.Input

	ToLocalTerraformBackendRepositoryMapOutput() LocalTerraformBackendRepositoryMapOutput
	ToLocalTerraformBackendRepositoryMapOutputWithContext(context.Context) LocalTerraformBackendRepositoryMapOutput
}

type LocalTerraformBackendRepositoryMap map[string]LocalTerraformBackendRepositoryInput

func (LocalTerraformBackendRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalTerraformBackendRepository)(nil)).Elem()
}

func (i LocalTerraformBackendRepositoryMap) ToLocalTerraformBackendRepositoryMapOutput() LocalTerraformBackendRepositoryMapOutput {
	return i.ToLocalTerraformBackendRepositoryMapOutputWithContext(context.Background())
}

func (i LocalTerraformBackendRepositoryMap) ToLocalTerraformBackendRepositoryMapOutputWithContext(ctx context.Context) LocalTerraformBackendRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalTerraformBackendRepositoryMapOutput)
}

type LocalTerraformBackendRepositoryOutput struct{ *pulumi.OutputState }

func (LocalTerraformBackendRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalTerraformBackendRepository)(nil)).Elem()
}

func (o LocalTerraformBackendRepositoryOutput) ToLocalTerraformBackendRepositoryOutput() LocalTerraformBackendRepositoryOutput {
	return o
}

func (o LocalTerraformBackendRepositoryOutput) ToLocalTerraformBackendRepositoryOutputWithContext(ctx context.Context) LocalTerraformBackendRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalTerraformBackendRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalTerraformBackendRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalTerraformBackendRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalTerraformBackendRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalTerraformBackendRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o LocalTerraformBackendRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalTerraformBackendRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalTerraformBackendRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalTerraformBackendRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalTerraformBackendRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalTerraformBackendRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalTerraformBackendRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalTerraformBackendRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalTerraformBackendRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalTerraformBackendRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalTerraformBackendRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalTerraformBackendRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalTerraformBackendRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalTerraformBackendRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalTerraformBackendRepository)(nil)).Elem()
}

func (o LocalTerraformBackendRepositoryArrayOutput) ToLocalTerraformBackendRepositoryArrayOutput() LocalTerraformBackendRepositoryArrayOutput {
	return o
}

func (o LocalTerraformBackendRepositoryArrayOutput) ToLocalTerraformBackendRepositoryArrayOutputWithContext(ctx context.Context) LocalTerraformBackendRepositoryArrayOutput {
	return o
}

func (o LocalTerraformBackendRepositoryArrayOutput) Index(i pulumi.IntInput) LocalTerraformBackendRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalTerraformBackendRepository {
		return vs[0].([]*LocalTerraformBackendRepository)[vs[1].(int)]
	}).(LocalTerraformBackendRepositoryOutput)
}

type LocalTerraformBackendRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalTerraformBackendRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalTerraformBackendRepository)(nil)).Elem()
}

func (o LocalTerraformBackendRepositoryMapOutput) ToLocalTerraformBackendRepositoryMapOutput() LocalTerraformBackendRepositoryMapOutput {
	return o
}

func (o LocalTerraformBackendRepositoryMapOutput) ToLocalTerraformBackendRepositoryMapOutputWithContext(ctx context.Context) LocalTerraformBackendRepositoryMapOutput {
	return o
}

func (o LocalTerraformBackendRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalTerraformBackendRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalTerraformBackendRepository {
		return vs[0].(map[string]*LocalTerraformBackendRepository)[vs[1].(string)]
	}).(LocalTerraformBackendRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTerraformBackendRepositoryInput)(nil)).Elem(), &LocalTerraformBackendRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTerraformBackendRepositoryArrayInput)(nil)).Elem(), LocalTerraformBackendRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalTerraformBackendRepositoryMapInput)(nil)).Elem(), LocalTerraformBackendRepositoryMap{})
	pulumi.RegisterOutputType(LocalTerraformBackendRepositoryOutput{})
	pulumi.RegisterOutputType(LocalTerraformBackendRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalTerraformBackendRepositoryMapOutput{})
}
