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
    /// ## # keycloak.openid.AudienceProtocolMapper
    /// 
    /// Allows for creating and managing audience protocol mappers within
    /// Keycloak. This mapper was added in Keycloak v4.6.0.Final.
    /// 
    /// Audience protocol mappers allow you add audiences to the `aud` claim
    /// within issued tokens. The audience can be a custom string, or it can be
    /// mapped to the ID of a pre-existing client.
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
    ///         var audienceMapper = new Keycloak.OpenId.AudienceProtocolMapper("audienceMapper", new Keycloak.OpenId.AudienceProtocolMapperArgs
    ///         {
    ///             ClientId = openidClient.Id,
    ///             IncludedCustomAudience = "foo",
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
    ///         var audienceMapper = new Keycloak.OpenId.AudienceProtocolMapper("audienceMapper", new Keycloak.OpenId.AudienceProtocolMapperArgs
    ///         {
    ///             ClientScopeId = clientScope.Id,
    ///             IncludedCustomAudience = "foo",
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
    /// - `included_client_audience` - (Required if `included_custom_audience` is not specified) A client ID to include within the token's `aud` claim.
    /// - `included_custom_audience` - (Required if `included_client_audience` is not specified) A custom audience to include within the token's `aud` claim.
    /// - `add_to_id_token` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
    /// - `add_to_access_token` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
    /// </summary>
    public partial class AudienceProtocolMapper : Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates if this claim should be added to the access token.
        /// </summary>
        [Output("addToAccessToken")]
        public Output<bool?> AddToAccessToken { get; private set; } = null!;

        /// <summary>
        /// Indicates if this claim should be added to the id token.
        /// </summary>
        [Output("addToIdToken")]
        public Output<bool?> AddToIdToken { get; private set; } = null!;

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
        /// A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        /// </summary>
        [Output("includedClientAudience")]
        public Output<string?> IncludedClientAudience { get; private set; } = null!;

        /// <summary>
        /// A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        /// </summary>
        [Output("includedCustomAudience")]
        public Output<string?> IncludedCustomAudience { get; private set; } = null!;

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
        /// Create a AudienceProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AudienceProtocolMapper(string name, AudienceProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper", name, args ?? new AudienceProtocolMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AudienceProtocolMapper(string name, Input<string> id, AudienceProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AudienceProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AudienceProtocolMapper Get(string name, Input<string> id, AudienceProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new AudienceProtocolMapper(name, id, state, options);
        }
    }

    public sealed class AudienceProtocolMapperArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if this claim should be added to the access token.
        /// </summary>
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        /// <summary>
        /// Indicates if this claim should be added to the id token.
        /// </summary>
        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

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
        /// A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        /// </summary>
        [Input("includedClientAudience")]
        public Input<string>? IncludedClientAudience { get; set; }

        /// <summary>
        /// A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        /// </summary>
        [Input("includedCustomAudience")]
        public Input<string>? IncludedCustomAudience { get; set; }

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

        public AudienceProtocolMapperArgs()
        {
        }
    }

    public sealed class AudienceProtocolMapperState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if this claim should be added to the access token.
        /// </summary>
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        /// <summary>
        /// Indicates if this claim should be added to the id token.
        /// </summary>
        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

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
        /// A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        /// </summary>
        [Input("includedClientAudience")]
        public Input<string>? IncludedClientAudience { get; set; }

        /// <summary>
        /// A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        /// </summary>
        [Input("includedCustomAudience")]
        public Input<string>? IncludedCustomAudience { get; set; }

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

        public AudienceProtocolMapperState()
        {
        }
    }
}
