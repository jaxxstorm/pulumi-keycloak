// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # .User
// 
// Allows for creating and managing Users within Keycloak.
// 
// This resource was created primarily to enable the acceptance tests for the `.Group` resource.
// Creating users within Keycloak is not recommended. Instead, users should be federated from external sources
// by configuring user federation providers or identity providers.
// 
// ### Argument Reference
// 
// The following arguments are supported:
// 
// - `realmId` - (Required) The realm this user belongs to.
// - `username` - (Required) The unique username of this user.
// - `initialPassword` (Optional) When given, the user's initial password will be set.
//    This attribute is only respected during initial user creation.
//     - `value` (Required) The initial password.
//     - `temporary` (Optional) If set to `true`, the initial password is set up for renewal on first use. Default to `false`.
// - `enabled` - (Optional) When false, this user cannot log in. Defaults to `true`.
// - `email` - (Optional) The user's email.
// - `firstName` - (Optional) The user's first name.
// - `lastName` - (Optional) The user's last name.
// 
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/user.html.markdown.
type User struct {
	pulumi.CustomResourceState

	Attributes pulumi.MapOutput `pulumi:"attributes"`
	Email pulumi.StringPtrOutput `pulumi:"email"`
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	FederatedIdentities UserFederatedIdentityArrayOutput `pulumi:"federatedIdentities"`
	FirstName pulumi.StringPtrOutput `pulumi:"firstName"`
	InitialPassword UserInitialPasswordPtrOutput `pulumi:"initialPassword"`
	LastName pulumi.StringPtrOutput `pulumi:"lastName"`
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.Username == nil {
		return nil, errors.New("missing required argument 'Username'")
	}
	if args == nil {
		args = &UserArgs{}
	}
	var resource User
	err := ctx.RegisterResource("keycloak:index/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("keycloak:index/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	Attributes map[string]interface{} `pulumi:"attributes"`
	Email *string `pulumi:"email"`
	Enabled *bool `pulumi:"enabled"`
	FederatedIdentities []UserFederatedIdentity `pulumi:"federatedIdentities"`
	FirstName *string `pulumi:"firstName"`
	InitialPassword *UserInitialPassword `pulumi:"initialPassword"`
	LastName *string `pulumi:"lastName"`
	RealmId *string `pulumi:"realmId"`
	Username *string `pulumi:"username"`
}

type UserState struct {
	Attributes pulumi.MapInput
	Email pulumi.StringPtrInput
	Enabled pulumi.BoolPtrInput
	FederatedIdentities UserFederatedIdentityArrayInput
	FirstName pulumi.StringPtrInput
	InitialPassword UserInitialPasswordPtrInput
	LastName pulumi.StringPtrInput
	RealmId pulumi.StringPtrInput
	Username pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	Attributes map[string]interface{} `pulumi:"attributes"`
	Email *string `pulumi:"email"`
	Enabled *bool `pulumi:"enabled"`
	FederatedIdentities []UserFederatedIdentity `pulumi:"federatedIdentities"`
	FirstName *string `pulumi:"firstName"`
	InitialPassword *UserInitialPassword `pulumi:"initialPassword"`
	LastName *string `pulumi:"lastName"`
	RealmId string `pulumi:"realmId"`
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	Attributes pulumi.MapInput
	Email pulumi.StringPtrInput
	Enabled pulumi.BoolPtrInput
	FederatedIdentities UserFederatedIdentityArrayInput
	FirstName pulumi.StringPtrInput
	InitialPassword UserInitialPasswordPtrInput
	LastName pulumi.StringPtrInput
	RealmId pulumi.StringInput
	Username pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

