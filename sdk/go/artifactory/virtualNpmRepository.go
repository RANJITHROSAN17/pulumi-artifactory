// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a virtual repository resource with specific npm features.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/npm+Registry#npmRegistry-VirtualnpmRegistry).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/v2/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewVirtualNpmRepository(ctx, "foo-npm", &artifactory.VirtualNpmRepositoryArgs{
// 			Description:     pulumi.String("A test virtual repo"),
// 			ExcludesPattern: pulumi.String("com/google/**"),
// 			IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
// 			Key:             pulumi.String("foo-npm"),
// 			Notes:           pulumi.String("Internal description"),
// 			Repositories:    pulumi.StringArray{},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/virtualNpmRepository:VirtualNpmRepository foo-npm foo-npm
// ```
type VirtualNpmRepository struct {
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

// NewVirtualNpmRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualNpmRepository(ctx *pulumi.Context,
	name string, args *VirtualNpmRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualNpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualNpmRepository
	err := ctx.RegisterResource("artifactory:index/virtualNpmRepository:VirtualNpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNpmRepository gets an existing VirtualNpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNpmRepositoryState, opts ...pulumi.ResourceOption) (*VirtualNpmRepository, error) {
	var resource VirtualNpmRepository
	err := ctx.ReadResource("artifactory:index/virtualNpmRepository:VirtualNpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNpmRepository resources.
type virtualNpmRepositoryState struct {
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

type VirtualNpmRepositoryState struct {
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

func (VirtualNpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNpmRepositoryState)(nil)).Elem()
}

type virtualNpmRepositoryArgs struct {
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

// The set of arguments for constructing a VirtualNpmRepository resource.
type VirtualNpmRepositoryArgs struct {
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

func (VirtualNpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNpmRepositoryArgs)(nil)).Elem()
}

type VirtualNpmRepositoryInput interface {
	pulumi.Input

	ToVirtualNpmRepositoryOutput() VirtualNpmRepositoryOutput
	ToVirtualNpmRepositoryOutputWithContext(ctx context.Context) VirtualNpmRepositoryOutput
}

func (*VirtualNpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNpmRepository)(nil)).Elem()
}

func (i *VirtualNpmRepository) ToVirtualNpmRepositoryOutput() VirtualNpmRepositoryOutput {
	return i.ToVirtualNpmRepositoryOutputWithContext(context.Background())
}

func (i *VirtualNpmRepository) ToVirtualNpmRepositoryOutputWithContext(ctx context.Context) VirtualNpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNpmRepositoryOutput)
}

// VirtualNpmRepositoryArrayInput is an input type that accepts VirtualNpmRepositoryArray and VirtualNpmRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualNpmRepositoryArrayInput` via:
//
//          VirtualNpmRepositoryArray{ VirtualNpmRepositoryArgs{...} }
type VirtualNpmRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualNpmRepositoryArrayOutput() VirtualNpmRepositoryArrayOutput
	ToVirtualNpmRepositoryArrayOutputWithContext(context.Context) VirtualNpmRepositoryArrayOutput
}

type VirtualNpmRepositoryArray []VirtualNpmRepositoryInput

func (VirtualNpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualNpmRepository)(nil)).Elem()
}

func (i VirtualNpmRepositoryArray) ToVirtualNpmRepositoryArrayOutput() VirtualNpmRepositoryArrayOutput {
	return i.ToVirtualNpmRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualNpmRepositoryArray) ToVirtualNpmRepositoryArrayOutputWithContext(ctx context.Context) VirtualNpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNpmRepositoryArrayOutput)
}

// VirtualNpmRepositoryMapInput is an input type that accepts VirtualNpmRepositoryMap and VirtualNpmRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualNpmRepositoryMapInput` via:
//
//          VirtualNpmRepositoryMap{ "key": VirtualNpmRepositoryArgs{...} }
type VirtualNpmRepositoryMapInput interface {
	pulumi.Input

	ToVirtualNpmRepositoryMapOutput() VirtualNpmRepositoryMapOutput
	ToVirtualNpmRepositoryMapOutputWithContext(context.Context) VirtualNpmRepositoryMapOutput
}

type VirtualNpmRepositoryMap map[string]VirtualNpmRepositoryInput

func (VirtualNpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualNpmRepository)(nil)).Elem()
}

func (i VirtualNpmRepositoryMap) ToVirtualNpmRepositoryMapOutput() VirtualNpmRepositoryMapOutput {
	return i.ToVirtualNpmRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualNpmRepositoryMap) ToVirtualNpmRepositoryMapOutputWithContext(ctx context.Context) VirtualNpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNpmRepositoryMapOutput)
}

type VirtualNpmRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualNpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNpmRepository)(nil)).Elem()
}

func (o VirtualNpmRepositoryOutput) ToVirtualNpmRepositoryOutput() VirtualNpmRepositoryOutput {
	return o
}

func (o VirtualNpmRepositoryOutput) ToVirtualNpmRepositoryOutputWithContext(ctx context.Context) VirtualNpmRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualNpmRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualNpmRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
func (o VirtualNpmRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualNpmRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualNpmRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualNpmRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// A free text field to add additional notes about the repository. These are only visible to the administrator.
func (o VirtualNpmRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
func (o VirtualNpmRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
func (o VirtualNpmRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
// repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualNpmRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualNpmRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualNpmRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
// repositories. A value of 0 indicates no caching.
func (o VirtualNpmRepositoryOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNpmRepository) pulumi.IntPtrOutput { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

type VirtualNpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualNpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualNpmRepository)(nil)).Elem()
}

func (o VirtualNpmRepositoryArrayOutput) ToVirtualNpmRepositoryArrayOutput() VirtualNpmRepositoryArrayOutput {
	return o
}

func (o VirtualNpmRepositoryArrayOutput) ToVirtualNpmRepositoryArrayOutputWithContext(ctx context.Context) VirtualNpmRepositoryArrayOutput {
	return o
}

func (o VirtualNpmRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualNpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualNpmRepository {
		return vs[0].([]*VirtualNpmRepository)[vs[1].(int)]
	}).(VirtualNpmRepositoryOutput)
}

type VirtualNpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualNpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualNpmRepository)(nil)).Elem()
}

func (o VirtualNpmRepositoryMapOutput) ToVirtualNpmRepositoryMapOutput() VirtualNpmRepositoryMapOutput {
	return o
}

func (o VirtualNpmRepositoryMapOutput) ToVirtualNpmRepositoryMapOutputWithContext(ctx context.Context) VirtualNpmRepositoryMapOutput {
	return o
}

func (o VirtualNpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualNpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualNpmRepository {
		return vs[0].(map[string]*VirtualNpmRepository)[vs[1].(string)]
	}).(VirtualNpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNpmRepositoryInput)(nil)).Elem(), &VirtualNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNpmRepositoryArrayInput)(nil)).Elem(), VirtualNpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNpmRepositoryMapInput)(nil)).Elem(), VirtualNpmRepositoryMap{})
	pulumi.RegisterOutputType(VirtualNpmRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualNpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualNpmRepositoryMapOutput{})
}
