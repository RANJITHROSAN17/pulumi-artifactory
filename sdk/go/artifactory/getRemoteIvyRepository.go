// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote Ivy repository.
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
//			_, err := artifactory.LookupRemoteIvyRepository(ctx, &artifactory.LookupRemoteIvyRepositoryArgs{
//				Key: "remote-ivy",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteIvyRepository(ctx *pulumi.Context, args *LookupRemoteIvyRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteIvyRepositoryResult, error) {
	var rv LookupRemoteIvyRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteIvyRepository:getRemoteIvyRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteIvyRepository.
type LookupRemoteIvyRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                         `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                          `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                         `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                         `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                         `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                         `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteIvyRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                       `pulumi:"description"`
	DownloadDirect            *bool                                         `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                         `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                       `pulumi:"excludesPattern"`
	// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
	FetchJarsEagerly *bool `pulumi:"fetchJarsEagerly"`
	// (Optional, Default: `false`) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
	FetchSourcesEagerly *bool `pulumi:"fetchSourcesEagerly"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases *bool `pulumi:"handleReleases"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots *bool   `pulumi:"handleSnapshots"`
	HardFail        *bool   `pulumi:"hardFail"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                              string   `pulumi:"key"`
	ListRemoteFolderItems            *bool    `pulumi:"listRemoteFolderItems"`
	LocalAddress                     *string  `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs     *int     `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList *string  `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds         *int     `pulumi:"missedCachePeriodSeconds"`
	Notes                            *string  `pulumi:"notes"`
	Offline                          *bool    `pulumi:"offline"`
	Password                         *string  `pulumi:"password"`
	PriorityResolution               *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments              []string `pulumi:"projectEnvironments"`
	ProjectKey                       *string  `pulumi:"projectKey"`
	PropertySets                     []string `pulumi:"propertySets"`
	Proxy                            *string  `pulumi:"proxy"`
	QueryParams                      *string  `pulumi:"queryParams"`
	// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
	RejectInvalidJars *bool `pulumi:"rejectInvalidJars"`
	// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
	RemoteRepoChecksumPolicyType *string `pulumi:"remoteRepoChecksumPolicyType"`
	RemoteRepoLayoutRef          *string `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                *string `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds  *int    `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration           *bool   `pulumi:"shareConfiguration"`
	SocketTimeoutMillis          *int    `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally        *bool   `pulumi:"storeArtifactsLocally"`
	// (Optional, Default: `true`) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`'
	SuppressPomConsistencyChecks      *bool   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string `pulumi:"url"`
	Username                          *string `pulumi:"username"`
	XrayIndex                         *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemoteIvyRepository.
type LookupRemoteIvyRepositoryResult struct {
	AllowAnyHostAuth          *bool                                        `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                         `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                        `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                        `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                        `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                        `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteIvyRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                      `pulumi:"description"`
	DownloadDirect            *bool                                        `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                        `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                      `pulumi:"excludesPattern"`
	// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
	FetchJarsEagerly *bool `pulumi:"fetchJarsEagerly"`
	// (Optional, Default: `false`) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
	FetchSourcesEagerly *bool `pulumi:"fetchSourcesEagerly"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases *bool `pulumi:"handleReleases"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots *bool `pulumi:"handleSnapshots"`
	HardFail        *bool `pulumi:"hardFail"`
	// The provider-assigned unique ID for this managed resource.
	Id                               string   `pulumi:"id"`
	IncludesPattern                  *string  `pulumi:"includesPattern"`
	Key                              string   `pulumi:"key"`
	ListRemoteFolderItems            *bool    `pulumi:"listRemoteFolderItems"`
	LocalAddress                     *string  `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs     *int     `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList *string  `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds         *int     `pulumi:"missedCachePeriodSeconds"`
	Notes                            *string  `pulumi:"notes"`
	Offline                          *bool    `pulumi:"offline"`
	PackageType                      string   `pulumi:"packageType"`
	Password                         *string  `pulumi:"password"`
	PriorityResolution               *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments              []string `pulumi:"projectEnvironments"`
	ProjectKey                       *string  `pulumi:"projectKey"`
	PropertySets                     []string `pulumi:"propertySets"`
	Proxy                            *string  `pulumi:"proxy"`
	QueryParams                      *string  `pulumi:"queryParams"`
	// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
	RejectInvalidJars *bool `pulumi:"rejectInvalidJars"`
	// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
	RemoteRepoChecksumPolicyType *string `pulumi:"remoteRepoChecksumPolicyType"`
	RemoteRepoLayoutRef          *string `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                *string `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds  *int    `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration           bool    `pulumi:"shareConfiguration"`
	SocketTimeoutMillis          *int    `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally        *bool   `pulumi:"storeArtifactsLocally"`
	// (Optional, Default: `true`) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`'
	SuppressPomConsistencyChecks      *bool   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string `pulumi:"url"`
	Username                          *string `pulumi:"username"`
	XrayIndex                         *bool   `pulumi:"xrayIndex"`
}

