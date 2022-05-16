// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can be used to manage Artifactory's LDAP Group settings for user authentication.
//
// LDAP Groups Add-on allows you to synchronize your LDAP groups with the system and leverage your existing organizational
// structure for managing group-based permissions.
//
// ## Import
//
// LDAP Group setting can be imported using the key, e.g.
//
// ```sh
//  $ pulumi import artifactory:index/ldapGroupSetting:LdapGroupSetting ldap_group_name ldap_group_name
// ```
type LdapGroupSetting struct {
	pulumi.CustomResourceState

	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute pulumi.StringOutput `pulumi:"descriptionAttribute"`
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn pulumi.StringPtrOutput `pulumi:"groupBaseDn"`
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
	GroupMemberAttribute pulumi.StringOutput `pulumi:"groupMemberAttribute"`
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute pulumi.StringOutput `pulumi:"groupNameAttribute"`
	// The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of
	// the ldap group setting XML block of system configuration.
	LdapSettingKey pulumi.StringOutput `pulumi:"ldapSettingKey"`
	// Ldap group setting name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
	// - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
	// - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
	// - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
	Strategy pulumi.StringOutput `pulumi:"strategy"`
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
	SubTree pulumi.BoolPtrOutput `pulumi:"subTree"`
}

// NewLdapGroupSetting registers a new resource with the given unique name, arguments, and options.
func NewLdapGroupSetting(ctx *pulumi.Context,
	name string, args *LdapGroupSettingArgs, opts ...pulumi.ResourceOption) (*LdapGroupSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DescriptionAttribute == nil {
		return nil, errors.New("invalid value for required argument 'DescriptionAttribute'")
	}
	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	if args.GroupMemberAttribute == nil {
		return nil, errors.New("invalid value for required argument 'GroupMemberAttribute'")
	}
	if args.GroupNameAttribute == nil {
		return nil, errors.New("invalid value for required argument 'GroupNameAttribute'")
	}
	if args.LdapSettingKey == nil {
		return nil, errors.New("invalid value for required argument 'LdapSettingKey'")
	}
	if args.Strategy == nil {
		return nil, errors.New("invalid value for required argument 'Strategy'")
	}
	var resource LdapGroupSetting
	err := ctx.RegisterResource("artifactory:index/ldapGroupSetting:LdapGroupSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLdapGroupSetting gets an existing LdapGroupSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLdapGroupSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LdapGroupSettingState, opts ...pulumi.ResourceOption) (*LdapGroupSetting, error) {
	var resource LdapGroupSetting
	err := ctx.ReadResource("artifactory:index/ldapGroupSetting:LdapGroupSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LdapGroupSetting resources.
type ldapGroupSettingState struct {
	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute *string `pulumi:"descriptionAttribute"`
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter *string `pulumi:"filter"`
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn *string `pulumi:"groupBaseDn"`
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
	GroupMemberAttribute *string `pulumi:"groupMemberAttribute"`
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute *string `pulumi:"groupNameAttribute"`
	// The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of
	// the ldap group setting XML block of system configuration.
	LdapSettingKey *string `pulumi:"ldapSettingKey"`
	// Ldap group setting name.
	Name *string `pulumi:"name"`
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
	// - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
	// - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
	// - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
	Strategy *string `pulumi:"strategy"`
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
	SubTree *bool `pulumi:"subTree"`
}

type LdapGroupSettingState struct {
	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute pulumi.StringPtrInput
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter pulumi.StringPtrInput
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn pulumi.StringPtrInput
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
	GroupMemberAttribute pulumi.StringPtrInput
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute pulumi.StringPtrInput
	// The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of
	// the ldap group setting XML block of system configuration.
	LdapSettingKey pulumi.StringPtrInput
	// Ldap group setting name.
	Name pulumi.StringPtrInput
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
	// - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
	// - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
	// - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
	Strategy pulumi.StringPtrInput
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
	SubTree pulumi.BoolPtrInput
}

func (LdapGroupSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*ldapGroupSettingState)(nil)).Elem()
}

type ldapGroupSettingArgs struct {
	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute string `pulumi:"descriptionAttribute"`
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter string `pulumi:"filter"`
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn *string `pulumi:"groupBaseDn"`
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
	GroupMemberAttribute string `pulumi:"groupMemberAttribute"`
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute string `pulumi:"groupNameAttribute"`
	// The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of
	// the ldap group setting XML block of system configuration.
	LdapSettingKey string `pulumi:"ldapSettingKey"`
	// Ldap group setting name.
	Name *string `pulumi:"name"`
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
	// - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
	// - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
	// - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
	Strategy string `pulumi:"strategy"`
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
	SubTree *bool `pulumi:"subTree"`
}

// The set of arguments for constructing a LdapGroupSetting resource.
type LdapGroupSettingArgs struct {
	// An attribute on the group entry which denoting the group description. Used when importing groups.
	DescriptionAttribute pulumi.StringInput
	// The LDAP filter used to search for group entries. Used for importing groups.
	Filter pulumi.StringInput
	// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
	GroupBaseDn pulumi.StringPtrInput
	// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
	GroupMemberAttribute pulumi.StringInput
	// Attribute on the group entry denoting the group name. Used when importing groups.
	GroupNameAttribute pulumi.StringInput
	// The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of
	// the ldap group setting XML block of system configuration.
	LdapSettingKey pulumi.StringInput
	// Ldap group setting name.
	Name pulumi.StringPtrInput
	// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
	// - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
	// - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
	// - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
	Strategy pulumi.StringInput
	// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
	SubTree pulumi.BoolPtrInput
}

func (LdapGroupSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ldapGroupSettingArgs)(nil)).Elem()
}

