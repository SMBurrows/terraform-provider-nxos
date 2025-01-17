// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type VRFAddressFamily struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Vrf    types.String `tfsdk:"vrf"`
	Type   types.String `tfsdk:"address_family"`
}

func (data VRFAddressFamily) getDn() string {
	return fmt.Sprintf("sys/inst-[%s]/dom-[%[1]s]/af-[%s]", data.Vrf.ValueString(), data.Type.ValueString())
}

func (data VRFAddressFamily) getClassName() string {
	return "rtctrlDomAf"
}

func (data VRFAddressFamily) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("type", data.Type.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *VRFAddressFamily) fromBody(res gjson.Result) {
	data.Type = types.StringValue(res.Get("*.attributes.type").String())
}

func (data *VRFAddressFamily) fromPlan(plan VRFAddressFamily) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Vrf = plan.Vrf
}
