// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    /// <summary>
    /// ## # keycloak..AttributeImporterIdentityProviderMapper
    /// 
    /// Allows to create and manage identity provider mappers within Keycloak.
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm` - (Required) The name of the realm.
    /// - `name` - (Required) The name of the mapper.
    /// - `identity_provider_alias` - (Required) The alias of the associated identity provider.
    /// - `user_attribute` - (Required) The user attribute name to store SAML attribute.
    /// - `attribute_name` - (Optional) The Name of attribute to search for in assertion. You can leave this blank and specify a friendly name instead.
    /// - `attribute_friendly_name` - (Optional) The friendly name of attribute to search for in assertion.  You can leave this blank and specify an attribute name instead.
    /// - `claim_name` - (Optional) The claim name.
    /// 
    /// &gt; This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_attribute_importer_identity_provider_mapper.html.markdown.
    /// </summary>
    public partial class AttributeImporterIdentityProviderMapper : Pulumi.CustomResource
    {
        /// <summary>
        /// Attribute Friendly Name
        /// </summary>
        [Output("attributeFriendlyName")]
        public Output<string?> AttributeFriendlyName { get; private set; } = null!;

        /// <summary>
        /// Attribute Name
        /// </summary>
        [Output("attributeName")]
        public Output<string?> AttributeName { get; private set; } = null!;

        /// <summary>
        /// Claim Name
        /// </summary>
        [Output("claimName")]
        public Output<string?> ClaimName { get; private set; } = null!;

        /// <summary>
        /// IDP Alias
        /// </summary>
        [Output("identityProviderAlias")]
        public Output<string> IdentityProviderAlias { get; private set; } = null!;

        /// <summary>
        /// IDP Mapper Name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Realm Name
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// User Attribute
        /// </summary>
        [Output("userAttribute")]
        public Output<string> UserAttribute { get; private set; } = null!;


        /// <summary>
        /// Create a AttributeImporterIdentityProviderMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AttributeImporterIdentityProviderMapper(string name, AttributeImporterIdentityProviderMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AttributeImporterIdentityProviderMapper(string name, Input<string> id, AttributeImporterIdentityProviderMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AttributeImporterIdentityProviderMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AttributeImporterIdentityProviderMapper Get(string name, Input<string> id, AttributeImporterIdentityProviderMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new AttributeImporterIdentityProviderMapper(name, id, state, options);
        }
    }

    public sealed class AttributeImporterIdentityProviderMapperArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Attribute Friendly Name
        /// </summary>
        [Input("attributeFriendlyName")]
        public Input<string>? AttributeFriendlyName { get; set; }

        /// <summary>
        /// Attribute Name
        /// </summary>
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        /// <summary>
        /// Claim Name
        /// </summary>
        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        /// <summary>
        /// IDP Alias
        /// </summary>
        [Input("identityProviderAlias", required: true)]
        public Input<string> IdentityProviderAlias { get; set; } = null!;

        /// <summary>
        /// IDP Mapper Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Realm Name
        /// </summary>
        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        /// <summary>
        /// User Attribute
        /// </summary>
        [Input("userAttribute", required: true)]
        public Input<string> UserAttribute { get; set; } = null!;

        public AttributeImporterIdentityProviderMapperArgs()
        {
        }
    }

    public sealed class AttributeImporterIdentityProviderMapperState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Attribute Friendly Name
        /// </summary>
        [Input("attributeFriendlyName")]
        public Input<string>? AttributeFriendlyName { get; set; }

        /// <summary>
        /// Attribute Name
        /// </summary>
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        /// <summary>
        /// Claim Name
        /// </summary>
        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        /// <summary>
        /// IDP Alias
        /// </summary>
        [Input("identityProviderAlias")]
        public Input<string>? IdentityProviderAlias { get; set; }

        /// <summary>
        /// IDP Mapper Name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Realm Name
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// User Attribute
        /// </summary>
        [Input("userAttribute")]
        public Input<string>? UserAttribute { get; set; }

        public AttributeImporterIdentityProviderMapperState()
        {
        }
    }
}
