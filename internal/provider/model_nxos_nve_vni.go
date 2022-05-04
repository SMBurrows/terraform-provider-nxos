// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
	"github.com/tidwall/gjson"
)

type NVEVNI struct {
	Device           types.String `tfsdk:"device"`
	Dn               types.String `tfsdk:"id"`
	Vni              types.Int64  `tfsdk:"vni"`
	AssociateVrfFlag types.Bool   `tfsdk:"associate_vrf"`
	McastGroup       types.String `tfsdk:"multicast_group"`
	MultisiteIngRepl types.String `tfsdk:"multisite_ingress_replication"`
	SuppressARP      types.String `tfsdk:"suppress_arp"`
}

func (data NVEVNI) getDn() string {
	return fmt.Sprintf("sys/eps/epId-[1]/nws/vni-[%v]", data.Vni.Value)
}

func (data NVEVNI) getClassName() string {
	return "nvoNw"
}

func (data NVEVNI) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("vni", strconv.FormatInt(data.Vni.Value, 10)).
		Set("associateVrfFlag", strconv.FormatBool(data.AssociateVrfFlag.Value)).
		Set("mcastGroup", data.McastGroup.Value).
		Set("multisiteIngRepl", data.MultisiteIngRepl.Value).
		Set("suppressARP", data.SuppressARP.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *NVEVNI) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Vni.Value = res.Get("*.attributes.vni").Int()
	data.AssociateVrfFlag.Value = helpers.ParseNxosBoolean(res.Get("*.attributes.associateVrfFlag").String())
	data.McastGroup.Value = res.Get("*.attributes.mcastGroup").String()
	data.MultisiteIngRepl.Value = res.Get("*.attributes.multisiteIngRepl").String()
	data.SuppressARP.Value = res.Get("*.attributes.suppressARP").String()
}

func (data *NVEVNI) fromPlan(plan NVEVNI) {
	data.Device = plan.Device
}
