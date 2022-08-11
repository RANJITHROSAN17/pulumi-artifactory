# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('artifactory')


class _ExportableConfig(types.ModuleType):
    @property
    def access_token(self) -> Optional[str]:
        """
        This is a access token that can be given to you by your admin under `Identity and Access`. If not set, the 'api_key'
        attribute value will be used.
        """
        return __config__.get('accessToken')

    @property
    def api_key(self) -> Optional[str]:
        """
        API token. Projects functionality will not work with any auth method other than access tokens
        """
        return __config__.get('apiKey')

    @property
    def check_license(self) -> bool:
        """
        Toggle for pre-flight checking of Artifactory Pro and Enterprise license. Default to `true`.
        """
        return __config__.get_bool('checkLicense') or False

    @property
    def url(self) -> Optional[str]:
        return __config__.get('url')

