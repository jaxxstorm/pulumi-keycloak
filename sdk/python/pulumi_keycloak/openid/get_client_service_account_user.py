# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetClientServiceAccountUserResult:
    """
    A collection of values returned by getClientServiceAccountUser.
    """
    def __init__(__self__, attributes=None, client_id=None, email=None, enabled=None, federated_identities=None, first_name=None, id=None, last_name=None, realm_id=None, username=None):
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        __self__.attributes = attributes
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        __self__.client_id = client_id
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        __self__.email = email
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        __self__.enabled = enabled
        if federated_identities and not isinstance(federated_identities, list):
            raise TypeError("Expected argument 'federated_identities' to be a list")
        __self__.federated_identities = federated_identities
        if first_name and not isinstance(first_name, str):
            raise TypeError("Expected argument 'first_name' to be a str")
        __self__.first_name = first_name
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if last_name and not isinstance(last_name, str):
            raise TypeError("Expected argument 'last_name' to be a str")
        __self__.last_name = last_name
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        __self__.realm_id = realm_id
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        __self__.username = username
class AwaitableGetClientServiceAccountUserResult(GetClientServiceAccountUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientServiceAccountUserResult(
            attributes=self.attributes,
            client_id=self.client_id,
            email=self.email,
            enabled=self.enabled,
            federated_identities=self.federated_identities,
            first_name=self.first_name,
            id=self.id,
            last_name=self.last_name,
            realm_id=self.realm_id,
            username=self.username)

def get_client_service_account_user(client_id=None,realm_id=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['clientId'] = client_id
    __args__['realmId'] = realm_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('keycloak:openid/getClientServiceAccountUser:getClientServiceAccountUser', __args__, opts=opts).value

    return AwaitableGetClientServiceAccountUserResult(
        attributes=__ret__.get('attributes'),
        client_id=__ret__.get('clientId'),
        email=__ret__.get('email'),
        enabled=__ret__.get('enabled'),
        federated_identities=__ret__.get('federatedIdentities'),
        first_name=__ret__.get('firstName'),
        id=__ret__.get('id'),
        last_name=__ret__.get('lastName'),
        realm_id=__ret__.get('realmId'),
        username=__ret__.get('username'))
