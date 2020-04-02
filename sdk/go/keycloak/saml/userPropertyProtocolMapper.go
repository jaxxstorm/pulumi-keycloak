// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package saml

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # saml.UserPropertyProtocolMapper
//
// Allows for creating and managing user property protocol mappers for
// SAML clients within Keycloak.
//
// SAML user property protocol mappers allow you to map properties of the Keycloak
// user model to an attribute in a SAML assertion. Protocol mappers
// can be defined for a single client, or they can be defined for a client scope which
// can be shared between multiple different clients.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm this protocol mapper exists within.
// - `clientId` - (Required if `clientScopeId` is not specified) The SAML client this protocol mapper is attached to.
// - `clientScopeId` - (Required if `clientId` is not specified) The SAML client scope this protocol mapper is attached to.
// - `name` - (Required) The display name of this protocol mapper in the GUI.
// - `userProperty` - (Required) The property of the Keycloak user model to map.
// - `friendlyName` - (Optional) An optional human-friendly name for this attribute.
// - `samlAttributeName` - (Required) The name of the SAML attribute.
// - `samlAttributeNameFormat` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_saml_user_property_protocol_mapper.html.markdown.
type UserPropertyProtocolMapper struct {
	pulumi.CustomResourceState

	ClientId                pulumi.StringPtrOutput `pulumi:"clientId"`
	ClientScopeId           pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	FriendlyName            pulumi.StringPtrOutput `pulumi:"friendlyName"`
	Name                    pulumi.StringOutput    `pulumi:"name"`
	RealmId                 pulumi.StringOutput    `pulumi:"realmId"`
	SamlAttributeName       pulumi.StringOutput    `pulumi:"samlAttributeName"`
	SamlAttributeNameFormat pulumi.StringOutput    `pulumi:"samlAttributeNameFormat"`
	UserProperty            pulumi.StringOutput    `pulumi:"userProperty"`
}

// NewUserPropertyProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserPropertyProtocolMapper(ctx *pulumi.Context,
	name string, args *UserPropertyProtocolMapperArgs, opts ...pulumi.ResourceOption) (*UserPropertyProtocolMapper, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.SamlAttributeName == nil {
		return nil, errors.New("missing required argument 'SamlAttributeName'")
	}
	if args == nil || args.SamlAttributeNameFormat == nil {
		return nil, errors.New("missing required argument 'SamlAttributeNameFormat'")
	}
	if args == nil || args.UserProperty == nil {
		return nil, errors.New("missing required argument 'UserProperty'")
	}
	if args == nil {
		args = &UserPropertyProtocolMapperArgs{}
	}
	var resource UserPropertyProtocolMapper
	err := ctx.RegisterResource("keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPropertyProtocolMapper gets an existing UserPropertyProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPropertyProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPropertyProtocolMapperState, opts ...pulumi.ResourceOption) (*UserPropertyProtocolMapper, error) {
	var resource UserPropertyProtocolMapper
	err := ctx.ReadResource("keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPropertyProtocolMapper resources.
type userPropertyProtocolMapperState struct {
	ClientId                *string `pulumi:"clientId"`
	ClientScopeId           *string `pulumi:"clientScopeId"`
	FriendlyName            *string `pulumi:"friendlyName"`
	Name                    *string `pulumi:"name"`
	RealmId                 *string `pulumi:"realmId"`
	SamlAttributeName       *string `pulumi:"samlAttributeName"`
	SamlAttributeNameFormat *string `pulumi:"samlAttributeNameFormat"`
	UserProperty            *string `pulumi:"userProperty"`
}

type UserPropertyProtocolMapperState struct {
	ClientId                pulumi.StringPtrInput
	ClientScopeId           pulumi.StringPtrInput
	FriendlyName            pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	RealmId                 pulumi.StringPtrInput
	SamlAttributeName       pulumi.StringPtrInput
	SamlAttributeNameFormat pulumi.StringPtrInput
	UserProperty            pulumi.StringPtrInput
}

func (UserPropertyProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPropertyProtocolMapperState)(nil)).Elem()
}

type userPropertyProtocolMapperArgs struct {
	ClientId                *string `pulumi:"clientId"`
	ClientScopeId           *string `pulumi:"clientScopeId"`
	FriendlyName            *string `pulumi:"friendlyName"`
	Name                    *string `pulumi:"name"`
	RealmId                 string  `pulumi:"realmId"`
	SamlAttributeName       string  `pulumi:"samlAttributeName"`
	SamlAttributeNameFormat string  `pulumi:"samlAttributeNameFormat"`
	UserProperty            string  `pulumi:"userProperty"`
}

// The set of arguments for constructing a UserPropertyProtocolMapper resource.
type UserPropertyProtocolMapperArgs struct {
	ClientId                pulumi.StringPtrInput
	ClientScopeId           pulumi.StringPtrInput
	FriendlyName            pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	RealmId                 pulumi.StringInput
	SamlAttributeName       pulumi.StringInput
	SamlAttributeNameFormat pulumi.StringInput
	UserProperty            pulumi.StringInput
}

func (UserPropertyProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPropertyProtocolMapperArgs)(nil)).Elem()
}
