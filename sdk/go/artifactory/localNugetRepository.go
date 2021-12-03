// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Nuget Repository Resource
//
// Creates a local Nuget repository and allows for the creation of a
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
// 		_, err := artifactory.NewLocalNugetRepository(ctx, "terraform_local_test_nuget_repo_basic", &artifactory.LocalNugetRepositoryArgs{
// 			ForceNugetAuthentication: pulumi.Bool(true),
// 			Key:                      pulumi.String("terraform-local-test-nuget-repo-basic"),
// 			MaxUniqueSnapshots:       pulumi.Int(5),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalNugetRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrOutput   `pulumi:"blackedOut"`
	Description            pulumi.StringPtrOutput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrOutput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringOutput    `pulumi:"excludesPattern"`
	// - Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication pulumi.BoolPtrOutput `pulumi:"forceNugetAuthentication"`
	IncludesPattern          pulumi.StringOutput  `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key pulumi.StringOutput `pulumi:"key"`
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrOutput      `pulumi:"maxUniqueSnapshots"`
	Notes              pulumi.StringPtrOutput   `pulumi:"notes"`
	PackageType        pulumi.StringOutput      `pulumi:"packageType"`
	PropertySets       pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef      pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex          pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewLocalNugetRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalNugetRepository(ctx *pulumi.Context,
	name string, args *LocalNugetRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalNugetRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalNugetRepository
	err := ctx.RegisterResource("artifactory:index/localNugetRepository:LocalNugetRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalNugetRepository gets an existing LocalNugetRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalNugetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalNugetRepositoryState, opts ...pulumi.ResourceOption) (*LocalNugetRepository, error) {
	var resource LocalNugetRepository
	err := ctx.ReadResource("artifactory:index/localNugetRepository:LocalNugetRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalNugetRepository resources.
type localNugetRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// - Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication *bool   `pulumi:"forceNugetAuthentication"`
	IncludesPattern          *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key *string `pulumi:"key"`
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int     `pulumi:"maxUniqueSnapshots"`
	Notes              *string  `pulumi:"notes"`
	PackageType        *string  `pulumi:"packageType"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

type LocalNugetRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	// - Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication pulumi.BoolPtrInput
	IncludesPattern          pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringPtrInput
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	Notes              pulumi.StringPtrInput
	PackageType        pulumi.StringPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (LocalNugetRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localNugetRepositoryState)(nil)).Elem()
}

type localNugetRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// - Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication *bool   `pulumi:"forceNugetAuthentication"`
	IncludesPattern          *string `pulumi:"includesPattern"`
	// - the identity key of the repo
	Key string `pulumi:"key"`
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots *int     `pulumi:"maxUniqueSnapshots"`
	Notes              *string  `pulumi:"notes"`
	PropertySets       []string `pulumi:"propertySets"`
	RepoLayoutRef      *string  `pulumi:"repoLayoutRef"`
	XrayIndex          *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalNugetRepository resource.
type LocalNugetRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	BlackedOut             pulumi.BoolPtrInput
	Description            pulumi.StringPtrInput
	DownloadDirect         pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	// - Force basic authentication credentials in order to use this repository.
	ForceNugetAuthentication pulumi.BoolPtrInput
	IncludesPattern          pulumi.StringPtrInput
	// - the identity key of the repo
	Key pulumi.StringInput
	// - The maximum number of unique snapshots of a single artifact to store.
	//   Once the number of snapshots exceeds this setting, older versions are removed.
	//   A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots pulumi.IntPtrInput
	Notes              pulumi.StringPtrInput
	PropertySets       pulumi.StringArrayInput
	RepoLayoutRef      pulumi.StringPtrInput
	XrayIndex          pulumi.BoolPtrInput
}

func (LocalNugetRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localNugetRepositoryArgs)(nil)).Elem()
}

