// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package openid

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ClientAuthenticationFlowBindingOverrides struct {
	BrowserId     *string `pulumi:"browserId"`
	DirectGrantId *string `pulumi:"directGrantId"`
}

// ClientAuthenticationFlowBindingOverridesInput is an input type that accepts ClientAuthenticationFlowBindingOverridesArgs and ClientAuthenticationFlowBindingOverridesOutput values.
// You can construct a concrete instance of `ClientAuthenticationFlowBindingOverridesInput` via:
//
// 		 ClientAuthenticationFlowBindingOverridesArgs{...}
//
type ClientAuthenticationFlowBindingOverridesInput interface {
	pulumi.Input

	ToClientAuthenticationFlowBindingOverridesOutput() ClientAuthenticationFlowBindingOverridesOutput
	ToClientAuthenticationFlowBindingOverridesOutputWithContext(context.Context) ClientAuthenticationFlowBindingOverridesOutput
}

type ClientAuthenticationFlowBindingOverridesArgs struct {
	BrowserId     pulumi.StringPtrInput `pulumi:"browserId"`
	DirectGrantId pulumi.StringPtrInput `pulumi:"directGrantId"`
}

func (ClientAuthenticationFlowBindingOverridesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (i ClientAuthenticationFlowBindingOverridesArgs) ToClientAuthenticationFlowBindingOverridesOutput() ClientAuthenticationFlowBindingOverridesOutput {
	return i.ToClientAuthenticationFlowBindingOverridesOutputWithContext(context.Background())
}

func (i ClientAuthenticationFlowBindingOverridesArgs) ToClientAuthenticationFlowBindingOverridesOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthenticationFlowBindingOverridesOutput)
}

func (i ClientAuthenticationFlowBindingOverridesArgs) ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput {
	return i.ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(context.Background())
}

func (i ClientAuthenticationFlowBindingOverridesArgs) ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthenticationFlowBindingOverridesOutput).ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx)
}

// ClientAuthenticationFlowBindingOverridesPtrInput is an input type that accepts ClientAuthenticationFlowBindingOverridesArgs, ClientAuthenticationFlowBindingOverridesPtr and ClientAuthenticationFlowBindingOverridesPtrOutput values.
// You can construct a concrete instance of `ClientAuthenticationFlowBindingOverridesPtrInput` via:
//
// 		 ClientAuthenticationFlowBindingOverridesArgs{...}
//
//  or:
//
// 		 nil
//
type ClientAuthenticationFlowBindingOverridesPtrInput interface {
	pulumi.Input

	ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput
	ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput
}

type clientAuthenticationFlowBindingOverridesPtrType ClientAuthenticationFlowBindingOverridesArgs

func ClientAuthenticationFlowBindingOverridesPtr(v *ClientAuthenticationFlowBindingOverridesArgs) ClientAuthenticationFlowBindingOverridesPtrInput {
	return (*clientAuthenticationFlowBindingOverridesPtrType)(v)
}

func (*clientAuthenticationFlowBindingOverridesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (i *clientAuthenticationFlowBindingOverridesPtrType) ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput {
	return i.ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(context.Background())
}

func (i *clientAuthenticationFlowBindingOverridesPtrType) ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthenticationFlowBindingOverridesPtrOutput)
}

type ClientAuthenticationFlowBindingOverridesOutput struct{ *pulumi.OutputState }

func (ClientAuthenticationFlowBindingOverridesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (o ClientAuthenticationFlowBindingOverridesOutput) ToClientAuthenticationFlowBindingOverridesOutput() ClientAuthenticationFlowBindingOverridesOutput {
	return o
}

func (o ClientAuthenticationFlowBindingOverridesOutput) ToClientAuthenticationFlowBindingOverridesOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesOutput {
	return o
}

func (o ClientAuthenticationFlowBindingOverridesOutput) ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput {
	return o.ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(context.Background())
}

func (o ClientAuthenticationFlowBindingOverridesOutput) ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput {
	return o.ApplyT(func(v ClientAuthenticationFlowBindingOverrides) *ClientAuthenticationFlowBindingOverrides {
		return &v
	}).(ClientAuthenticationFlowBindingOverridesPtrOutput)
}
func (o ClientAuthenticationFlowBindingOverridesOutput) BrowserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientAuthenticationFlowBindingOverrides) *string { return v.BrowserId }).(pulumi.StringPtrOutput)
}

