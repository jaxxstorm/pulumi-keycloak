// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public static partial class Invokes
    {
        /// <summary>
        /// ## # keycloak..Realm data source
        /// 
        /// This data source can be used to fetch properties of a Keycloak realm for
        /// usage with other resources.
        /// 
        /// ### Argument Reference
        /// 
        /// The following arguments are supported:
        /// 
        /// - `realm` - (Required) The realm name.
        /// 
        /// ### Attributes Reference
        /// 
        /// See the docs for the `keycloak..Realm` resource for details on the exported attributes.
        /// 
        /// &gt; This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/d/keycloak_realm.html.markdown.
        /// </summary>
        [Obsolete("Use GetRealm.InvokeAsync() instead")]
        public static Task<GetRealmResult> GetRealm(GetRealmArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRealmResult>("keycloak:index/getRealm:getRealm", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetRealm
    {
        /// <summary>
        /// ## # keycloak..Realm data source
        /// 
        /// This data source can be used to fetch properties of a Keycloak realm for
        /// usage with other resources.
        /// 
        /// ### Argument Reference
        /// 
        /// The following arguments are supported:
        /// 
        /// - `realm` - (Required) The realm name.
        /// 
        /// ### Attributes Reference
        /// 
        /// See the docs for the `keycloak..Realm` resource for details on the exported attributes.
        /// 
        /// &gt; This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/d/keycloak_realm.html.markdown.
        /// </summary>
        public static Task<GetRealmResult> InvokeAsync(GetRealmArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRealmResult>("keycloak:index/getRealm:getRealm", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetRealmArgs : Pulumi.InvokeArgs
    {
        [Input("attributes")]
        private Dictionary<string, object>? _attributes;
        public Dictionary<string, object> Attributes
        {
            get => _attributes ?? (_attributes = new Dictionary<string, object>());
            set => _attributes = value;
        }

        [Input("displayNameHtml")]
        public string? DisplayNameHtml { get; set; }

        [Input("internationalizations")]
        private List<Inputs.GetRealmInternationalizationsArgs>? _internationalizations;
        public List<Inputs.GetRealmInternationalizationsArgs> Internationalizations
        {
            get => _internationalizations ?? (_internationalizations = new List<Inputs.GetRealmInternationalizationsArgs>());
            set => _internationalizations = value;
        }

        [Input("realm", required: true)]
        public string Realm { get; set; } = null!;

        [Input("securityDefenses")]
        private List<Inputs.GetRealmSecurityDefensesArgs>? _securityDefenses;
        public List<Inputs.GetRealmSecurityDefensesArgs> SecurityDefenses
        {
            get => _securityDefenses ?? (_securityDefenses = new List<Inputs.GetRealmSecurityDefensesArgs>());
            set => _securityDefenses = value;
        }

        [Input("smtpServers")]
        private List<Inputs.GetRealmSmtpServersArgs>? _smtpServers;
        public List<Inputs.GetRealmSmtpServersArgs> SmtpServers
        {
            get => _smtpServers ?? (_smtpServers = new List<Inputs.GetRealmSmtpServersArgs>());
            set => _smtpServers = value;
        }

        public GetRealmArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetRealmResult
    {
        public readonly string AccessCodeLifespan;
        public readonly string AccessCodeLifespanLogin;
        public readonly string AccessCodeLifespanUserAction;
        public readonly string AccessTokenLifespan;
        public readonly string AccessTokenLifespanForImplicitFlow;
        public readonly string AccountTheme;
        public readonly string ActionTokenGeneratedByAdminLifespan;
        public readonly string ActionTokenGeneratedByUserLifespan;
        public readonly string AdminTheme;
        public readonly ImmutableDictionary<string, object> Attributes;
        public readonly string BrowserFlow;
        public readonly string ClientAuthenticationFlow;
        public readonly string DirectGrantFlow;
        public readonly string DisplayName;
        public readonly string? DisplayNameHtml;
        public readonly string DockerAuthenticationFlow;
        public readonly bool DuplicateEmailsAllowed;
        public readonly bool EditUsernameAllowed;
        public readonly string EmailTheme;
        public readonly bool Enabled;
        public readonly ImmutableArray<Outputs.GetRealmInternationalizationsResult> Internationalizations;
        public readonly string LoginTheme;
        public readonly bool LoginWithEmailAllowed;
        public readonly string OfflineSessionIdleTimeout;
        public readonly string OfflineSessionMaxLifespan;
        public readonly string PasswordPolicy;
        public readonly string Realm;
        public readonly int RefreshTokenMaxReuse;
        public readonly bool RegistrationAllowed;
        public readonly bool RegistrationEmailAsUsername;
        public readonly string RegistrationFlow;
        public readonly bool RememberMe;
        public readonly string ResetCredentialsFlow;
        public readonly bool ResetPasswordAllowed;
        public readonly ImmutableArray<Outputs.GetRealmSecurityDefensesResult> SecurityDefenses;
        public readonly ImmutableArray<Outputs.GetRealmSmtpServersResult> SmtpServers;
        public readonly string SslRequired;
        public readonly string SsoSessionIdleTimeout;
        public readonly string SsoSessionMaxLifespan;
        public readonly bool VerifyEmail;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetRealmResult(
            string accessCodeLifespan,
            string accessCodeLifespanLogin,
            string accessCodeLifespanUserAction,
            string accessTokenLifespan,
            string accessTokenLifespanForImplicitFlow,
            string accountTheme,
            string actionTokenGeneratedByAdminLifespan,
            string actionTokenGeneratedByUserLifespan,
            string adminTheme,
            ImmutableDictionary<string, object> attributes,
            string browserFlow,
            string clientAuthenticationFlow,
            string directGrantFlow,
            string displayName,
            string? displayNameHtml,
            string dockerAuthenticationFlow,
            bool duplicateEmailsAllowed,
            bool editUsernameAllowed,
            string emailTheme,
            bool enabled,
            ImmutableArray<Outputs.GetRealmInternationalizationsResult> internationalizations,
            string loginTheme,
            bool loginWithEmailAllowed,
            string offlineSessionIdleTimeout,
            string offlineSessionMaxLifespan,
            string passwordPolicy,
            string realm,
            int refreshTokenMaxReuse,
            bool registrationAllowed,
            bool registrationEmailAsUsername,
            string registrationFlow,
            bool rememberMe,
            string resetCredentialsFlow,
            bool resetPasswordAllowed,
            ImmutableArray<Outputs.GetRealmSecurityDefensesResult> securityDefenses,
            ImmutableArray<Outputs.GetRealmSmtpServersResult> smtpServers,
            string sslRequired,
            string ssoSessionIdleTimeout,
            string ssoSessionMaxLifespan,
            bool verifyEmail,
            string id)
        {
            AccessCodeLifespan = accessCodeLifespan;
            AccessCodeLifespanLogin = accessCodeLifespanLogin;
            AccessCodeLifespanUserAction = accessCodeLifespanUserAction;
            AccessTokenLifespan = accessTokenLifespan;
            AccessTokenLifespanForImplicitFlow = accessTokenLifespanForImplicitFlow;
            AccountTheme = accountTheme;
            ActionTokenGeneratedByAdminLifespan = actionTokenGeneratedByAdminLifespan;
            ActionTokenGeneratedByUserLifespan = actionTokenGeneratedByUserLifespan;
            AdminTheme = adminTheme;
            Attributes = attributes;
            BrowserFlow = browserFlow;
            ClientAuthenticationFlow = clientAuthenticationFlow;
            DirectGrantFlow = directGrantFlow;
            DisplayName = displayName;
            DisplayNameHtml = displayNameHtml;
            DockerAuthenticationFlow = dockerAuthenticationFlow;
            DuplicateEmailsAllowed = duplicateEmailsAllowed;
            EditUsernameAllowed = editUsernameAllowed;
            EmailTheme = emailTheme;
            Enabled = enabled;
            Internationalizations = internationalizations;
            LoginTheme = loginTheme;
            LoginWithEmailAllowed = loginWithEmailAllowed;
            OfflineSessionIdleTimeout = offlineSessionIdleTimeout;
            OfflineSessionMaxLifespan = offlineSessionMaxLifespan;
            PasswordPolicy = passwordPolicy;
            Realm = realm;
            RefreshTokenMaxReuse = refreshTokenMaxReuse;
            RegistrationAllowed = registrationAllowed;
            RegistrationEmailAsUsername = registrationEmailAsUsername;
            RegistrationFlow = registrationFlow;
            RememberMe = rememberMe;
            ResetCredentialsFlow = resetCredentialsFlow;
            ResetPasswordAllowed = resetPasswordAllowed;
            SecurityDefenses = securityDefenses;
            SmtpServers = smtpServers;
            SslRequired = sslRequired;
            SsoSessionIdleTimeout = ssoSessionIdleTimeout;
            SsoSessionMaxLifespan = ssoSessionMaxLifespan;
            VerifyEmail = verifyEmail;
            Id = id;
        }
    }

    namespace Inputs
    {

    public sealed class GetRealmInternationalizationsArgs : Pulumi.InvokeArgs
    {
        [Input("defaultLocale")]
        public string? DefaultLocale { get; set; }

        [Input("supportedLocales")]
        private List<string>? _supportedLocales;
        public List<string> SupportedLocales
        {
            get => _supportedLocales ?? (_supportedLocales = new List<string>());
            set => _supportedLocales = value;
        }

        public GetRealmInternationalizationsArgs()
        {
        }
    }

    public sealed class GetRealmSecurityDefensesArgs : Pulumi.InvokeArgs
    {
        [Input("bruteForceDetections")]
        private List<GetRealmSecurityDefensesBruteForceDetectionsArgs>? _bruteForceDetections;
        public List<GetRealmSecurityDefensesBruteForceDetectionsArgs> BruteForceDetections
        {
            get => _bruteForceDetections ?? (_bruteForceDetections = new List<GetRealmSecurityDefensesBruteForceDetectionsArgs>());
            set => _bruteForceDetections = value;
        }

        [Input("headers")]
        private List<GetRealmSecurityDefensesHeadersArgs>? _headers;
        public List<GetRealmSecurityDefensesHeadersArgs> Headers
        {
            get => _headers ?? (_headers = new List<GetRealmSecurityDefensesHeadersArgs>());
            set => _headers = value;
        }

        public GetRealmSecurityDefensesArgs()
        {
        }
    }

    public sealed class GetRealmSecurityDefensesBruteForceDetectionsArgs : Pulumi.InvokeArgs
    {
        [Input("failureResetTimeSeconds")]
        public int? FailureResetTimeSeconds { get; set; }

        [Input("maxFailureWaitSeconds")]
        public int? MaxFailureWaitSeconds { get; set; }

        [Input("maxLoginFailures")]
        public int? MaxLoginFailures { get; set; }

        [Input("minimumQuickLoginWaitSeconds")]
        public int? MinimumQuickLoginWaitSeconds { get; set; }

        [Input("permanentLockout")]
        public bool? PermanentLockout { get; set; }

        [Input("quickLoginCheckMilliSeconds")]
        public int? QuickLoginCheckMilliSeconds { get; set; }

        [Input("waitIncrementSeconds")]
        public int? WaitIncrementSeconds { get; set; }

        public GetRealmSecurityDefensesBruteForceDetectionsArgs()
        {
        }
    }

    public sealed class GetRealmSecurityDefensesHeadersArgs : Pulumi.InvokeArgs
    {
        [Input("contentSecurityPolicy")]
        public string? ContentSecurityPolicy { get; set; }

        [Input("contentSecurityPolicyReportOnly")]
        public string? ContentSecurityPolicyReportOnly { get; set; }

        [Input("strictTransportSecurity")]
        public string? StrictTransportSecurity { get; set; }

        [Input("xContentTypeOptions")]
        public string? XContentTypeOptions { get; set; }

        [Input("xFrameOptions")]
        public string? XFrameOptions { get; set; }

        [Input("xRobotsTag")]
        public string? XRobotsTag { get; set; }

        [Input("xXssProtection")]
        public string? XXssProtection { get; set; }

        public GetRealmSecurityDefensesHeadersArgs()
        {
        }
    }

    public sealed class GetRealmSmtpServersArgs : Pulumi.InvokeArgs
    {
        [Input("auths")]
        private List<GetRealmSmtpServersAuthsArgs>? _auths;
        public List<GetRealmSmtpServersAuthsArgs> Auths
        {
            get => _auths ?? (_auths = new List<GetRealmSmtpServersAuthsArgs>());
            set => _auths = value;
        }

        [Input("envelopeFrom")]
        public string? EnvelopeFrom { get; set; }

        [Input("from")]
        public string? From { get; set; }

        [Input("fromDisplayName")]
        public string? FromDisplayName { get; set; }

        [Input("host")]
        public string? Host { get; set; }

        [Input("port")]
        public string? Port { get; set; }

        [Input("replyTo")]
        public string? ReplyTo { get; set; }

        [Input("replyToDisplayName")]
        public string? ReplyToDisplayName { get; set; }

        [Input("ssl")]
        public bool? Ssl { get; set; }

        [Input("starttls")]
        public bool? Starttls { get; set; }

        public GetRealmSmtpServersArgs()
        {
        }
    }

    public sealed class GetRealmSmtpServersAuthsArgs : Pulumi.InvokeArgs
    {
        [Input("password")]
        public string? Password { get; set; }

        [Input("username")]
        public string? Username { get; set; }

        public GetRealmSmtpServersAuthsArgs()
        {
        }
    }
    }

    namespace Outputs
    {

    [OutputType]
    public sealed class GetRealmInternationalizationsResult
    {
        public readonly string DefaultLocale;
        public readonly ImmutableArray<string> SupportedLocales;

        [OutputConstructor]
        private GetRealmInternationalizationsResult(
            string defaultLocale,
            ImmutableArray<string> supportedLocales)
        {
            DefaultLocale = defaultLocale;
            SupportedLocales = supportedLocales;
        }
    }

    [OutputType]
    public sealed class GetRealmSecurityDefensesBruteForceDetectionsResult
    {
        public readonly int FailureResetTimeSeconds;
        public readonly int MaxFailureWaitSeconds;
        public readonly int MaxLoginFailures;
        public readonly int MinimumQuickLoginWaitSeconds;
        public readonly bool PermanentLockout;
        public readonly int QuickLoginCheckMilliSeconds;
        public readonly int WaitIncrementSeconds;

        [OutputConstructor]
        private GetRealmSecurityDefensesBruteForceDetectionsResult(
            int failureResetTimeSeconds,
            int maxFailureWaitSeconds,
            int maxLoginFailures,
            int minimumQuickLoginWaitSeconds,
            bool permanentLockout,
            int quickLoginCheckMilliSeconds,
            int waitIncrementSeconds)
        {
            FailureResetTimeSeconds = failureResetTimeSeconds;
            MaxFailureWaitSeconds = maxFailureWaitSeconds;
            MaxLoginFailures = maxLoginFailures;
            MinimumQuickLoginWaitSeconds = minimumQuickLoginWaitSeconds;
            PermanentLockout = permanentLockout;
            QuickLoginCheckMilliSeconds = quickLoginCheckMilliSeconds;
            WaitIncrementSeconds = waitIncrementSeconds;
        }
    }

    [OutputType]
    public sealed class GetRealmSecurityDefensesHeadersResult
    {
        public readonly string ContentSecurityPolicy;
        public readonly string ContentSecurityPolicyReportOnly;
        public readonly string StrictTransportSecurity;
        public readonly string XContentTypeOptions;
        public readonly string XFrameOptions;
        public readonly string XRobotsTag;
        public readonly string XXssProtection;

        [OutputConstructor]
        private GetRealmSecurityDefensesHeadersResult(
            string contentSecurityPolicy,
            string contentSecurityPolicyReportOnly,
            string strictTransportSecurity,
            string xContentTypeOptions,
            string xFrameOptions,
            string xRobotsTag,
            string xXssProtection)
        {
            ContentSecurityPolicy = contentSecurityPolicy;
            ContentSecurityPolicyReportOnly = contentSecurityPolicyReportOnly;
            StrictTransportSecurity = strictTransportSecurity;
            XContentTypeOptions = xContentTypeOptions;
            XFrameOptions = xFrameOptions;
            XRobotsTag = xRobotsTag;
            XXssProtection = xXssProtection;
        }
    }

    [OutputType]
    public sealed class GetRealmSecurityDefensesResult
    {
        public readonly ImmutableArray<GetRealmSecurityDefensesBruteForceDetectionsResult> BruteForceDetections;
        public readonly ImmutableArray<GetRealmSecurityDefensesHeadersResult> Headers;

        [OutputConstructor]
        private GetRealmSecurityDefensesResult(
            ImmutableArray<GetRealmSecurityDefensesBruteForceDetectionsResult> bruteForceDetections,
            ImmutableArray<GetRealmSecurityDefensesHeadersResult> headers)
        {
            BruteForceDetections = bruteForceDetections;
            Headers = headers;
        }
    }

    [OutputType]
    public sealed class GetRealmSmtpServersAuthsResult
    {
        public readonly string Password;
        public readonly string Username;

        [OutputConstructor]
        private GetRealmSmtpServersAuthsResult(
            string password,
            string username)
        {
            Password = password;
            Username = username;
        }
    }

    [OutputType]
    public sealed class GetRealmSmtpServersResult
    {
        public readonly ImmutableArray<GetRealmSmtpServersAuthsResult> Auths;
        public readonly string EnvelopeFrom;
        public readonly string From;
        public readonly string FromDisplayName;
        public readonly string Host;
        public readonly string Port;
        public readonly string ReplyTo;
        public readonly string ReplyToDisplayName;
        public readonly bool Ssl;
        public readonly bool Starttls;

        [OutputConstructor]
        private GetRealmSmtpServersResult(
            ImmutableArray<GetRealmSmtpServersAuthsResult> auths,
            string envelopeFrom,
            string from,
            string fromDisplayName,
            string host,
            string port,
            string replyTo,
            string replyToDisplayName,
            bool ssl,
            bool starttls)
        {
            Auths = auths;
            EnvelopeFrom = envelopeFrom;
            From = from;
            FromDisplayName = fromDisplayName;
            Host = host;
            Port = port;
            ReplyTo = replyTo;
            ReplyToDisplayName = replyToDisplayName;
            Ssl = ssl;
            Starttls = starttls;
        }
    }
    }
}
