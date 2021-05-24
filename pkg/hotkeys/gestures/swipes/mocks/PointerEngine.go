// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	swipes "github.com/echocrow/Mouser/pkg/hotkeys/gestures/swipes"
	vec "github.com/echocrow/Mouser/pkg/vec"
	mock "github.com/stretchr/testify/mock"
)

// PointerEngine is an autogenerated mock type for the PointerEngine type
type PointerEngine struct {
	mock.Mock
}

// GetPointerPos provides a mock function with given fields:
func (_m *PointerEngine) GetPointerPos() vec.Vec2D {
	ret := _m.Called()

	var r0 vec.Vec2D
	if rf, ok := ret.Get(0).(func() vec.Vec2D); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(vec.Vec2D)
	}

	return r0
}

// Init provides a mock function with given fields: ptEvs
func (_m *PointerEngine) Init(ptEvs chan<- swipes.PointerEvent) {
	_m.Called(ptEvs)
}

// Pause provides a mock function with given fields:
func (_m *PointerEngine) Pause() {
	_m.Called()
}

// Resume provides a mock function with given fields:
func (_m *PointerEngine) Resume() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *PointerEngine) Stop() {
	_m.Called()
}