func (o ClientAuthenticationFlowBindingOverridesOutput) DirectGrantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientAuthenticationFlowBindingOverrides) *string { return v.DirectGrantId }).(pulumi.StringPtrOutput)
}

type ClientAuthenticationFlowBindingOverridesPtrOutput struct{ *pulumi.OutputState }

func (ClientAuthenticationFlowBindingOverridesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (o ClientAuthenticationFlowBindingOverridesPtrOutput) ToClientAuthenticationFlowBindingOverridesPtrOutput() ClientAuthenticationFlowBindingOverridesPtrOutput {
	return o
}

func (o ClientAuthenticationFlowBindingOverridesPtrOutput) ToClientAuthenticationFlowBindingOverridesPtrOutputWithContext(ctx context.Context) ClientAuthenticationFlowBindingOverridesPtrOutput {
	return o
}

func (o ClientAuthenticationFlowBindingOverridesPtrOutput) Elem() ClientAuthenticationFlowBindingOverridesOutput {
	return o.ApplyT(func(v *ClientAuthenticationFlowBindingOverrides) ClientAuthenticationFlowBindingOverrides { return *v }).(ClientAuthenticationFlowBindingOverridesOutput)
}

func (o ClientAuthenticationFlowBindingOverridesPtrOutput) BrowserId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthenticationFlowBindingOverrides) *string {
		if v == nil {
			return nil
		}
		return v.BrowserId
	}).(pulumi.StringPtrOutput)
}

func (o ClientAuthenticationFlowBindingOverridesPtrOutput) DirectGrantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthenticationFlowBindingOverrides) *string {
		if v == nil {
			return nil
		}
		return v.DirectGrantId
	}).(pulumi.StringPtrOutput)
}

type ClientAuthorization struct {
	AllowRemoteResourceManagement *bool  `pulumi:"allowRemoteResourceManagement"`
	KeepDefaults                  *bool  `pulumi:"keepDefaults"`
	PolicyEnforcementMode         string `pulumi:"policyEnforcementMode"`
}

// ClientAuthorizationInput is an input type that accepts ClientAuthorizationArgs and ClientAuthorizationOutput values.
// You can construct a concrete instance of `ClientAuthorizationInput` via:
//
// 		 ClientAuthorizationArgs{...}
//
type ClientAuthorizationInput interface {
	pulumi.Input

	ToClientAuthorizationOutput() ClientAuthorizationOutput
	ToClientAuthorizationOutputWithContext(context.Context) ClientAuthorizationOutput
}

type ClientAuthorizationArgs struct {
	AllowRemoteResourceManagement pulumi.BoolPtrInput `pulumi:"allowRemoteResourceManagement"`
	KeepDefaults                  pulumi.BoolPtrInput `pulumi:"keepDefaults"`
	PolicyEnforcementMode         pulumi.StringInput  `pulumi:"policyEnforcementMode"`
}

func (ClientAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthorization)(nil)).Elem()
}

func (i ClientAuthorizationArgs) ToClientAuthorizationOutput() ClientAuthorizationOutput {
	return i.ToClientAuthorizationOutputWithContext(context.Background())
}

func (i ClientAuthorizationArgs) ToClientAuthorizationOutputWithContext(ctx context.Context) ClientAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationOutput)
}

func (i ClientAuthorizationArgs) ToClientAuthorizationPtrOutput() ClientAuthorizationPtrOutput {
	return i.ToClientAuthorizationPtrOutputWithContext(context.Background())
}

func (i ClientAuthorizationArgs) ToClientAuthorizationPtrOutputWithContext(ctx context.Context) ClientAuthorizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationOutput).ToClientAuthorizationPtrOutputWithContext(ctx)
}

// ClientAuthorizationPtrInput is an input type that accepts ClientAuthorizationArgs, ClientAuthorizationPtr and ClientAuthorizationPtrOutput values.
// You can construct a concrete instance of `ClientAuthorizationPtrInput` via:
//
// 		 ClientAuthorizationArgs{...}
//
//  or:
//
// 		 nil
//
type ClientAuthorizationPtrInput interface {
	pulumi.Input

	ToClientAuthorizationPtrOutput() ClientAuthorizationPtrOutput
	ToClientAuthorizationPtrOutputWithContext(context.Context) ClientAuthorizationPtrOutput
}

