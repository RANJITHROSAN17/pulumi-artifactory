// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a federated Debian repository.
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
//			_, err := artifactory.NewFederatedDebianRepository(ctx, "terraform-federated-test-debian-repo", &artifactory.FederatedDebianRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-debian-repo"),
//				Members: FederatedDebianRepositoryMemberArray{
//					&FederatedDebianRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-debian-repo"),
//					},
//					&FederatedDebianRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-debian-repo-2"),
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
//	$ pulumi import artifactory:index/federatedDebianRepository:FederatedDebianRepository terraform-federated-test-debian-repo terraform-federated-test-debian-repo
//
// ```
type FederatedDebianRepository struct {
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
	IncludesPattern         pulumi.StringOutput      `pulumi:"includesPattern"`
	IndexCompressionFormats pulumi.StringArrayOutput `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedDebianRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                     `pulumi:"notes"`
	PackageType pulumi.StringOutput                        `pulumi:"packageType"`
	// Used to sign index files in Debian artifacts.
	PrimaryKeypairRef pulumi.StringPtrOutput `pulumi:"primaryKeypairRef"`
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
	// Used to sign index files in Debian artifacts.
	SecondaryKeypairRef pulumi.StringPtrOutput `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrOutput `pulumi:"trivialLayout"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewFederatedDebianRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedDebianRepository(ctx *pulumi.Context,
	name string, args *FederatedDebianRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedDebianRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedDebianRepository
	err := ctx.RegisterResource("artifactory:index/federatedDebianRepository:FederatedDebianRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedDebianRepository gets an existing FederatedDebianRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedDebianRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedDebianRepositoryState, opts ...pulumi.ResourceOption) (*FederatedDebianRepository, error) {
	var resource FederatedDebianRepository
	err := ctx.ReadResource("artifactory:index/federatedDebianRepository:FederatedDebianRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedDebianRepository resources.
type federatedDebianRepositoryState struct {
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
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     []FederatedDebianRepositoryMember `pulumi:"members"`
	Notes       *string                           `pulumi:"notes"`
	PackageType *string                           `pulumi:"packageType"`
	// Used to sign index files in Debian artifacts.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
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
	// Used to sign index files in Debian artifacts.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type FederatedDebianRepositoryState struct {
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
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members     FederatedDebianRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Used to sign index files in Debian artifacts.
	PrimaryKeypairRef pulumi.StringPtrInput
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
	// Used to sign index files in Debian artifacts.
	SecondaryKeypairRef pulumi.StringPtrInput
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedDebianRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDebianRepositoryState)(nil)).Elem()
}

type federatedDebianRepositoryArgs struct {
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
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedDebianRepositoryMember `pulumi:"members"`
	Notes   *string                           `pulumi:"notes"`
	// Used to sign index files in Debian artifacts.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
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
	// Used to sign index files in Debian artifacts.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout *bool `pulumi:"trivialLayout"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedDebianRepository resource.
type FederatedDebianRepositoryArgs struct {
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
	IncludesPattern         pulumi.StringPtrInput
	IndexCompressionFormats pulumi.StringArrayInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedDebianRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Used to sign index files in Debian artifacts.
	PrimaryKeypairRef pulumi.StringPtrInput
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
	// Used to sign index files in Debian artifacts.
	SecondaryKeypairRef pulumi.StringPtrInput
	// When set, the repository will use the deprecated trivial layout.
	//
	// Deprecated: You shouldn't be using this
	TrivialLayout pulumi.BoolPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedDebianRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedDebianRepositoryArgs)(nil)).Elem()
}

type FederatedDebianRepositoryInput interface {
	pulumi.Input

	ToFederatedDebianRepositoryOutput() FederatedDebianRepositoryOutput
	ToFederatedDebianRepositoryOutputWithContext(ctx context.Context) FederatedDebianRepositoryOutput
}

func (*FederatedDebianRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDebianRepository)(nil)).Elem()
}

func (i *FederatedDebianRepository) ToFederatedDebianRepositoryOutput() FederatedDebianRepositoryOutput {
	return i.ToFederatedDebianRepositoryOutputWithContext(context.Background())
}

func (i *FederatedDebianRepository) ToFederatedDebianRepositoryOutputWithContext(ctx context.Context) FederatedDebianRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDebianRepositoryOutput)
}

// FederatedDebianRepositoryArrayInput is an input type that accepts FederatedDebianRepositoryArray and FederatedDebianRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedDebianRepositoryArrayInput` via:
//
//	FederatedDebianRepositoryArray{ FederatedDebianRepositoryArgs{...} }
type FederatedDebianRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedDebianRepositoryArrayOutput() FederatedDebianRepositoryArrayOutput
	ToFederatedDebianRepositoryArrayOutputWithContext(context.Context) FederatedDebianRepositoryArrayOutput
}

type FederatedDebianRepositoryArray []FederatedDebianRepositoryInput

