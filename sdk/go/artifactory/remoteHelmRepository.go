// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Remote Repository Resource
//
// Provides an Artifactory remote `helm` repository resource. This provides helm specific fields and is the only way to get them.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Package+Management),
// although helm is (currently) not listed as a supported format
//
// ## Example Usage
// ### Additional Examples
// Includes only new and relevant fields, for anything else, see: generic repo.
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
// 		_, err := artifactory.NewRemoteHelmRepository(ctx, "helm_remote", &artifactory.RemoteHelmRepositoryArgs{
// 			HelmChartsBaseUrl: pulumi.String("https://foo.com"),
// 			Key:               pulumi.String("helm-remote-foo25"),
// 			Url:               pulumi.String("https://repo.chartcenter.io/"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RemoteHelmRepository struct {
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
	BypassHeadRequests     pulumi.BoolOutput                                `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate   pulumi.StringOutput                              `pulumi:"clientTlsCertificate"`
	ContentSynchronisation RemoteHelmRepositoryContentSynchronisationOutput `pulumi:"contentSynchronisation"`
	Description            pulumi.StringOutput                              `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolOutput   `pulumi:"enableCookieManagement"`
	ExcludesPattern        pulumi.StringOutput `pulumi:"excludesPattern"`
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs pulumi.IntOutput  `pulumi:"failedRetrievalCachePeriodSecs"`
	HardFail                       pulumi.BoolOutput `pulumi:"hardFail"`
	// - No documentation is available. Hopefully you know what this means
	HelmChartsBaseUrl pulumi.StringOutput `pulumi:"helmChartsBaseUrl"`
	IncludesPattern   pulumi.StringOutput `pulumi:"includesPattern"`
	// The repository identifier. Must be unique system-wide
	Key          pulumi.StringOutput    `pulumi:"key"`
	LocalAddress pulumi.StringPtrOutput `pulumi:"localAddress"`
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

