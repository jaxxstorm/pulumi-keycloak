// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # ldap.UserFederation
//
// Allows for creating and managing LDAP user federation providers within Keycloak.
//
// Keycloak can use an LDAP user federation provider to federate users to Keycloak
// from a directory system such as LDAP or Active Directory. Federated users
// will exist within the realm and will be able to log in to clients. Federated
// users can have their attributes defined using mappers.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm that this provider will provide user federation for.
// - `name` - (Required) Display name of the provider when displayed in the console.
// - `enabled` - (Optional) When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
// - `priority` - (Optional) Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
// - `importEnabled` - (Optional) When `true`, LDAP users will be imported into the Keycloak database. Defaults to `true`.
// - `editMode` - (Optional) Can be one of `READ_ONLY`, `WRITABLE`, or `UNSYNCED`. `UNSYNCED` allows user data to be imported but not synced back to LDAP. Defaults to `READ_ONLY`.
// - `syncRegistrations` - (Optional) When `true`, newly created users will be synced back to LDAP. Defaults to `false`.
// - `vendor` - (Optional) Can be one of `OTHER`, `EDIRECTORY`, `AD`, `RHDS`, or `TIVOLI`. When this is selected in the GUI, it provides reasonable defaults for other fields. When used with the Keycloak API, this attribute does nothing, but is still required. Defaults to `OPTIONAL`.
// - `usernameLdapAttribute` - (Required) Name of the LDAP attribute to use as the Keycloak username.
// - `rdnLdapAttribute` - (Required) Name of the LDAP attribute to use as the relative distinguished name.
// - `uuidLdapAttribute` - (Required) Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
// - `userObjectClasses` - (Required) Array of all values of LDAP objectClass attribute for users in LDAP. Must contain at least one.
// - `connectionUrl` - (Required) Connection URL to the LDAP server.
// - `usersDn` - (Required) Full DN of LDAP tree where your users are.
// - `bindDn` - (Optional) DN of LDAP admin, which will be used by Keycloak to access LDAP server. This attribute must be set if `bindCredential` is set.
// - `bindCredential` - (Optional) Password of LDAP admin. This attribute must be set if `bindDn` is set.
// - `customUserSearchFilter` - (Optional) Additional LDAP filter for filtering searched users. Must begin with `(` and end with `)`.
// - `searchScope` - (Optional) Can be one of `ONE_LEVEL` or `SUBTREE`:
//     - `ONE_LEVEL`: Only search for users in the DN specified by `userDn`.
//     - `SUBTREE`: Search entire LDAP subtree.
// - `validatePasswordPolicy` - (Optional) When `true`, Keycloak will validate passwords using the realm policy before updating it.
// - `useTruststoreSpi` - (Optional) Can be one of `ALWAYS`, `ONLY_FOR_LDAPS`, or `NEVER`:
//     - `ALWAYS` - Always use the truststore SPI for LDAP connections.
//     - `NEVER` - Never use the truststore SPI for LDAP connections.
//     - `ONLY_FOR_LDAPS` - Only use the truststore SPI if your LDAP connection uses the ldaps protocol.
// - `connectionTimeout` - (Optional) LDAP connection timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
// - `readTimeout` - (Optional) LDAP read timeout in the format of a [Go duration string](https://golang.org/pkg/time/#Duration.String).
// - `pagination` - (Optional) When true, Keycloak assumes the LDAP server supports pagination. Defaults to `true`.
// - `batchSizeForSync` - (Optional) The number of users to sync within a single transaction. Defaults to `1000`.
// - `fullSyncPeriod` - (Optional) How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
// - `changedSyncPeriod` - (Optional) How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
// - `cachePolicy` - (Optional) Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
type UserFederation struct {
	pulumi.CustomResourceState

	// The number of users to sync within a single transaction.
	BatchSizeForSync pulumi.IntPtrOutput `pulumi:"batchSizeForSync"`
	// Password of LDAP admin.
	BindCredential pulumi.StringPtrOutput `pulumi:"bindCredential"`
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server.
	BindDn      pulumi.StringPtrOutput `pulumi:"bindDn"`
	CachePolicy pulumi.StringPtrOutput `pulumi:"cachePolicy"`
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users
	// sync.
	ChangedSyncPeriod pulumi.IntPtrOutput `pulumi:"changedSyncPeriod"`
	// LDAP connection timeout (duration string)
	ConnectionTimeout pulumi.StringPtrOutput `pulumi:"connectionTimeout"`
	// Connection URL to the LDAP server.
	ConnectionUrl pulumi.StringOutput `pulumi:"connectionUrl"`
	// Additional LDAP filter for filtering searched users. Must begin with '(' and end with ')'.
	CustomUserSearchFilter pulumi.StringPtrOutput `pulumi:"customUserSearchFilter"`
	// READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
	EditMode pulumi.StringPtrOutput `pulumi:"editMode"`
	// When false, this provider will not be used when performing queries for users.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod pulumi.IntPtrOutput `pulumi:"fullSyncPeriod"`
	// When true, LDAP users will be imported into the Keycloak database.
	ImportEnabled pulumi.BoolPtrOutput `pulumi:"importEnabled"`
	// Settings regarding kerberos authentication for this realm.
	Kerberos UserFederationKerberosPtrOutput `pulumi:"kerberos"`
	// Display name of the provider when displayed in the console.
	Name pulumi.StringOutput `pulumi:"name"`
	// When true, Keycloak assumes the LDAP server supports pagination.
	Pagination pulumi.BoolPtrOutput `pulumi:"pagination"`
	// Priority of this provider when looking up users. Lower values are first.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute pulumi.StringOutput `pulumi:"rdnLdapAttribute"`
	// LDAP read timeout (duration string)
	ReadTimeout pulumi.StringPtrOutput `pulumi:"readTimeout"`
	// The realm this provider will provide user federation for.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
	SearchScope pulumi.StringPtrOutput `pulumi:"searchScope"`
	// When true, newly created users will be synced back to LDAP.
	SyncRegistrations pulumi.BoolPtrOutput   `pulumi:"syncRegistrations"`
	UseTruststoreSpi  pulumi.StringPtrOutput `pulumi:"useTruststoreSpi"`
	// All values of LDAP objectClass attribute for users in LDAP.
	UserObjectClasses pulumi.StringArrayOutput `pulumi:"userObjectClasses"`
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute pulumi.StringOutput `pulumi:"usernameLdapAttribute"`
	// Full DN of LDAP tree where your users are.
	UsersDn pulumi.StringOutput `pulumi:"usersDn"`
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute pulumi.StringOutput `pulumi:"uuidLdapAttribute"`
	// When true, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy pulumi.BoolPtrOutput `pulumi:"validatePasswordPolicy"`
	// LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
	Vendor pulumi.StringPtrOutput `pulumi:"vendor"`
}

