// Code generated by mockery. DO NOT EDIT.

package mgmtv1alpha1connect

import (
	context "context"

	connect "connectrpc.com/connect"

	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	mock "github.com/stretchr/testify/mock"
)

// MockConnectionServiceClient is an autogenerated mock type for the ConnectionServiceClient type
type MockConnectionServiceClient struct {
	mock.Mock
}

type MockConnectionServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConnectionServiceClient) EXPECT() *MockConnectionServiceClient_Expecter {
	return &MockConnectionServiceClient_Expecter{mock: &_m.Mock}
}

// CheckConnectionConfig provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) CheckConnectionConfig(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) (*connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckConnectionConfig")
	}

	var r0 *connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) (*connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) *connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_CheckConnectionConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckConnectionConfig'
type MockConnectionServiceClient_CheckConnectionConfig_Call struct {
	*mock.Call
}

// CheckConnectionConfig is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]
func (_e *MockConnectionServiceClient_Expecter) CheckConnectionConfig(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_CheckConnectionConfig_Call {
	return &MockConnectionServiceClient_CheckConnectionConfig_Call{Call: _e.mock.On("CheckConnectionConfig", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_CheckConnectionConfig_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest])) *MockConnectionServiceClient_CheckConnectionConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_CheckConnectionConfig_Call) Return(_a0 *connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse], _a1 error) *MockConnectionServiceClient_CheckConnectionConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_CheckConnectionConfig_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.CheckConnectionConfigRequest]) (*connect.Response[mgmtv1alpha1.CheckConnectionConfigResponse], error)) *MockConnectionServiceClient_CheckConnectionConfig_Call {
	_c.Call.Return(run)
	return _c
}

// CheckSqlQuery provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) CheckSqlQuery(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) (*connect.Response[mgmtv1alpha1.CheckSqlQueryResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CheckSqlQuery")
	}

	var r0 *connect.Response[mgmtv1alpha1.CheckSqlQueryResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) (*connect.Response[mgmtv1alpha1.CheckSqlQueryResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) *connect.Response[mgmtv1alpha1.CheckSqlQueryResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.CheckSqlQueryResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_CheckSqlQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckSqlQuery'
type MockConnectionServiceClient_CheckSqlQuery_Call struct {
	*mock.Call
}

// CheckSqlQuery is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]
func (_e *MockConnectionServiceClient_Expecter) CheckSqlQuery(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_CheckSqlQuery_Call {
	return &MockConnectionServiceClient_CheckSqlQuery_Call{Call: _e.mock.On("CheckSqlQuery", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_CheckSqlQuery_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest])) *MockConnectionServiceClient_CheckSqlQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_CheckSqlQuery_Call) Return(_a0 *connect.Response[mgmtv1alpha1.CheckSqlQueryResponse], _a1 error) *MockConnectionServiceClient_CheckSqlQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_CheckSqlQuery_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.CheckSqlQueryRequest]) (*connect.Response[mgmtv1alpha1.CheckSqlQueryResponse], error)) *MockConnectionServiceClient_CheckSqlQuery_Call {
	_c.Call.Return(run)
	return _c
}

