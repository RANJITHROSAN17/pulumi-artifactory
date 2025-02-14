# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetPermissionTargetResult',
    'AwaitableGetPermissionTargetResult',
    'get_permission_target',
    'get_permission_target_output',
]

@pulumi.output_type
class GetPermissionTargetResult:
    """
    A collection of values returned by getPermissionTarget.
    """
    def __init__(__self__, build=None, id=None, name=None, release_bundle=None, repo=None):
        if build and not isinstance(build, dict):
            raise TypeError("Expected argument 'build' to be a dict")
        pulumi.set(__self__, "build", build)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if release_bundle and not isinstance(release_bundle, dict):
            raise TypeError("Expected argument 'release_bundle' to be a dict")
        pulumi.set(__self__, "release_bundle", release_bundle)
        if repo and not isinstance(repo, dict):
            raise TypeError("Expected argument 'repo' to be a dict")
        pulumi.set(__self__, "repo", repo)

    @property
    @pulumi.getter
    def build(self) -> Optional['outputs.GetPermissionTargetBuildResult']:
        """
        Same as repo but for artifactory-build-info permissions.
        """
        return pulumi.get(self, "build")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="releaseBundle")
    def release_bundle(self) -> Optional['outputs.GetPermissionTargetReleaseBundleResult']:
        """
        Same as repo but for release-bundles permissions.
        """
        return pulumi.get(self, "release_bundle")

    @property
    @pulumi.getter
    def repo(self) -> Optional['outputs.GetPermissionTargetRepoResult']:
        """
        Repository permission configuration.
        """
        return pulumi.get(self, "repo")


class AwaitableGetPermissionTargetResult(GetPermissionTargetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPermissionTargetResult(
            build=self.build,
            id=self.id,
            name=self.name,
            release_bundle=self.release_bundle,
            repo=self.repo)


def get_permission_target(build: Optional[pulumi.InputType['GetPermissionTargetBuildArgs']] = None,
                          name: Optional[str] = None,
                          release_bundle: Optional[pulumi.InputType['GetPermissionTargetReleaseBundleArgs']] = None,
                          repo: Optional[pulumi.InputType['GetPermissionTargetRepoArgs']] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPermissionTargetResult:
    """
    ## # Artifactory Permission Target Data Source

    Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    target1 = artifactory.get_permission_target(name="my_permission")
    ```


    :param pulumi.InputType['GetPermissionTargetBuildArgs'] build: Same as repo but for artifactory-build-info permissions.
    :param str name: Name of the permission target.
    :param pulumi.InputType['GetPermissionTargetReleaseBundleArgs'] release_bundle: Same as repo but for release-bundles permissions.
    :param pulumi.InputType['GetPermissionTargetRepoArgs'] repo: Repository permission configuration.
    """
    __args__ = dict()
    __args__['build'] = build
    __args__['name'] = name
    __args__['releaseBundle'] = release_bundle
    __args__['repo'] = repo
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('artifactory:index/getPermissionTarget:getPermissionTarget', __args__, opts=opts, typ=GetPermissionTargetResult).value

    return AwaitableGetPermissionTargetResult(
        build=__ret__.build,
        id=__ret__.id,
        name=__ret__.name,
        release_bundle=__ret__.release_bundle,
        repo=__ret__.repo)


@_utilities.lift_output_func(get_permission_target)
def get_permission_target_output(build: Optional[pulumi.Input[Optional[pulumi.InputType['GetPermissionTargetBuildArgs']]]] = None,
                                 name: Optional[pulumi.Input[str]] = None,
                                 release_bundle: Optional[pulumi.Input[Optional[pulumi.InputType['GetPermissionTargetReleaseBundleArgs']]]] = None,
                                 repo: Optional[pulumi.Input[Optional[pulumi.InputType['GetPermissionTargetRepoArgs']]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPermissionTargetResult]:
    """
    ## # Artifactory Permission Target Data Source

    Provides an Artifactory permission target data source. This can be used to read the configuration of permission targets in artifactory.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_artifactory as artifactory

    target1 = artifactory.get_permission_target(name="my_permission")
    ```


    :param pulumi.InputType['GetPermissionTargetBuildArgs'] build: Same as repo but for artifactory-build-info permissions.
    :param str name: Name of the permission target.
    :param pulumi.InputType['GetPermissionTargetReleaseBundleArgs'] release_bundle: Same as repo but for release-bundles permissions.
    :param pulumi.InputType['GetPermissionTargetRepoArgs'] repo: Repository permission configuration.
    """
    ...
