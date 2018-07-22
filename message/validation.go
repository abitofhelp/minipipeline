// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

// The Validation type collects information for the validation stage in the pipeline.
type Validation struct {
	Staged
}

// Function NewValidation creates a new, initialized instance.
// Returns nil on error.
func NewValidation(payload string) *Validation {
	var instance *Validation = nil

	staged := NewStaged(payload)
	if staged != nil {
		instance = &Validation{*staged}
	}

	return instance
}