// NewRemoteHelmRepository registers a new resource with the given unique name, arguments, and options.
func NewRemoteHelmRepository(ctx *pulumi.Context,
	name string, args *RemoteHelmRepositoryArgs, opts ...pulumi.ResourceOption) (*RemoteHelmRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HelmChartsBaseUrl == nil {
		return nil, errors.New("invalid value for required argument 'HelmChartsBaseUrl'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	var resource RemoteHelmRepository
	err := ctx.RegisterResource("artifactory:index/remoteHelmRepository:RemoteHelmRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemoteHelmRepository gets an existing RemoteHelmRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemoteHelmRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemoteHelmRepositoryState, opts ...pulumi.ResourceOption) (*RemoteHelmRepository, error) {
	var resource RemoteHelmRepository
	err := ctx.ReadResource("artifactory:index/remoteHelmRepository:RemoteHelmRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemoteHelmRepository resources.
type remoteHelmRepositoryState struct {
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
	BypassHeadRequests     *bool                                       `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate   *string                                     `pulumi:"clientTlsCertificate"`
	ContentSynchronisation *RemoteHelmRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                     `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement *bool   `pulumi:"enableCookieManagement"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs *int  `pulumi:"failedRetrievalCachePeriodSecs"`
	HardFail                       *bool `pulumi:"hardFail"`
	// - No documentation is available. Hopefully you know what this means
	HelmChartsBaseUrl *string `pulumi:"helmChartsBaseUrl"`
	IncludesPattern   *string `pulumi:"includesPattern"`
	// The repository identifier. Must be unique system-wide
	Key          *string `pulumi:"key"`
	LocalAddress *string `pulumi:"localAddress"`
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

type RemoteHelmRepositoryState struct {
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
	ContentSynchronisation RemoteHelmRepositoryContentSynchronisationPtrInput
	Description            pulumi.StringPtrInput
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	// Deprecated: This field is not returned in a get payload but is offered on the UI. It's inserted here for inclusive and informational reasons. It does not function
	FailedRetrievalCachePeriodSecs pulumi.IntPtrInput
	HardFail                       pulumi.BoolPtrInput
	// - No documentation is available. Hopefully you know what this means
	HelmChartsBaseUrl pulumi.StringPtrInput
	IncludesPattern   pulumi.StringPtrInput
	// The repository identifier. Must be unique system-wide
	Key          pulumi.StringPtrInput
	LocalAddress pulumi.StringPtrInput
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

func (RemoteHelmRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteHelmRepositoryState)(nil)).Elem()
}

type remoteHelmRepositoryArgs struct {
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
	BypassHeadRequests     *bool                                       `pulumi:"bypassHeadRequests"`
	ClientTlsCertificate   *string                                     `pulumi:"clientTlsCertificate"`
	ContentSynchronisation *RemoteHelmRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                     `pulumi:"description"`
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement *bool   `pulumi:"enableCookieManagement"`
	ExcludesPattern        *string `pulumi:"excludesPattern"`
	HardFail               *bool   `pulumi:"hardFail"`
	// - No documentation is available. Hopefully you know what this means
	HelmChartsBaseUrl string  `pulumi:"helmChartsBaseUrl"`
	IncludesPattern   *string `pulumi:"includesPattern"`
	// The repository identifier. Must be unique system-wide
	Key          string  `pulumi:"key"`
	LocalAddress *string `pulumi:"localAddress"`
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

// The set of arguments for constructing a RemoteHelmRepository resource.
type RemoteHelmRepositoryArgs struct {
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
	ContentSynchronisation RemoteHelmRepositoryContentSynchronisationPtrInput
	Description            pulumi.StringPtrInput
	// Enables cookie management if the remote repository uses cookies to manage client state.
	EnableCookieManagement pulumi.BoolPtrInput
	ExcludesPattern        pulumi.StringPtrInput
	HardFail               pulumi.BoolPtrInput
	// - No documentation is available. Hopefully you know what this means
	HelmChartsBaseUrl pulumi.StringInput
	IncludesPattern   pulumi.StringPtrInput
	// The repository identifier. Must be unique system-wide
	Key          pulumi.StringInput
	LocalAddress pulumi.StringPtrInput
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

func (RemoteHelmRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remoteHelmRepositoryArgs)(nil)).Elem()
}

type RemoteHelmRepositoryInput interface {
	pulumi.Input

	ToRemoteHelmRepositoryOutput() RemoteHelmRepositoryOutput
	ToRemoteHelmRepositoryOutputWithContext(ctx context.Context) RemoteHelmRepositoryOutput
}

func (*RemoteHelmRepository) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteHelmRepository)(nil))
}

func (i *RemoteHelmRepository) ToRemoteHelmRepositoryOutput() RemoteHelmRepositoryOutput {
	return i.ToRemoteHelmRepositoryOutputWithContext(context.Background())
}

func (i *RemoteHelmRepository) ToRemoteHelmRepositoryOutputWithContext(ctx context.Context) RemoteHelmRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteHelmRepositoryOutput)
}

func (i *RemoteHelmRepository) ToRemoteHelmRepositoryPtrOutput() RemoteHelmRepositoryPtrOutput {
	return i.ToRemoteHelmRepositoryPtrOutputWithContext(context.Background())
}

func (i *RemoteHelmRepository) ToRemoteHelmRepositoryPtrOutputWithContext(ctx context.Context) RemoteHelmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteHelmRepositoryPtrOutput)
}

type RemoteHelmRepositoryPtrInput interface {
	pulumi.Input

	ToRemoteHelmRepositoryPtrOutput() RemoteHelmRepositoryPtrOutput
	ToRemoteHelmRepositoryPtrOutputWithContext(ctx context.Context) RemoteHelmRepositoryPtrOutput
}

type remoteHelmRepositoryPtrType RemoteHelmRepositoryArgs

func (*remoteHelmRepositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteHelmRepository)(nil))
}

func (i *remoteHelmRepositoryPtrType) ToRemoteHelmRepositoryPtrOutput() RemoteHelmRepositoryPtrOutput {
	return i.ToRemoteHelmRepositoryPtrOutputWithContext(context.Background())
}

func (i *remoteHelmRepositoryPtrType) ToRemoteHelmRepositoryPtrOutputWithContext(ctx context.Context) RemoteHelmRepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteHelmRepositoryPtrOutput)
}

// RemoteHelmRepositoryArrayInput is an input type that accepts RemoteHelmRepositoryArray and RemoteHelmRepositoryArrayOutput values.
// You can construct a concrete instance of `RemoteHelmRepositoryArrayInput` via:
//
//          RemoteHelmRepositoryArray{ RemoteHelmRepositoryArgs{...} }
type RemoteHelmRepositoryArrayInput interface {
	pulumi.Input

	ToRemoteHelmRepositoryArrayOutput() RemoteHelmRepositoryArrayOutput
	ToRemoteHelmRepositoryArrayOutputWithContext(context.Context) RemoteHelmRepositoryArrayOutput
}

type RemoteHelmRepositoryArray []RemoteHelmRepositoryInput

func (RemoteHelmRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RemoteHelmRepository)(nil)).Elem()
}

func (i RemoteHelmRepositoryArray) ToRemoteHelmRepositoryArrayOutput() RemoteHelmRepositoryArrayOutput {
	return i.ToRemoteHelmRepositoryArrayOutputWithContext(context.Background())
}

func (i RemoteHelmRepositoryArray) ToRemoteHelmRepositoryArrayOutputWithContext(ctx context.Context) RemoteHelmRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteHelmRepositoryArrayOutput)
}

// RemoteHelmRepositoryMapInput is an input type that accepts RemoteHelmRepositoryMap and RemoteHelmRepositoryMapOutput values.
// You can construct a concrete instance of `RemoteHelmRepositoryMapInput` via:
//
//          RemoteHelmRepositoryMap{ "key": RemoteHelmRepositoryArgs{...} }
type RemoteHelmRepositoryMapInput interface {
	pulumi.Input

	ToRemoteHelmRepositoryMapOutput() RemoteHelmRepositoryMapOutput
	ToRemoteHelmRepositoryMapOutputWithContext(context.Context) RemoteHelmRepositoryMapOutput
}

type RemoteHelmRepositoryMap map[string]RemoteHelmRepositoryInput

func (RemoteHelmRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RemoteHelmRepository)(nil)).Elem()
}

func (i RemoteHelmRepositoryMap) ToRemoteHelmRepositoryMapOutput() RemoteHelmRepositoryMapOutput {
	return i.ToRemoteHelmRepositoryMapOutputWithContext(context.Background())
}

func (i RemoteHelmRepositoryMap) ToRemoteHelmRepositoryMapOutputWithContext(ctx context.Context) RemoteHelmRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoteHelmRepositoryMapOutput)
}

type RemoteHelmRepositoryOutput struct{ *pulumi.OutputState }

func (RemoteHelmRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoteHelmRepository)(nil))
}

func (o RemoteHelmRepositoryOutput) ToRemoteHelmRepositoryOutput() RemoteHelmRepositoryOutput {
	return o
}

func (o RemoteHelmRepositoryOutput) ToRemoteHelmRepositoryOutputWithContext(ctx context.Context) RemoteHelmRepositoryOutput {
	return o
}

func (o RemoteHelmRepositoryOutput) ToRemoteHelmRepositoryPtrOutput() RemoteHelmRepositoryPtrOutput {
	return o.ToRemoteHelmRepositoryPtrOutputWithContext(context.Background())
}

func (o RemoteHelmRepositoryOutput) ToRemoteHelmRepositoryPtrOutputWithContext(ctx context.Context) RemoteHelmRepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RemoteHelmRepository) *RemoteHelmRepository {
		return &v
	}).(RemoteHelmRepositoryPtrOutput)
}

type RemoteHelmRepositoryPtrOutput struct{ *pulumi.OutputState }

func (RemoteHelmRepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemoteHelmRepository)(nil))
}

func (o RemoteHelmRepositoryPtrOutput) ToRemoteHelmRepositoryPtrOutput() RemoteHelmRepositoryPtrOutput {
	return o
}

func (o RemoteHelmRepositoryPtrOutput) ToRemoteHelmRepositoryPtrOutputWithContext(ctx context.Context) RemoteHelmRepositoryPtrOutput {
	return o
}

func (o RemoteHelmRepositoryPtrOutput) Elem() RemoteHelmRepositoryOutput {
	return o.ApplyT(func(v *RemoteHelmRepository) RemoteHelmRepository {
		if v != nil {
			return *v
		}
		var ret RemoteHelmRepository
		return ret
	}).(RemoteHelmRepositoryOutput)
}

type RemoteHelmRepositoryArrayOutput struct{ *pulumi.OutputState }

func (RemoteHelmRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RemoteHelmRepository)(nil))
}

func (o RemoteHelmRepositoryArrayOutput) ToRemoteHelmRepositoryArrayOutput() RemoteHelmRepositoryArrayOutput {
	return o
}

func (o RemoteHelmRepositoryArrayOutput) ToRemoteHelmRepositoryArrayOutputWithContext(ctx context.Context) RemoteHelmRepositoryArrayOutput {
	return o
}

func (o RemoteHelmRepositoryArrayOutput) Index(i pulumi.IntInput) RemoteHelmRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RemoteHelmRepository {
		return vs[0].([]RemoteHelmRepository)[vs[1].(int)]
	}).(RemoteHelmRepositoryOutput)
}

type RemoteHelmRepositoryMapOutput struct{ *pulumi.OutputState }

func (RemoteHelmRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RemoteHelmRepository)(nil))
}

func (o RemoteHelmRepositoryMapOutput) ToRemoteHelmRepositoryMapOutput() RemoteHelmRepositoryMapOutput {
	return o
}

func (o RemoteHelmRepositoryMapOutput) ToRemoteHelmRepositoryMapOutputWithContext(ctx context.Context) RemoteHelmRepositoryMapOutput {
	return o
}

func (o RemoteHelmRepositoryMapOutput) MapIndex(k pulumi.StringInput) RemoteHelmRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RemoteHelmRepository {
		return vs[0].(map[string]RemoteHelmRepository)[vs[1].(string)]
	}).(RemoteHelmRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteHelmRepositoryInput)(nil)).Elem(), &RemoteHelmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteHelmRepositoryPtrInput)(nil)).Elem(), &RemoteHelmRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteHelmRepositoryArrayInput)(nil)).Elem(), RemoteHelmRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RemoteHelmRepositoryMapInput)(nil)).Elem(), RemoteHelmRepositoryMap{})
	pulumi.RegisterOutputType(RemoteHelmRepositoryOutput{})
	pulumi.RegisterOutputType(RemoteHelmRepositoryPtrOutput{})
	pulumi.RegisterOutputType(RemoteHelmRepositoryArrayOutput{})
	pulumi.RegisterOutputType(RemoteHelmRepositoryMapOutput{})
}
