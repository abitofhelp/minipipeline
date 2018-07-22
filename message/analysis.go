// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

import (
	"go.uber.org/zap"
)

// A fast logging package
var logger zap.Logger

// The Analysis type collects information for the analysis stage in the pipeline.
type Analysis struct {
	Staged
}

// Function NewAnalysis creates a new, initialized instance.
// Returns nil on error.
func NewAnalysis(payload string) *Analysis {
	var instance *Analysis = nil

	staged := NewStaged(payload)
	if staged != nil {
		instance = &Analysis{*staged}
	}

	return instance
}
