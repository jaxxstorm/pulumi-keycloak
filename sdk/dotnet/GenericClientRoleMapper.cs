// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public partial class GenericClientRoleMapper : Pulumi.CustomResource
    {
        /// <summary>
        /// The destination client of the client role. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The destination client scope of the client role. Cannot be used at the same time as client_id.
        /// </summary>
        [Output("clientScopeId")]
        public Output<string?> ClientScopeId { get; private set; } = null!;

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// Id of the role to assign
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;


        /// <summary>
        /// Create a GenericClientRoleMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GenericClientRoleMapper(string name, GenericClientRoleMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/genericClientRoleMapper:GenericClientRoleMapper", name, args ?? new GenericClientRoleMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GenericClientRoleMapper(string name, Input<string> id, GenericClientRoleMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/genericClientRoleMapper:GenericClientRoleMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GenericClientRoleMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GenericClientRoleMapper Get(string name, Input<string> id, GenericClientRoleMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new GenericClientRoleMapper(name, id, state, options);
        }
    }

    public sealed class GenericClientRoleMapperArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination client of the client role. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The destination client scope of the client role. Cannot be used at the same time as client_id.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// Id of the role to assign
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        public GenericClientRoleMapperArgs()
        {
        }
    }

    public sealed class GenericClientRoleMapperState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination client of the client role. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The destination client scope of the client role. Cannot be used at the same time as client_id.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// Id of the role to assign
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        public GenericClientRoleMapperState()
        {
        }
    }
}
