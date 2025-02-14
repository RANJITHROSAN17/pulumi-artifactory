// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Bower repository.
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
//			_, err := artifactory.LookupFederatedBowerRepository(ctx, &artifactory.LookupFederatedBowerRepositoryArgs{
//				Key: "federated-test-bower-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedBowerRepository(ctx *pulumi.Context, args *LookupFederatedBowerRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedBowerRepositoryResult, error) {
	var rv LookupFederatedBowerRepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedBowerRepository:getFederatedBowerRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedBowerRepository.
type LookupFederatedBowerRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedBowerRepositoryMember `pulumi:"members"`
	Notes               *string                             `pulumi:"notes"`
	PriorityResolution  *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments []string                            `pulumi:"projectEnvironments"`
	ProjectKey          *string                             `pulumi:"projectKey"`
	PropertySets        []string                            `pulumi:"propertySets"`
	RepoLayoutRef       *string                             `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                               `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedBowerRepository.
type LookupFederatedBowerRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string `pulumi:"id"`
	IncludesPattern string `pulumi:"includesPattern"`
	Key             string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedBowerRepositoryMember `pulumi:"members"`
	Notes               *string                             `pulumi:"notes"`
	PackageType         string                              `pulumi:"packageType"`
	PriorityResolution  *bool                               `pulumi:"priorityResolution"`
	ProjectEnvironments []string                            `pulumi:"projectEnvironments"`
	ProjectKey          *string                             `pulumi:"projectKey"`
	PropertySets        []string                            `pulumi:"propertySets"`
	RepoLayoutRef       *string                             `pulumi:"repoLayoutRef"`
	XrayIndex           *bool                               `pulumi:"xrayIndex"`
}

func LookupFederatedBowerRepositoryOutput(ctx *pulumi.Context, args LookupFederatedBowerRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedBowerRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFederatedBowerRepositoryResult, error) {
			args := v.(LookupFederatedBowerRepositoryArgs)
			r, err := LookupFederatedBowerRepository(ctx, &args, opts...)
			var s LookupFederatedBowerRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFederatedBowerRepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedBowerRepository.
type LookupFederatedBowerRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedBowerRepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                       `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                         `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                     `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                       `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                     `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput                       `pulumi:"repoLayoutRef"`
	XrayIndex           pulumi.BoolPtrInput                         `pulumi:"xrayIndex"`
}

func (LookupFederatedBowerRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedBowerRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedBowerRepository.
type LookupFederatedBowerRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedBowerRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedBowerRepositoryResult)(nil)).Elem()
}

func (o LookupFederatedBowerRepositoryResultOutput) ToLookupFederatedBowerRepositoryResultOutput() LookupFederatedBowerRepositoryResultOutput {
	return o
}

func (o LookupFederatedBowerRepositoryResultOutput) ToLookupFederatedBowerRepositoryResultOutputWithContext(ctx context.Context) LookupFederatedBowerRepositoryResultOutput {
	return o
}

func (o LookupFederatedBowerRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedBowerRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedBowerRepositoryResultOutput) Members() GetFederatedBowerRepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) []GetFederatedBowerRepositoryMember { return v.Members }).(GetFederatedBowerRepositoryMemberArrayOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedBowerRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedBowerRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedBowerRepositoryResultOutput{})
}
