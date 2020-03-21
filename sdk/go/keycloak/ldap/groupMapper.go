// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ldap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # ldap.GroupMapper
//
// Allows for creating and managing group mappers for Keycloak users federated
// via LDAP.
//
// The LDAP group mapper can be used to map an LDAP user's groups from some DN
// to Keycloak groups. This group mapper will also create the groups within Keycloak
// if they do not already exist.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm that this LDAP mapper will exist in.
// - `ldapUserFederationId` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
// - `name` - (Required) Display name of this mapper when displayed in the console.
// - `ldapGroupsDn` - (Required) The LDAP DN where groups can be found.
// - `groupNameLdapAttribute` - (Required) The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
// - `groupObjectClasses` - (Required) Array of strings representing the object classes for the group. Must contain at least one.
// - `preserveGroupInheritance` - (Optional) When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
// - `ignoreMissingGroups` - (Optional) When `true`, missing groups in the hierarchy will be ignored.
// - `membershipLdapAttribute` - (Required) The name of the LDAP attribute that is used for membership mappings.
// - `membershipAttributeType` - (Optional) Can be one of `DN` or `UID`. Defaults to `DN`.
// - `membershipUserLdapAttribute` - (Required) The name of the LDAP attribute on a user that is used for membership mappings.
// - `groupsLdapFilter` - (Optional) When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
// - `mode` - (Optional) Can be one of `READ_ONLY` or `LDAP_ONLY`. Defaults to `READ_ONLY`.
// - `userRolesRetrieveStrategy` - (Optional) Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
// - `memberofLdapAttribute` - (Optional) Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
// - `mappedGroupAttributes` - (Optional) Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
// - `dropNonExistingGroupsDuringSync` - (Optional) When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_ldap_group_mapper.html.markdown.
type GroupMapper struct {
	pulumi.CustomResourceState

	DropNonExistingGroupsDuringSync pulumi.BoolPtrOutput `pulumi:"dropNonExistingGroupsDuringSync"`
	GroupNameLdapAttribute pulumi.StringOutput `pulumi:"groupNameLdapAttribute"`
	GroupObjectClasses pulumi.StringArrayOutput `pulumi:"groupObjectClasses"`
	GroupsLdapFilter pulumi.StringPtrOutput `pulumi:"groupsLdapFilter"`
	IgnoreMissingGroups pulumi.BoolPtrOutput `pulumi:"ignoreMissingGroups"`
	LdapGroupsDn pulumi.StringOutput `pulumi:"ldapGroupsDn"`
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	MappedGroupAttributes pulumi.StringArrayOutput `pulumi:"mappedGroupAttributes"`
	MemberofLdapAttribute pulumi.StringPtrOutput `pulumi:"memberofLdapAttribute"`
	MembershipAttributeType pulumi.StringPtrOutput `pulumi:"membershipAttributeType"`
	MembershipLdapAttribute pulumi.StringOutput `pulumi:"membershipLdapAttribute"`
	MembershipUserLdapAttribute pulumi.StringOutput `pulumi:"membershipUserLdapAttribute"`
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	PreserveGroupInheritance pulumi.BoolPtrOutput `pulumi:"preserveGroupInheritance"`
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	UserRolesRetrieveStrategy pulumi.StringPtrOutput `pulumi:"userRolesRetrieveStrategy"`
}

