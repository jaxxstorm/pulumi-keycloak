// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package authentication

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Flow struct {
	pulumi.CustomResourceState

	Alias       pulumi.StringOutput    `pulumi:"alias"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	ProviderId  pulumi.StringPtrOutput `pulumi:"providerId"`
	RealmId     pulumi.StringOutput    `pulumi:"realmId"`
}

// NewFlow registers a new resource with the given unique name, arguments, and options.
func NewFlow(ctx *pulumi.Context,
	name string, args *FlowArgs, opts ...pulumi.ResourceOption) (*Flow, error) {
	if args == nil || args.Alias == nil {
		return nil, errors.New("missing required argument 'Alias'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &FlowArgs{}
	}
	var resource Flow
	err := ctx.RegisterResource("keycloak:authentication/flow:Flow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlow gets an existing Flow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlowState, opts ...pulumi.ResourceOption) (*Flow, error) {
	var resource Flow
	err := ctx.ReadResource("keycloak:authentication/flow:Flow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Flow resources.
type flowState struct {
	Alias       *string `pulumi:"alias"`
	Description *string `pulumi:"description"`
	ProviderId  *string `pulumi:"providerId"`
	RealmId     *string `pulumi:"realmId"`
}

type FlowState struct {
	Alias       pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	ProviderId  pulumi.StringPtrInput
	RealmId     pulumi.StringPtrInput
}

func (FlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*flowState)(nil)).Elem()
}

type flowArgs struct {
	Alias       string  `pulumi:"alias"`
	Description *string `pulumi:"description"`
	ProviderId  *string `pulumi:"providerId"`
	RealmId     string  `pulumi:"realmId"`
}

// The set of arguments for constructing a Flow resource.
type FlowArgs struct {
	Alias       pulumi.StringInput
	Description pulumi.StringPtrInput
	ProviderId  pulumi.StringPtrInput
	RealmId     pulumi.StringInput
}

func (FlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flowArgs)(nil)).Elem()
}
