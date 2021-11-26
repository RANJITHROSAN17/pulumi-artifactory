// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface AccessTokenAdminToken {
    instanceId: string;
}

export interface OauthSettingsOauthProvider {
    /**
     * OAuth user info endpoint for the IdP.
     */
    apiUrl: string;
    /**
     * OAuth authorization endpoint for the IdP.
     */
    authUrl: string;
    /**
     * OAuth client ID configured on the IdP.
     */
    clientId: string;
    /**
     * OAuth client secret configured on the IdP.
     */
    clientSecret: string;
    /**
     * Enable the Artifactory OAuth provider.  Default value is `true`.
     */
    enabled?: boolean;
    /**
     * Name of the Artifactory OAuth provider.
     */
    name: string;
    /**
     * OAuth token endpoint for the IdP.
     */
    tokenUrl: string;
    /**
     * Type of OAuth provider. (e.g., `github`, `google`, `cloudfoundry`, or `openId`)
     */
    type: string;
}

export interface PermissionTargetBuild {
    /**
     * -
     */
    actions?: outputs.PermissionTargetBuildActions;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: string[];
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: string[];
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: string[];
}

export interface PermissionTargetBuildActions {
    /**
     * Groups this permission applies for.
     */
    groups?: outputs.PermissionTargetBuildActionsGroup[];
    /**
     * Users this permission target applies for.
     */
    users?: outputs.PermissionTargetBuildActionsUser[];
}

