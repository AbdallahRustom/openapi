/*
 * Npcf_EventExposure
 *
 * PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.523 V17.7.0; 5G System; Policy Control Event Exposure Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.523/
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents a combination of S-NSSAI and DNN(s).
type SnssaiDnnCombination struct {
	Snssai *Snssai  `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	Dnns   []string `json:"dnns,omitempty" yaml:"dnns" bson:"dnns,omitempty"`
}