type clientAuthorizationPtrType ClientAuthorizationArgs

func ClientAuthorizationPtr(v *ClientAuthorizationArgs) ClientAuthorizationPtrInput {
	return (*clientAuthorizationPtrType)(v)
}

func (*clientAuthorizationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorization)(nil)).Elem()
}

func (i *clientAuthorizationPtrType) ToClientAuthorizationPtrOutput() ClientAuthorizationPtrOutput {
	return i.ToClientAuthorizationPtrOutputWithContext(context.Background())
}

func (i *clientAuthorizationPtrType) ToClientAuthorizationPtrOutputWithContext(ctx context.Context) ClientAuthorizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientAuthorizationPtrOutput)
}

type ClientAuthorizationOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientAuthorization)(nil)).Elem()
}

func (o ClientAuthorizationOutput) ToClientAuthorizationOutput() ClientAuthorizationOutput {
	return o
}

func (o ClientAuthorizationOutput) ToClientAuthorizationOutputWithContext(ctx context.Context) ClientAuthorizationOutput {
	return o
}

func (o ClientAuthorizationOutput) ToClientAuthorizationPtrOutput() ClientAuthorizationPtrOutput {
	return o.ToClientAuthorizationPtrOutputWithContext(context.Background())
}

func (o ClientAuthorizationOutput) ToClientAuthorizationPtrOutputWithContext(ctx context.Context) ClientAuthorizationPtrOutput {
	return o.ApplyT(func(v ClientAuthorization) *ClientAuthorization {
		return &v
	}).(ClientAuthorizationPtrOutput)
}
func (o ClientAuthorizationOutput) AllowRemoteResourceManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClientAuthorization) *bool { return v.AllowRemoteResourceManagement }).(pulumi.BoolPtrOutput)
}

func (o ClientAuthorizationOutput) KeepDefaults() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClientAuthorization) *bool { return v.KeepDefaults }).(pulumi.BoolPtrOutput)
}

func (o ClientAuthorizationOutput) PolicyEnforcementMode() pulumi.StringOutput {
	return o.ApplyT(func(v ClientAuthorization) string { return v.PolicyEnforcementMode }).(pulumi.StringOutput)
}

type ClientAuthorizationPtrOutput struct{ *pulumi.OutputState }

func (ClientAuthorizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientAuthorization)(nil)).Elem()
}

func (o ClientAuthorizationPtrOutput) ToClientAuthorizationPtrOutput() ClientAuthorizationPtrOutput {
	return o
}

func (o ClientAuthorizationPtrOutput) ToClientAuthorizationPtrOutputWithContext(ctx context.Context) ClientAuthorizationPtrOutput {
	return o
}

func (o ClientAuthorizationPtrOutput) Elem() ClientAuthorizationOutput {
	return o.ApplyT(func(v *ClientAuthorization) ClientAuthorization { return *v }).(ClientAuthorizationOutput)
}

func (o ClientAuthorizationPtrOutput) AllowRemoteResourceManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClientAuthorization) *bool {
		if v == nil {
			return nil
		}
		return v.AllowRemoteResourceManagement
	}).(pulumi.BoolPtrOutput)
}

func (o ClientAuthorizationPtrOutput) KeepDefaults() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClientAuthorization) *bool {
		if v == nil {
			return nil
		}
		return v.KeepDefaults
	}).(pulumi.BoolPtrOutput)
}

func (o ClientAuthorizationPtrOutput) PolicyEnforcementMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientAuthorization) *string {
		if v == nil {
			return nil
		}
		return &v.PolicyEnforcementMode
	}).(pulumi.StringPtrOutput)
}

type ClientGroupPolicyGroup struct {
	ExtendChildren bool   `pulumi:"extendChildren"`
	Id             string `pulumi:"id"`
	Path           string `pulumi:"path"`
}

