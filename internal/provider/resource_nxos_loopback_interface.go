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
var _ resource.Resource = &LoopbackInterfaceResource{}
var _ resource.ResourceWithImportState = &LoopbackInterfaceResource{}

func NewLoopbackInterfaceResource() resource.Resource {
	return &LoopbackInterfaceResource{}
}

type LoopbackInterfaceResource struct {
	data *NxosProviderData
}

func (r *LoopbackInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_loopback_interface"
}

func (r *LoopbackInterfaceResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage a loopback interface.", "l3LbRtdIf", "Layer%203/l3:LbRtdIf/").AddChildren("loopback_interface_vrf").String,

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
			"interface_id": {
				MarkdownDescription: helpers.NewAttributeDescription("Must match first field in the output of `show intf brief`. Example: `lo123`.").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.RequiresReplace(),
				},
			},
			"admin_state": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative state.").AddStringEnumDescription("up", "down").AddDefaultValueDescription("up").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("up", "down"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("up"),
				},
			},
			"description": {
				MarkdownDescription: helpers.NewAttributeDescription("Interface description.").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (r *LoopbackInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *LoopbackInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state LoopbackInterface

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

func (r *LoopbackInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LoopbackInterface

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

func (r *LoopbackInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state LoopbackInterface

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

func (r *LoopbackInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LoopbackInterface

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

func (r *LoopbackInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