func (FederatedDebianRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDebianRepository)(nil)).Elem()
}

func (i FederatedDebianRepositoryArray) ToFederatedDebianRepositoryArrayOutput() FederatedDebianRepositoryArrayOutput {
	return i.ToFederatedDebianRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedDebianRepositoryArray) ToFederatedDebianRepositoryArrayOutputWithContext(ctx context.Context) FederatedDebianRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDebianRepositoryArrayOutput)
}

// FederatedDebianRepositoryMapInput is an input type that accepts FederatedDebianRepositoryMap and FederatedDebianRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedDebianRepositoryMapInput` via:
//
//	FederatedDebianRepositoryMap{ "key": FederatedDebianRepositoryArgs{...} }
type FederatedDebianRepositoryMapInput interface {
	pulumi.Input

	ToFederatedDebianRepositoryMapOutput() FederatedDebianRepositoryMapOutput
	ToFederatedDebianRepositoryMapOutputWithContext(context.Context) FederatedDebianRepositoryMapOutput
}

type FederatedDebianRepositoryMap map[string]FederatedDebianRepositoryInput

func (FederatedDebianRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDebianRepository)(nil)).Elem()
}

func (i FederatedDebianRepositoryMap) ToFederatedDebianRepositoryMapOutput() FederatedDebianRepositoryMapOutput {
	return i.ToFederatedDebianRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedDebianRepositoryMap) ToFederatedDebianRepositoryMapOutputWithContext(ctx context.Context) FederatedDebianRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedDebianRepositoryMapOutput)
}

type FederatedDebianRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedDebianRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedDebianRepository)(nil)).Elem()
}

func (o FederatedDebianRepositoryOutput) ToFederatedDebianRepositoryOutput() FederatedDebianRepositoryOutput {
	return o
}

func (o FederatedDebianRepositoryOutput) ToFederatedDebianRepositoryOutputWithContext(ctx context.Context) FederatedDebianRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedDebianRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedDebianRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o FederatedDebianRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedDebianRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o FederatedDebianRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedDebianRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o FederatedDebianRepositoryOutput) IndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringArrayOutput { return v.IndexCompressionFormats }).(pulumi.StringArrayOutput)
}

// the identity key of the repo.
func (o FederatedDebianRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedDebianRepositoryOutput) Members() FederatedDebianRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) FederatedDebianRepositoryMemberArrayOutput { return v.Members }).(FederatedDebianRepositoryMemberArrayOutput)
}

func (o FederatedDebianRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedDebianRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Used to sign index files in Debian artifacts.
func (o FederatedDebianRepositoryOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedDebianRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
func (o FederatedDebianRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. When assigning repository to a project, repository key must be prefixed
// with project key, separated by a dash.
func (o FederatedDebianRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedDebianRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o FederatedDebianRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Used to sign index files in Debian artifacts.
func (o FederatedDebianRepositoryOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.StringPtrOutput { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

// When set, the repository will use the deprecated trivial layout.
//
// Deprecated: You shouldn't be using this
func (o FederatedDebianRepositoryOutput) TrivialLayout() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.TrivialLayout }).(pulumi.BoolPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedDebianRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedDebianRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedDebianRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedDebianRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedDebianRepository)(nil)).Elem()
}

func (o FederatedDebianRepositoryArrayOutput) ToFederatedDebianRepositoryArrayOutput() FederatedDebianRepositoryArrayOutput {
	return o
}

func (o FederatedDebianRepositoryArrayOutput) ToFederatedDebianRepositoryArrayOutputWithContext(ctx context.Context) FederatedDebianRepositoryArrayOutput {
	return o
}

func (o FederatedDebianRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedDebianRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedDebianRepository {
		return vs[0].([]*FederatedDebianRepository)[vs[1].(int)]
	}).(FederatedDebianRepositoryOutput)
}

type FederatedDebianRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedDebianRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedDebianRepository)(nil)).Elem()
}

func (o FederatedDebianRepositoryMapOutput) ToFederatedDebianRepositoryMapOutput() FederatedDebianRepositoryMapOutput {
	return o
}

func (o FederatedDebianRepositoryMapOutput) ToFederatedDebianRepositoryMapOutputWithContext(ctx context.Context) FederatedDebianRepositoryMapOutput {
	return o
}

func (o FederatedDebianRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedDebianRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedDebianRepository {
		return vs[0].(map[string]*FederatedDebianRepository)[vs[1].(string)]
	}).(FederatedDebianRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDebianRepositoryInput)(nil)).Elem(), &FederatedDebianRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDebianRepositoryArrayInput)(nil)).Elem(), FederatedDebianRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedDebianRepositoryMapInput)(nil)).Elem(), FederatedDebianRepositoryMap{})
	pulumi.RegisterOutputType(FederatedDebianRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedDebianRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedDebianRepositoryMapOutput{})
}
