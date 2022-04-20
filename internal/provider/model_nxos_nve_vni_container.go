// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type NVEVNIContainer struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
}

func (data NVEVNIContainer) getDn() string {
	return "sys/eps/epId-[1]/nws"
}

func (data NVEVNIContainer) getClassName() string {
	return "nvoNws"
}

func (data NVEVNIContainer) toBody() nxos.Body {
	return nxos.Body{Str: `{"` + data.getClassName() + `":{"attributes":{}}}`}
}

func (data *NVEVNIContainer) fromBody(res gjson.Result) {
	data.Dn.Value = res.Get("*.attributes.dn").String()
}

func (data *NVEVNIContainer) fromPlan(plan NVEVNIContainer) {
	data.Device = plan.Device
}