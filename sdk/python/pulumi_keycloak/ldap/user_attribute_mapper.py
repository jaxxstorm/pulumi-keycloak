# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UserAttributeMapper(pulumi.CustomResource):
    always_read_value_from_ldap: pulumi.Output[bool]
    """
    When true, the value fetched from LDAP will override the value stored in Keycloak.
    """
    is_mandatory_in_ldap: pulumi.Output[bool]
    """
    When true, this attribute must exist in LDAP.
    """
    ldap_attribute: pulumi.Output[str]
    """
    Name of the mapped attribute on LDAP object.
    """
    ldap_user_federation_id: pulumi.Output[str]
    """
    The ldap user federation provider to attach this mapper to.
    """
    name: pulumi.Output[str]
    """
    Display name of the mapper when displayed in the console.
    """
    read_only: pulumi.Output[bool]
    """
    When true, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak.
    """
    realm_id: pulumi.Output[str]
    """
    The realm in which the ldap user federation provider exists.
    """
    user_model_attribute: pulumi.Output[str]
    """
    Name of the UserModel property or attribute you want to map the LDAP attribute into.
    """
    def __init__(__self__, resource_name, opts=None, always_read_value_from_ldap=None, is_mandatory_in_ldap=None, ldap_attribute=None, ldap_user_federation_id=None, name=None, read_only=None, realm_id=None, user_model_attribute=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # ldap.UserAttributeMapper

        Allows for creating and managing user attribute mappers for Keycloak users
        federated via LDAP.

        The LDAP user attribute mapper can be used to map a single LDAP attribute
        to an attribute on the Keycloak user model.

        ### Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            enabled=True,
            realm="test")
        ldap_user_federation = keycloak.ldap.UserFederation("ldapUserFederation",
            bind_credential="admin",
            bind_dn="cn=admin,dc=example,dc=org",
            connection_url="ldap://openldap",
            rdn_ldap_attribute="cn",
            realm_id=realm.id,
            user_object_classes=[
                "simpleSecurityObject",
                "organizationalRole",
            ],
            username_ldap_attribute="cn",
            users_dn="dc=example,dc=org",
            uuid_ldap_attribute="entryDN")
        ldap_user_attribute_mapper = keycloak.ldap.UserAttributeMapper("ldapUserAttributeMapper",
            ldap_attribute="bar",
            ldap_user_federation_id=ldap_user_federation.id,
            realm_id=realm.id,
            user_model_attribute="foo")
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm that this LDAP mapper will exist in.
        - `ldap_user_federation_id` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
        - `name` - (Required) Display name of this mapper when displayed in the console.
        - `user_model_attribute` - (Required) Name of the user property or attribute you want to map the LDAP attribute into.
        - `ldap_attribute` - (Required) Name of the mapped attribute on the LDAP object.
        - `read_only` - (Optional) When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
        - `always_read_value_from_ldap` - (Optional) When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
        - `is_mandatory_in_ldap` - (Optional) When `true`, this attribute must exist in LDAP. Defaults to `false`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] always_read_value_from_ldap: When true, the value fetched from LDAP will override the value stored in Keycloak.
        :param pulumi.Input[bool] is_mandatory_in_ldap: When true, this attribute must exist in LDAP.
        :param pulumi.Input[str] ldap_attribute: Name of the mapped attribute on LDAP object.
        :param pulumi.Input[str] ldap_user_federation_id: The ldap user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of the mapper when displayed in the console.
        :param pulumi.Input[bool] read_only: When true, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak.
        :param pulumi.Input[str] realm_id: The realm in which the ldap user federation provider exists.
        :param pulumi.Input[str] user_model_attribute: Name of the UserModel property or attribute you want to map the LDAP attribute into.
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

            __props__['always_read_value_from_ldap'] = always_read_value_from_ldap
            __props__['is_mandatory_in_ldap'] = is_mandatory_in_ldap
            if ldap_attribute is None:
                raise TypeError("Missing required property 'ldap_attribute'")
            __props__['ldap_attribute'] = ldap_attribute
            if ldap_user_federation_id is None:
                raise TypeError("Missing required property 'ldap_user_federation_id'")
            __props__['ldap_user_federation_id'] = ldap_user_federation_id
            __props__['name'] = name
            __props__['read_only'] = read_only
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if user_model_attribute is None:
                raise TypeError("Missing required property 'user_model_attribute'")
            __props__['user_model_attribute'] = user_model_attribute
        super(UserAttributeMapper, __self__).__init__(
            'keycloak:ldap/userAttributeMapper:UserAttributeMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, always_read_value_from_ldap=None, is_mandatory_in_ldap=None, ldap_attribute=None, ldap_user_federation_id=None, name=None, read_only=None, realm_id=None, user_model_attribute=None):
        """
        Get an existing UserAttributeMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] always_read_value_from_ldap: When true, the value fetched from LDAP will override the value stored in Keycloak.
        :param pulumi.Input[bool] is_mandatory_in_ldap: When true, this attribute must exist in LDAP.
        :param pulumi.Input[str] ldap_attribute: Name of the mapped attribute on LDAP object.
        :param pulumi.Input[str] ldap_user_federation_id: The ldap user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of the mapper when displayed in the console.
        :param pulumi.Input[bool] read_only: When true, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak.
        :param pulumi.Input[str] realm_id: The realm in which the ldap user federation provider exists.
        :param pulumi.Input[str] user_model_attribute: Name of the UserModel property or attribute you want to map the LDAP attribute into.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["always_read_value_from_ldap"] = always_read_value_from_ldap
        __props__["is_mandatory_in_ldap"] = is_mandatory_in_ldap
        __props__["ldap_attribute"] = ldap_attribute
        __props__["ldap_user_federation_id"] = ldap_user_federation_id
        __props__["name"] = name
        __props__["read_only"] = read_only
        __props__["realm_id"] = realm_id
        __props__["user_model_attribute"] = user_model_attribute
        return UserAttributeMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

