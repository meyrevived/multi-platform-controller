// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TokenRequest token request
//
// swagger:model TokenRequest
type TokenRequest struct {

	// The refresh token to request the new Access Token
	// Required: true
	RefreshToken *string `json:"refreshToken"`

	// Source type of the token request (web or cli)
	// Required: true
	// Enum: ["web","cli"]
	Source *string `json:"source"`
}

// Validate validates this token request
func (m *TokenRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRefreshToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenRequest) validateRefreshToken(formats strfmt.Registry) error {

	if err := validate.Required("refreshToken", "body", m.RefreshToken); err != nil {
		return err
	}

	return nil
}

var tokenRequestTypeSourcePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["web","cli"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tokenRequestTypeSourcePropEnum = append(tokenRequestTypeSourcePropEnum, v)
	}
}

const (

	// TokenRequestSourceWeb captures enum value "web"
	TokenRequestSourceWeb string = "web"

	// TokenRequestSourceCli captures enum value "cli"
	TokenRequestSourceCli string = "cli"
)

// prop value enum
func (m *TokenRequest) validateSourceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tokenRequestTypeSourcePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TokenRequest) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	// value enum
	if err := m.validateSourceEnum("source", "body", *m.Source); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this token request based on context it is used
func (m *TokenRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TokenRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenRequest) UnmarshalBinary(b []byte) error {
	var res TokenRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}