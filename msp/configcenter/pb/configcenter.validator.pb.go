// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: configcenter.proto

package pb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetGroupRequest) Validate() error {
	if this.TenantID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantID", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantID))
	}
	if !(this.PageNum > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageNum", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageNum))
	}
	if !(this.PageSize > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '0'`, this.PageSize))
	}
	return nil
}
func (this *GetGroupResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *Groups) Validate() error {
	return nil
}
func (this *GetGroupPropertiesRequest) Validate() error {
	if this.TenantID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantID", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantID))
	}
	if this.GroupID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("GroupID", fmt.Errorf(`value '%v' must not be an empty string`, this.GroupID))
	}
	return nil
}
func (this *GetGroupPropertiesResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *SaveGroupPropertiesRequest) Validate() error {
	if this.TenantID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TenantID", fmt.Errorf(`value '%v' must not be an empty string`, this.TenantID))
	}
	if this.GroupID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("GroupID", fmt.Errorf(`value '%v' must not be an empty string`, this.GroupID))
	}
	for _, item := range this.Properties {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Properties", err)
			}
		}
	}
	return nil
}
func (this *SaveGroupPropertiesResponse) Validate() error {
	return nil
}
func (this *GroupProperties) Validate() error {
	for _, item := range this.Properties {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Properties", err)
			}
		}
	}
	return nil
}
func (this *Property) Validate() error {
	return nil
}
