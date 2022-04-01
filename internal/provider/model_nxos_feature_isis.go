// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type FeatureISIS struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data FeatureISIS) getDn() string {
	return "sys/fm/isis"
}

func (data FeatureISIS) getClassName() string {
	return "fmIsis"
}

func (data FeatureISIS) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *FeatureISIS) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
}

func (data *FeatureISIS) fromPlan(plan FeatureISIS) {
	data.Device = plan.Device
}