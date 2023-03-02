// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a federated Conda repository.
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
//			_, err := artifactory.NewFederatedCondaRepository(ctx, "terraform-federated-test-conda-repo", &artifactory.FederatedCondaRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-conda-repo"),
//				Members: artifactory.FederatedCondaRepositoryMemberArray{
//					&artifactory.FederatedCondaRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-conda-repo"),
//					},
//					&artifactory.FederatedCondaRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-conda-repo-2"),
//					},
//				},
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
// Federated repositories can be imported using their name, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/federatedCondaRepository:FederatedCondaRepository terraform-federated-test-conda-repo terraform-federated-test-conda-repo
//
// ```
type FederatedCondaRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrOutput   `pulumi:"cdnRedirect"`
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
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedCondaRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                    `pulumi:"notes"`
	PackageType pulumi.StringOutput                       `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

// NewFederatedCondaRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedCondaRepository(ctx *pulumi.Context,
	name string, args *FederatedCondaRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedCondaRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedCondaRepository
	err := ctx.RegisterResource("artifactory:index/federatedCondaRepository:FederatedCondaRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedCondaRepository gets an existing FederatedCondaRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedCondaRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedCondaRepositoryState, opts ...pulumi.ResourceOption) (*FederatedCondaRepository, error) {
	var resource FederatedCondaRepository
	err := ctx.ReadResource("artifactory:index/federatedCondaRepository:FederatedCondaRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedCondaRepository resources.
type federatedCondaRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool   `pulumi:"cdnRedirect"`
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
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     []FederatedCondaRepositoryMember `pulumi:"members"`
	Notes       *string                          `pulumi:"notes"`
	PackageType *string                          `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

type FederatedCondaRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
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
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedCondaRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

func (FederatedCondaRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCondaRepositoryState)(nil)).Elem()
}

type federatedCondaRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool   `pulumi:"cdnRedirect"`
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
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedCondaRepositoryMember `pulumi:"members"`
	Notes   *string                          `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool `pulumi:"priorityResolution"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

// The set of arguments for constructing a FederatedCondaRepository resource.
type FederatedCondaRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
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
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedCondaRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
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

func (FederatedCondaRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCondaRepositoryArgs)(nil)).Elem()
}

type FederatedCondaRepositoryInput interface {
	pulumi.Input

	ToFederatedCondaRepositoryOutput() FederatedCondaRepositoryOutput
	ToFederatedCondaRepositoryOutputWithContext(ctx context.Context) FederatedCondaRepositoryOutput
}

func (*FederatedCondaRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCondaRepository)(nil)).Elem()
}

func (i *FederatedCondaRepository) ToFederatedCondaRepositoryOutput() FederatedCondaRepositoryOutput {
	return i.ToFederatedCondaRepositoryOutputWithContext(context.Background())
}

func (i *FederatedCondaRepository) ToFederatedCondaRepositoryOutputWithContext(ctx context.Context) FederatedCondaRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCondaRepositoryOutput)
}

// FederatedCondaRepositoryArrayInput is an input type that accepts FederatedCondaRepositoryArray and FederatedCondaRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedCondaRepositoryArrayInput` via:
//
//	FederatedCondaRepositoryArray{ FederatedCondaRepositoryArgs{...} }
type FederatedCondaRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedCondaRepositoryArrayOutput() FederatedCondaRepositoryArrayOutput
	ToFederatedCondaRepositoryArrayOutputWithContext(context.Context) FederatedCondaRepositoryArrayOutput
}

type FederatedCondaRepositoryArray []FederatedCondaRepositoryInput

func (FederatedCondaRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCondaRepository)(nil)).Elem()
}

func (i FederatedCondaRepositoryArray) ToFederatedCondaRepositoryArrayOutput() FederatedCondaRepositoryArrayOutput {
	return i.ToFederatedCondaRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedCondaRepositoryArray) ToFederatedCondaRepositoryArrayOutputWithContext(ctx context.Context) FederatedCondaRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCondaRepositoryArrayOutput)
}

// FederatedCondaRepositoryMapInput is an input type that accepts FederatedCondaRepositoryMap and FederatedCondaRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedCondaRepositoryMapInput` via:
//
//	FederatedCondaRepositoryMap{ "key": FederatedCondaRepositoryArgs{...} }
type FederatedCondaRepositoryMapInput interface {
	pulumi.Input

	ToFederatedCondaRepositoryMapOutput() FederatedCondaRepositoryMapOutput
	ToFederatedCondaRepositoryMapOutputWithContext(context.Context) FederatedCondaRepositoryMapOutput
}

type FederatedCondaRepositoryMap map[string]FederatedCondaRepositoryInput

func (FederatedCondaRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCondaRepository)(nil)).Elem()
}

func (i FederatedCondaRepositoryMap) ToFederatedCondaRepositoryMapOutput() FederatedCondaRepositoryMapOutput {
	return i.ToFederatedCondaRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedCondaRepositoryMap) ToFederatedCondaRepositoryMapOutputWithContext(ctx context.Context) FederatedCondaRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedCondaRepositoryMapOutput)
}

type FederatedCondaRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedCondaRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedCondaRepository)(nil)).Elem()
}

func (o FederatedCondaRepositoryOutput) ToFederatedCondaRepositoryOutput() FederatedCondaRepositoryOutput {
	return o
}

func (o FederatedCondaRepositoryOutput) ToFederatedCondaRepositoryOutputWithContext(ctx context.Context) FederatedCondaRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedCondaRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedCondaRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedCondaRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o FederatedCondaRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedCondaRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o FederatedCondaRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedCondaRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o FederatedCondaRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedCondaRepositoryOutput) Members() FederatedCondaRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) FederatedCondaRepositoryMemberArrayOutput { return v.Members }).(FederatedCondaRepositoryMemberArrayOutput)
}

func (o FederatedCondaRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedCondaRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedCondaRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o FederatedCondaRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedCondaRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedCondaRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o FederatedCondaRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedCondaRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedCondaRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedCondaRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedCondaRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedCondaRepository)(nil)).Elem()
}

func (o FederatedCondaRepositoryArrayOutput) ToFederatedCondaRepositoryArrayOutput() FederatedCondaRepositoryArrayOutput {
	return o
}

func (o FederatedCondaRepositoryArrayOutput) ToFederatedCondaRepositoryArrayOutputWithContext(ctx context.Context) FederatedCondaRepositoryArrayOutput {
	return o
}

func (o FederatedCondaRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedCondaRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedCondaRepository {
		return vs[0].([]*FederatedCondaRepository)[vs[1].(int)]
	}).(FederatedCondaRepositoryOutput)
}

type FederatedCondaRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedCondaRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedCondaRepository)(nil)).Elem()
}

func (o FederatedCondaRepositoryMapOutput) ToFederatedCondaRepositoryMapOutput() FederatedCondaRepositoryMapOutput {
	return o
}

func (o FederatedCondaRepositoryMapOutput) ToFederatedCondaRepositoryMapOutputWithContext(ctx context.Context) FederatedCondaRepositoryMapOutput {
	return o
}

func (o FederatedCondaRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedCondaRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedCondaRepository {
		return vs[0].(map[string]*FederatedCondaRepository)[vs[1].(string)]
	}).(FederatedCondaRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCondaRepositoryInput)(nil)).Elem(), &FederatedCondaRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCondaRepositoryArrayInput)(nil)).Elem(), FederatedCondaRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedCondaRepositoryMapInput)(nil)).Elem(), FederatedCondaRepositoryMap{})
	pulumi.RegisterOutputType(FederatedCondaRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedCondaRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedCondaRepositoryMapOutput{})
}
