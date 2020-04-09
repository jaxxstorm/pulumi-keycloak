// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type GroupMemberships struct {
	pulumi.CustomResourceState

	GroupId pulumi.StringPtrOutput   `pulumi:"groupId"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	RealmId pulumi.StringOutput      `pulumi:"realmId"`
}

// NewGroupMemberships registers a new resource with the given unique name, arguments, and options.
func NewGroupMemberships(ctx *pulumi.Context,
	name string, args *GroupMembershipsArgs, opts ...pulumi.ResourceOption) (*GroupMemberships, error) {
	if args == nil || args.Members == nil {
		return nil, errors.New("missing required argument 'Members'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &GroupMembershipsArgs{}
	}
	var resource GroupMemberships
	err := ctx.RegisterResource("keycloak:index/groupMemberships:GroupMemberships", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMemberships gets an existing GroupMemberships resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMemberships(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMembershipsState, opts ...pulumi.ResourceOption) (*GroupMemberships, error) {
	var resource GroupMemberships
	err := ctx.ReadResource("keycloak:index/groupMemberships:GroupMemberships", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMemberships resources.
type groupMembershipsState struct {
	GroupId *string  `pulumi:"groupId"`
	Members []string `pulumi:"members"`
	RealmId *string  `pulumi:"realmId"`
}

type GroupMembershipsState struct {
	GroupId pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	RealmId pulumi.StringPtrInput
}

func (GroupMembershipsState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipsState)(nil)).Elem()
}

type groupMembershipsArgs struct {
	GroupId *string  `pulumi:"groupId"`
	Members []string `pulumi:"members"`
	RealmId string   `pulumi:"realmId"`
}

// The set of arguments for constructing a GroupMemberships resource.
type GroupMembershipsArgs struct {
	GroupId pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	RealmId pulumi.StringInput
}

func (GroupMembershipsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMembershipsArgs)(nil)).Elem()
}
