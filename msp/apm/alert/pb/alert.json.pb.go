// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: alert.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*QueryAlertRuleRequest)(nil)
var _ json.Unmarshaler = (*QueryAlertRuleRequest)(nil)
var _ json.Marshaler = (*QueryAlertRuleResponse)(nil)
var _ json.Unmarshaler = (*QueryAlertRuleResponse)(nil)
var _ json.Marshaler = (*QueryAlertRequest)(nil)
var _ json.Unmarshaler = (*QueryAlertRequest)(nil)
var _ json.Marshaler = (*QueryAlertResponse)(nil)
var _ json.Unmarshaler = (*QueryAlertResponse)(nil)
var _ json.Marshaler = (*QueryAlertData)(nil)
var _ json.Unmarshaler = (*QueryAlertData)(nil)
var _ json.Marshaler = (*GetAlertRequest)(nil)
var _ json.Unmarshaler = (*GetAlertRequest)(nil)
var _ json.Marshaler = (*GetAlertResponse)(nil)
var _ json.Unmarshaler = (*GetAlertResponse)(nil)
var _ json.Marshaler = (*CreateAlertRequest)(nil)
var _ json.Unmarshaler = (*CreateAlertRequest)(nil)
var _ json.Marshaler = (*CreateAlertResponse)(nil)
var _ json.Unmarshaler = (*CreateAlertResponse)(nil)
var _ json.Marshaler = (*CreateAlertData)(nil)
var _ json.Unmarshaler = (*CreateAlertData)(nil)
var _ json.Marshaler = (*UpdateAlertRequest)(nil)
var _ json.Unmarshaler = (*UpdateAlertRequest)(nil)
var _ json.Marshaler = (*UpdateAlertResponse)(nil)
var _ json.Unmarshaler = (*UpdateAlertResponse)(nil)
var _ json.Marshaler = (*UpdateAlertEnableRequest)(nil)
var _ json.Unmarshaler = (*UpdateAlertEnableRequest)(nil)
var _ json.Marshaler = (*UpdateAlertEnableResponse)(nil)
var _ json.Unmarshaler = (*UpdateAlertEnableResponse)(nil)
var _ json.Marshaler = (*DeleteAlertRequest)(nil)
var _ json.Unmarshaler = (*DeleteAlertRequest)(nil)
var _ json.Marshaler = (*DeleteAlertResponse)(nil)
var _ json.Unmarshaler = (*DeleteAlertResponse)(nil)
var _ json.Marshaler = (*DeleteAlertData)(nil)
var _ json.Unmarshaler = (*DeleteAlertData)(nil)
var _ json.Marshaler = (*QueryCustomizeMetricRequest)(nil)
var _ json.Unmarshaler = (*QueryCustomizeMetricRequest)(nil)
var _ json.Marshaler = (*QueryCustomizeMetricResponse)(nil)
var _ json.Unmarshaler = (*QueryCustomizeMetricResponse)(nil)
var _ json.Marshaler = (*QueryCustomizeNotifyTargetRequest)(nil)
var _ json.Unmarshaler = (*QueryCustomizeNotifyTargetRequest)(nil)
var _ json.Marshaler = (*QueryCustomizeNotifyTargetResponse)(nil)
var _ json.Unmarshaler = (*QueryCustomizeNotifyTargetResponse)(nil)
var _ json.Marshaler = (*QueryCustomizeAlertsRequest)(nil)
var _ json.Unmarshaler = (*QueryCustomizeAlertsRequest)(nil)
var _ json.Marshaler = (*QueryCustomizeAlertsResponse)(nil)
var _ json.Unmarshaler = (*QueryCustomizeAlertsResponse)(nil)
var _ json.Marshaler = (*GetCustomizeAlertRequest)(nil)
var _ json.Unmarshaler = (*GetCustomizeAlertRequest)(nil)
var _ json.Marshaler = (*GetCustomizeAlertResponse)(nil)
var _ json.Unmarshaler = (*GetCustomizeAlertResponse)(nil)
var _ json.Marshaler = (*CreateCustomizeAlertRequest)(nil)
var _ json.Unmarshaler = (*CreateCustomizeAlertRequest)(nil)
var _ json.Marshaler = (*CreateCustomizeAlertResponse)(nil)
var _ json.Unmarshaler = (*CreateCustomizeAlertResponse)(nil)
var _ json.Marshaler = (*CreateCustomizeAlertData)(nil)
var _ json.Unmarshaler = (*CreateCustomizeAlertData)(nil)
var _ json.Marshaler = (*UpdateCustomizeAlertRequest)(nil)
var _ json.Unmarshaler = (*UpdateCustomizeAlertRequest)(nil)
var _ json.Marshaler = (*UpdateCustomizeAlertResponse)(nil)
var _ json.Unmarshaler = (*UpdateCustomizeAlertResponse)(nil)
var _ json.Marshaler = (*UpdateCustomizeAlertEnableRequest)(nil)
var _ json.Unmarshaler = (*UpdateCustomizeAlertEnableRequest)(nil)
var _ json.Marshaler = (*UpdateCustomizeAlertEnableResponse)(nil)
var _ json.Unmarshaler = (*UpdateCustomizeAlertEnableResponse)(nil)
var _ json.Marshaler = (*DeleteCustomizeAlertRequest)(nil)
var _ json.Unmarshaler = (*DeleteCustomizeAlertRequest)(nil)
var _ json.Marshaler = (*DeleteCustomizeAlertResponse)(nil)
var _ json.Unmarshaler = (*DeleteCustomizeAlertResponse)(nil)
var _ json.Marshaler = (*DeleteCustomizeAlertData)(nil)
var _ json.Unmarshaler = (*DeleteCustomizeAlertData)(nil)
var _ json.Marshaler = (*GetAlertRecordAttrsRequest)(nil)
var _ json.Unmarshaler = (*GetAlertRecordAttrsRequest)(nil)
var _ json.Marshaler = (*GetAlertRecordAttrsResponse)(nil)
var _ json.Unmarshaler = (*GetAlertRecordAttrsResponse)(nil)
var _ json.Marshaler = (*GetAlertRecordsRequest)(nil)
var _ json.Unmarshaler = (*GetAlertRecordsRequest)(nil)
var _ json.Marshaler = (*GetAlertRecordsResponse)(nil)
var _ json.Unmarshaler = (*GetAlertRecordsResponse)(nil)
var _ json.Marshaler = (*GetAlertRecordsData)(nil)
var _ json.Unmarshaler = (*GetAlertRecordsData)(nil)
var _ json.Marshaler = (*GetAlertRecordRequest)(nil)
var _ json.Unmarshaler = (*GetAlertRecordRequest)(nil)
var _ json.Marshaler = (*GetAlertRecordResponse)(nil)
var _ json.Unmarshaler = (*GetAlertRecordResponse)(nil)
var _ json.Marshaler = (*GetAlertHistoriesRequest)(nil)
var _ json.Unmarshaler = (*GetAlertHistoriesRequest)(nil)
var _ json.Marshaler = (*GetAlertHistoriesResponse)(nil)
var _ json.Unmarshaler = (*GetAlertHistoriesResponse)(nil)
var _ json.Marshaler = (*CreateAlertRecordIssueRequest)(nil)
var _ json.Unmarshaler = (*CreateAlertRecordIssueRequest)(nil)
var _ json.Marshaler = (*CreateAlertRecordIssueResponse)(nil)
var _ json.Unmarshaler = (*CreateAlertRecordIssueResponse)(nil)
var _ json.Marshaler = (*UpdateAlertRecordIssueRequest)(nil)
var _ json.Unmarshaler = (*UpdateAlertRecordIssueRequest)(nil)
var _ json.Marshaler = (*UpdateAlertRecordIssueResponse)(nil)
var _ json.Unmarshaler = (*UpdateAlertRecordIssueResponse)(nil)
var _ json.Marshaler = (*DashboardPreviewRequest)(nil)
var _ json.Unmarshaler = (*DashboardPreviewRequest)(nil)
var _ json.Marshaler = (*DashboardPreviewResponse)(nil)
var _ json.Unmarshaler = (*DashboardPreviewResponse)(nil)

