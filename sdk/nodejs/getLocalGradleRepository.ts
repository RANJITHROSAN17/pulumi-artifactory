// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getLocalGradleRepository(args: GetLocalGradleRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetLocalGradleRepositoryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("artifactory:index/getLocalGradleRepository:getLocalGradleRepository", {
        "archiveBrowsingEnabled": args.archiveBrowsingEnabled,
        "blackedOut": args.blackedOut,
        "cdnRedirect": args.cdnRedirect,
        "checksumPolicyType": args.checksumPolicyType,
        "description": args.description,
        "downloadDirect": args.downloadDirect,
        "excludesPattern": args.excludesPattern,
        "handleReleases": args.handleReleases,
        "handleSnapshots": args.handleSnapshots,
        "includesPattern": args.includesPattern,
        "key": args.key,
        "maxUniqueSnapshots": args.maxUniqueSnapshots,
        "notes": args.notes,
        "priorityResolution": args.priorityResolution,
        "projectEnvironments": args.projectEnvironments,
        "projectKey": args.projectKey,
        "propertySets": args.propertySets,
        "repoLayoutRef": args.repoLayoutRef,
        "snapshotVersionBehavior": args.snapshotVersionBehavior,
        "suppressPomConsistencyChecks": args.suppressPomConsistencyChecks,
        "xrayIndex": args.xrayIndex,
    }, opts);
}

/**
 * A collection of arguments for invoking getLocalGradleRepository.
 */
export interface GetLocalGradleRepositoryArgs {
    archiveBrowsingEnabled?: boolean;
    blackedOut?: boolean;
    cdnRedirect?: boolean;
    checksumPolicyType?: string;
    description?: string;
    downloadDirect?: boolean;
    excludesPattern?: string;
    handleReleases?: boolean;
    handleSnapshots?: boolean;
    includesPattern?: string;
    key: string;
    maxUniqueSnapshots?: number;
    notes?: string;
    priorityResolution?: boolean;
    projectEnvironments?: string[];
    projectKey?: string;
    propertySets?: string[];
    repoLayoutRef?: string;
    snapshotVersionBehavior?: string;
    suppressPomConsistencyChecks?: boolean;
    xrayIndex?: boolean;
}

/**
 * A collection of values returned by getLocalGradleRepository.
 */
export interface GetLocalGradleRepositoryResult {
    readonly archiveBrowsingEnabled?: boolean;
    readonly blackedOut?: boolean;
    readonly cdnRedirect?: boolean;
    readonly checksumPolicyType?: string;
    readonly description?: string;
    readonly downloadDirect?: boolean;
    readonly excludesPattern: string;
    readonly handleReleases?: boolean;
    readonly handleSnapshots?: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includesPattern: string;
    readonly key: string;
    readonly maxUniqueSnapshots?: number;
    readonly notes?: string;
    readonly packageType: string;
    readonly priorityResolution?: boolean;
    readonly projectEnvironments: string[];
    readonly projectKey?: string;
    readonly propertySets?: string[];
    readonly repoLayoutRef?: string;
    readonly snapshotVersionBehavior?: string;
    readonly suppressPomConsistencyChecks?: boolean;
    readonly xrayIndex?: boolean;
}
export function getLocalGradleRepositoryOutput(args: GetLocalGradleRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocalGradleRepositoryResult> {
    return pulumi.output(args).apply((a: any) => getLocalGradleRepository(a, opts))
}

/**
 * A collection of arguments for invoking getLocalGradleRepository.
 */
export interface GetLocalGradleRepositoryOutputArgs {
    archiveBrowsingEnabled?: pulumi.Input<boolean>;
    blackedOut?: pulumi.Input<boolean>;
    cdnRedirect?: pulumi.Input<boolean>;
    checksumPolicyType?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    downloadDirect?: pulumi.Input<boolean>;
    excludesPattern?: pulumi.Input<string>;
    handleReleases?: pulumi.Input<boolean>;
    handleSnapshots?: pulumi.Input<boolean>;
    includesPattern?: pulumi.Input<string>;
    key: pulumi.Input<string>;
    maxUniqueSnapshots?: pulumi.Input<number>;
    notes?: pulumi.Input<string>;
    priorityResolution?: pulumi.Input<boolean>;
    projectEnvironments?: pulumi.Input<pulumi.Input<string>[]>;
    projectKey?: pulumi.Input<string>;
    propertySets?: pulumi.Input<pulumi.Input<string>[]>;
    repoLayoutRef?: pulumi.Input<string>;
    snapshotVersionBehavior?: pulumi.Input<string>;
    suppressPomConsistencyChecks?: pulumi.Input<boolean>;
    xrayIndex?: pulumi.Input<boolean>;
}
