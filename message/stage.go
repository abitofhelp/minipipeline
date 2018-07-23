// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

import (
	"github.com/abitofhelp/minipipeline/stage"
	"time"
)

// Type Stage implements the IStage interface to provide the required data.
type Stage struct {
	Basic

	//  Field stage indicates the kind of stage.
	stage stage.Stages
}

// Method Stage returns the stage in the pipeline that the message is related to.
func (s *Stage) Stage() stage.Stages {
	return s.stage
}

// Function NewStage creates a new, initialized instance with a payload and stage.
// Parameter stage is the kind of stage (i.e. Intake).
// Parameter payload is the string being passed in the message.
// Returns nil on error.
func NewStage(stage stage.Stages, payload string, departedUtc time.Time) *IStage {
	var step *Stage = nil

	basic := NewBasic(payload, departedUtc)
	if basic != nil {
		step = &Stage{*basic, stage}
	}
	return step
}
