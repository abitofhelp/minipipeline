// Copyright (c) 2018 A Bit of Help, Inc. - All Rights Reserved, Worldwide.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.

package main

/*
import (
	"errors"
	"github.com/abitofhelp/minipipeline/stage/intake"
	. "github.com/abitofhelp/minipipeline/stage"
)

// Default stage implementation to use.
const (
	// Constant kDefaultName is the default stage name to use.
	defaultName = "intake"
)

var (
	// Variable instance is a singleton instance of the factory.
	instance *Factory
)

// Type Factory provides for constructing concrete instances of IStep.
type Factory struct {
	// Variable constructors is a map of constructors for the concrete steps.
	constructors map[string]func() IStep
}

// Method New creates an instance of the default kind of stage.
func (f *Factory) New() IStep {
	ret, _ := f.Create(defaultName)
	return ret
}

// Method Create returns an instance of a specific kind of stage using its name.
func (f *Factory) Create(name string) (IStep, error) {
	constructor, ok := f.constructors[name]
	if !ok {
		return nil, errors.New("stage not found")
	}
	return constructor(), nil
}

// Method Register maintains a collection of known kinds of steps and their constructors.
func (f *Factory) Register(name string, constructor func() IStep) {
	f.constructors[name] = constructor
}

// Function NewFactory creates a singleton instance of the factory for IStep.
func NewFactory() *Factory {
	if instance == nil {
		instance = &Factory{constructors: map[string]func() IStep{}}
	}

	return instance
}

// Function init registers each kind of stage that can exist in the pipeline.
func init() {
	NewFactory().Register("intake", func() IStep { return intake.New() })
}
*/