// NewGroupMapper registers a new resource with the given unique name, arguments, and options.
func NewGroupMapper(ctx *pulumi.Context,
	name string, args *GroupMapperArgs, opts ...pulumi.ResourceOption) (*GroupMapper, error) {
	if args == nil || args.GroupNameLdapAttribute == nil {
		return nil, errors.New("missing required argument 'GroupNameLdapAttribute'")
	}
	if args == nil || args.GroupObjectClasses == nil {
		return nil, errors.New("missing required argument 'GroupObjectClasses'")
	}
	if args == nil || args.LdapGroupsDn == nil {
		return nil, errors.New("missing required argument 'LdapGroupsDn'")
	}
	if args == nil || args.LdapUserFederationId == nil {
		return nil, errors.New("missing required argument 'LdapUserFederationId'")
	}
	if args == nil || args.MembershipLdapAttribute == nil {
		return nil, errors.New("missing required argument 'MembershipLdapAttribute'")
	}
	if args == nil || args.MembershipUserLdapAttribute == nil {
		return nil, errors.New("missing required argument 'MembershipUserLdapAttribute'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &GroupMapperArgs{}
	}
	var resource GroupMapper
	err := ctx.RegisterResource("keycloak:ldap/groupMapper:GroupMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupMapper gets an existing GroupMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupMapperState, opts ...pulumi.ResourceOption) (*GroupMapper, error) {
	var resource GroupMapper
	err := ctx.ReadResource("keycloak:ldap/groupMapper:GroupMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupMapper resources.
type groupMapperState struct {
	DropNonExistingGroupsDuringSync *bool `pulumi:"dropNonExistingGroupsDuringSync"`
	GroupNameLdapAttribute *string `pulumi:"groupNameLdapAttribute"`
	GroupObjectClasses []string `pulumi:"groupObjectClasses"`
	GroupsLdapFilter *string `pulumi:"groupsLdapFilter"`
	IgnoreMissingGroups *bool `pulumi:"ignoreMissingGroups"`
	LdapGroupsDn *string `pulumi:"ldapGroupsDn"`
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	MappedGroupAttributes []string `pulumi:"mappedGroupAttributes"`
	MemberofLdapAttribute *string `pulumi:"memberofLdapAttribute"`
	MembershipAttributeType *string `pulumi:"membershipAttributeType"`
	MembershipLdapAttribute *string `pulumi:"membershipLdapAttribute"`
	MembershipUserLdapAttribute *string `pulumi:"membershipUserLdapAttribute"`
	Mode *string `pulumi:"mode"`
	// Display name of the mapper when displayed in the console.
	Name *string `pulumi:"name"`
	PreserveGroupInheritance *bool `pulumi:"preserveGroupInheritance"`
	// The realm in which the ldap user federation provider exists.
	RealmId *string `pulumi:"realmId"`
	UserRolesRetrieveStrategy *string `pulumi:"userRolesRetrieveStrategy"`
}

type GroupMapperState struct {
	DropNonExistingGroupsDuringSync pulumi.BoolPtrInput
	GroupNameLdapAttribute pulumi.StringPtrInput
	GroupObjectClasses pulumi.StringArrayInput
	GroupsLdapFilter pulumi.StringPtrInput
	IgnoreMissingGroups pulumi.BoolPtrInput
	LdapGroupsDn pulumi.StringPtrInput
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	MappedGroupAttributes pulumi.StringArrayInput
	MemberofLdapAttribute pulumi.StringPtrInput
	MembershipAttributeType pulumi.StringPtrInput
	MembershipLdapAttribute pulumi.StringPtrInput
	MembershipUserLdapAttribute pulumi.StringPtrInput
	Mode pulumi.StringPtrInput
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringPtrInput
	PreserveGroupInheritance pulumi.BoolPtrInput
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringPtrInput
	UserRolesRetrieveStrategy pulumi.StringPtrInput
}

func (GroupMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMapperState)(nil)).Elem()
}

type groupMapperArgs struct {
	DropNonExistingGroupsDuringSync *bool `pulumi:"dropNonExistingGroupsDuringSync"`
	GroupNameLdapAttribute string `pulumi:"groupNameLdapAttribute"`
	GroupObjectClasses []string `pulumi:"groupObjectClasses"`
	GroupsLdapFilter *string `pulumi:"groupsLdapFilter"`
	IgnoreMissingGroups *bool `pulumi:"ignoreMissingGroups"`
	LdapGroupsDn string `pulumi:"ldapGroupsDn"`
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	MappedGroupAttributes []string `pulumi:"mappedGroupAttributes"`
	MemberofLdapAttribute *string `pulumi:"memberofLdapAttribute"`
	MembershipAttributeType *string `pulumi:"membershipAttributeType"`
	MembershipLdapAttribute string `pulumi:"membershipLdapAttribute"`
	MembershipUserLdapAttribute string `pulumi:"membershipUserLdapAttribute"`
	Mode *string `pulumi:"mode"`
	// Display name of the mapper when displayed in the console.
	Name *string `pulumi:"name"`
	PreserveGroupInheritance *bool `pulumi:"preserveGroupInheritance"`
	// The realm in which the ldap user federation provider exists.
	RealmId string `pulumi:"realmId"`
	UserRolesRetrieveStrategy *string `pulumi:"userRolesRetrieveStrategy"`
}

// The set of arguments for constructing a GroupMapper resource.
type GroupMapperArgs struct {
	DropNonExistingGroupsDuringSync pulumi.BoolPtrInput
	GroupNameLdapAttribute pulumi.StringInput
	GroupObjectClasses pulumi.StringArrayInput
	GroupsLdapFilter pulumi.StringPtrInput
	IgnoreMissingGroups pulumi.BoolPtrInput
	LdapGroupsDn pulumi.StringInput
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	MappedGroupAttributes pulumi.StringArrayInput
	MemberofLdapAttribute pulumi.StringPtrInput
	MembershipAttributeType pulumi.StringPtrInput
	MembershipLdapAttribute pulumi.StringInput
	MembershipUserLdapAttribute pulumi.StringInput
	Mode pulumi.StringPtrInput
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringPtrInput
	PreserveGroupInheritance pulumi.BoolPtrInput
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringInput
	UserRolesRetrieveStrategy pulumi.StringPtrInput
}

func (GroupMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupMapperArgs)(nil)).Elem()
}