// CreateConnection provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) CreateConnection(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) (*connect.Response[mgmtv1alpha1.CreateConnectionResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateConnection")
	}

	var r0 *connect.Response[mgmtv1alpha1.CreateConnectionResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) (*connect.Response[mgmtv1alpha1.CreateConnectionResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) *connect.Response[mgmtv1alpha1.CreateConnectionResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.CreateConnectionResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_CreateConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateConnection'
type MockConnectionServiceClient_CreateConnection_Call struct {
	*mock.Call
}

// CreateConnection is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.CreateConnectionRequest]
func (_e *MockConnectionServiceClient_Expecter) CreateConnection(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_CreateConnection_Call {
	return &MockConnectionServiceClient_CreateConnection_Call{Call: _e.mock.On("CreateConnection", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_CreateConnection_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.CreateConnectionRequest])) *MockConnectionServiceClient_CreateConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.CreateConnectionRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_CreateConnection_Call) Return(_a0 *connect.Response[mgmtv1alpha1.CreateConnectionResponse], _a1 error) *MockConnectionServiceClient_CreateConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_CreateConnection_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.CreateConnectionRequest]) (*connect.Response[mgmtv1alpha1.CreateConnectionResponse], error)) *MockConnectionServiceClient_CreateConnection_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteConnection provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) DeleteConnection(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) (*connect.Response[mgmtv1alpha1.DeleteConnectionResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for DeleteConnection")
	}

	var r0 *connect.Response[mgmtv1alpha1.DeleteConnectionResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) (*connect.Response[mgmtv1alpha1.DeleteConnectionResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) *connect.Response[mgmtv1alpha1.DeleteConnectionResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.DeleteConnectionResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_DeleteConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteConnection'
type MockConnectionServiceClient_DeleteConnection_Call struct {
	*mock.Call
}

// DeleteConnection is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]
func (_e *MockConnectionServiceClient_Expecter) DeleteConnection(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_DeleteConnection_Call {
	return &MockConnectionServiceClient_DeleteConnection_Call{Call: _e.mock.On("DeleteConnection", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_DeleteConnection_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.DeleteConnectionRequest])) *MockConnectionServiceClient_DeleteConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.DeleteConnectionRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_DeleteConnection_Call) Return(_a0 *connect.Response[mgmtv1alpha1.DeleteConnectionResponse], _a1 error) *MockConnectionServiceClient_DeleteConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_DeleteConnection_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.DeleteConnectionRequest]) (*connect.Response[mgmtv1alpha1.DeleteConnectionResponse], error)) *MockConnectionServiceClient_DeleteConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnection provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) GetConnection(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetConnection")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetConnectionResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionRequest]) *connect.Response[mgmtv1alpha1.GetConnectionResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetConnectionResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_GetConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnection'
type MockConnectionServiceClient_GetConnection_Call struct {
	*mock.Call
}

// GetConnection is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionRequest]
func (_e *MockConnectionServiceClient_Expecter) GetConnection(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_GetConnection_Call {
	return &MockConnectionServiceClient_GetConnection_Call{Call: _e.mock.On("GetConnection", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_GetConnection_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionRequest])) *MockConnectionServiceClient_GetConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_GetConnection_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetConnectionResponse], _a1 error) *MockConnectionServiceClient_GetConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_GetConnection_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionResponse], error)) *MockConnectionServiceClient_GetConnection_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnectionDataStream provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) GetConnectionDataStream(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest]) (*connect.ServerStreamForClient[mgmtv1alpha1.GetConnectionDataStreamResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.ServerStreamForClient[mgmtv1alpha1.GetConnectionDataStreamResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest]) (*connect.ServerStreamForClient[mgmtv1alpha1.GetConnectionDataStreamResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest]) *connect.ServerStreamForClient[mgmtv1alpha1.GetConnectionDataStreamResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.ServerStreamForClient[mgmtv1alpha1.GetConnectionDataStreamResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_GetConnectionDataStream_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionDataStream'
type MockConnectionServiceClient_GetConnectionDataStream_Call struct {
	*mock.Call
}

// GetConnectionDataStream is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest]
func (_e *MockConnectionServiceClient_Expecter) GetConnectionDataStream(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_GetConnectionDataStream_Call {
	return &MockConnectionServiceClient_GetConnectionDataStream_Call{Call: _e.mock.On("GetConnectionDataStream", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_GetConnectionDataStream_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest])) *MockConnectionServiceClient_GetConnectionDataStream_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_GetConnectionDataStream_Call) Return(_a0 *connect.ServerStreamForClient[mgmtv1alpha1.GetConnectionDataStreamResponse], _a1 error) *MockConnectionServiceClient_GetConnectionDataStream_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_GetConnectionDataStream_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionDataStreamRequest]) (*connect.ServerStreamForClient[mgmtv1alpha1.GetConnectionDataStreamResponse], error)) *MockConnectionServiceClient_GetConnectionDataStream_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnectionForeignConstraints provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) GetConnectionForeignConstraints(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse], error) {
	ret := _m.Called(_a0, _a1)

	var r0 *connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) *connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_GetConnectionForeignConstraints_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionForeignConstraints'
type MockConnectionServiceClient_GetConnectionForeignConstraints_Call struct {
	*mock.Call
}

// GetConnectionForeignConstraints is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]
func (_e *MockConnectionServiceClient_Expecter) GetConnectionForeignConstraints(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_GetConnectionForeignConstraints_Call {
	return &MockConnectionServiceClient_GetConnectionForeignConstraints_Call{Call: _e.mock.On("GetConnectionForeignConstraints", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_GetConnectionForeignConstraints_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest])) *MockConnectionServiceClient_GetConnectionForeignConstraints_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_GetConnectionForeignConstraints_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse], _a1 error) *MockConnectionServiceClient_GetConnectionForeignConstraints_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_GetConnectionForeignConstraints_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionForeignConstraintsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionForeignConstraintsResponse], error)) *MockConnectionServiceClient_GetConnectionForeignConstraints_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnectionSchema provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) GetConnectionSchema(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionSchema")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) *connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_GetConnectionSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionSchema'
type MockConnectionServiceClient_GetConnectionSchema_Call struct {
	*mock.Call
}

// GetConnectionSchema is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]
func (_e *MockConnectionServiceClient_Expecter) GetConnectionSchema(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_GetConnectionSchema_Call {
	return &MockConnectionServiceClient_GetConnectionSchema_Call{Call: _e.mock.On("GetConnectionSchema", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_GetConnectionSchema_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest])) *MockConnectionServiceClient_GetConnectionSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_GetConnectionSchema_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse], _a1 error) *MockConnectionServiceClient_GetConnectionSchema_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_GetConnectionSchema_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionSchemaRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionSchemaResponse], error)) *MockConnectionServiceClient_GetConnectionSchema_Call {
	_c.Call.Return(run)
	return _c
}

