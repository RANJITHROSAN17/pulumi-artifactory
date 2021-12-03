// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Remote npm Repository Resource
//
// Provides an Artifactory remote `npm` repository resource. This provides npm specific fields and is the only way to get them
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/npm+Registry)
//
// ## Example Usage
//
// Create a new Artifactory remote npm repository called my-remote-npm
// for brevity sake, only npm specific fields are included; for other fields see documentation for
// generic repo.
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewRemoteNpmRepository(ctx, "thing", &artifactory.RemoteNpmRepositoryArgs{
// 			Key:                              pulumi.String("remote-thing-npm"),
// 			ListRemoteFolderItems:            pulumi.Bool(true),
// 			MismatchingMimeTypesOverrideList: pulumi.String("application/json,application/xml"),
// 			Url:                              pulumi.String("https://registry.npmjs.org/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RemoteNpmRepository struct {
	pulumi.CustomResourceState

	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth         pulumi.BoolOutput   `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs pulumi.IntPtrOutput `pulumi:"assumedOfflinePeriodSecs"`
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut pulumi.BoolOutput `pulumi:"blackedOut"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes pulumi.BoolOutput `pulumi:"blockMismatchingMimeTypes"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests     pulumi.BoolOutput                               `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate   pulumi.StringOutput                             `pulumi:"clientTlsCertificate"`
	ContentSynchronisation RemoteNpmRepositoryContentSynchronisationOutput `pulumi:"contentSynchronisation"`
	Description            pulumi.StringOutput                             `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolOutput   `pulumi:"enableCookieManagement"`
	ExcludesPattern        pulumi.StringOutput `pulumi:"excludesPattern"`
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs pulumi.IntOutput    `pulumi:"failedRetrievalCachePeriodSecs"`
	HardFail                       pulumi.BoolOutput   `pulumi:"hardFail"`
	IncludesPattern                pulumi.StringOutput `pulumi:"includesPattern"`
	// The repository identifier. Must be unique system-wide
	Key pulumi.StringOutput `pulumi:"key"`
	// - No documentation could be found. This field exist in the API but not in the UI
	ListRemoteFolderItems pulumi.BoolPtrOutput   `pulumi:"listRemoteFolderItems"`
	LocalAddress          pulumi.StringPtrOutput `pulumi:"localAddress"`
	// - No documentation could be found. This field exist in the API but not in the UI
	MismatchingMimeTypesOverrideList pulumi.StringPtrOutput `pulumi:"mismatchingMimeTypesOverrideList"`
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds pulumi.IntOutput       `pulumi:"missedCachePeriodSeconds"`
	Notes                    pulumi.StringPtrOutput `pulumi:"notes"`
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline              pulumi.BoolOutput        `pulumi:"offline"`
	PackageType          pulumi.StringOutput      `pulumi:"packageType"`
	Password             pulumi.StringPtrOutput   `pulumi:"password"`
	PriorityResolution   pulumi.BoolOutput        `pulumi:"priorityResolution"`
	PropagateQueryParams pulumi.BoolPtrOutput     `pulumi:"propagateQueryParams"`
	PropertySets         pulumi.StringArrayOutput `pulumi:"propertySets"`
	Proxy                pulumi.StringOutput      `pulumi:"proxy"`
	RemoteRepoLayoutRef  pulumi.StringOutput      `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef        pulumi.StringOutput      `pulumi:"repoLayoutRef"`
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds pulumi.IntOutput  `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration          pulumi.BoolOutput `pulumi:"shareConfiguration"`
	SocketTimeoutMillis         pulumi.IntOutput  `pulumi:"socketTimeoutMillis"`
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally pulumi.BoolOutput `pulumi:"storeArtifactsLocally"`
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               pulumi.BoolOutput      `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodEnabled pulumi.BoolOutput      `pulumi:"unusedArtifactsCleanupPeriodEnabled"`
	UnusedArtifactsCleanupPeriodHours   pulumi.IntOutput       `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                                 pulumi.StringOutput    `pulumi:"url"`
	Username                            pulumi.StringPtrOutput `pulumi:"username"`
	XrayIndex                           pulumi.BoolOutput      `pulumi:"xrayIndex"`
}

// NewRemoteNpmRepository registers a new resource with the given unique name, arguments, and options.
func NewRemoteNpmRepository(ctx *pulumi.Context,
	name string, args *RemoteNpmRepositoryArgs, opts ...pulumi.ResourceOption) (*RemoteNpmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource RemoteNpmRepository
	err := ctx.RegisterResource("artifactory:index/remoteNpmRepository:RemoteNpmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemoteNpmRepository gets an existing RemoteNpmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemoteNpmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteNpmRepositoryState, opts ...pulumi.ResourceOption) (*RemoteNpmRepository, error) {
	var resource RemoteNpmRepository
	err := ctx.ReadResource("artifactory:index/remoteNpmRepository:RemoteNpmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemoteNpmRepository resources.
type remoteNpmRepositoryState struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth         *bool `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs *int  `pulumi:"assumedOfflinePeriodSecs"`
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut *bool `pulumi:"blackedOut"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes *bool `pulumi:"blockMismatchingMimeTypes"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests     *bool                                      `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate   *string                                    `pulumi:"clientTlsCertificate"`
	ContentSynchronisation *RemoteNpmRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                    `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement *bool   `pulumi:"enableCookieManagement"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs *int    `pulumi:"failedRetrievalCachePeriodSecs"`
	HardFail                       *bool   `pulumi:"hardFail"`
	IncludesPattern                *string `pulumi:"includesPattern"`
	// The repository identifier. Must be unique system-wide
	Key *string `pulumi:"key"`
	// - No documentation could be found. This field exist in the API but not in the UI
	ListRemoteFolderItems *bool   `pulumi:"listRemoteFolderItems"`
	LocalAddress          *string `pulumi:"localAddress"`
	// - No documentation could be found. This field exist in the API but not in the UI
	MismatchingMimeTypesOverrideList *string `pulumi:"mismatchingMimeTypesOverrideList"`
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds *int    `pulumi:"missedCachePeriodSeconds"`
	Notes                    *string `pulumi:"notes"`
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline              *bool    `pulumi:"offline"`
	PackageType          *string  `pulumi:"packageType"`
	Password             *string  `pulumi:"password"`
	PriorityResolution   *bool    `pulumi:"priorityResolution"`
	PropagateQueryParams *bool    `pulumi:"propagateQueryParams"`
	PropertySets         []string `pulumi:"propertySets"`
	Proxy                *string  `pulumi:"proxy"`
	RemoteRepoLayoutRef  *string  `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef        *string  `pulumi:"repoLayoutRef"`
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds *int  `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration          *bool `pulumi:"shareConfiguration"`
	SocketTimeoutMillis         *int  `pulumi:"socketTimeoutMillis"`
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally *bool `pulumi:"storeArtifactsLocally"`
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodEnabled *bool   `pulumi:"unusedArtifactsCleanupPeriodEnabled"`
	UnusedArtifactsCleanupPeriodHours   *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                                 *string `pulumi:"url"`
	Username                            *string `pulumi:"username"`
	XrayIndex                           *bool   `pulumi:"xrayIndex"`
}

type RemoteNpmRepositoryState struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth         pulumi.BoolPtrInput
	AssumedOfflinePeriodSecs pulumi.IntPtrInput
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests     pulumi.BoolPtrInput
	ClientTlsCertificate   pulumi.StringPtrInput
	ContentSynchronisation RemoteNpmRepositoryContentSynchronisationPtrInput
	Description            pulumi.StringPtrInput
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs pulumi.IntPtrInput
	HardFail                       pulumi.BoolPtrInput
	IncludesPattern                pulumi.StringPtrInput
	// The repository identifier. Must be unique system-wide
	Key pulumi.StringPtrInput
	// - No documentation could be found. This field exist in the API but not in the UI
	ListRemoteFolderItems pulumi.BoolPtrInput
	LocalAddress          pulumi.StringPtrInput
	// - No documentation could be found. This field exist in the API but not in the UI
	MismatchingMimeTypesOverrideList pulumi.StringPtrInput
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds pulumi.IntPtrInput
	Notes                    pulumi.StringPtrInput
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline              pulumi.BoolPtrInput
	PackageType          pulumi.StringPtrInput
	Password             pulumi.StringPtrInput
	PriorityResolution   pulumi.BoolPtrInput
	PropagateQueryParams pulumi.BoolPtrInput
	PropertySets         pulumi.StringArrayInput
	Proxy                pulumi.StringPtrInput
	RemoteRepoLayoutRef  pulumi.StringPtrInput
	RepoLayoutRef        pulumi.StringPtrInput
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
	ShareConfiguration          pulumi.BoolPtrInput
	SocketTimeoutMillis         pulumi.IntPtrInput
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally pulumi.BoolPtrInput
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodEnabled pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodHours   pulumi.IntPtrInput
	Url                                 pulumi.StringPtrInput
	Username                            pulumi.StringPtrInput
	XrayIndex                           pulumi.BoolPtrInput
}

func (RemoteNpmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteNpmRepositoryState)(nil)).Elem()
}

type remoteNpmRepositoryArgs struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth         *bool `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs *int  `pulumi:"assumedOfflinePeriodSecs"`
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut *bool `pulumi:"blackedOut"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes *bool `pulumi:"blockMismatchingMimeTypes"`
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests     *bool                                      `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate   *string                                    `pulumi:"clientTlsCertificate"`
	ContentSynchronisation *RemoteNpmRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                    `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement *bool   `pulumi:"enableCookieManagement"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	HardFail               *bool   `pulumi:"hardFail"`
	IncludesPattern        *string `pulumi:"includesPattern"`
	// The repository identifier. Must be unique system-wide
	Key string `pulumi:"key"`
	// - No documentation could be found. This field exist in the API but not in the UI
	ListRemoteFolderItems *bool   `pulumi:"listRemoteFolderItems"`
	LocalAddress          *string `pulumi:"localAddress"`
	// - No documentation could be found. This field exist in the API but not in the UI
	MismatchingMimeTypesOverrideList *string `pulumi:"mismatchingMimeTypesOverrideList"`
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds *int    `pulumi:"missedCachePeriodSeconds"`
	Notes                    *string `pulumi:"notes"`
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline              *bool    `pulumi:"offline"`
	Password             *string  `pulumi:"password"`
	PriorityResolution   *bool    `pulumi:"priorityResolution"`
	PropagateQueryParams *bool    `pulumi:"propagateQueryParams"`
	PropertySets         []string `pulumi:"propertySets"`
	Proxy                *string  `pulumi:"proxy"`
	RemoteRepoLayoutRef  *string  `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef        *string  `pulumi:"repoLayoutRef"`
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds *int  `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration          *bool `pulumi:"shareConfiguration"`
	SocketTimeoutMillis         *int  `pulumi:"socketTimeoutMillis"`
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally *bool `pulumi:"storeArtifactsLocally"`
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               *bool   `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodEnabled *bool   `pulumi:"unusedArtifactsCleanupPeriodEnabled"`
	UnusedArtifactsCleanupPeriodHours   *int    `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                                 string  `pulumi:"url"`
	Username                            *string `pulumi:"username"`
	XrayIndex                           *bool   `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a RemoteNpmRepository resource.
type RemoteNpmRepositoryArgs struct {
	// Also known as 'Lenient Host Authentication', Allow credentials of this repository to be used on requests redirected to
	// any other host.
	AllowAnyHostAuth         pulumi.BoolPtrInput
	AssumedOfflinePeriodSecs pulumi.IntPtrInput
	// (A.K.A 'Ignore Repository' on the UI) When set, the repository or its local cache do not participate in artifact
	// resolution.
	BlackedOut pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BlockMismatchingMimeTypes pulumi.BoolPtrInput
	// Before caching an artifact, Artifactory first sends a HEAD request to the remote resource. In some remote resources,
	// HEAD requests are disallowed and therefore rejected, even though downloading the artifact is allowed. When checked,
	// Artifactory will bypass the HEAD request and cache the artifact directly using a GET request.
	BypassHeadRequests     pulumi.BoolPtrInput
	ClientTlsCertificate   pulumi.StringPtrInput
	ContentSynchronisation RemoteNpmRepositoryContentSynchronisationPtrInput
	Description            pulumi.StringPtrInput
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	HardFail               pulumi.BoolPtrInput
	IncludesPattern        pulumi.StringPtrInput
	// The repository identifier. Must be unique system-wide
	Key pulumi.StringInput
	// - No documentation could be found. This field exist in the API but not in the UI
	ListRemoteFolderItems pulumi.BoolPtrInput
	LocalAddress          pulumi.StringPtrInput
	// - No documentation could be found. This field exist in the API but not in the UI
	MismatchingMimeTypesOverrideList pulumi.StringPtrInput
	// This is actually the missedRetrievalCachePeriodSecs in the API
	MissedCachePeriodSeconds pulumi.IntPtrInput
	Notes                    pulumi.StringPtrInput
	// If set, Artifactory does not try to fetch remote artifacts. Only locally-cached artifacts are retrieved.
	Offline              pulumi.BoolPtrInput
	Password             pulumi.StringPtrInput
	PriorityResolution   pulumi.BoolPtrInput
	PropagateQueryParams pulumi.BoolPtrInput
	PropertySets         pulumi.StringArrayInput
	Proxy                pulumi.StringPtrInput
	RemoteRepoLayoutRef  pulumi.StringPtrInput
	RepoLayoutRef        pulumi.StringPtrInput
	// The metadataRetrievalTimeoutSecs field not allowed to be bigger then retrievalCachePeriodSecs field.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
	ShareConfiguration          pulumi.BoolPtrInput
	SocketTimeoutMillis         pulumi.IntPtrInput
	// When set, the repository should store cached artifacts locally. When not set, artifacts are not stored locally, and
	// direct repository-to-client streaming is used. This can be useful for multi-server setups over a high-speed LAN, with
	// one Artifactory caching certain data on central storage, and streaming it directly to satellite pass-though Artifactory
	// servers.
	StoreArtifactsLocally pulumi.BoolPtrInput
	// When set, remote artifacts are fetched along with their properties.
	SynchronizeProperties               pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodEnabled pulumi.BoolPtrInput
	UnusedArtifactsCleanupPeriodHours   pulumi.IntPtrInput
	Url                                 pulumi.StringInput
	Username                            pulumi.StringPtrInput
	XrayIndex                           pulumi.BoolPtrInput
}

func (RemoteNpmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteNpmRepositoryArgs)(nil)).Elem()
}

type RemoteNpmRepositoryInput interface {
	pulumi.Input

	ToRemoteNpmRepositoryOutput() RemoteNpmRepositoryOutput
	ToRemoteNpmRepositoryOutputWithContext(ctx context.Context) RemoteNpmRepositoryOutput
}

func (*RemoteNpmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteNpmRepository)(nil))
}

func (i *RemoteNpmRepository) ToRemoteNpmRepositoryOutput() RemoteNpmRepositoryOutput {
	return i.ToRemoteNpmRepositoryOutputWithContext(context.Background())
}

func (i *RemoteNpmRepository) ToRemoteNpmRepositoryOutputWithContext(ctx context.Context) RemoteNpmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteNpmRepositoryOutput)
}

func (i *RemoteNpmRepository) ToRemoteNpmRepositoryPtrOutput() RemoteNpmRepositoryPtrOutput {
	return i.ToRemoteNpmRepositoryPtrOutputWithContext(context.Background())
}

func (i *RemoteNpmRepository) ToRemoteNpmRepositoryPtrOutputWithContext(ctx context.Context) RemoteNpmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteNpmRepositoryPtrOutput)
}

type RemoteNpmRepositoryPtrInput interface {
	pulumi.Input

	ToRemoteNpmRepositoryPtrOutput() RemoteNpmRepositoryPtrOutput
	ToRemoteNpmRepositoryPtrOutputWithContext(ctx context.Context) RemoteNpmRepositoryPtrOutput
}

type remoteNpmRepositoryPtrType RemoteNpmRepositoryArgs

func (*remoteNpmRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteNpmRepository)(nil))
}

func (i *remoteNpmRepositoryPtrType) ToRemoteNpmRepositoryPtrOutput() RemoteNpmRepositoryPtrOutput {
	return i.ToRemoteNpmRepositoryPtrOutputWithContext(context.Background())
}

func (i *remoteNpmRepositoryPtrType) ToRemoteNpmRepositoryPtrOutputWithContext(ctx context.Context) RemoteNpmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteNpmRepositoryPtrOutput)
}

// RemoteNpmRepositoryArrayInput is an input type that accepts RemoteNpmRepositoryArray and RemoteNpmRepositoryArrayOutput values.
// You can construct a concrete instance of `RemoteNpmRepositoryArrayInput` via:
//
//          RemoteNpmRepositoryArray{ RemoteNpmRepositoryArgs{...} }
type RemoteNpmRepositoryArrayInput interface {
	pulumi.Input

	ToRemoteNpmRepositoryArrayOutput() RemoteNpmRepositoryArrayOutput
	ToRemoteNpmRepositoryArrayOutputWithContext(context.Context) RemoteNpmRepositoryArrayOutput
}

type RemoteNpmRepositoryArray []RemoteNpmRepositoryInput

func (RemoteNpmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteNpmRepository)(nil)).Elem()
}

func (i RemoteNpmRepositoryArray) ToRemoteNpmRepositoryArrayOutput() RemoteNpmRepositoryArrayOutput {
	return i.ToRemoteNpmRepositoryArrayOutputWithContext(context.Background())
}

func (i RemoteNpmRepositoryArray) ToRemoteNpmRepositoryArrayOutputWithContext(ctx context.Context) RemoteNpmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteNpmRepositoryArrayOutput)
}

// RemoteNpmRepositoryMapInput is an input type that accepts RemoteNpmRepositoryMap and RemoteNpmRepositoryMapOutput values.
// You can construct a concrete instance of `RemoteNpmRepositoryMapInput` via:
//
//          RemoteNpmRepositoryMap{ "key": RemoteNpmRepositoryArgs{...} }
type RemoteNpmRepositoryMapInput interface {
	pulumi.Input

	ToRemoteNpmRepositoryMapOutput() RemoteNpmRepositoryMapOutput
	ToRemoteNpmRepositoryMapOutputWithContext(context.Context) RemoteNpmRepositoryMapOutput
}

type RemoteNpmRepositoryMap map[string]RemoteNpmRepositoryInput

func (RemoteNpmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteNpmRepository)(nil)).Elem()
}

func (i RemoteNpmRepositoryMap) ToRemoteNpmRepositoryMapOutput() RemoteNpmRepositoryMapOutput {
	return i.ToRemoteNpmRepositoryMapOutputWithContext(context.Background())
}

func (i RemoteNpmRepositoryMap) ToRemoteNpmRepositoryMapOutputWithContext(ctx context.Context) RemoteNpmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteNpmRepositoryMapOutput)
}

type RemoteNpmRepositoryOutput struct{ *pulumi.OutputState }

func (RemoteNpmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteNpmRepository)(nil))
}

func (o RemoteNpmRepositoryOutput) ToRemoteNpmRepositoryOutput() RemoteNpmRepositoryOutput {
	return o
}

func (o RemoteNpmRepositoryOutput) ToRemoteNpmRepositoryOutputWithContext(ctx context.Context) RemoteNpmRepositoryOutput {
	return o
}

func (o RemoteNpmRepositoryOutput) ToRemoteNpmRepositoryPtrOutput() RemoteNpmRepositoryPtrOutput {
	return o.ToRemoteNpmRepositoryPtrOutputWithContext(context.Background())
}

func (o RemoteNpmRepositoryOutput) ToRemoteNpmRepositoryPtrOutputWithContext(ctx context.Context) RemoteNpmRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemoteNpmRepository) *RemoteNpmRepository {
		return &v
	}).(RemoteNpmRepositoryPtrOutput)
}

type RemoteNpmRepositoryPtrOutput struct{ *pulumi.OutputState }

func (RemoteNpmRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteNpmRepository)(nil))
}

func (o RemoteNpmRepositoryPtrOutput) ToRemoteNpmRepositoryPtrOutput() RemoteNpmRepositoryPtrOutput {
	return o
}

func (o RemoteNpmRepositoryPtrOutput) ToRemoteNpmRepositoryPtrOutputWithContext(ctx context.Context) RemoteNpmRepositoryPtrOutput {
	return o
}

func (o RemoteNpmRepositoryPtrOutput) Elem() RemoteNpmRepositoryOutput {
	return o.ApplyT(func(v *RemoteNpmRepository) RemoteNpmRepository {
		if v != nil {
			return *v
		}
		var ret RemoteNpmRepository
		return ret
	}).(RemoteNpmRepositoryOutput)
}

type RemoteNpmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (RemoteNpmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemoteNpmRepository)(nil))
}

func (o RemoteNpmRepositoryArrayOutput) ToRemoteNpmRepositoryArrayOutput() RemoteNpmRepositoryArrayOutput {
	return o
}

func (o RemoteNpmRepositoryArrayOutput) ToRemoteNpmRepositoryArrayOutputWithContext(ctx context.Context) RemoteNpmRepositoryArrayOutput {
	return o
}

func (o RemoteNpmRepositoryArrayOutput) Index(i pulumi.IntInput) RemoteNpmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RemoteNpmRepository {
		return vs[0].([]RemoteNpmRepository)[vs[1].(int)]
	}).(RemoteNpmRepositoryOutput)
}

type RemoteNpmRepositoryMapOutput struct{ *pulumi.OutputState }

func (RemoteNpmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RemoteNpmRepository)(nil))
}

func (o RemoteNpmRepositoryMapOutput) ToRemoteNpmRepositoryMapOutput() RemoteNpmRepositoryMapOutput {
	return o
}

func (o RemoteNpmRepositoryMapOutput) ToRemoteNpmRepositoryMapOutputWithContext(ctx context.Context) RemoteNpmRepositoryMapOutput {
	return o
}

func (o RemoteNpmRepositoryMapOutput) MapIndex(k pulumi.StringInput) RemoteNpmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RemoteNpmRepository {
		return vs[0].(map[string]RemoteNpmRepository)[vs[1].(string)]
	}).(RemoteNpmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteNpmRepositoryInput)(nil)).Elem(), &RemoteNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteNpmRepositoryPtrInput)(nil)).Elem(), &RemoteNpmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteNpmRepositoryArrayInput)(nil)).Elem(), RemoteNpmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteNpmRepositoryMapInput)(nil)).Elem(), RemoteNpmRepositoryMap{})
	pulumi.RegisterOutputType(RemoteNpmRepositoryOutput{})
	pulumi.RegisterOutputType(RemoteNpmRepositoryPtrOutput{})
	pulumi.RegisterOutputType(RemoteNpmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(RemoteNpmRepositoryMapOutput{})
}
