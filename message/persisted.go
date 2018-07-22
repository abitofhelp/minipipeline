// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

// The Persisted type collects information for the persisted stage in the pipeline.
type Persisted struct {
	Staged
}

// Function NewPersisted creates a new, initialized instance.
// Returns nil on error.
func NewPersisted(payload string) *Persisted {
	var instance *Persisted = nil

	staged := NewStaged(payload)
	if staged != nil {
		instance = &Persisted{*staged}
	}

	return instance
}