// ClientGroupPolicyGroupInput is an input type that accepts ClientGroupPolicyGroupArgs and ClientGroupPolicyGroupOutput values.
// You can construct a concrete instance of `ClientGroupPolicyGroupInput` via:
//
// 		 ClientGroupPolicyGroupArgs{...}
//
type ClientGroupPolicyGroupInput interface {
	pulumi.Input

	ToClientGroupPolicyGroupOutput() ClientGroupPolicyGroupOutput
	ToClientGroupPolicyGroupOutputWithContext(context.Context) ClientGroupPolicyGroupOutput
}

type ClientGroupPolicyGroupArgs struct {
	ExtendChildren pulumi.BoolInput   `pulumi:"extendChildren"`
	Id             pulumi.StringInput `pulumi:"id"`
	Path           pulumi.StringInput `pulumi:"path"`
}

func (ClientGroupPolicyGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientGroupPolicyGroup)(nil)).Elem()
}

func (i ClientGroupPolicyGroupArgs) ToClientGroupPolicyGroupOutput() ClientGroupPolicyGroupOutput {
	return i.ToClientGroupPolicyGroupOutputWithContext(context.Background())
}

func (i ClientGroupPolicyGroupArgs) ToClientGroupPolicyGroupOutputWithContext(ctx context.Context) ClientGroupPolicyGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGroupPolicyGroupOutput)
}

// ClientGroupPolicyGroupArrayInput is an input type that accepts ClientGroupPolicyGroupArray and ClientGroupPolicyGroupArrayOutput values.
// You can construct a concrete instance of `ClientGroupPolicyGroupArrayInput` via:
//
// 		 ClientGroupPolicyGroupArray{ ClientGroupPolicyGroupArgs{...} }
//
type ClientGroupPolicyGroupArrayInput interface {
	pulumi.Input

	ToClientGroupPolicyGroupArrayOutput() ClientGroupPolicyGroupArrayOutput
	ToClientGroupPolicyGroupArrayOutputWithContext(context.Context) ClientGroupPolicyGroupArrayOutput
}

type ClientGroupPolicyGroupArray []ClientGroupPolicyGroupInput

func (ClientGroupPolicyGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientGroupPolicyGroup)(nil)).Elem()
}

func (i ClientGroupPolicyGroupArray) ToClientGroupPolicyGroupArrayOutput() ClientGroupPolicyGroupArrayOutput {
	return i.ToClientGroupPolicyGroupArrayOutputWithContext(context.Background())
}

func (i ClientGroupPolicyGroupArray) ToClientGroupPolicyGroupArrayOutputWithContext(ctx context.Context) ClientGroupPolicyGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientGroupPolicyGroupArrayOutput)
}

type ClientGroupPolicyGroupOutput struct{ *pulumi.OutputState }

func (ClientGroupPolicyGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientGroupPolicyGroup)(nil)).Elem()
}

func (o ClientGroupPolicyGroupOutput) ToClientGroupPolicyGroupOutput() ClientGroupPolicyGroupOutput {
	return o
}

func (o ClientGroupPolicyGroupOutput) ToClientGroupPolicyGroupOutputWithContext(ctx context.Context) ClientGroupPolicyGroupOutput {
	return o
}

func (o ClientGroupPolicyGroupOutput) ExtendChildren() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientGroupPolicyGroup) bool { return v.ExtendChildren }).(pulumi.BoolOutput)
}

func (o ClientGroupPolicyGroupOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ClientGroupPolicyGroup) string { return v.Id }).(pulumi.StringOutput)
}

func (o ClientGroupPolicyGroupOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v ClientGroupPolicyGroup) string { return v.Path }).(pulumi.StringOutput)
}

type ClientGroupPolicyGroupArrayOutput struct{ *pulumi.OutputState }

func (ClientGroupPolicyGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientGroupPolicyGroup)(nil)).Elem()
}

func (o ClientGroupPolicyGroupArrayOutput) ToClientGroupPolicyGroupArrayOutput() ClientGroupPolicyGroupArrayOutput {
	return o
}

func (o ClientGroupPolicyGroupArrayOutput) ToClientGroupPolicyGroupArrayOutputWithContext(ctx context.Context) ClientGroupPolicyGroupArrayOutput {
	return o
}

