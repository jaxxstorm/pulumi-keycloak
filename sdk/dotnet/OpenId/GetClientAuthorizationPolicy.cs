// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public static class GetClientAuthorizationPolicy
    {
        public static Task<GetClientAuthorizationPolicyResult> InvokeAsync(GetClientAuthorizationPolicyArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClientAuthorizationPolicyResult>("keycloak:openid/getClientAuthorizationPolicy:getClientAuthorizationPolicy", args ?? new GetClientAuthorizationPolicyArgs(), options.WithVersion());
    }


    public sealed class GetClientAuthorizationPolicyArgs : Pulumi.InvokeArgs
    {
        [Input("logic")]
        public string? Logic { get; set; }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        [Input("resourceServerId", required: true)]
        public string ResourceServerId { get; set; } = null!;

        public GetClientAuthorizationPolicyArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClientAuthorizationPolicyResult
    {
        public readonly string DecisionStrategy;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? Logic;
        public readonly string Name;
        public readonly string Owner;
        public readonly ImmutableArray<string> Policies;
        public readonly string RealmId;
        public readonly string ResourceServerId;
        public readonly ImmutableArray<string> Resources;
        public readonly ImmutableArray<string> Scopes;
        public readonly string Type;

        [OutputConstructor]
        private GetClientAuthorizationPolicyResult(
            string decisionStrategy,

            string id,

            string? logic,

            string name,

            string owner,

            ImmutableArray<string> policies,

            string realmId,

            string resourceServerId,

            ImmutableArray<string> resources,

            ImmutableArray<string> scopes,

            string type)
        {
            DecisionStrategy = decisionStrategy;
            Id = id;
            Logic = logic;
            Name = name;
            Owner = owner;
            Policies = policies;
            RealmId = realmId;
            ResourceServerId = resourceServerId;
            Resources = resources;
            Scopes = scopes;
            Type = type;
        }
    }
}
