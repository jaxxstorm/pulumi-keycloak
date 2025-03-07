// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientAuthorizationPermission struct {
	pulumi.CustomResourceState

	DecisionStrategy pulumi.StringPtrOutput   `pulumi:"decisionStrategy"`
	Description      pulumi.StringPtrOutput   `pulumi:"description"`
	Name             pulumi.StringOutput      `pulumi:"name"`
	Policies         pulumi.StringArrayOutput `pulumi:"policies"`
	RealmId          pulumi.StringOutput      `pulumi:"realmId"`
	ResourceServerId pulumi.StringOutput      `pulumi:"resourceServerId"`
	Resources        pulumi.StringArrayOutput `pulumi:"resources"`
	Scopes           pulumi.StringArrayOutput `pulumi:"scopes"`
	Type             pulumi.StringPtrOutput   `pulumi:"type"`
}

// NewClientAuthorizationPermission registers a new resource with the given unique name, arguments, and options.
func NewClientAuthorizationPermission(ctx *pulumi.Context,
	name string, args *ClientAuthorizationPermissionArgs, opts ...pulumi.ResourceOption) (*ClientAuthorizationPermission, error) {
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.ResourceServerId == nil {
		return nil, errors.New("missing required argument 'ResourceServerId'")
	}
	if args == nil {
		args = &ClientAuthorizationPermissionArgs{}
	}
	var resource ClientAuthorizationPermission
	err := ctx.RegisterResource("keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClientAuthorizationPermission gets an existing ClientAuthorizationPermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClientAuthorizationPermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientAuthorizationPermissionState, opts ...pulumi.ResourceOption) (*ClientAuthorizationPermission, error) {
	var resource ClientAuthorizationPermission
	err := ctx.ReadResource("keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClientAuthorizationPermission resources.
type clientAuthorizationPermissionState struct {
	DecisionStrategy *string  `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Name             *string  `pulumi:"name"`
	Policies         []string `pulumi:"policies"`
	RealmId          *string  `pulumi:"realmId"`
	ResourceServerId *string  `pulumi:"resourceServerId"`
	Resources        []string `pulumi:"resources"`
	Scopes           []string `pulumi:"scopes"`
	Type             *string  `pulumi:"type"`
}

type ClientAuthorizationPermissionState struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Policies         pulumi.StringArrayInput
	RealmId          pulumi.StringPtrInput
	ResourceServerId pulumi.StringPtrInput
	Resources        pulumi.StringArrayInput
	Scopes           pulumi.StringArrayInput
	Type             pulumi.StringPtrInput
}

func (ClientAuthorizationPermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationPermissionState)(nil)).Elem()
}

type clientAuthorizationPermissionArgs struct {
	DecisionStrategy *string  `pulumi:"decisionStrategy"`
	Description      *string  `pulumi:"description"`
	Name             *string  `pulumi:"name"`
	Policies         []string `pulumi:"policies"`
	RealmId          string   `pulumi:"realmId"`
	ResourceServerId string   `pulumi:"resourceServerId"`
	Resources        []string `pulumi:"resources"`
	Scopes           []string `pulumi:"scopes"`
	Type             *string  `pulumi:"type"`
}

// The set of arguments for constructing a ClientAuthorizationPermission resource.
type ClientAuthorizationPermissionArgs struct {
	DecisionStrategy pulumi.StringPtrInput
	Description      pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Policies         pulumi.StringArrayInput
	RealmId          pulumi.StringInput
	ResourceServerId pulumi.StringInput
	Resources        pulumi.StringArrayInput
	Scopes           pulumi.StringArrayInput
	Type             pulumi.StringPtrInput
}

func (ClientAuthorizationPermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientAuthorizationPermissionArgs)(nil)).Elem()
}
