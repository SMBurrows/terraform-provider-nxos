// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type VRFRouteTargetDirection struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	Vrf     types.String `tfsdk:"vrf"`
	Af_type types.String `tfsdk:"address_family"`
	Rt_type types.String `tfsdk:"route_target_address_family"`
	Type    types.String `tfsdk:"direction"`
}

func (data VRFRouteTargetDirection) getDn() string {
	return fmt.Sprintf("sys/inst-[%s]/dom-[%[1]s]/af-[%s]/ctrl-[%s]/rttp-[%s]", data.Vrf.ValueString(), data.Af_type.ValueString(), data.Rt_type.ValueString(), data.Type.ValueString())
}

func (data VRFRouteTargetDirection) getClassName() string {
	return "rtctrlRttP"
}

func (data VRFRouteTargetDirection) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("type", data.Type.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *VRFRouteTargetDirection) fromBody(res gjson.Result) {
	data.Type = types.StringValue(res.Get("*.attributes.type").String())
}

func (data *VRFRouteTargetDirection) fromPlan(plan VRFRouteTargetDirection) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Vrf = plan.Vrf
	data.Af_type = plan.Af_type
	data.Rt_type = plan.Rt_type
}
