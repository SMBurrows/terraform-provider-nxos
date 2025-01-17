// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type BGPPeerTemplateMaxPrefix struct {
	Device      types.String `tfsdk:"device"`
	Dn          types.String `tfsdk:"id"`
	Asn         types.String `tfsdk:"asn"`
	Name        types.String `tfsdk:"template_name"`
	Type        types.String `tfsdk:"address_family"`
	Action      types.String `tfsdk:"action"`
	MaxPfx      types.Int64  `tfsdk:"maximum_prefix"`
	RestartTime types.Int64  `tfsdk:"restart_time"`
	Thresh      types.Int64  `tfsdk:"threshold"`
}

func (data BGPPeerTemplateMaxPrefix) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[default]/peercont-[%s]/af-[%s]/maxpfxp", data.Name.ValueString(), data.Type.ValueString())
}

func (data BGPPeerTemplateMaxPrefix) getClassName() string {
	return "bgpMaxPfxP"
}

func (data BGPPeerTemplateMaxPrefix) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("action", data.Action.ValueString()).
		Set("maxPfx", strconv.FormatInt(data.MaxPfx.ValueInt64(), 10)).
		Set("restartTime", strconv.FormatInt(data.RestartTime.ValueInt64(), 10)).
		Set("thresh", strconv.FormatInt(data.Thresh.ValueInt64(), 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *BGPPeerTemplateMaxPrefix) fromBody(res gjson.Result) {
	data.Action = types.StringValue(res.Get("*.attributes.action").String())
	data.MaxPfx = types.Int64Value(res.Get("*.attributes.maxPfx").Int())
	data.RestartTime = types.Int64Value(res.Get("*.attributes.restartTime").Int())
	data.Thresh = types.Int64Value(res.Get("*.attributes.thresh").Int())
}

func (data *BGPPeerTemplateMaxPrefix) fromPlan(plan BGPPeerTemplateMaxPrefix) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Asn = plan.Asn
	data.Name = plan.Name
	data.Type = plan.Type
}
