// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Saml
{
    public static class GetClientInstallationProvider
    {
        public static Task<GetClientInstallationProviderResult> InvokeAsync(GetClientInstallationProviderArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetClientInstallationProviderResult>("keycloak:saml/getClientInstallationProvider:getClientInstallationProvider", args ?? new GetClientInstallationProviderArgs(), options.WithVersion());
    }


    public sealed class GetClientInstallationProviderArgs : Pulumi.InvokeArgs
    {
        [Input("clientId", required: true)]
        public string ClientId { get; set; } = null!;

        [Input("providerId", required: true)]
        public string ProviderId { get; set; } = null!;

        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        public GetClientInstallationProviderArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetClientInstallationProviderResult
    {
        public readonly string ClientId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ProviderId;
        public readonly string RealmId;
        public readonly string Value;

        [OutputConstructor]
        private GetClientInstallationProviderResult(
            string clientId,

            string id,

            string providerId,

            string realmId,

            string value)
        {
            ClientId = clientId;
            Id = id;
            ProviderId = providerId;
            RealmId = realmId;
            Value = value;
        }
    }
}