type LocalNugetRepositoryInput interface {
	pulumi.Input

	ToLocalNugetRepositoryOutput() LocalNugetRepositoryOutput
	ToLocalNugetRepositoryOutputWithContext(ctx context.Context) LocalNugetRepositoryOutput
}

func (*LocalNugetRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNugetRepository)(nil))
}

func (i *LocalNugetRepository) ToLocalNugetRepositoryOutput() LocalNugetRepositoryOutput {
	return i.ToLocalNugetRepositoryOutputWithContext(context.Background())
}

func (i *LocalNugetRepository) ToLocalNugetRepositoryOutputWithContext(ctx context.Context) LocalNugetRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNugetRepositoryOutput)
}

func (i *LocalNugetRepository) ToLocalNugetRepositoryPtrOutput() LocalNugetRepositoryPtrOutput {
	return i.ToLocalNugetRepositoryPtrOutputWithContext(context.Background())
}

func (i *LocalNugetRepository) ToLocalNugetRepositoryPtrOutputWithContext(ctx context.Context) LocalNugetRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNugetRepositoryPtrOutput)
}

type LocalNugetRepositoryPtrInput interface {
	pulumi.Input

	ToLocalNugetRepositoryPtrOutput() LocalNugetRepositoryPtrOutput
	ToLocalNugetRepositoryPtrOutputWithContext(ctx context.Context) LocalNugetRepositoryPtrOutput
}

type localNugetRepositoryPtrType LocalNugetRepositoryArgs

func (*localNugetRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNugetRepository)(nil))
}

func (i *localNugetRepositoryPtrType) ToLocalNugetRepositoryPtrOutput() LocalNugetRepositoryPtrOutput {
	return i.ToLocalNugetRepositoryPtrOutputWithContext(context.Background())
}

func (i *localNugetRepositoryPtrType) ToLocalNugetRepositoryPtrOutputWithContext(ctx context.Context) LocalNugetRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNugetRepositoryPtrOutput)
}

// LocalNugetRepositoryArrayInput is an input type that accepts LocalNugetRepositoryArray and LocalNugetRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalNugetRepositoryArrayInput` via:
//
//          LocalNugetRepositoryArray{ LocalNugetRepositoryArgs{...} }
type LocalNugetRepositoryArrayInput interface {
	pulumi.Input

	ToLocalNugetRepositoryArrayOutput() LocalNugetRepositoryArrayOutput
	ToLocalNugetRepositoryArrayOutputWithContext(context.Context) LocalNugetRepositoryArrayOutput
}

type LocalNugetRepositoryArray []LocalNugetRepositoryInput

func (LocalNugetRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalNugetRepository)(nil)).Elem()
}

func (i LocalNugetRepositoryArray) ToLocalNugetRepositoryArrayOutput() LocalNugetRepositoryArrayOutput {
	return i.ToLocalNugetRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalNugetRepositoryArray) ToLocalNugetRepositoryArrayOutputWithContext(ctx context.Context) LocalNugetRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNugetRepositoryArrayOutput)
}

// LocalNugetRepositoryMapInput is an input type that accepts LocalNugetRepositoryMap and LocalNugetRepositoryMapOutput values.
// You can construct a concrete instance of `LocalNugetRepositoryMapInput` via:
//
//          LocalNugetRepositoryMap{ "key": LocalNugetRepositoryArgs{...} }
type LocalNugetRepositoryMapInput interface {
	pulumi.Input

	ToLocalNugetRepositoryMapOutput() LocalNugetRepositoryMapOutput
	ToLocalNugetRepositoryMapOutputWithContext(context.Context) LocalNugetRepositoryMapOutput
}

type LocalNugetRepositoryMap map[string]LocalNugetRepositoryInput

func (LocalNugetRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalNugetRepository)(nil)).Elem()
}

func (i LocalNugetRepositoryMap) ToLocalNugetRepositoryMapOutput() LocalNugetRepositoryMapOutput {
	return i.ToLocalNugetRepositoryMapOutputWithContext(context.Background())
}

