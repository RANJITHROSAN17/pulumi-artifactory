// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a federated Rpm repository.
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
//			_, err := artifactory.NewFederatedRpmRepository(ctx, "terraform-federated-test-rpm-repo", &artifactory.FederatedRpmRepositoryArgs{
//				Key: pulumi.String("terraform-federated-test-rpm-repo"),
//				Members: artifactory.FederatedRpmRepositoryMemberArray{
//					&artifactory.FederatedRpmRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl.org/artifactory/terraform-federated-test-rpm-repo"),
//					},
//					&artifactory.FederatedRpmRepositoryMemberArgs{
//						Enabled: pulumi.Bool(true),
//						Url:     pulumi.String("http://tempurl2.org/artifactory/terraform-federated-test-rpm-repo-2"),
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
//	$ pulumi import artifactory:index/federatedRpmRepository:FederatedRpmRepository terraform-federated-test-rpm-repo terraform-federated-test-rpm-repo
//
// ```
type FederatedRpmRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut           pulumi.BoolPtrOutput   `pulumi:"blackedOut"`
	CalculateYumMetadata pulumi.BoolPtrOutput   `pulumi:"calculateYumMetadata"`
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect          pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	EnableFileListsIndexing pulumi.BoolPtrOutput `pulumi:"enableFileListsIndexing"`
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
	Members     FederatedRpmRepositoryMemberArrayOutput `pulumi:"members"`
	Notes       pulumi.StringPtrOutput                  `pulumi:"notes"`
	PackageType pulumi.StringOutput                     `pulumi:"packageType"`
	// Primary keypair used to sign artifacts.
	PrimaryKeypairRef pulumi.StringPtrOutput `pulumi:"primaryKeypairRef"`
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef pulumi.StringPtrOutput `pulumi:"secondaryKeypairRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
	// A comma separated list of XML file names containing RPM group component definitions. Artifactory includes the group
	// definitions as part of the calculated RPM metadata, as well as automatically generating a gzipped version of the group
	// files, if required.
	YumGroupFileNames pulumi.StringPtrOutput `pulumi:"yumGroupFileNames"`
	// The depth, relative to the repository's root folder, where RPM metadata is created. This is useful when your repository
	// contains multiple RPM repositories under parallel hierarchies. For example, if your RPMs are stored under
	// 'fedora/linux/$releasever/$basearch', specify a depth of 4.
	YumRootDepth pulumi.IntPtrOutput `pulumi:"yumRootDepth"`
}

// NewFederatedRpmRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedRpmRepository(ctx *pulumi.Context,
	name string, args *FederatedRpmRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedRpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	var resource FederatedRpmRepository
	err := ctx.RegisterResource("artifactory:index/federatedRpmRepository:FederatedRpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedRpmRepository gets an existing FederatedRpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedRpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedRpmRepositoryState, opts ...pulumi.ResourceOption) (*FederatedRpmRepository, error) {
	var resource FederatedRpmRepository
	err := ctx.ReadResource("artifactory:index/federatedRpmRepository:FederatedRpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedRpmRepository resources.
type federatedRpmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut           *bool   `pulumi:"blackedOut"`
	CalculateYumMetadata *bool   `pulumi:"calculateYumMetadata"`
	Description          *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect          *bool `pulumi:"downloadDirect"`
	EnableFileListsIndexing *bool `pulumi:"enableFileListsIndexing"`
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
	Members     []FederatedRpmRepositoryMember `pulumi:"members"`
	Notes       *string                        `pulumi:"notes"`
	PackageType *string                        `pulumi:"packageType"`
	// Primary keypair used to sign artifacts.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
	// A comma separated list of XML file names containing RPM group component definitions. Artifactory includes the group
	// definitions as part of the calculated RPM metadata, as well as automatically generating a gzipped version of the group
	// files, if required.
	YumGroupFileNames *string `pulumi:"yumGroupFileNames"`
	// The depth, relative to the repository's root folder, where RPM metadata is created. This is useful when your repository
	// contains multiple RPM repositories under parallel hierarchies. For example, if your RPMs are stored under
	// 'fedora/linux/$releasever/$basearch', specify a depth of 4.
	YumRootDepth *int `pulumi:"yumRootDepth"`
}

type FederatedRpmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut           pulumi.BoolPtrInput
	CalculateYumMetadata pulumi.BoolPtrInput
	Description          pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect          pulumi.BoolPtrInput
	EnableFileListsIndexing pulumi.BoolPtrInput
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
	Members     FederatedRpmRepositoryMemberArrayInput
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Primary keypair used to sign artifacts.
	PrimaryKeypairRef pulumi.StringPtrInput
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
	// A comma separated list of XML file names containing RPM group component definitions. Artifactory includes the group
	// definitions as part of the calculated RPM metadata, as well as automatically generating a gzipped version of the group
	// files, if required.
	YumGroupFileNames pulumi.StringPtrInput
	// The depth, relative to the repository's root folder, where RPM metadata is created. This is useful when your repository
	// contains multiple RPM repositories under parallel hierarchies. For example, if your RPMs are stored under
	// 'fedora/linux/$releasever/$basearch', specify a depth of 4.
	YumRootDepth pulumi.IntPtrInput
}

func (FederatedRpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedRpmRepositoryState)(nil)).Elem()
}

type federatedRpmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut           *bool   `pulumi:"blackedOut"`
	CalculateYumMetadata *bool   `pulumi:"calculateYumMetadata"`
	Description          *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect          *bool `pulumi:"downloadDirect"`
	EnableFileListsIndexing *bool `pulumi:"enableFileListsIndexing"`
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
	Members []FederatedRpmRepositoryMember `pulumi:"members"`
	Notes   *string                        `pulumi:"notes"`
	// Primary keypair used to sign artifacts.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
	// A comma separated list of XML file names containing RPM group component definitions. Artifactory includes the group
	// definitions as part of the calculated RPM metadata, as well as automatically generating a gzipped version of the group
	// files, if required.
	YumGroupFileNames *string `pulumi:"yumGroupFileNames"`
	// The depth, relative to the repository's root folder, where RPM metadata is created. This is useful when your repository
	// contains multiple RPM repositories under parallel hierarchies. For example, if your RPMs are stored under
	// 'fedora/linux/$releasever/$basearch', specify a depth of 4.
	YumRootDepth *int `pulumi:"yumRootDepth"`
}

// The set of arguments for constructing a FederatedRpmRepository resource.
type FederatedRpmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut           pulumi.BoolPtrInput
	CalculateYumMetadata pulumi.BoolPtrInput
	Description          pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect          pulumi.BoolPtrInput
	EnableFileListsIndexing pulumi.BoolPtrInput
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
	Members FederatedRpmRepositoryMemberArrayInput
	Notes   pulumi.StringPtrInput
	// Primary keypair used to sign artifacts.
	PrimaryKeypairRef pulumi.StringPtrInput
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
	// Secondary keypair used to sign artifacts.
	SecondaryKeypairRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
	// A comma separated list of XML file names containing RPM group component definitions. Artifactory includes the group
	// definitions as part of the calculated RPM metadata, as well as automatically generating a gzipped version of the group
	// files, if required.
	YumGroupFileNames pulumi.StringPtrInput
	// The depth, relative to the repository's root folder, where RPM metadata is created. This is useful when your repository
	// contains multiple RPM repositories under parallel hierarchies. For example, if your RPMs are stored under
	// 'fedora/linux/$releasever/$basearch', specify a depth of 4.
	YumRootDepth pulumi.IntPtrInput
}

func (FederatedRpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedRpmRepositoryArgs)(nil)).Elem()
}

type FederatedRpmRepositoryInput interface {
	pulumi.Input

	ToFederatedRpmRepositoryOutput() FederatedRpmRepositoryOutput
	ToFederatedRpmRepositoryOutputWithContext(ctx context.Context) FederatedRpmRepositoryOutput
}

func (*FederatedRpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedRpmRepository)(nil)).Elem()
}

func (i *FederatedRpmRepository) ToFederatedRpmRepositoryOutput() FederatedRpmRepositoryOutput {
	return i.ToFederatedRpmRepositoryOutputWithContext(context.Background())
}

func (i *FederatedRpmRepository) ToFederatedRpmRepositoryOutputWithContext(ctx context.Context) FederatedRpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedRpmRepositoryOutput)
}

// FederatedRpmRepositoryArrayInput is an input type that accepts FederatedRpmRepositoryArray and FederatedRpmRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedRpmRepositoryArrayInput` via:
//
//	FederatedRpmRepositoryArray{ FederatedRpmRepositoryArgs{...} }
type FederatedRpmRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedRpmRepositoryArrayOutput() FederatedRpmRepositoryArrayOutput
	ToFederatedRpmRepositoryArrayOutputWithContext(context.Context) FederatedRpmRepositoryArrayOutput
}

