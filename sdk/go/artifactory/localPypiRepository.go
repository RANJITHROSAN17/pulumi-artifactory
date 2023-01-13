// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a local Pypi repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewLocalPypiRepository(ctx, "terraform-local-test-pypi-repo", &artifactory.LocalPypiRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-pypi-repo"),
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
//	$ pulumi import artifactory:index/localPypiRepository:LocalPypiRepository terraform-local-test-pypi-repo terraform-local-test-pypi-repo
//
// ```
type LocalPypiRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  pulumi.BoolPtrOutput   `pulumi:"blackedOut"`
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
	Key         pulumi.StringOutput    `pulumi:"key"`
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

// NewLocalPypiRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalPypiRepository(ctx *pulumi.Context,
	name string, args *LocalPypiRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalPypiRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalPypiRepository
	err := ctx.RegisterResource("artifactory:index/localPypiRepository:LocalPypiRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalPypiRepository gets an existing LocalPypiRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalPypiRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalPypiRepositoryState, opts ...pulumi.ResourceOption) (*LocalPypiRepository, error) {
	var resource LocalPypiRepository
	err := ctx.ReadResource("artifactory:index/localPypiRepository:LocalPypiRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalPypiRepository resources.
type localPypiRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  *bool   `pulumi:"blackedOut"`
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
	Key         *string `pulumi:"key"`
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

type LocalPypiRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  pulumi.BoolPtrInput
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
	Key         pulumi.StringPtrInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

func (LocalPypiRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localPypiRepositoryState)(nil)).Elem()
}

type localPypiRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  *bool   `pulumi:"blackedOut"`
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
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

// The set of arguments for constructing a LocalPypiRepository resource.
type LocalPypiRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut  pulumi.BoolPtrInput
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
	Key   pulumi.StringInput
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
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

func (LocalPypiRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localPypiRepositoryArgs)(nil)).Elem()
}

type LocalPypiRepositoryInput interface {
	pulumi.Input

	ToLocalPypiRepositoryOutput() LocalPypiRepositoryOutput
	ToLocalPypiRepositoryOutputWithContext(ctx context.Context) LocalPypiRepositoryOutput
}

func (*LocalPypiRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalPypiRepository)(nil)).Elem()
}

func (i *LocalPypiRepository) ToLocalPypiRepositoryOutput() LocalPypiRepositoryOutput {
	return i.ToLocalPypiRepositoryOutputWithContext(context.Background())
}

func (i *LocalPypiRepository) ToLocalPypiRepositoryOutputWithContext(ctx context.Context) LocalPypiRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPypiRepositoryOutput)
}

// LocalPypiRepositoryArrayInput is an input type that accepts LocalPypiRepositoryArray and LocalPypiRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalPypiRepositoryArrayInput` via:
//
//	LocalPypiRepositoryArray{ LocalPypiRepositoryArgs{...} }
type LocalPypiRepositoryArrayInput interface {
	pulumi.Input

	ToLocalPypiRepositoryArrayOutput() LocalPypiRepositoryArrayOutput
	ToLocalPypiRepositoryArrayOutputWithContext(context.Context) LocalPypiRepositoryArrayOutput
}

type LocalPypiRepositoryArray []LocalPypiRepositoryInput

func (LocalPypiRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalPypiRepository)(nil)).Elem()
}

func (i LocalPypiRepositoryArray) ToLocalPypiRepositoryArrayOutput() LocalPypiRepositoryArrayOutput {
	return i.ToLocalPypiRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalPypiRepositoryArray) ToLocalPypiRepositoryArrayOutputWithContext(ctx context.Context) LocalPypiRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPypiRepositoryArrayOutput)
}

// LocalPypiRepositoryMapInput is an input type that accepts LocalPypiRepositoryMap and LocalPypiRepositoryMapOutput values.
// You can construct a concrete instance of `LocalPypiRepositoryMapInput` via:
//
//	LocalPypiRepositoryMap{ "key": LocalPypiRepositoryArgs{...} }
type LocalPypiRepositoryMapInput interface {
	pulumi.Input

	ToLocalPypiRepositoryMapOutput() LocalPypiRepositoryMapOutput
	ToLocalPypiRepositoryMapOutputWithContext(context.Context) LocalPypiRepositoryMapOutput
}

type LocalPypiRepositoryMap map[string]LocalPypiRepositoryInput

func (LocalPypiRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalPypiRepository)(nil)).Elem()
}

func (i LocalPypiRepositoryMap) ToLocalPypiRepositoryMapOutput() LocalPypiRepositoryMapOutput {
	return i.ToLocalPypiRepositoryMapOutputWithContext(context.Background())
}

func (i LocalPypiRepositoryMap) ToLocalPypiRepositoryMapOutputWithContext(ctx context.Context) LocalPypiRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalPypiRepositoryMapOutput)
}

type LocalPypiRepositoryOutput struct{ *pulumi.OutputState }

func (LocalPypiRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalPypiRepository)(nil)).Elem()
}

func (o LocalPypiRepositoryOutput) ToLocalPypiRepositoryOutput() LocalPypiRepositoryOutput {
	return o
}

func (o LocalPypiRepositoryOutput) ToLocalPypiRepositoryOutputWithContext(ctx context.Context) LocalPypiRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalPypiRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalPypiRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LocalPypiRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalPypiRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o LocalPypiRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalPypiRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalPypiRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o LocalPypiRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalPypiRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalPypiRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o LocalPypiRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalPypiRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalPypiRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalPypiRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalPypiRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalPypiRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalPypiRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalPypiRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalPypiRepository)(nil)).Elem()
}

func (o LocalPypiRepositoryArrayOutput) ToLocalPypiRepositoryArrayOutput() LocalPypiRepositoryArrayOutput {
	return o
}

func (o LocalPypiRepositoryArrayOutput) ToLocalPypiRepositoryArrayOutputWithContext(ctx context.Context) LocalPypiRepositoryArrayOutput {
	return o
}

func (o LocalPypiRepositoryArrayOutput) Index(i pulumi.IntInput) LocalPypiRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalPypiRepository {
		return vs[0].([]*LocalPypiRepository)[vs[1].(int)]
	}).(LocalPypiRepositoryOutput)
}

type LocalPypiRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalPypiRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalPypiRepository)(nil)).Elem()
}

func (o LocalPypiRepositoryMapOutput) ToLocalPypiRepositoryMapOutput() LocalPypiRepositoryMapOutput {
	return o
}

func (o LocalPypiRepositoryMapOutput) ToLocalPypiRepositoryMapOutputWithContext(ctx context.Context) LocalPypiRepositoryMapOutput {
	return o
}

func (o LocalPypiRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalPypiRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalPypiRepository {
		return vs[0].(map[string]*LocalPypiRepository)[vs[1].(string)]
	}).(LocalPypiRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPypiRepositoryInput)(nil)).Elem(), &LocalPypiRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPypiRepositoryArrayInput)(nil)).Elem(), LocalPypiRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalPypiRepositoryMapInput)(nil)).Elem(), LocalPypiRepositoryMap{})
	pulumi.RegisterOutputType(LocalPypiRepositoryOutput{})
	pulumi.RegisterOutputType(LocalPypiRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalPypiRepositoryMapOutput{})
}
