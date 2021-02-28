// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	swipes "github.com/birdkid/mouser/pkg/hotkeys/gestures/swipes"
	mock "github.com/stretchr/testify/mock"
)

// Monitor is an autogenerated mock type for the Monitor type
type Monitor struct {
	mock.Mock
}

// Init provides a mock function with given fields:
func (_m *Monitor) Init() <-chan swipes.Event {
	ret := _m.Called()

	var r0 <-chan swipes.Event
	if rf, ok := ret.Get(0).(func() <-chan swipes.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan swipes.Event)
		}
	}

	return r0
}

// Pause provides a mock function with given fields:
func (_m *Monitor) Pause() {
	_m.Called()
}

// Restart provides a mock function with given fields:
func (_m *Monitor) Restart() {
	_m.Called()
}

// Stop provides a mock function with given fields:
func (_m *Monitor) Stop() {
	_m.Called()
}