func (i LocalNugetRepositoryMap) ToLocalNugetRepositoryMapOutputWithContext(ctx context.Context) LocalNugetRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalNugetRepositoryMapOutput)
}

type LocalNugetRepositoryOutput struct{ *pulumi.OutputState }

func (LocalNugetRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalNugetRepository)(nil))
}

func (o LocalNugetRepositoryOutput) ToLocalNugetRepositoryOutput() LocalNugetRepositoryOutput {
	return o
}

func (o LocalNugetRepositoryOutput) ToLocalNugetRepositoryOutputWithContext(ctx context.Context) LocalNugetRepositoryOutput {
	return o
}

func (o LocalNugetRepositoryOutput) ToLocalNugetRepositoryPtrOutput() LocalNugetRepositoryPtrOutput {
	return o.ToLocalNugetRepositoryPtrOutputWithContext(context.Background())
}

func (o LocalNugetRepositoryOutput) ToLocalNugetRepositoryPtrOutputWithContext(ctx context.Context) LocalNugetRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalNugetRepository) *LocalNugetRepository {
		return &v
	}).(LocalNugetRepositoryPtrOutput)
}

type LocalNugetRepositoryPtrOutput struct{ *pulumi.OutputState }

func (LocalNugetRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalNugetRepository)(nil))
}

func (o LocalNugetRepositoryPtrOutput) ToLocalNugetRepositoryPtrOutput() LocalNugetRepositoryPtrOutput {
	return o
}

func (o LocalNugetRepositoryPtrOutput) ToLocalNugetRepositoryPtrOutputWithContext(ctx context.Context) LocalNugetRepositoryPtrOutput {
	return o
}

func (o LocalNugetRepositoryPtrOutput) Elem() LocalNugetRepositoryOutput {
	return o.ApplyT(func(v *LocalNugetRepository) LocalNugetRepository {
		if v != nil {
			return *v
		}
		var ret LocalNugetRepository
		return ret
	}).(LocalNugetRepositoryOutput)
}

type LocalNugetRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalNugetRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalNugetRepository)(nil))
}

func (o LocalNugetRepositoryArrayOutput) ToLocalNugetRepositoryArrayOutput() LocalNugetRepositoryArrayOutput {
	return o
}

func (o LocalNugetRepositoryArrayOutput) ToLocalNugetRepositoryArrayOutputWithContext(ctx context.Context) LocalNugetRepositoryArrayOutput {
	return o
}

func (o LocalNugetRepositoryArrayOutput) Index(i pulumi.IntInput) LocalNugetRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocalNugetRepository {
		return vs[0].([]LocalNugetRepository)[vs[1].(int)]
	}).(LocalNugetRepositoryOutput)
}

type LocalNugetRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalNugetRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalNugetRepository)(nil))
}

func (o LocalNugetRepositoryMapOutput) ToLocalNugetRepositoryMapOutput() LocalNugetRepositoryMapOutput {
	return o
}

func (o LocalNugetRepositoryMapOutput) ToLocalNugetRepositoryMapOutputWithContext(ctx context.Context) LocalNugetRepositoryMapOutput {
	return o
}

func (o LocalNugetRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalNugetRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocalNugetRepository {
		return vs[0].(map[string]LocalNugetRepository)[vs[1].(string)]
	}).(LocalNugetRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNugetRepositoryInput)(nil)).Elem(), &LocalNugetRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNugetRepositoryPtrInput)(nil)).Elem(), &LocalNugetRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNugetRepositoryArrayInput)(nil)).Elem(), LocalNugetRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalNugetRepositoryMapInput)(nil)).Elem(), LocalNugetRepositoryMap{})
	pulumi.RegisterOutputType(LocalNugetRepositoryOutput{})
	pulumi.RegisterOutputType(LocalNugetRepositoryPtrOutput{})
	pulumi.RegisterOutputType(LocalNugetRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalNugetRepositoryMapOutput{})
}
