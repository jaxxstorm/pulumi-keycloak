// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package ldap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # ldap.MsadUserAccountControlMapper
//
// Allows for creating and managing MSAD user account control mappers for Keycloak
// users federated via LDAP.
//
// The MSAD (Microsoft Active Directory) user account control mapper is specific
// to LDAP user federation providers that are pulling from AD, and it can propagate
// AD user state to Keycloak in order to enforce settings like expired passwords
// or disabled accounts.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm that this LDAP mapper will exist in.
// - `ldapUserFederationId` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
// - `name` - (Required) Display name of this mapper when displayed in the console.
// - `ldapPasswordPolicyHintsEnabled` - (Optional) When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_ldap_msad_user_account_control_mapper.html.markdown.
type MsadUserAccountControlMapper struct {
	pulumi.CustomResourceState

	LdapPasswordPolicyHintsEnabled pulumi.BoolPtrOutput `pulumi:"ldapPasswordPolicyHintsEnabled"`
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringOutput `pulumi:"ldapUserFederationId"`
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
}

// NewMsadUserAccountControlMapper registers a new resource with the given unique name, arguments, and options.
func NewMsadUserAccountControlMapper(ctx *pulumi.Context,
	name string, args *MsadUserAccountControlMapperArgs, opts ...pulumi.ResourceOption) (*MsadUserAccountControlMapper, error) {
	if args == nil || args.LdapUserFederationId == nil {
		return nil, errors.New("missing required argument 'LdapUserFederationId'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &MsadUserAccountControlMapperArgs{}
	}
	var resource MsadUserAccountControlMapper
	err := ctx.RegisterResource("keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMsadUserAccountControlMapper gets an existing MsadUserAccountControlMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMsadUserAccountControlMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MsadUserAccountControlMapperState, opts ...pulumi.ResourceOption) (*MsadUserAccountControlMapper, error) {
	var resource MsadUserAccountControlMapper
	err := ctx.ReadResource("keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MsadUserAccountControlMapper resources.
type msadUserAccountControlMapperState struct {
	LdapPasswordPolicyHintsEnabled *bool `pulumi:"ldapPasswordPolicyHintsEnabled"`
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId *string `pulumi:"ldapUserFederationId"`
	// Display name of the mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm in which the ldap user federation provider exists.
	RealmId *string `pulumi:"realmId"`
}

type MsadUserAccountControlMapperState struct {
	LdapPasswordPolicyHintsEnabled pulumi.BoolPtrInput
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringPtrInput
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringPtrInput
}

func (MsadUserAccountControlMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*msadUserAccountControlMapperState)(nil)).Elem()
}

type msadUserAccountControlMapperArgs struct {
	LdapPasswordPolicyHintsEnabled *bool `pulumi:"ldapPasswordPolicyHintsEnabled"`
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId string `pulumi:"ldapUserFederationId"`
	// Display name of the mapper when displayed in the console.
	Name *string `pulumi:"name"`
	// The realm in which the ldap user federation provider exists.
	RealmId string `pulumi:"realmId"`
}

// The set of arguments for constructing a MsadUserAccountControlMapper resource.
type MsadUserAccountControlMapperArgs struct {
	LdapPasswordPolicyHintsEnabled pulumi.BoolPtrInput
	// The ldap user federation provider to attach this mapper to.
	LdapUserFederationId pulumi.StringInput
	// Display name of the mapper when displayed in the console.
	Name pulumi.StringPtrInput
	// The realm in which the ldap user federation provider exists.
	RealmId pulumi.StringInput
}

func (MsadUserAccountControlMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*msadUserAccountControlMapperArgs)(nil)).Elem()
}
