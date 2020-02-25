// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package Saml

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # Saml.UserAttributeProtocolMapper
// 
// Allows for creating and managing user attribute protocol mappers for
// SAML clients within Keycloak.
// 
// SAML user attribute protocol mappers allow you to map custom attributes defined
// for a user within Keycloak to an attribute in a SAML assertion. Protocol mappers
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
// - `userAttribute` - (Required) The custom user attribute to map.
// - `friendlyName` - (Optional) An optional human-friendly name for this attribute.
// - `samlAttributeName` - (Required) The name of the SAML attribute.
// - `samlAttributeNameFormat` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
// 
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/saml_user_attribute_protocol_mapper.html.markdown.
type UserAttributeProtocolMapper struct {
	pulumi.CustomResourceState

	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	Name pulumi.StringOutput `pulumi:"name"`
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	SamlAttributeName pulumi.StringOutput `pulumi:"samlAttributeName"`
	SamlAttributeNameFormat pulumi.StringOutput `pulumi:"samlAttributeNameFormat"`
	UserAttribute pulumi.StringOutput `pulumi:"userAttribute"`
}

// NewUserAttributeProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserAttributeProtocolMapper(ctx *pulumi.Context,
	name string, args *UserAttributeProtocolMapperArgs, opts ...pulumi.ResourceOption) (*UserAttributeProtocolMapper, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.SamlAttributeName == nil {
		return nil, errors.New("missing required argument 'SamlAttributeName'")
	}
	if args == nil || args.SamlAttributeNameFormat == nil {
		return nil, errors.New("missing required argument 'SamlAttributeNameFormat'")
	}
	if args == nil || args.UserAttribute == nil {
		return nil, errors.New("missing required argument 'UserAttribute'")
	}
	if args == nil {
		args = &UserAttributeProtocolMapperArgs{}
	}
	var resource UserAttributeProtocolMapper
	err := ctx.RegisterResource("keycloak:Saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserAttributeProtocolMapper gets an existing UserAttributeProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserAttributeProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserAttributeProtocolMapperState, opts ...pulumi.ResourceOption) (*UserAttributeProtocolMapper, error) {
	var resource UserAttributeProtocolMapper
	err := ctx.ReadResource("keycloak:Saml/userAttributeProtocolMapper:UserAttributeProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserAttributeProtocolMapper resources.
type userAttributeProtocolMapperState struct {
	ClientId *string `pulumi:"clientId"`
	ClientScopeId *string `pulumi:"clientScopeId"`
	FriendlyName *string `pulumi:"friendlyName"`
	Name *string `pulumi:"name"`
	RealmId *string `pulumi:"realmId"`
	SamlAttributeName *string `pulumi:"samlAttributeName"`
	SamlAttributeNameFormat *string `pulumi:"samlAttributeNameFormat"`
	UserAttribute *string `pulumi:"userAttribute"`
}

type UserAttributeProtocolMapperState struct {
	ClientId pulumi.StringPtrInput
	ClientScopeId pulumi.StringPtrInput
	FriendlyName pulumi.StringPtrInput
	Name pulumi.StringPtrInput
	RealmId pulumi.StringPtrInput
	SamlAttributeName pulumi.StringPtrInput
	SamlAttributeNameFormat pulumi.StringPtrInput
	UserAttribute pulumi.StringPtrInput
}

func (UserAttributeProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userAttributeProtocolMapperState)(nil)).Elem()
}

type userAttributeProtocolMapperArgs struct {
	ClientId *string `pulumi:"clientId"`
	ClientScopeId *string `pulumi:"clientScopeId"`
	FriendlyName *string `pulumi:"friendlyName"`
	Name *string `pulumi:"name"`
	RealmId string `pulumi:"realmId"`
	SamlAttributeName string `pulumi:"samlAttributeName"`
	SamlAttributeNameFormat string `pulumi:"samlAttributeNameFormat"`
	UserAttribute string `pulumi:"userAttribute"`
}

// The set of arguments for constructing a UserAttributeProtocolMapper resource.
type UserAttributeProtocolMapperArgs struct {
	ClientId pulumi.StringPtrInput
	ClientScopeId pulumi.StringPtrInput
	FriendlyName pulumi.StringPtrInput
	Name pulumi.StringPtrInput
	RealmId pulumi.StringInput
	SamlAttributeName pulumi.StringInput
	SamlAttributeNameFormat pulumi.StringInput
	UserAttribute pulumi.StringInput
}

func (UserAttributeProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userAttributeProtocolMapperArgs)(nil)).Elem()
}

