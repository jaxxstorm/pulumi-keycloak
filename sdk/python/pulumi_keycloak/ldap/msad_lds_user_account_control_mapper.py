# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class MsadLdsUserAccountControlMapper(pulumi.CustomResource):
    ldap_user_federation_id: pulumi.Output[str]
    """
    The ldap user federation provider to attach this mapper to.
    """
    name: pulumi.Output[str]
    """
    Display name of the mapper when displayed in the console.
    """
    realm_id: pulumi.Output[str]
    """
    The realm in which the ldap user federation provider exists.
    """
    def __init__(__self__, resource_name, opts=None, ldap_user_federation_id=None, name=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a MsadLdsUserAccountControlMapper resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ldap user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of the mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm in which the ldap user federation provider exists.
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

            if ldap_user_federation_id is None:
                raise TypeError("Missing required property 'ldap_user_federation_id'")
            __props__['ldap_user_federation_id'] = ldap_user_federation_id
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(MsadLdsUserAccountControlMapper, __self__).__init__(
            'keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, ldap_user_federation_id=None, name=None, realm_id=None):
        """
        Get an existing MsadLdsUserAccountControlMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ldap user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of the mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm in which the ldap user federation provider exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["ldap_user_federation_id"] = ldap_user_federation_id
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        return MsadLdsUserAccountControlMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

