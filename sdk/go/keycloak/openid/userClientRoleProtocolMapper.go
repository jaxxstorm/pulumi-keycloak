// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type UserClientRoleProtocolMapper struct {
	pulumi.CustomResourceState

	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrOutput `pulumi:"addToAccessToken"`
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrOutput `pulumi:"addToIdToken"`
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrOutput `pulumi:"addToUserinfo"`
	ClaimName     pulumi.StringOutput  `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrOutput `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// Client ID for role mappings.
	ClientIdForRoleMappings pulumi.StringPtrOutput `pulumi:"clientIdForRoleMappings"`
	// Prefix that will be added to each client role.
	ClientRolePrefix pulumi.StringPtrOutput `pulumi:"clientRolePrefix"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued pulumi.BoolPtrOutput `pulumi:"multivalued"`
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewUserClientRoleProtocolMapper registers a new resource with the given unique name, arguments, and options.
func NewUserClientRoleProtocolMapper(ctx *pulumi.Context,
	name string, args *UserClientRoleProtocolMapperArgs, opts ...pulumi.ResourceOption) (*UserClientRoleProtocolMapper, error) {
	if args == nil || args.ClaimName == nil {
		return nil, errors.New("missing required argument 'ClaimName'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &UserClientRoleProtocolMapperArgs{}
	}
	var resource UserClientRoleProtocolMapper
	err := ctx.RegisterResource("keycloak:openid/userClientRoleProtocolMapper:UserClientRoleProtocolMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserClientRoleProtocolMapper gets an existing UserClientRoleProtocolMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserClientRoleProtocolMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserClientRoleProtocolMapperState, opts ...pulumi.ResourceOption) (*UserClientRoleProtocolMapper, error) {
	var resource UserClientRoleProtocolMapper
	err := ctx.ReadResource("keycloak:openid/userClientRoleProtocolMapper:UserClientRoleProtocolMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserClientRoleProtocolMapper resources.
type userClientRoleProtocolMapperState struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo *bool   `pulumi:"addToUserinfo"`
	ClaimName     *string `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// Client ID for role mappings.
	ClientIdForRoleMappings *string `pulumi:"clientIdForRoleMappings"`
	// Prefix that will be added to each client role.
	ClientRolePrefix *string `pulumi:"clientRolePrefix"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued *bool `pulumi:"multivalued"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId *string `pulumi:"realmId"`
}

type UserClientRoleProtocolMapperState struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrInput
	ClaimName     pulumi.StringPtrInput
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// Client ID for role mappings.
	ClientIdForRoleMappings pulumi.StringPtrInput
	// Prefix that will be added to each client role.
	ClientRolePrefix pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringPtrInput
}

func (UserClientRoleProtocolMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userClientRoleProtocolMapperState)(nil)).Elem()
}

type userClientRoleProtocolMapperArgs struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken *bool `pulumi:"addToAccessToken"`
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken *bool `pulumi:"addToIdToken"`
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo *bool  `pulumi:"addToUserinfo"`
	ClaimName     string `pulumi:"claimName"`
	// Claim type used when serializing tokens.
	ClaimValueType *string `pulumi:"claimValueType"`
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// Client ID for role mappings.
	ClientIdForRoleMappings *string `pulumi:"clientIdForRoleMappings"`
	// Prefix that will be added to each client role.
	ClientRolePrefix *string `pulumi:"clientRolePrefix"`
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued *bool `pulumi:"multivalued"`
	// A human-friendly name that will appear in the Keycloak console.
	Name *string `pulumi:"name"`
	// The realm id where the associated client or client scope exists.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a UserClientRoleProtocolMapper resource.
type UserClientRoleProtocolMapperArgs struct {
	// Indicates if the attribute should be a claim in the access token.
	AddToAccessToken pulumi.BoolPtrInput
	// Indicates if the attribute should be a claim in the id token.
	AddToIdToken pulumi.BoolPtrInput
	// Indicates if the attribute should appear in the userinfo response body.
	AddToUserinfo pulumi.BoolPtrInput
	ClaimName     pulumi.StringInput
	// Claim type used when serializing tokens.
	ClaimValueType pulumi.StringPtrInput
	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// Client ID for role mappings.
	ClientIdForRoleMappings pulumi.StringPtrInput
	// Prefix that will be added to each client role.
	ClientRolePrefix pulumi.StringPtrInput
	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// Indicates whether this attribute is a single value or an array of values.
	Multivalued pulumi.BoolPtrInput
	// A human-friendly name that will appear in the Keycloak console.
	Name pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringInput
}

func (UserClientRoleProtocolMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userClientRoleProtocolMapperArgs)(nil)).Elem()
}
