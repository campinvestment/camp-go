// Copyright 2025 CAMP Investment Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package camp

// TradesOpen contains the open asks and bids in the market place.
type TradesOpen struct {
	Asks []Trade `json:"asks"`
	Bids []Trade `json:"bids"`
}
