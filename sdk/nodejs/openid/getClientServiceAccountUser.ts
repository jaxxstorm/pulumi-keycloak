// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getClientServiceAccountUser(args: GetClientServiceAccountUserArgs, opts?: pulumi.InvokeOptions): Promise<GetClientServiceAccountUserResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("keycloak:openid/getClientServiceAccountUser:getClientServiceAccountUser", {
        "clientId": args.clientId,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getClientServiceAccountUser.
 */
export interface GetClientServiceAccountUserArgs {
    readonly clientId: string;
    readonly realmId: string;
}

/**
 * A collection of values returned by getClientServiceAccountUser.
 */
export interface GetClientServiceAccountUserResult {
    readonly attributes: {[key: string]: any};
    readonly clientId: string;
    readonly email: string;
    readonly enabled: boolean;
    readonly federatedIdentities: outputs.openid.GetClientServiceAccountUserFederatedIdentity[];
    readonly firstName: string;
    readonly lastName: string;
    readonly realmId: string;
    readonly username: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