func (o ClientGroupPolicyGroupArrayOutput) Index(i pulumi.IntInput) ClientGroupPolicyGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientGroupPolicyGroup {
		return vs[0].([]ClientGroupPolicyGroup)[vs[1].(int)]
	}).(ClientGroupPolicyGroupOutput)
}

type ClientRolePolicyRole struct {
	Id       string `pulumi:"id"`
	Required bool   `pulumi:"required"`
}

// ClientRolePolicyRoleInput is an input type that accepts ClientRolePolicyRoleArgs and ClientRolePolicyRoleOutput values.
// You can construct a concrete instance of `ClientRolePolicyRoleInput` via:
//
// 		 ClientRolePolicyRoleArgs{...}
//
type ClientRolePolicyRoleInput interface {
	pulumi.Input

	ToClientRolePolicyRoleOutput() ClientRolePolicyRoleOutput
	ToClientRolePolicyRoleOutputWithContext(context.Context) ClientRolePolicyRoleOutput
}

type ClientRolePolicyRoleArgs struct {
	Id       pulumi.StringInput `pulumi:"id"`
	Required pulumi.BoolInput   `pulumi:"required"`
}

func (ClientRolePolicyRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRolePolicyRole)(nil)).Elem()
}

func (i ClientRolePolicyRoleArgs) ToClientRolePolicyRoleOutput() ClientRolePolicyRoleOutput {
	return i.ToClientRolePolicyRoleOutputWithContext(context.Background())
}

func (i ClientRolePolicyRoleArgs) ToClientRolePolicyRoleOutputWithContext(ctx context.Context) ClientRolePolicyRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRolePolicyRoleOutput)
}

// ClientRolePolicyRoleArrayInput is an input type that accepts ClientRolePolicyRoleArray and ClientRolePolicyRoleArrayOutput values.
// You can construct a concrete instance of `ClientRolePolicyRoleArrayInput` via:
//
// 		 ClientRolePolicyRoleArray{ ClientRolePolicyRoleArgs{...} }
//
type ClientRolePolicyRoleArrayInput interface {
	pulumi.Input

	ToClientRolePolicyRoleArrayOutput() ClientRolePolicyRoleArrayOutput
	ToClientRolePolicyRoleArrayOutputWithContext(context.Context) ClientRolePolicyRoleArrayOutput
}

type ClientRolePolicyRoleArray []ClientRolePolicyRoleInput

func (ClientRolePolicyRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientRolePolicyRole)(nil)).Elem()
}

func (i ClientRolePolicyRoleArray) ToClientRolePolicyRoleArrayOutput() ClientRolePolicyRoleArrayOutput {
	return i.ToClientRolePolicyRoleArrayOutputWithContext(context.Background())
}

func (i ClientRolePolicyRoleArray) ToClientRolePolicyRoleArrayOutputWithContext(ctx context.Context) ClientRolePolicyRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientRolePolicyRoleArrayOutput)
}

type ClientRolePolicyRoleOutput struct{ *pulumi.OutputState }

func (ClientRolePolicyRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientRolePolicyRole)(nil)).Elem()
}

func (o ClientRolePolicyRoleOutput) ToClientRolePolicyRoleOutput() ClientRolePolicyRoleOutput {
	return o
}

func (o ClientRolePolicyRoleOutput) ToClientRolePolicyRoleOutputWithContext(ctx context.Context) ClientRolePolicyRoleOutput {
	return o
}

func (o ClientRolePolicyRoleOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ClientRolePolicyRole) string { return v.Id }).(pulumi.StringOutput)
}

func (o ClientRolePolicyRoleOutput) Required() pulumi.BoolOutput {
	return o.ApplyT(func(v ClientRolePolicyRole) bool { return v.Required }).(pulumi.BoolOutput)
}

type ClientRolePolicyRoleArrayOutput struct{ *pulumi.OutputState }

func (ClientRolePolicyRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClientRolePolicyRole)(nil)).Elem()
}

func (o ClientRolePolicyRoleArrayOutput) ToClientRolePolicyRoleArrayOutput() ClientRolePolicyRoleArrayOutput {
	return o
}

func (o ClientRolePolicyRoleArrayOutput) ToClientRolePolicyRoleArrayOutputWithContext(ctx context.Context) ClientRolePolicyRoleArrayOutput {
	return o
}