type FederatedRpmRepositoryArray []FederatedRpmRepositoryInput

func (FederatedRpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedRpmRepository)(nil)).Elem()
}

func (i FederatedRpmRepositoryArray) ToFederatedRpmRepositoryArrayOutput() FederatedRpmRepositoryArrayOutput {
	return i.ToFederatedRpmRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedRpmRepositoryArray) ToFederatedRpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedRpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedRpmRepositoryArrayOutput)
}

// FederatedRpmRepositoryMapInput is an input type that accepts FederatedRpmRepositoryMap and FederatedRpmRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedRpmRepositoryMapInput` via:
//
//	FederatedRpmRepositoryMap{ "key": FederatedRpmRepositoryArgs{...} }
type FederatedRpmRepositoryMapInput interface {
	pulumi.Input

	ToFederatedRpmRepositoryMapOutput() FederatedRpmRepositoryMapOutput
	ToFederatedRpmRepositoryMapOutputWithContext(context.Context) FederatedRpmRepositoryMapOutput
}

type FederatedRpmRepositoryMap map[string]FederatedRpmRepositoryInput

func (FederatedRpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedRpmRepository)(nil)).Elem()
}

func (i FederatedRpmRepositoryMap) ToFederatedRpmRepositoryMapOutput() FederatedRpmRepositoryMapOutput {
	return i.ToFederatedRpmRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedRpmRepositoryMap) ToFederatedRpmRepositoryMapOutputWithContext(ctx context.Context) FederatedRpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedRpmRepositoryMapOutput)
}

type FederatedRpmRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedRpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedRpmRepository)(nil)).Elem()
}

func (o FederatedRpmRepositoryOutput) ToFederatedRpmRepositoryOutput() FederatedRpmRepositoryOutput {
	return o
}

func (o FederatedRpmRepositoryOutput) ToFederatedRpmRepositoryOutputWithContext(ctx context.Context) FederatedRpmRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedRpmRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedRpmRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o FederatedRpmRepositoryOutput) CalculateYumMetadata() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.BoolPtrOutput { return v.CalculateYumMetadata }).(pulumi.BoolPtrOutput)
}

func (o FederatedRpmRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedRpmRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o FederatedRpmRepositoryOutput) EnableFileListsIndexing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.BoolPtrOutput { return v.EnableFileListsIndexing }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o FederatedRpmRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o FederatedRpmRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o FederatedRpmRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedRpmRepositoryOutput) Members() FederatedRpmRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) FederatedRpmRepositoryMemberArrayOutput { return v.Members }).(FederatedRpmRepositoryMemberArrayOutput)
}

func (o FederatedRpmRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedRpmRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Primary keypair used to sign artifacts.
func (o FederatedRpmRepositoryOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringPtrOutput { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedRpmRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o FederatedRpmRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedRpmRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedRpmRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o FederatedRpmRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Secondary keypair used to sign artifacts.
func (o FederatedRpmRepositoryOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringPtrOutput { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedRpmRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

// A comma separated list of XML file names containing RPM group component definitions. Artifactory includes the group
// definitions as part of the calculated RPM metadata, as well as automatically generating a gzipped version of the group
// files, if required.
func (o FederatedRpmRepositoryOutput) YumGroupFileNames() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.StringPtrOutput { return v.YumGroupFileNames }).(pulumi.StringPtrOutput)
}

// The depth, relative to the repository's root folder, where RPM metadata is created. This is useful when your repository
// contains multiple RPM repositories under parallel hierarchies. For example, if your RPMs are stored under
// 'fedora/linux/$releasever/$basearch', specify a depth of 4.
func (o FederatedRpmRepositoryOutput) YumRootDepth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FederatedRpmRepository) pulumi.IntPtrOutput { return v.YumRootDepth }).(pulumi.IntPtrOutput)
}

type FederatedRpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedRpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedRpmRepository)(nil)).Elem()
}

func (o FederatedRpmRepositoryArrayOutput) ToFederatedRpmRepositoryArrayOutput() FederatedRpmRepositoryArrayOutput {
	return o
}

func (o FederatedRpmRepositoryArrayOutput) ToFederatedRpmRepositoryArrayOutputWithContext(ctx context.Context) FederatedRpmRepositoryArrayOutput {
	return o
}

func (o FederatedRpmRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedRpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedRpmRepository {
		return vs[0].([]*FederatedRpmRepository)[vs[1].(int)]
	}).(FederatedRpmRepositoryOutput)
}

type FederatedRpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedRpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedRpmRepository)(nil)).Elem()
}

func (o FederatedRpmRepositoryMapOutput) ToFederatedRpmRepositoryMapOutput() FederatedRpmRepositoryMapOutput {
	return o
}

func (o FederatedRpmRepositoryMapOutput) ToFederatedRpmRepositoryMapOutputWithContext(ctx context.Context) FederatedRpmRepositoryMapOutput {
	return o
}

func (o FederatedRpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedRpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedRpmRepository {
		return vs[0].(map[string]*FederatedRpmRepository)[vs[1].(string)]
	}).(FederatedRpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedRpmRepositoryInput)(nil)).Elem(), &FederatedRpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedRpmRepositoryArrayInput)(nil)).Elem(), FederatedRpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedRpmRepositoryMapInput)(nil)).Elem(), FederatedRpmRepositoryMap{})
	pulumi.RegisterOutputType(FederatedRpmRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedRpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedRpmRepositoryMapOutput{})
}
