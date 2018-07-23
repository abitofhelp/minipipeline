// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

import (
	"github.com/abitofhelp/minipipeline/stage"
)

// Interface IStage defines the methods that must be implemented for a
// message flowing through a stage in the pipeline.
type IStage interface {
	IBasic

	// Method Stage returns the stage in the pipeline that the message is related to.
	Stage() stage.Stages
}
