// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Local Helm Repository Resource
//
// Creates a local helm repository.
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
// 		_, err := artifactory.NewLocalHelmRepository(ctx, "terraform_local_test_helm_repo", &artifactory.LocalHelmRepositoryArgs{
// 			Key: pulumi.String("terraform-local-test-helm-repo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LocalHelmRepository struct {
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
	Key           pulumi.StringOutput      `pulumi:"key"`
	Notes         pulumi.StringPtrOutput   `pulumi:"notes"`
	PackageType   pulumi.StringOutput      `pulumi:"packageType"`
	PropertySets  pulumi.StringArrayOutput `pulumi:"propertySets"`
	RepoLayoutRef pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	XrayIndex     pulumi.BoolOutput        `pulumi:"xrayIndex"`
}

// NewLocalHelmRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalHelmRepository(ctx *pulumi.Context,
	name string, args *LocalHelmRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalHelmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource LocalHelmRepository
	err := ctx.RegisterResource("artifactory:index/localHelmRepository:LocalHelmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalHelmRepository gets an existing LocalHelmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalHelmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalHelmRepositoryState, opts ...pulumi.ResourceOption) (*LocalHelmRepository, error) {
	var resource LocalHelmRepository
	err := ctx.ReadResource("artifactory:index/localHelmRepository:LocalHelmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalHelmRepository resources.
type localHelmRepositoryState struct {
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
	Key           *string  `pulumi:"key"`
	Notes         *string  `pulumi:"notes"`
	PackageType   *string  `pulumi:"packageType"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

type LocalHelmRepositoryState struct {
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
	Key           pulumi.StringPtrInput
	Notes         pulumi.StringPtrInput
	PackageType   pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (LocalHelmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localHelmRepositoryState)(nil)).Elem()
}

type localHelmRepositoryArgs struct {
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
	Key           string   `pulumi:"key"`
	Notes         *string  `pulumi:"notes"`
	PropertySets  []string `pulumi:"propertySets"`
	RepoLayoutRef *string  `pulumi:"repoLayoutRef"`
	XrayIndex     *bool    `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalHelmRepository resource.
type LocalHelmRepositoryArgs struct {
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
	Key           pulumi.StringInput
	Notes         pulumi.StringPtrInput
	PropertySets  pulumi.StringArrayInput
	RepoLayoutRef pulumi.StringPtrInput
	XrayIndex     pulumi.BoolPtrInput
}

func (LocalHelmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localHelmRepositoryArgs)(nil)).Elem()
}

type LocalHelmRepositoryInput interface {
	pulumi.Input

	ToLocalHelmRepositoryOutput() LocalHelmRepositoryOutput
	ToLocalHelmRepositoryOutputWithContext(ctx context.Context) LocalHelmRepositoryOutput
}

func (*LocalHelmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalHelmRepository)(nil))
}

func (i *LocalHelmRepository) ToLocalHelmRepositoryOutput() LocalHelmRepositoryOutput {
	return i.ToLocalHelmRepositoryOutputWithContext(context.Background())
}

func (i *LocalHelmRepository) ToLocalHelmRepositoryOutputWithContext(ctx context.Context) LocalHelmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalHelmRepositoryOutput)
}

func (i *LocalHelmRepository) ToLocalHelmRepositoryPtrOutput() LocalHelmRepositoryPtrOutput {
	return i.ToLocalHelmRepositoryPtrOutputWithContext(context.Background())
}

func (i *LocalHelmRepository) ToLocalHelmRepositoryPtrOutputWithContext(ctx context.Context) LocalHelmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalHelmRepositoryPtrOutput)
}

type LocalHelmRepositoryPtrInput interface {
	pulumi.Input

	ToLocalHelmRepositoryPtrOutput() LocalHelmRepositoryPtrOutput
	ToLocalHelmRepositoryPtrOutputWithContext(ctx context.Context) LocalHelmRepositoryPtrOutput
}

type localHelmRepositoryPtrType LocalHelmRepositoryArgs

func (*localHelmRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalHelmRepository)(nil))
}

func (i *localHelmRepositoryPtrType) ToLocalHelmRepositoryPtrOutput() LocalHelmRepositoryPtrOutput {
	return i.ToLocalHelmRepositoryPtrOutputWithContext(context.Background())
}

func (i *localHelmRepositoryPtrType) ToLocalHelmRepositoryPtrOutputWithContext(ctx context.Context) LocalHelmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalHelmRepositoryPtrOutput)
}

// LocalHelmRepositoryArrayInput is an input type that accepts LocalHelmRepositoryArray and LocalHelmRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalHelmRepositoryArrayInput` via:
//
//          LocalHelmRepositoryArray{ LocalHelmRepositoryArgs{...} }
type LocalHelmRepositoryArrayInput interface {
	pulumi.Input

	ToLocalHelmRepositoryArrayOutput() LocalHelmRepositoryArrayOutput
	ToLocalHelmRepositoryArrayOutputWithContext(context.Context) LocalHelmRepositoryArrayOutput
}

type LocalHelmRepositoryArray []LocalHelmRepositoryInput

func (LocalHelmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalHelmRepository)(nil)).Elem()
}

