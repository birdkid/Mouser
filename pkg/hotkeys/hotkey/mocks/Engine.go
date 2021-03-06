// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	hotkey "github.com/echocrow/Mouser/pkg/hotkeys/hotkey"
	mock "github.com/stretchr/testify/mock"
)

// Engine is an autogenerated mock type for the Engine type
type Engine struct {
	mock.Mock
}

// IDFromEvent provides a mock function with given fields: _a0
func (_m *Engine) IDFromEvent(_a0 hotkey.EngineEvent) (hotkey.ID, error) {
	ret := _m.Called(_a0)

	var r0 hotkey.ID
	if rf, ok := ret.Get(0).(func(hotkey.EngineEvent) hotkey.ID); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(hotkey.ID)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(hotkey.EngineEvent) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: id, key
func (_m *Engine) Register(id hotkey.ID, key hotkey.KeyName) error {
	ret := _m.Called(id, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(hotkey.ID, hotkey.KeyName) error); ok {
		r0 = rf(id, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unregister provides a mock function with given fields: id
func (_m *Engine) Unregister(id hotkey.ID) {
	_m.Called(id)
}
