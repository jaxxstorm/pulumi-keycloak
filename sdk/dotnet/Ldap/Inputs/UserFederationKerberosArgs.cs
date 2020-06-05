// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Ldap.Inputs
{

    public sealed class UserFederationKerberosArgs : Pulumi.ResourceArgs
    {
        [Input("kerberosRealm", required: true)]
        public Input<string> KerberosRealm { get; set; } = null!;

        [Input("keyTab", required: true)]
        public Input<string> KeyTab { get; set; } = null!;

        [Input("serverPrincipal", required: true)]
        public Input<string> ServerPrincipal { get; set; } = null!;

        [Input("useKerberosForPasswordAuthentication")]
        public Input<bool>? UseKerberosForPasswordAuthentication { get; set; }

        public UserFederationKerberosArgs()
        {
        }
    }
}
