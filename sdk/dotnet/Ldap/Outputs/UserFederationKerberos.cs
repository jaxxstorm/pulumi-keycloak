// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Ldap.Outputs
{

    [OutputType]
    public sealed class UserFederationKerberos
    {
        public readonly string KerberosRealm;
        public readonly string KeyTab;
        public readonly string ServerPrincipal;
        public readonly bool? UseKerberosForPasswordAuthentication;

        [OutputConstructor]
        private UserFederationKerberos(
            string kerberosRealm,

            string keyTab,

            string serverPrincipal,

            bool? useKerberosForPasswordAuthentication)
        {
            KerberosRealm = kerberosRealm;
            KeyTab = keyTab;
            ServerPrincipal = serverPrincipal;
            UseKerberosForPasswordAuthentication = useKerberosForPasswordAuthentication;
        }
    }
}
