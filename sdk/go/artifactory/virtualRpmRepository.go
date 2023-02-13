// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a virtual Rpm repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/RPM+Repositories).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"os"
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewKeypair(ctx, "primary-keypair", &artifactory.KeypairArgs{
//				PairName:   pulumi.String("primary-keypair"),
//				PairType:   pulumi.String("GPG"),
//				Alias:      pulumi.String("foo-alias-1"),
//				PrivateKey: readFileOrPanic("samples/gpg.priv"),
//				PublicKey:  readFileOrPanic("samples/gpg.pub"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewKeypair(ctx, "secondary-keypair", &artifactory.KeypairArgs{
//				PairName:   pulumi.String("secondary-keypair"),
//				PairType:   pulumi.String("GPG"),
//				Alias:      pulumi.String("foo-alias-2"),
//				PrivateKey: readFileOrPanic("samples/gpg.priv"),
//				PublicKey:  readFileOrPanic("samples/gpg.pub"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewVirtualRpmRepository(ctx, "foo-rpm-virtual", &artifactory.VirtualRpmRepositoryArgs{
//				Key:                 pulumi.String("foo-rpm-virtual"),
//				PrimaryKeypairRef:   primary_keypair.PairName,
//				SecondaryKeypairRef: secondary_keypair.PairName,
//			}, pulumi.DependsOn([]pulumi.Resource{
//				primary_keypair,
//				secondary_keypair,
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
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/virtualRpmRepository:VirtualRpmRepository foo-rpm-virtual foo-rpm-virtual
//
// ```
type VirtualRpmRepository struct {
	pulumi.CustomResourceState

	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrOutput `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrOutput `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringOutput `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef pulumi.StringPtrOutput `pulumi:"primaryKeypairRef"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrOutput `pulumi:"secondaryKeypairRef"`
}

// NewVirtualRpmRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualRpmRepository(ctx *pulumi.Context,
	name string, args *VirtualRpmRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualRpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualRpmRepository
	err := ctx.RegisterResource("artifactory:index/virtualRpmRepository:VirtualRpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualRpmRepository gets an existing VirtualRpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualRpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualRpmRepositoryState, opts ...pulumi.ResourceOption) (*VirtualRpmRepository, error) {
	var resource VirtualRpmRepository
	err := ctx.ReadResource("artifactory:index/virtualRpmRepository:VirtualRpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualRpmRepository resources.
type virtualRpmRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key *string `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType *string `pulumi:"packageType"`
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

type VirtualRpmRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringPtrInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringPtrInput
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrInput
}

func (VirtualRpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRpmRepositoryState)(nil)).Elem()
}

type virtualRpmRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key string `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef *string `pulumi:"secondaryKeypairRef"`
}

// The set of arguments for constructing a VirtualRpmRepository resource.
type VirtualRpmRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
	// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// The primary GPG key to be used to sign packages.
	PrimaryKeypairRef pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// The secondary GPG key to be used to sign packages.
	SecondaryKeypairRef pulumi.StringPtrInput
}

func (VirtualRpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRpmRepositoryArgs)(nil)).Elem()
}

type VirtualRpmRepositoryInput interface {
	pulumi.Input

	ToVirtualRpmRepositoryOutput() VirtualRpmRepositoryOutput
	ToVirtualRpmRepositoryOutputWithContext(ctx context.Context) VirtualRpmRepositoryOutput
}

func (*VirtualRpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRpmRepository)(nil)).Elem()
}

func (i *VirtualRpmRepository) ToVirtualRpmRepositoryOutput() VirtualRpmRepositoryOutput {
	return i.ToVirtualRpmRepositoryOutputWithContext(context.Background())
}

func (i *VirtualRpmRepository) ToVirtualRpmRepositoryOutputWithContext(ctx context.Context) VirtualRpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRpmRepositoryOutput)
}

// VirtualRpmRepositoryArrayInput is an input type that accepts VirtualRpmRepositoryArray and VirtualRpmRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualRpmRepositoryArrayInput` via:
//
//	VirtualRpmRepositoryArray{ VirtualRpmRepositoryArgs{...} }
type VirtualRpmRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualRpmRepositoryArrayOutput() VirtualRpmRepositoryArrayOutput
	ToVirtualRpmRepositoryArrayOutputWithContext(context.Context) VirtualRpmRepositoryArrayOutput
}

