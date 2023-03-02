// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a virtual Maven repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Maven+Repository).
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
//			bar, err := artifactory.NewLocalMavenRepository(ctx, "bar", &artifactory.LocalMavenRepositoryArgs{
//				Key:           pulumi.String("bar"),
//				RepoLayoutRef: pulumi.String("maven-2-default"),
//			})
//			if err != nil {
//				return err
//			}
//			baz, err := artifactory.NewRemoteMavenRepository(ctx, "baz", &artifactory.RemoteMavenRepositoryArgs{
//				Key:           pulumi.String("baz"),
//				RepoLayoutRef: pulumi.String("maven-2-default"),
//				Url:           pulumi.String("https://search.maven.com/"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewMavenRepository(ctx, "maven-virt-repo", &artifactory.MavenRepositoryArgs{
//				Description:                          pulumi.String("A test virtual repo"),
//				ExcludesPattern:                      pulumi.String("com/google/**"),
//				ForceMavenAuthentication:             pulumi.Bool(true),
//				IncludesPattern:                      pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				Key:                                  pulumi.String("maven-virt-repo"),
//				Notes:                                pulumi.String("Internal description"),
//				PomRepositoryReferencesCleanupPolicy: pulumi.String("discard_active_reference"),
//				RepoLayoutRef:                        pulumi.String("maven-2-default"),
//				Repositories: pulumi.StringArray{
//					bar.Key,
//					baz.Key,
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
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/mavenRepository:MavenRepository maven-virt-repo maven-virt-repo
//
// ```
type MavenRepository struct {
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
	// Forces authentication when fetching from remote repos.
	ForceMavenAuthentication pulumi.BoolOutput `pulumi:"forceMavenAuthentication"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringOutput `pulumi:"key"`
	// The keypair used to sign artifacts
	KeyPair pulumi.StringPtrOutput `pulumi:"keyPair"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringOutput `pulumi:"packageType"`
	// One of: `"discardActiveReference", "discardAnyReference", "nothing"`
	PomRepositoryReferencesCleanupPolicy pulumi.StringOutput `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
}

// NewMavenRepository registers a new resource with the given unique name, arguments, and options.
func NewMavenRepository(ctx *pulumi.Context,
	name string, args *MavenRepositoryArgs, opts ...pulumi.ResourceOption) (*MavenRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource MavenRepository
	err := ctx.RegisterResource("artifactory:index/mavenRepository:MavenRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMavenRepository gets an existing MavenRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMavenRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MavenRepositoryState, opts ...pulumi.ResourceOption) (*MavenRepository, error) {
	var resource MavenRepository
	err := ctx.ReadResource("artifactory:index/mavenRepository:MavenRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MavenRepository resources.
type mavenRepositoryState struct {
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
	// Forces authentication when fetching from remote repos.
	ForceMavenAuthentication *bool `pulumi:"forceMavenAuthentication"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key *string `pulumi:"key"`
	// The keypair used to sign artifacts
	KeyPair *string `pulumi:"keyPair"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType *string `pulumi:"packageType"`
	// One of: `"discardActiveReference", "discardAnyReference", "nothing"`
	PomRepositoryReferencesCleanupPolicy *string `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
}

type MavenRepositoryState struct {
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
	// Forces authentication when fetching from remote repos.
	ForceMavenAuthentication pulumi.BoolPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringPtrInput
	// The keypair used to sign artifacts
	KeyPair pulumi.StringPtrInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
	PackageType pulumi.StringPtrInput
	// One of: `"discardActiveReference", "discardAnyReference", "nothing"`
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
}

func (MavenRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*mavenRepositoryState)(nil)).Elem()
}

