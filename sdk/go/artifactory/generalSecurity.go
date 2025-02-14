// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can be used to manage Artifactory's general security settings.
//
// Only a single `GeneralSecurity` resource is meant to be defined.
//
// ~>The `GeneralSecurity` resource utilizes endpoints which are blocked/removed in SaaS environments (i.e. in Artifactory online), rendering this resource incompatible with Artifactory SaaS environments.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewGeneralSecurity(ctx, "security", &artifactory.GeneralSecurityArgs{
//				EnableAnonymousAccess: pulumi.Bool(true),
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
// Current general security settings can be imported using `security` as the `ID`, e.g.
//
// ```sh
//
//	$ pulumi import artifactory:index/generalSecurity:GeneralSecurity security security
//
// ```
//
//	~>The `artifactory_general_security` resource uses endpoints that are undocumented and may not work with SaaS environments, or may change without notice.
type GeneralSecurity struct {
	pulumi.CustomResourceState

	EnableAnonymousAccess pulumi.BoolPtrOutput `pulumi:"enableAnonymousAccess"`
}

// NewGeneralSecurity registers a new resource with the given unique name, arguments, and options.
func NewGeneralSecurity(ctx *pulumi.Context,
	name string, args *GeneralSecurityArgs, opts ...pulumi.ResourceOption) (*GeneralSecurity, error) {
	if args == nil {
		args = &GeneralSecurityArgs{}
	}

	var resource GeneralSecurity
	err := ctx.RegisterResource("artifactory:index/generalSecurity:GeneralSecurity", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeneralSecurity gets an existing GeneralSecurity resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeneralSecurity(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeneralSecurityState, opts ...pulumi.ResourceOption) (*GeneralSecurity, error) {
	var resource GeneralSecurity
	err := ctx.ReadResource("artifactory:index/generalSecurity:GeneralSecurity", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GeneralSecurity resources.
type generalSecurityState struct {
	EnableAnonymousAccess *bool `pulumi:"enableAnonymousAccess"`
}

type GeneralSecurityState struct {
	EnableAnonymousAccess pulumi.BoolPtrInput
}

func (GeneralSecurityState) ElementType() reflect.Type {
	return reflect.TypeOf((*generalSecurityState)(nil)).Elem()
}

type generalSecurityArgs struct {
	EnableAnonymousAccess *bool `pulumi:"enableAnonymousAccess"`
}

// The set of arguments for constructing a GeneralSecurity resource.
type GeneralSecurityArgs struct {
	EnableAnonymousAccess pulumi.BoolPtrInput
}

func (GeneralSecurityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*generalSecurityArgs)(nil)).Elem()
}

type GeneralSecurityInput interface {
	pulumi.Input

	ToGeneralSecurityOutput() GeneralSecurityOutput
	ToGeneralSecurityOutputWithContext(ctx context.Context) GeneralSecurityOutput
}

func (*GeneralSecurity) ElementType() reflect.Type {
	return reflect.TypeOf((**GeneralSecurity)(nil)).Elem()
}

func (i *GeneralSecurity) ToGeneralSecurityOutput() GeneralSecurityOutput {
	return i.ToGeneralSecurityOutputWithContext(context.Background())
}

func (i *GeneralSecurity) ToGeneralSecurityOutputWithContext(ctx context.Context) GeneralSecurityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneralSecurityOutput)
}

// GeneralSecurityArrayInput is an input type that accepts GeneralSecurityArray and GeneralSecurityArrayOutput values.
// You can construct a concrete instance of `GeneralSecurityArrayInput` via:
//
//	GeneralSecurityArray{ GeneralSecurityArgs{...} }
type GeneralSecurityArrayInput interface {
	pulumi.Input

	ToGeneralSecurityArrayOutput() GeneralSecurityArrayOutput
	ToGeneralSecurityArrayOutputWithContext(context.Context) GeneralSecurityArrayOutput
}

type GeneralSecurityArray []GeneralSecurityInput

func (GeneralSecurityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GeneralSecurity)(nil)).Elem()
}

