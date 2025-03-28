/*
 * Nhss_imsSDM
 *
 * Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.562 HSS Services for interworking with IMS, version 17.6.0
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.562/
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Reference Location Information for the user in fixed access networks
type ReferenceLocationInformation struct {
	AccessType  string `json:"accessType,omitempty" yaml:"accessType" bson:"accessType,omitempty"`
	AccessInfo  string `json:"accessInfo,omitempty" yaml:"accessInfo" bson:"accessInfo,omitempty"`
	AccessValue string `json:"accessValue,omitempty" yaml:"accessValue" bson:"accessValue,omitempty"`
}
