# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DockerV2RepositoryArgs', 'DockerV2Repository']

@pulumi.input_type
class DockerV2RepositoryArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 block_pushing_schema1: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_unique_tags: Optional[pulumi.Input[int]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 tag_retention: Optional[pulumi.Input[int]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a DockerV2Repository resource.
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[bool] block_pushing_schema1: - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        :param pulumi.Input[int] max_unique_tags: - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
               Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
               This only applies to manifest v2
        :param pulumi.Input[int] tag_retention: - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
        """
        pulumi.set(__self__, "key", key)
        if archive_browsing_enabled is not None:
            pulumi.set(__self__, "archive_browsing_enabled", archive_browsing_enabled)
        if blacked_out is not None:
            pulumi.set(__self__, "blacked_out", blacked_out)
        if block_pushing_schema1 is not None:
            pulumi.set(__self__, "block_pushing_schema1", block_pushing_schema1)
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
        if max_unique_tags is not None:
            pulumi.set(__self__, "max_unique_tags", max_unique_tags)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if property_sets is not None:
            pulumi.set(__self__, "property_sets", property_sets)
        if repo_layout_ref is not None:
            pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if tag_retention is not None:
            pulumi.set(__self__, "tag_retention", tag_retention)
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
    @pulumi.getter(name="blockPushingSchema1")
    def block_pushing_schema1(self) -> Optional[pulumi.Input[bool]]:
        """
        - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        """
        return pulumi.get(self, "block_pushing_schema1")

    @block_pushing_schema1.setter
    def block_pushing_schema1(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_pushing_schema1", value)

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
        return pulumi.get(self, "index_compression_formats")

    @index_compression_formats.setter
    def index_compression_formats(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "index_compression_formats", value)

    @property
    @pulumi.getter(name="maxUniqueTags")
    def max_unique_tags(self) -> Optional[pulumi.Input[int]]:
        """
        - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
        Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
        This only applies to manifest v2
        """
        return pulumi.get(self, "max_unique_tags")

    @max_unique_tags.setter
    def max_unique_tags(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_unique_tags", value)

    @property
    @pulumi.getter
    def notes(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "notes")

    @notes.setter
    def notes(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notes", value)

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
    @pulumi.getter(name="tagRetention")
    def tag_retention(self) -> Optional[pulumi.Input[int]]:
        """
        - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
        """
        return pulumi.get(self, "tag_retention")

    @tag_retention.setter
    def tag_retention(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tag_retention", value)

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "xray_index")

    @xray_index.setter
    def xray_index(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_index", value)


@pulumi.input_type
class _DockerV2RepositoryState:
    def __init__(__self__, *,
                 api_version: Optional[pulumi.Input[str]] = None,
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 block_pushing_schema1: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 max_unique_tags: Optional[pulumi.Input[int]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 package_type: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 tag_retention: Optional[pulumi.Input[int]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering DockerV2Repository resources.
        :param pulumi.Input[str] api_version: The Docker API version to use. This cannot be set
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[bool] block_pushing_schema1: - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[int] max_unique_tags: - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
               Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
               This only applies to manifest v2
        :param pulumi.Input[int] tag_retention: - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
        """
        if api_version is not None:
            pulumi.set(__self__, "api_version", api_version)
        if archive_browsing_enabled is not None:
            pulumi.set(__self__, "archive_browsing_enabled", archive_browsing_enabled)
        if blacked_out is not None:
            pulumi.set(__self__, "blacked_out", blacked_out)
        if block_pushing_schema1 is not None:
            pulumi.set(__self__, "block_pushing_schema1", block_pushing_schema1)
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
        if max_unique_tags is not None:
            pulumi.set(__self__, "max_unique_tags", max_unique_tags)
        if notes is not None:
            pulumi.set(__self__, "notes", notes)
        if package_type is not None:
            pulumi.set(__self__, "package_type", package_type)
        if property_sets is not None:
            pulumi.set(__self__, "property_sets", property_sets)
        if repo_layout_ref is not None:
            pulumi.set(__self__, "repo_layout_ref", repo_layout_ref)
        if tag_retention is not None:
            pulumi.set(__self__, "tag_retention", tag_retention)
        if xray_index is not None:
            pulumi.set(__self__, "xray_index", xray_index)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        The Docker API version to use. This cannot be set
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

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
    @pulumi.getter(name="blockPushingSchema1")
    def block_pushing_schema1(self) -> Optional[pulumi.Input[bool]]:
        """
        - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        """
        return pulumi.get(self, "block_pushing_schema1")

    @block_pushing_schema1.setter
    def block_pushing_schema1(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "block_pushing_schema1", value)

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
    @pulumi.getter(name="maxUniqueTags")
    def max_unique_tags(self) -> Optional[pulumi.Input[int]]:
        """
        - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
        Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
        This only applies to manifest v2
        """
        return pulumi.get(self, "max_unique_tags")

    @max_unique_tags.setter
    def max_unique_tags(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_unique_tags", value)

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
    @pulumi.getter(name="tagRetention")
    def tag_retention(self) -> Optional[pulumi.Input[int]]:
        """
        - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
        """
        return pulumi.get(self, "tag_retention")

    @tag_retention.setter
    def tag_retention(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tag_retention", value)

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "xray_index")

    @xray_index.setter
    def xray_index(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "xray_index", value)


class DockerV2Repository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 block_pushing_schema1: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 max_unique_tags: Optional[pulumi.Input[int]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 tag_retention: Optional[pulumi.Input[int]] = None,
                 xray_index: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## # Artifactory Local Docker V2 Repository Resource

        Creates a local Docker v2 repository

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        foo = artifactory.DockerV2Repository("foo",
            key="foo",
            max_unique_tags=5,
            tag_retention=3)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[bool] block_pushing_schema1: - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[int] max_unique_tags: - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
               Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
               This only applies to manifest v2
        :param pulumi.Input[int] tag_retention: - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DockerV2RepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Artifactory Local Docker V2 Repository Resource

        Creates a local Docker v2 repository

        ## Example Usage

        ```python
        import pulumi
        import pulumi_artifactory as artifactory

        foo = artifactory.DockerV2Repository("foo",
            key="foo",
            max_unique_tags=5,
            tag_retention=3)
        ```

        :param str resource_name: The name of the resource.
        :param DockerV2RepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DockerV2RepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
                 blacked_out: Optional[pulumi.Input[bool]] = None,
                 block_pushing_schema1: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 download_direct: Optional[pulumi.Input[bool]] = None,
                 excludes_pattern: Optional[pulumi.Input[str]] = None,
                 includes_pattern: Optional[pulumi.Input[str]] = None,
                 index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 max_unique_tags: Optional[pulumi.Input[int]] = None,
                 notes: Optional[pulumi.Input[str]] = None,
                 property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 repo_layout_ref: Optional[pulumi.Input[str]] = None,
                 tag_retention: Optional[pulumi.Input[int]] = None,
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
            __props__ = DockerV2RepositoryArgs.__new__(DockerV2RepositoryArgs)

            __props__.__dict__["archive_browsing_enabled"] = archive_browsing_enabled
            __props__.__dict__["blacked_out"] = blacked_out
            __props__.__dict__["block_pushing_schema1"] = block_pushing_schema1
            __props__.__dict__["description"] = description
            __props__.__dict__["download_direct"] = download_direct
            __props__.__dict__["excludes_pattern"] = excludes_pattern
            __props__.__dict__["includes_pattern"] = includes_pattern
            __props__.__dict__["index_compression_formats"] = index_compression_formats
            if key is None and not opts.urn:
                raise TypeError("Missing required property 'key'")
            __props__.__dict__["key"] = key
            __props__.__dict__["max_unique_tags"] = max_unique_tags
            __props__.__dict__["notes"] = notes
            __props__.__dict__["property_sets"] = property_sets
            __props__.__dict__["repo_layout_ref"] = repo_layout_ref
            __props__.__dict__["tag_retention"] = tag_retention
            __props__.__dict__["xray_index"] = xray_index
            __props__.__dict__["api_version"] = None
            __props__.__dict__["package_type"] = None
        super(DockerV2Repository, __self__).__init__(
            'artifactory:index/dockerV2Repository:DockerV2Repository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_version: Optional[pulumi.Input[str]] = None,
            archive_browsing_enabled: Optional[pulumi.Input[bool]] = None,
            blacked_out: Optional[pulumi.Input[bool]] = None,
            block_pushing_schema1: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            download_direct: Optional[pulumi.Input[bool]] = None,
            excludes_pattern: Optional[pulumi.Input[str]] = None,
            includes_pattern: Optional[pulumi.Input[str]] = None,
            index_compression_formats: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            key: Optional[pulumi.Input[str]] = None,
            max_unique_tags: Optional[pulumi.Input[int]] = None,
            notes: Optional[pulumi.Input[str]] = None,
            package_type: Optional[pulumi.Input[str]] = None,
            property_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            repo_layout_ref: Optional[pulumi.Input[str]] = None,
            tag_retention: Optional[pulumi.Input[int]] = None,
            xray_index: Optional[pulumi.Input[bool]] = None) -> 'DockerV2Repository':
        """
        Get an existing DockerV2Repository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_version: The Docker API version to use. This cannot be set
        :param pulumi.Input[bool] archive_browsing_enabled: When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
               therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
               security (e.g., cross-site scripting attacks).
        :param pulumi.Input[bool] block_pushing_schema1: - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        :param pulumi.Input[str] key: - the identity key of the repo
        :param pulumi.Input[int] max_unique_tags: - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
               Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
               This only applies to manifest v2
        :param pulumi.Input[int] tag_retention: - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DockerV2RepositoryState.__new__(_DockerV2RepositoryState)

        __props__.__dict__["api_version"] = api_version
        __props__.__dict__["archive_browsing_enabled"] = archive_browsing_enabled
        __props__.__dict__["blacked_out"] = blacked_out
        __props__.__dict__["block_pushing_schema1"] = block_pushing_schema1
        __props__.__dict__["description"] = description
        __props__.__dict__["download_direct"] = download_direct
        __props__.__dict__["excludes_pattern"] = excludes_pattern
        __props__.__dict__["includes_pattern"] = includes_pattern
        __props__.__dict__["index_compression_formats"] = index_compression_formats
        __props__.__dict__["key"] = key
        __props__.__dict__["max_unique_tags"] = max_unique_tags
        __props__.__dict__["notes"] = notes
        __props__.__dict__["package_type"] = package_type
        __props__.__dict__["property_sets"] = property_sets
        __props__.__dict__["repo_layout_ref"] = repo_layout_ref
        __props__.__dict__["tag_retention"] = tag_retention
        __props__.__dict__["xray_index"] = xray_index
        return DockerV2Repository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> pulumi.Output[str]:
        """
        The Docker API version to use. This cannot be set
        """
        return pulumi.get(self, "api_version")

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
    @pulumi.getter(name="blockPushingSchema1")
    def block_pushing_schema1(self) -> pulumi.Output[bool]:
        """
        - When set, Artifactory will block the pushing of Docker images with manifest v2 schema 1 to this repository.
        """
        return pulumi.get(self, "block_pushing_schema1")

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
        return pulumi.get(self, "index_compression_formats")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        - the identity key of the repo
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter(name="maxUniqueTags")
    def max_unique_tags(self) -> pulumi.Output[Optional[int]]:
        """
        - The maximum number of unique tags of a single Docker image to store in this repository.\n" +
        Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there is no limit.
        This only applies to manifest v2
        """
        return pulumi.get(self, "max_unique_tags")

    @property
    @pulumi.getter
    def notes(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter(name="packageType")
    def package_type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "package_type")

    @property
    @pulumi.getter(name="propertySets")
    def property_sets(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "property_sets")

    @property
    @pulumi.getter(name="repoLayoutRef")
    def repo_layout_ref(self) -> pulumi.Output[str]:
        return pulumi.get(self, "repo_layout_ref")

    @property
    @pulumi.getter(name="tagRetention")
    def tag_retention(self) -> pulumi.Output[Optional[int]]:
        """
        - If greater than 1, overwritten tags will be saved by their digest, up to the set up number. This only applies to manifest V2
        """
        return pulumi.get(self, "tag_retention")

    @property
    @pulumi.getter(name="xrayIndex")
    def xray_index(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "xray_index")

