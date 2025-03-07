// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    /// <summary>
    /// ## # keycloak.openid.HardcodedClaimProtocolMapper
    /// 
    /// Allows for creating and managing hardcoded claim protocol mappers within
    /// Keycloak.
    /// 
    /// Hardcoded claim protocol mappers allow you to define a claim with a hardcoded
    /// value. Protocol mappers can be defined for a single client, or they can
    /// be defined for a client scope which can be shared between multiple different
    /// clients.
    /// 
    /// ### Example Usage (Client)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var realm = new Keycloak.Realm("realm", new Keycloak.RealmArgs
    ///         {
    ///             Enabled = true,
    ///             Realm = "my-realm",
    ///         });
    ///         var openidClient = new Keycloak.OpenId.Client("openidClient", new Keycloak.OpenId.ClientArgs
    ///         {
    ///             AccessType = "CONFIDENTIAL",
    ///             ClientId = "test-client",
    ///             Enabled = true,
    ///             RealmId = realm.Id,
    ///             ValidRedirectUris = 
    ///             {
    ///                 "http://localhost:8080/openid-callback",
    ///             },
    ///         });
    ///         var hardcodedClaimMapper = new Keycloak.OpenId.HardcodedClaimProtocolMapper("hardcodedClaimMapper", new Keycloak.OpenId.HardcodedClaimProtocolMapperArgs
    ///         {
    ///             ClaimName = "foo",
    ///             ClaimValue = "bar",
    ///             ClientId = openidClient.Id,
    ///             RealmId = realm.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Example Usage (Client Scope)
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var realm = new Keycloak.Realm("realm", new Keycloak.RealmArgs
    ///         {
    ///             Enabled = true,
    ///             Realm = "my-realm",
    ///         });
    ///         var clientScope = new Keycloak.OpenId.ClientScope("clientScope", new Keycloak.OpenId.ClientScopeArgs
    ///         {
    ///             RealmId = realm.Id,
    ///         });
    ///         var hardcodedClaimMapper = new Keycloak.OpenId.HardcodedClaimProtocolMapper("hardcodedClaimMapper", new Keycloak.OpenId.HardcodedClaimProtocolMapperArgs
    ///         {
    ///             ClaimName = "foo",
    ///             ClaimValue = "bar",
    ///             ClientScopeId = clientScope.Id,
    ///             RealmId = realm.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this protocol mapper exists within.
    /// - `client_id` - (Required if `client_scope_id` is not specified) The client this protocol mapper is attached to.
    /// - `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
    /// - `name` - (Required) The display name of this protocol mapper in the GUI.
    /// - `claim_name` - (Required) The name of the claim to insert into a token.
    /// - `claim_value` - (Required) The hardcoded value of the claim.
    /// - `claim_value_type` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
    /// - `add_to_id_token` - (Optional) Indicates if the property should be added as a claim to the id token. Defaults to `true`.
    /// - `add_to_access_token` - (Optional) Indicates if the property should be added as a claim to the access token. Defaults to `true`.
    /// - `add_to_userinfo` - (Optional) Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
    /// </summary>
    public partial class HardcodedClaimProtocolMapper : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates if the attribute should be a claim in the access token.
        /// </summary>
        [Output("addToAccessToken")]
        public Output<bool?> AddToAccessToken { get; private set; } = null!;

        /// <summary>
        /// Indicates if the attribute should be a claim in the id token.
        /// </summary>
        [Output("addToIdToken")]
        public Output<bool?> AddToIdToken { get; private set; } = null!;

        /// <summary>
        /// Indicates if the attribute should appear in the userinfo response body.
        /// </summary>
        [Output("addToUserinfo")]
        public Output<bool?> AddToUserinfo { get; private set; } = null!;

        [Output("claimName")]
        public Output<string> ClaimName { get; private set; } = null!;

        [Output("claimValue")]
        public Output<string> ClaimValue { get; private set; } = null!;

        /// <summary>
        /// Claim type used when serializing tokens.
        /// </summary>
        [Output("claimValueType")]
        public Output<string?> ClaimValueType { get; private set; } = null!;

        /// <summary>
        /// The mapper's associated client. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The mapper's associated client scope. Cannot be used at the same time as client_id.
        /// </summary>
        [Output("clientScopeId")]
        public Output<string?> ClientScopeId { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a HardcodedClaimProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HardcodedClaimProtocolMapper(string name, HardcodedClaimProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper", name, args ?? new HardcodedClaimProtocolMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HardcodedClaimProtocolMapper(string name, Input<string> id, HardcodedClaimProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HardcodedClaimProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HardcodedClaimProtocolMapper Get(string name, Input<string> id, HardcodedClaimProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new HardcodedClaimProtocolMapper(name, id, state, options);
        }
    }

    public sealed class HardcodedClaimProtocolMapperArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the attribute should be a claim in the access token.
        /// </summary>
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        /// <summary>
        /// Indicates if the attribute should be a claim in the id token.
        /// </summary>
        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

        /// <summary>
        /// Indicates if the attribute should appear in the userinfo response body.
        /// </summary>
        [Input("addToUserinfo")]
        public Input<bool>? AddToUserinfo { get; set; }

        [Input("claimName", required: true)]
        public Input<string> ClaimName { get; set; } = null!;

        [Input("claimValue", required: true)]
        public Input<string> ClaimValue { get; set; } = null!;

        /// <summary>
        /// Claim type used when serializing tokens.
        /// </summary>
        [Input("claimValueType")]
        public Input<string>? ClaimValueType { get; set; }

        /// <summary>
        /// The mapper's associated client. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The mapper's associated client scope. Cannot be used at the same time as client_id.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public HardcodedClaimProtocolMapperArgs()
        {
        }
    }

    public sealed class HardcodedClaimProtocolMapperState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the attribute should be a claim in the access token.
        /// </summary>
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        /// <summary>
        /// Indicates if the attribute should be a claim in the id token.
        /// </summary>
        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

        /// <summary>
        /// Indicates if the attribute should appear in the userinfo response body.
        /// </summary>
        [Input("addToUserinfo")]
        public Input<bool>? AddToUserinfo { get; set; }

        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        [Input("claimValue")]
        public Input<string>? ClaimValue { get; set; }

        /// <summary>
        /// Claim type used when serializing tokens.
        /// </summary>
        [Input("claimValueType")]
        public Input<string>? ClaimValueType { get; set; }

        /// <summary>
        /// The mapper's associated client. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The mapper's associated client scope. Cannot be used at the same time as client_id.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public HardcodedClaimProtocolMapperState()
        {
        }
    }
}
