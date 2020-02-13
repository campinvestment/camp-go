// Copyright 2019 Tokenomy Technologies Pte. Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package v2

import "github.com/tokenomy/tokenomy-go"

//
// OpenOrders contains the open ask and bid orders in the market place.
//
type OpenOrders struct {
	Asks []tokenomy.Order `json:"asks"`
	Bids []tokenomy.Order `json:"bids"`
}
