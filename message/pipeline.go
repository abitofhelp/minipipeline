// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

// Package message provides the messages for the pipeline.
package message

// The Pipeline type collects information from each stage in the pipeline,
// and for the overall process.
type Pipeline struct {
	Basic

	Intake
	Analysis
	Transformation
	Validation
	Persisted
}

// Function NewPipeline creates a new, initialized instance.
// Returns nil on error.
func NewPipeline(payload string) *Pipeline {
	var instance *Pipeline = nil

	// Basic message is for the pipeline, overall, and not a single stage...
	basic := NewBasic(payload)

	// Stages...
	intake := NewIntake(payload)
	analysis := NewAnalysis(payload)
	transformation := NewTransformation(payload)
	validation := NewValidation(payload)
	persisted := NewPersisted(payload)

	if ContainsNil(basic, intake, analysis, transformation, validation, persisted) {
		logger.Error("at least one of the pipeline stages was nil, so the pipeline cannot be built.")
		return nil
	}

	instance = &Pipeline{
		*basic,
		*intake,
		*analysis,
		*transformation,
		*validation,
		*persisted,
	}

	return instance
}
