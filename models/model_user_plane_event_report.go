/*
 * 3gpp-as-session-with-qos
 *
 * API for setting us an AS session with required QoS.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.122 V17.9.0 T8 reference point for Northbound APIs
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.122/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Represents an event report for user plane.
type UserPlaneEventReport struct {
	Event            UserPlaneEvent    `json:"event" yaml:"event" bson:"event,omitempty"`
	AccumulatedUsage *AccumulatedUsage `json:"accumulatedUsage,omitempty" yaml:"accumulatedUsage" bson:"accumulatedUsage,omitempty"`
	// Identifies the affected flows that were sent during event subscription. It might be omitted when the reported event applies to all the flows sent during the subscription.
	FlowIds []int32 `json:"flowIds,omitempty" yaml:"flowIds" bson:"flowIds,omitempty"`
	// The currently applied QoS reference. Applicable for event QOS_NOT_GUARANTEED or SUCCESSFUL_RESOURCES_ALLOCATION.
	AppliedQosRef string     `json:"appliedQosRef,omitempty" yaml:"appliedQosRef" bson:"appliedQosRef,omitempty"`
	PlmnId        *PlmnIdNid `json:"plmnId,omitempty" yaml:"plmnId" bson:"plmnId,omitempty"`
	// Contains the QoS Monitoring Reporting information
	QosMonReports []QosMonitoringReport `json:"qosMonReports,omitempty" yaml:"qosMonReports" bson:"qosMonReports,omitempty"`
	RatType       RatType               `json:"ratType,omitempty" yaml:"ratType" bson:"ratType,omitempty"`
}