func (o ClientRolePolicyRoleArrayOutput) Index(i pulumi.IntInput) ClientRolePolicyRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClientRolePolicyRole {
		return vs[0].([]ClientRolePolicyRole)[vs[1].(int)]
	}).(ClientRolePolicyRoleOutput)
}

type GetClientAuthenticationFlowBindingOverrides struct {
	BrowserId     string `pulumi:"browserId"`
	DirectGrantId string `pulumi:"directGrantId"`
}

// GetClientAuthenticationFlowBindingOverridesInput is an input type that accepts GetClientAuthenticationFlowBindingOverridesArgs and GetClientAuthenticationFlowBindingOverridesOutput values.
// You can construct a concrete instance of `GetClientAuthenticationFlowBindingOverridesInput` via:
//
// 		 GetClientAuthenticationFlowBindingOverridesArgs{...}
//
type GetClientAuthenticationFlowBindingOverridesInput interface {
	pulumi.Input

	ToGetClientAuthenticationFlowBindingOverridesOutput() GetClientAuthenticationFlowBindingOverridesOutput
	ToGetClientAuthenticationFlowBindingOverridesOutputWithContext(context.Context) GetClientAuthenticationFlowBindingOverridesOutput
}

type GetClientAuthenticationFlowBindingOverridesArgs struct {
	BrowserId     pulumi.StringInput `pulumi:"browserId"`
	DirectGrantId pulumi.StringInput `pulumi:"directGrantId"`
}

func (GetClientAuthenticationFlowBindingOverridesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (i GetClientAuthenticationFlowBindingOverridesArgs) ToGetClientAuthenticationFlowBindingOverridesOutput() GetClientAuthenticationFlowBindingOverridesOutput {
	return i.ToGetClientAuthenticationFlowBindingOverridesOutputWithContext(context.Background())
}

func (i GetClientAuthenticationFlowBindingOverridesArgs) ToGetClientAuthenticationFlowBindingOverridesOutputWithContext(ctx context.Context) GetClientAuthenticationFlowBindingOverridesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClientAuthenticationFlowBindingOverridesOutput)
}

type GetClientAuthenticationFlowBindingOverridesOutput struct{ *pulumi.OutputState }

func (GetClientAuthenticationFlowBindingOverridesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientAuthenticationFlowBindingOverrides)(nil)).Elem()
}

func (o GetClientAuthenticationFlowBindingOverridesOutput) ToGetClientAuthenticationFlowBindingOverridesOutput() GetClientAuthenticationFlowBindingOverridesOutput {
	return o
}

func (o GetClientAuthenticationFlowBindingOverridesOutput) ToGetClientAuthenticationFlowBindingOverridesOutputWithContext(ctx context.Context) GetClientAuthenticationFlowBindingOverridesOutput {
	return o
}

func (o GetClientAuthenticationFlowBindingOverridesOutput) BrowserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthenticationFlowBindingOverrides) string { return v.BrowserId }).(pulumi.StringOutput)
}

func (o GetClientAuthenticationFlowBindingOverridesOutput) DirectGrantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthenticationFlowBindingOverrides) string { return v.DirectGrantId }).(pulumi.StringOutput)
}

type GetClientAuthorization struct {
	AllowRemoteResourceManagement bool   `pulumi:"allowRemoteResourceManagement"`
	PolicyEnforcementMode         string `pulumi:"policyEnforcementMode"`
}

// GetClientAuthorizationInput is an input type that accepts GetClientAuthorizationArgs and GetClientAuthorizationOutput values.
// You can construct a concrete instance of `GetClientAuthorizationInput` via:
//
// 		 GetClientAuthorizationArgs{...}
//
type GetClientAuthorizationInput interface {
	pulumi.Input

	ToGetClientAuthorizationOutput() GetClientAuthorizationOutput
	ToGetClientAuthorizationOutputWithContext(context.Context) GetClientAuthorizationOutput
}

type GetClientAuthorizationArgs struct {
	AllowRemoteResourceManagement pulumi.BoolInput   `pulumi:"allowRemoteResourceManagement"`
	PolicyEnforcementMode         pulumi.StringInput `pulumi:"policyEnforcementMode"`
}

