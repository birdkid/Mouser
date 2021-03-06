// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	log "github.com/echocrow/Mouser/pkg/log"
	mock "github.com/stretchr/testify/mock"

	monitor "github.com/echocrow/Mouser/pkg/hotkeys/monitor"
)

// Engine is an autogenerated mock type for the Engine type
type Engine struct {
	mock.Mock
}

// Deinit provides a mock function with given fields:
func (_m *Engine) Deinit() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Init provides a mock function with given fields:
func (_m *Engine) Init() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetLogCb provides a mock function with given fields: logCb
func (_m *Engine) SetLogCb(logCb log.Callback) {
	_m.Called(logCb)
}

// Start provides a mock function with given fields: m
func (_m *Engine) Start(m *monitor.Monitor) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(*monitor.Monitor) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Engine) Stop() {
	_m.Called()
}
