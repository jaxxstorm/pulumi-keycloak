# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ClientScope(pulumi.CustomResource):
    consent_screen_text: pulumi.Output[str]
    description: pulumi.Output[str]
    name: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, consent_screen_text=None, description=None, name=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # openid.ClientScope

        Allows for creating and managing Keycloak client scopes that can be attached to
        clients that use the OpenID Connect protocol.

        Client Scopes can be used to share common protocol and role mappings between multiple
        clients within a realm. They can also be used by clients to conditionally request
        claims or roles for a user based on the OAuth 2.0 `scope` parameter.

        ### Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            enabled=True,
            realm="my-realm")
        openid_client_scope = keycloak.openid.ClientScope("openidClientScope",
            description="When requested, this scope will map a user's group memberships to a claim",
            realm_id=realm.id)
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this client scope belongs to.
        - `name` - (Required) The display name of this client scope in the GUI.
        - `description` - (Optional) The description of this client scope in the GUI.
        - `consent_screen_text` - (Optional) When set, a consent screen will be displayed to users
        authenticating to clients with this scope attached. The consent screen will display the string
        value of this attribute.

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

            __props__['consent_screen_text'] = consent_screen_text
            __props__['description'] = description
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(ClientScope, __self__).__init__(
            'keycloak:openid/clientScope:ClientScope',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, consent_screen_text=None, description=None, name=None, realm_id=None):
        """
        Get an existing ClientScope resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["consent_screen_text"] = consent_screen_text
        __props__["description"] = description
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        return ClientScope(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