func (GetClientAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientAuthorization)(nil)).Elem()
}

func (i GetClientAuthorizationArgs) ToGetClientAuthorizationOutput() GetClientAuthorizationOutput {
	return i.ToGetClientAuthorizationOutputWithContext(context.Background())
}

func (i GetClientAuthorizationArgs) ToGetClientAuthorizationOutputWithContext(ctx context.Context) GetClientAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClientAuthorizationOutput)
}

type GetClientAuthorizationOutput struct{ *pulumi.OutputState }

func (GetClientAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientAuthorization)(nil)).Elem()
}

func (o GetClientAuthorizationOutput) ToGetClientAuthorizationOutput() GetClientAuthorizationOutput {
	return o
}

func (o GetClientAuthorizationOutput) ToGetClientAuthorizationOutputWithContext(ctx context.Context) GetClientAuthorizationOutput {
	return o
}

func (o GetClientAuthorizationOutput) AllowRemoteResourceManagement() pulumi.BoolOutput {
	return o.ApplyT(func(v GetClientAuthorization) bool { return v.AllowRemoteResourceManagement }).(pulumi.BoolOutput)
}

func (o GetClientAuthorizationOutput) PolicyEnforcementMode() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientAuthorization) string { return v.PolicyEnforcementMode }).(pulumi.StringOutput)
}

type GetClientServiceAccountUserFederatedIdentity struct {
	IdentityProvider string `pulumi:"identityProvider"`
	UserId           string `pulumi:"userId"`
	UserName         string `pulumi:"userName"`
}

// GetClientServiceAccountUserFederatedIdentityInput is an input type that accepts GetClientServiceAccountUserFederatedIdentityArgs and GetClientServiceAccountUserFederatedIdentityOutput values.
// You can construct a concrete instance of `GetClientServiceAccountUserFederatedIdentityInput` via:
//
// 		 GetClientServiceAccountUserFederatedIdentityArgs{...}
//
type GetClientServiceAccountUserFederatedIdentityInput interface {
	pulumi.Input

	ToGetClientServiceAccountUserFederatedIdentityOutput() GetClientServiceAccountUserFederatedIdentityOutput
	ToGetClientServiceAccountUserFederatedIdentityOutputWithContext(context.Context) GetClientServiceAccountUserFederatedIdentityOutput
}

type GetClientServiceAccountUserFederatedIdentityArgs struct {
	IdentityProvider pulumi.StringInput `pulumi:"identityProvider"`
	UserId           pulumi.StringInput `pulumi:"userId"`
	UserName         pulumi.StringInput `pulumi:"userName"`
}

func (GetClientServiceAccountUserFederatedIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientServiceAccountUserFederatedIdentity)(nil)).Elem()
}

func (i GetClientServiceAccountUserFederatedIdentityArgs) ToGetClientServiceAccountUserFederatedIdentityOutput() GetClientServiceAccountUserFederatedIdentityOutput {
	return i.ToGetClientServiceAccountUserFederatedIdentityOutputWithContext(context.Background())
}

func (i GetClientServiceAccountUserFederatedIdentityArgs) ToGetClientServiceAccountUserFederatedIdentityOutputWithContext(ctx context.Context) GetClientServiceAccountUserFederatedIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClientServiceAccountUserFederatedIdentityOutput)
}

// GetClientServiceAccountUserFederatedIdentityArrayInput is an input type that accepts GetClientServiceAccountUserFederatedIdentityArray and GetClientServiceAccountUserFederatedIdentityArrayOutput values.
// You can construct a concrete instance of `GetClientServiceAccountUserFederatedIdentityArrayInput` via:
//
// 		 GetClientServiceAccountUserFederatedIdentityArray{ GetClientServiceAccountUserFederatedIdentityArgs{...} }
//
type GetClientServiceAccountUserFederatedIdentityArrayInput interface {
	pulumi.Input

	ToGetClientServiceAccountUserFederatedIdentityArrayOutput() GetClientServiceAccountUserFederatedIdentityArrayOutput
	ToGetClientServiceAccountUserFederatedIdentityArrayOutputWithContext(context.Context) GetClientServiceAccountUserFederatedIdentityArrayOutput
}

