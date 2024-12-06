/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type PartnerSubList struct {
	// Total
	Total int64 `json:"total,omitempty"`
	// Subordinate list
	List []PartnerSub `json:"list,omitempty"`
}
