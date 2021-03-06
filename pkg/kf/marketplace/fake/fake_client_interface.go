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
// Source: github.com/google/kf/pkg/kf/marketplace/fake (interfaces: ClientInterface)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	marketplace "github.com/google/kf/pkg/kf/marketplace"
	v1beta1 "github.com/poy/service-catalog/pkg/apis/servicecatalog/v1beta1"
	reflect "reflect"
)

// FakeClientInterface is a mock of ClientInterface interface
type FakeClientInterface struct {
	ctrl     *gomock.Controller
	recorder *FakeClientInterfaceMockRecorder
}

// FakeClientInterfaceMockRecorder is the mock recorder for FakeClientInterface
type FakeClientInterfaceMockRecorder struct {
	mock *FakeClientInterface
}

// NewFakeClientInterface creates a new mock instance
func NewFakeClientInterface(ctrl *gomock.Controller) *FakeClientInterface {
	mock := &FakeClientInterface{ctrl: ctrl}
	mock.recorder = &FakeClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *FakeClientInterface) EXPECT() *FakeClientInterfaceMockRecorder {
	return m.recorder
}

// BrokerName mocks base method
func (m *FakeClientInterface) BrokerName(arg0 v1beta1.ServiceInstance) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BrokerName", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BrokerName indicates an expected call of BrokerName
func (mr *FakeClientInterfaceMockRecorder) BrokerName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BrokerName", reflect.TypeOf((*FakeClientInterface)(nil).BrokerName), arg0)
}

// ListClusterPlans mocks base method
func (m *FakeClientInterface) ListClusterPlans(arg0 marketplace.ListPlanOptions) ([]v1beta1.ClusterServicePlan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusterPlans", arg0)
	ret0, _ := ret[0].([]v1beta1.ClusterServicePlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusterPlans indicates an expected call of ListClusterPlans
func (mr *FakeClientInterfaceMockRecorder) ListClusterPlans(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusterPlans", reflect.TypeOf((*FakeClientInterface)(nil).ListClusterPlans), arg0)
}

// ListNamespacedPlans mocks base method
func (m *FakeClientInterface) ListNamespacedPlans(arg0 string, arg1 marketplace.ListPlanOptions) ([]v1beta1.ServicePlan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespacedPlans", arg0, arg1)
	ret0, _ := ret[0].([]v1beta1.ServicePlan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespacedPlans indicates an expected call of ListNamespacedPlans
func (mr *FakeClientInterfaceMockRecorder) ListNamespacedPlans(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespacedPlans", reflect.TypeOf((*FakeClientInterface)(nil).ListNamespacedPlans), arg0, arg1)
}

// Marketplace mocks base method
func (m *FakeClientInterface) Marketplace(arg0 string) (*marketplace.KfMarketplace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marketplace", arg0)
	ret0, _ := ret[0].(*marketplace.KfMarketplace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marketplace indicates an expected call of Marketplace
func (mr *FakeClientInterfaceMockRecorder) Marketplace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marketplace", reflect.TypeOf((*FakeClientInterface)(nil).Marketplace), arg0)
}
