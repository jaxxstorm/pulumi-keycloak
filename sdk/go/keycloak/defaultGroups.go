// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # .DefaultGroups
//
// Allows for managing a realm's default groups.
//
// Note that you should not use `.DefaultGroups` with a group with memberships managed
// by `.GroupMemberships`.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm this group exists in.
// - `groupIds` - (Required) A set of group ids that should be default groups on the realm referenced by `realmId`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_default_groups.html.markdown.
type DefaultGroups struct {
	pulumi.CustomResourceState

	GroupIds pulumi.StringArrayOutput `pulumi:"groupIds"`
	RealmId  pulumi.StringOutput      `pulumi:"realmId"`
}

// NewDefaultGroups registers a new resource with the given unique name, arguments, and options.
func NewDefaultGroups(ctx *pulumi.Context,
	name string, args *DefaultGroupsArgs, opts ...pulumi.ResourceOption) (*DefaultGroups, error) {
	if args == nil || args.GroupIds == nil {
		return nil, errors.New("missing required argument 'GroupIds'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &DefaultGroupsArgs{}
	}
	var resource DefaultGroups
	err := ctx.RegisterResource("keycloak:index/defaultGroups:DefaultGroups", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultGroups gets an existing DefaultGroups resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultGroups(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultGroupsState, opts ...pulumi.ResourceOption) (*DefaultGroups, error) {
	var resource DefaultGroups
	err := ctx.ReadResource("keycloak:index/defaultGroups:DefaultGroups", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultGroups resources.
type defaultGroupsState struct {
	GroupIds []string `pulumi:"groupIds"`
	RealmId  *string  `pulumi:"realmId"`
}

type DefaultGroupsState struct {
	GroupIds pulumi.StringArrayInput
	RealmId  pulumi.StringPtrInput
}

func (DefaultGroupsState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultGroupsState)(nil)).Elem()
}

type defaultGroupsArgs struct {
	GroupIds []string `pulumi:"groupIds"`
	RealmId  string   `pulumi:"realmId"`
}

// The set of arguments for constructing a DefaultGroups resource.
type DefaultGroupsArgs struct {
	GroupIds pulumi.StringArrayInput
	RealmId  pulumi.StringInput
}

func (DefaultGroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultGroupsArgs)(nil)).Elem()
}
