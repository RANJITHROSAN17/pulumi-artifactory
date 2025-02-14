// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a local Generic repository.
//
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
//			_, err := artifactory.NewLocalGenericRepository(ctx, "terraform-local-test-generic-repo", &artifactory.LocalGenericRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-generic-repo"),
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
//	$ pulumi import artifactory:index/localGenericRepository:LocalGenericRepository terraform-local-test-generic-repo terraform-local-test-generic-repo
//
// ```
type LocalGenericRepository struct {
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
	// A mandatory identifier for the repository that must be unique.
	// It cannot begin with a number or contain spaces or special characters.
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

// NewLocalGenericRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalGenericRepository(ctx *pulumi.Context,
	name string, args *LocalGenericRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalGenericRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalGenericRepository
	err := ctx.RegisterResource("artifactory:index/localGenericRepository:LocalGenericRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalGenericRepository gets an existing LocalGenericRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalGenericRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalGenericRepositoryState, opts ...pulumi.ResourceOption) (*LocalGenericRepository, error) {
	var resource LocalGenericRepository
	err := ctx.ReadResource("artifactory:index/localGenericRepository:LocalGenericRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalGenericRepository resources.
type localGenericRepositoryState struct {
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
	// A mandatory identifier for the repository that must be unique.
	// It cannot begin with a number or contain spaces or special characters.
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

type LocalGenericRepositoryState struct {
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
	// A mandatory identifier for the repository that must be unique.
	// It cannot begin with a number or contain spaces or special characters.
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

func (LocalGenericRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localGenericRepositoryState)(nil)).Elem()
}

type localGenericRepositoryArgs struct {
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
	// A mandatory identifier for the repository that must be unique.
	// It cannot begin with a number or contain spaces or special characters.
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

// The set of arguments for constructing a LocalGenericRepository resource.
type LocalGenericRepositoryArgs struct {
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
	// A mandatory identifier for the repository that must be unique.
	// It cannot begin with a number or contain spaces or special characters.
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

func (LocalGenericRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localGenericRepositoryArgs)(nil)).Elem()
}

type LocalGenericRepositoryInput interface {
	pulumi.Input

	ToLocalGenericRepositoryOutput() LocalGenericRepositoryOutput
	ToLocalGenericRepositoryOutputWithContext(ctx context.Context) LocalGenericRepositoryOutput
}

func (*LocalGenericRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGenericRepository)(nil)).Elem()
}

func (i *LocalGenericRepository) ToLocalGenericRepositoryOutput() LocalGenericRepositoryOutput {
	return i.ToLocalGenericRepositoryOutputWithContext(context.Background())
}

func (i *LocalGenericRepository) ToLocalGenericRepositoryOutputWithContext(ctx context.Context) LocalGenericRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGenericRepositoryOutput)
}

// LocalGenericRepositoryArrayInput is an input type that accepts LocalGenericRepositoryArray and LocalGenericRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalGenericRepositoryArrayInput` via:
//
//	LocalGenericRepositoryArray{ LocalGenericRepositoryArgs{...} }
type LocalGenericRepositoryArrayInput interface {
	pulumi.Input

	ToLocalGenericRepositoryArrayOutput() LocalGenericRepositoryArrayOutput
	ToLocalGenericRepositoryArrayOutputWithContext(context.Context) LocalGenericRepositoryArrayOutput
}

type LocalGenericRepositoryArray []LocalGenericRepositoryInput

func (LocalGenericRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalGenericRepository)(nil)).Elem()
}

func (i LocalGenericRepositoryArray) ToLocalGenericRepositoryArrayOutput() LocalGenericRepositoryArrayOutput {
	return i.ToLocalGenericRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalGenericRepositoryArray) ToLocalGenericRepositoryArrayOutputWithContext(ctx context.Context) LocalGenericRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGenericRepositoryArrayOutput)
}

