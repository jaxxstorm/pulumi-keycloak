// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public static class GetClient
    {
        /// <summary>
        /// ## # keycloak.openid.Client data source
        /// 
        /// This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.
        /// 
        /// ### Example Usage
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Keycloak = Pulumi.Keycloak;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var realmManagement = Output.Create(Keycloak.OpenId.GetClient.InvokeAsync(new Keycloak.OpenId.GetClientArgs
        ///         {
        ///             RealmId = "my-realm",
        ///             ClientId = "realm-management",
        ///         }));
        ///         var admin = realmManagement.Apply(realmManagement =&gt; Output.Create(Keycloak.GetRole.InvokeAsync(new Keycloak.GetRoleArgs
        ///         {
        ///             RealmId = "my-realm",
        ///             ClientId = realmManagement.Id,
        ///             Name = "realm-admin",
        ///         })));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ### Argument Reference
        /// 
        /// The following arguments are supported:
        /// 
        /// - `realm_id` - (Required) The realm id.
        /// - `client_id` - (Required) The client id.
        /// 
        /// ### Attributes Reference
        /// 
        /// See the docs for the `keycloak.openid.Client` resource for details on the exported attributes.
        /// </summary>
        public static Task<GetClientResult> InvokeAsync(GetClientArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClientResult>("keycloak:openid/getClient:getClient", args ?? new GetClientArgs(), options.WithVersion());
    }


    public sealed class GetClientArgs : Pulumi.InvokeArgs
    {
        [Input("clientId", required: true)]
        public string ClientId { get; set; } = null!;

        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetClientArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClientResult
    {
        public readonly string AccessType;
        public readonly Outputs.GetClientAuthenticationFlowBindingOverridesResult AuthenticationFlowBindingOverrides;
        public readonly Outputs.GetClientAuthorizationResult Authorization;
        public readonly string ClientId;
        public readonly string ClientSecret;
        public readonly bool ConsentRequired;
        public readonly string Description;
        public readonly bool DirectAccessGrantsEnabled;
        public readonly bool Enabled;
        public readonly bool FullScopeAllowed;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool ImplicitFlowEnabled;
        public readonly string LoginTheme;
        public readonly string Name;
        public readonly string RealmId;
        public readonly string ResourceServerId;
        public readonly string RootUrl;
        public readonly string ServiceAccountUserId;
        public readonly bool ServiceAccountsEnabled;
        public readonly bool StandardFlowEnabled;
        public readonly ImmutableArray<string> ValidRedirectUris;
        public readonly ImmutableArray<string> WebOrigins;

        [OutputConstructor]
        private GetClientResult(
            string accessType,

            Outputs.GetClientAuthenticationFlowBindingOverridesResult authenticationFlowBindingOverrides,

            Outputs.GetClientAuthorizationResult authorization,

            string clientId,

            string clientSecret,

            bool consentRequired,

            string description,

            bool directAccessGrantsEnabled,

            bool enabled,

            bool fullScopeAllowed,

            string id,

            bool implicitFlowEnabled,

            string loginTheme,

            string name,

            string realmId,

            string resourceServerId,

            string rootUrl,

            string serviceAccountUserId,

            bool serviceAccountsEnabled,

            bool standardFlowEnabled,

            ImmutableArray<string> validRedirectUris,

            ImmutableArray<string> webOrigins)
        {
            AccessType = accessType;
            AuthenticationFlowBindingOverrides = authenticationFlowBindingOverrides;
            Authorization = authorization;
            ClientId = clientId;
            ClientSecret = clientSecret;
            ConsentRequired = consentRequired;
            Description = description;
            DirectAccessGrantsEnabled = directAccessGrantsEnabled;
            Enabled = enabled;
            FullScopeAllowed = fullScopeAllowed;
            Id = id;
            ImplicitFlowEnabled = implicitFlowEnabled;
            LoginTheme = loginTheme;
            Name = name;
            RealmId = realmId;
            ResourceServerId = resourceServerId;
            RootUrl = rootUrl;
            ServiceAccountUserId = serviceAccountUserId;
            ServiceAccountsEnabled = serviceAccountsEnabled;
            StandardFlowEnabled = standardFlowEnabled;
            ValidRedirectUris = validRedirectUris;
            WebOrigins = webOrigins;
        }
    }
}
