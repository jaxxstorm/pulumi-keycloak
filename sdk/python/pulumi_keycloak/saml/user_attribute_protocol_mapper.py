# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UserAttributeProtocolMapper(pulumi.CustomResource):
    client_id: pulumi.Output[str]
    client_scope_id: pulumi.Output[str]
    friendly_name: pulumi.Output[str]
    name: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    saml_attribute_name: pulumi.Output[str]
    saml_attribute_name_format: pulumi.Output[str]
    user_attribute: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, client_id=None, client_scope_id=None, friendly_name=None, name=None, realm_id=None, saml_attribute_name=None, saml_attribute_name_format=None, user_attribute=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # saml.UserAttributeProtocolMapper

        Allows for creating and managing user attribute protocol mappers for
        SAML clients within Keycloak.

        SAML user attribute protocol mappers allow you to map custom attributes defined
        for a user within Keycloak to an attribute in a SAML assertion. Protocol mappers
        can be defined for a single client, or they can be defined for a client scope which
        can be shared between multiple different clients.

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required if `client_scope_id` is not specified) The SAML client this protocol mapper is attached to.
        - `client_scope_id` - (Required if `client_id` is not specified) The SAML client scope this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `user_attribute` - (Required) The custom user attribute to map.
        - `friendly_name` - (Optional) An optional human-friendly name for this attribute.
        - `saml_attribute_name` - (Required) The name of the SAML attribute.
        - `saml_attribute_name_format` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_saml_user_attribute_protocol_mapper.html.markdown.

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
            __props__['client_scope_id'] = client_scope_id
            __props__['friendly_name'] = friendly_name
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if saml_attribute_name is None:
                raise TypeError("Missing required property 'saml_attribute_name'")
            __props__['saml_attribute_name'] = saml_attribute_name
            if saml_attribute_name_format is None:
                raise TypeError("Missing required property 'saml_attribute_name_format'")
            __props__['saml_attribute_name_format'] = saml_attribute_name_format
            if user_attribute is None:
                raise TypeError("Missing required property 'user_attribute'")
            __props__['user_attribute'] = user_attribute
        super(UserAttributeProtocolMapper, __self__).__init__(
            'keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, client_id=None, client_scope_id=None, friendly_name=None, name=None, realm_id=None, saml_attribute_name=None, saml_attribute_name_format=None, user_attribute=None):
        """
        Get an existing UserAttributeProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["client_scope_id"] = client_scope_id
        __props__["friendly_name"] = friendly_name
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        __props__["saml_attribute_name"] = saml_attribute_name
        __props__["saml_attribute_name_format"] = saml_attribute_name_format
        __props__["user_attribute"] = user_attribute
        return UserAttributeProtocolMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

