// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Maven repository.
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
//			_, err := artifactory.LookupFederatedMavenRepository(ctx, &artifactory.LookupFederatedMavenRepositoryArgs{
//				Key: "federated-test-maven-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedMavenRepository(ctx *pulumi.Context, args *LookupFederatedMavenRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedMavenRepositoryResult, error) {
	var rv LookupFederatedMavenRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedMavenRepository:getFederatedMavenRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedMavenRepository.
type LookupFederatedMavenRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	HandleReleases         *bool   `pulumi:"handleReleases"`
	HandleSnapshots        *bool   `pulumi:"handleSnapshots"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                string `pulumi:"key"`
	MaxUniqueSnapshots *int   `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members                      []GetFederatedMavenRepositoryMember `pulumi:"members"`
	Notes                        *string                             `pulumi:"notes"`
	PriorityResolution           *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments          []string                            `pulumi:"projectEnvironments"`
	ProjectKey                   *string                             `pulumi:"projectKey"`
	PropertySets                 []string                            `pulumi:"propertySets"`
	RepoLayoutRef                *string                             `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string                             `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool                               `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool                               `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedMavenRepository.
type LookupFederatedMavenRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     *string `pulumi:"checksumPolicyType"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	HandleReleases         *bool   `pulumi:"handleReleases"`
	HandleSnapshots        *bool   `pulumi:"handleSnapshots"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string `pulumi:"id"`
	IncludesPattern    string `pulumi:"includesPattern"`
	Key                string `pulumi:"key"`
	MaxUniqueSnapshots *int   `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members                      []GetFederatedMavenRepositoryMember `pulumi:"members"`
	Notes                        *string                             `pulumi:"notes"`
	PackageType                  string                              `pulumi:"packageType"`
	PriorityResolution           *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments          []string                            `pulumi:"projectEnvironments"`
	ProjectKey                   *string                             `pulumi:"projectKey"`
	PropertySets                 []string                            `pulumi:"propertySets"`
	RepoLayoutRef                *string                             `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      *string                             `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks *bool                               `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    *bool                               `pulumi:"xrayIndex"`
}

func LookupFederatedMavenRepositoryOutput(ctx *pulumi.Context, args LookupFederatedMavenRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedMavenRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedMavenRepositoryResult, error) {
			args := v.(LookupFederatedMavenRepositoryArgs)
			r, err := LookupFederatedMavenRepository(ctx, &args, opts...)
			var s LookupFederatedMavenRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedMavenRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedMavenRepository.
type LookupFederatedMavenRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	ChecksumPolicyType     pulumi.StringPtrInput `pulumi:"checksumPolicyType"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	HandleReleases         pulumi.BoolPtrInput   `pulumi:"handleReleases"`
	HandleSnapshots        pulumi.BoolPtrInput   `pulumi:"handleSnapshots"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                pulumi.StringInput `pulumi:"key"`
	MaxUniqueSnapshots pulumi.IntPtrInput `pulumi:"maxUniqueSnapshots"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members                      GetFederatedMavenRepositoryMemberArrayInput `pulumi:"members"`
	Notes                        pulumi.StringPtrInput                       `pulumi:"notes"`
	PriorityResolution           pulumi.BoolPtrInput                         `pulumi:"priorityResolution"`
	ProjectEnvironments          pulumi.StringArrayInput                     `pulumi:"projectEnvironments"`
	ProjectKey                   pulumi.StringPtrInput                       `pulumi:"projectKey"`
	PropertySets                 pulumi.StringArrayInput                     `pulumi:"propertySets"`
	RepoLayoutRef                pulumi.StringPtrInput                       `pulumi:"repoLayoutRef"`
	SnapshotVersionBehavior      pulumi.StringPtrInput                       `pulumi:"snapshotVersionBehavior"`
	SuppressPomConsistencyChecks pulumi.BoolPtrInput                         `pulumi:"suppressPomConsistencyChecks"`
	XrayIndex                    pulumi.BoolPtrInput                         `pulumi:"xrayIndex"`
}

func (LookupFederatedMavenRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedMavenRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedMavenRepository.
type LookupFederatedMavenRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedMavenRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedMavenRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedMavenRepositoryResultOutput) ToLookupFederatedMavenRepositoryResultOutput() LookupFederatedMavenRepositoryResultOutput {
	return o
}

func (o LookupFederatedMavenRepositoryResultOutput) ToLookupFederatedMavenRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedMavenRepositoryResultOutput {
	return o
}

func (o LookupFederatedMavenRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) ChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *string { return v.ChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedMavenRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedMavenRepositoryResultOutput) Members() GetFederatedMavenRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) []GetFederatedMavenRepositoryMember { return v.Members }).(GetFederatedMavenRepositoryMemberArrayOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) SnapshotVersionBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *string { return v.SnapshotVersionBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedMavenRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedMavenRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedMavenRepositoryResultOutput{})
}
