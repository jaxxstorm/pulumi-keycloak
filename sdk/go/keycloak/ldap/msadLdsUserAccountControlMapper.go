// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type MsadLdsUserAccountControlMapper struct {
	pulumi.CustomResourceState

	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewMsadLdsUserAccountControlMapper registers a new resource with the given unique name, arguments, and options.
func NewMsadLdsUserAccountControlMapper(ctx *pulumi.Context,
	name string, args *MsadLdsUserAccountControlMapperArgs, opts ...pulumi.ResourceOption) (*MsadLdsUserAccountControlMapper, error) {
	if args == nil || args.LdapUserFederationId == nil {
		return nil, errors.New("missing required argument 'LdapUserFederationId'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &MsadLdsUserAccountControlMapperArgs{}
	}
	var resource MsadLdsUserAccountControlMapper
	err := ctx.RegisterResource("keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMsadLdsUserAccountControlMapper gets an existing MsadLdsUserAccountControlMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMsadLdsUserAccountControlMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MsadLdsUserAccountControlMapperState, opts ...pulumi.ResourceOption) (*MsadLdsUserAccountControlMapper, error) {
	var resource MsadLdsUserAccountControlMapper
	err := ctx.ReadResource("keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MsadLdsUserAccountControlMapper resources.
type msadLdsUserAccountControlMapperState struct {
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Display name of the mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm in which the ldap user federation provider exists.
	RealmId *string `pulumi:"realmId"`
}

type MsadLdsUserAccountControlMapperState struct {
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringPtrInput
}

func (MsadLdsUserAccountControlMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*msadLdsUserAccountControlMapperState)(nil)).Elem()
}

type msadLdsUserAccountControlMapperArgs struct {
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Display name of the mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm in which the ldap user federation provider exists.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a MsadLdsUserAccountControlMapper resource.
type MsadLdsUserAccountControlMapperArgs struct {
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringInput
}

func (MsadLdsUserAccountControlMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*msadLdsUserAccountControlMapperArgs)(nil)).Elem()
}
