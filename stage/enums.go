// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package stage

// The Stage enumeration values.
const (
	Unknown Stage = iota
	Intake
	Analysis
	Transformation
	Validation
	Delivery
)

// The Stage enumeration type.
type Stage uint8

// The function String() provides the name for a stage.
// It returns the name of a stage or "Undefined" if the stage does not exist.
func (stage Stage) String() string {
	// Declare an array of strings.
	// ... operator counts how many
	// items in the array (6).
	names := [...]string{
		"Unknown",
		"Intake",
		"Analysis",
		"Transformation",
		"Validation",
		"Delivery"}
	// → `stage`: It's one of the
	// values of Stage constants.
	// If the constant is Intake,
	// then stage is 0.
	// Let's prevent panicking in case
	// `stage` is out of range of Stage.
	if stage < Unknown || stage > Delivery {
		return "Undefined"
	}
	// Return the name of a Stage
	// constant from the names array
	// above.
	return names[stage]
}
