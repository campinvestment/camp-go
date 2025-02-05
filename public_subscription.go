// Copyright 2025 CAMP Investment Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package camp

// PublicSubscription contains list of pairs that currently subscribed for
// each topic: "depths", "ticker", and "trades".
type PublicSubscription struct {
	Depths    []string `json:"depths"`
	Ticker    []string `json:"ticker"`
	Trades    []string `json:"trades"`
	Summaries bool     `json:"summaries"`
}
