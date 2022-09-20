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
var _ resource.Resource = &PhysicalInterfaceResource{}
var _ resource.ResourceWithImportState = &PhysicalInterfaceResource{}

func NewPhysicalInterfaceResource() resource.Resource {
	return &PhysicalInterfaceResource{}
}

type PhysicalInterfaceResource struct {
	data NxosProviderData
}

func (r *PhysicalInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_physical_interface"
}

func (r *PhysicalInterfaceResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This resource can manage a physical interface.", "l1PhysIf", "System/l1:PhysIf/").AddChildren("physical_interface_vrf").String,

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
				MarkdownDescription: helpers.NewAttributeDescription("Must match first field in the output of `show intf brief`. Example: `eth1/1`.").String,
				Type:                types.StringType,
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.RequiresReplace(),
				},
			},
			"fec_mode": {
				MarkdownDescription: helpers.NewAttributeDescription("FEC mode.").AddStringEnumDescription("fc-fec", "rs-fec", "fec-off", "auto", "rs-ieee", "rs-cons16", "kp-fec").AddDefaultValueDescription("auto").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("fc-fec", "rs-fec", "fec-off", "auto", "rs-ieee", "rs-cons16", "kp-fec"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"access_vlan": {
				MarkdownDescription: helpers.NewAttributeDescription("Access VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.").AddDefaultValueDescription("vlan-1").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("vlan-1"),
				},
			},
			"admin_state": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port state.").AddStringEnumDescription("up", "down").AddDefaultValueDescription("up").String,
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
			"auto_negotiation": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port auto-negotiation.").AddStringEnumDescription("on", "off", "25G").AddDefaultValueDescription("on").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("on", "off", "25G"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("on"),
				},
			},
			"bandwidth": {
				MarkdownDescription: helpers.NewAttributeDescription("The bandwidth parameter for a routed interface, port channel, or subinterface.").AddIntegerRangeDescription(0, 3200000000).AddDefaultValueDescription("0").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 3200000000),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(0),
				},
			},
			"delay": {
				MarkdownDescription: helpers.NewAttributeDescription("The administrative port delay time.").AddIntegerRangeDescription(1, 16777215).AddDefaultValueDescription("1").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(1, 16777215),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(1),
				},
			},
			"description": {
				MarkdownDescription: helpers.NewAttributeDescription("Interface description.").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
			"duplex": {
				MarkdownDescription: helpers.NewAttributeDescription("Duplex.").AddStringEnumDescription("auto", "full", "half").AddDefaultValueDescription("auto").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("auto", "full", "half"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"layer": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port layer.").AddStringEnumDescription("Layer2", "Layer3").AddDefaultValueDescription("Layer2").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("Layer2", "Layer3"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("Layer2"),
				},
			},
			"link_logging": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative link logging.").AddStringEnumDescription("default", "enable", "disable").AddDefaultValueDescription("default").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("default", "enable", "disable"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("default"),
				},
			},
			"link_debounce_down": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port link debounce interval.").AddIntegerRangeDescription(0, 20000).AddDefaultValueDescription("100").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 20000),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(100),
				},
			},
			"link_debounce_up": {
				MarkdownDescription: helpers.NewAttributeDescription("Link Debounce Interval - LinkUp Event.").AddIntegerRangeDescription(0, 20000).AddDefaultValueDescription("0").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(0, 20000),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(0),
				},
			},
			"medium": {
				MarkdownDescription: helpers.NewAttributeDescription("The administrative port medium type.").AddStringEnumDescription("broadcast", "p2p").AddDefaultValueDescription("broadcast").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("broadcast", "p2p"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("broadcast"),
				},
			},
			"mode": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port mode.").AddStringEnumDescription("access", "trunk", "fex-fabric", "dot1q-tunnel", "promiscuous", "host", "trunk_secondary", "trunk_promiscuous", "vntag").AddDefaultValueDescription("access").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("access", "trunk", "fex-fabric", "dot1q-tunnel", "promiscuous", "host", "trunk_secondary", "trunk_promiscuous", "vntag"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("access"),
				},
			},
			"mtu": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port MTU.").AddIntegerRangeDescription(576, 9216).AddDefaultValueDescription("1500").String,
				Type:                types.Int64Type,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator(576, 9216),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.IntegerDefaultModifier(1500),
				},
			},
			"native_vlan": {
				MarkdownDescription: helpers.NewAttributeDescription("Native VLAN. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.").AddDefaultValueDescription("vlan-1").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("vlan-1"),
				},
			},
			"speed": {
				MarkdownDescription: helpers.NewAttributeDescription("Administrative port speed.").AddStringEnumDescription("unknown", "100M", "1G", "10G", "40G", "auto", "auto 100M", "auto 100M 1G", "100G", "25G", "10M", "50G", "200G", "400G", "2.5G", "5G", "auto 2.5G 5G 10G", "auto 100M 1G 2.5G 5G").AddDefaultValueDescription("auto").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("unknown", "100M", "1G", "10G", "40G", "auto", "auto 100M", "auto 100M 1G", "100G", "25G", "10M", "50G", "200G", "400G", "2.5G", "5G", "auto 2.5G 5G 10G", "auto 100M 1G 2.5G 5G"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"speed_group": {
				MarkdownDescription: helpers.NewAttributeDescription("Speed group.").AddStringEnumDescription("unknown", "1000", "10000", "40000", "auto", "25000").AddDefaultValueDescription("auto").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("unknown", "1000", "10000", "40000", "auto", "25000"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("auto"),
				},
			},
			"trunk_vlans": {
				MarkdownDescription: helpers.NewAttributeDescription("List of trunk VLANs.").AddDefaultValueDescription("1-4094").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("1-4094"),
				},
			},
			"uni_directional_ethernet": {
				MarkdownDescription: helpers.NewAttributeDescription("UDE (Uni-Directional Ethernet).").AddStringEnumDescription("disable", "send-only", "receive-only").AddDefaultValueDescription("disable").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator("disable", "send-only", "receive-only"),
				},
				PlanModifiers: tfsdk.AttributePlanModifiers{
					helpers.StringDefaultModifier("disable"),
				},
			},
			"user_configured_flags": {
				MarkdownDescription: helpers.NewAttributeDescription("Port User Config Flags.").String,
				Type:                types.StringType,
				Optional:            true,
				Computed:            true,
			},
		},
	}, nil
}

func (r *PhysicalInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	data, ok := req.ProviderData.(NxosProviderData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected data, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.data = data
}

func (r *PhysicalInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state PhysicalInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody()
	_, err := r.data.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.data.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)
	state.Dn.Value = plan.getDn()

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *PhysicalInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PhysicalInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.Value))

	res, err := r.data.client.GetDn(state.Dn.Value, nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[state.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *PhysicalInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PhysicalInterface

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody()
	_, err := r.data.client.Post(plan.getDn(), body.Str, nxos.OverrideUrl(r.data.devices[plan.Device.Value]))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.data.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"), nxos.OverrideUrl(r.data.devices[plan.Device.Value]))
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

func (r *PhysicalInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PhysicalInterface

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.Value))

	res, err := r.data.client.DeleteDn(state.Dn.Value, nxos.OverrideUrl(r.data.devices[state.Device.Value]))
	if err != nil {
		errCode := res.Get("imdata.0.error.attributes.code").Str
		// Ignore errors of type "Cannot delete object"
		if errCode != "1" && errCode != "107" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.Value))

	resp.State.RemoveResource(ctx)
}

func (r *PhysicalInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
