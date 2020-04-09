// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package openid

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## # openid.Client
//
// Allows for creating and managing Keycloak clients that use the OpenID Connect protocol.
//
// Clients are entities that can use Keycloak for user authentication. Typically,
// clients are applications that redirect users to Keycloak for authentication
// in order to take advantage of Keycloak's user sessions for SSO.
//
// ### Argument Reference
//
// The following arguments are supported:
//
// - `realmId` - (Required) The realm this client is attached to.
// - `clientId` - (Required) The unique ID of this client, referenced in the URI during authentication and in issued tokens.
// - `name` - (Optional) The display name of this client in the GUI.
// - `enabled` - (Optional) When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
// - `description` - (Optional) The description of this client in the GUI.
// - `accessType` - (Required) Specifies the type of client, which can be one of the following:
//     - `CONFIDENTIAL` - Used for server-side clients that require both client ID and secret when authenticating.
//       This client should be used for applications using the Authorization Code or Client Credentials grant flows.
//     - `PUBLIC` - Used for browser-only applications that do not require a client secret, and instead rely only on authorized redirect
//       URIs for security. This client should be used for applications using the Implicit grant flow.
//     - `BEARER-ONLY` - Used for services that never initiate a login. This client will only allow bearer token requests.
// - `clientSecret` - (Optional) The secret for clients with an `accessType` of `CONFIDENTIAL` or `BEARER-ONLY`. This value is sensitive and
// should be treated with the same care as a password. If omitted, Keycloak will generate a GUID for this attribute.
// - `standardFlowEnabled` - (Optional) When `true`, the OAuth2 Authorization Code Grant will be enabled for this client. Defaults to `false`.
// - `implicitFlowEnabled` - (Optional) When `true`, the OAuth2 Implicit Grant will be enabled for this client. Defaults to `false`.
// - `directAccessGrantsEnabled` - (Optional) When `true`, the OAuth2 Resource Owner Password Grant will be enabled for this client. Defaults to `false`.
// - `serviceAccountsEnabled` - (Optional) When `true`, the OAuth2 Client Credentials grant will be enabled for this client. Defaults to `false`.
// - `validRedirectUris` - (Optional) A list of valid URIs a browser is permitted to redirect to after a successful login or logout. Simple
// wildcards in the form of an asterisk can be used here. This attribute must be set if either `standardFlowEnabled` or `implicitFlowEnabled`
// is set to `true`.
// - `webOrigins` - (Optional) A list of allowed CORS origins. `+` can be used to permit all valid redirect URIs, and `*` can be used to permit all origins.
// - `adminUrl` - (Optional) URL to the admin interface of the client.
// - `baseUrl` - (Optional) Default URL to use when the auth server needs to redirect or link back to the client.
// - `pkceCodeChallengeMethod` - (Optional) The challenge method to use for Proof Key for Code Exchange. Can be either `plain` or `S256` or set to empty value ``.
// - `fullScopeAllowed` - (Optional) - Allow to include all roles mappings in the access token.
//
// ### Attributes Reference
//
// In addition to the arguments listed above, the following computed attributes are exported:
//
// - `serviceAccountUserId` - When service accounts are enabled for this client, this attribute is the unique ID for the Keycloak user that represents this service account.
//
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/keycloak_openid_client.html.markdown.
type Client struct {
	pulumi.CustomResourceState

	AccessTokenLifespan                 pulumi.StringPtrOutput       `pulumi:"accessTokenLifespan"`
	AccessType                          pulumi.StringOutput          `pulumi:"accessType"`
	AdminUrl                            pulumi.StringPtrOutput       `pulumi:"adminUrl"`
	Authorization                       ClientAuthorizationPtrOutput `pulumi:"authorization"`
	BaseUrl                             pulumi.StringPtrOutput       `pulumi:"baseUrl"`
	ClientId                            pulumi.StringOutput          `pulumi:"clientId"`
	ClientSecret                        pulumi.StringOutput          `pulumi:"clientSecret"`
	ConsentRequired                     pulumi.BoolPtrOutput         `pulumi:"consentRequired"`
	Description                         pulumi.StringPtrOutput       `pulumi:"description"`
	DirectAccessGrantsEnabled           pulumi.BoolPtrOutput         `pulumi:"directAccessGrantsEnabled"`
	Enabled                             pulumi.BoolPtrOutput         `pulumi:"enabled"`
	ExcludeSessionStateFromAuthResponse pulumi.BoolPtrOutput         `pulumi:"excludeSessionStateFromAuthResponse"`
	FullScopeAllowed                    pulumi.BoolPtrOutput         `pulumi:"fullScopeAllowed"`
	ImplicitFlowEnabled                 pulumi.BoolPtrOutput         `pulumi:"implicitFlowEnabled"`
	Name                                pulumi.StringOutput          `pulumi:"name"`
	PkceCodeChallengeMethod             pulumi.StringPtrOutput       `pulumi:"pkceCodeChallengeMethod"`
	RealmId                             pulumi.StringOutput          `pulumi:"realmId"`
	ResourceServerId                    pulumi.StringOutput          `pulumi:"resourceServerId"`
	ServiceAccountUserId                pulumi.StringOutput          `pulumi:"serviceAccountUserId"`
	ServiceAccountsEnabled              pulumi.BoolPtrOutput         `pulumi:"serviceAccountsEnabled"`
	StandardFlowEnabled                 pulumi.BoolPtrOutput         `pulumi:"standardFlowEnabled"`
	ValidRedirectUris                   pulumi.StringArrayOutput     `pulumi:"validRedirectUris"`
	WebOrigins                          pulumi.StringArrayOutput     `pulumi:"webOrigins"`
}

