# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class Role(pulumi.CustomResource):
    client_id: pulumi.Output[str]
    composite_roles: pulumi.Output[list]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, client_id=None, composite_roles=None, description=None, name=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # .Role

        Allows for creating and managing roles within Keycloak.

        Roles allow you define privileges within Keycloak and map them to users
        and groups.

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this role exists within.
        - `client_id` - (Optional) When specified, this role will be created as
          a client role attached to the client with the provided ID
        - `name` - (Required) The name of the role
        - `description` - (Optional) The description of the role
        - `composite_roles` - (Optional) When specified, this role will be a
          composite role, composed of all roles that have an ID present within
          this list.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_role.html.markdown.

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

            __props__['client_id'] = client_id
            __props__['composite_roles'] = composite_roles
            __props__['description'] = description
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(Role, __self__).__init__(
            'keycloak:index/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, client_id=None, composite_roles=None, description=None, name=None, realm_id=None):
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["composite_roles"] = composite_roles
        __props__["description"] = description
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        return Role(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

