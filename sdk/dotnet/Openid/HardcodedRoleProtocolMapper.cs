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
    /// ## # keycloak.openid.HardcodedRoleProtocolMapper
    /// 
    /// Allows for creating and managing hardcoded role protocol mappers within
    /// Keycloak.
    /// 
    /// Hardcoded role protocol mappers allow you to specify a single role to
    /// always map to an access token for a client. Protocol mappers can be
    /// defined for a single client, or they can be defined for a client scope
    /// which can be shared between multiple different clients.
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this protocol mapper exists within.
    /// - `client_id` - (Required if `client_scope_id` is not specified) The client this protocol mapper is attached to.
    /// - `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
    /// - `name` - (Required) The display name of this protocol mapper in the
    ///   GUI.
    /// - `role_id` - (Required) The ID of the role to map to an access token.
    /// 
    /// &gt; This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_openid_hardcoded_role_protocol_mapper.html.markdown.
    /// </summary>
    public partial class HardcodedRoleProtocolMapper : Pulumi.CustomResource
    {
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

        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;


        /// <summary>
        /// Create a HardcodedRoleProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HardcodedRoleProtocolMapper(string name, HardcodedRoleProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private HardcodedRoleProtocolMapper(string name, Input<string> id, HardcodedRoleProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/hardcodedRoleProtocolMapper:HardcodedRoleProtocolMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HardcodedRoleProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HardcodedRoleProtocolMapper Get(string name, Input<string> id, HardcodedRoleProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new HardcodedRoleProtocolMapper(name, id, state, options);
        }
    }

    public sealed class HardcodedRoleProtocolMapperArgs : Pulumi.ResourceArgs
    {
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

        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        public HardcodedRoleProtocolMapperArgs()
        {
        }
    }

    public sealed class HardcodedRoleProtocolMapperState : Pulumi.ResourceArgs
    {
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

        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        public HardcodedRoleProtocolMapperState()
        {
        }
    }
}