func (i GeneralSecurityArray) ToGeneralSecurityArrayOutput() GeneralSecurityArrayOutput {
	return i.ToGeneralSecurityArrayOutputWithContext(context.Background())
}

func (i GeneralSecurityArray) ToGeneralSecurityArrayOutputWithContext(ctx context.Context) GeneralSecurityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneralSecurityArrayOutput)
}

// GeneralSecurityMapInput is an input type that accepts GeneralSecurityMap and GeneralSecurityMapOutput values.
// You can construct a concrete instance of `GeneralSecurityMapInput` via:
//
//	GeneralSecurityMap{ "key": GeneralSecurityArgs{...} }
type GeneralSecurityMapInput interface {
	pulumi.Input

	ToGeneralSecurityMapOutput() GeneralSecurityMapOutput
	ToGeneralSecurityMapOutputWithContext(context.Context) GeneralSecurityMapOutput
}

type GeneralSecurityMap map[string]GeneralSecurityInput

func (GeneralSecurityMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GeneralSecurity)(nil)).Elem()
}

func (i GeneralSecurityMap) ToGeneralSecurityMapOutput() GeneralSecurityMapOutput {
	return i.ToGeneralSecurityMapOutputWithContext(context.Background())
}

func (i GeneralSecurityMap) ToGeneralSecurityMapOutputWithContext(ctx context.Context) GeneralSecurityMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneralSecurityMapOutput)
}

type GeneralSecurityOutput struct{ *pulumi.OutputState }

func (GeneralSecurityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeneralSecurity)(nil)).Elem()
}

func (o GeneralSecurityOutput) ToGeneralSecurityOutput() GeneralSecurityOutput {
	return o
}

func (o GeneralSecurityOutput) ToGeneralSecurityOutputWithContext(ctx context.Context) GeneralSecurityOutput {
	return o
}

func (o GeneralSecurityOutput) EnableAnonymousAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GeneralSecurity) pulumi.BoolPtrOutput { return v.EnableAnonymousAccess }).(pulumi.BoolPtrOutput)
}

type GeneralSecurityArrayOutput struct{ *pulumi.OutputState }

func (GeneralSecurityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GeneralSecurity)(nil)).Elem()
}

func (o GeneralSecurityArrayOutput) ToGeneralSecurityArrayOutput() GeneralSecurityArrayOutput {
	return o
}

func (o GeneralSecurityArrayOutput) ToGeneralSecurityArrayOutputWithContext(ctx context.Context) GeneralSecurityArrayOutput {
	return o
}

func (o GeneralSecurityArrayOutput) Index(i pulumi.IntInput) GeneralSecurityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GeneralSecurity {
		return vs[0].([]*GeneralSecurity)[vs[1].(int)]
	}).(GeneralSecurityOutput)
}

type GeneralSecurityMapOutput struct{ *pulumi.OutputState }

func (GeneralSecurityMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GeneralSecurity)(nil)).Elem()
}

func (o GeneralSecurityMapOutput) ToGeneralSecurityMapOutput() GeneralSecurityMapOutput {
	return o
}

func (o GeneralSecurityMapOutput) ToGeneralSecurityMapOutputWithContext(ctx context.Context) GeneralSecurityMapOutput {
	return o
}

func (o GeneralSecurityMapOutput) MapIndex(k pulumi.StringInput) GeneralSecurityOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GeneralSecurity {
		return vs[0].(map[string]*GeneralSecurity)[vs[1].(string)]
	}).(GeneralSecurityOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GeneralSecurityInput)(nil)).Elem(), &GeneralSecurity{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeneralSecurityArrayInput)(nil)).Elem(), GeneralSecurityArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeneralSecurityMapInput)(nil)).Elem(), GeneralSecurityMap{})
	pulumi.RegisterOutputType(GeneralSecurityOutput{})
	pulumi.RegisterOutputType(GeneralSecurityArrayOutput{})
	pulumi.RegisterOutputType(GeneralSecurityMapOutput{})
}
