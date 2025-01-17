// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &PIMSSMRangeResource{}
var _ resource.ResourceWithImportState = &PIMSSMRangeResource{}

func NewPIMSSMRangeResource() resource.Resource {
	return &PIMSSMRangeResource{}
}

type PIMSSMRangeResource struct {
	data *NxosProviderData
}

func (r *PIMSSMRangeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_pim_ssm_range"
}

func (r *PIMSSMRangeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage the PIM SSM range configuration.", "pimSSMRangeP", "Layer%203/pim:SSMRangeP/").AddParents("pim_ssm_policy").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vrf_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("VRF name.").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"group_list_1": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Group list 1.").AddDefaultValueDescription("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("0.0.0.0"),
				},
			},
			"group_list_2": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Group list 2.").AddDefaultValueDescription("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("0.0.0.0"),
				},
			},
			"group_list_3": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Group list 3.").AddDefaultValueDescription("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("0.0.0.0"),
				},
			},
			"group_list_4": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Group list 4.").AddDefaultValueDescription("0.0.0.0").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					helpers.StringDefaultModifier("0.0.0.0"),
				},
			},
			"prefix_list": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Prefix list name.").String,
				Optional:            true,
				Computed:            true,
			},
			"route_map": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Route map name.").String,
				Optional:            true,
				Computed:            true,
			},
			"ssm_none": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Exclude standard SSM range (232.0.0.0/8).").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					helpers.BooleanDefaultModifier(false),
				},
			},
		},
	}
}

func (r *PIMSSMRangeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	r.data = req.ProviderData.(*NxosProviderData)
}

func (r *PIMSSMRangeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state PIMSSMRange

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

func (r *PIMSSMRangeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PIMSSMRange

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

func (r *PIMSSMRangeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PIMSSMRange

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

func (r *PIMSSMRangeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PIMSSMRange

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

func (r *PIMSSMRangeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
