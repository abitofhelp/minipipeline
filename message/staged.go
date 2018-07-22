// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

import (
	"github.com/abitofhelp/minipipeline/stage"
)

// Type Staged implements the IStage interface to provide the minimum
// set of information a message transiting the pipeline.
type Staged struct {
	Basic

	stage stage.Stage
}

// Method Stage returns the stage in the pipeline that the message is related to.
func (s *Staged) Stage() stage.Stage {
	return s.stage
}

// Function NewStaged creates a new, initialized instance.
// Returns nil on error.
func NewStaged(payload string) *Staged {
	var staged *Staged = nil

	basic := NewBasic(payload)
	if basic != nil {
		staged = &Staged{*basic, stage.Unknown}
	}
	return staged
}
