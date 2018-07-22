// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

// The Transformation type collects information for the transformation stage in the pipeline.
type Transformation struct {
	Staged
}

// Function NewTransformation creates a new, initialized instance.
// Returns nil on error.
func NewTransformation(payload string) *Transformation {
	var instance *Transformation = nil

	staged := NewStaged(payload)
	if staged != nil {
		instance = &Transformation{*staged}
	}

	return instance
}
