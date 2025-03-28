/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.520 V17.9.0; 5G System; Network Data Analytics Services.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.520/
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import (
	"time"
)

// Contains historical data related to an analytics subscription.
type HistoricalData struct {
	// string with format 'date-time' as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty" yaml:"startTime" bson:"startTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTime *time.Time `json:"endTime,omitempty" yaml:"endTime" bson:"endTime,omitempty"`
	// Information about subscriptions with the data sources.
	SubsWithSources []SpecificDataSubscription `json:"subsWithSources,omitempty" yaml:"subsWithSources" bson:"subsWithSources,omitempty"`
	// Historical data related to the analytics.
	Data []DataNotification `json:"data" yaml:"data" bson:"data,omitempty"`
}
