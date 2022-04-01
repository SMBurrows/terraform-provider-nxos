// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

type resourceISISVRFType struct{}

func (t resourceISISVRFType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the IS-IS VRF configuration.", "isisDom", "Routing%20and%20Forwarding/isis:Dom/").AddParents("isis_instance").AddReferences("vrf").String,

		Attributes: map[string]tfsdk.Attribute{
			"device": {
				MarkdownDescription: "A device name from the provider configuration.",
				Type:                types.StringType,
				Optional:            true,
			},
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			"instance_name": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS instance name.").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"name": {
				MarkdownDescription: helpers.NewAttributeDescription("VRF name.").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
			},
			"admin_state": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative state.").AddDefaultValueDescription("enabled").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("enabled"),
				},
			},
			"authentication_check_l1": {
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Check for ISIS on Level-1.").AddDefaultValueDescription("true").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.BooleanDefaultModifier(true),
				},
			},
			"authentication_check_l2": {
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Check for ISIS on Level-2.").AddDefaultValueDescription("true").String,
				Type:                types.BoolType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.BooleanDefaultModifier(true),
				},
			},
			"authentication_key_l1": {
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Key for IS-IS on Level-1.").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"authentication_key_l2": {
				MarkdownDescription: helpers.NewAttributeDescription("Authentication Key for IS-IS on Level-2.").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"authentication_type_l1": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Authentication-Type for Level-1.").AddDefaultValueDescription("unknown").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("unknown"),
				},
			},
			"authentication_type_l2": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Authentication-Type for Level-2.").AddDefaultValueDescription("unknown").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("unknown"),
				},
			},
			"bandwidth_reference": {
				MarkdownDescription: helpers.NewAttributeDescription("The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost metric.").AddIntegerRangeDescription(0, 4294967295).AddDefaultValueDescription("40000").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 4294967295),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(40000),
				},
			},
			"banwidth_reference_unit": {
				MarkdownDescription: helpers.NewAttributeDescription("Bandwidth reference unit.").AddDefaultValueDescription("mbps").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("mbps"),
				},
			},
			"is_type": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS domain type.").AddDefaultValueDescription("l12").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("l12"),
				},
			},
			"metric_type": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS metric type.").AddDefaultValueDescription("wide").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("wide"),
				},
			},
			"mtu": {
				MarkdownDescription: helpers.NewAttributeDescription("The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352 bytes.").AddIntegerRangeDescription(256, 4352).AddDefaultValueDescription("1492").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(256, 4352),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(1492),
				},
			},
			"net": {
				MarkdownDescription: helpers.NewAttributeDescription("Holds IS-IS domain NET (address) value.").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"passive_default": {
				MarkdownDescription: helpers.NewAttributeDescription("IS-IS Domain passive-interface default level.").AddDefaultValueDescription("unknown").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("unknown"),
				},
			},
		},
	}, nil
}

func (t resourceISISVRFType) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resourceISISVRF{
		provider: provider,
	}, diags
}

type resourceISISVRF struct {
	provider provider
}

func (r resourceISISVRF) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state ISISVRF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody()
	_, err := r.provider.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.provider.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.provider.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.provider.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceISISVRF) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state ISISVRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.Value))

	res, err := r.provider.client.GetDn(state.Dn.Value, nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.provider.devices[state.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceISISVRF) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state ISISVRF

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody()
	_, err := r.provider.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.provider.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.provider.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.provider.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resourceISISVRF) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state ISISVRF

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.Value))

	res, err := r.provider.client.DeleteDn(state.Dn.Value, nxos.OverrideUrl(r.provider.devices[state.Device.Value]))
	if err != nil {
		errCode := res.Get("imdata.0.error.attributes.code").Str
		// Ignore errors of type "Cannot delete object"
		if errCode != "107" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.Value))

	resp.State.RemoveResource(ctx)
}

func (r resourceISISVRF) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}