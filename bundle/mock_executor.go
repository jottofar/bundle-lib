// Code generated by mockery v1.0.0

// Generated

package bundle

import mock "github.com/stretchr/testify/mock"

// MockExecutor is an autogenerated mock type for the Executor type
type MockExecutor struct {
	mock.Mock
}

// Bind provides a mock function with given fields: instance, parameters, bindingID
func (_m *MockExecutor) Bind(instance *ServiceInstance, parameters *Parameters, bindingID string) <-chan StatusMessage {
	ret := _m.Called(instance, parameters, bindingID)

	var r0 <-chan StatusMessage
	if rf, ok := ret.Get(0).(func(*ServiceInstance, *Parameters, string) <-chan StatusMessage); ok {
		r0 = rf(instance, parameters, bindingID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan StatusMessage)
		}
	}

	return r0
}

// Deprovision provides a mock function with given fields: instance
func (_m *MockExecutor) Deprovision(instance *ServiceInstance) <-chan StatusMessage {
	ret := _m.Called(instance)

	var r0 <-chan StatusMessage
	if rf, ok := ret.Get(0).(func(*ServiceInstance) <-chan StatusMessage); ok {
		r0 = rf(instance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan StatusMessage)
		}
	}

	return r0
}

// ExtractedCredentials provides a mock function with given fields:
func (_m *MockExecutor) ExtractedCredentials() *ExtractedCredentials {
	ret := _m.Called()

	var r0 *ExtractedCredentials
	if rf, ok := ret.Get(0).(func() *ExtractedCredentials); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ExtractedCredentials)
		}
	}

	return r0
}

// LastStatus provides a mock function with given fields:
func (_m *MockExecutor) LastStatus() StatusMessage {
	ret := _m.Called()

	var r0 StatusMessage
	if rf, ok := ret.Get(0).(func() StatusMessage); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(StatusMessage)
	}

	return r0
}

// PodName provides a mock function with given fields:
func (_m *MockExecutor) PodName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Provision provides a mock function with given fields: _a0
func (_m *MockExecutor) Provision(_a0 *ServiceInstance) <-chan StatusMessage {
	ret := _m.Called(_a0)

	var r0 <-chan StatusMessage
	if rf, ok := ret.Get(0).(func(*ServiceInstance) <-chan StatusMessage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan StatusMessage)
		}
	}

	return r0
}

// Unbind provides a mock function with given fields: instance, parameters, bindingID
func (_m *MockExecutor) Unbind(instance *ServiceInstance, parameters *Parameters, bindingID string) <-chan StatusMessage {
	ret := _m.Called(instance, parameters, bindingID)

	var r0 <-chan StatusMessage
	if rf, ok := ret.Get(0).(func(*ServiceInstance, *Parameters, string) <-chan StatusMessage); ok {
		r0 = rf(instance, parameters, bindingID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan StatusMessage)
		}
	}

	return r0
}

// Update provides a mock function with given fields: instance
func (_m *MockExecutor) Update(instance *ServiceInstance) <-chan StatusMessage {
	ret := _m.Called(instance)

	var r0 <-chan StatusMessage
	if rf, ok := ret.Get(0).(func(*ServiceInstance) <-chan StatusMessage); ok {
		r0 = rf(instance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan StatusMessage)
		}
	}

	return r0
}
