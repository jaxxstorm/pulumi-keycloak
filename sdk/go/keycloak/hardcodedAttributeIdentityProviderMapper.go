// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type HardcodedAttributeIdentityProviderMapper struct {
	pulumi.CustomResourceState

	// OIDC Claim
	AttributeName pulumi.StringPtrOutput `pulumi:"attributeName"`
	// User Attribute
	AttributeValue pulumi.StringPtrOutput `pulumi:"attributeValue"`
	// IDP Alias
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Realm Name
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Is Attribute Related To a User Session
	UserSession pulumi.BoolOutput `pulumi:"userSession"`
}

// NewHardcodedAttributeIdentityProviderMapper registers a new resource with the given unique name, arguments, and options.
func NewHardcodedAttributeIdentityProviderMapper(ctx *pulumi.Context,
	name string, args *HardcodedAttributeIdentityProviderMapperArgs, opts ...pulumi.ResourceOption) (*HardcodedAttributeIdentityProviderMapper, error) {
	if args == nil || args.IdentityProviderAlias == nil {
		return nil, errors.New("missing required argument 'IdentityProviderAlias'")
	}
	if args == nil || args.Realm == nil {
		return nil, errors.New("missing required argument 'Realm'")
	}
	if args == nil || args.UserSession == nil {
		return nil, errors.New("missing required argument 'UserSession'")
	}
	if args == nil {
		args = &HardcodedAttributeIdentityProviderMapperArgs{}
	}
	var resource HardcodedAttributeIdentityProviderMapper
	err := ctx.RegisterResource("keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHardcodedAttributeIdentityProviderMapper gets an existing HardcodedAttributeIdentityProviderMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHardcodedAttributeIdentityProviderMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HardcodedAttributeIdentityProviderMapperState, opts ...pulumi.ResourceOption) (*HardcodedAttributeIdentityProviderMapper, error) {
	var resource HardcodedAttributeIdentityProviderMapper
	err := ctx.ReadResource("keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HardcodedAttributeIdentityProviderMapper resources.
type hardcodedAttributeIdentityProviderMapperState struct {
	// OIDC Claim
	AttributeName *string `pulumi:"attributeName"`
	// User Attribute
	AttributeValue *string `pulumi:"attributeValue"`
	// IDP Alias
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm *string `pulumi:"realm"`
	// Is Attribute Related To a User Session
	UserSession *bool `pulumi:"userSession"`
}

type HardcodedAttributeIdentityProviderMapperState struct {
	// OIDC Claim
	AttributeName pulumi.StringPtrInput
	// User Attribute
	AttributeValue pulumi.StringPtrInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringPtrInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringPtrInput
	// Is Attribute Related To a User Session
	UserSession pulumi.BoolPtrInput
}

func (HardcodedAttributeIdentityProviderMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedAttributeIdentityProviderMapperState)(nil)).Elem()
}

type hardcodedAttributeIdentityProviderMapperArgs struct {
	// OIDC Claim
	AttributeName *string `pulumi:"attributeName"`
	// User Attribute
	AttributeValue *string `pulumi:"attributeValue"`
	// IDP Alias
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm string `pulumi:"realm"`
	// Is Attribute Related To a User Session
	UserSession bool `pulumi:"userSession"`
}

// The set of arguments for constructing a HardcodedAttributeIdentityProviderMapper resource.
type HardcodedAttributeIdentityProviderMapperArgs struct {
	// OIDC Claim
	AttributeName pulumi.StringPtrInput
	// User Attribute
	AttributeValue pulumi.StringPtrInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringInput
	// Is Attribute Related To a User Session
	UserSession pulumi.BoolInput
}

func (HardcodedAttributeIdentityProviderMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hardcodedAttributeIdentityProviderMapperArgs)(nil)).Elem()
}
