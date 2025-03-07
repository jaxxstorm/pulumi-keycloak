// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientDefaultScopes struct {
	pulumi.CustomResourceState

	ClientId      pulumi.StringOutput      `pulumi:"clientId"`
	DefaultScopes pulumi.StringArrayOutput `pulumi:"defaultScopes"`
	RealmId       pulumi.StringOutput      `pulumi:"realmId"`
}

// NewClientDefaultScopes registers a new resource with the given unique name, arguments, and options.
func NewClientDefaultScopes(ctx *pulumi.Context,
	name string, args *ClientDefaultScopesArgs, opts ...pulumi.ResourceOption) (*ClientDefaultScopes, error) {
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.DefaultScopes == nil {
		return nil, errors.New("missing required argument 'DefaultScopes'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &ClientDefaultScopesArgs{}
	}
	var resource ClientDefaultScopes
	err := ctx.RegisterResource("keycloak:openid/clientDefaultScopes:ClientDefaultScopes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientDefaultScopes gets an existing ClientDefaultScopes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientDefaultScopes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientDefaultScopesState, opts ...pulumi.ResourceOption) (*ClientDefaultScopes, error) {
	var resource ClientDefaultScopes
	err := ctx.ReadResource("keycloak:openid/clientDefaultScopes:ClientDefaultScopes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientDefaultScopes resources.
type clientDefaultScopesState struct {
	ClientId      *string  `pulumi:"clientId"`
	DefaultScopes []string `pulumi:"defaultScopes"`
	RealmId       *string  `pulumi:"realmId"`
}

type ClientDefaultScopesState struct {
	ClientId      pulumi.StringPtrInput
	DefaultScopes pulumi.StringArrayInput
	RealmId       pulumi.StringPtrInput
}

func (ClientDefaultScopesState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientDefaultScopesState)(nil)).Elem()
}

type clientDefaultScopesArgs struct {
	ClientId      string   `pulumi:"clientId"`
	DefaultScopes []string `pulumi:"defaultScopes"`
	RealmId       string   `pulumi:"realmId"`
}

// The set of arguments for constructing a ClientDefaultScopes resource.
type ClientDefaultScopesArgs struct {
	ClientId      pulumi.StringInput
	DefaultScopes pulumi.StringArrayInput
	RealmId       pulumi.StringInput
}

func (ClientDefaultScopesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientDefaultScopesArgs)(nil)).Elem()
}
