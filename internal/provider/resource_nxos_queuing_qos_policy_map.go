// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &QueuingQOSPolicyMapResource{}
var _ resource.ResourceWithImportState = &QueuingQOSPolicyMapResource{}

func NewQueuingQOSPolicyMapResource() resource.Resource {
	return &QueuingQOSPolicyMapResource{}
}

type QueuingQOSPolicyMapResource struct {
	data *NxosProviderData
}

func (r *QueuingQOSPolicyMapResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_queuing_qos_policy_map"
}

func (r *QueuingQOSPolicyMapResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the queuing QoS policy map configuration.", "ipqosPMapInst", "Qos/ipqos:PMapInst/").AddChildren("queuing_qos_policy_map_match_class_map").String,

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
					resource.UseStateForUnknown(),
				},
			},
			"name": {
				MarkdownDescription: helpers.NewAttributeDescription("Policy map name.").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.RequiresReplace(),
				},
			},
			"match_type": {
				MarkdownDescription: helpers.NewAttributeDescription("Match type.").AddStringEnumDescription("match-any", "match-all", "match-first").AddDefaultValueDescription("match-all").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("match-any", "match-all", "match-first"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("match-all"),
				},
			},
		},
	}, nil
}

func (r *QueuingQOSPolicyMapResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *QueuingQOSPolicyMapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state QueuingQOSPolicyMap

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody()
	_, err := r.data.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.data.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)
	state.Dn = types.StringValue(plan.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *QueuingQOSPolicyMapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state QueuingQOSPolicyMap

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.ValueString()))

	res, err := r.data.client.GetDn(state.Dn.ValueString(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[state.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *QueuingQOSPolicyMapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state QueuingQOSPolicyMap

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody()
	_, err := r.data.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.data.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[plan.Device.ValueString()]))
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

func (r *QueuingQOSPolicyMapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state QueuingQOSPolicyMap

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.ValueString()))

	res, err := r.data.client.DeleteDn(state.Dn.ValueString(), nxos.OverrideUrl(r.data.devices[state.Device.ValueString()]))
	if err != nil {
		errCode := res.Get("imdata.0.error.attributes.code").Str
		// Ignore errors of type "Cannot delete object"
		if errCode != "1" && errCode != "107" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.ValueString()))

	resp.State.RemoveResource(ctx)
}

func (r *QueuingQOSPolicyMapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
