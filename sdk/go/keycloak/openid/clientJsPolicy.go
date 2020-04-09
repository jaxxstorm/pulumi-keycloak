// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientJsPolicy struct {
	pulumi.CustomResourceState

	Code             pulumi.StringOutput    `pulumi:"code"`
	DecisionStrategy pulumi.StringOutput    `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	Logic            pulumi.StringPtrOutput `pulumi:"logic"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	RealmId          pulumi.StringOutput    `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput    `pulumi:"resourceServerId"`
	Type             pulumi.StringPtrOutput `pulumi:"type"`
}

// NewClientJsPolicy registers a new resource with the given unique name, arguments, and options.
func NewClientJsPolicy(ctx *pulumi.Context,
	name string, args *ClientJsPolicyArgs, opts ...pulumi.ResourceOption) (*ClientJsPolicy, error) {
	if args == nil || args.Code == nil {
		return nil, errors.New("missing required argument 'Code'")
	}
	if args == nil || args.DecisionStrategy == nil {
		return nil, errors.New("missing required argument 'DecisionStrategy'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.ResourceServerId == nil {
		return nil, errors.New("missing required argument 'ResourceServerId'")
	}
	if args == nil {
		args = &ClientJsPolicyArgs{}
	}
	var resource ClientJsPolicy
	err := ctx.RegisterResource("keycloak:openid/clientJsPolicy:ClientJsPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientJsPolicy gets an existing ClientJsPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientJsPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientJsPolicyState, opts ...pulumi.ResourceOption) (*ClientJsPolicy, error) {
	var resource ClientJsPolicy
	err := ctx.ReadResource("keycloak:openid/clientJsPolicy:ClientJsPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientJsPolicy resources.
type clientJsPolicyState struct {
	Code             *string `pulumi:"code"`
	DecisionStrategy *string `pulumi:"decisionStrategy"`
	Description      *string `pulumi:"description"`
	Logic            *string `pulumi:"logic"`
	Name             *string `pulumi:"name"`
	RealmId          *string `pulumi:"realmId"`
	ResourceServerId *string `pulumi:"resourceServerId"`
	Type             *string `pulumi:"type"`
}

type ClientJsPolicyState struct {
	Code             pulumi.StringPtrInput
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
	Type             pulumi.StringPtrInput
}

func (ClientJsPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientJsPolicyState)(nil)).Elem()
}

type clientJsPolicyArgs struct {
	Code             string  `pulumi:"code"`
	DecisionStrategy string  `pulumi:"decisionStrategy"`
	Description      *string `pulumi:"description"`
	Logic            *string `pulumi:"logic"`
	Name             *string `pulumi:"name"`
	RealmId          string  `pulumi:"realmId"`
	ResourceServerId string  `pulumi:"resourceServerId"`
	Type             *string `pulumi:"type"`
}

// The set of arguments for constructing a ClientJsPolicy resource.
type ClientJsPolicyArgs struct {
	Code             pulumi.StringInput
	DecisionStrategy pulumi.StringInput
	Description      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
	Type             pulumi.StringPtrInput
}

func (ClientJsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientJsPolicyArgs)(nil)).Elem()
}
