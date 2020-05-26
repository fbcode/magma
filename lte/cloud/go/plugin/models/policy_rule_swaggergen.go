// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PolicyRule policy rule
// swagger:model policy_rule
type PolicyRule struct {

	// Subscribers which have been assigned this policy not as part of a base name
	AssignedSubscribers []SubscriberID `json:"assigned_subscribers,omitempty"`

	// flow list
	// Required: true
	FlowList []*FlowDescription `json:"flow_list"`

	// id
	// Required: true
	ID PolicyID `json:"id"`

	// monitoring key
	MonitoringKey string `json:"monitoring_key,omitempty"`

	// priority
	// Required: true
	Priority *uint32 `json:"priority"`

	// qos
	Qos *FlowQos `json:"qos,omitempty"`

	// rating group
	RatingGroup uint32 `json:"rating_group,omitempty"`

	// redirect
	Redirect *RedirectInformation `json:"redirect,omitempty"`

	// tracking type
	// Enum: [ONLY_OCS ONLY_PCRF OCS_AND_PCRF NO_TRACKING]
	TrackingType string `json:"tracking_type,omitempty"`
}

// Validate validates this policy rule
func (m *PolicyRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedSubscribers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrackingType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyRule) validateAssignedSubscribers(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignedSubscribers) { // not required
		return nil
	}

	for i := 0; i < len(m.AssignedSubscribers); i++ {

		if err := m.AssignedSubscribers[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assigned_subscribers" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PolicyRule) validateFlowList(formats strfmt.Registry) error {

	if err := validate.Required("flow_list", "body", m.FlowList); err != nil {
		return err
	}

	for i := 0; i < len(m.FlowList); i++ {
		if swag.IsZero(m.FlowList[i]) { // not required
			continue
		}

		if m.FlowList[i] != nil {
			if err := m.FlowList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("flow_list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PolicyRule) validateID(formats strfmt.Registry) error {

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *PolicyRule) validatePriority(formats strfmt.Registry) error {

	if err := validate.Required("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *PolicyRule) validateQos(formats strfmt.Registry) error {

	if swag.IsZero(m.Qos) { // not required
		return nil
	}

	if m.Qos != nil {
		if err := m.Qos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos")
			}
			return err
		}
	}

	return nil
}

func (m *PolicyRule) validateRedirect(formats strfmt.Registry) error {

	if swag.IsZero(m.Redirect) { // not required
		return nil
	}

	if m.Redirect != nil {
		if err := m.Redirect.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redirect")
			}
			return err
		}
	}

	return nil
}

var policyRuleTypeTrackingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ONLY_OCS","ONLY_PCRF","OCS_AND_PCRF","NO_TRACKING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyRuleTypeTrackingTypePropEnum = append(policyRuleTypeTrackingTypePropEnum, v)
	}
}

const (

	// PolicyRuleTrackingTypeONLYOCS captures enum value "ONLY_OCS"
	PolicyRuleTrackingTypeONLYOCS string = "ONLY_OCS"

	// PolicyRuleTrackingTypeONLYPCRF captures enum value "ONLY_PCRF"
	PolicyRuleTrackingTypeONLYPCRF string = "ONLY_PCRF"

	// PolicyRuleTrackingTypeOCSANDPCRF captures enum value "OCS_AND_PCRF"
	PolicyRuleTrackingTypeOCSANDPCRF string = "OCS_AND_PCRF"

	// PolicyRuleTrackingTypeNOTRACKING captures enum value "NO_TRACKING"
	PolicyRuleTrackingTypeNOTRACKING string = "NO_TRACKING"
)

// prop value enum
func (m *PolicyRule) validateTrackingTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, policyRuleTypeTrackingTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PolicyRule) validateTrackingType(formats strfmt.Registry) error {

	if swag.IsZero(m.TrackingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrackingTypeEnum("tracking_type", "body", m.TrackingType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRule) UnmarshalBinary(b []byte) error {
	var res PolicyRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}