// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GenericClientRoleMapper struct {
	pulumi.CustomResourceState

	// The destination client of the client role. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The destination client scope of the client role. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrOutput `pulumi:"clientScopeId"`
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// Id of the role to assign
	RoleId pulumi.StringOutput `pulumi:"roleId"`
}

// NewGenericClientRoleMapper registers a new resource with the given unique name, arguments, and options.
func NewGenericClientRoleMapper(ctx *pulumi.Context,
	name string, args *GenericClientRoleMapperArgs, opts ...pulumi.ResourceOption) (*GenericClientRoleMapper, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.RoleId == nil {
		return nil, errors.New("missing required argument 'RoleId'")
	}
	if args == nil {
		args = &GenericClientRoleMapperArgs{}
	}
	var resource GenericClientRoleMapper
	err := ctx.RegisterResource("keycloak:index/genericClientRoleMapper:GenericClientRoleMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGenericClientRoleMapper gets an existing GenericClientRoleMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGenericClientRoleMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GenericClientRoleMapperState, opts ...pulumi.ResourceOption) (*GenericClientRoleMapper, error) {
	var resource GenericClientRoleMapper
	err := ctx.ReadResource("keycloak:index/genericClientRoleMapper:GenericClientRoleMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GenericClientRoleMapper resources.
type genericClientRoleMapperState struct {
	// The destination client of the client role. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The destination client scope of the client role. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The realm id where the associated client or client scope exists.
	RealmId *string `pulumi:"realmId"`
	// Id of the role to assign
	RoleId *string `pulumi:"roleId"`
}

type GenericClientRoleMapperState struct {
	// The destination client of the client role. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The destination client scope of the client role. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringPtrInput
	// Id of the role to assign
	RoleId pulumi.StringPtrInput
}

func (GenericClientRoleMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*genericClientRoleMapperState)(nil)).Elem()
}

type genericClientRoleMapperArgs struct {
	// The destination client of the client role. Cannot be used at the same time as client_scope_id.
	ClientId *string `pulumi:"clientId"`
	// The destination client scope of the client role. Cannot be used at the same time as client_id.
	ClientScopeId *string `pulumi:"clientScopeId"`
	// The realm id where the associated client or client scope exists.
	RealmId string `pulumi:"realmId"`
	// Id of the role to assign
	RoleId string `pulumi:"roleId"`
}

// The set of arguments for constructing a GenericClientRoleMapper resource.
type GenericClientRoleMapperArgs struct {
	// The destination client of the client role. Cannot be used at the same time as client_scope_id.
	ClientId pulumi.StringPtrInput
	// The destination client scope of the client role. Cannot be used at the same time as client_id.
	ClientScopeId pulumi.StringPtrInput
	// The realm id where the associated client or client scope exists.
	RealmId pulumi.StringInput
	// Id of the role to assign
	RoleId pulumi.StringInput
}

func (GenericClientRoleMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*genericClientRoleMapperArgs)(nil)).Elem()
}
