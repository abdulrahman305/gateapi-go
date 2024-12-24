/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type FuturesAccount struct {
	// total is the balance after the user's accumulated deposit, withdraw, profit and loss (including realized profit and loss, fund, fee and referral rebate), excluding unrealized profit and loss.  total = SUM(history_dnw, history_pnl, history_fee, history_refr, history_fund)
	Total string `json:"total,omitempty"`
	// Unrealized PNL
	UnrealisedPnl string `json:"unrealised_pnl,omitempty"`
	// Position margin
	PositionMargin string `json:"position_margin,omitempty"`
	// Order margin of unfinished orders
	OrderMargin string `json:"order_margin,omitempty"`
	// The available balance for transferring or trading(including bonus.  Bonus can't be be withdrawn. The transfer amount needs to deduct the bonus)
	Available string `json:"available,omitempty"`
	// POINT amount
	Point string `json:"point,omitempty"`
	// Settle currency
	Currency string `json:"currency,omitempty"`
	// Whether dual mode is enabled
	InDualMode bool `json:"in_dual_mode,omitempty"`
	// Whether portfolio margin account mode is enabled
	EnableCredit bool `json:"enable_credit,omitempty"`
	// Initial margin position, applicable to the portfolio margin account model
	PositionInitialMargin string `json:"position_initial_margin,omitempty"`
	// The maintenance deposit occupied by the position is suitable for the new classic account margin model and unified account model
	MaintenanceMargin string `json:"maintenance_margin,omitempty"`
	// Perpetual Contract Bonus
	Bonus string `json:"bonus,omitempty"`
	// Classic account margin mode, true-new mode, false-old mode
	EnableEvolvedClassic bool `json:"enable_evolved_classic,omitempty"`
	// Full -warehouse hanging order deposit, suitable for the new classic account margin model
	CrossOrderMargin string `json:"cross_order_margin,omitempty"`
	// The initial security deposit of the full warehouse is suitable for the new classic account margin model
	CrossInitialMargin string `json:"cross_initial_margin,omitempty"`
	// Maintain deposit in full warehouse, suitable for new classic account margin models
	CrossMaintenanceMargin string `json:"cross_maintenance_margin,omitempty"`
	// The full warehouse does not achieve profit and loss, suitable for the new classic account margin model
	CrossUnrealisedPnl string `json:"cross_unrealised_pnl,omitempty"`
	// Full warehouse available amount, suitable for the new classic account margin model
	CrossAvailable string `json:"cross_available,omitempty"`
	// Ware -position margin, suitable for the new classic account margin model
	IsolatedPositionMargin string `json:"isolated_position_margin,omitempty"`
	// Whether to open a new two-way position mode
	EnableNewDualMode bool `json:"enable_new_dual_mode,omitempty"`
	// Margin mode, 0-classic margin mode, 1-cross-currency margin mode, 2-combined margin mode
	MarginMode int32                 `json:"margin_mode,omitempty"`
	History    FuturesAccountHistory `json:"history,omitempty"`
}
