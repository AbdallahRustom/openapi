/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.514 V17.9.0; 5G System; Policy Authorization Service; Stage 3.
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.514/
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the new TSC user plane node information and may contain the DS-TT port and/or NW-TT port management information.
type PduSessionTsnBridge struct {
	TsnBridgeInfo       *TsnBridgeInfo             `json:"tsnBridgeInfo" yaml:"tsnBridgeInfo" bson:"tsnBridgeInfo,omitempty"`
	TsnBridgeManCont    *BridgeManagementContainer `json:"tsnBridgeManCont,omitempty" yaml:"tsnBridgeManCont" bson:"tsnBridgeManCont,omitempty"`
	TsnPortManContDstt  *PortManagementContainer   `json:"tsnPortManContDstt,omitempty" yaml:"tsnPortManContDstt" bson:"tsnPortManContDstt,omitempty"`
	TsnPortManContNwtts []PortManagementContainer  `json:"tsnPortManContNwtts,omitempty" yaml:"tsnPortManContNwtts" bson:"tsnPortManContNwtts,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	UeIpv4Addr string `json:"ueIpv4Addr,omitempty" yaml:"ueIpv4Addr" bson:"ueIpv4Addr,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn    string  `json:"dnn,omitempty" yaml:"dnn" bson:"dnn,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty" yaml:"snssai" bson:"snssai,omitempty"`
	// IPv4 address domain identifier.
	IpDomain         string `json:"ipDomain,omitempty" yaml:"ipDomain" bson:"ipDomain,omitempty"`
	UeIpv6AddrPrefix string `json:"ueIpv6AddrPrefix,omitempty" yaml:"ueIpv6AddrPrefix" bson:"ueIpv6AddrPrefix,omitempty"`
}