// QueryAlertRuleRequest implement json.Marshaler.
func (m *QueryAlertRuleRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAlertRuleRequest implement json.Marshaler.
func (m *QueryAlertRuleRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryAlertRuleResponse implement json.Marshaler.
func (m *QueryAlertRuleResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAlertRuleResponse implement json.Marshaler.
func (m *QueryAlertRuleResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryAlertRequest implement json.Marshaler.
func (m *QueryAlertRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAlertRequest implement json.Marshaler.
func (m *QueryAlertRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryAlertResponse implement json.Marshaler.
func (m *QueryAlertResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAlertResponse implement json.Marshaler.
func (m *QueryAlertResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryAlertData implement json.Marshaler.
func (m *QueryAlertData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryAlertData implement json.Marshaler.
func (m *QueryAlertData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertRequest implement json.Marshaler.
func (m *GetAlertRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertRequest implement json.Marshaler.
func (m *GetAlertRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertResponse implement json.Marshaler.
func (m *GetAlertResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertResponse implement json.Marshaler.
func (m *GetAlertResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateAlertRequest implement json.Marshaler.
func (m *CreateAlertRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateAlertRequest implement json.Marshaler.
func (m *CreateAlertRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateAlertResponse implement json.Marshaler.
func (m *CreateAlertResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateAlertResponse implement json.Marshaler.
func (m *CreateAlertResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateAlertData implement json.Marshaler.
func (m *CreateAlertData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateAlertData implement json.Marshaler.
func (m *CreateAlertData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateAlertRequest implement json.Marshaler.
func (m *UpdateAlertRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateAlertRequest implement json.Marshaler.
func (m *UpdateAlertRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateAlertResponse implement json.Marshaler.
func (m *UpdateAlertResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateAlertResponse implement json.Marshaler.
func (m *UpdateAlertResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateAlertEnableRequest implement json.Marshaler.
func (m *UpdateAlertEnableRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateAlertEnableRequest implement json.Marshaler.
func (m *UpdateAlertEnableRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateAlertEnableResponse implement json.Marshaler.
func (m *UpdateAlertEnableResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateAlertEnableResponse implement json.Marshaler.
func (m *UpdateAlertEnableResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteAlertRequest implement json.Marshaler.
func (m *DeleteAlertRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteAlertRequest implement json.Marshaler.
func (m *DeleteAlertRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteAlertResponse implement json.Marshaler.
func (m *DeleteAlertResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteAlertResponse implement json.Marshaler.
func (m *DeleteAlertResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteAlertData implement json.Marshaler.
func (m *DeleteAlertData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteAlertData implement json.Marshaler.
func (m *DeleteAlertData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryCustomizeMetricRequest implement json.Marshaler.
func (m *QueryCustomizeMetricRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryCustomizeMetricRequest implement json.Marshaler.
func (m *QueryCustomizeMetricRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryCustomizeMetricResponse implement json.Marshaler.
func (m *QueryCustomizeMetricResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryCustomizeMetricResponse implement json.Marshaler.
func (m *QueryCustomizeMetricResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryCustomizeNotifyTargetRequest implement json.Marshaler.
func (m *QueryCustomizeNotifyTargetRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryCustomizeNotifyTargetRequest implement json.Marshaler.
func (m *QueryCustomizeNotifyTargetRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryCustomizeNotifyTargetResponse implement json.Marshaler.
func (m *QueryCustomizeNotifyTargetResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryCustomizeNotifyTargetResponse implement json.Marshaler.
func (m *QueryCustomizeNotifyTargetResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryCustomizeAlertsRequest implement json.Marshaler.
func (m *QueryCustomizeAlertsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryCustomizeAlertsRequest implement json.Marshaler.
func (m *QueryCustomizeAlertsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// QueryCustomizeAlertsResponse implement json.Marshaler.
func (m *QueryCustomizeAlertsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// QueryCustomizeAlertsResponse implement json.Marshaler.
func (m *QueryCustomizeAlertsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCustomizeAlertRequest implement json.Marshaler.
func (m *GetCustomizeAlertRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCustomizeAlertRequest implement json.Marshaler.
func (m *GetCustomizeAlertRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetCustomizeAlertResponse implement json.Marshaler.
func (m *GetCustomizeAlertResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetCustomizeAlertResponse implement json.Marshaler.
func (m *GetCustomizeAlertResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateCustomizeAlertRequest implement json.Marshaler.
func (m *CreateCustomizeAlertRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateCustomizeAlertRequest implement json.Marshaler.
func (m *CreateCustomizeAlertRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateCustomizeAlertResponse implement json.Marshaler.
func (m *CreateCustomizeAlertResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateCustomizeAlertResponse implement json.Marshaler.
func (m *CreateCustomizeAlertResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateCustomizeAlertData implement json.Marshaler.
func (m *CreateCustomizeAlertData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateCustomizeAlertData implement json.Marshaler.
func (m *CreateCustomizeAlertData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCustomizeAlertRequest implement json.Marshaler.
func (m *UpdateCustomizeAlertRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCustomizeAlertRequest implement json.Marshaler.
func (m *UpdateCustomizeAlertRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCustomizeAlertResponse implement json.Marshaler.
func (m *UpdateCustomizeAlertResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCustomizeAlertResponse implement json.Marshaler.
func (m *UpdateCustomizeAlertResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCustomizeAlertEnableRequest implement json.Marshaler.
func (m *UpdateCustomizeAlertEnableRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCustomizeAlertEnableRequest implement json.Marshaler.
func (m *UpdateCustomizeAlertEnableRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateCustomizeAlertEnableResponse implement json.Marshaler.
func (m *UpdateCustomizeAlertEnableResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateCustomizeAlertEnableResponse implement json.Marshaler.
func (m *UpdateCustomizeAlertEnableResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteCustomizeAlertRequest implement json.Marshaler.
func (m *DeleteCustomizeAlertRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteCustomizeAlertRequest implement json.Marshaler.
func (m *DeleteCustomizeAlertRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteCustomizeAlertResponse implement json.Marshaler.
func (m *DeleteCustomizeAlertResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteCustomizeAlertResponse implement json.Marshaler.
func (m *DeleteCustomizeAlertResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteCustomizeAlertData implement json.Marshaler.
func (m *DeleteCustomizeAlertData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteCustomizeAlertData implement json.Marshaler.
func (m *DeleteCustomizeAlertData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertRecordAttrsRequest implement json.Marshaler.
func (m *GetAlertRecordAttrsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertRecordAttrsRequest implement json.Marshaler.
func (m *GetAlertRecordAttrsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertRecordAttrsResponse implement json.Marshaler.
func (m *GetAlertRecordAttrsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertRecordAttrsResponse implement json.Marshaler.
func (m *GetAlertRecordAttrsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertRecordsRequest implement json.Marshaler.
func (m *GetAlertRecordsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertRecordsRequest implement json.Marshaler.
func (m *GetAlertRecordsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertRecordsResponse implement json.Marshaler.
func (m *GetAlertRecordsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertRecordsResponse implement json.Marshaler.
func (m *GetAlertRecordsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertRecordsData implement json.Marshaler.
func (m *GetAlertRecordsData) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertRecordsData implement json.Marshaler.
func (m *GetAlertRecordsData) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertRecordRequest implement json.Marshaler.
func (m *GetAlertRecordRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertRecordRequest implement json.Marshaler.
func (m *GetAlertRecordRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertRecordResponse implement json.Marshaler.
func (m *GetAlertRecordResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertRecordResponse implement json.Marshaler.
func (m *GetAlertRecordResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertHistoriesRequest implement json.Marshaler.
func (m *GetAlertHistoriesRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertHistoriesRequest implement json.Marshaler.
func (m *GetAlertHistoriesRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetAlertHistoriesResponse implement json.Marshaler.
func (m *GetAlertHistoriesResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetAlertHistoriesResponse implement json.Marshaler.
func (m *GetAlertHistoriesResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateAlertRecordIssueRequest implement json.Marshaler.
func (m *CreateAlertRecordIssueRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateAlertRecordIssueRequest implement json.Marshaler.
func (m *CreateAlertRecordIssueRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateAlertRecordIssueResponse implement json.Marshaler.
func (m *CreateAlertRecordIssueResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateAlertRecordIssueResponse implement json.Marshaler.
func (m *CreateAlertRecordIssueResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateAlertRecordIssueRequest implement json.Marshaler.
func (m *UpdateAlertRecordIssueRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateAlertRecordIssueRequest implement json.Marshaler.
func (m *UpdateAlertRecordIssueRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateAlertRecordIssueResponse implement json.Marshaler.
func (m *UpdateAlertRecordIssueResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateAlertRecordIssueResponse implement json.Marshaler.
func (m *UpdateAlertRecordIssueResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DashboardPreviewRequest implement json.Marshaler.
func (m *DashboardPreviewRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DashboardPreviewRequest implement json.Marshaler.
func (m *DashboardPreviewRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DashboardPreviewResponse implement json.Marshaler.
func (m *DashboardPreviewResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DashboardPreviewResponse implement json.Marshaler.
func (m *DashboardPreviewResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}
