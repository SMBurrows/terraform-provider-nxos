// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type SpanningTreeInterface struct {
	Device     types.String `tfsdk:"device"`
	Dn         types.String `tfsdk:"id"`
	Id         types.String `tfsdk:"interface_id"`
	AdminSt    types.String `tfsdk:"admin_state"`
	Bpdufilter types.String `tfsdk:"bpdu_filter"`
	Bpduguard  types.String `tfsdk:"bpdu_guard"`
	Cost       types.Int64  `tfsdk:"cost"`
	Ctrl       types.String `tfsdk:"ctrl"`
	Guard      types.String `tfsdk:"guard"`
	LinkType   types.String `tfsdk:"link_type"`
	Mode       types.String `tfsdk:"mode"`
	Priority   types.Int64  `tfsdk:"priority"`
}

func (data SpanningTreeInterface) getDn() string {
	return fmt.Sprintf("sys/stp/inst/if-[%s]", data.Id.ValueString())
}

func (data SpanningTreeInterface) getClassName() string {
	return "stpIf"
}

func (data SpanningTreeInterface) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("id", data.Id.ValueString()).
		Set("adminSt", data.AdminSt.ValueString()).
		Set("bpdufilter", data.Bpdufilter.ValueString()).
		Set("bpduguard", data.Bpduguard.ValueString()).
		Set("cost", strconv.FormatInt(data.Cost.ValueInt64(), 10)).
		Set("ctrl", data.Ctrl.ValueString()).
		Set("guard", data.Guard.ValueString()).
		Set("linkType", data.LinkType.ValueString()).
		Set("mode", data.Mode.ValueString()).
		Set("priority", strconv.FormatInt(data.Priority.ValueInt64(), 10))
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *SpanningTreeInterface) fromBody(res gjson.Result) {
	data.Id = types.StringValue(res.Get("*.attributes.id").String())
	data.AdminSt = types.StringValue(res.Get("*.attributes.adminSt").String())
	data.Bpdufilter = types.StringValue(res.Get("*.attributes.bpdufilter").String())
	data.Bpduguard = types.StringValue(res.Get("*.attributes.bpduguard").String())
	data.Cost = types.Int64Value(res.Get("*.attributes.cost").Int())
	data.Ctrl = types.StringValue(res.Get("*.attributes.ctrl").String())
	data.Guard = types.StringValue(res.Get("*.attributes.guard").String())
	data.LinkType = types.StringValue(res.Get("*.attributes.linkType").String())
	data.Mode = types.StringValue(res.Get("*.attributes.mode").String())
	data.Priority = types.Int64Value(res.Get("*.attributes.priority").Int())
}

func (data *SpanningTreeInterface) fromPlan(plan SpanningTreeInterface) {
	data.Device = plan.Device
	data.Dn = plan.Dn
}