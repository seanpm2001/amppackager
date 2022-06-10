//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
MIT License

Copyright (c) 2021 Manabu Sonoda

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package zones

import (
	net "net"

	api "github.com/mimuret/golang-iij-dpf/pkg/api"
	core "github.com/mimuret/golang-iij-dpf/pkg/apis/dpf/v1/core"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttributeMeta) DeepCopyInto(out *AttributeMeta) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttributeMeta.
func (in *AttributeMeta) DeepCopy() *AttributeMeta {
	if in == nil {
		return nil
	}
	out := new(AttributeMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Contract) DeepCopyInto(out *Contract) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Contract = in.Contract
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Contract.
func (in *Contract) DeepCopy() *Contract {
	if in == nil {
		return nil
	}
	out := new(Contract)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Contract) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CurrentRecordList) DeepCopyInto(out *CurrentRecordList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Record, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CurrentRecordList.
func (in *CurrentRecordList) DeepCopy() *CurrentRecordList {
	if in == nil {
		return nil
	}
	out := new(CurrentRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *CurrentRecordList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultTTL) DeepCopyInto(out *DefaultTTL) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultTTL.
func (in *DefaultTTL) DeepCopy() *DefaultTTL {
	if in == nil {
		return nil
	}
	out := new(DefaultTTL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *DefaultTTL) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultTTLDiff) DeepCopyInto(out *DefaultTTLDiff) {
	*out = *in
	if in.New != nil {
		in, out := &in.New, &out.New
		*out = new(DefaultTTL)
		**out = **in
	}
	if in.Old != nil {
		in, out := &in.Old, &out.Old
		*out = new(DefaultTTL)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultTTLDiff.
func (in *DefaultTTLDiff) DeepCopy() *DefaultTTLDiff {
	if in == nil {
		return nil
	}
	out := new(DefaultTTLDiff)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DefaultTTLDiffList) DeepCopyInto(out *DefaultTTLDiffList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DefaultTTLDiff, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DefaultTTLDiffList.
func (in *DefaultTTLDiffList) DeepCopy() *DefaultTTLDiffList {
	if in == nil {
		return nil
	}
	out := new(DefaultTTLDiffList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *DefaultTTLDiffList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dnssec) DeepCopyInto(out *Dnssec) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dnssec.
func (in *Dnssec) DeepCopy() *Dnssec {
	if in == nil {
		return nil
	}
	out := new(Dnssec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Dnssec) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DnssecKskRollover) DeepCopyInto(out *DnssecKskRollover) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DnssecKskRollover.
func (in *DnssecKskRollover) DeepCopy() *DnssecKskRollover {
	if in == nil {
		return nil
	}
	out := new(DnssecKskRollover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *DnssecKskRollover) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DsRecord) DeepCopyInto(out *DsRecord) {
	*out = *in
	in.TransitAt.DeepCopyInto(&out.TransitAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DsRecord.
func (in *DsRecord) DeepCopy() *DsRecord {
	if in == nil {
		return nil
	}
	out := new(DsRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DsRecordList) DeepCopyInto(out *DsRecordList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DsRecord, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DsRecordList.
func (in *DsRecordList) DeepCopy() *DsRecordList {
	if in == nil {
		return nil
	}
	out := new(DsRecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *DsRecordList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *History) DeepCopyInto(out *History) {
	*out = *in
	in.CommittedAt.DeepCopyInto(&out.CommittedAt)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new History.
func (in *History) DeepCopy() *History {
	if in == nil {
		return nil
	}
	out := new(History)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistoryList) DeepCopyInto(out *HistoryList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]History, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistoryList.
func (in *HistoryList) DeepCopy() *HistoryList {
	if in == nil {
		return nil
	}
	out := new(HistoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *HistoryList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistoryText) DeepCopyInto(out *HistoryText) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	in.History.DeepCopyInto(&out.History)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistoryText.
func (in *HistoryText) DeepCopy() *HistoryText {
	if in == nil {
		return nil
	}
	out := new(HistoryText)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *HistoryText) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogList) DeepCopyInto(out *LogList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]core.Log, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogList.
func (in *LogList) DeepCopy() *LogList {
	if in == nil {
		return nil
	}
	out := new(LogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *LogList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedDnsList) DeepCopyInto(out *ManagedDnsList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedDnsList.
func (in *ManagedDnsList) DeepCopy() *ManagedDnsList {
	if in == nil {
		return nil
	}
	out := new(ManagedDnsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *ManagedDnsList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Record) DeepCopyInto(out *Record) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.RData != nil {
		in, out := &in.RData, &out.RData
		*out = make(RecordRDATASlice, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Record.
func (in *Record) DeepCopy() *Record {
	if in == nil {
		return nil
	}
	out := new(Record)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *Record) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordDiff) DeepCopyInto(out *RecordDiff) {
	*out = *in
	if in.New != nil {
		in, out := &in.New, &out.New
		*out = new(Record)
		(*in).DeepCopyInto(*out)
	}
	if in.Old != nil {
		in, out := &in.Old, &out.Old
		*out = new(Record)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordDiff.
func (in *RecordDiff) DeepCopy() *RecordDiff {
	if in == nil {
		return nil
	}
	out := new(RecordDiff)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordDiffList) DeepCopyInto(out *RecordDiffList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RecordDiff, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordDiffList.
func (in *RecordDiffList) DeepCopy() *RecordDiffList {
	if in == nil {
		return nil
	}
	out := new(RecordDiffList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *RecordDiffList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordList) DeepCopyInto(out *RecordList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	out.Count = in.Count
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Record, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordList.
func (in *RecordList) DeepCopy() *RecordList {
	if in == nil {
		return nil
	}
	out := new(RecordList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *RecordList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecordRDATA) DeepCopyInto(out *RecordRDATA) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordRDATA.
func (in *RecordRDATA) DeepCopy() *RecordRDATA {
	if in == nil {
		return nil
	}
	out := new(RecordRDATA)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RecordRDATASlice) DeepCopyInto(out *RecordRDATASlice) {
	{
		in := &in
		*out = make(RecordRDATASlice, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecordRDATASlice.
func (in RecordRDATASlice) DeepCopy() RecordRDATASlice {
	if in == nil {
		return nil
	}
	out := new(RecordRDATASlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneApply) DeepCopyInto(out *ZoneApply) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneApply.
func (in *ZoneApply) DeepCopy() *ZoneApply {
	if in == nil {
		return nil
	}
	out := new(ZoneApply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *ZoneApply) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneProxy) DeepCopyInto(out *ZoneProxy) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneProxy.
func (in *ZoneProxy) DeepCopy() *ZoneProxy {
	if in == nil {
		return nil
	}
	out := new(ZoneProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *ZoneProxy) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneProxyHealthCheck) DeepCopyInto(out *ZoneProxyHealthCheck) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = make(net.IP, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneProxyHealthCheck.
func (in *ZoneProxyHealthCheck) DeepCopy() *ZoneProxyHealthCheck {
	if in == nil {
		return nil
	}
	out := new(ZoneProxyHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZoneProxyHealthCheckList) DeepCopyInto(out *ZoneProxyHealthCheckList) {
	*out = *in
	out.AttributeMeta = in.AttributeMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ZoneProxyHealthCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneProxyHealthCheckList.
func (in *ZoneProxyHealthCheckList) DeepCopy() *ZoneProxyHealthCheckList {
	if in == nil {
		return nil
	}
	out := new(ZoneProxyHealthCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new api.Object.
func (in *ZoneProxyHealthCheckList) DeepCopyObject() api.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}