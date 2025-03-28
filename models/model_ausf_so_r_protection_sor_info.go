/*
 * Nausf_SoRProtection Service
 *
 * AUSF SoR Protection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.
 *
 * Source file: 3GPP TS 29.509 V17.6.0; 5G System; Authentication Server Services
 * Url: https://www.3gpp.org/ftp/Specs/archive/29_series/29.509
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// Contains the Steering Information.
type AusfSoRProtectionSorInfo struct {
	SteeringContainer *AusfSoRProtectionSteeringContainer `json:"steeringContainer,omitempty" yaml:"steeringContainer" bson:"steeringContainer,omitempty"`
	// Contains indication whether the acknowledgement from UE is needed.
	AckInd bool `json:"ackInd" yaml:"ackInd" bson:"ackInd,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	SorHeader string `json:"sorHeader,omitempty" yaml:"sorHeader" bson:"sorHeader,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	SorTransparentInfo string `json:"sorTransparentInfo,omitempty" yaml:"sorTransparentInfo" bson:"sorTransparentInfo,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures string `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures,omitempty"`
}