func (i LocalHelmRepositoryArray) ToLocalHelmRepositoryArrayOutput() LocalHelmRepositoryArrayOutput {
	return i.ToLocalHelmRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalHelmRepositoryArray) ToLocalHelmRepositoryArrayOutputWithContext(ctx context.Context) LocalHelmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalHelmRepositoryArrayOutput)
}

// LocalHelmRepositoryMapInput is an input type that accepts LocalHelmRepositoryMap and LocalHelmRepositoryMapOutput values.
// You can construct a concrete instance of `LocalHelmRepositoryMapInput` via:
//
//          LocalHelmRepositoryMap{ "key": LocalHelmRepositoryArgs{...} }
type LocalHelmRepositoryMapInput interface {
	pulumi.Input

	ToLocalHelmRepositoryMapOutput() LocalHelmRepositoryMapOutput
	ToLocalHelmRepositoryMapOutputWithContext(context.Context) LocalHelmRepositoryMapOutput
}

type LocalHelmRepositoryMap map[string]LocalHelmRepositoryInput

func (LocalHelmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalHelmRepository)(nil)).Elem()
}

func (i LocalHelmRepositoryMap) ToLocalHelmRepositoryMapOutput() LocalHelmRepositoryMapOutput {
	return i.ToLocalHelmRepositoryMapOutputWithContext(context.Background())
}

func (i LocalHelmRepositoryMap) ToLocalHelmRepositoryMapOutputWithContext(ctx context.Context) LocalHelmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalHelmRepositoryMapOutput)
}

type LocalHelmRepositoryOutput struct{ *pulumi.OutputState }

func (LocalHelmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocalHelmRepository)(nil))
}

func (o LocalHelmRepositoryOutput) ToLocalHelmRepositoryOutput() LocalHelmRepositoryOutput {
	return o
}

func (o LocalHelmRepositoryOutput) ToLocalHelmRepositoryOutputWithContext(ctx context.Context) LocalHelmRepositoryOutput {
	return o
}

func (o LocalHelmRepositoryOutput) ToLocalHelmRepositoryPtrOutput() LocalHelmRepositoryPtrOutput {
	return o.ToLocalHelmRepositoryPtrOutputWithContext(context.Background())
}

func (o LocalHelmRepositoryOutput) ToLocalHelmRepositoryPtrOutputWithContext(ctx context.Context) LocalHelmRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LocalHelmRepository) *LocalHelmRepository {
		return &v
	}).(LocalHelmRepositoryPtrOutput)
}

type LocalHelmRepositoryPtrOutput struct{ *pulumi.OutputState }

func (LocalHelmRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalHelmRepository)(nil))
}

func (o LocalHelmRepositoryPtrOutput) ToLocalHelmRepositoryPtrOutput() LocalHelmRepositoryPtrOutput {
	return o
}

func (o LocalHelmRepositoryPtrOutput) ToLocalHelmRepositoryPtrOutputWithContext(ctx context.Context) LocalHelmRepositoryPtrOutput {
	return o
}

func (o LocalHelmRepositoryPtrOutput) Elem() LocalHelmRepositoryOutput {
	return o.ApplyT(func(v *LocalHelmRepository) LocalHelmRepository {
		if v != nil {
			return *v
		}
		var ret LocalHelmRepository
		return ret
	}).(LocalHelmRepositoryOutput)
}

type LocalHelmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalHelmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocalHelmRepository)(nil))
}

func (o LocalHelmRepositoryArrayOutput) ToLocalHelmRepositoryArrayOutput() LocalHelmRepositoryArrayOutput {
	return o
}

func (o LocalHelmRepositoryArrayOutput) ToLocalHelmRepositoryArrayOutputWithContext(ctx context.Context) LocalHelmRepositoryArrayOutput {
	return o
}

func (o LocalHelmRepositoryArrayOutput) Index(i pulumi.IntInput) LocalHelmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocalHelmRepository {
		return vs[0].([]LocalHelmRepository)[vs[1].(int)]
	}).(LocalHelmRepositoryOutput)
}

type LocalHelmRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalHelmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LocalHelmRepository)(nil))
}

func (o LocalHelmRepositoryMapOutput) ToLocalHelmRepositoryMapOutput() LocalHelmRepositoryMapOutput {
	return o
}

func (o LocalHelmRepositoryMapOutput) ToLocalHelmRepositoryMapOutputWithContext(ctx context.Context) LocalHelmRepositoryMapOutput {
	return o
}

func (o LocalHelmRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalHelmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LocalHelmRepository {
		return vs[0].(map[string]LocalHelmRepository)[vs[1].(string)]
	}).(LocalHelmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalHelmRepositoryInput)(nil)).Elem(), &LocalHelmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalHelmRepositoryPtrInput)(nil)).Elem(), &LocalHelmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalHelmRepositoryArrayInput)(nil)).Elem(), LocalHelmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalHelmRepositoryMapInput)(nil)).Elem(), LocalHelmRepositoryMap{})
	pulumi.RegisterOutputType(LocalHelmRepositoryOutput{})
	pulumi.RegisterOutputType(LocalHelmRepositoryPtrOutput{})
	pulumi.RegisterOutputType(LocalHelmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalHelmRepositoryMapOutput{})
}