type LdapGroupSettingInput interface {
	pulumi.Input

	ToLdapGroupSettingOutput() LdapGroupSettingOutput
	ToLdapGroupSettingOutputWithContext(ctx context.Context) LdapGroupSettingOutput
}

func (*LdapGroupSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapGroupSetting)(nil)).Elem()
}

func (i *LdapGroupSetting) ToLdapGroupSettingOutput() LdapGroupSettingOutput {
	return i.ToLdapGroupSettingOutputWithContext(context.Background())
}

func (i *LdapGroupSetting) ToLdapGroupSettingOutputWithContext(ctx context.Context) LdapGroupSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapGroupSettingOutput)
}

// LdapGroupSettingArrayInput is an input type that accepts LdapGroupSettingArray and LdapGroupSettingArrayOutput values.
// You can construct a concrete instance of `LdapGroupSettingArrayInput` via:
//
//          LdapGroupSettingArray{ LdapGroupSettingArgs{...} }
type LdapGroupSettingArrayInput interface {
	pulumi.Input

	ToLdapGroupSettingArrayOutput() LdapGroupSettingArrayOutput
	ToLdapGroupSettingArrayOutputWithContext(context.Context) LdapGroupSettingArrayOutput
}

type LdapGroupSettingArray []LdapGroupSettingInput

func (LdapGroupSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LdapGroupSetting)(nil)).Elem()
}

func (i LdapGroupSettingArray) ToLdapGroupSettingArrayOutput() LdapGroupSettingArrayOutput {
	return i.ToLdapGroupSettingArrayOutputWithContext(context.Background())
}

func (i LdapGroupSettingArray) ToLdapGroupSettingArrayOutputWithContext(ctx context.Context) LdapGroupSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapGroupSettingArrayOutput)
}

// LdapGroupSettingMapInput is an input type that accepts LdapGroupSettingMap and LdapGroupSettingMapOutput values.
// You can construct a concrete instance of `LdapGroupSettingMapInput` via:
//
//          LdapGroupSettingMap{ "key": LdapGroupSettingArgs{...} }
type LdapGroupSettingMapInput interface {
	pulumi.Input

	ToLdapGroupSettingMapOutput() LdapGroupSettingMapOutput
	ToLdapGroupSettingMapOutputWithContext(context.Context) LdapGroupSettingMapOutput
}

type LdapGroupSettingMap map[string]LdapGroupSettingInput

func (LdapGroupSettingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LdapGroupSetting)(nil)).Elem()
}

func (i LdapGroupSettingMap) ToLdapGroupSettingMapOutput() LdapGroupSettingMapOutput {
	return i.ToLdapGroupSettingMapOutputWithContext(context.Background())
}

func (i LdapGroupSettingMap) ToLdapGroupSettingMapOutputWithContext(ctx context.Context) LdapGroupSettingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LdapGroupSettingMapOutput)
}

type LdapGroupSettingOutput struct{ *pulumi.OutputState }

func (LdapGroupSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LdapGroupSetting)(nil)).Elem()
}

func (o LdapGroupSettingOutput) ToLdapGroupSettingOutput() LdapGroupSettingOutput {
	return o
}

func (o LdapGroupSettingOutput) ToLdapGroupSettingOutputWithContext(ctx context.Context) LdapGroupSettingOutput {
	return o
}

// An attribute on the group entry which denoting the group description. Used when importing groups.
func (o LdapGroupSettingOutput) DescriptionAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSetting) pulumi.StringOutput { return v.DescriptionAttribute }).(pulumi.StringOutput)
}

