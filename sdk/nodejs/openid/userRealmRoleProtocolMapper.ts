// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## # keycloak.openid.UserRealmRoleProtocolMapper
 *
 * Allows for creating and managing user realm role protocol mappers within
 * Keycloak.
 *
 * User realm role protocol mappers allow you to define a claim containing the list of the realm roles.
 * Protocol mappers can be defined for a single client, or they can
 * be defined for a client scope which can be shared between multiple different
 * clients.
 *
 * ### Example Usage (Client)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "my-realm",
 * });
 * const openidClient = new keycloak.openid.Client("openidClient", {
 *     accessType: "CONFIDENTIAL",
 *     clientId: "test-client",
 *     enabled: true,
 *     realmId: realm.id,
 *     validRedirectUris: ["http://localhost:8080/openid-callback"],
 * });
 * const userRealmRoleMapper = new keycloak.openid.UserRealmRoleProtocolMapper("userRealmRoleMapper", {
 *     claimName: "foo",
 *     clientId: openidClient.id,
 *     realmId: realm.id,
 * });
 * ```
 *
 * ### Example Usage (Client Scope)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "my-realm",
 * });
 * const clientScope = new keycloak.openid.ClientScope("clientScope", {
 *     realmId: realm.id,
 * });
 * const userRealmRoleMapper = new keycloak.openid.UserRealmRoleProtocolMapper("userRealmRoleMapper", {
 *     claimName: "foo",
 *     clientScopeId: clientScope.id,
 *     realmId: realm.id,
 * });
 * ```
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm this protocol mapper exists within.
 * - `clientId` - (Required if `clientScopeId` is not specified) The client this protocol mapper is attached to.
 * - `clientScopeId` - (Required if `clientId` is not specified) The client scope this protocol mapper is attached to.
 * - `name` - (Required) The display name of this protocol mapper in the GUI.
 * - `claimName` - (Required) The name of the claim to insert into a token.
 * - `claimValueType` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
 * - `multivalued` - (Optional) Indicates if attribute supports multiple values. If true, then the list of all values of this attribute will be set as claim. If false, then just first value will be set as claim. Defaults to `true`.
 * - `realmRolePrefix` - (Optional) A prefix for each Realm Role.
 * - `addToIdToken` - (Optional) Indicates if the property should be added as a claim to the id token. Defaults to `true`.
 * - `addToAccessToken` - (Optional) Indicates if the property should be added as a claim to the access token. Defaults to `true`.
 * - `addToUserinfo` - (Optional) Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
 */
export class UserRealmRoleProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing UserRealmRoleProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserRealmRoleProtocolMapperState, opts?: pulumi.CustomResourceOptions): UserRealmRoleProtocolMapper {
        return new UserRealmRoleProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/userRealmRoleProtocolMapper:UserRealmRoleProtocolMapper';

    /**
     * Returns true if the given object is an instance of UserRealmRoleProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserRealmRoleProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserRealmRoleProtocolMapper.__pulumiType;
    }

    /**
     * Indicates if the attribute should be a claim in the access token.
     */
    public readonly addToAccessToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the attribute should be a claim in the id token.
     */
    public readonly addToIdToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the attribute should appear in the userinfo response body.
     */
    public readonly addToUserinfo!: pulumi.Output<boolean | undefined>;
    public readonly claimName!: pulumi.Output<string>;
    /**
     * Claim type used when serializing tokens.
     */
    public readonly claimValueType!: pulumi.Output<string | undefined>;
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether this attribute is a single value or an array of values.
     */
    public readonly multivalued!: pulumi.Output<boolean | undefined>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * Prefix that will be added to each realm role.
     */
    public readonly realmRolePrefix!: pulumi.Output<string | undefined>;

    /**
     * Create a UserRealmRoleProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserRealmRoleProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserRealmRoleProtocolMapperArgs | UserRealmRoleProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserRealmRoleProtocolMapperState | undefined;
            inputs["addToAccessToken"] = state ? state.addToAccessToken : undefined;
            inputs["addToIdToken"] = state ? state.addToIdToken : undefined;
            inputs["addToUserinfo"] = state ? state.addToUserinfo : undefined;
            inputs["claimName"] = state ? state.claimName : undefined;
            inputs["claimValueType"] = state ? state.claimValueType : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            inputs["multivalued"] = state ? state.multivalued : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["realmRolePrefix"] = state ? state.realmRolePrefix : undefined;
        } else {
            const args = argsOrState as UserRealmRoleProtocolMapperArgs | undefined;
            if (!args || args.claimName === undefined) {
                throw new Error("Missing required property 'claimName'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["addToAccessToken"] = args ? args.addToAccessToken : undefined;
            inputs["addToIdToken"] = args ? args.addToIdToken : undefined;
            inputs["addToUserinfo"] = args ? args.addToUserinfo : undefined;
            inputs["claimName"] = args ? args.claimName : undefined;
            inputs["claimValueType"] = args ? args.claimValueType : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            inputs["multivalued"] = args ? args.multivalued : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["realmRolePrefix"] = args ? args.realmRolePrefix : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserRealmRoleProtocolMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserRealmRoleProtocolMapper resources.
 */
export interface UserRealmRoleProtocolMapperState {
    /**
     * Indicates if the attribute should be a claim in the access token.
     */
    readonly addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the attribute should be a claim in the id token.
     */
    readonly addToIdToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the attribute should appear in the userinfo response body.
     */
    readonly addToUserinfo?: pulumi.Input<boolean>;
    readonly claimName?: pulumi.Input<string>;
    /**
     * Claim type used when serializing tokens.
     */
    readonly claimValueType?: pulumi.Input<string>;
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * Indicates whether this attribute is a single value or an array of values.
     */
    readonly multivalued?: pulumi.Input<boolean>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    readonly realmId?: pulumi.Input<string>;
    /**
     * Prefix that will be added to each realm role.
     */
    readonly realmRolePrefix?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserRealmRoleProtocolMapper resource.
 */
export interface UserRealmRoleProtocolMapperArgs {
    /**
     * Indicates if the attribute should be a claim in the access token.
     */
    readonly addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the attribute should be a claim in the id token.
     */
    readonly addToIdToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the attribute should appear in the userinfo response body.
     */
    readonly addToUserinfo?: pulumi.Input<boolean>;
    readonly claimName: pulumi.Input<string>;
    /**
     * Claim type used when serializing tokens.
     */
    readonly claimValueType?: pulumi.Input<string>;
    /**
     * The mapper's associated client. Cannot be used at the same time as client_scope_id.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The mapper's associated client scope. Cannot be used at the same time as client_id.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * Indicates whether this attribute is a single value or an array of values.
     */
    readonly multivalued?: pulumi.Input<boolean>;
    /**
     * A human-friendly name that will appear in the Keycloak console.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm id where the associated client or client scope exists.
     */
    readonly realmId: pulumi.Input<string>;
    /**
     * Prefix that will be added to each realm role.
     */
    readonly realmRolePrefix?: pulumi.Input<string>;
}
