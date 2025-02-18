// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trace.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetSpansRequest) Validate() error {
	if this.TraceID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TraceID", fmt.Errorf(`value '%v' must not be an empty string`, this.TraceID))
	}
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetTracesRequest) Validate() error {
	if this.ScopeID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ScopeID", fmt.Errorf(`value '%v' must not be an empty string`, this.ScopeID))
	}
	return nil
}
func (this *GetSpansResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetTracesResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *Span) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Trace) Validate() error {
	return nil
}
