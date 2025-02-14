// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLocalPypiRepository(ctx *pulumi.Context, args *LookupLocalPypiRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalPypiRepositoryResult, error) {
	var rv LookupLocalPypiRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalPypiRepository:getLocalPypiRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalPypiRepository.
type LookupLocalPypiRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool    `pulumi:"blackedOut"`
	CdnRedirect            *bool    `pulumi:"cdnRedirect"`
	Description            *string  `pulumi:"description"`
	DownloadDirect         *bool    `pulumi:"downloadDirect"`
	ExcludesPattern        *string  `pulumi:"excludesPattern"`
	IncludesPattern        *string  `pulumi:"includesPattern"`
	Key                    string   `pulumi:"key"`
	Notes                  *string  `pulumi:"notes"`
	PriorityResolution     *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments    []string `pulumi:"projectEnvironments"`
	ProjectKey             *string  `pulumi:"projectKey"`
	PropertySets           []string `pulumi:"propertySets"`
	RepoLayoutRef          *string  `pulumi:"repoLayoutRef"`
	XrayIndex              *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalPypiRepository.
type LookupLocalPypiRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string   `pulumi:"id"`
	IncludesPattern     string   `pulumi:"includesPattern"`
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
}

func LookupLocalPypiRepositoryOutput(ctx *pulumi.Context, args LookupLocalPypiRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalPypiRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalPypiRepositoryResult, error) {
			args := v.(LookupLocalPypiRepositoryArgs)
			r, err := LookupLocalPypiRepository(ctx, &args, opts...)
			var s LookupLocalPypiRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalPypiRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalPypiRepository.
type LookupLocalPypiRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput     `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput     `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput     `pulumi:"cdnRedirect"`
	Description            pulumi.StringPtrInput   `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput     `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput   `pulumi:"includesPattern"`
	Key                    pulumi.StringInput      `pulumi:"key"`
	Notes                  pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution     pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments    pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey             pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets           pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef          pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	XrayIndex              pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupLocalPypiRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalPypiRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalPypiRepository.
type LookupLocalPypiRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalPypiRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalPypiRepositoryResult)(nil)).Elem()
}

func (o LookupLocalPypiRepositoryResultOutput) ToLookupLocalPypiRepositoryResultOutput() LookupLocalPypiRepositoryResultOutput {
	return o
}

func (o LookupLocalPypiRepositoryResultOutput) ToLookupLocalPypiRepositoryResultOutputWithContext(ctx context.Context) LookupLocalPypiRepositoryResultOutput {
	return o
}

func (o LookupLocalPypiRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalPypiRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalPypiRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalPypiRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalPypiRepositoryResultOutput{})
}
