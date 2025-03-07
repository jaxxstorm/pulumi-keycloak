// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # keycloak.ldap.UserAttributeMapper
 *
 * Allows for creating and managing user attribute mappers for Keycloak users
 * federated via LDAP.
 *
 * The LDAP user attribute mapper can be used to map a single LDAP attribute
 * to an attribute on the Keycloak user model.
 *
 * ### Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "test",
 * });
 * const ldapUserFederation = new keycloak.ldap.UserFederation("ldapUserFederation", {
 *     bindCredential: "admin",
 *     bindDn: "cn=admin,dc=example,dc=org",
 *     connectionUrl: "ldap://openldap",
 *     rdnLdapAttribute: "cn",
 *     realmId: realm.id,
 *     userObjectClasses: [
 *         "simpleSecurityObject",
 *         "organizationalRole",
 *     ],
 *     usernameLdapAttribute: "cn",
 *     usersDn: "dc=example,dc=org",
 *     uuidLdapAttribute: "entryDN",
 * });
 * const ldapUserAttributeMapper = new keycloak.ldap.UserAttributeMapper("ldapUserAttributeMapper", {
 *     ldapAttribute: "bar",
 *     ldapUserFederationId: ldapUserFederation.id,
 *     realmId: realm.id,
 *     userModelAttribute: "foo",
 * });
 * ```
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm that this LDAP mapper will exist in.
 * - `ldapUserFederationId` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
 * - `name` - (Required) Display name of this mapper when displayed in the console.
 * - `userModelAttribute` - (Required) Name of the user property or attribute you want to map the LDAP attribute into.
 * - `ldapAttribute` - (Required) Name of the mapped attribute on the LDAP object.
 * - `readOnly` - (Optional) When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
 * - `alwaysReadValueFromLdap` - (Optional) When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
 * - `isMandatoryInLdap` - (Optional) When `true`, this attribute must exist in LDAP. Defaults to `false`.
 */
export class UserAttributeMapper extends pulumi.CustomResource {
    /**
     * Get an existing UserAttributeMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserAttributeMapperState, opts?: pulumi.CustomResourceOptions): UserAttributeMapper {
        return new UserAttributeMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:ldap/userAttributeMapper:UserAttributeMapper';

    /**
     * Returns true if the given object is an instance of UserAttributeMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserAttributeMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserAttributeMapper.__pulumiType;
    }

    /**
     * When true, the value fetched from LDAP will override the value stored in Keycloak.
     */
    public readonly alwaysReadValueFromLdap!: pulumi.Output<boolean | undefined>;
    /**
     * When true, this attribute must exist in LDAP.
     */
    public readonly isMandatoryInLdap!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the mapped attribute on LDAP object.
     */
    public readonly ldapAttribute!: pulumi.Output<string>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    public readonly ldapUserFederationId!: pulumi.Output<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * When true, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak.
     */
    public readonly readOnly!: pulumi.Output<boolean | undefined>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * Name of the UserModel property or attribute you want to map the LDAP attribute into.
     */
    public readonly userModelAttribute!: pulumi.Output<string>;

    /**
     * Create a UserAttributeMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserAttributeMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserAttributeMapperArgs | UserAttributeMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserAttributeMapperState | undefined;
            inputs["alwaysReadValueFromLdap"] = state ? state.alwaysReadValueFromLdap : undefined;
            inputs["isMandatoryInLdap"] = state ? state.isMandatoryInLdap : undefined;
            inputs["ldapAttribute"] = state ? state.ldapAttribute : undefined;
            inputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["readOnly"] = state ? state.readOnly : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["userModelAttribute"] = state ? state.userModelAttribute : undefined;
        } else {
            const args = argsOrState as UserAttributeMapperArgs | undefined;
            if (!args || args.ldapAttribute === undefined) {
                throw new Error("Missing required property 'ldapAttribute'");
            }
            if (!args || args.ldapUserFederationId === undefined) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.userModelAttribute === undefined) {
                throw new Error("Missing required property 'userModelAttribute'");
            }
            inputs["alwaysReadValueFromLdap"] = args ? args.alwaysReadValueFromLdap : undefined;
            inputs["isMandatoryInLdap"] = args ? args.isMandatoryInLdap : undefined;
            inputs["ldapAttribute"] = args ? args.ldapAttribute : undefined;
            inputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["readOnly"] = args ? args.readOnly : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["userModelAttribute"] = args ? args.userModelAttribute : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserAttributeMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserAttributeMapper resources.
 */
export interface UserAttributeMapperState {
    /**
     * When true, the value fetched from LDAP will override the value stored in Keycloak.
     */
    readonly alwaysReadValueFromLdap?: pulumi.Input<boolean>;
    /**
     * When true, this attribute must exist in LDAP.
     */
    readonly isMandatoryInLdap?: pulumi.Input<boolean>;
    /**
     * Name of the mapped attribute on LDAP object.
     */
    readonly ldapAttribute?: pulumi.Input<string>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    readonly ldapUserFederationId?: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * When true, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak.
     */
    readonly readOnly?: pulumi.Input<boolean>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    readonly realmId?: pulumi.Input<string>;
    /**
     * Name of the UserModel property or attribute you want to map the LDAP attribute into.
     */
    readonly userModelAttribute?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserAttributeMapper resource.
 */
export interface UserAttributeMapperArgs {
    /**
     * When true, the value fetched from LDAP will override the value stored in Keycloak.
     */
    readonly alwaysReadValueFromLdap?: pulumi.Input<boolean>;
    /**
     * When true, this attribute must exist in LDAP.
     */
    readonly isMandatoryInLdap?: pulumi.Input<boolean>;
    /**
     * Name of the mapped attribute on LDAP object.
     */
    readonly ldapAttribute: pulumi.Input<string>;
    /**
     * The ldap user federation provider to attach this mapper to.
     */
    readonly ldapUserFederationId: pulumi.Input<string>;
    /**
     * Display name of the mapper when displayed in the console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * When true, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak.
     */
    readonly readOnly?: pulumi.Input<boolean>;
    /**
     * The realm in which the ldap user federation provider exists.
     */
    readonly realmId: pulumi.Input<string>;
    /**
     * Name of the UserModel property or attribute you want to map the LDAP attribute into.
     */
    readonly userModelAttribute: pulumi.Input<string>;
}
