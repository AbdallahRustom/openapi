/*
 * Nchf_ConvergedCharging
 *
 * ConvergedCharging Service    © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.
 *
 * Source file: 3GPP TS 32.291 V17.9.0: Telecommunication management; Charging management;  5G system, charging service; Stage 3.
 * Url: http://www.3gpp.org/ftp/Specs/archive/32_series/32.291/
 *
 * API version: 3.1.6
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type FinalUnitIndication struct {
	FinalUnitAction           FinalUnitAction `json:"finalUnitAction" yaml:"finalUnitAction" bson:"finalUnitAction,omitempty"`
	RestrictionFilterRule     string          `json:"restrictionFilterRule,omitempty" yaml:"restrictionFilterRule" bson:"restrictionFilterRule,omitempty"`
	RestrictionFilterRuleList []string        `json:"restrictionFilterRuleList,omitempty" yaml:"restrictionFilterRuleList" bson:"restrictionFilterRuleList,omitempty"`
	FilterId                  string          `json:"filterId,omitempty" yaml:"filterId" bson:"filterId,omitempty"`
	FilterIdList              []string        `json:"filterIdList,omitempty" yaml:"filterIdList" bson:"filterIdList,omitempty"`
	RedirectServer            *RedirectServer `json:"redirectServer,omitempty" yaml:"redirectServer" bson:"redirectServer,omitempty"`
}