// NewUserFederation registers a new resource with the given unique name, arguments, and options.
func NewUserFederation(ctx *pulumi.Context,
	name string, args *UserFederationArgs, opts ...pulumi.ResourceOption) (*UserFederation, error) {
	if args == nil || args.ConnectionUrl == nil {
		return nil, errors.New("missing required argument 'ConnectionUrl'")
	}
	if args == nil || args.RdnLdapAttribute == nil {
		return nil, errors.New("missing required argument 'RdnLdapAttribute'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil || args.UserObjectClasses == nil {
		return nil, errors.New("missing required argument 'UserObjectClasses'")
	}
	if args == nil || args.UsernameLdapAttribute == nil {
		return nil, errors.New("missing required argument 'UsernameLdapAttribute'")
	}
	if args == nil || args.UsersDn == nil {
		return nil, errors.New("missing required argument 'UsersDn'")
	}
	if args == nil || args.UuidLdapAttribute == nil {
		return nil, errors.New("missing required argument 'UuidLdapAttribute'")
	}
	if args == nil {
		args = &UserFederationArgs{}
	}
	var resource UserFederation
	err := ctx.RegisterResource("keycloak:ldap/userFederation:UserFederation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserFederation gets an existing UserFederation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserFederation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserFederationState, opts ...pulumi.ResourceOption) (*UserFederation, error) {
	var resource UserFederation
	err := ctx.ReadResource("keycloak:ldap/userFederation:UserFederation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserFederation resources.
type userFederationState struct {
	// The number of users to sync within a single transaction.
	BatchSizeForSync *int `pulumi:"batchSizeForSync"`
	// Password of LDAP admin.
	BindCredential *string `pulumi:"bindCredential"`
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server.
	BindDn      *string `pulumi:"bindDn"`
	CachePolicy *string `pulumi:"cachePolicy"`
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users
	// sync.
	ChangedSyncPeriod *int `pulumi:"changedSyncPeriod"`
	// LDAP connection timeout (duration string)
	ConnectionTimeout *string `pulumi:"connectionTimeout"`
	// Connection URL to the LDAP server.
	ConnectionUrl *string `pulumi:"connectionUrl"`
	// Additional LDAP filter for filtering searched users. Must begin with '(' and end with ')'.
	CustomUserSearchFilter *string `pulumi:"customUserSearchFilter"`
	// READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
	EditMode *string `pulumi:"editMode"`
	// When false, this provider will not be used when performing queries for users.
	Enabled *bool `pulumi:"enabled"`
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod *int `pulumi:"fullSyncPeriod"`
	// When true, LDAP users will be imported into the Keycloak database.
	ImportEnabled *bool `pulumi:"importEnabled"`
	// Settings regarding kerberos authentication for this realm.
	Kerberos *UserFederationKerberos `pulumi:"kerberos"`
	// Display name of the provider when displayed in the console.
	Name *string `pulumi:"name"`
	// When true, Keycloak assumes the LDAP server supports pagination.
	Pagination *bool `pulumi:"pagination"`
	// Priority of this provider when looking up users. Lower values are first.
	Priority *int `pulumi:"priority"`
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute *string `pulumi:"rdnLdapAttribute"`
	// LDAP read timeout (duration string)
	ReadTimeout *string `pulumi:"readTimeout"`
	// The realm this provider will provide user federation for.
	RealmId *string `pulumi:"realmId"`
	// ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
	SearchScope *string `pulumi:"searchScope"`
	// When true, newly created users will be synced back to LDAP.
	SyncRegistrations *bool   `pulumi:"syncRegistrations"`
	UseTruststoreSpi  *string `pulumi:"useTruststoreSpi"`
	// All values of LDAP objectClass attribute for users in LDAP.
	UserObjectClasses []string `pulumi:"userObjectClasses"`
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute *string `pulumi:"usernameLdapAttribute"`
	// Full DN of LDAP tree where your users are.
	UsersDn *string `pulumi:"usersDn"`
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute *string `pulumi:"uuidLdapAttribute"`
	// When true, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy *bool `pulumi:"validatePasswordPolicy"`
	// LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
	Vendor *string `pulumi:"vendor"`
}

type UserFederationState struct {
	// The number of users to sync within a single transaction.
	BatchSizeForSync pulumi.IntPtrInput
	// Password of LDAP admin.
	BindCredential pulumi.StringPtrInput
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server.
	BindDn      pulumi.StringPtrInput
	CachePolicy pulumi.StringPtrInput
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users
	// sync.
	ChangedSyncPeriod pulumi.IntPtrInput
	// LDAP connection timeout (duration string)
	ConnectionTimeout pulumi.StringPtrInput
	// Connection URL to the LDAP server.
	ConnectionUrl pulumi.StringPtrInput
	// Additional LDAP filter for filtering searched users. Must begin with '(' and end with ')'.
	CustomUserSearchFilter pulumi.StringPtrInput
	// READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
	EditMode pulumi.StringPtrInput
	// When false, this provider will not be used when performing queries for users.
	Enabled pulumi.BoolPtrInput
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod pulumi.IntPtrInput
	// When true, LDAP users will be imported into the Keycloak database.
	ImportEnabled pulumi.BoolPtrInput
	// Settings regarding kerberos authentication for this realm.
	Kerberos UserFederationKerberosPtrInput
	// Display name of the provider when displayed in the console.
	Name pulumi.StringPtrInput
	// When true, Keycloak assumes the LDAP server supports pagination.
	Pagination pulumi.BoolPtrInput
	// Priority of this provider when looking up users. Lower values are first.
	Priority pulumi.IntPtrInput
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute pulumi.StringPtrInput
	// LDAP read timeout (duration string)
	ReadTimeout pulumi.StringPtrInput
	// The realm this provider will provide user federation for.
	RealmId pulumi.StringPtrInput
	// ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
	SearchScope pulumi.StringPtrInput
	// When true, newly created users will be synced back to LDAP.
	SyncRegistrations pulumi.BoolPtrInput
	UseTruststoreSpi  pulumi.StringPtrInput
	// All values of LDAP objectClass attribute for users in LDAP.
	UserObjectClasses pulumi.StringArrayInput
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute pulumi.StringPtrInput
	// Full DN of LDAP tree where your users are.
	UsersDn pulumi.StringPtrInput
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute pulumi.StringPtrInput
	// When true, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy pulumi.BoolPtrInput
	// LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
	Vendor pulumi.StringPtrInput
}

func (UserFederationState) ElementType() reflect.Type {
	return reflect.TypeOf((*userFederationState)(nil)).Elem()
}

type userFederationArgs struct {
	// The number of users to sync within a single transaction.
	BatchSizeForSync *int `pulumi:"batchSizeForSync"`
	// Password of LDAP admin.
	BindCredential *string `pulumi:"bindCredential"`
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server.
	BindDn      *string `pulumi:"bindDn"`
	CachePolicy *string `pulumi:"cachePolicy"`
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users
	// sync.
	ChangedSyncPeriod *int `pulumi:"changedSyncPeriod"`
	// LDAP connection timeout (duration string)
	ConnectionTimeout *string `pulumi:"connectionTimeout"`
	// Connection URL to the LDAP server.
	ConnectionUrl string `pulumi:"connectionUrl"`
	// Additional LDAP filter for filtering searched users. Must begin with '(' and end with ')'.
	CustomUserSearchFilter *string `pulumi:"customUserSearchFilter"`
	// READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
	EditMode *string `pulumi:"editMode"`
	// When false, this provider will not be used when performing queries for users.
	Enabled *bool `pulumi:"enabled"`
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod *int `pulumi:"fullSyncPeriod"`
	// When true, LDAP users will be imported into the Keycloak database.
	ImportEnabled *bool `pulumi:"importEnabled"`
	// Settings regarding kerberos authentication for this realm.
	Kerberos *UserFederationKerberos `pulumi:"kerberos"`
	// Display name of the provider when displayed in the console.
	Name *string `pulumi:"name"`
	// When true, Keycloak assumes the LDAP server supports pagination.
	Pagination *bool `pulumi:"pagination"`
	// Priority of this provider when looking up users. Lower values are first.
	Priority *int `pulumi:"priority"`
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute string `pulumi:"rdnLdapAttribute"`
	// LDAP read timeout (duration string)
	ReadTimeout *string `pulumi:"readTimeout"`
	// The realm this provider will provide user federation for.
	RealmId string `pulumi:"realmId"`
	// ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
	SearchScope *string `pulumi:"searchScope"`
	// When true, newly created users will be synced back to LDAP.
	SyncRegistrations *bool   `pulumi:"syncRegistrations"`
	UseTruststoreSpi  *string `pulumi:"useTruststoreSpi"`
	// All values of LDAP objectClass attribute for users in LDAP.
	UserObjectClasses []string `pulumi:"userObjectClasses"`
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute string `pulumi:"usernameLdapAttribute"`
	// Full DN of LDAP tree where your users are.
	UsersDn string `pulumi:"usersDn"`
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute string `pulumi:"uuidLdapAttribute"`
	// When true, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy *bool `pulumi:"validatePasswordPolicy"`
	// LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
	Vendor *string `pulumi:"vendor"`
}

// The set of arguments for constructing a UserFederation resource.
type UserFederationArgs struct {
	// The number of users to sync within a single transaction.
	BatchSizeForSync pulumi.IntPtrInput
	// Password of LDAP admin.
	BindCredential pulumi.StringPtrInput
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server.
	BindDn      pulumi.StringPtrInput
	CachePolicy pulumi.StringPtrInput
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users
	// sync.
	ChangedSyncPeriod pulumi.IntPtrInput
	// LDAP connection timeout (duration string)
	ConnectionTimeout pulumi.StringPtrInput
	// Connection URL to the LDAP server.
	ConnectionUrl pulumi.StringInput
	// Additional LDAP filter for filtering searched users. Must begin with '(' and end with ')'.
	CustomUserSearchFilter pulumi.StringPtrInput
	// READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
	EditMode pulumi.StringPtrInput
	// When false, this provider will not be used when performing queries for users.
	Enabled pulumi.BoolPtrInput
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod pulumi.IntPtrInput
	// When true, LDAP users will be imported into the Keycloak database.
	ImportEnabled pulumi.BoolPtrInput
	// Settings regarding kerberos authentication for this realm.
	Kerberos UserFederationKerberosPtrInput
	// Display name of the provider when displayed in the console.
	Name pulumi.StringPtrInput
	// When true, Keycloak assumes the LDAP server supports pagination.
	Pagination pulumi.BoolPtrInput
	// Priority of this provider when looking up users. Lower values are first.
	Priority pulumi.IntPtrInput
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute pulumi.StringInput
	// LDAP read timeout (duration string)
	ReadTimeout pulumi.StringPtrInput
	// The realm this provider will provide user federation for.
	RealmId pulumi.StringInput
	// ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
	SearchScope pulumi.StringPtrInput
	// When true, newly created users will be synced back to LDAP.
	SyncRegistrations pulumi.BoolPtrInput
	UseTruststoreSpi  pulumi.StringPtrInput
	// All values of LDAP objectClass attribute for users in LDAP.
	UserObjectClasses pulumi.StringArrayInput
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute pulumi.StringInput
	// Full DN of LDAP tree where your users are.
	UsersDn pulumi.StringInput
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UuidLdapAttribute pulumi.StringInput
	// When true, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy pulumi.BoolPtrInput
	// LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
	Vendor pulumi.StringPtrInput
}

func (UserFederationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userFederationArgs)(nil)).Elem()
}
