// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientGroupPolicy struct {
	pulumi.CustomResourceState

	DecisionStrategy pulumi.StringOutput               `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput            `pulumi:"description"`
	Groups           ClientGroupPolicyGroupArrayOutput `pulumi:"groups"`
	GroupsClaim      pulumi.StringPtrOutput            `pulumi:"groupsClaim"`
	Logic            pulumi.StringPtrOutput            `pulumi:"logic"`
	Name             pulumi.StringOutput               `pulumi:"name"`
	RealmId          pulumi.StringOutput               `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput               `pulumi:"resourceServerId"`
}

// NewClientGroupPolicy registers a new resource with the given unique name, arguments, and options.
func NewClientGroupPolicy(ctx *pulumi.Context,
	name string, args *ClientGroupPolicyArgs, opts ...pulumi.ResourceOption) (*ClientGroupPolicy, error) {
	if args == nil || args.DecisionStrategy == nil {
		return nil, errors.New("missing required argument 'DecisionStrategy'")
	}
	if args == nil || args.Groups == nil {
		return nil, errors.New("missing required argument 'Groups'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.ResourceServerId == nil {
		return nil, errors.New("missing required argument 'ResourceServerId'")
	}
	if args == nil {
		args = &ClientGroupPolicyArgs{}
	}
	var resource ClientGroupPolicy
	err := ctx.RegisterResource("keycloak:openid/clientGroupPolicy:ClientGroupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientGroupPolicy gets an existing ClientGroupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientGroupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientGroupPolicyState, opts ...pulumi.ResourceOption) (*ClientGroupPolicy, error) {
	var resource ClientGroupPolicy
	err := ctx.ReadResource("keycloak:openid/clientGroupPolicy:ClientGroupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientGroupPolicy resources.
type clientGroupPolicyState struct {
	DecisionStrategy *string                  `pulumi:"decisionStrategy"`
	Description      *string                  `pulumi:"description"`
	Groups           []ClientGroupPolicyGroup `pulumi:"groups"`
	GroupsClaim      *string                  `pulumi:"groupsClaim"`
	Logic            *string                  `pulumi:"logic"`
	Name             *string                  `pulumi:"name"`
	RealmId          *string                  `pulumi:"realmId"`
	ResourceServerId *string                  `pulumi:"resourceServerId"`
}

type ClientGroupPolicyState struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Groups           ClientGroupPolicyGroupArrayInput
	GroupsClaim      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
}

func (ClientGroupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientGroupPolicyState)(nil)).Elem()
}

type clientGroupPolicyArgs struct {
	DecisionStrategy string                   `pulumi:"decisionStrategy"`
	Description      *string                  `pulumi:"description"`
	Groups           []ClientGroupPolicyGroup `pulumi:"groups"`
	GroupsClaim      *string                  `pulumi:"groupsClaim"`
	Logic            *string                  `pulumi:"logic"`
	Name             *string                  `pulumi:"name"`
	RealmId          string                   `pulumi:"realmId"`
	ResourceServerId string                   `pulumi:"resourceServerId"`
}

// The set of arguments for constructing a ClientGroupPolicy resource.
type ClientGroupPolicyArgs struct {
	DecisionStrategy pulumi.StringInput
	Description      pulumi.StringPtrInput
	Groups           ClientGroupPolicyGroupArrayInput
	GroupsClaim      pulumi.StringPtrInput
	Logic            pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
}

func (ClientGroupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientGroupPolicyArgs)(nil)).Elem()
}
