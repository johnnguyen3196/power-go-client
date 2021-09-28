// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VPNConnectionCreate v p n connection create
// swagger:model VPNConnectionCreate
type VPNConnectionCreate struct {

	// connection state of the VPN Connection
	// Required: true
	ConnectionState *bool `json:"connectionState"`

	// unique identifier of IKEPolicy selected for this VPNConnection
	// Required: true
	IkePolicy *string `json:"ikePolicy"`

	// unique identifier of IPSecPolicy selected for this VPNConnection
	// Required: true
	IPSecPolicy *string `json:"ipSecPolicy"`

	// Mode used by this VPNConnection, either policy-based, or route-based, this attribute is set at the creation and cannot be updated later.
	// Required: true
	// Enum: [policy route]
	Mode *string `json:"mode"`

	// VPN Connection name
	// Required: true
	Name *string `json:"name"`

	// an array of network IDs to attach to this VPNConnection
	// Required: true
	Networks []string `json:"networks"`

	// peer gateway address
	// Required: true
	// Format: ipv4
	PeerGatewayAddress PeerGatewayAddress `json:"peerGatewayAddress"`

	// an array of strings containing CIDR of peer subnets
	// Required: true
	PeerSubnets []string `json:"peerSubnets"`
}

// Validate validates this v p n connection create
func (m *VPNConnectionCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIkePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPSecPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerGatewayAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerSubnets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VPNConnectionCreate) validateConnectionState(formats strfmt.Registry) error {

	if err := validate.Required("connectionState", "body", m.ConnectionState); err != nil {
		return err
	}

	return nil
}

func (m *VPNConnectionCreate) validateIkePolicy(formats strfmt.Registry) error {

	if err := validate.Required("ikePolicy", "body", m.IkePolicy); err != nil {
		return err
	}

	return nil
}

func (m *VPNConnectionCreate) validateIPSecPolicy(formats strfmt.Registry) error {

	if err := validate.Required("ipSecPolicy", "body", m.IPSecPolicy); err != nil {
		return err
	}

	return nil
}

var vPNConnectionCreateTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["policy","route"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vPNConnectionCreateTypeModePropEnum = append(vPNConnectionCreateTypeModePropEnum, v)
	}
}

const (

	// VPNConnectionCreateModePolicy captures enum value "policy"
	VPNConnectionCreateModePolicy string = "policy"

	// VPNConnectionCreateModeRoute captures enum value "route"
	VPNConnectionCreateModeRoute string = "route"
)

// prop value enum
func (m *VPNConnectionCreate) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, vPNConnectionCreateTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VPNConnectionCreate) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("mode", "body", m.Mode); err != nil {
		return err
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *VPNConnectionCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VPNConnectionCreate) validateNetworks(formats strfmt.Registry) error {

	if err := validate.Required("networks", "body", m.Networks); err != nil {
		return err
	}

	return nil
}

func (m *VPNConnectionCreate) validatePeerGatewayAddress(formats strfmt.Registry) error {

	if err := m.PeerGatewayAddress.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("peerGatewayAddress")
		}
		return err
	}

	return nil
}

func (m *VPNConnectionCreate) validatePeerSubnets(formats strfmt.Registry) error {

	if err := validate.Required("peerSubnets", "body", m.PeerSubnets); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VPNConnectionCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VPNConnectionCreate) UnmarshalBinary(b []byte) error {
	var res VPNConnectionCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
