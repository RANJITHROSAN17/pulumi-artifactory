// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLocalMavenRepository(ctx *pulumi.Context, args *LookupLocalMavenRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupLocalMavenRepositoryResult, error) {
	var rv LookupLocalMavenRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalMavenRepository:getLocalMavenRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalMavenRepository.
type LookupLocalMavenRepositoryArgs struct {
	ArchiveBrowsingEnabled       *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut                   *bool    `pulumi:"blackedOut"`
	CdnRedirect                  *bool    `pulumi:"cdnRedirect"`
	ChecksumPolicyType           *string  `pulumi:"checksumPolicyType"`
	Description                  *string  `pulumi:"description"`
	DownloadDirect               *bool    `pulumi:"downloadDirect"`
	ExcludesPattern              *string  `pulumi:"excludesPattern"`
	HandleReleases               *bool    `pulumi:"handleReleases"`
	HandleSnapshots              *bool    `pulumi:"handleSnapshots"`
	IncludesPattern              *string  `pulumi:"includesPattern"`
	Key                          string   `pulumi:"key"`
	MaxUniqueSnapshots           *int     `pulumi:"maxUniqueSnapshots"`
	Notes                        *string  `pulumi:"notes"`
	PriorityResolution           *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments          []string `pulumi:"projectEnvironments"`
	ProjectKey                   *string  `pulumi:"projectKey"`
	PropertySets                 []string `pulumi:"propertySets"`
	RepoLayoutRef                *string  `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string  `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool    `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalMavenRepository.
type LookupLocalMavenRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	HandleReleases         *bool   `pulumi:"handleReleases"`
	HandleSnapshots        *bool   `pulumi:"handleSnapshots"`
	// The provider-assigned unique ID for this managed resource.
	Id                           string   `pulumi:"id"`
	IncludesPattern              string   `pulumi:"includesPattern"`
	Key                          string   `pulumi:"key"`
	MaxUniqueSnapshots           *int     `pulumi:"maxUniqueSnapshots"`
	Notes                        *string  `pulumi:"notes"`
	PackageType                  string   `pulumi:"packageType"`
	PriorityResolution           *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments          []string `pulumi:"projectEnvironments"`
	ProjectKey                   *string  `pulumi:"projectKey"`
	PropertySets                 []string `pulumi:"propertySets"`
	RepoLayoutRef                *string  `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string  `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool    `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool    `pulumi:"xrayIndex"`
}

func LookupLocalMavenRepositoryOutput(ctx *pulumi.Context, args LookupLocalMavenRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupLocalMavenRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLocalMavenRepositoryResult, error) {
			args := v.(LookupLocalMavenRepositoryArgs)
			r, err := LookupLocalMavenRepository(ctx, &args, opts...)
			var s LookupLocalMavenRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLocalMavenRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalMavenRepository.
type LookupLocalMavenRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled       pulumi.BoolPtrInput     `pulumi:"archiveBrowsingEnabled"`
	BlackedOut                   pulumi.BoolPtrInput     `pulumi:"blackedOut"`
	CdnRedirect                  pulumi.BoolPtrInput     `pulumi:"cdnRedirect"`
	ChecksumPolicyType           pulumi.StringPtrInput   `pulumi:"checksumPolicyType"`
	Description                  pulumi.StringPtrInput   `pulumi:"description"`
	DownloadDirect               pulumi.BoolPtrInput     `pulumi:"downloadDirect"`
	ExcludesPattern              pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	HandleReleases               pulumi.BoolPtrInput     `pulumi:"handleReleases"`
	HandleSnapshots              pulumi.BoolPtrInput     `pulumi:"handleSnapshots"`
	IncludesPattern              pulumi.StringPtrInput   `pulumi:"includesPattern"`
	Key                          pulumi.StringInput      `pulumi:"key"`
	MaxUniqueSnapshots           pulumi.IntPtrInput      `pulumi:"maxUniqueSnapshots"`
	Notes                        pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution           pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments          pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                   pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets                 pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef                pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      pulumi.StringPtrInput   `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks pulumi.BoolPtrInput     `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupLocalMavenRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalMavenRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalMavenRepository.
type LookupLocalMavenRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupLocalMavenRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLocalMavenRepositoryResult)(nil)).Elem()
}

func (o LookupLocalMavenRepositoryResultOutput) ToLookupLocalMavenRepositoryResultOutput() LookupLocalMavenRepositoryResultOutput {
	return o
}

func (o LookupLocalMavenRepositoryResultOutput) ToLookupLocalMavenRepositoryResultOutputWithContext(ctx context.Context) LookupLocalMavenRepositoryResultOutput {
	return o
}

func (o LookupLocalMavenRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *string { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLocalMavenRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *string { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupLocalMavenRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupLocalMavenRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLocalMavenRepositoryResultOutput{})
}