export interface PermissionTargetBuildActionsGroup {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetBuildActionsUser {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetReleaseBundle {
    /**
     * -
     */
    actions?: outputs.PermissionTargetReleaseBundleActions;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: string[];
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: string[];
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: string[];
}

export interface PermissionTargetReleaseBundleActions {
    /**
     * Groups this permission applies for.
     */
    groups?: outputs.PermissionTargetReleaseBundleActionsGroup[];
    /**
     * Users this permission target applies for.
     */
    users?: outputs.PermissionTargetReleaseBundleActionsUser[];
}

export interface PermissionTargetReleaseBundleActionsGroup {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetReleaseBundleActionsUser {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetRepo {
    /**
     * -
     */
    actions?: outputs.PermissionTargetRepoActions;
    /**
     * Pattern of artifacts to exclude
     */
    excludesPatterns?: string[];
    /**
     * Pattern of artifacts to include
     */
    includesPatterns?: string[];
    /**
     * List of repositories this permission target is applicable for
     */
    repositories: string[];
}

export interface PermissionTargetRepoActions {
    /**
     * Groups this permission applies for.
     */
    groups?: outputs.PermissionTargetRepoActionsGroup[];
    /**
     * Users this permission target applies for.
     */
    users?: outputs.PermissionTargetRepoActionsUser[];
}

export interface PermissionTargetRepoActionsGroup {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetRepoActionsUser {
    /**
     * Name of permission
     */
    name: string;
    permissions: string[];
}

export interface PermissionTargetsBuild {
    actions?: outputs.PermissionTargetsBuildActions;
    excludesPatterns?: string[];
    includesPatterns?: string[];
    repositories: string[];
}

export interface PermissionTargetsBuildActions {
    groups?: outputs.PermissionTargetsBuildActionsGroup[];
    users?: outputs.PermissionTargetsBuildActionsUser[];
}

export interface PermissionTargetsBuildActionsGroup {
    name: string;
    permissions: string[];
}

export interface PermissionTargetsBuildActionsUser {
    name: string;
    permissions: string[];
}

export interface PermissionTargetsReleaseBundle {
    actions?: outputs.PermissionTargetsReleaseBundleActions;
    excludesPatterns?: string[];
    includesPatterns?: string[];
    repositories: string[];
}

export interface PermissionTargetsReleaseBundleActions {
    groups?: outputs.PermissionTargetsReleaseBundleActionsGroup[];
    users?: outputs.PermissionTargetsReleaseBundleActionsUser[];
}

export interface PermissionTargetsReleaseBundleActionsGroup {
    name: string;
    permissions: string[];
}

export interface PermissionTargetsReleaseBundleActionsUser {
    name: string;
    permissions: string[];
}

export interface PermissionTargetsRepo {
    actions?: outputs.PermissionTargetsRepoActions;
    excludesPatterns?: string[];
    includesPatterns?: string[];
    repositories: string[];
}

export interface PermissionTargetsRepoActions {
    groups?: outputs.PermissionTargetsRepoActionsGroup[];
    users?: outputs.PermissionTargetsRepoActionsUser[];
}

export interface PermissionTargetsRepoActionsGroup {
    name: string;
    permissions: string[];
}

export interface PermissionTargetsRepoActionsUser {
    name: string;
    permissions: string[];
}

export interface PushReplicationReplication {
    enabled: boolean;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    password: string;
    pathPrefix?: string;
    socketTimeoutMillis: number;
    syncDeletes: boolean;
    syncProperties: boolean;
    syncStatistics: boolean;
    url?: string;
    username?: string;
}

export interface RemoteCargoRepositoryContentSynchronisation {
    enabled?: boolean;
}

export interface RemoteDockerRepositoryContentSynchronisation {
    enabled?: boolean;
}

export interface RemoteHelmRepositoryContentSynchronisation {
    enabled?: boolean;
}

export interface RemoteNpmRepositoryContentSynchronisation {
    enabled?: boolean;
}

export interface RemoteRepositoryContentSynchronisation {
    enabled?: boolean;
}

export interface ReplicationConfigReplication {
    enabled: boolean;
    /**
     * Requires password encryption to be turned off `POST /api/system/decrypt`
     */
    password: string;
    pathPrefix?: string;
    socketTimeoutMillis: number;
    syncDeletes: boolean;
    syncProperties: boolean;
    syncStatistics: boolean;
    url?: string;
    username?: string;
}

export interface XrayPolicyRule {
    /**
     * (Required) Nested block describing the actions to be applied by the policy. Described below.
     */
    actions?: outputs.XrayPolicyRuleActions;
    /**
     * (Required) Nested block describing the criteria for the policy. Described below.
     */
    criteria: outputs.XrayPolicyRuleCriteria;
    /**
     * (Required) Name of the rule
     */
    name: string;
    /**
     * (Required) Integer describing the rule priority
     */
    priority: number;
}

export interface XrayPolicyRuleActions {
    /**
     * (Optional) Nested block describing artifacts that should be blocked for download if a violation is triggered. Described below.
     */
    blockDownload: outputs.XrayPolicyRuleActionsBlockDownload;
    /**
     * (Optional) The severity of violation to be triggered if the `criteria` are met.
     */
    customSeverity?: string;
    /**
     * (Optional) Whether or not the related CI build should be marked as failed if a violation is triggered. This option is only available when the policy is applied to an `xrayWatch` resource with a `type` of `builds`.
     */
    failBuild?: boolean;
    /**
     * (Optional) A list of email addressed that will get emailed when a violation is triggered.
     */
    mails?: string[];
    /**
     * (Optional) A list of Xray-configured webhook URLs to be invoked if a violation is triggered.
     */
    webhooks?: string[];
}

export interface XrayPolicyRuleActionsBlockDownload {
    /**
     * Whether or not to block download of artifacts that meet the artifact and severity `filters` for the associated `xrayWatch` resource.
     */
    active: boolean;
    /**
     * Whether or not to block download of artifacts that meet the artifact `filters` for the associated `xrayWatch` resource but have not been scanned yet.
     */
    unscanned: boolean;
}

export interface XrayPolicyRuleCriteria {
    /**
     * (Optional) Whether or not to allow components whose license cannot be determined (`true` or `false`).
     */
    allowUnknown?: boolean;
    /**
     * (Optional) A list of OSS license names that may be attached to a component.
     */
    allowedLicenses?: string[];
    /**
     * (Optional) A list of OSS license names that may not be attached to a component.
     */
    bannedLicenses?: string[];
    /**
     * (Optional) Nested block describing a CVS score range to be impacted. Defined below.
     */
    cvssRange?: outputs.XrayPolicyRuleCriteriaCvssRange;
    /**
     * (Optional) The minimum security vulnerability severity that will be impacted by the policy.
     */
    minSeverity?: string;
}

export interface XrayPolicyRuleCriteriaCvssRange {
    /**
     * (Required) The beginning of the range of CVS scores (from 1-10) to flag.
     */
    from: number;
    /**
     * (Required) The end of the range of CVS scores (from 1-10) to flag.
     */
    to: number;
}

export interface XrayWatchAssignedPolicy {
    /**
     * The name of the policy that will be applied
     */
    name: string;
    /**
     * The type of the policy
     */
    type: string;
}

export interface XrayWatchResource {
    /**
     * The ID number of a binary manager resource
     */
    binMgrId?: string;
    /**
     * Nested argument describing filters to be applied. Defined below.
     */
    filters?: outputs.XrayWatchResourceFilter[];
    /**
     * A name describing the resource
     */
    name: string;
    /**
     * Type of repository (e.g. local or remote)
     */
    repoType: string;
    /**
     * Type of resource to be watched
     */
    type: string;
}

export interface XrayWatchResourceFilter {
    /**
     * The type of filter, such as `regex` or `package-type`
     */
    type: string;
    /**
     * The value of the filter, such as the text of the regex or name of the package type
     */
    value: string;
}