// The LDAP filter used to search for group entries. Used for importing groups.
func (o LdapGroupSettingOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSetting) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// A search base for group entry DNs, relative to the DN on the LDAP server’s URL (and not relative to the LDAP Setting’s “Search Base”). Used when importing groups.
func (o LdapGroupSettingOutput) GroupBaseDn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LdapGroupSetting) pulumi.StringPtrOutput { return v.GroupBaseDn }).(pulumi.StringPtrOutput)
}

// A multi-value attribute on the group entry containing user DNs or IDs of the group members (e.g., uniqueMember,member).
func (o LdapGroupSettingOutput) GroupMemberAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSetting) pulumi.StringOutput { return v.GroupMemberAttribute }).(pulumi.StringOutput)
}

// Attribute on the group entry denoting the group name. Used when importing groups.
func (o LdapGroupSettingOutput) GroupNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSetting) pulumi.StringOutput { return v.GroupNameAttribute }).(pulumi.StringOutput)
}

// The LDAP setting key you want to use for group retrieval. The value for this field corresponds to 'enabledLdap' field of
// the ldap group setting XML block of system configuration.
func (o LdapGroupSettingOutput) LdapSettingKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSetting) pulumi.StringOutput { return v.LdapSettingKey }).(pulumi.StringOutput)
}

// Ldap group setting name.
func (o LdapGroupSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The JFrog Platform Deployment (JPD) supports three ways of mapping groups to LDAP schemas:
// - STATIC: Group objects are aware of their members, however, the users are not aware of the groups they belong to. Each group object such as groupOfNames or groupOfUniqueNames holds its respective member attributes, typically member or uniqueMember, which is a user DN.
// - DYNAMIC: User objects are aware of what groups they belong to, but the group objects are not aware of their members. Each user object contains a custom attribute, such as group, that holds the group DNs or group names of which the user is a member.
// - HIERARCHICAL: The user's DN is indicative of the groups the user belongs to by using group names as part of user DN hierarchy. Each user DN contains a list of ou's or custom attributes that make up the group association. For example, uid=user1,ou=developers,ou=uk,dc=jfrog,dc=org indicates that user1 belongs to two groups: uk and developers.
func (o LdapGroupSettingOutput) Strategy() pulumi.StringOutput {
	return o.ApplyT(func(v *LdapGroupSetting) pulumi.StringOutput { return v.Strategy }).(pulumi.StringOutput)
}

// When set, enables deep search through the sub-tree of the LDAP URL + Search Base. True by default.
func (o LdapGroupSettingOutput) SubTree() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LdapGroupSetting) pulumi.BoolPtrOutput { return v.SubTree }).(pulumi.BoolPtrOutput)
}

type LdapGroupSettingArrayOutput struct{ *pulumi.OutputState }

func (LdapGroupSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LdapGroupSetting)(nil)).Elem()
}

func (o LdapGroupSettingArrayOutput) ToLdapGroupSettingArrayOutput() LdapGroupSettingArrayOutput {
	return o
}

func (o LdapGroupSettingArrayOutput) ToLdapGroupSettingArrayOutputWithContext(ctx context.Context) LdapGroupSettingArrayOutput {
	return o
}

func (o LdapGroupSettingArrayOutput) Index(i pulumi.IntInput) LdapGroupSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LdapGroupSetting {
		return vs[0].([]*LdapGroupSetting)[vs[1].(int)]
	}).(LdapGroupSettingOutput)
}

type LdapGroupSettingMapOutput struct{ *pulumi.OutputState }

func (LdapGroupSettingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LdapGroupSetting)(nil)).Elem()
}

func (o LdapGroupSettingMapOutput) ToLdapGroupSettingMapOutput() LdapGroupSettingMapOutput {
	return o
}

func (o LdapGroupSettingMapOutput) ToLdapGroupSettingMapOutputWithContext(ctx context.Context) LdapGroupSettingMapOutput {
	return o
}

func (o LdapGroupSettingMapOutput) MapIndex(k pulumi.StringInput) LdapGroupSettingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LdapGroupSetting {
		return vs[0].(map[string]*LdapGroupSetting)[vs[1].(string)]
	}).(LdapGroupSettingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LdapGroupSettingInput)(nil)).Elem(), &LdapGroupSetting{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdapGroupSettingArrayInput)(nil)).Elem(), LdapGroupSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LdapGroupSettingMapInput)(nil)).Elem(), LdapGroupSettingMap{})
	pulumi.RegisterOutputType(LdapGroupSettingOutput{})
	pulumi.RegisterOutputType(LdapGroupSettingArrayOutput{})
	pulumi.RegisterOutputType(LdapGroupSettingMapOutput{})
}
