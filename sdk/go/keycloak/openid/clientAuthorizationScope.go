// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientAuthorizationScope struct {
	pulumi.CustomResourceState

	DisplayName      pulumi.StringPtrOutput `pulumi:"displayName"`
	IconUri          pulumi.StringPtrOutput `pulumi:"iconUri"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	RealmId          pulumi.StringOutput    `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput    `pulumi:"resourceServerId"`
}

// NewClientAuthorizationScope registers a new resource with the given unique name, arguments, and options.
func NewClientAuthorizationScope(ctx *pulumi.Context,
	name string, args *ClientAuthorizationScopeArgs, opts ...pulumi.ResourceOption) (*ClientAuthorizationScope, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.ResourceServerId == nil {
		return nil, errors.New("missing required argument 'ResourceServerId'")
	}
	if args == nil {
		args = &ClientAuthorizationScopeArgs{}
	}
	var resource ClientAuthorizationScope
	err := ctx.RegisterResource("keycloak:openid/clientAuthorizationScope:ClientAuthorizationScope", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientAuthorizationScope gets an existing ClientAuthorizationScope resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientAuthorizationScope(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientAuthorizationScopeState, opts ...pulumi.ResourceOption) (*ClientAuthorizationScope, error) {
	var resource ClientAuthorizationScope
	err := ctx.ReadResource("keycloak:openid/clientAuthorizationScope:ClientAuthorizationScope", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientAuthorizationScope resources.
type clientAuthorizationScopeState struct {
	DisplayName      *string `pulumi:"displayName"`
	IconUri          *string `pulumi:"iconUri"`
	Name             *string `pulumi:"name"`
	RealmId          *string `pulumi:"realmId"`
	ResourceServerId *string `pulumi:"resourceServerId"`
}

type ClientAuthorizationScopeState struct {
	DisplayName      pulumi.StringPtrInput
	IconUri          pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
}

func (ClientAuthorizationScopeState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationScopeState)(nil)).Elem()
}

type clientAuthorizationScopeArgs struct {
	DisplayName      *string `pulumi:"displayName"`
	IconUri          *string `pulumi:"iconUri"`
	Name             *string `pulumi:"name"`
	RealmId          string  `pulumi:"realmId"`
	ResourceServerId string  `pulumi:"resourceServerId"`
}

// The set of arguments for constructing a ClientAuthorizationScope resource.
type ClientAuthorizationScopeArgs struct {
	DisplayName      pulumi.StringPtrInput
	IconUri          pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
}

func (ClientAuthorizationScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationScopeArgs)(nil)).Elem()
}
