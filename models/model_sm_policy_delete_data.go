/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.512 V17.11.0; 5G System; Session Management Policy Control Service.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.512/
 *
 * API version: 1.2.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Contains the parameters to be sent to the PCF when an individual SM policy is deleted.
type SmPolicyDeleteData struct {
	UserLocationInfo *UserLocation `json:"userLocationInfo,omitempty" yaml:"userLocationInfo" bson:"userLocationInfo,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UeTimeZone     string     `json:"ueTimeZone,omitempty" yaml:"ueTimeZone" bson:"ueTimeZone,omitempty"`
	ServingNetwork *PlmnIdNid `json:"servingNetwork,omitempty" yaml:"servingNetwork" bson:"servingNetwork,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	UserLocationInfoTime *time.Time `json:"userLocationInfoTime,omitempty" yaml:"userLocationInfoTime" bson:"userLocationInfoTime,omitempty"`
	// Contains the RAN and/or NAS release cause.
	RanNasRelCauses []RanNasRelCause `json:"ranNasRelCauses,omitempty" yaml:"ranNasRelCauses" bson:"ranNasRelCauses,omitempty"`
	// Contains the usage report
	AccuUsageReports []AccuUsageReport  `json:"accuUsageReports,omitempty" yaml:"accuUsageReports" bson:"accuUsageReports,omitempty"`
	PduSessRelCause  PduSessionRelCause `json:"pduSessRelCause,omitempty" yaml:"pduSessRelCause" bson:"pduSessRelCause,omitempty"`
}
