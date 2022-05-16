// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a virtual Git LFS repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Git+LFS+Repositories#GitLFSRepositories-VirtualRepositories).
//
// ## Import
//
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/virtualGitlfsRepository:VirtualGitlfsRepository foo-gitlfs foo-gitlfs
// ```
type VirtualGitlfsRepository struct {
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

// NewVirtualGitlfsRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualGitlfsRepository(ctx *pulumi.Context,
	name string, args *VirtualGitlfsRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualGitlfsRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource VirtualGitlfsRepository
	err := ctx.RegisterResource("artifactory:index/virtualGitlfsRepository:VirtualGitlfsRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualGitlfsRepository gets an existing VirtualGitlfsRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualGitlfsRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualGitlfsRepositoryState, opts ...pulumi.ResourceOption) (*VirtualGitlfsRepository, error) {
	var resource VirtualGitlfsRepository
	err := ctx.ReadResource("artifactory:index/virtualGitlfsRepository:VirtualGitlfsRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualGitlfsRepository resources.
type virtualGitlfsRepositoryState struct {
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

type VirtualGitlfsRepositoryState struct {
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

func (VirtualGitlfsRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGitlfsRepositoryState)(nil)).Elem()
}

type virtualGitlfsRepositoryArgs struct {
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

// The set of arguments for constructing a VirtualGitlfsRepository resource.
type VirtualGitlfsRepositoryArgs struct {
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

func (VirtualGitlfsRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualGitlfsRepositoryArgs)(nil)).Elem()
}

type VirtualGitlfsRepositoryInput interface {
	pulumi.Input

	ToVirtualGitlfsRepositoryOutput() VirtualGitlfsRepositoryOutput
	ToVirtualGitlfsRepositoryOutputWithContext(ctx context.Context) VirtualGitlfsRepositoryOutput
}

func (*VirtualGitlfsRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGitlfsRepository)(nil)).Elem()
}

func (i *VirtualGitlfsRepository) ToVirtualGitlfsRepositoryOutput() VirtualGitlfsRepositoryOutput {
	return i.ToVirtualGitlfsRepositoryOutputWithContext(context.Background())
}

func (i *VirtualGitlfsRepository) ToVirtualGitlfsRepositoryOutputWithContext(ctx context.Context) VirtualGitlfsRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGitlfsRepositoryOutput)
}

// VirtualGitlfsRepositoryArrayInput is an input type that accepts VirtualGitlfsRepositoryArray and VirtualGitlfsRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualGitlfsRepositoryArrayInput` via:
//
//          VirtualGitlfsRepositoryArray{ VirtualGitlfsRepositoryArgs{...} }
type VirtualGitlfsRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualGitlfsRepositoryArrayOutput() VirtualGitlfsRepositoryArrayOutput
	ToVirtualGitlfsRepositoryArrayOutputWithContext(context.Context) VirtualGitlfsRepositoryArrayOutput
}

type VirtualGitlfsRepositoryArray []VirtualGitlfsRepositoryInput

func (VirtualGitlfsRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGitlfsRepository)(nil)).Elem()
}

func (i VirtualGitlfsRepositoryArray) ToVirtualGitlfsRepositoryArrayOutput() VirtualGitlfsRepositoryArrayOutput {
	return i.ToVirtualGitlfsRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualGitlfsRepositoryArray) ToVirtualGitlfsRepositoryArrayOutputWithContext(ctx context.Context) VirtualGitlfsRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGitlfsRepositoryArrayOutput)
}

// VirtualGitlfsRepositoryMapInput is an input type that accepts VirtualGitlfsRepositoryMap and VirtualGitlfsRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualGitlfsRepositoryMapInput` via:
//
//          VirtualGitlfsRepositoryMap{ "key": VirtualGitlfsRepositoryArgs{...} }
type VirtualGitlfsRepositoryMapInput interface {
	pulumi.Input

	ToVirtualGitlfsRepositoryMapOutput() VirtualGitlfsRepositoryMapOutput
	ToVirtualGitlfsRepositoryMapOutputWithContext(context.Context) VirtualGitlfsRepositoryMapOutput
}

type VirtualGitlfsRepositoryMap map[string]VirtualGitlfsRepositoryInput

func (VirtualGitlfsRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGitlfsRepository)(nil)).Elem()
}

func (i VirtualGitlfsRepositoryMap) ToVirtualGitlfsRepositoryMapOutput() VirtualGitlfsRepositoryMapOutput {
	return i.ToVirtualGitlfsRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualGitlfsRepositoryMap) ToVirtualGitlfsRepositoryMapOutputWithContext(ctx context.Context) VirtualGitlfsRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualGitlfsRepositoryMapOutput)
}

type VirtualGitlfsRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualGitlfsRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualGitlfsRepository)(nil)).Elem()
}

func (o VirtualGitlfsRepositoryOutput) ToVirtualGitlfsRepositoryOutput() VirtualGitlfsRepositoryOutput {
	return o
}

func (o VirtualGitlfsRepositoryOutput) ToVirtualGitlfsRepositoryOutputWithContext(ctx context.Context) VirtualGitlfsRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualGitlfsRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualGitlfsRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
func (o VirtualGitlfsRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o VirtualGitlfsRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When used, only
// artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o VirtualGitlfsRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualGitlfsRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// A free text field to add additional notes about the repository. These are only visible to the administrator.
func (o VirtualGitlfsRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
func (o VirtualGitlfsRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD"
func (o VirtualGitlfsRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 3 - 10 lowercase alphanumeric characters. When assigning
// repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualGitlfsRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o VirtualGitlfsRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualGitlfsRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated
// repositories. A value of 0 indicates no caching.
func (o VirtualGitlfsRepositoryOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualGitlfsRepository) pulumi.IntPtrOutput { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

type VirtualGitlfsRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualGitlfsRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualGitlfsRepository)(nil)).Elem()
}

func (o VirtualGitlfsRepositoryArrayOutput) ToVirtualGitlfsRepositoryArrayOutput() VirtualGitlfsRepositoryArrayOutput {
	return o
}

func (o VirtualGitlfsRepositoryArrayOutput) ToVirtualGitlfsRepositoryArrayOutputWithContext(ctx context.Context) VirtualGitlfsRepositoryArrayOutput {
	return o
}

func (o VirtualGitlfsRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualGitlfsRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualGitlfsRepository {
		return vs[0].([]*VirtualGitlfsRepository)[vs[1].(int)]
	}).(VirtualGitlfsRepositoryOutput)
}

type VirtualGitlfsRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualGitlfsRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualGitlfsRepository)(nil)).Elem()
}

func (o VirtualGitlfsRepositoryMapOutput) ToVirtualGitlfsRepositoryMapOutput() VirtualGitlfsRepositoryMapOutput {
	return o
}

func (o VirtualGitlfsRepositoryMapOutput) ToVirtualGitlfsRepositoryMapOutputWithContext(ctx context.Context) VirtualGitlfsRepositoryMapOutput {
	return o
}

func (o VirtualGitlfsRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualGitlfsRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualGitlfsRepository {
		return vs[0].(map[string]*VirtualGitlfsRepository)[vs[1].(string)]
	}).(VirtualGitlfsRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGitlfsRepositoryInput)(nil)).Elem(), &VirtualGitlfsRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGitlfsRepositoryArrayInput)(nil)).Elem(), VirtualGitlfsRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualGitlfsRepositoryMapInput)(nil)).Elem(), VirtualGitlfsRepositoryMap{})
	pulumi.RegisterOutputType(VirtualGitlfsRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualGitlfsRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualGitlfsRepositoryMapOutput{})
}
