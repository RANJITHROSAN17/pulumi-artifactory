# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DebianRepositoryArgs', 'DebianRepository']

@pulumi.input_type
class DebianRepositoryArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 primary_keypair_ref: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 secondary_keypair_ref: Optional[pulumi.Input[str]] = None,
                 trivial_layout: Optional[pulumi.Input[bool]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a DebianRepository resource.
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] index_compression_formats: - If you're creating this repo, then maybe you know?
        :param pulumi.Input[str] primary_keypair_ref: - The RSA key to be used to sign packages
        :param pulumi.Input[str] secondary_keypair_ref: - Not really clear what this does
        :param pulumi.Input[bool] trivial_layout: - Apparently this is a deprecated repo layout
        """
        pulumi.set(__self__, "key", key)
        if archive_browsing_enabled is not None:
            pulumi.set(__self__, "archive_browsing_enabled", archive_browsing_enabled)
        if blacked_out is not None:
            pulumi.set(__self__, "blacked_out", blacked_out)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if download_direct is not None:
            pulumi.set(__self__, "download_direct", download_direct)
        if excludes_pattern is not None:
            pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if includes_pattern is not None:
            pulumi.set(__self__, "includes_pattern", includes_pattern)
        if index_compression_formats is not None:
            pulumi.set(__self__, "index_compression_formats", index_compression_formats)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if primary_keypair_ref is not None:
            pulumi.set(__self__, "primary_keypair_ref", primary_keypair_ref)
        if property_sets is not None:
            pulumi.set(__self__, "property_sets", property_sets)
        if repo_layout_ref is not None:
            pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if secondary_keypair_ref is not None:
            pulumi.set(__self__, "secondary_keypair_ref", secondary_keypair_ref)
        if trivial_layout is not None:
            warnings.warn("""You shouldn't be using this""", DeprecationWarning)
            pulumi.log.warn("""trivial_layout is deprecated: You shouldn't be using this""")
        if trivial_layout is not None:
            pulumi.set(__self__, "trivial_layout", trivial_layout)
        if xray_index is not None:
            pulumi.set(__self__, "xray_index", xray_index)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        - the identity key of the repo
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter(name="archiveBrowsingEnabled")
    def archive_browsing_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        security (e.g., cross-site scripting attacks).
        """
        return pulumi.get(self, "archive_browsing_enabled")

    @archive_browsing_enabled.setter
    def archive_browsing_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "archive_browsing_enabled", value)

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "blacked_out")

    @blacked_out.setter
    def blacked_out(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blacked_out", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "download_direct")

    @download_direct.setter
    def download_direct(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "download_direct", value)

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "excludes_pattern")

    @excludes_pattern.setter
    def excludes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "excludes_pattern", value)

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "includes_pattern")

    @includes_pattern.setter
    def includes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "includes_pattern", value)

    @property
    @pulumi.getter(name="indexCompressionFormats")
    def index_compression_formats(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        - If you're creating this repo, then maybe you know?
        """
        return pulumi.get(self, "index_compression_formats")

    @index_compression_formats.setter
    def index_compression_formats(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "index_compression_formats", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter(name="primaryKeypairRef")
    def primary_keypair_ref(self) -> Optional[pulumi.Input[str]]:
        """
        - The RSA key to be used to sign packages
        """
        return pulumi.get(self, "primary_keypair_ref")

    @primary_keypair_ref.setter
    def primary_keypair_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_keypair_ref", value)

    @property
    @pulumi.getter(name="propertySets")
    def property_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "property_sets")

    @property_sets.setter
    def property_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "property_sets", value)

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repo_layout_ref")

    @repo_layout_ref.setter
    def repo_layout_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_layout_ref", value)

    @property
    @pulumi.getter(name="secondaryKeypairRef")
    def secondary_keypair_ref(self) -> Optional[pulumi.Input[str]]:
        """
        - Not really clear what this does
        """
        return pulumi.get(self, "secondary_keypair_ref")

    @secondary_keypair_ref.setter
    def secondary_keypair_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_keypair_ref", value)

    @property
    @pulumi.getter(name="trivialLayout")
    def trivial_layout(self) -> Optional[pulumi.Input[bool]]:
        """
        - Apparently this is a deprecated repo layout
        """
        return pulumi.get(self, "trivial_layout")

    @trivial_layout.setter
    def trivial_layout(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "trivial_layout", value)

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "xray_index")

    @xray_index.setter
    def xray_index(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_index", value)


@pulumi.input_type
class _DebianRepositoryState:
    def __init__(__self__, *,
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 package_type: Optional[pulumi.Input[str]] = None,
                 primary_keypair_ref: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 secondary_keypair_ref: Optional[pulumi.Input[str]] = None,
                 trivial_layout: Optional[pulumi.Input[bool]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering DebianRepository resources.
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] index_compression_formats: - If you're creating this repo, then maybe you know?
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[str] primary_keypair_ref: - The RSA key to be used to sign packages
        :param pulumi.Input[str] secondary_keypair_ref: - Not really clear what this does
        :param pulumi.Input[bool] trivial_layout: - Apparently this is a deprecated repo layout
        """
        if archive_browsing_enabled is not None:
            pulumi.set(__self__, "archive_browsing_enabled", archive_browsing_enabled)
        if blacked_out is not None:
            pulumi.set(__self__, "blacked_out", blacked_out)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if download_direct is not None:
            pulumi.set(__self__, "download_direct", download_direct)
        if excludes_pattern is not None:
            pulumi.set(__self__, "excludes_pattern", excludes_pattern)
        if includes_pattern is not None:
            pulumi.set(__self__, "includes_pattern", includes_pattern)
        if index_compression_formats is not None:
            pulumi.set(__self__, "index_compression_formats", index_compression_formats)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if package_type is not None:
            pulumi.set(__self__, "package_type", package_type)
        if primary_keypair_ref is not None:
            pulumi.set(__self__, "primary_keypair_ref", primary_keypair_ref)
        if property_sets is not None:
            pulumi.set(__self__, "property_sets", property_sets)
        if repo_layout_ref is not None:
            pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if secondary_keypair_ref is not None:
            pulumi.set(__self__, "secondary_keypair_ref", secondary_keypair_ref)
        if trivial_layout is not None:
            warnings.warn("""You shouldn't be using this""", DeprecationWarning)
            pulumi.log.warn("""trivial_layout is deprecated: You shouldn't be using this""")
        if trivial_layout is not None:
            pulumi.set(__self__, "trivial_layout", trivial_layout)
        if xray_index is not None:
            pulumi.set(__self__, "xray_index", xray_index)

    @property
    @pulumi.getter(name="archiveBrowsingEnabled")
    def archive_browsing_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        security (e.g., cross-site scripting attacks).
        """
        return pulumi.get(self, "archive_browsing_enabled")

    @archive_browsing_enabled.setter
    def archive_browsing_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "archive_browsing_enabled", value)

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "blacked_out")

    @blacked_out.setter
    def blacked_out(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blacked_out", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "download_direct")

    @download_direct.setter
    def download_direct(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "download_direct", value)

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "excludes_pattern")

    @excludes_pattern.setter
    def excludes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "excludes_pattern", value)

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "includes_pattern")

    @includes_pattern.setter
    def includes_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "includes_pattern", value)

    @property
    @pulumi.getter(name="indexCompressionFormats")
    def index_compression_formats(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        - If you're creating this repo, then maybe you know?
        """
        return pulumi.get(self, "index_compression_formats")

    @index_compression_formats.setter
    def index_compression_formats(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "index_compression_formats", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        - the identity key of the repo
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "package_type")

    @package_type.setter
    def package_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "package_type", value)

    @property
    @pulumi.getter(name="primaryKeypairRef")
    def primary_keypair_ref(self) -> Optional[pulumi.Input[str]]:
        """
        - The RSA key to be used to sign packages
        """
        return pulumi.get(self, "primary_keypair_ref")

    @primary_keypair_ref.setter
    def primary_keypair_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_keypair_ref", value)

    @property
    @pulumi.getter(name="propertySets")
    def property_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "property_sets")

    @property_sets.setter
    def property_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "property_sets", value)

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "repo_layout_ref")

    @repo_layout_ref.setter
    def repo_layout_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repo_layout_ref", value)

    @property
    @pulumi.getter(name="secondaryKeypairRef")
    def secondary_keypair_ref(self) -> Optional[pulumi.Input[str]]:
        """
        - Not really clear what this does
        """
        return pulumi.get(self, "secondary_keypair_ref")

    @secondary_keypair_ref.setter
    def secondary_keypair_ref(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secondary_keypair_ref", value)

    @property
    @pulumi.getter(name="trivialLayout")
    def trivial_layout(self) -> Optional[pulumi.Input[bool]]:
        """
        - Apparently this is a deprecated repo layout
        """
        return pulumi.get(self, "trivial_layout")

    @trivial_layout.setter
    def trivial_layout(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "trivial_layout", value)

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "xray_index")

    @xray_index.setter
    def xray_index(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_index", value)


class DebianRepository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 primary_keypair_ref: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 secondary_keypair_ref: Optional[pulumi.Input[str]] = None,
                 trivial_layout: Optional[pulumi.Input[bool]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## # Artifactory Local Debian Repository Resource

        Creates a local Debian repository and allows for the creation of a GPG key

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        some_keypair_gpg1 = artifactory.Keypair("some-keypairGPG1",
            pair_name=f"some-keypair{random_id['randid']['id']}",
            pair_type="GPG",
            alias="foo-alias1",
            private_key=(lambda path: open(path).read())("samples/gpg.priv"),
            public_key=(lambda path: open(path).read())("samples/gpg.pub"))
        some_keypair_gpg2 = artifactory.Keypair("some-keypairGPG2",
            pair_name=f"some-keypair4{random_id['randid']['id']}",
            pair_type="GPG",
            alias="foo-alias2",
            private_key=(lambda path: open(path).read())("samples/gpg.priv"),
            public_key=(lambda path: open(path).read())("samples/gpg.pub"))
        my_debian_repo = artifactory.DebianRepository("my-debian-repo",
            key="my-debian-repo",
            primary_keypair_ref=some_keypair_gpg1.pair_name,
            secondary_keypair_ref=some_keypair_gpg2.pair_name,
            index_compression_formats=[
                "bz2",
                "lzma",
                "xz",
            ],
            trivial_layout=True,
            opts=pulumi.ResourceOptions(depends_on=[
                    some_keypair_gpg1,
                    some_keypair_gpg2,
                ]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] index_compression_formats: - If you're creating this repo, then maybe you know?
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[str] primary_keypair_ref: - The RSA key to be used to sign packages
        :param pulumi.Input[str] secondary_keypair_ref: - Not really clear what this does
        :param pulumi.Input[bool] trivial_layout: - Apparently this is a deprecated repo layout
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DebianRepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Artifactory Local Debian Repository Resource

        Creates a local Debian repository and allows for the creation of a GPG key

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        some_keypair_gpg1 = artifactory.Keypair("some-keypairGPG1",
            pair_name=f"some-keypair{random_id['randid']['id']}",
            pair_type="GPG",
            alias="foo-alias1",
            private_key=(lambda path: open(path).read())("samples/gpg.priv"),
            public_key=(lambda path: open(path).read())("samples/gpg.pub"))
        some_keypair_gpg2 = artifactory.Keypair("some-keypairGPG2",
            pair_name=f"some-keypair4{random_id['randid']['id']}",
            pair_type="GPG",
            alias="foo-alias2",
            private_key=(lambda path: open(path).read())("samples/gpg.priv"),
            public_key=(lambda path: open(path).read())("samples/gpg.pub"))
        my_debian_repo = artifactory.DebianRepository("my-debian-repo",
            key="my-debian-repo",
            primary_keypair_ref=some_keypair_gpg1.pair_name,
            secondary_keypair_ref=some_keypair_gpg2.pair_name,
            index_compression_formats=[
                "bz2",
                "lzma",
                "xz",
            ],
            trivial_layout=True,
            opts=pulumi.ResourceOptions(depends_on=[
                    some_keypair_gpg1,
                    some_keypair_gpg2,
                ]))
        ```

        :param str resource_name: The name of the resource.
        :param DebianRepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DebianRepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 primary_keypair_ref: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 secondary_keypair_ref: Optional[pulumi.Input[str]] = None,
                 trivial_layout: Optional[pulumi.Input[bool]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None,
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
            __props__ = DebianRepositoryArgs.__new__(DebianRepositoryArgs)

            __props__.__dict__["archive_browsing_enabled"] = archive_browsing_enabled
            __props__.__dict__["blacked_out"] = blacked_out
            __props__.__dict__["description"] = description
            __props__.__dict__["download_direct"] = download_direct
            __props__.__dict__["excludes_pattern"] = excludes_pattern
            __props__.__dict__["includes_pattern"] = includes_pattern
            __props__.__dict__["index_compression_formats"] = index_compression_formats
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["notes"] = notes
            __props__.__dict__["primary_keypair_ref"] = primary_keypair_ref
            __props__.__dict__["property_sets"] = property_sets
            __props__.__dict__["repo_layout_ref"] = repo_layout_ref
            __props__.__dict__["secondary_keypair_ref"] = secondary_keypair_ref
            if trivial_layout is not None and not opts.urn:
                warnings.warn("""You shouldn't be using this""", DeprecationWarning)
                pulumi.log.warn("""trivial_layout is deprecated: You shouldn't be using this""")
            __props__.__dict__["trivial_layout"] = trivial_layout
            __props__.__dict__["xray_index"] = xray_index
            __props__.__dict__["package_type"] = None
        super(DebianRepository, __self__).__init__(
            'artifactory:index/debianRepository:DebianRepository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
            blacked_out: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            download_direct: Optional[pulumi.Input[bool]] = None,
            excludes_pattern: Optional[pulumi.Input[str]] = None,
            includes_pattern: Optional[pulumi.Input[str]] = None,
            index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            key: Optional[pulumi.Input[str]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            package_type: Optional[pulumi.Input[str]] = None,
            primary_keypair_ref: Optional[pulumi.Input[str]] = None,
            property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            repo_layout_ref: Optional[pulumi.Input[str]] = None,
            secondary_keypair_ref: Optional[pulumi.Input[str]] = None,
            trivial_layout: Optional[pulumi.Input[bool]] = None,
            xray_index: Optional[pulumi.Input[bool]] = None) -> 'DebianRepository':
        """
        Get an existing DebianRepository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] index_compression_formats: - If you're creating this repo, then maybe you know?
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[str] primary_keypair_ref: - The RSA key to be used to sign packages
        :param pulumi.Input[str] secondary_keypair_ref: - Not really clear what this does
        :param pulumi.Input[bool] trivial_layout: - Apparently this is a deprecated repo layout
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DebianRepositoryState.__new__(_DebianRepositoryState)

        __props__.__dict__["archive_browsing_enabled"] = archive_browsing_enabled
        __props__.__dict__["blacked_out"] = blacked_out
        __props__.__dict__["description"] = description
        __props__.__dict__["download_direct"] = download_direct
        __props__.__dict__["excludes_pattern"] = excludes_pattern
        __props__.__dict__["includes_pattern"] = includes_pattern
        __props__.__dict__["index_compression_formats"] = index_compression_formats
        __props__.__dict__["key"] = key
        __props__.__dict__["notes"] = notes
        __props__.__dict__["package_type"] = package_type
        __props__.__dict__["primary_keypair_ref"] = primary_keypair_ref
        __props__.__dict__["property_sets"] = property_sets
        __props__.__dict__["repo_layout_ref"] = repo_layout_ref
        __props__.__dict__["secondary_keypair_ref"] = secondary_keypair_ref
        __props__.__dict__["trivial_layout"] = trivial_layout
        __props__.__dict__["xray_index"] = xray_index
        return DebianRepository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="archiveBrowsingEnabled")
    def archive_browsing_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
        therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
        security (e.g., cross-site scripting attacks).
        """
        return pulumi.get(self, "archive_browsing_enabled")

    @property
    @pulumi.getter(name="blackedOut")
    def blacked_out(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "blacked_out")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="downloadDirect")
    def download_direct(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "download_direct")

    @property
    @pulumi.getter(name="excludesPattern")
    def excludes_pattern(self) -> pulumi.Output[str]:
        return pulumi.get(self, "excludes_pattern")

    @property
    @pulumi.getter(name="includesPattern")
    def includes_pattern(self) -> pulumi.Output[str]:
        return pulumi.get(self, "includes_pattern")

    @property
    @pulumi.getter(name="indexCompressionFormats")
    def index_compression_formats(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        - If you're creating this repo, then maybe you know?
        """
        return pulumi.get(self, "index_compression_formats")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        - the identity key of the repo
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="primaryKeypairRef")
    def primary_keypair_ref(self) -> pulumi.Output[Optional[str]]:
        """
        - The RSA key to be used to sign packages
        """
        return pulumi.get(self, "primary_keypair_ref")

    @property
    @pulumi.getter(name="propertySets")
    def property_sets(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "property_sets")

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> pulumi.Output[str]:
        return pulumi.get(self, "repo_layout_ref")

    @property
    @pulumi.getter(name="secondaryKeypairRef")
    def secondary_keypair_ref(self) -> pulumi.Output[Optional[str]]:
        """
        - Not really clear what this does
        """
        return pulumi.get(self, "secondary_keypair_ref")

    @property
    @pulumi.getter(name="trivialLayout")
    def trivial_layout(self) -> pulumi.Output[Optional[bool]]:
        """
        - Apparently this is a deprecated repo layout
        """
        return pulumi.get(self, "trivial_layout")

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "xray_index")

