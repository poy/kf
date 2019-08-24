// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/kf/pkg/client/clientset/versioned/typed/kf/v1alpha1 (interfaces: KfV1alpha1Interface,RouteInterface)

// Package route is a generated GoMock package.
package route

import (
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/google/kf/pkg/apis/kf/v1alpha1"
	v1alpha10 "github.com/google/kf/pkg/client/clientset/versioned/typed/kf/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// FakeKfAlpha1Interface is a mock of KfV1alpha1Interface interface
type FakeKfAlpha1Interface struct {
	ctrl     *gomock.Controller
	recorder *FakeKfAlpha1InterfaceMockRecorder
}

// FakeKfAlpha1InterfaceMockRecorder is the mock recorder for FakeKfAlpha1Interface
type FakeKfAlpha1InterfaceMockRecorder struct {
	mock *FakeKfAlpha1Interface
}

// NewFakeKfAlpha1Interface creates a new mock instance
func NewFakeKfAlpha1Interface(ctrl *gomock.Controller) *FakeKfAlpha1Interface {
	mock := &FakeKfAlpha1Interface{ctrl: ctrl}
	mock.recorder = &FakeKfAlpha1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeKfAlpha1Interface) EXPECT() *FakeKfAlpha1InterfaceMockRecorder {
	return m.recorder
}

// Apps mocks base method
func (m *FakeKfAlpha1Interface) Apps(arg0 string) v1alpha10.AppInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apps", arg0)
	ret0, _ := ret[0].(v1alpha10.AppInterface)
	return ret0
}

// Apps indicates an expected call of Apps
func (mr *FakeKfAlpha1InterfaceMockRecorder) Apps(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apps", reflect.TypeOf((*FakeKfAlpha1Interface)(nil).Apps), arg0)
}

// RESTClient mocks base method
func (m *FakeKfAlpha1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *FakeKfAlpha1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*FakeKfAlpha1Interface)(nil).RESTClient))
}

// RouteClaims mocks base method
func (m *FakeKfAlpha1Interface) RouteClaims(arg0 string) v1alpha10.RouteClaimInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RouteClaims", arg0)
	ret0, _ := ret[0].(v1alpha10.RouteClaimInterface)
	return ret0
}

// RouteClaims indicates an expected call of RouteClaims
func (mr *FakeKfAlpha1InterfaceMockRecorder) RouteClaims(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteClaims", reflect.TypeOf((*FakeKfAlpha1Interface)(nil).RouteClaims), arg0)
}

// Routes mocks base method
func (m *FakeKfAlpha1Interface) Routes(arg0 string) v1alpha10.RouteInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Routes", arg0)
	ret0, _ := ret[0].(v1alpha10.RouteInterface)
	return ret0
}

// Routes indicates an expected call of Routes
func (mr *FakeKfAlpha1InterfaceMockRecorder) Routes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Routes", reflect.TypeOf((*FakeKfAlpha1Interface)(nil).Routes), arg0)
}

// Sources mocks base method
func (m *FakeKfAlpha1Interface) Sources(arg0 string) v1alpha10.SourceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sources", arg0)
	ret0, _ := ret[0].(v1alpha10.SourceInterface)
	return ret0
}

// Sources indicates an expected call of Sources
func (mr *FakeKfAlpha1InterfaceMockRecorder) Sources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sources", reflect.TypeOf((*FakeKfAlpha1Interface)(nil).Sources), arg0)
}

// Spaces mocks base method
func (m *FakeKfAlpha1Interface) Spaces() v1alpha10.SpaceInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Spaces")
	ret0, _ := ret[0].(v1alpha10.SpaceInterface)
	return ret0
}

// Spaces indicates an expected call of Spaces
func (mr *FakeKfAlpha1InterfaceMockRecorder) Spaces() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Spaces", reflect.TypeOf((*FakeKfAlpha1Interface)(nil).Spaces))
}

// FakeRouteInterface is a mock of RouteInterface interface
type FakeRouteInterface struct {
	ctrl     *gomock.Controller
	recorder *FakeRouteInterfaceMockRecorder
}

// FakeRouteInterfaceMockRecorder is the mock recorder for FakeRouteInterface
type FakeRouteInterfaceMockRecorder struct {
	mock *FakeRouteInterface
}

// NewFakeRouteInterface creates a new mock instance
func NewFakeRouteInterface(ctrl *gomock.Controller) *FakeRouteInterface {
	mock := &FakeRouteInterface{ctrl: ctrl}
	mock.recorder = &FakeRouteInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeRouteInterface) EXPECT() *FakeRouteInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *FakeRouteInterface) Create(arg0 *v1alpha1.Route) (*v1alpha1.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1alpha1.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *FakeRouteInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*FakeRouteInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *FakeRouteInterface) Delete(arg0 string, arg1 *v1.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *FakeRouteInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*FakeRouteInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *FakeRouteInterface) DeleteCollection(arg0 *v1.DeleteOptions, arg1 v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *FakeRouteInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*FakeRouteInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *FakeRouteInterface) Get(arg0 string, arg1 v1.GetOptions) (*v1alpha1.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *FakeRouteInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*FakeRouteInterface)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *FakeRouteInterface) List(arg0 v1.ListOptions) (*v1alpha1.RouteList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1alpha1.RouteList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *FakeRouteInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*FakeRouteInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *FakeRouteInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1alpha1.Route, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1alpha1.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *FakeRouteInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*FakeRouteInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *FakeRouteInterface) Update(arg0 *v1alpha1.Route) (*v1alpha1.Route, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1alpha1.Route)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *FakeRouteInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*FakeRouteInterface)(nil).Update), arg0)
}

// Watch mocks base method
func (m *FakeRouteInterface) Watch(arg0 v1.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *FakeRouteInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*FakeRouteInterface)(nil).Watch), arg0)
}
