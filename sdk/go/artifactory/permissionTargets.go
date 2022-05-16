// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Artifactory permission target resource. This can be used to create and manage Artifactory permission targets.
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
// 		_, err := artifactory.NewPermissionTarget(ctx, "test-perm", &artifactory.PermissionTargetArgs{
// 			Build: &PermissionTargetBuildArgs{
// 				Actions: &PermissionTargetBuildActionsArgs{
// 					Users: PermissionTargetBuildActionsUserArray{
// 						&PermissionTargetBuildActionsUserArgs{
// 							Name: pulumi.String("anonymous"),
// 							Permissions: pulumi.StringArray{
// 								pulumi.String("read"),
// 								pulumi.String("write"),
// 							},
// 						},
// 					},
// 				},
// 				IncludesPatterns: pulumi.StringArray{
// 					pulumi.String("**"),
// 				},
// 				Repositories: pulumi.StringArray{
// 					pulumi.String("artifactory-build-info"),
// 				},
// 			},
// 			ReleaseBundle: &PermissionTargetReleaseBundleArgs{
// 				Actions: &PermissionTargetReleaseBundleActionsArgs{
// 					Users: PermissionTargetReleaseBundleActionsUserArray{
// 						&PermissionTargetReleaseBundleActionsUserArgs{
// 							Name: pulumi.String("anonymous"),
// 							Permissions: pulumi.StringArray{
// 								pulumi.String("read"),
// 							},
// 						},
// 					},
// 				},
// 				IncludesPatterns: pulumi.StringArray{
// 					pulumi.String("**"),
// 				},
// 				Repositories: pulumi.StringArray{
// 					pulumi.String("release-bundles"),
// 				},
// 			},
// 			Repo: &PermissionTargetRepoArgs{
// 				Actions: &PermissionTargetRepoActionsArgs{
// 					Groups: PermissionTargetRepoActionsGroupArray{
// 						&PermissionTargetRepoActionsGroupArgs{
// 							Name: pulumi.String("readers"),
// 							Permissions: pulumi.StringArray{
// 								pulumi.String("read"),
// 							},
// 						},
// 					},
// 					Users: PermissionTargetRepoActionsUserArray{
// 						&PermissionTargetRepoActionsUserArgs{
// 							Name: pulumi.String("anonymous"),
// 							Permissions: pulumi.StringArray{
// 								pulumi.String("read"),
// 								pulumi.String("write"),
// 							},
// 						},
// 					},
// 				},
// 				ExcludesPatterns: pulumi.StringArray{
// 					pulumi.String("bar/**"),
// 				},
// 				IncludesPatterns: pulumi.StringArray{
// 					pulumi.String("foo/**"),
// 				},
// 				Repositories: pulumi.StringArray{
// 					pulumi.String("example-repo-local"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Permissions
//
// The provider supports the following `permission` enums:
//
// * `read`
// * `write`
// * `annotate`
// * `delete`
// * `manage`
//
// The values can be mapped to the permissions from the official [documentation](https://www.jfrog.com/confluence/display/JFROG/Permissions):
//
// * `read` - matches `Read` permissions.
// * `write` - matches `  Deploy / Cache / Create ` permissions.
// * `annotate` - matches `Annotate` permissions.
// * `delete` - matches `Delete / Overwrite` permissions.
// * `manage` - matches `Manage` permissions.
//
// ## Import
//
// Permission targets can be imported using their name, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/permissionTargets:PermissionTargets terraform-test-permission mypermission
// ```
type PermissionTargets struct {
	pulumi.CustomResourceState

	// As for repo but for artifactory-build-info permssions.
	Build PermissionTargetsBuildPtrOutput `pulumi:"build"`
	// Name of permission.
	Name pulumi.StringOutput `pulumi:"name"`
	// As for repo for for release-bundles permissions.
	ReleaseBundle PermissionTargetsReleaseBundlePtrOutput `pulumi:"releaseBundle"`
	// Repository permission configuration.
	Repo PermissionTargetsRepoPtrOutput `pulumi:"repo"`
}

