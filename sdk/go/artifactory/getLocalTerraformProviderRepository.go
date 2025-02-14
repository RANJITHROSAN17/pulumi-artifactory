// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLocalTerraformProviderRepository(ctx *pulumi.Context, args *LookupLocalTerraformProviderRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalTerraformProviderRepositoryResult, error) {
	var rv LookupLocalTerraformProviderRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalTerraformProviderRepository:getLocalTerraformProviderRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalTerraformProviderRepository.
type LookupLocalTerraformProviderRepositoryArgs struct {
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

// A collection of values returned by getLocalTerraformProviderRepository.
type LookupLocalTerraformProviderRepositoryResult struct {
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

func LookupLocalTerraformProviderRepositoryOutput(ctx *pulumi.Context, args LookupLocalTerraformProviderRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalTerraformProviderRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalTerraformProviderRepositoryResult, error) {
			args := v.(LookupLocalTerraformProviderRepositoryArgs)
			r, err := LookupLocalTerraformProviderRepository(ctx, &args, opts...)
			var s LookupLocalTerraformProviderRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalTerraformProviderRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalTerraformProviderRepository.
type LookupLocalTerraformProviderRepositoryOutputArgs struct {
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

func (LookupLocalTerraformProviderRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalTerraformProviderRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalTerraformProviderRepository.
type LookupLocalTerraformProviderRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalTerraformProviderRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalTerraformProviderRepositoryResult)(nil)).Elem()
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) ToLookupLocalTerraformProviderRepositoryResultOutput() LookupLocalTerraformProviderRepositoryResultOutput {
	return o
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) ToLookupLocalTerraformProviderRepositoryResultOutputWithContext(ctx context.Context) LookupLocalTerraformProviderRepositoryResultOutput {
	return o
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalTerraformProviderRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalTerraformProviderRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalTerraformProviderRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalTerraformProviderRepositoryResultOutput{})
}
