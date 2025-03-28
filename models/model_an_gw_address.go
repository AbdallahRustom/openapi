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

// Describes the address of the access network gateway control node.
type AnGwAddress struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	AnGwIpv4Addr string `json:"anGwIpv4Addr,omitempty" yaml:"anGwIpv4Addr" bson:"anGwIpv4Addr,omitempty"`
	AnGwIpv6Addr string `json:"anGwIpv6Addr,omitempty" yaml:"anGwIpv6Addr" bson:"anGwIpv6Addr,omitempty"`
}
