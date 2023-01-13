// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a local RPM repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//	"io/ioutil"
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := ioutil.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewKeypair(ctx, "some-keypair-gpg-1", &artifactory.KeypairArgs{
//				PairName:   pulumi.String(fmt.Sprintf("some-keypair%v", random_id.Randid.Id)),
//				PairType:   pulumi.String("GPG"),
//				Alias:      pulumi.String("foo-alias1"),
//				PrivateKey: readFileOrPanic("samples/gpg.priv"),
//				PublicKey:  readFileOrPanic("samples/gpg.pub"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewKeypair(ctx, "some-keypair-gpg-2", &artifactory.KeypairArgs{
//				PairName:   pulumi.String(fmt.Sprintf("some-keypair%v", random_id.Randid.Id)),
//				PairType:   pulumi.String("GPG"),
//				Alias:      pulumi.String("foo-alias2"),
//				PrivateKey: readFileOrPanic("samples/gpg.priv"),
//				PublicKey:  readFileOrPanic("samples/gpg.pub"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewLocalRpmRepository(ctx, "terraform-local-test-rpm-repo-basic", &artifactory.LocalRpmRepositoryArgs{
//				Key:                     pulumi.String("terraform-local-test-rpm-repo-basic"),
//				YumRootDepth:            pulumi.Int(5),
//				CalculateYumMetadata:    pulumi.Bool(true),
//				EnableFileListsIndexing: pulumi.Bool(true),
//				YumGroupFileNames:       pulumi.String("file-1.xml,file-2.xml"),
//				PrimaryKeypairRef:       pulumi.Any(artifactory_keypair.SomeKeypairGPG1.Pair_name),
//				SecondaryKeypairRef:     pulumi.Any(artifactory_keypair.SomeKeypairGPG2.Pair_name),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				some_keypair_gpg_1,
//				some_keypair_gpg_2,
//			}))
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
//	$ pulumi import artifactory:index/localRpmRepository:LocalRpmRepository terraform-local-test-rpm-repo-basic terraform-local-test-rpm-repo-basic
//
// ```
type LocalRpmRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// Default: false.
	CalculateYumMetadata pulumi.BoolPtrOutput   `pulumi:"calculateYumMetadata"`
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// Default: false.
	EnableFileListsIndexing pulumi.BoolPtrOutput `pulumi:"enableFileListsIndexing"`
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
	// The primary GPG key to be used to sign packages.
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
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrOutput `pulumi:"secondaryKeypairRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
	// A comma separated list of XML file names containing RPM group component definitions.
	// Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
	// generating a gzipped version of the group files, if required. Default is empty string.
	YumGroupFileNames pulumi.StringPtrOutput `pulumi:"yumGroupFileNames"`
	// The depth, relative to the repository's root folder, where RPM metadata is created.
	// This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
	// your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots
	// exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
	// snapshots are not cleaned up.
	YumRootDepth pulumi.IntPtrOutput `pulumi:"yumRootDepth"`
}

// NewLocalRpmRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalRpmRepository(ctx *pulumi.Context,
	name string, args *LocalRpmRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalRpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalRpmRepository
	err := ctx.RegisterResource("artifactory:index/localRpmRepository:LocalRpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalRpmRepository gets an existing LocalRpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalRpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalRpmRepositoryState, opts ...pulumi.ResourceOption) (*LocalRpmRepository, error) {
	var resource LocalRpmRepository
	err := ctx.ReadResource("artifactory:index/localRpmRepository:LocalRpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalRpmRepository resources.
type localRpmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// Default: false.
	CalculateYumMetadata *bool   `pulumi:"calculateYumMetadata"`
	Description          *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// Default: false.
	EnableFileListsIndexing *bool `pulumi:"enableFileListsIndexing"`
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
	// The primary GPG key to be used to sign packages.
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
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
	// A comma separated list of XML file names containing RPM group component definitions.
	// Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
	// generating a gzipped version of the group files, if required. Default is empty string.
	YumGroupFileNames *string `pulumi:"yumGroupFileNames"`
	// The depth, relative to the repository's root folder, where RPM metadata is created.
	// This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
	// your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots
	// exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
	// snapshots are not cleaned up.
	YumRootDepth *int `pulumi:"yumRootDepth"`
}

type LocalRpmRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// Default: false.
	CalculateYumMetadata pulumi.BoolPtrInput
	Description          pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// Default: false.
	EnableFileListsIndexing pulumi.BoolPtrInput
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
	// The primary GPG key to be used to sign packages.
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
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
	// A comma separated list of XML file names containing RPM group component definitions.
	// Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
	// generating a gzipped version of the group files, if required. Default is empty string.
	YumGroupFileNames pulumi.StringPtrInput
	// The depth, relative to the repository's root folder, where RPM metadata is created.
	// This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
	// your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots
	// exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
	// snapshots are not cleaned up.
	YumRootDepth pulumi.IntPtrInput
}

func (LocalRpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localRpmRepositoryState)(nil)).Elem()
}

type localRpmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// Default: false.
	CalculateYumMetadata *bool   `pulumi:"calculateYumMetadata"`
	Description          *string `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// Default: false.
	EnableFileListsIndexing *bool `pulumi:"enableFileListsIndexing"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key   string  `pulumi:"key"`
	Notes *string `pulumi:"notes"`
	// The primary GPG key to be used to sign packages.
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
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
	// A comma separated list of XML file names containing RPM group component definitions.
	// Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
	// generating a gzipped version of the group files, if required. Default is empty string.
	YumGroupFileNames *string `pulumi:"yumGroupFileNames"`
	// The depth, relative to the repository's root folder, where RPM metadata is created.
	// This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
	// your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots
	// exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
	// snapshots are not cleaned up.
	YumRootDepth *int `pulumi:"yumRootDepth"`
}

// The set of arguments for constructing a LocalRpmRepository resource.
type LocalRpmRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// Default: false.
	CalculateYumMetadata pulumi.BoolPtrInput
	Description          pulumi.StringPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// Default: false.
	EnableFileListsIndexing pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key   pulumi.StringInput
	Notes pulumi.StringPtrInput
	// The primary GPG key to be used to sign packages.
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
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
	// A comma separated list of XML file names containing RPM group component definitions.
	// Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
	// generating a gzipped version of the group files, if required. Default is empty string.
	YumGroupFileNames pulumi.StringPtrInput
	// The depth, relative to the repository's root folder, where RPM metadata is created.
	// This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
	// your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots
	// exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
	// snapshots are not cleaned up.
	YumRootDepth pulumi.IntPtrInput
}

func (LocalRpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localRpmRepositoryArgs)(nil)).Elem()
}

type LocalRpmRepositoryInput interface {
	pulumi.Input

	ToLocalRpmRepositoryOutput() LocalRpmRepositoryOutput
	ToLocalRpmRepositoryOutputWithContext(ctx context.Context) LocalRpmRepositoryOutput
}

func (*LocalRpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalRpmRepository)(nil)).Elem()
}

func (i *LocalRpmRepository) ToLocalRpmRepositoryOutput() LocalRpmRepositoryOutput {
	return i.ToLocalRpmRepositoryOutputWithContext(context.Background())
}

func (i *LocalRpmRepository) ToLocalRpmRepositoryOutputWithContext(ctx context.Context) LocalRpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalRpmRepositoryOutput)
}

// LocalRpmRepositoryArrayInput is an input type that accepts LocalRpmRepositoryArray and LocalRpmRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalRpmRepositoryArrayInput` via:
//
//	LocalRpmRepositoryArray{ LocalRpmRepositoryArgs{...} }
type LocalRpmRepositoryArrayInput interface {
	pulumi.Input

	ToLocalRpmRepositoryArrayOutput() LocalRpmRepositoryArrayOutput
	ToLocalRpmRepositoryArrayOutputWithContext(context.Context) LocalRpmRepositoryArrayOutput
}

type LocalRpmRepositoryArray []LocalRpmRepositoryInput

func (LocalRpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalRpmRepository)(nil)).Elem()
}