// NewPermissionTargets registers a new resource with the given unique name, arguments, and options.
func NewPermissionTargets(ctx *pulumi.Context,
	name string, args *PermissionTargetsArgs, opts ...pulumi.ResourceOption) (*PermissionTargets, error) {
	if args == nil {
		args = &PermissionTargetsArgs{}
	}

	var resource PermissionTargets
	err := ctx.RegisterResource("artifactory:index/permissionTargets:PermissionTargets", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPermissionTargets gets an existing PermissionTargets resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPermissionTargets(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PermissionTargetsState, opts ...pulumi.ResourceOption) (*PermissionTargets, error) {
	var resource PermissionTargets
	err := ctx.ReadResource("artifactory:index/permissionTargets:PermissionTargets", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PermissionTargets resources.
type permissionTargetsState struct {
	// As for repo but for artifactory-build-info permssions.
	Build *PermissionTargetsBuild `pulumi:"build"`
	// Name of permission.
	Name *string `pulumi:"name"`
	// As for repo for for release-bundles permissions.
	ReleaseBundle *PermissionTargetsReleaseBundle `pulumi:"releaseBundle"`
	// Repository permission configuration.
	Repo *PermissionTargetsRepo `pulumi:"repo"`
}

type PermissionTargetsState struct {
	// As for repo but for artifactory-build-info permssions.
	Build PermissionTargetsBuildPtrInput
	// Name of permission.
	Name pulumi.StringPtrInput
	// As for repo for for release-bundles permissions.
	ReleaseBundle PermissionTargetsReleaseBundlePtrInput
	// Repository permission configuration.
	Repo PermissionTargetsRepoPtrInput
}

func (PermissionTargetsState) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionTargetsState)(nil)).Elem()
}

type permissionTargetsArgs struct {
	// As for repo but for artifactory-build-info permssions.
	Build *PermissionTargetsBuild `pulumi:"build"`
	// Name of permission.
	Name *string `pulumi:"name"`
	// As for repo for for release-bundles permissions.
	ReleaseBundle *PermissionTargetsReleaseBundle `pulumi:"releaseBundle"`
	// Repository permission configuration.
	Repo *PermissionTargetsRepo `pulumi:"repo"`
}

// The set of arguments for constructing a PermissionTargets resource.
type PermissionTargetsArgs struct {
	// As for repo but for artifactory-build-info permssions.
	Build PermissionTargetsBuildPtrInput
	// Name of permission.
	Name pulumi.StringPtrInput
	// As for repo for for release-bundles permissions.
	ReleaseBundle PermissionTargetsReleaseBundlePtrInput
	// Repository permission configuration.
	Repo PermissionTargetsRepoPtrInput
}

func (PermissionTargetsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*permissionTargetsArgs)(nil)).Elem()
}

type PermissionTargetsInput interface {
	pulumi.Input

	ToPermissionTargetsOutput() PermissionTargetsOutput
	ToPermissionTargetsOutputWithContext(ctx context.Context) PermissionTargetsOutput
}

func (*PermissionTargets) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionTargets)(nil)).Elem()
}

func (i *PermissionTargets) ToPermissionTargetsOutput() PermissionTargetsOutput {
	return i.ToPermissionTargetsOutputWithContext(context.Background())
}

func (i *PermissionTargets) ToPermissionTargetsOutputWithContext(ctx context.Context) PermissionTargetsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionTargetsOutput)
}

// PermissionTargetsArrayInput is an input type that accepts PermissionTargetsArray and PermissionTargetsArrayOutput values.
// You can construct a concrete instance of `PermissionTargetsArrayInput` via:
//
//          PermissionTargetsArray{ PermissionTargetsArgs{...} }
type PermissionTargetsArrayInput interface {
	pulumi.Input

	ToPermissionTargetsArrayOutput() PermissionTargetsArrayOutput
	ToPermissionTargetsArrayOutputWithContext(context.Context) PermissionTargetsArrayOutput
}

type PermissionTargetsArray []PermissionTargetsInput

func (PermissionTargetsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PermissionTargets)(nil)).Elem()
}

func (i PermissionTargetsArray) ToPermissionTargetsArrayOutput() PermissionTargetsArrayOutput {
	return i.ToPermissionTargetsArrayOutputWithContext(context.Background())
}

func (i PermissionTargetsArray) ToPermissionTargetsArrayOutputWithContext(ctx context.Context) PermissionTargetsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionTargetsArrayOutput)
}

