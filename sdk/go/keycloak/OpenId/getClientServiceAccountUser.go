// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package OpenId

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func GetClientServiceAccountUser(ctx *pulumi.Context, args *GetClientServiceAccountUserArgs, opts ...pulumi.InvokeOption) (*GetClientServiceAccountUserResult, error) {
	var rv GetClientServiceAccountUserResult
	err := ctx.Invoke("keycloak:OpenId/getClientServiceAccountUser:getClientServiceAccountUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getClientServiceAccountUser.
type GetClientServiceAccountUserArgs struct {
	ClientId string `pulumi:"clientId"`
	RealmId string `pulumi:"realmId"`
}


// A collection of values returned by getClientServiceAccountUser.
type GetClientServiceAccountUserResult struct {
	Attributes map[string]interface{} `pulumi:"attributes"`
	ClientId string `pulumi:"clientId"`
	Email string `pulumi:"email"`
	Enabled bool `pulumi:"enabled"`
	FederatedIdentities []GetClientServiceAccountUserFederatedIdentity `pulumi:"federatedIdentities"`
	FirstName string `pulumi:"firstName"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	LastName string `pulumi:"lastName"`
	RealmId string `pulumi:"realmId"`
	Username string `pulumi:"username"`
}