func (i LocalRpmRepositoryArray) ToLocalRpmRepositoryArrayOutput() LocalRpmRepositoryArrayOutput {
	return i.ToLocalRpmRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalRpmRepositoryArray) ToLocalRpmRepositoryArrayOutputWithContext(ctx context.Context) LocalRpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalRpmRepositoryArrayOutput)
}

// LocalRpmRepositoryMapInput is an input type that accepts LocalRpmRepositoryMap and LocalRpmRepositoryMapOutput values.
// You can construct a concrete instance of `LocalRpmRepositoryMapInput` via:
//
//	LocalRpmRepositoryMap{ "key": LocalRpmRepositoryArgs{...} }
type LocalRpmRepositoryMapInput interface {
	pulumi.Input

	ToLocalRpmRepositoryMapOutput() LocalRpmRepositoryMapOutput
	ToLocalRpmRepositoryMapOutputWithContext(context.Context) LocalRpmRepositoryMapOutput
}

type LocalRpmRepositoryMap map[string]LocalRpmRepositoryInput

func (LocalRpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalRpmRepository)(nil)).Elem()
}

func (i LocalRpmRepositoryMap) ToLocalRpmRepositoryMapOutput() LocalRpmRepositoryMapOutput {
	return i.ToLocalRpmRepositoryMapOutputWithContext(context.Background())
}