// NewClient registers a new resource with the given unique name, arguments, and options.
func NewClient(ctx *pulumi.Context,
	name string, args *ClientArgs, opts ...pulumi.ResourceOption) (*Client, error) {
	if args == nil || args.AccessType == nil {
		return nil, errors.New("missing required argument 'AccessType'")
	}
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.RealmId == nil {
		return nil, errors.New("missing required argument 'RealmId'")
	}
	if args == nil {
		args = &ClientArgs{}
	}
	var resource Client
	err := ctx.RegisterResource("keycloak:openid/client:Client", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClient gets an existing Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClientState, opts ...pulumi.ResourceOption) (*Client, error) {
	var resource Client
	err := ctx.ReadResource("keycloak:openid/client:Client", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Client resources.
type clientState struct {
	AccessTokenLifespan                 *string              `pulumi:"accessTokenLifespan"`
	AccessType                          *string              `pulumi:"accessType"`
	AdminUrl                            *string              `pulumi:"adminUrl"`
	Authorization                       *ClientAuthorization `pulumi:"authorization"`
	BaseUrl                             *string              `pulumi:"baseUrl"`
	ClientId                            *string              `pulumi:"clientId"`
	ClientSecret                        *string              `pulumi:"clientSecret"`
	ConsentRequired                     *bool                `pulumi:"consentRequired"`
	Description                         *string              `pulumi:"description"`
	DirectAccessGrantsEnabled           *bool                `pulumi:"directAccessGrantsEnabled"`
	Enabled                             *bool                `pulumi:"enabled"`
	ExcludeSessionStateFromAuthResponse *bool                `pulumi:"excludeSessionStateFromAuthResponse"`
	FullScopeAllowed                    *bool                `pulumi:"fullScopeAllowed"`
	ImplicitFlowEnabled                 *bool                `pulumi:"implicitFlowEnabled"`
	Name                                *string              `pulumi:"name"`
	PkceCodeChallengeMethod             *string              `pulumi:"pkceCodeChallengeMethod"`
	RealmId                             *string              `pulumi:"realmId"`
	ResourceServerId                    *string              `pulumi:"resourceServerId"`
	ServiceAccountUserId                *string              `pulumi:"serviceAccountUserId"`
	ServiceAccountsEnabled              *bool                `pulumi:"serviceAccountsEnabled"`
	StandardFlowEnabled                 *bool                `pulumi:"standardFlowEnabled"`
	ValidRedirectUris                   []string             `pulumi:"validRedirectUris"`
	WebOrigins                          []string             `pulumi:"webOrigins"`
}

type ClientState struct {
	AccessTokenLifespan                 pulumi.StringPtrInput
	AccessType                          pulumi.StringPtrInput
	AdminUrl                            pulumi.StringPtrInput
	Authorization                       ClientAuthorizationPtrInput
	BaseUrl                             pulumi.StringPtrInput
	ClientId                            pulumi.StringPtrInput
	ClientSecret                        pulumi.StringPtrInput
	ConsentRequired                     pulumi.BoolPtrInput
	Description                         pulumi.StringPtrInput
	DirectAccessGrantsEnabled           pulumi.BoolPtrInput
	Enabled                             pulumi.BoolPtrInput
	ExcludeSessionStateFromAuthResponse pulumi.BoolPtrInput
	FullScopeAllowed                    pulumi.BoolPtrInput
	ImplicitFlowEnabled                 pulumi.BoolPtrInput
	Name                                pulumi.StringPtrInput
	PkceCodeChallengeMethod             pulumi.StringPtrInput
	RealmId                             pulumi.StringPtrInput
	ResourceServerId                    pulumi.StringPtrInput
	ServiceAccountUserId                pulumi.StringPtrInput
	ServiceAccountsEnabled              pulumi.BoolPtrInput
	StandardFlowEnabled                 pulumi.BoolPtrInput
	ValidRedirectUris                   pulumi.StringArrayInput
	WebOrigins                          pulumi.StringArrayInput
}

func (ClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*clientState)(nil)).Elem()
}

type clientArgs struct {
	AccessTokenLifespan                 *string              `pulumi:"accessTokenLifespan"`
	AccessType                          string               `pulumi:"accessType"`
	AdminUrl                            *string              `pulumi:"adminUrl"`
	Authorization                       *ClientAuthorization `pulumi:"authorization"`
	BaseUrl                             *string              `pulumi:"baseUrl"`
	ClientId                            string               `pulumi:"clientId"`
	ClientSecret                        *string              `pulumi:"clientSecret"`
	ConsentRequired                     *bool                `pulumi:"consentRequired"`
	Description                         *string              `pulumi:"description"`
	DirectAccessGrantsEnabled           *bool                `pulumi:"directAccessGrantsEnabled"`
	Enabled                             *bool                `pulumi:"enabled"`
	ExcludeSessionStateFromAuthResponse *bool                `pulumi:"excludeSessionStateFromAuthResponse"`
	FullScopeAllowed                    *bool                `pulumi:"fullScopeAllowed"`
	ImplicitFlowEnabled                 *bool                `pulumi:"implicitFlowEnabled"`
	Name                                *string              `pulumi:"name"`
	PkceCodeChallengeMethod             *string              `pulumi:"pkceCodeChallengeMethod"`
	RealmId                             string               `pulumi:"realmId"`
	ServiceAccountsEnabled              *bool                `pulumi:"serviceAccountsEnabled"`
	StandardFlowEnabled                 *bool                `pulumi:"standardFlowEnabled"`
	ValidRedirectUris                   []string             `pulumi:"validRedirectUris"`
	WebOrigins                          []string             `pulumi:"webOrigins"`
}

// The set of arguments for constructing a Client resource.
type ClientArgs struct {
	AccessTokenLifespan                 pulumi.StringPtrInput
	AccessType                          pulumi.StringInput
	AdminUrl                            pulumi.StringPtrInput
	Authorization                       ClientAuthorizationPtrInput
	BaseUrl                             pulumi.StringPtrInput
	ClientId                            pulumi.StringInput
	ClientSecret                        pulumi.StringPtrInput
	ConsentRequired                     pulumi.BoolPtrInput
	Description                         pulumi.StringPtrInput
	DirectAccessGrantsEnabled           pulumi.BoolPtrInput
	Enabled                             pulumi.BoolPtrInput
	ExcludeSessionStateFromAuthResponse pulumi.BoolPtrInput
	FullScopeAllowed                    pulumi.BoolPtrInput
	ImplicitFlowEnabled                 pulumi.BoolPtrInput
	Name                                pulumi.StringPtrInput
	PkceCodeChallengeMethod             pulumi.StringPtrInput
	RealmId                             pulumi.StringInput
	ServiceAccountsEnabled              pulumi.BoolPtrInput
	StandardFlowEnabled                 pulumi.BoolPtrInput
	ValidRedirectUris                   pulumi.StringArrayInput
	WebOrigins                          pulumi.StringArrayInput
}

func (ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clientArgs)(nil)).Elem()
}
