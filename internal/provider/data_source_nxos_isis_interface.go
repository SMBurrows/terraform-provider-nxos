// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ISISInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &ISISInterfaceDataSource{}
)

func NewISISInterfaceDataSource() datasource.DataSource {
	return &ISISInterfaceDataSource{}
}

type ISISInterfaceDataSource struct {
	data *NxosProviderData
}

func (d *ISISInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_isis_interface"
}

func (d *ISISInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewResourceDescription("This data source can read the IS-IS interface configuration.", "isisInternalIf", "Routing%20and%20Forwarding/isis:InternalIf/").String,

		Attributes: map[string]schema.Attribute{
			"device": schema.StringAttribute{
				MarkdownDescription: "A device name from the provider configuration.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "The distinguished name of the object.",
				Computed:            true,
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `eth1/1`.",
				Required:            true,
			},
			"authentication_check": schema.BoolAttribute{
				MarkdownDescription: "Authentication Check for ISIS without specific level.",
				Computed:            true,
			},
			"authentication_check_l1": schema.BoolAttribute{
				MarkdownDescription: "Authentication Check for ISIS on Level-1.",
				Computed:            true,
			},
			"authentication_check_l2": schema.BoolAttribute{
				MarkdownDescription: "Authentication Check for ISIS on Level-2.",
				Computed:            true,
			},
			"authentication_key": schema.StringAttribute{
				MarkdownDescription: "Authentication Key for IS-IS without specific level.",
				Computed:            true,
			},
			"authentication_key_l1": schema.StringAttribute{
				MarkdownDescription: "Authentication Key for IS-IS on Level-1.",
				Computed:            true,
			},
			"authentication_key_l2": schema.StringAttribute{
				MarkdownDescription: "Authentication Key for IS-IS on Level-2.",
				Computed:            true,
			},
			"authentication_type": schema.StringAttribute{
				MarkdownDescription: "IS-IS Authentication-Type without specific level.",
				Computed:            true,
			},
			"authentication_type_l1": schema.StringAttribute{
				MarkdownDescription: "IS-IS Authentication-Type for Level-1.",
				Computed:            true,
			},
			"authentication_type_l2": schema.StringAttribute{
				MarkdownDescription: "IS-IS Authentication-Type for Level-2.",
				Computed:            true,
			},
			"circuit_type": schema.StringAttribute{
				MarkdownDescription: "Circuit type.",
				Computed:            true,
			},
			"vrf": schema.StringAttribute{
				MarkdownDescription: "VRF.",
				Computed:            true,
			},
			"hello_interval": schema.Int64Attribute{
				MarkdownDescription: "Hello interval.",
				Computed:            true,
			},
			"hello_interval_l1": schema.Int64Attribute{
				MarkdownDescription: "Hello interval Level-1.",
				Computed:            true,
			},
			"hello_interval_l2": schema.Int64Attribute{
				MarkdownDescription: "Hello interval Level-2.",
				Computed:            true,
			},
			"hello_multiplier": schema.Int64Attribute{
				MarkdownDescription: "Hello multiplier.",
				Computed:            true,
			},
			"hello_multiplier_l1": schema.Int64Attribute{
				MarkdownDescription: "Hello multiplier Level-1.",
				Computed:            true,
			},
			"hello_multiplier_l2": schema.Int64Attribute{
				MarkdownDescription: "Hello multiplier Level-2.",
				Computed:            true,
			},
			"hello_padding": schema.StringAttribute{
				MarkdownDescription: "Hello padding.",
				Computed:            true,
			},
			"metric_l1": schema.Int64Attribute{
				MarkdownDescription: "Interface metric Level-1.",
				Computed:            true,
			},
			"metric_l2": schema.Int64Attribute{
				MarkdownDescription: "Interface metric Level-2.",
				Computed:            true,
			},
			"mtu_check": schema.BoolAttribute{
				MarkdownDescription: "MTU Check for IS-IS without specific level.",
				Computed:            true,
			},
			"mtu_check_l1": schema.BoolAttribute{
				MarkdownDescription: "MTU Check for IS-IS on Level-1.",
				Computed:            true,
			},
			"mtu_check_l2": schema.BoolAttribute{
				MarkdownDescription: "MTU Check for IS-IS on Level-2.",
				Computed:            true,
			},
			"network_type_p2p": schema.StringAttribute{
				MarkdownDescription: "Enabling Point-to-Point Network Type on IS-IS Interface.",
				Computed:            true,
			},
			"passive": schema.StringAttribute{
				MarkdownDescription: "IS-IS Passive Interface Info.",
				Computed:            true,
			},
			"priority_l1": schema.Int64Attribute{
				MarkdownDescription: "Circuit priority.",
				Computed:            true,
			},
			"priority_l2": schema.Int64Attribute{
				MarkdownDescription: "Circuit priority.",
				Computed:            true,
			},
		},
	}
}

func (d *ISISInterfaceDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.data = req.ProviderData.(*NxosProviderData)
}

func (d *ISISInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config, state ISISInterface

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.data.client.GetDn(config.getDn(), nxos.OverrideUrl(d.data.devices[config.Device.ValueString()]))

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(config)
	state.Dn = types.StringValue(config.getDn())

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
