// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	utils "github.com/aeternity/aepp-sdk-go/v5/utils"
)

// ChannelDepositTx channel deposit tx
// swagger:model ChannelDepositTx
type ChannelDepositTx struct {

	// amount
	// Required: true
	Amount utils.BigInt `json:"amount"`

	// channel id
	// Required: true
	ChannelID *string `json:"channel_id"`

	// fee
	// Required: true
	Fee utils.BigInt `json:"fee"`

	// from id
	// Required: true
	FromID *string `json:"from_id"`

	// nonce
	// Required: true
	Nonce *uint64 `json:"nonce"`

	// Channel's next round
	// Required: true
	Round *uint64 `json:"round"`

	// Root hash of the channel's internal state tree after the deposit had been applied to it
	// Required: true
	StateHash *string `json:"state_hash"`

	// ttl
	TTL uint64 `json:"ttl,omitempty"`
}

// Validate validates this channel deposit tx
func (m *ChannelDepositTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFromID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonce(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRound(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStateHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelDepositTx) validateAmount(formats strfmt.Registry) error {

	if err := m.Amount.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("amount")
		}
		return err
	}

	return nil
}

func (m *ChannelDepositTx) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("channel_id", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelDepositTx) validateFee(formats strfmt.Registry) error {

	if err := m.Fee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("fee")
		}
		return err
	}

	return nil
}

func (m *ChannelDepositTx) validateFromID(formats strfmt.Registry) error {

	if err := validate.Required("from_id", "body", m.FromID); err != nil {
		return err
	}

	return nil
}

func (m *ChannelDepositTx) validateNonce(formats strfmt.Registry) error {

	if err := validate.Required("nonce", "body", m.Nonce); err != nil {
		return err
	}

	return nil
}

func (m *ChannelDepositTx) validateRound(formats strfmt.Registry) error {

	if err := validate.Required("round", "body", m.Round); err != nil {
		return err
	}

	return nil
}

func (m *ChannelDepositTx) validateStateHash(formats strfmt.Registry) error {

	if err := validate.Required("state_hash", "body", m.StateHash); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelDepositTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelDepositTx) UnmarshalBinary(b []byte) error {
	var res ChannelDepositTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}