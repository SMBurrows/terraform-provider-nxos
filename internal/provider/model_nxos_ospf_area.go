// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"
	
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type OSPFArea struct {
	Dn types.String `tfsdk:"id"`
	Inst types.String `tfsdk:"instance_name"`
	Name types.String `tfsdk:"vrf_name"`
	Id types.String `tfsdk:"area_id"`
	AuthType types.String `tfsdk:"authentication_type"`
	Cost types.Int64 `tfsdk:"cost"`
	Type types.String `tfsdk:"type"`
}

func (data OSPFArea) getDn() string {
	return fmt.Sprintf("sys/ospf/inst-[%s]/dom-[%s]/area-[%s]", data.Inst.Value, data.Name.Value, data.Id.Value)
}

func (data OSPFArea) getClassName() string {
	return "ospfArea"
}

func (data OSPFArea) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", data.Id.Value).
		Set("authType", data.AuthType.Value).
		Set("cost", strconv.FormatInt(data.Cost.Value, 10)).
		Set("type", data.Type.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *OSPFArea) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.Id.Value = res.Get("*.attributes.id").String()
	data.AuthType.Value = res.Get("*.attributes.authType").String()
	data.Cost.Value = res.Get("*.attributes.cost").Int()
	data.Type.Value = res.Get("*.attributes.type").String()
}

func (data *OSPFArea) fromPlan(plan OSPFArea) {
	data.Inst.Value = plan.Inst.Value
	data.Name.Value = plan.Name.Value
}