// LocalGenericRepositoryMapInput is an input type that accepts LocalGenericRepositoryMap and LocalGenericRepositoryMapOutput values.
// You can construct a concrete instance of `LocalGenericRepositoryMapInput` via:
//
//	LocalGenericRepositoryMap{ "key": LocalGenericRepositoryArgs{...} }
type LocalGenericRepositoryMapInput interface {
	pulumi.Input

	ToLocalGenericRepositoryMapOutput() LocalGenericRepositoryMapOutput
	ToLocalGenericRepositoryMapOutputWithContext(context.Context) LocalGenericRepositoryMapOutput
}

type LocalGenericRepositoryMap map[string]LocalGenericRepositoryInput

func (LocalGenericRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalGenericRepository)(nil)).Elem()
}

func (i LocalGenericRepositoryMap) ToLocalGenericRepositoryMapOutput() LocalGenericRepositoryMapOutput {
	return i.ToLocalGenericRepositoryMapOutputWithContext(context.Background())
}

func (i LocalGenericRepositoryMap) ToLocalGenericRepositoryMapOutputWithContext(ctx context.Context) LocalGenericRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGenericRepositoryMapOutput)
}

type LocalGenericRepositoryOutput struct{ *pulumi.OutputState }

func (LocalGenericRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGenericRepository)(nil)).Elem()
}

func (o LocalGenericRepositoryOutput) ToLocalGenericRepositoryOutput() LocalGenericRepositoryOutput {
	return o
}

func (o LocalGenericRepositoryOutput) ToLocalGenericRepositoryOutputWithContext(ctx context.Context) LocalGenericRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalGenericRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalGenericRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalGenericRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalGenericRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalGenericRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o LocalGenericRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalGenericRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// A mandatory identifier for the repository that must be unique.
// It cannot begin with a number or contain spaces or special characters.
func (o LocalGenericRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalGenericRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalGenericRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalGenericRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV", "PROD", or one of custom environment. Before
// Artifactory 7.53.1, up to 2 values ("DEV" and "PROD") are allowed. From 7.53.1 onward, only one value is allowed. The
// attribute should only be used if the repository is already assigned to the existing project. If not, the attribute will
// be ignored by Artifactory, but will remain in the Terraform state, which will create state drift during the update.
func (o LocalGenericRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalGenericRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalGenericRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalGenericRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalGenericRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalGenericRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalGenericRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalGenericRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalGenericRepository)(nil)).Elem()
}

func (o LocalGenericRepositoryArrayOutput) ToLocalGenericRepositoryArrayOutput() LocalGenericRepositoryArrayOutput {
	return o
}

func (o LocalGenericRepositoryArrayOutput) ToLocalGenericRepositoryArrayOutputWithContext(ctx context.Context) LocalGenericRepositoryArrayOutput {
	return o
}

func (o LocalGenericRepositoryArrayOutput) Index(i pulumi.IntInput) LocalGenericRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalGenericRepository {
		return vs[0].([]*LocalGenericRepository)[vs[1].(int)]
	}).(LocalGenericRepositoryOutput)
}

type LocalGenericRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalGenericRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalGenericRepository)(nil)).Elem()
}

func (o LocalGenericRepositoryMapOutput) ToLocalGenericRepositoryMapOutput() LocalGenericRepositoryMapOutput {
	return o
}

func (o LocalGenericRepositoryMapOutput) ToLocalGenericRepositoryMapOutputWithContext(ctx context.Context) LocalGenericRepositoryMapOutput {
	return o
}

func (o LocalGenericRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalGenericRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalGenericRepository {
		return vs[0].(map[string]*LocalGenericRepository)[vs[1].(string)]
	}).(LocalGenericRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGenericRepositoryInput)(nil)).Elem(), &LocalGenericRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGenericRepositoryArrayInput)(nil)).Elem(), LocalGenericRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGenericRepositoryMapInput)(nil)).Elem(), LocalGenericRepositoryMap{})
	pulumi.RegisterOutputType(LocalGenericRepositoryOutput{})
	pulumi.RegisterOutputType(LocalGenericRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalGenericRepositoryMapOutput{})
}
