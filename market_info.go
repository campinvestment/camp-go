// Copyright 2025 CAMP Investment Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package camp

import "github.com/shuLhan/share/lib/math/big"

// MarketInfo contains the pair information to be consumed by public.
type MarketInfo struct {
	PriceMinimum  *big.Rat `json:"price_minimum"`
	AmountMinimum *big.Rat `json:"amount_minimum"`

	ID        string `json:"id"`
	Symbol    string `json:"symbol"` // Deprecated, use Pair field.
	Pair      string `json:"pair"`
	CoinAsset string `json:"coin_asset"`
	BaseAsset string `json:"base_asset"`

	PricePrecision  int `json:"price_precision"`
	AmountPrecision int `json:"amount_precision"`

	IsActive bool `json:"is_active"`
}
