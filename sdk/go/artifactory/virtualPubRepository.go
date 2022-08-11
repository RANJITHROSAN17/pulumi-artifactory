// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a virtual Pub repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Pub+Repositories#PubRepositories-SettingupaVirtualRepository).
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
//			_, err := artifactory.NewVirtualPubRepository(ctx, "foo-pub", &artifactory.VirtualPubRepositoryArgs{
//				Description:     pulumi.String("A test virtual repo"),
//				ExcludesPattern: pulumi.String("com/google/**"),
//				IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				Key:             pulumi.String("foo-pub"),
//				Notes:           pulumi.String("Internal description"),
//				Repositories:    pulumi.StringArray{},
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
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/virtualPubRepository:VirtualPubRepository foo-pub foo-pub
//
// ```
type VirtualPubRepository struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringOutput `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrOutput `pulumi:"retrievalCachePeriodSeconds"`
}

// NewVirtualPubRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualPubRepository(ctx *pulumi.Context,
	name string, args *VirtualPubRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualPubRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualPubRepository
	err := ctx.RegisterResource("artifactory:index/virtualPubRepository:VirtualPubRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualPubRepository gets an existing VirtualPubRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualPubRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualPubRepositoryState, opts ...pulumi.ResourceOption) (*VirtualPubRepository, error) {
	var resource VirtualPubRepository
	err := ctx.ReadResource("artifactory:index/virtualPubRepository:VirtualPubRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualPubRepository resources.
type virtualPubRepositoryState struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key *string `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType *string `pulumi:"packageType"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

type VirtualPubRepositoryState struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringPtrInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualPubRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualPubRepositoryState)(nil)).Elem()
}

type virtualPubRepositoryArgs struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key string `pulumi:"key"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

// The set of arguments for constructing a VirtualPubRepository resource.
type VirtualPubRepositoryArgs struct {
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
	// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
	// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
	// repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
	// repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualPubRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualPubRepositoryArgs)(nil)).Elem()
}

type VirtualPubRepositoryInput interface {
	pulumi.Input

	ToVirtualPubRepositoryOutput() VirtualPubRepositoryOutput
	ToVirtualPubRepositoryOutputWithContext(ctx context.Context) VirtualPubRepositoryOutput
}

func (*VirtualPubRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualPubRepository)(nil)).Elem()
}

func (i *VirtualPubRepository) ToVirtualPubRepositoryOutput() VirtualPubRepositoryOutput {
	return i.ToVirtualPubRepositoryOutputWithContext(context.Background())
}

func (i *VirtualPubRepository) ToVirtualPubRepositoryOutputWithContext(ctx context.Context) VirtualPubRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPubRepositoryOutput)
}

// VirtualPubRepositoryArrayInput is an input type that accepts VirtualPubRepositoryArray and VirtualPubRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualPubRepositoryArrayInput` via:
//
//	VirtualPubRepositoryArray{ VirtualPubRepositoryArgs{...} }
type VirtualPubRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualPubRepositoryArrayOutput() VirtualPubRepositoryArrayOutput
	ToVirtualPubRepositoryArrayOutputWithContext(context.Context) VirtualPubRepositoryArrayOutput
}

type VirtualPubRepositoryArray []VirtualPubRepositoryInput

func (VirtualPubRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualPubRepository)(nil)).Elem()
}

func (i VirtualPubRepositoryArray) ToVirtualPubRepositoryArrayOutput() VirtualPubRepositoryArrayOutput {
	return i.ToVirtualPubRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualPubRepositoryArray) ToVirtualPubRepositoryArrayOutputWithContext(ctx context.Context) VirtualPubRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPubRepositoryArrayOutput)
}

// VirtualPubRepositoryMapInput is an input type that accepts VirtualPubRepositoryMap and VirtualPubRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualPubRepositoryMapInput` via:
//
//	VirtualPubRepositoryMap{ "key": VirtualPubRepositoryArgs{...} }
type VirtualPubRepositoryMapInput interface {
	pulumi.Input

	ToVirtualPubRepositoryMapOutput() VirtualPubRepositoryMapOutput
	ToVirtualPubRepositoryMapOutputWithContext(context.Context) VirtualPubRepositoryMapOutput
}

type VirtualPubRepositoryMap map[string]VirtualPubRepositoryInput

func (VirtualPubRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualPubRepository)(nil)).Elem()
}

func (i VirtualPubRepositoryMap) ToVirtualPubRepositoryMapOutput() VirtualPubRepositoryMapOutput {
	return i.ToVirtualPubRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualPubRepositoryMap) ToVirtualPubRepositoryMapOutputWithContext(ctx context.Context) VirtualPubRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualPubRepositoryMapOutput)
}

type VirtualPubRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualPubRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualPubRepository)(nil)).Elem()
}

func (o VirtualPubRepositoryOutput) ToVirtualPubRepositoryOutput() VirtualPubRepositoryOutput {
	return o
}

func (o VirtualPubRepositoryOutput) ToVirtualPubRepositoryOutputWithContext(ctx context.Context) VirtualPubRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualPubRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualPubRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
func (o VirtualPubRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualPubRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualPubRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualPubRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// A free text field to add additional notes about the repository. These are only visible to the administrator.
func (o VirtualPubRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
func (o VirtualPubRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
func (o VirtualPubRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
// repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualPubRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualPubRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualPubRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
// repositories. A value of 0 indicates no caching.
func (o VirtualPubRepositoryOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualPubRepository) pulumi.IntPtrOutput { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

type VirtualPubRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualPubRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualPubRepository)(nil)).Elem()
}

func (o VirtualPubRepositoryArrayOutput) ToVirtualPubRepositoryArrayOutput() VirtualPubRepositoryArrayOutput {
	return o
}

func (o VirtualPubRepositoryArrayOutput) ToVirtualPubRepositoryArrayOutputWithContext(ctx context.Context) VirtualPubRepositoryArrayOutput {
	return o
}

func (o VirtualPubRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualPubRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualPubRepository {
		return vs[0].([]*VirtualPubRepository)[vs[1].(int)]
	}).(VirtualPubRepositoryOutput)
}

type VirtualPubRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualPubRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualPubRepository)(nil)).Elem()
}

func (o VirtualPubRepositoryMapOutput) ToVirtualPubRepositoryMapOutput() VirtualPubRepositoryMapOutput {
	return o
}

func (o VirtualPubRepositoryMapOutput) ToVirtualPubRepositoryMapOutputWithContext(ctx context.Context) VirtualPubRepositoryMapOutput {
	return o
}

func (o VirtualPubRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualPubRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualPubRepository {
		return vs[0].(map[string]*VirtualPubRepository)[vs[1].(string)]
	}).(VirtualPubRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPubRepositoryInput)(nil)).Elem(), &VirtualPubRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPubRepositoryArrayInput)(nil)).Elem(), VirtualPubRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualPubRepositoryMapInput)(nil)).Elem(), VirtualPubRepositoryMap{})
	pulumi.RegisterOutputType(VirtualPubRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualPubRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualPubRepositoryMapOutput{})
}
