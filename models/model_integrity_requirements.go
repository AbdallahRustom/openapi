/*
 * LMF Location
 *
 * LMF Location Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.572 V17.9.0; 5G System; Location Management Services; Stage 3
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.572/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// integrity requirements.
type IntegrityRequirements struct {
	// Contains the time-to-alert
	TimeToAlert int32 `json:"timeToAlert,omitempty" yaml:"timeToAlert" bson:"timeToAlert,omitempty"`
	// Contains the target integrity risk
	TargetIntegrityRisk int32       `json:"targetIntegrityRisk,omitempty" yaml:"targetIntegrityRisk" bson:"targetIntegrityRisk,omitempty"`
	AlertLimit          *AlertLimit `json:"alertLimit,omitempty" yaml:"alertLimit" bson:"alertLimit,omitempty"`
}