func (i LocalRpmRepositoryMap) ToLocalRpmRepositoryMapOutputWithContext(ctx context.Context) LocalRpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalRpmRepositoryMapOutput)
}

type LocalRpmRepositoryOutput struct{ *pulumi.OutputState }

func (LocalRpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalRpmRepository)(nil)).Elem()
}

func (o LocalRpmRepositoryOutput) ToLocalRpmRepositoryOutput() LocalRpmRepositoryOutput {
	return o
}

func (o LocalRpmRepositoryOutput) ToLocalRpmRepositoryOutputWithContext(ctx context.Context) LocalRpmRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalRpmRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalRpmRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// Default: false.
func (o LocalRpmRepositoryOutput) CalculateYumMetadata() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.BoolPtrOutput { return v.CalculateYumMetadata }).(pulumi.BoolPtrOutput)
}

func (o LocalRpmRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalRpmRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// Default: false.
func (o LocalRpmRepositoryOutput) EnableFileListsIndexing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.BoolPtrOutput { return v.EnableFileListsIndexing }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*. By default no
// artifacts are excluded.
func (o LocalRpmRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalRpmRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalRpmRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

func (o LocalRpmRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalRpmRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// The primary GPG key to be used to sign packages.
func (o LocalRpmRepositoryOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringPtrOutput { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalRpmRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o LocalRpmRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalRpmRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalRpmRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalRpmRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The secondary GPG key to be used to sign packages.
func (o LocalRpmRepositoryOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringPtrOutput { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalRpmRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

// A comma separated list of XML file names containing RPM group component definitions.
// Artifactory includes the group definitions as part of the calculated RPM metadata, as well as automatically
// generating a gzipped version of the group files, if required. Default is empty string.
func (o LocalRpmRepositoryOutput) YumGroupFileNames() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.StringPtrOutput { return v.YumGroupFileNames }).(pulumi.StringPtrOutput)
}

// The depth, relative to the repository's root folder, where RPM metadata is created.
// This is useful when your repository contains multiple RPM repositories under parallel hierarchies. For example, if
// your RPMs are stored under 'fedora/linux/$releasever/$basearch', specify a depth of 4. Once the number of snapshots
// exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique
// snapshots are not cleaned up.
func (o LocalRpmRepositoryOutput) YumRootDepth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LocalRpmRepository) pulumi.IntPtrOutput { return v.YumRootDepth }).(pulumi.IntPtrOutput)
}

type LocalRpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalRpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalRpmRepository)(nil)).Elem()
}

func (o LocalRpmRepositoryArrayOutput) ToLocalRpmRepositoryArrayOutput() LocalRpmRepositoryArrayOutput {
	return o
}

func (o LocalRpmRepositoryArrayOutput) ToLocalRpmRepositoryArrayOutputWithContext(ctx context.Context) LocalRpmRepositoryArrayOutput {
	return o
}

func (o LocalRpmRepositoryArrayOutput) Index(i pulumi.IntInput) LocalRpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalRpmRepository {
		return vs[0].([]*LocalRpmRepository)[vs[1].(int)]
	}).(LocalRpmRepositoryOutput)
}

type LocalRpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalRpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalRpmRepository)(nil)).Elem()
}

func (o LocalRpmRepositoryMapOutput) ToLocalRpmRepositoryMapOutput() LocalRpmRepositoryMapOutput {
	return o
}

func (o LocalRpmRepositoryMapOutput) ToLocalRpmRepositoryMapOutputWithContext(ctx context.Context) LocalRpmRepositoryMapOutput {
	return o
}

func (o LocalRpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalRpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalRpmRepository {
		return vs[0].(map[string]*LocalRpmRepository)[vs[1].(string)]
	}).(LocalRpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalRpmRepositoryInput)(nil)).Elem(), &LocalRpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalRpmRepositoryArrayInput)(nil)).Elem(), LocalRpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalRpmRepositoryMapInput)(nil)).Elem(), LocalRpmRepositoryMap{})
	pulumi.RegisterOutputType(LocalRpmRepositoryOutput{})
	pulumi.RegisterOutputType(LocalRpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalRpmRepositoryMapOutput{})
}
