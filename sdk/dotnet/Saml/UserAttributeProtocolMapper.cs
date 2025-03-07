// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Saml
{
    /// <summary>
    /// ## # keycloak.saml.UserAttributeProtocolMapper
    /// 
    /// Allows for creating and managing user attribute protocol mappers for
    /// SAML clients within Keycloak.
    /// 
    /// SAML user attribute protocol mappers allow you to map custom attributes defined
    /// for a user within Keycloak to an attribute in a SAML assertion. Protocol mappers
    /// can be defined for a single client, or they can be defined for a client scope which
    /// can be shared between multiple different clients.
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
    ///         var samlClient = new Keycloak.Saml.Client("samlClient", new Keycloak.Saml.ClientArgs
    ///         {
    ///             ClientId = "test-saml-client",
    ///             RealmId = keycloak_realm.Test.Id,
    ///         });
    ///         var samlUserAttributeMapper = new Keycloak.Saml.UserAttributeProtocolMapper("samlUserAttributeMapper", new Keycloak.Saml.UserAttributeProtocolMapperArgs
    ///         {
    ///             ClientId = samlClient.Id,
    ///             RealmId = keycloak_realm.Test.Id,
    ///             SamlAttributeName = "displayName",
    ///             SamlAttributeNameFormat = "Unspecified",
    ///             UserAttribute = "displayName",
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
    /// - `client_id` - (Required if `client_scope_id` is not specified) The SAML client this protocol mapper is attached to.
    /// - `client_scope_id` - (Required if `client_id` is not specified) The SAML client scope this protocol mapper is attached to.
    /// - `name` - (Required) The display name of this protocol mapper in the GUI.
    /// - `user_attribute` - (Required) The custom user attribute to map.
    /// - `friendly_name` - (Optional) An optional human-friendly name for this attribute.
    /// - `saml_attribute_name` - (Required) The name of the SAML attribute.
    /// - `saml_attribute_name_format` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
    /// </summary>
    public partial class UserAttributeProtocolMapper : Pulumi.CustomResource
    {
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        [Output("clientScopeId")]
        public Output<string?> ClientScopeId { get; private set; } = null!;

        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("samlAttributeName")]
        public Output<string> SamlAttributeName { get; private set; } = null!;

        [Output("samlAttributeNameFormat")]
        public Output<string> SamlAttributeNameFormat { get; private set; } = null!;

        [Output("userAttribute")]
        public Output<string> UserAttribute { get; private set; } = null!;


        /// <summary>
        /// Create a UserAttributeProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserAttributeProtocolMapper(string name, UserAttributeProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, args ?? new UserAttributeProtocolMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserAttributeProtocolMapper(string name, Input<string> id, UserAttributeProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserAttributeProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserAttributeProtocolMapper Get(string name, Input<string> id, UserAttributeProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new UserAttributeProtocolMapper(name, id, state, options);
        }
    }

    public sealed class UserAttributeProtocolMapperArgs : Pulumi.ResourceArgs
    {
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("samlAttributeName", required: true)]
        public Input<string> SamlAttributeName { get; set; } = null!;

        [Input("samlAttributeNameFormat", required: true)]
        public Input<string> SamlAttributeNameFormat { get; set; } = null!;

        [Input("userAttribute", required: true)]
        public Input<string> UserAttribute { get; set; } = null!;

        public UserAttributeProtocolMapperArgs()
        {
        }
    }

    public sealed class UserAttributeProtocolMapperState : Pulumi.ResourceArgs
    {
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("samlAttributeName")]
        public Input<string>? SamlAttributeName { get; set; }

        [Input("samlAttributeNameFormat")]
        public Input<string>? SamlAttributeNameFormat { get; set; }

        [Input("userAttribute")]
        public Input<string>? UserAttribute { get; set; }

        public UserAttributeProtocolMapperState()
        {
        }
    }
}