type VirtualRpmRepositoryArray []VirtualRpmRepositoryInput

func (VirtualRpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualRpmRepository)(nil)).Elem()
}

func (i VirtualRpmRepositoryArray) ToVirtualRpmRepositoryArrayOutput() VirtualRpmRepositoryArrayOutput {
	return i.ToVirtualRpmRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualRpmRepositoryArray) ToVirtualRpmRepositoryArrayOutputWithContext(ctx context.Context) VirtualRpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRpmRepositoryArrayOutput)
}

// VirtualRpmRepositoryMapInput is an input type that accepts VirtualRpmRepositoryMap and VirtualRpmRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualRpmRepositoryMapInput` via:
//
//	VirtualRpmRepositoryMap{ "key": VirtualRpmRepositoryArgs{...} }
type VirtualRpmRepositoryMapInput interface {
	pulumi.Input

	ToVirtualRpmRepositoryMapOutput() VirtualRpmRepositoryMapOutput
	ToVirtualRpmRepositoryMapOutputWithContext(context.Context) VirtualRpmRepositoryMapOutput
}

type VirtualRpmRepositoryMap map[string]VirtualRpmRepositoryInput

func (VirtualRpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualRpmRepository)(nil)).Elem()
}

func (i VirtualRpmRepositoryMap) ToVirtualRpmRepositoryMapOutput() VirtualRpmRepositoryMapOutput {
	return i.ToVirtualRpmRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualRpmRepositoryMap) ToVirtualRpmRepositoryMapOutputWithContext(ctx context.Context) VirtualRpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualRpmRepositoryMapOutput)
}

type VirtualRpmRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualRpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualRpmRepository)(nil)).Elem()
}

func (o VirtualRpmRepositoryOutput) ToVirtualRpmRepositoryOutput() VirtualRpmRepositoryOutput {
	return o
}

func (o VirtualRpmRepositoryOutput) ToVirtualRpmRepositoryOutputWithContext(ctx context.Context) VirtualRpmRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualRpmRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualRpmRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
func (o VirtualRpmRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualRpmRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualRpmRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualRpmRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// A free text field to add additional notes about the repository. These are only visible to the administrator.
func (o VirtualRpmRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
func (o VirtualRpmRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// The primary GPG key to be used to sign packages.
func (o VirtualRpmRepositoryOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o VirtualRpmRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 10 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualRpmRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualRpmRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualRpmRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// The secondary GPG key to be used to sign packages.
func (o VirtualRpmRepositoryOutput) SecondaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualRpmRepository) pulumi.StringPtrOutput { return v.SecondaryKeypairRef }).(pulumi.StringPtrOutput)
}

type VirtualRpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualRpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualRpmRepository)(nil)).Elem()
}

func (o VirtualRpmRepositoryArrayOutput) ToVirtualRpmRepositoryArrayOutput() VirtualRpmRepositoryArrayOutput {
	return o
}

func (o VirtualRpmRepositoryArrayOutput) ToVirtualRpmRepositoryArrayOutputWithContext(ctx context.Context) VirtualRpmRepositoryArrayOutput {
	return o
}

func (o VirtualRpmRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualRpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualRpmRepository {
		return vs[0].([]*VirtualRpmRepository)[vs[1].(int)]
	}).(VirtualRpmRepositoryOutput)
}

type VirtualRpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualRpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualRpmRepository)(nil)).Elem()
}

func (o VirtualRpmRepositoryMapOutput) ToVirtualRpmRepositoryMapOutput() VirtualRpmRepositoryMapOutput {
	return o
}

func (o VirtualRpmRepositoryMapOutput) ToVirtualRpmRepositoryMapOutputWithContext(ctx context.Context) VirtualRpmRepositoryMapOutput {
	return o
}

func (o VirtualRpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualRpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualRpmRepository {
		return vs[0].(map[string]*VirtualRpmRepository)[vs[1].(string)]
	}).(VirtualRpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualRpmRepositoryInput)(nil)).Elem(), &VirtualRpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualRpmRepositoryArrayInput)(nil)).Elem(), VirtualRpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualRpmRepositoryMapInput)(nil)).Elem(), VirtualRpmRepositoryMap{})
	pulumi.RegisterOutputType(VirtualRpmRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualRpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualRpmRepositoryMapOutput{})
}
