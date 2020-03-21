# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Group(pulumi.CustomResource):
    attributes: pulumi.Output[dict]
    name: pulumi.Output[str]
    parent_id: pulumi.Output[str]
    path: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, attributes=None, name=None, parent_id=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # .Group

        Allows for creating and managing Groups within Keycloak.

        Groups provide a logical wrapping for users within Keycloak. Users within a
        group can share attributes and roles, and group membership can be mapped
        to a claim.

        Attributes can also be defined on Groups.

        Groups can also be federated from external data sources, such as LDAP or Active Directory.
        This resource **should not** be used to manage groups that were created this way.

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this group exists in.
        - `parent_id` - (Optional) The ID of this group's parent. If omitted, this group will be defined at the root level.
        - `name` - (Required) The name of the group.
        - `attributes` - (Optional) A dict of key/value pairs to set as custom attributes for the group.

        ### Attributes Reference

        In addition to the arguments listed above, the following computed attributes are exported:

        - `path` - The complete path of the group. For example, the child group's path in the example configuration would be `/parent-group/child-group`.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_group.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['attributes'] = attributes
            __props__['name'] = name
            __props__['parent_id'] = parent_id
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            __props__['path'] = None
        super(Group, __self__).__init__(
            'keycloak:index/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, attributes=None, name=None, parent_id=None, path=None, realm_id=None):
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attributes"] = attributes
        __props__["name"] = name
        __props__["parent_id"] = parent_id
        __props__["path"] = path
        __props__["realm_id"] = realm_id
        return Group(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

