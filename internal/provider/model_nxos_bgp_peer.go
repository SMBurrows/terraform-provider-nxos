// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type BGPPeer struct {
	Device    types.String `tfsdk:"device"`
	Dn        types.String `tfsdk:"id"`
	Bgp_asn   types.String `tfsdk:"asn"`
	Vrf_name  types.String `tfsdk:"vrf"`
	Addr      types.String `tfsdk:"address"`
	Asn       types.String `tfsdk:"remote_asn"`
	Name      types.String `tfsdk:"description"`
	PeerImp   types.String `tfsdk:"peer_template"`
	PeerType  types.String `tfsdk:"peer_type"`
	SrcIf     types.String `tfsdk:"source_interface"`
	HoldIntvl types.Int64  `tfsdk:"hold_time"`
	KaIntvl   types.Int64  `tfsdk:"keep_alive"`
}

func (data BGPPeer) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/peer-[%s]", data.Vrf_name.ValueString(), data.Addr.ValueString())
}

func (data BGPPeer) getClassName() string {
	return "bgpPeer"
}

func (data BGPPeer) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("addr", data.Addr.ValueString()).
		Set("asn", data.Asn.ValueString()).
		Set("name", data.Name.ValueString()).
		Set("peerImp", data.PeerImp.ValueString()).
		Set("peerType", data.PeerType.ValueString()).
		Set("srcIf", data.SrcIf.ValueString()).
		Set("holdIntvl", strconv.FormatInt(data.HoldIntvl.ValueInt64(), 10)).
		Set("kaIntvl", strconv.FormatInt(data.KaIntvl.ValueInt64(), 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *BGPPeer) fromBody(res gjson.Result) {
	data.Addr = types.StringValue(res.Get("*.attributes.addr").String())
	data.Asn = types.StringValue(res.Get("*.attributes.asn").String())
	data.Name = types.StringValue(res.Get("*.attributes.name").String())
	data.PeerImp = types.StringValue(res.Get("*.attributes.peerImp").String())
	data.PeerType = types.StringValue(res.Get("*.attributes.peerType").String())
	data.SrcIf = types.StringValue(res.Get("*.attributes.srcIf").String())
	data.HoldIntvl = types.Int64Value(res.Get("*.attributes.holdIntvl").Int())
	data.KaIntvl = types.Int64Value(res.Get("*.attributes.kaIntvl").Int())
}

func (data *BGPPeer) fromPlan(plan BGPPeer) {
	data.Device = plan.Device
	data.Dn = plan.Dn
	data.Bgp_asn = plan.Bgp_asn
	data.Vrf_name = plan.Vrf_name
}
