# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SamlSettingsArgs', 'SamlSettings']

@pulumi.input_type
class SamlSettingsArgs:
    def __init__(__self__, *,
                 login_url: pulumi.Input[str],
                 logout_url: pulumi.Input[str],
                 service_provider_name: pulumi.Input[str],
                 allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
                 auto_redirect: Optional[pulumi.Input[bool]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 email_attribute: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 group_attribute: Optional[pulumi.Input[str]] = None,
                 no_auto_user_creation: Optional[pulumi.Input[bool]] = None,
                 sync_groups: Optional[pulumi.Input[bool]] = None,
                 verify_audience_restriction: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SamlSettings resource.
        :param pulumi.Input[str] login_url: Service provider login url configured on the IdP.
        :param pulumi.Input[str] logout_url: Service provider logout url, or where to redirect after user logs out.
        :param pulumi.Input[str] service_provider_name: Name of the service provider configured on the .
        :param pulumi.Input[bool] allow_user_to_access_profile: Allow persisted users to access their profile.  Default value is `true`.
        :param pulumi.Input[bool] auto_redirect: Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        :param pulumi.Input[str] certificate: SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
        :param pulumi.Input[str] email_attribute: Name of the attribute in the SAML response from the IdP that contains the user's email.
        :param pulumi.Input[bool] enable: Enable SAML SSO.  Default value is `true`.
        :param pulumi.Input[str] group_attribute: Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
        :param pulumi.Input[bool] no_auto_user_creation: When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        :param pulumi.Input[bool] sync_groups: Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        :param pulumi.Input[bool] verify_audience_restriction: Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        """
        pulumi.set(__self__, "login_url", login_url)
        pulumi.set(__self__, "logout_url", logout_url)
        pulumi.set(__self__, "service_provider_name", service_provider_name)
        if allow_user_to_access_profile is not None:
            pulumi.set(__self__, "allow_user_to_access_profile", allow_user_to_access_profile)
        if auto_redirect is not None:
            pulumi.set(__self__, "auto_redirect", auto_redirect)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if email_attribute is not None:
            pulumi.set(__self__, "email_attribute", email_attribute)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if group_attribute is not None:
            pulumi.set(__self__, "group_attribute", group_attribute)
        if no_auto_user_creation is not None:
            pulumi.set(__self__, "no_auto_user_creation", no_auto_user_creation)
        if sync_groups is not None:
            pulumi.set(__self__, "sync_groups", sync_groups)
        if verify_audience_restriction is not None:
            pulumi.set(__self__, "verify_audience_restriction", verify_audience_restriction)

    @property
    @pulumi.getter(name="loginUrl")
    def login_url(self) -> pulumi.Input[str]:
        """
        Service provider login url configured on the IdP.
        """
        return pulumi.get(self, "login_url")

    @login_url.setter
    def login_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "login_url", value)

    @property
    @pulumi.getter(name="logoutUrl")
    def logout_url(self) -> pulumi.Input[str]:
        """
        Service provider logout url, or where to redirect after user logs out.
        """
        return pulumi.get(self, "logout_url")

    @logout_url.setter
    def logout_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "logout_url", value)

    @property
    @pulumi.getter(name="serviceProviderName")
    def service_provider_name(self) -> pulumi.Input[str]:
        """
        Name of the service provider configured on the .
        """
        return pulumi.get(self, "service_provider_name")

    @service_provider_name.setter
    def service_provider_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_provider_name", value)

    @property
    @pulumi.getter(name="allowUserToAccessProfile")
    def allow_user_to_access_profile(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow persisted users to access their profile.  Default value is `true`.
        """
        return pulumi.get(self, "allow_user_to_access_profile")

    @allow_user_to_access_profile.setter
    def allow_user_to_access_profile(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_user_to_access_profile", value)

    @property
    @pulumi.getter(name="autoRedirect")
    def auto_redirect(self) -> Optional[pulumi.Input[bool]]:
        """
        Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        """
        return pulumi.get(self, "auto_redirect")

    @auto_redirect.setter
    def auto_redirect(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_redirect", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="emailAttribute")
    def email_attribute(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the attribute in the SAML response from the IdP that contains the user's email.
        """
        return pulumi.get(self, "email_attribute")

    @email_attribute.setter
    def email_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_attribute", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SAML SSO.  Default value is `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="groupAttribute")
    def group_attribute(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
        """
        return pulumi.get(self, "group_attribute")

    @group_attribute.setter
    def group_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_attribute", value)

    @property
    @pulumi.getter(name="noAutoUserCreation")
    def no_auto_user_creation(self) -> Optional[pulumi.Input[bool]]:
        """
        When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        """
        return pulumi.get(self, "no_auto_user_creation")

    @no_auto_user_creation.setter
    def no_auto_user_creation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_auto_user_creation", value)

    @property
    @pulumi.getter(name="syncGroups")
    def sync_groups(self) -> Optional[pulumi.Input[bool]]:
        """
        Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        """
        return pulumi.get(self, "sync_groups")

    @sync_groups.setter
    def sync_groups(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sync_groups", value)

    @property
    @pulumi.getter(name="verifyAudienceRestriction")
    def verify_audience_restriction(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        """
        return pulumi.get(self, "verify_audience_restriction")

    @verify_audience_restriction.setter
    def verify_audience_restriction(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verify_audience_restriction", value)


@pulumi.input_type
class _SamlSettingsState:
    def __init__(__self__, *,
                 allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
                 auto_redirect: Optional[pulumi.Input[bool]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 email_attribute: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 group_attribute: Optional[pulumi.Input[str]] = None,
                 login_url: Optional[pulumi.Input[str]] = None,
                 logout_url: Optional[pulumi.Input[str]] = None,
                 no_auto_user_creation: Optional[pulumi.Input[bool]] = None,
                 service_provider_name: Optional[pulumi.Input[str]] = None,
                 sync_groups: Optional[pulumi.Input[bool]] = None,
                 verify_audience_restriction: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering SamlSettings resources.
        :param pulumi.Input[bool] allow_user_to_access_profile: Allow persisted users to access their profile.  Default value is `true`.
        :param pulumi.Input[bool] auto_redirect: Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        :param pulumi.Input[str] certificate: SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
        :param pulumi.Input[str] email_attribute: Name of the attribute in the SAML response from the IdP that contains the user's email.
        :param pulumi.Input[bool] enable: Enable SAML SSO.  Default value is `true`.
        :param pulumi.Input[str] group_attribute: Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
        :param pulumi.Input[str] login_url: Service provider login url configured on the IdP.
        :param pulumi.Input[str] logout_url: Service provider logout url, or where to redirect after user logs out.
        :param pulumi.Input[bool] no_auto_user_creation: When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        :param pulumi.Input[str] service_provider_name: Name of the service provider configured on the .
        :param pulumi.Input[bool] sync_groups: Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        :param pulumi.Input[bool] verify_audience_restriction: Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        """
        if allow_user_to_access_profile is not None:
            pulumi.set(__self__, "allow_user_to_access_profile", allow_user_to_access_profile)
        if auto_redirect is not None:
            pulumi.set(__self__, "auto_redirect", auto_redirect)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if email_attribute is not None:
            pulumi.set(__self__, "email_attribute", email_attribute)
        if enable is not None:
            pulumi.set(__self__, "enable", enable)
        if group_attribute is not None:
            pulumi.set(__self__, "group_attribute", group_attribute)
        if login_url is not None:
            pulumi.set(__self__, "login_url", login_url)
        if logout_url is not None:
            pulumi.set(__self__, "logout_url", logout_url)
        if no_auto_user_creation is not None:
            pulumi.set(__self__, "no_auto_user_creation", no_auto_user_creation)
        if service_provider_name is not None:
            pulumi.set(__self__, "service_provider_name", service_provider_name)
        if sync_groups is not None:
            pulumi.set(__self__, "sync_groups", sync_groups)
        if verify_audience_restriction is not None:
            pulumi.set(__self__, "verify_audience_restriction", verify_audience_restriction)

    @property
    @pulumi.getter(name="allowUserToAccessProfile")
    def allow_user_to_access_profile(self) -> Optional[pulumi.Input[bool]]:
        """
        Allow persisted users to access their profile.  Default value is `true`.
        """
        return pulumi.get(self, "allow_user_to_access_profile")

    @allow_user_to_access_profile.setter
    def allow_user_to_access_profile(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_user_to_access_profile", value)

    @property
    @pulumi.getter(name="autoRedirect")
    def auto_redirect(self) -> Optional[pulumi.Input[bool]]:
        """
        Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        """
        return pulumi.get(self, "auto_redirect")

    @auto_redirect.setter
    def auto_redirect(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_redirect", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="emailAttribute")
    def email_attribute(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the attribute in the SAML response from the IdP that contains the user's email.
        """
        return pulumi.get(self, "email_attribute")

    @email_attribute.setter
    def email_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email_attribute", value)

    @property
    @pulumi.getter
    def enable(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable SAML SSO.  Default value is `true`.
        """
        return pulumi.get(self, "enable")

    @enable.setter
    def enable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable", value)

    @property
    @pulumi.getter(name="groupAttribute")
    def group_attribute(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
        """
        return pulumi.get(self, "group_attribute")

    @group_attribute.setter
    def group_attribute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_attribute", value)

    @property
    @pulumi.getter(name="loginUrl")
    def login_url(self) -> Optional[pulumi.Input[str]]:
        """
        Service provider login url configured on the IdP.
        """
        return pulumi.get(self, "login_url")

    @login_url.setter
    def login_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "login_url", value)

    @property
    @pulumi.getter(name="logoutUrl")
    def logout_url(self) -> Optional[pulumi.Input[str]]:
        """
        Service provider logout url, or where to redirect after user logs out.
        """
        return pulumi.get(self, "logout_url")

    @logout_url.setter
    def logout_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logout_url", value)

    @property
    @pulumi.getter(name="noAutoUserCreation")
    def no_auto_user_creation(self) -> Optional[pulumi.Input[bool]]:
        """
        When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        """
        return pulumi.get(self, "no_auto_user_creation")

    @no_auto_user_creation.setter
    def no_auto_user_creation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "no_auto_user_creation", value)

    @property
    @pulumi.getter(name="serviceProviderName")
    def service_provider_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the service provider configured on the .
        """
        return pulumi.get(self, "service_provider_name")

    @service_provider_name.setter
    def service_provider_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_provider_name", value)

    @property
    @pulumi.getter(name="syncGroups")
    def sync_groups(self) -> Optional[pulumi.Input[bool]]:
        """
        Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        """
        return pulumi.get(self, "sync_groups")

    @sync_groups.setter
    def sync_groups(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sync_groups", value)

    @property
    @pulumi.getter(name="verifyAudienceRestriction")
    def verify_audience_restriction(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        """
        return pulumi.get(self, "verify_audience_restriction")

    @verify_audience_restriction.setter
    def verify_audience_restriction(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verify_audience_restriction", value)


class SamlSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
                 auto_redirect: Optional[pulumi.Input[bool]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 email_attribute: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 group_attribute: Optional[pulumi.Input[str]] = None,
                 login_url: Optional[pulumi.Input[str]] = None,
                 logout_url: Optional[pulumi.Input[str]] = None,
                 no_auto_user_creation: Optional[pulumi.Input[bool]] = None,
                 service_provider_name: Optional[pulumi.Input[str]] = None,
                 sync_groups: Optional[pulumi.Input[bool]] = None,
                 verify_audience_restriction: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## # Artifactory SAML SSO Settings Resource

        This resource can be used to manage Artifactory's SAML SSO settings.

        Only a single `SamlSettings` resource is meant to be defined.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory SAML SSO settings
        saml = artifactory.SamlSettings("saml",
            allow_user_to_access_profile=True,
            auto_redirect=True,
            certificate="test-certificate",
            email_attribute="email",
            enable=True,
            group_attribute="groups",
            login_url="test-login-url",
            logout_url="test-logout-url",
            no_auto_user_creation=False,
            service_provider_name="okta",
            sync_groups=True,
            verify_audience_restriction=True)
        ```

        ## Import

        Current SAML SSO settings can be imported using `saml_settings` as the `ID`, e.g.

        ```sh
         $ pulumi import artifactory:index/samlSettings:SamlSettings saml saml_settings
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_user_to_access_profile: Allow persisted users to access their profile.  Default value is `true`.
        :param pulumi.Input[bool] auto_redirect: Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        :param pulumi.Input[str] certificate: SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
        :param pulumi.Input[str] email_attribute: Name of the attribute in the SAML response from the IdP that contains the user's email.
        :param pulumi.Input[bool] enable: Enable SAML SSO.  Default value is `true`.
        :param pulumi.Input[str] group_attribute: Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
        :param pulumi.Input[str] login_url: Service provider login url configured on the IdP.
        :param pulumi.Input[str] logout_url: Service provider logout url, or where to redirect after user logs out.
        :param pulumi.Input[bool] no_auto_user_creation: When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        :param pulumi.Input[str] service_provider_name: Name of the service provider configured on the .
        :param pulumi.Input[bool] sync_groups: Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        :param pulumi.Input[bool] verify_audience_restriction: Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SamlSettingsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Artifactory SAML SSO Settings Resource

        This resource can be used to manage Artifactory's SAML SSO settings.

        Only a single `SamlSettings` resource is meant to be defined.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        # Configure Artifactory SAML SSO settings
        saml = artifactory.SamlSettings("saml",
            allow_user_to_access_profile=True,
            auto_redirect=True,
            certificate="test-certificate",
            email_attribute="email",
            enable=True,
            group_attribute="groups",
            login_url="test-login-url",
            logout_url="test-logout-url",
            no_auto_user_creation=False,
            service_provider_name="okta",
            sync_groups=True,
            verify_audience_restriction=True)
        ```

        ## Import

        Current SAML SSO settings can be imported using `saml_settings` as the `ID`, e.g.

        ```sh
         $ pulumi import artifactory:index/samlSettings:SamlSettings saml saml_settings
        ```

        :param str resource_name: The name of the resource.
        :param SamlSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SamlSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
                 auto_redirect: Optional[pulumi.Input[bool]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 email_attribute: Optional[pulumi.Input[str]] = None,
                 enable: Optional[pulumi.Input[bool]] = None,
                 group_attribute: Optional[pulumi.Input[str]] = None,
                 login_url: Optional[pulumi.Input[str]] = None,
                 logout_url: Optional[pulumi.Input[str]] = None,
                 no_auto_user_creation: Optional[pulumi.Input[bool]] = None,
                 service_provider_name: Optional[pulumi.Input[str]] = None,
                 sync_groups: Optional[pulumi.Input[bool]] = None,
                 verify_audience_restriction: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SamlSettingsArgs.__new__(SamlSettingsArgs)

            __props__.__dict__["allow_user_to_access_profile"] = allow_user_to_access_profile
            __props__.__dict__["auto_redirect"] = auto_redirect
            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["email_attribute"] = email_attribute
            __props__.__dict__["enable"] = enable
            __props__.__dict__["group_attribute"] = group_attribute
            if login_url is None and not opts.urn:
                raise TypeError("Missing required property 'login_url'")
            __props__.__dict__["login_url"] = login_url
            if logout_url is None and not opts.urn:
                raise TypeError("Missing required property 'logout_url'")
            __props__.__dict__["logout_url"] = logout_url
            __props__.__dict__["no_auto_user_creation"] = no_auto_user_creation
            if service_provider_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_provider_name'")
            __props__.__dict__["service_provider_name"] = service_provider_name
            __props__.__dict__["sync_groups"] = sync_groups
            __props__.__dict__["verify_audience_restriction"] = verify_audience_restriction
        super(SamlSettings, __self__).__init__(
            'artifactory:index/samlSettings:SamlSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_user_to_access_profile: Optional[pulumi.Input[bool]] = None,
            auto_redirect: Optional[pulumi.Input[bool]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            email_attribute: Optional[pulumi.Input[str]] = None,
            enable: Optional[pulumi.Input[bool]] = None,
            group_attribute: Optional[pulumi.Input[str]] = None,
            login_url: Optional[pulumi.Input[str]] = None,
            logout_url: Optional[pulumi.Input[str]] = None,
            no_auto_user_creation: Optional[pulumi.Input[bool]] = None,
            service_provider_name: Optional[pulumi.Input[str]] = None,
            sync_groups: Optional[pulumi.Input[bool]] = None,
            verify_audience_restriction: Optional[pulumi.Input[bool]] = None) -> 'SamlSettings':
        """
        Get an existing SamlSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_user_to_access_profile: Allow persisted users to access their profile.  Default value is `true`.
        :param pulumi.Input[bool] auto_redirect: Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        :param pulumi.Input[str] certificate: SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
        :param pulumi.Input[str] email_attribute: Name of the attribute in the SAML response from the IdP that contains the user's email.
        :param pulumi.Input[bool] enable: Enable SAML SSO.  Default value is `true`.
        :param pulumi.Input[str] group_attribute: Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
        :param pulumi.Input[str] login_url: Service provider login url configured on the IdP.
        :param pulumi.Input[str] logout_url: Service provider logout url, or where to redirect after user logs out.
        :param pulumi.Input[bool] no_auto_user_creation: When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        :param pulumi.Input[str] service_provider_name: Name of the service provider configured on the .
        :param pulumi.Input[bool] sync_groups: Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        :param pulumi.Input[bool] verify_audience_restriction: Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SamlSettingsState.__new__(_SamlSettingsState)

        __props__.__dict__["allow_user_to_access_profile"] = allow_user_to_access_profile
        __props__.__dict__["auto_redirect"] = auto_redirect
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["email_attribute"] = email_attribute
        __props__.__dict__["enable"] = enable
        __props__.__dict__["group_attribute"] = group_attribute
        __props__.__dict__["login_url"] = login_url
        __props__.__dict__["logout_url"] = logout_url
        __props__.__dict__["no_auto_user_creation"] = no_auto_user_creation
        __props__.__dict__["service_provider_name"] = service_provider_name
        __props__.__dict__["sync_groups"] = sync_groups
        __props__.__dict__["verify_audience_restriction"] = verify_audience_restriction
        return SamlSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowUserToAccessProfile")
    def allow_user_to_access_profile(self) -> pulumi.Output[Optional[bool]]:
        """
        Allow persisted users to access their profile.  Default value is `true`.
        """
        return pulumi.get(self, "allow_user_to_access_profile")

    @property
    @pulumi.getter(name="autoRedirect")
    def auto_redirect(self) -> pulumi.Output[Optional[bool]]:
        """
        Auto redirect to login through the IdP when clicking on Artifactory's login link.  Default value is `false`.
        """
        return pulumi.get(self, "auto_redirect")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[Optional[str]]:
        """
        SAML certificate that contains the public key for the IdP service provider.  Used by Artifactory to verify sign-in requests.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="emailAttribute")
    def email_attribute(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the attribute in the SAML response from the IdP that contains the user's email.
        """
        return pulumi.get(self, "email_attribute")

    @property
    @pulumi.getter
    def enable(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable SAML SSO.  Default value is `true`.
        """
        return pulumi.get(self, "enable")

    @property
    @pulumi.getter(name="groupAttribute")
    def group_attribute(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the attribute in the SAML response from the IdP that contains the user's group memberships.
        """
        return pulumi.get(self, "group_attribute")

    @property
    @pulumi.getter(name="loginUrl")
    def login_url(self) -> pulumi.Output[str]:
        """
        Service provider login url configured on the IdP.
        """
        return pulumi.get(self, "login_url")

    @property
    @pulumi.getter(name="logoutUrl")
    def logout_url(self) -> pulumi.Output[str]:
        """
        Service provider logout url, or where to redirect after user logs out.
        """
        return pulumi.get(self, "logout_url")

    @property
    @pulumi.getter(name="noAutoUserCreation")
    def no_auto_user_creation(self) -> pulumi.Output[Optional[bool]]:
        """
        When automatic user creation is off, authenticated users are not automatically created inside Artifactory. Instead, for every request from an SSO user, the user is temporarily associated with default groups (if such groups are defined), and the permissions for these groups apply. Without auto-user creation, you must manually create the user inside Artifactory to manage user permissions not attached to their default groups. Default value is `false`.
        """
        return pulumi.get(self, "no_auto_user_creation")

    @property
    @pulumi.getter(name="serviceProviderName")
    def service_provider_name(self) -> pulumi.Output[str]:
        """
        Name of the service provider configured on the .
        """
        return pulumi.get(self, "service_provider_name")

    @property
    @pulumi.getter(name="syncGroups")
    def sync_groups(self) -> pulumi.Output[Optional[bool]]:
        """
        Associate user with Artifactory groups based on the `group_attribute` provided in the SAML response from the identity provider.  Default value is `false`.
        """
        return pulumi.get(self, "sync_groups")

    @property
    @pulumi.getter(name="verifyAudienceRestriction")
    def verify_audience_restriction(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable "audience", or who the SAML assertion is intended for.  Ensures that the correct service provider intended for Artifactory is used on the IdP.  Default value is `true`.
        """
        return pulumi.get(self, "verify_audience_restriction")

