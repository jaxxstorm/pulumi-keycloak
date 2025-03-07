// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";

export interface GetRealmInternationalization {
    defaultLocale?: string;
    supportedLocales?: string[];
}

export interface GetRealmSecurityDefense {
    bruteForceDetections?: inputs.GetRealmSecurityDefenseBruteForceDetection[];
    headers?: inputs.GetRealmSecurityDefenseHeader[];
}

export interface GetRealmSecurityDefenseBruteForceDetection {
    failureResetTimeSeconds?: number;
    maxFailureWaitSeconds?: number;
    maxLoginFailures?: number;
    minimumQuickLoginWaitSeconds?: number;
    permanentLockout?: boolean;
    quickLoginCheckMilliSeconds?: number;
    waitIncrementSeconds?: number;
}

export interface GetRealmSecurityDefenseHeader {
    contentSecurityPolicy?: string;
    contentSecurityPolicyReportOnly?: string;
    strictTransportSecurity?: string;
    xContentTypeOptions?: string;
    xFrameOptions?: string;
    xRobotsTag?: string;
    xXssProtection?: string;
}

export interface GetRealmSmtpServer {
    auths?: inputs.GetRealmSmtpServerAuth[];
    envelopeFrom?: string;
    from?: string;
    fromDisplayName?: string;
    host?: string;
    port?: string;
    replyTo?: string;
    replyToDisplayName?: string;
    ssl?: boolean;
    starttls?: boolean;
}

export interface GetRealmSmtpServerAuth {
    password?: string;
    username?: string;
}

export interface RealmInternationalization {
    defaultLocale: pulumi.Input<string>;
    supportedLocales: pulumi.Input<pulumi.Input<string>[]>;
}

export interface RealmSecurityDefenses {
    bruteForceDetection?: pulumi.Input<inputs.RealmSecurityDefensesBruteForceDetection>;
    headers?: pulumi.Input<inputs.RealmSecurityDefensesHeaders>;
}

export interface RealmSecurityDefensesBruteForceDetection {
    failureResetTimeSeconds?: pulumi.Input<number>;
    maxFailureWaitSeconds?: pulumi.Input<number>;
    maxLoginFailures?: pulumi.Input<number>;
    minimumQuickLoginWaitSeconds?: pulumi.Input<number>;
    permanentLockout?: pulumi.Input<boolean>;
    quickLoginCheckMilliSeconds?: pulumi.Input<number>;
    waitIncrementSeconds?: pulumi.Input<number>;
}

export interface RealmSecurityDefensesHeaders {
    contentSecurityPolicy?: pulumi.Input<string>;
    contentSecurityPolicyReportOnly?: pulumi.Input<string>;
    strictTransportSecurity?: pulumi.Input<string>;
    xContentTypeOptions?: pulumi.Input<string>;
    xFrameOptions?: pulumi.Input<string>;
    xRobotsTag?: pulumi.Input<string>;
    xXssProtection?: pulumi.Input<string>;
}

export interface RealmSmtpServer {
    auth?: pulumi.Input<inputs.RealmSmtpServerAuth>;
    envelopeFrom?: pulumi.Input<string>;
    from: pulumi.Input<string>;
    fromDisplayName?: pulumi.Input<string>;
    host: pulumi.Input<string>;
    port?: pulumi.Input<string>;
    replyTo?: pulumi.Input<string>;
    replyToDisplayName?: pulumi.Input<string>;
    ssl?: pulumi.Input<boolean>;
    starttls?: pulumi.Input<boolean>;
}

export interface RealmSmtpServerAuth {
    password: pulumi.Input<string>;
    username: pulumi.Input<string>;
}

export interface UserFederatedIdentity {
    identityProvider: pulumi.Input<string>;
    userId: pulumi.Input<string>;
    userName: pulumi.Input<string>;
}

export interface UserInitialPassword {
    temporary?: pulumi.Input<boolean>;
    value: pulumi.Input<string>;
}

export namespace ldap {
    export interface UserFederationKerberos {
        kerberosRealm: pulumi.Input<string>;
        keyTab: pulumi.Input<string>;
        serverPrincipal: pulumi.Input<string>;
        useKerberosForPasswordAuthentication?: pulumi.Input<boolean>;
    }
}

export namespace openid {
    export interface ClientAuthenticationFlowBindingOverrides {
        browserId?: pulumi.Input<string>;
        directGrantId?: pulumi.Input<string>;
    }

    export interface ClientAuthorization {
        allowRemoteResourceManagement?: pulumi.Input<boolean>;
        keepDefaults?: pulumi.Input<boolean>;
        policyEnforcementMode: pulumi.Input<string>;
    }

    export interface ClientGroupPolicyGroup {
        extendChildren: pulumi.Input<boolean>;
        id: pulumi.Input<string>;
        path: pulumi.Input<string>;
    }

    export interface ClientRolePolicyRole {
        id: pulumi.Input<string>;
        required: pulumi.Input<boolean>;
    }
}
