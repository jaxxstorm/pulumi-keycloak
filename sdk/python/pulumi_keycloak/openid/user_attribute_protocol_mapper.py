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
    add_to_access_token: pulumi.Output[bool]
    """
    Indicates if the attribute should be a claim in the access token.
    """
    add_to_id_token: pulumi.Output[bool]
    """
    Indicates if the attribute should be a claim in the id token.
    """
    add_to_userinfo: pulumi.Output[bool]
    """
    Indicates if the attribute should appear in the userinfo response body.
    """
    aggregate_attributes: pulumi.Output[bool]
    """
    Indicates if attribute values should be aggregated within the group attributes
    """
    claim_name: pulumi.Output[str]
    claim_value_type: pulumi.Output[str]
    """
    Claim type used when serializing tokens.
    """
    client_id: pulumi.Output[str]
    """
    The mapper's associated client. Cannot be used at the same time as client_scope_id.
    """
    client_scope_id: pulumi.Output[str]
    """
    The mapper's associated client scope. Cannot be used at the same time as client_id.
    """
    multivalued: pulumi.Output[bool]
    """
    Indicates whether this attribute is a single value or an array of values.
    """
    name: pulumi.Output[str]
    """
    A human-friendly name that will appear in the Keycloak console.
    """
    realm_id: pulumi.Output[str]
    """
    The realm id where the associated client or client scope exists.
    """
    user_attribute: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, add_to_access_token=None, add_to_id_token=None, add_to_userinfo=None, aggregate_attributes=None, claim_name=None, claim_value_type=None, client_id=None, client_scope_id=None, multivalued=None, name=None, realm_id=None, user_attribute=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # openid.UserAttributeProtocolMapper

        Allows for creating and managing user attribute protocol mappers within
        Keycloak.

        User attribute protocol mappers allow you to map custom attributes defined
        for a user within Keycloak to a claim in a token. Protocol mappers can be
        defined for a single client, or they can be defined for a client scope which
        can be shared between multiple different clients.

        ### Example Usage (Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            enabled=True,
            realm="my-realm")
        openid_client = keycloak.openid.Client("openidClient",
            access_type="CONFIDENTIAL",
            client_id="test-client",
            enabled=True,
            realm_id=realm.id,
            valid_redirect_uris=["http://localhost:8080/openid-callback"])
        user_attribute_mapper = keycloak.openid.UserAttributeProtocolMapper("userAttributeMapper",
            claim_name="bar",
            client_id=openid_client.id,
            realm_id=realm.id,
            user_attribute="foo")
        ```

        ### Example Usage (Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            enabled=True,
            realm="my-realm")
        client_scope = keycloak.openid.ClientScope("clientScope", realm_id=realm.id)
        user_attribute_mapper = keycloak.openid.UserAttributeProtocolMapper("userAttributeMapper",
            claim_name="bar",
            client_scope_id=client_scope.id,
            realm_id=realm.id,
            user_attribute="foo")
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required if `client_scope_id` is not specified) The client this protocol mapper is attached to.
        - `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `user_attribute` - (Required) The custom user attribute to map a claim for.
        - `claim_name` - (Required) The name of the claim to insert into a token.
        - `claim_value_type` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
        - `multivalued` - (Optional) Indicates whether this attribute is a single value or an array of values. Defaults to `false`.
        - `add_to_id_token` - (Optional) Indicates if the attribute should be added as a claim to the id token. Defaults to `true`.
        - `add_to_access_token` - (Optional) Indicates if the attribute should be added as a claim to the access token. Defaults to `true`.
        - `add_to_userinfo` - (Optional) Indicates if the attribute should be added as a claim to the UserInfo response body. Defaults to `true`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the attribute should be a claim in the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the attribute should be a claim in the id token.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the attribute should appear in the userinfo response body.
        :param pulumi.Input[bool] aggregate_attributes: Indicates if attribute values should be aggregated within the group attributes
        :param pulumi.Input[str] claim_value_type: Claim type used when serializing tokens.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[bool] multivalued: Indicates whether this attribute is a single value or an array of values.
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
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

            __props__['add_to_access_token'] = add_to_access_token
            __props__['add_to_id_token'] = add_to_id_token
            __props__['add_to_userinfo'] = add_to_userinfo
            __props__['aggregate_attributes'] = aggregate_attributes
            if claim_name is None:
                raise TypeError("Missing required property 'claim_name'")
            __props__['claim_name'] = claim_name
            __props__['claim_value_type'] = claim_value_type
            __props__['client_id'] = client_id
            __props__['client_scope_id'] = client_scope_id
            __props__['multivalued'] = multivalued
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if user_attribute is None:
                raise TypeError("Missing required property 'user_attribute'")
            __props__['user_attribute'] = user_attribute
        super(UserAttributeProtocolMapper, __self__).__init__(
            'keycloak:openid/userAttributeProtocolMapper:UserAttributeProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, add_to_access_token=None, add_to_id_token=None, add_to_userinfo=None, aggregate_attributes=None, claim_name=None, claim_value_type=None, client_id=None, client_scope_id=None, multivalued=None, name=None, realm_id=None, user_attribute=None):
        """
        Get an existing UserAttributeProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the attribute should be a claim in the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the attribute should be a claim in the id token.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the attribute should appear in the userinfo response body.
        :param pulumi.Input[bool] aggregate_attributes: Indicates if attribute values should be aggregated within the group attributes
        :param pulumi.Input[str] claim_value_type: Claim type used when serializing tokens.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[bool] multivalued: Indicates whether this attribute is a single value or an array of values.
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["add_to_access_token"] = add_to_access_token
        __props__["add_to_id_token"] = add_to_id_token
        __props__["add_to_userinfo"] = add_to_userinfo
        __props__["aggregate_attributes"] = aggregate_attributes
        __props__["claim_name"] = claim_name
        __props__["claim_value_type"] = claim_value_type
        __props__["client_id"] = client_id
        __props__["client_scope_id"] = client_scope_id
        __props__["multivalued"] = multivalued
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        __props__["user_attribute"] = user_attribute
        return UserAttributeProtocolMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