// PermissionTargetsMapInput is an input type that accepts PermissionTargetsMap and PermissionTargetsMapOutput values.
// You can construct a concrete instance of `PermissionTargetsMapInput` via:
//
//          PermissionTargetsMap{ "key": PermissionTargetsArgs{...} }
type PermissionTargetsMapInput interface {
	pulumi.Input

	ToPermissionTargetsMapOutput() PermissionTargetsMapOutput
	ToPermissionTargetsMapOutputWithContext(context.Context) PermissionTargetsMapOutput
}

type PermissionTargetsMap map[string]PermissionTargetsInput

func (PermissionTargetsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PermissionTargets)(nil)).Elem()
}

func (i PermissionTargetsMap) ToPermissionTargetsMapOutput() PermissionTargetsMapOutput {
	return i.ToPermissionTargetsMapOutputWithContext(context.Background())
}

func (i PermissionTargetsMap) ToPermissionTargetsMapOutputWithContext(ctx context.Context) PermissionTargetsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionTargetsMapOutput)
}

type PermissionTargetsOutput struct{ *pulumi.OutputState }

func (PermissionTargetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionTargets)(nil)).Elem()
}

func (o PermissionTargetsOutput) ToPermissionTargetsOutput() PermissionTargetsOutput {
	return o
}

func (o PermissionTargetsOutput) ToPermissionTargetsOutputWithContext(ctx context.Context) PermissionTargetsOutput {
	return o
}

// As for repo but for artifactory-build-info permssions.
func (o PermissionTargetsOutput) Build() PermissionTargetsBuildPtrOutput {
	return o.ApplyT(func(v *PermissionTargets) PermissionTargetsBuildPtrOutput { return v.Build }).(PermissionTargetsBuildPtrOutput)
}

// Name of permission.
func (o PermissionTargetsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PermissionTargets) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// As for repo for for release-bundles permissions.
func (o PermissionTargetsOutput) ReleaseBundle() PermissionTargetsReleaseBundlePtrOutput {
	return o.ApplyT(func(v *PermissionTargets) PermissionTargetsReleaseBundlePtrOutput { return v.ReleaseBundle }).(PermissionTargetsReleaseBundlePtrOutput)
}

// Repository permission configuration.
func (o PermissionTargetsOutput) Repo() PermissionTargetsRepoPtrOutput {
	return o.ApplyT(func(v *PermissionTargets) PermissionTargetsRepoPtrOutput { return v.Repo }).(PermissionTargetsRepoPtrOutput)
}

type PermissionTargetsArrayOutput struct{ *pulumi.OutputState }

func (PermissionTargetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PermissionTargets)(nil)).Elem()
}

func (o PermissionTargetsArrayOutput) ToPermissionTargetsArrayOutput() PermissionTargetsArrayOutput {
	return o
}

func (o PermissionTargetsArrayOutput) ToPermissionTargetsArrayOutputWithContext(ctx context.Context) PermissionTargetsArrayOutput {
	return o
}

func (o PermissionTargetsArrayOutput) Index(i pulumi.IntInput) PermissionTargetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PermissionTargets {
		return vs[0].([]*PermissionTargets)[vs[1].(int)]
	}).(PermissionTargetsOutput)
}

type PermissionTargetsMapOutput struct{ *pulumi.OutputState }

func (PermissionTargetsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PermissionTargets)(nil)).Elem()
}

func (o PermissionTargetsMapOutput) ToPermissionTargetsMapOutput() PermissionTargetsMapOutput {
	return o
}

func (o PermissionTargetsMapOutput) ToPermissionTargetsMapOutputWithContext(ctx context.Context) PermissionTargetsMapOutput {
	return o
}

func (o PermissionTargetsMapOutput) MapIndex(k pulumi.StringInput) PermissionTargetsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PermissionTargets {
		return vs[0].(map[string]*PermissionTargets)[vs[1].(string)]
	}).(PermissionTargetsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionTargetsInput)(nil)).Elem(), &PermissionTargets{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionTargetsArrayInput)(nil)).Elem(), PermissionTargetsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PermissionTargetsMapInput)(nil)).Elem(), PermissionTargetsMap{})
	pulumi.RegisterOutputType(PermissionTargetsOutput{})
	pulumi.RegisterOutputType(PermissionTargetsArrayOutput{})
	pulumi.RegisterOutputType(PermissionTargetsMapOutput{})
}
