// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public static partial class Invokes
    {
        public static Task<GetClientServiceAccountUserResult> GetClientServiceAccountUser(GetClientServiceAccountUserArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClientServiceAccountUserResult>("keycloak:openid/getClientServiceAccountUser:getClientServiceAccountUser", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetClientServiceAccountUserArgs : Pulumi.InvokeArgs
    {
        [Input("clientId", required: true)]
        public string ClientId { get; set; } = null!;

        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetClientServiceAccountUserArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetClientServiceAccountUserResult
    {
        public readonly ImmutableDictionary<string, object> Attributes;
        public readonly string ClientId;
        public readonly string Email;
        public readonly bool Enabled;
        public readonly ImmutableArray<Outputs.GetClientServiceAccountUserFederatedIdentitiesResult> FederatedIdentities;
        public readonly string FirstName;
        public readonly string LastName;
        public readonly string RealmId;
        public readonly string Username;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetClientServiceAccountUserResult(
            ImmutableDictionary<string, object> attributes,
            string clientId,
            string email,
            bool enabled,
            ImmutableArray<Outputs.GetClientServiceAccountUserFederatedIdentitiesResult> federatedIdentities,
            string firstName,
            string lastName,
            string realmId,
            string username,
            string id)
        {
            Attributes = attributes;
            ClientId = clientId;
            Email = email;
            Enabled = enabled;
            FederatedIdentities = federatedIdentities;
            FirstName = firstName;
            LastName = lastName;
            RealmId = realmId;
            Username = username;
            Id = id;
        }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetClientServiceAccountUserFederatedIdentitiesResult
    {
        public readonly string IdentityProvider;
        public readonly string UserId;
        public readonly string UserName;

        [OutputConstructor]
        private GetClientServiceAccountUserFederatedIdentitiesResult(
            string identityProvider,
            string userId,
            string userName)
        {
            IdentityProvider = identityProvider;
            UserId = userId;
            UserName = userName;
        }
    }
    }
}
