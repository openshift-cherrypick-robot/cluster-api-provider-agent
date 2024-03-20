// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// FinalizingStage Cluster finalizing stage managed by controller
//
// swagger:model finalizing-stage
type FinalizingStage string

func NewFinalizingStage(value FinalizingStage) *FinalizingStage {
	return &value
}

// Pointer returns a pointer to a freshly-allocated FinalizingStage.
func (m FinalizingStage) Pointer() *FinalizingStage {
	return &m
}

const (

	// FinalizingStageWaitingForFinalizing captures enum value "Waiting for finalizing"
	FinalizingStageWaitingForFinalizing FinalizingStage = "Waiting for finalizing"

	// FinalizingStageWaitingForClusterOperators captures enum value "Waiting for cluster operators"
	FinalizingStageWaitingForClusterOperators FinalizingStage = "Waiting for cluster operators"

	// FinalizingStageAddingRouterCa captures enum value "Adding router ca"
	FinalizingStageAddingRouterCa FinalizingStage = "Adding router ca"

	// FinalizingStageWaitingForOlmOperators captures enum value "Waiting for olm operators"
	FinalizingStageWaitingForOlmOperators FinalizingStage = "Waiting for olm operators"

	// FinalizingStageApplyingManifests captures enum value "Applying manifests"
	FinalizingStageApplyingManifests FinalizingStage = "Applying manifests"

	// FinalizingStageWaitingForOlmOperatorsCsvInitialization captures enum value "Waiting for olm operators csv initialization"
	FinalizingStageWaitingForOlmOperatorsCsvInitialization FinalizingStage = "Waiting for olm operators csv initialization"

	// FinalizingStageWaitingForOlmOperatorsCsv captures enum value "Waiting for olm operators csv"
	FinalizingStageWaitingForOlmOperatorsCsv FinalizingStage = "Waiting for olm operators csv"

	// FinalizingStageDone captures enum value "Done"
	FinalizingStageDone FinalizingStage = "Done"
)

// for schema
var finalizingStageEnum []interface{}

func init() {
	var res []FinalizingStage
	if err := json.Unmarshal([]byte(`["Waiting for finalizing","Waiting for cluster operators","Adding router ca","Waiting for olm operators","Applying manifests","Waiting for olm operators csv initialization","Waiting for olm operators csv","Done"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		finalizingStageEnum = append(finalizingStageEnum, v)
	}
}

func (m FinalizingStage) validateFinalizingStageEnum(path, location string, value FinalizingStage) error {
	if err := validate.EnumCase(path, location, value, finalizingStageEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this finalizing stage
func (m FinalizingStage) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateFinalizingStageEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this finalizing stage based on context it is used
func (m FinalizingStage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}