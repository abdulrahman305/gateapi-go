/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Loan records
type UnifiedLoanRecord struct {
	// id
	Id int64 `json:"id,omitempty"`
	// type: borrow - borrow, repay - repay
	Type string `json:"type,omitempty"`
	// Repayment type: none - no repayment type, manual_repay - manual repayment, auto_repay - automatic repayment, cancel_auto_repay - automatic repayment after cancellation
	RepaymentType string `json:"repayment_type,omitempty"`
	// Loan type, returned when querying loan records. manual_borrow - Manual repayment , auto_borrow - Automatic repayment
	BorrowType string `json:"borrow_type,omitempty"`
	// Currency pair
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Currency
	Currency string `json:"currency,omitempty"`
	// The amount of lending or repaying
	Amount string `json:"amount,omitempty"`
	// Created time
	CreateTime int64 `json:"create_time,omitempty"`
}
