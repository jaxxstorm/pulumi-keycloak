// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type UserTemplateImporterIdentityProviderMapper struct {
	pulumi.CustomResourceState

	// IDP Alias
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Realm Name
	Realm pulumi.StringOutput `pulumi:"realm"`
	// Username For Template Import
	Template pulumi.StringPtrOutput `pulumi:"template"`
}

// NewUserTemplateImporterIdentityProviderMapper registers a new resource with the given unique name, arguments, and options.
func NewUserTemplateImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, args *UserTemplateImporterIdentityProviderMapperArgs, opts ...pulumi.ResourceOption) (*UserTemplateImporterIdentityProviderMapper, error) {
	if args == nil || args.IdentityProviderAlias == nil {
		return nil, errors.New("missing required argument 'IdentityProviderAlias'")
	}
	if args == nil || args.Realm == nil {
		return nil, errors.New("missing required argument 'Realm'")
	}
	if args == nil {
		args = &UserTemplateImporterIdentityProviderMapperArgs{}
	}
	var resource UserTemplateImporterIdentityProviderMapper
	err := ctx.RegisterResource("keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserTemplateImporterIdentityProviderMapper gets an existing UserTemplateImporterIdentityProviderMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserTemplateImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserTemplateImporterIdentityProviderMapperState, opts ...pulumi.ResourceOption) (*UserTemplateImporterIdentityProviderMapper, error) {
	var resource UserTemplateImporterIdentityProviderMapper
	err := ctx.ReadResource("keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserTemplateImporterIdentityProviderMapper resources.
type userTemplateImporterIdentityProviderMapperState struct {
	// IDP Alias
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm *string `pulumi:"realm"`
	// Username For Template Import
	Template *string `pulumi:"template"`
}

type UserTemplateImporterIdentityProviderMapperState struct {
	// IDP Alias
	IdentityProviderAlias pulumi.StringPtrInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringPtrInput
	// Username For Template Import
	Template pulumi.StringPtrInput
}

func (UserTemplateImporterIdentityProviderMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*userTemplateImporterIdentityProviderMapperState)(nil)).Elem()
}

type userTemplateImporterIdentityProviderMapperArgs struct {
	// IDP Alias
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm string `pulumi:"realm"`
	// Username For Template Import
	Template *string `pulumi:"template"`
}

// The set of arguments for constructing a UserTemplateImporterIdentityProviderMapper resource.
type UserTemplateImporterIdentityProviderMapperArgs struct {
	// IDP Alias
	IdentityProviderAlias pulumi.StringInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringInput
	// Username For Template Import
	Template pulumi.StringPtrInput
}

func (UserTemplateImporterIdentityProviderMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userTemplateImporterIdentityProviderMapperArgs)(nil)).Elem()
}
