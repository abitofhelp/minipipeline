// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

import (
	"github.com/abitofhelp/minipipeline/stage"
)

// The Intake type collects information for the intake stage in the pipeline.
type Intake struct {
	Staged
}

// Function NewIntake creates a new, initialized instance.
// Returns nil on error.
func NewIntake(payload string) *Intake {
	var instance *Intake = nil

	staged := NewStaged(payload)
	if staged != nil {
		staged.stage = stage.Intake
		instance = &Intake{*staged}
	}

	return instance
}
