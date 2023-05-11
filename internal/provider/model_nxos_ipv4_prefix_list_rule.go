// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type IPv4PrefixListRule struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
}

func (data IPv4PrefixListRule) getDn() string {
	return fmt.Sprintf("sys/rpm/pfxlistv4-[%s]", data.Name.ValueString())
}

func (data IPv4PrefixListRule) getClassName() string {
	return "rtpfxRuleV4"
}

func (data IPv4PrefixListRule) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.ValueString())
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *IPv4PrefixListRule) fromBody(res gjson.Result) {
	data.Name = types.StringValue(res.Get("*.attributes.name").String())
}

func (data *IPv4PrefixListRule) fromPlan(plan IPv4PrefixListRule) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}