// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Federated Conda Repository Resource
//
// Creates a federated Conda repository
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
// 		_, err := artifactory.NewFederatedCondaRepository(ctx, "terraform-federated-test-conda-repo", &artifactory.FederatedCondaRepositoryArgs{
// 			Key: pulumi.String("terraform-federated-test-conda-repo"),
// 			Members: FederatedCondaRepositoryMemberArray{
// 				&FederatedCondaRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-conda-repo"),
// 				},
// 				&FederatedCondaRepositoryMemberArgs{
// 					Enable: true,
// 					Url:    pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-conda-repo-2"),
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
type FederatedCondaRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrOutput   `pulumi:"blackedOut"`
	Description            pulumi.StringPtrOutput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrOutput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringOutput    `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringOutput    `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key pulumi.StringOutput `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedCondaRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                    `pulumi:"notes"`
	PackageType pulumi.StringOutput                       `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
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
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key *string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     []FederatedCondaRepositoryMember `pulumi:"members"`
	Notes       *string                          `pulumi:"notes"`
	PackageType *string                          `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type FederatedCondaRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringPtrInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members     FederatedCondaRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (FederatedCondaRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedCondaRepositoryState)(nil)).Elem()
}

type federatedCondaRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key string `pulumi:"key"`
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members []FederatedCondaRepositoryMember `pulumi:"members"`
	Notes   *string                          `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution *bool    `pulumi:"priorityResolution"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedCondaRepository resource.
type FederatedCondaRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringInput
	// - The list of Federated members and must contain this repository URL (configured base URL + `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set. Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository) to set up Federated repositories correctly.
	Members FederatedCondaRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution pulumi.BoolPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
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
//          FederatedCondaRepositoryArray{ FederatedCondaRepositoryArgs{...} }
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
//          FederatedCondaRepositoryMap{ "key": FederatedCondaRepositoryArgs{...} }
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
