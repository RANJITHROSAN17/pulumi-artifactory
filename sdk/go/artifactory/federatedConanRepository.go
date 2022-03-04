// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Conan Repository Resource
//
// Creates a federated Conan repository
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewFederatedConanRepository(ctx, "terraform-federated-test-conan-repo", &artifactory.FederatedConanRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-conan-repo"),
// 			Members: FederatedConanRepositoryMemberArray{
// 				&FederatedConanRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-conan-repo"),
// 				},
// 				&FederatedConanRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-conan-repo-2"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type FederatedConanRepository struct {
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
	// - the identity key of the repo
	Key pulumi.StringOutput `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedConanRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                    `pulumi:"notes"`
	PackageType pulumi.StringOutput                       `pulumi:"packageType"`
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
	RepoLayoutRef pulumi.StringOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewFederatedConanRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedConanRepository(ctx *pulumi.Context,
	name string, args *FederatedConanRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedConanRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedConanRepository
	err := ctx.RegisterResource("artifactory:index/federatedConanRepository:FederatedConanRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedConanRepository gets an existing FederatedConanRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedConanRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedConanRepositoryState, opts ...pulumi.ResourceOption) (*FederatedConanRepository, error) {
	var resource FederatedConanRepository
	err := ctx.ReadResource("artifactory:index/federatedConanRepository:FederatedConanRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedConanRepository resources.
type federatedConanRepositoryState struct {
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
	// - the identity key of the repo
	Key *string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     []FederatedConanRepositoryMember `pulumi:"members"`
	Notes       *string                          `pulumi:"notes"`
	PackageType *string                          `pulumi:"packageType"`
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
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type FederatedConanRepositoryState struct {
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
	// - the identity key of the repo
	Key pulumi.StringPtrInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedConanRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
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
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedConanRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedConanRepositoryState)(nil)).Elem()
}

type federatedConanRepositoryArgs struct {
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
	// - the identity key of the repo
	Key string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members []FederatedConanRepositoryMember `pulumi:"members"`
	Notes   *string                          `pulumi:"notes"`
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
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedConanRepository resource.
type FederatedConanRepositoryArgs struct {
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
	// - the identity key of the repo
	Key pulumi.StringInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members FederatedConanRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
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
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedConanRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedConanRepositoryArgs)(nil)).Elem()
}

type FederatedConanRepositoryInput interface {
	pulumi.Input

	ToFederatedConanRepositoryOutput() FederatedConanRepositoryOutput
	ToFederatedConanRepositoryOutputWithContext(ctx context.Context) FederatedConanRepositoryOutput
}

func (*FederatedConanRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedConanRepository)(nil)).Elem()
}

func (i *FederatedConanRepository) ToFederatedConanRepositoryOutput() FederatedConanRepositoryOutput {
	return i.ToFederatedConanRepositoryOutputWithContext(context.Background())
}

func (i *FederatedConanRepository) ToFederatedConanRepositoryOutputWithContext(ctx context.Context) FederatedConanRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedConanRepositoryOutput)
}

// FederatedConanRepositoryArrayInput is an input type that accepts FederatedConanRepositoryArray and FederatedConanRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedConanRepositoryArrayInput` via:
//
//          FederatedConanRepositoryArray{ FederatedConanRepositoryArgs{...} }
type FederatedConanRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedConanRepositoryArrayOutput() FederatedConanRepositoryArrayOutput
	ToFederatedConanRepositoryArrayOutputWithContext(context.Context) FederatedConanRepositoryArrayOutput
}

type FederatedConanRepositoryArray []FederatedConanRepositoryInput

func (FederatedConanRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedConanRepository)(nil)).Elem()
}

func (i FederatedConanRepositoryArray) ToFederatedConanRepositoryArrayOutput() FederatedConanRepositoryArrayOutput {
	return i.ToFederatedConanRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedConanRepositoryArray) ToFederatedConanRepositoryArrayOutputWithContext(ctx context.Context) FederatedConanRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedConanRepositoryArrayOutput)
}

// FederatedConanRepositoryMapInput is an input type that accepts FederatedConanRepositoryMap and FederatedConanRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedConanRepositoryMapInput` via:
//
//          FederatedConanRepositoryMap{ "key": FederatedConanRepositoryArgs{...} }
type FederatedConanRepositoryMapInput interface {
	pulumi.Input

	ToFederatedConanRepositoryMapOutput() FederatedConanRepositoryMapOutput
	ToFederatedConanRepositoryMapOutputWithContext(context.Context) FederatedConanRepositoryMapOutput
}

type FederatedConanRepositoryMap map[string]FederatedConanRepositoryInput

func (FederatedConanRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedConanRepository)(nil)).Elem()
}

func (i FederatedConanRepositoryMap) ToFederatedConanRepositoryMapOutput() FederatedConanRepositoryMapOutput {
	return i.ToFederatedConanRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedConanRepositoryMap) ToFederatedConanRepositoryMapOutputWithContext(ctx context.Context) FederatedConanRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedConanRepositoryMapOutput)
}

type FederatedConanRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedConanRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedConanRepository)(nil)).Elem()
}

func (o FederatedConanRepositoryOutput) ToFederatedConanRepositoryOutput() FederatedConanRepositoryOutput {
	return o
}

func (o FederatedConanRepositoryOutput) ToFederatedConanRepositoryOutputWithContext(ctx context.Context) FederatedConanRepositoryOutput {
	return o
}

type FederatedConanRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedConanRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedConanRepository)(nil)).Elem()
}

func (o FederatedConanRepositoryArrayOutput) ToFederatedConanRepositoryArrayOutput() FederatedConanRepositoryArrayOutput {
	return o
}

func (o FederatedConanRepositoryArrayOutput) ToFederatedConanRepositoryArrayOutputWithContext(ctx context.Context) FederatedConanRepositoryArrayOutput {
	return o
}

func (o FederatedConanRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedConanRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedConanRepository {
		return vs[0].([]*FederatedConanRepository)[vs[1].(int)]
	}).(FederatedConanRepositoryOutput)
}

type FederatedConanRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedConanRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedConanRepository)(nil)).Elem()
}

func (o FederatedConanRepositoryMapOutput) ToFederatedConanRepositoryMapOutput() FederatedConanRepositoryMapOutput {
	return o
}

func (o FederatedConanRepositoryMapOutput) ToFederatedConanRepositoryMapOutputWithContext(ctx context.Context) FederatedConanRepositoryMapOutput {
	return o
}

func (o FederatedConanRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedConanRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedConanRepository {
		return vs[0].(map[string]*FederatedConanRepository)[vs[1].(string)]
	}).(FederatedConanRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedConanRepositoryInput)(nil)).Elem(), &FederatedConanRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedConanRepositoryArrayInput)(nil)).Elem(), FederatedConanRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedConanRepositoryMapInput)(nil)).Elem(), FederatedConanRepositoryMap{})
	pulumi.RegisterOutputType(FederatedConanRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedConanRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedConanRepositoryMapOutput{})
}
