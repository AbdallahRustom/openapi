/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Dispersion Area
type DispersionArea struct {
	TaiList  []Tai  `json:"taiList,omitempty" yaml:"taiList" bson:"taiList,omitempty"`
	NcgiList []Ncgi `json:"ncgiList,omitempty" yaml:"ncgiList" bson:"ncgiList,omitempty"`
	EcgiList []Ecgi `json:"ecgiList,omitempty" yaml:"ecgiList" bson:"ecgiList,omitempty"`
	N3gaInd  bool   `json:"n3gaInd,omitempty" yaml:"n3gaInd" bson:"n3gaInd,omitempty"`
}
