// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ClientPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ClientPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientPolicyState, opts?: pulumi.CustomResourceOptions): ClientPolicy {
        return new ClientPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientPolicy:ClientPolicy';

    /**
     * Returns true if the given object is an instance of ClientPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientPolicy.__pulumiType;
    }

    public readonly clients!: pulumi.Output<string[]>;
    public readonly decisionStrategy!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly logic!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly resourceServerId!: pulumi.Output<string>;

    /**
     * Create a ClientPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientPolicyArgs | ClientPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClientPolicyState | undefined;
            inputs["clients"] = state ? state.clients : undefined;
            inputs["decisionStrategy"] = state ? state.decisionStrategy : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["logic"] = state ? state.logic : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["resourceServerId"] = state ? state.resourceServerId : undefined;
        } else {
            const args = argsOrState as ClientPolicyArgs | undefined;
            if (!args || args.clients === undefined) {
                throw new Error("Missing required property 'clients'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.resourceServerId === undefined) {
                throw new Error("Missing required property 'resourceServerId'");
            }
            inputs["clients"] = args ? args.clients : undefined;
            inputs["decisionStrategy"] = args ? args.decisionStrategy : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["logic"] = args ? args.logic : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["resourceServerId"] = args ? args.resourceServerId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClientPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientPolicy resources.
 */
export interface ClientPolicyState {
    readonly clients?: pulumi.Input<pulumi.Input<string>[]>;
    readonly decisionStrategy?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly logic?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly realmId?: pulumi.Input<string>;
    readonly resourceServerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientPolicy resource.
 */
export interface ClientPolicyArgs {
    readonly clients: pulumi.Input<pulumi.Input<string>[]>;
    readonly decisionStrategy?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly logic?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
    readonly resourceServerId: pulumi.Input<string>;
}
