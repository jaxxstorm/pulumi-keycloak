// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public partial class ClientGroupPolicy : Pulumi.CustomResource
    {
        [Output("decisionStrategy")]
        public Output<string> DecisionStrategy { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("groups")]
        public Output<ImmutableArray<Outputs.ClientGroupPolicyGroups>> Groups { get; private set; } = null!;

        [Output("groupsClaim")]
        public Output<string?> GroupsClaim { get; private set; } = null!;

        [Output("logic")]
        public Output<string?> Logic { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("resourceServerId")]
        public Output<string> ResourceServerId { get; private set; } = null!;


        /// <summary>
        /// Create a ClientGroupPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientGroupPolicy(string name, ClientGroupPolicyArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientGroupPolicy:ClientGroupPolicy", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ClientGroupPolicy(string name, Input<string> id, ClientGroupPolicyState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientGroupPolicy:ClientGroupPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientGroupPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientGroupPolicy Get(string name, Input<string> id, ClientGroupPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientGroupPolicy(name, id, state, options);
        }
    }

    public sealed class ClientGroupPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("decisionStrategy", required: true)]
        public Input<string> DecisionStrategy { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("groups", required: true)]
        private InputList<Inputs.ClientGroupPolicyGroupsArgs>? _groups;
        public InputList<Inputs.ClientGroupPolicyGroupsArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.ClientGroupPolicyGroupsArgs>());
            set => _groups = value;
        }

        [Input("groupsClaim")]
        public Input<string>? GroupsClaim { get; set; }

        [Input("logic")]
        public Input<string>? Logic { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("resourceServerId", required: true)]
        public Input<string> ResourceServerId { get; set; } = null!;

        public ClientGroupPolicyArgs()
        {
        }
    }

    public sealed class ClientGroupPolicyState : Pulumi.ResourceArgs
    {
        [Input("decisionStrategy")]
        public Input<string>? DecisionStrategy { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("groups")]
        private InputList<Inputs.ClientGroupPolicyGroupsGetArgs>? _groups;
        public InputList<Inputs.ClientGroupPolicyGroupsGetArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.ClientGroupPolicyGroupsGetArgs>());
            set => _groups = value;
        }

        [Input("groupsClaim")]
        public Input<string>? GroupsClaim { get; set; }

        [Input("logic")]
        public Input<string>? Logic { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("resourceServerId")]
        public Input<string>? ResourceServerId { get; set; }

        public ClientGroupPolicyState()
        {
        }
    }

    namespace Inputs
    {

    public sealed class ClientGroupPolicyGroupsArgs : Pulumi.ResourceArgs
    {
        [Input("extendChildren", required: true)]
        public Input<bool> ExtendChildren { get; set; } = null!;

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public ClientGroupPolicyGroupsArgs()
        {
        }
    }

    public sealed class ClientGroupPolicyGroupsGetArgs : Pulumi.ResourceArgs
    {
        [Input("extendChildren", required: true)]
        public Input<bool> ExtendChildren { get; set; } = null!;

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public ClientGroupPolicyGroupsGetArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class ClientGroupPolicyGroups
    {
        public readonly bool ExtendChildren;
        public readonly string Id;
        public readonly string Path;

        [OutputConstructor]
        private ClientGroupPolicyGroups(
            bool extendChildren,
            string id,
            string path)
        {
            ExtendChildren = extendChildren;
            Id = id;
            Path = path;
        }
    }
    }
}
