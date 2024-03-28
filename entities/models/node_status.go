//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2024 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NodeStatus The definition of a backup node status response body
//
// swagger:model NodeStatus
type NodeStatus struct {

	// Weaviate batch statistics.
	BatchStats *BatchStats `json:"batchStats,omitempty"`

	// The gitHash of Weaviate.
	GitHash string `json:"gitHash,omitempty"`

	// The name of the node.
	Name string `json:"name,omitempty"`

	// The list of the shards with it's statistics.
	Shards []*NodeShardStatus `json:"shards"`

	// Weaviate overall statistics.
	Stats *NodeStats `json:"stats,omitempty"`

	// Node's status.
	// Enum: [HEALTHY UNHEALTHY UNAVAILABLE TIMEOUT]
	Status *string `json:"status,omitempty"`

	// The version of Weaviate.
	Version string `json:"version,omitempty"`
}

// Validate validates this node status
func (m *NodeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBatchStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeStatus) validateBatchStats(formats strfmt.Registry) error {
	if swag.IsZero(m.BatchStats) { // not required
		return nil
	}

	if m.BatchStats != nil {
		if err := m.BatchStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("batchStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("batchStats")
			}
			return err
		}
	}

	return nil
}

func (m *NodeStatus) validateShards(formats strfmt.Registry) error {
	if swag.IsZero(m.Shards) { // not required
		return nil
	}

	for i := 0; i < len(m.Shards); i++ {
		if swag.IsZero(m.Shards[i]) { // not required
			continue
		}

		if m.Shards[i] != nil {
			if err := m.Shards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NodeStatus) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

var nodeStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HEALTHY","UNHEALTHY","UNAVAILABLE","TIMEOUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeStatusTypeStatusPropEnum = append(nodeStatusTypeStatusPropEnum, v)
	}
}

const (

	// NodeStatusStatusHEALTHY captures enum value "HEALTHY"
	NodeStatusStatusHEALTHY string = "HEALTHY"

	// NodeStatusStatusUNHEALTHY captures enum value "UNHEALTHY"
	NodeStatusStatusUNHEALTHY string = "UNHEALTHY"

	// NodeStatusStatusUNAVAILABLE captures enum value "UNAVAILABLE"
	NodeStatusStatusUNAVAILABLE string = "UNAVAILABLE"

	// NodeStatusStatusTIMEOUT captures enum value "TIMEOUT"
	NodeStatusStatusTIMEOUT string = "TIMEOUT"
)

// prop value enum
func (m *NodeStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeStatusTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NodeStatus) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this node status based on the context it is used
func (m *NodeStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBatchStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeStatus) contextValidateBatchStats(ctx context.Context, formats strfmt.Registry) error {

	if m.BatchStats != nil {
		if err := m.BatchStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("batchStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("batchStats")
			}
			return err
		}
	}

	return nil
}

func (m *NodeStatus) contextValidateShards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Shards); i++ {

		if m.Shards[i] != nil {
			if err := m.Shards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("shards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("shards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NodeStatus) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {
		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeStatus) UnmarshalBinary(b []byte) error {
	var res NodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