type mavenRepositoryArgs struct {
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
	// Forces authentication when fetching from remote repos.
	ForceMavenAuthentication *bool `pulumi:"forceMavenAuthentication"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key string `pulumi:"key"`
	// The keypair used to sign artifacts
	KeyPair *string `pulumi:"keyPair"`
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes *string `pulumi:"notes"`
	// One of: `"discardActiveReference", "discardAnyReference", "nothing"`
	PomRepositoryReferencesCleanupPolicy *string `pulumi:"pomRepositoryReferencesCleanupPolicy"`
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
}

// The set of arguments for constructing a MavenRepository resource.
type MavenRepositoryArgs struct {
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
	// Forces authentication when fetching from remote repos.
	ForceMavenAuthentication pulumi.BoolPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
	Key pulumi.StringInput
	// The keypair used to sign artifacts
	KeyPair pulumi.StringPtrInput
	// A free text field to add additional notes about the repository. These are only visible to the administrator.
	Notes pulumi.StringPtrInput
	// One of: `"discardActiveReference", "discardAnyReference", "nothing"`
	PomRepositoryReferencesCleanupPolicy pulumi.StringPtrInput
	// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
	// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
	// will remain in the Terraform state, which will create state drift during the update.
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
}

func (MavenRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mavenRepositoryArgs)(nil)).Elem()
}

type MavenRepositoryInput interface {
	pulumi.Input

	ToMavenRepositoryOutput() MavenRepositoryOutput
	ToMavenRepositoryOutputWithContext(ctx context.Context) MavenRepositoryOutput
}

func (*MavenRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**MavenRepository)(nil)).Elem()
}

func (i *MavenRepository) ToMavenRepositoryOutput() MavenRepositoryOutput {
	return i.ToMavenRepositoryOutputWithContext(context.Background())
}

func (i *MavenRepository) ToMavenRepositoryOutputWithContext(ctx context.Context) MavenRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MavenRepositoryOutput)
}

// MavenRepositoryArrayInput is an input type that accepts MavenRepositoryArray and MavenRepositoryArrayOutput values.
// You can construct a concrete instance of `MavenRepositoryArrayInput` via:
//
//	MavenRepositoryArray{ MavenRepositoryArgs{...} }
type MavenRepositoryArrayInput interface {
	pulumi.Input

	ToMavenRepositoryArrayOutput() MavenRepositoryArrayOutput
	ToMavenRepositoryArrayOutputWithContext(context.Context) MavenRepositoryArrayOutput
}

type MavenRepositoryArray []MavenRepositoryInput

func (MavenRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MavenRepository)(nil)).Elem()
}

func (i MavenRepositoryArray) ToMavenRepositoryArrayOutput() MavenRepositoryArrayOutput {
	return i.ToMavenRepositoryArrayOutputWithContext(context.Background())
}

func (i MavenRepositoryArray) ToMavenRepositoryArrayOutputWithContext(ctx context.Context) MavenRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MavenRepositoryArrayOutput)
}

// MavenRepositoryMapInput is an input type that accepts MavenRepositoryMap and MavenRepositoryMapOutput values.
// You can construct a concrete instance of `MavenRepositoryMapInput` via:
//
//	MavenRepositoryMap{ "key": MavenRepositoryArgs{...} }
type MavenRepositoryMapInput interface {
	pulumi.Input

	ToMavenRepositoryMapOutput() MavenRepositoryMapOutput
	ToMavenRepositoryMapOutputWithContext(context.Context) MavenRepositoryMapOutput
}

type MavenRepositoryMap map[string]MavenRepositoryInput

func (MavenRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MavenRepository)(nil)).Elem()
}

func (i MavenRepositoryMap) ToMavenRepositoryMapOutput() MavenRepositoryMapOutput {
	return i.ToMavenRepositoryMapOutputWithContext(context.Background())
}

func (i MavenRepositoryMap) ToMavenRepositoryMapOutputWithContext(ctx context.Context) MavenRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MavenRepositoryMapOutput)
}

type MavenRepositoryOutput struct{ *pulumi.OutputState }

func (MavenRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MavenRepository)(nil)).Elem()
}

func (o MavenRepositoryOutput) ToMavenRepositoryOutput() MavenRepositoryOutput {
	return o
}

func (o MavenRepositoryOutput) ToMavenRepositoryOutputWithContext(ctx context.Context) MavenRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o MavenRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.BoolPtrOutput { return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts }).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o MavenRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// A free text field that describes the content and purpose of the repository. If you choose to insert a link into this
// field, clicking the link will prompt the user to confirm that they might be redirected to a new domain.
func (o MavenRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o MavenRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// Forces authentication when fetching from remote repos.
func (o MavenRepositoryOutput) ForceMavenAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.BoolOutput { return v.ForceMavenAuthentication }).(pulumi.BoolOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o MavenRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o MavenRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The keypair used to sign artifacts
func (o MavenRepositoryOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringPtrOutput { return v.KeyPair }).(pulumi.StringPtrOutput)
}

// A free text field to add additional notes about the repository. These are only visible to the administrator.
func (o MavenRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The Package Type. This must be specified when the repository is created, and once set, cannot be changed.
func (o MavenRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// One of: `"discardActiveReference", "discardAnyReference", "nothing"`
func (o MavenRepositoryOutput) PomRepositoryReferencesCleanupPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringOutput { return v.PomRepositoryReferencesCleanupPolicy }).(pulumi.StringOutput)
}

// Project environment for assigning this repository to. Allow values: "DEV" or "PROD". The attribute should only be used
// if the repository is already assigned to the existing project. If not, the attribute will be ignored by Artifactory, but
// will remain in the Terraform state, which will create state drift during the update.
func (o MavenRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 20 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o MavenRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the local repository
func (o MavenRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o MavenRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MavenRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

type MavenRepositoryArrayOutput struct{ *pulumi.OutputState }

func (MavenRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MavenRepository)(nil)).Elem()
}

func (o MavenRepositoryArrayOutput) ToMavenRepositoryArrayOutput() MavenRepositoryArrayOutput {
	return o
}

func (o MavenRepositoryArrayOutput) ToMavenRepositoryArrayOutputWithContext(ctx context.Context) MavenRepositoryArrayOutput {
	return o
}

func (o MavenRepositoryArrayOutput) Index(i pulumi.IntInput) MavenRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MavenRepository {
		return vs[0].([]*MavenRepository)[vs[1].(int)]
	}).(MavenRepositoryOutput)
}

type MavenRepositoryMapOutput struct{ *pulumi.OutputState }

func (MavenRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MavenRepository)(nil)).Elem()
}

func (o MavenRepositoryMapOutput) ToMavenRepositoryMapOutput() MavenRepositoryMapOutput {
	return o
}

func (o MavenRepositoryMapOutput) ToMavenRepositoryMapOutputWithContext(ctx context.Context) MavenRepositoryMapOutput {
	return o
}

func (o MavenRepositoryMapOutput) MapIndex(k pulumi.StringInput) MavenRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MavenRepository {
		return vs[0].(map[string]*MavenRepository)[vs[1].(string)]
	}).(MavenRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MavenRepositoryInput)(nil)).Elem(), &MavenRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*MavenRepositoryArrayInput)(nil)).Elem(), MavenRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MavenRepositoryMapInput)(nil)).Elem(), MavenRepositoryMap{})
	pulumi.RegisterOutputType(MavenRepositoryOutput{})
	pulumi.RegisterOutputType(MavenRepositoryArrayOutput{})
	pulumi.RegisterOutputType(MavenRepositoryMapOutput{})
}