func LookupRemoteIvyRepositoryOutput(ctx *pulumi.Context, args LookupRemoteIvyRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteIvyRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteIvyRepositoryResult, error) {
			args := v.(LookupRemoteIvyRepositoryArgs)
			r, err := LookupRemoteIvyRepository(ctx, &args, opts...)
			var s LookupRemoteIvyRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteIvyRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteIvyRepository.
type LookupRemoteIvyRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                  `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                   `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                  `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                  `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                  `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                  `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteIvyRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                `pulumi:"description"`
	DownloadDirect            pulumi.BoolPtrInput                                  `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                  `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                `pulumi:"excludesPattern"`
	// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
	FetchJarsEagerly pulumi.BoolPtrInput `pulumi:"fetchJarsEagerly"`
	// (Optional, Default: `false`) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
	FetchSourcesEagerly pulumi.BoolPtrInput `pulumi:"fetchSourcesEagerly"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
	HandleReleases pulumi.BoolPtrInput `pulumi:"handleReleases"`
	// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
	HandleSnapshots pulumi.BoolPtrInput   `pulumi:"handleSnapshots"`
	HardFail        pulumi.BoolPtrInput   `pulumi:"hardFail"`
	IncludesPattern pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                              pulumi.StringInput      `pulumi:"key"`
	ListRemoteFolderItems            pulumi.BoolPtrInput     `pulumi:"listRemoteFolderItems"`
	LocalAddress                     pulumi.StringPtrInput   `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs     pulumi.IntPtrInput      `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList pulumi.StringPtrInput   `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds         pulumi.IntPtrInput      `pulumi:"missedCachePeriodSeconds"`
	Notes                            pulumi.StringPtrInput   `pulumi:"notes"`
	Offline                          pulumi.BoolPtrInput     `pulumi:"offline"`
	Password                         pulumi.StringPtrInput   `pulumi:"password"`
	PriorityResolution               pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments              pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                       pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets                     pulumi.StringArrayInput `pulumi:"propertySets"`
	Proxy                            pulumi.StringPtrInput   `pulumi:"proxy"`
	QueryParams                      pulumi.StringPtrInput   `pulumi:"queryParams"`
	// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
	RejectInvalidJars pulumi.BoolPtrInput `pulumi:"rejectInvalidJars"`
	// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
	RemoteRepoChecksumPolicyType pulumi.StringPtrInput `pulumi:"remoteRepoChecksumPolicyType"`
	RemoteRepoLayoutRef          pulumi.StringPtrInput `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                pulumi.StringPtrInput `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds  pulumi.IntPtrInput    `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration           pulumi.BoolPtrInput   `pulumi:"shareConfiguration"`
	SocketTimeoutMillis          pulumi.IntPtrInput    `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally        pulumi.BoolPtrInput   `pulumi:"storeArtifactsLocally"`
	// (Optional, Default: `true`) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`'
	SuppressPomConsistencyChecks      pulumi.BoolPtrInput   `pulumi:"suppressPomConsistencyChecks"`
	SynchronizeProperties             pulumi.BoolPtrInput   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours pulumi.IntPtrInput    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               pulumi.StringPtrInput `pulumi:"url"`
	Username                          pulumi.StringPtrInput `pulumi:"username"`
	XrayIndex                         pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupRemoteIvyRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteIvyRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteIvyRepository.
type LookupRemoteIvyRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteIvyRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteIvyRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteIvyRepositoryResultOutput) ToLookupRemoteIvyRepositoryResultOutput() LookupRemoteIvyRepositoryResultOutput {
	return o
}

func (o LookupRemoteIvyRepositoryResultOutput) ToLookupRemoteIvyRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteIvyRepositoryResultOutput {
	return o
}

func (o LookupRemoteIvyRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) ContentSynchronisation() GetRemoteIvyRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) GetRemoteIvyRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteIvyRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// (Optional, Default: `false`) When set, if a POM is requested, Artifactory attempts to fetch the corresponding jar in the background. This will accelerate first access time to the jar when it is subsequently requested.
func (o LookupRemoteIvyRepositoryResultOutput) FetchJarsEagerly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.FetchJarsEagerly }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `false`) When set, if a binaries jar is requested, Artifactory attempts to fetch the corresponding source jar in the background. This will accelerate first access time to the source jar when it is subsequently requested.
func (o LookupRemoteIvyRepositoryResultOutput) FetchSourcesEagerly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.FetchSourcesEagerly }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `true`) If set, Artifactory allows you to deploy release artifacts into this repository.
func (o LookupRemoteIvyRepositoryResultOutput) HandleReleases() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.HandleReleases }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `true`) If set, Artifactory allows you to deploy snapshot artifacts into this repository.
func (o LookupRemoteIvyRepositoryResultOutput) HandleSnapshots() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.HandleSnapshots }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteIvyRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

// (Optional, Default: `false`) Reject the caching of jar files that are found to be invalid. For example, pseudo jars retrieved behind a "captive portal".
func (o LookupRemoteIvyRepositoryResultOutput) RejectInvalidJars() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.RejectInvalidJars }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `generate-if-absent`) Checking the Checksum effectively verifies the integrity of a deployed resource. The Checksum Policy determines how the system behaves when a client checksum for a remote resource is missing or conflicts with the locally calculated checksum. Available policies are `generate-if-absent`, `fail`, `ignore-and-generate`, and `pass-thru`.
func (o LookupRemoteIvyRepositoryResultOutput) RemoteRepoChecksumPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.RemoteRepoChecksumPolicyType }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

// (Optional, Default: `true`) - By default, the system keeps your repositories healthy by refusing POMs with incorrect coordinates (path). If the groupId:artifactId:version information inside the POM does not match the deployed path, Artifactory rejects the deployment with a "409 Conflict" error. You can disable this behavior by setting this attribute to `true`'
func (o LookupRemoteIvyRepositoryResultOutput) SuppressPomConsistencyChecks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.SuppressPomConsistencyChecks }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteIvyRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteIvyRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteIvyRepositoryResultOutput{})
}