// GetConnections provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) GetConnections(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionsResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetConnections")
	}

	var r0 *connect.Response[mgmtv1alpha1.GetConnectionsResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionsResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) *connect.Response[mgmtv1alpha1.GetConnectionsResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.GetConnectionsResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_GetConnections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnections'
type MockConnectionServiceClient_GetConnections_Call struct {
	*mock.Call
}

// GetConnections is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.GetConnectionsRequest]
func (_e *MockConnectionServiceClient_Expecter) GetConnections(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_GetConnections_Call {
	return &MockConnectionServiceClient_GetConnections_Call{Call: _e.mock.On("GetConnections", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_GetConnections_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.GetConnectionsRequest])) *MockConnectionServiceClient_GetConnections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.GetConnectionsRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_GetConnections_Call) Return(_a0 *connect.Response[mgmtv1alpha1.GetConnectionsResponse], _a1 error) *MockConnectionServiceClient_GetConnections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_GetConnections_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.GetConnectionsRequest]) (*connect.Response[mgmtv1alpha1.GetConnectionsResponse], error)) *MockConnectionServiceClient_GetConnections_Call {
	_c.Call.Return(run)
	return _c
}

// IsConnectionNameAvailable provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) IsConnectionNameAvailable(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for IsConnectionNameAvailable")
	}

	var r0 *connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) *connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_IsConnectionNameAvailable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsConnectionNameAvailable'
type MockConnectionServiceClient_IsConnectionNameAvailable_Call struct {
	*mock.Call
}

// IsConnectionNameAvailable is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]
func (_e *MockConnectionServiceClient_Expecter) IsConnectionNameAvailable(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_IsConnectionNameAvailable_Call {
	return &MockConnectionServiceClient_IsConnectionNameAvailable_Call{Call: _e.mock.On("IsConnectionNameAvailable", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_IsConnectionNameAvailable_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest])) *MockConnectionServiceClient_IsConnectionNameAvailable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_IsConnectionNameAvailable_Call) Return(_a0 *connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse], _a1 error) *MockConnectionServiceClient_IsConnectionNameAvailable_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_IsConnectionNameAvailable_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.IsConnectionNameAvailableRequest]) (*connect.Response[mgmtv1alpha1.IsConnectionNameAvailableResponse], error)) *MockConnectionServiceClient_IsConnectionNameAvailable_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateConnection provides a mock function with given fields: _a0, _a1
func (_m *MockConnectionServiceClient) UpdateConnection(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) (*connect.Response[mgmtv1alpha1.UpdateConnectionResponse], error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConnection")
	}

	var r0 *connect.Response[mgmtv1alpha1.UpdateConnectionResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) (*connect.Response[mgmtv1alpha1.UpdateConnectionResponse], error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) *connect.Response[mgmtv1alpha1.UpdateConnectionResponse]); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*connect.Response[mgmtv1alpha1.UpdateConnectionResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConnectionServiceClient_UpdateConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateConnection'
type MockConnectionServiceClient_UpdateConnection_Call struct {
	*mock.Call
}

// UpdateConnection is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]
func (_e *MockConnectionServiceClient_Expecter) UpdateConnection(_a0 interface{}, _a1 interface{}) *MockConnectionServiceClient_UpdateConnection_Call {
	return &MockConnectionServiceClient_UpdateConnection_Call{Call: _e.mock.On("UpdateConnection", _a0, _a1)}
}

func (_c *MockConnectionServiceClient_UpdateConnection_Call) Run(run func(_a0 context.Context, _a1 *connect.Request[mgmtv1alpha1.UpdateConnectionRequest])) *MockConnectionServiceClient_UpdateConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*connect.Request[mgmtv1alpha1.UpdateConnectionRequest]))
	})
	return _c
}

func (_c *MockConnectionServiceClient_UpdateConnection_Call) Return(_a0 *connect.Response[mgmtv1alpha1.UpdateConnectionResponse], _a1 error) *MockConnectionServiceClient_UpdateConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConnectionServiceClient_UpdateConnection_Call) RunAndReturn(run func(context.Context, *connect.Request[mgmtv1alpha1.UpdateConnectionRequest]) (*connect.Response[mgmtv1alpha1.UpdateConnectionResponse], error)) *MockConnectionServiceClient_UpdateConnection_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConnectionServiceClient creates a new instance of MockConnectionServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConnectionServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConnectionServiceClient {
	mock := &MockConnectionServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