type GetClientServiceAccountUserFederatedIdentityArray []GetClientServiceAccountUserFederatedIdentityInput

func (GetClientServiceAccountUserFederatedIdentityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClientServiceAccountUserFederatedIdentity)(nil)).Elem()
}

func (i GetClientServiceAccountUserFederatedIdentityArray) ToGetClientServiceAccountUserFederatedIdentityArrayOutput() GetClientServiceAccountUserFederatedIdentityArrayOutput {
	return i.ToGetClientServiceAccountUserFederatedIdentityArrayOutputWithContext(context.Background())
}

func (i GetClientServiceAccountUserFederatedIdentityArray) ToGetClientServiceAccountUserFederatedIdentityArrayOutputWithContext(ctx context.Context) GetClientServiceAccountUserFederatedIdentityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetClientServiceAccountUserFederatedIdentityArrayOutput)
}

type GetClientServiceAccountUserFederatedIdentityOutput struct{ *pulumi.OutputState }

func (GetClientServiceAccountUserFederatedIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientServiceAccountUserFederatedIdentity)(nil)).Elem()
}

func (o GetClientServiceAccountUserFederatedIdentityOutput) ToGetClientServiceAccountUserFederatedIdentityOutput() GetClientServiceAccountUserFederatedIdentityOutput {
	return o
}

func (o GetClientServiceAccountUserFederatedIdentityOutput) ToGetClientServiceAccountUserFederatedIdentityOutputWithContext(ctx context.Context) GetClientServiceAccountUserFederatedIdentityOutput {
	return o
}

func (o GetClientServiceAccountUserFederatedIdentityOutput) IdentityProvider() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserFederatedIdentity) string { return v.IdentityProvider }).(pulumi.StringOutput)
}

func (o GetClientServiceAccountUserFederatedIdentityOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserFederatedIdentity) string { return v.UserId }).(pulumi.StringOutput)
}

func (o GetClientServiceAccountUserFederatedIdentityOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientServiceAccountUserFederatedIdentity) string { return v.UserName }).(pulumi.StringOutput)
}

type GetClientServiceAccountUserFederatedIdentityArrayOutput struct{ *pulumi.OutputState }

func (GetClientServiceAccountUserFederatedIdentityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetClientServiceAccountUserFederatedIdentity)(nil)).Elem()
}

func (o GetClientServiceAccountUserFederatedIdentityArrayOutput) ToGetClientServiceAccountUserFederatedIdentityArrayOutput() GetClientServiceAccountUserFederatedIdentityArrayOutput {
	return o
}

func (o GetClientServiceAccountUserFederatedIdentityArrayOutput) ToGetClientServiceAccountUserFederatedIdentityArrayOutputWithContext(ctx context.Context) GetClientServiceAccountUserFederatedIdentityArrayOutput {
	return o
}

func (o GetClientServiceAccountUserFederatedIdentityArrayOutput) Index(i pulumi.IntInput) GetClientServiceAccountUserFederatedIdentityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetClientServiceAccountUserFederatedIdentity {
		return vs[0].([]GetClientServiceAccountUserFederatedIdentity)[vs[1].(int)]
	}).(GetClientServiceAccountUserFederatedIdentityOutput)
}

func init() {
	pulumi.RegisterOutputType(ClientAuthenticationFlowBindingOverridesOutput{})
	pulumi.RegisterOutputType(ClientAuthenticationFlowBindingOverridesPtrOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationOutput{})
	pulumi.RegisterOutputType(ClientAuthorizationPtrOutput{})
	pulumi.RegisterOutputType(ClientGroupPolicyGroupOutput{})
	pulumi.RegisterOutputType(ClientGroupPolicyGroupArrayOutput{})
	pulumi.RegisterOutputType(ClientRolePolicyRoleOutput{})
	pulumi.RegisterOutputType(ClientRolePolicyRoleArrayOutput{})
	pulumi.RegisterOutputType(GetClientAuthenticationFlowBindingOverridesOutput{})
	pulumi.RegisterOutputType(GetClientAuthorizationOutput{})
	pulumi.RegisterOutputType(GetClientServiceAccountUserFederatedIdentityOutput{})
	pulumi.RegisterOutputType(GetClientServiceAccountUserFederatedIdentityArrayOutput{})
}
