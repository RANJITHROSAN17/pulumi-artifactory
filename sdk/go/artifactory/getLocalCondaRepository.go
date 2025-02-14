// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLocalCondaRepository(ctx *pulumi.Context, args *LookupLocalCondaRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalCondaRepositoryResult, error) {
	var rv LookupLocalCondaRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalCondaRepository:getLocalCondaRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalCondaRepository.
type LookupLocalCondaRepositoryArgs struct {
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

// A collection of values returned by getLocalCondaRepository.
type LookupLocalCondaRepositoryResult struct {
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

func LookupLocalCondaRepositoryOutput(ctx *pulumi.Context, args LookupLocalCondaRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalCondaRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalCondaRepositoryResult, error) {
			args := v.(LookupLocalCondaRepositoryArgs)
			r, err := LookupLocalCondaRepository(ctx, &args, opts...)
			var s LookupLocalCondaRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalCondaRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalCondaRepository.
type LookupLocalCondaRepositoryOutputArgs struct {
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

func (LookupLocalCondaRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalCondaRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalCondaRepository.
type LookupLocalCondaRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalCondaRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalCondaRepositoryResult)(nil)).Elem()
}

func (o LookupLocalCondaRepositoryResultOutput) ToLookupLocalCondaRepositoryResultOutput() LookupLocalCondaRepositoryResultOutput {
	return o
}

func (o LookupLocalCondaRepositoryResultOutput) ToLookupLocalCondaRepositoryResultOutputWithContext(ctx context.Context) LookupLocalCondaRepositoryResultOutput {
	return o
}

func (o LookupLocalCondaRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalCondaRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalCondaRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalCondaRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalCondaRepositoryResultOutput{})
}
