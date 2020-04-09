// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientServiceAccountRealmRole struct {
	pulumi.CustomResourceState

	RealmId              pulumi.StringOutput `pulumi:"realmId"`
	Role                 pulumi.StringOutput `pulumi:"role"`
	ServiceAccountUserId pulumi.StringOutput `pulumi:"serviceAccountUserId"`
}

// NewClientServiceAccountRealmRole registers a new resource with the given unique name, arguments, and options.
func NewClientServiceAccountRealmRole(ctx *pulumi.Context,
	name string, args *ClientServiceAccountRealmRoleArgs, opts ...pulumi.ResourceOption) (*ClientServiceAccountRealmRole, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil || args.ServiceAccountUserId == nil {
		return nil, errors.New("missing required argument 'ServiceAccountUserId'")
	}
	if args == nil {
		args = &ClientServiceAccountRealmRoleArgs{}
	}
	var resource ClientServiceAccountRealmRole
	err := ctx.RegisterResource("keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientServiceAccountRealmRole gets an existing ClientServiceAccountRealmRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientServiceAccountRealmRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientServiceAccountRealmRoleState, opts ...pulumi.ResourceOption) (*ClientServiceAccountRealmRole, error) {
	var resource ClientServiceAccountRealmRole
	err := ctx.ReadResource("keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientServiceAccountRealmRole resources.
type clientServiceAccountRealmRoleState struct {
	RealmId              *string `pulumi:"realmId"`
	Role                 *string `pulumi:"role"`
	ServiceAccountUserId *string `pulumi:"serviceAccountUserId"`
}

type ClientServiceAccountRealmRoleState struct {
	RealmId              pulumi.StringPtrInput
	Role                 pulumi.StringPtrInput
	ServiceAccountUserId pulumi.StringPtrInput
}

func (ClientServiceAccountRealmRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientServiceAccountRealmRoleState)(nil)).Elem()
}

type clientServiceAccountRealmRoleArgs struct {
	RealmId              string `pulumi:"realmId"`
	Role                 string `pulumi:"role"`
	ServiceAccountUserId string `pulumi:"serviceAccountUserId"`
}

// The set of arguments for constructing a ClientServiceAccountRealmRole resource.
type ClientServiceAccountRealmRoleArgs struct {
	RealmId              pulumi.StringInput
	Role                 pulumi.StringInput
	ServiceAccountUserId pulumi.StringInput
}

func (ClientServiceAccountRealmRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientServiceAccountRealmRoleArgs)(nil)).Elem()
}
