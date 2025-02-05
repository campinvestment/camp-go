// Copyright 2025 CAMP Investment Technologies Ltd. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package camp

// Asset contains basic information of an asset
type Asset struct {
	Name             string `json:"name"`
	Logo             string `json:"logo"`
	IsWithoutNetwork bool   `json:"is_without_network"`
}
