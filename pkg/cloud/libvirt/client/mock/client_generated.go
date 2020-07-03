// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mock is a generated GoMock package.
package mock

import (
	"context"
	gomock "github.com/golang/mock/gomock"
	libvirt_go "github.com/libvirt/libvirt-go"
	client "github.com/openshift/cluster-api-provider-libvirt/pkg/cloud/libvirt/client"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// CreateDomain mocks base method
func (m *MockClient) CreateDomain(ctx context.Context, arg0 client.CreateDomainInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDomain", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateDomain indicates an expected call of CreateDomain
func (mr *MockClientMockRecorder) CreateDomain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDomain", reflect.TypeOf((*MockClient)(nil).CreateDomain), arg0)
}

// DeleteDomain mocks base method
func (m *MockClient) DeleteDomain(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDomain", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDomain indicates an expected call of DeleteDomain
func (mr *MockClientMockRecorder) DeleteDomain(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDomain", reflect.TypeOf((*MockClient)(nil).DeleteDomain), name)
}

// DomainExists mocks base method
func (m *MockClient) DomainExists(name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DomainExists", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DomainExists indicates an expected call of DomainExists
func (mr *MockClientMockRecorder) DomainExists(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DomainExists", reflect.TypeOf((*MockClient)(nil).DomainExists), name)
}

// LookupDomainByName mocks base method
func (m *MockClient) LookupDomainByName(name string) (*libvirt_go.Domain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupDomainByName", name)
	ret0, _ := ret[0].(*libvirt_go.Domain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupDomainByName indicates an expected call of LookupDomainByName
func (mr *MockClientMockRecorder) LookupDomainByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupDomainByName", reflect.TypeOf((*MockClient)(nil).LookupDomainByName), name)
}

// CreateVolume mocks base method
func (m *MockClient) CreateVolume(arg0 client.CreateVolumeInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockClientMockRecorder) CreateVolume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockClient)(nil).CreateVolume), arg0)
}

// VolumeExists mocks base method
func (m *MockClient) VolumeExists(name string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VolumeExists", name)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VolumeExists indicates an expected call of VolumeExists
func (mr *MockClientMockRecorder) VolumeExists(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VolumeExists", reflect.TypeOf((*MockClient)(nil).VolumeExists), name)
}

// DeleteVolume mocks base method
func (m *MockClient) DeleteVolume(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVolume indicates an expected call of DeleteVolume
func (mr *MockClientMockRecorder) DeleteVolume(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockClient)(nil).DeleteVolume), name)
}

// GetDHCPLeasesByNetwork mocks base method
func (m *MockClient) GetDHCPLeasesByNetwork(networkName string) ([]libvirt_go.NetworkDHCPLease, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDHCPLeasesByNetwork", networkName)
	ret0, _ := ret[0].([]libvirt_go.NetworkDHCPLease)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDHCPLeasesByNetwork indicates an expected call of GetDHCPLeasesByNetwork
func (mr *MockClientMockRecorder) GetDHCPLeasesByNetwork(networkName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDHCPLeasesByNetwork", reflect.TypeOf((*MockClient)(nil).GetDHCPLeasesByNetwork), networkName)
}

// LookupDomainHostnameByDHCPLease mocks base method
func (m *MockClient) LookupDomainHostnameByDHCPLease(domIPAddress, networkName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LookupDomainHostnameByDHCPLease", domIPAddress, networkName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookupDomainHostnameByDHCPLease indicates an expected call of LookupDomainHostnameByDHCPLease
func (mr *MockClientMockRecorder) LookupDomainHostnameByDHCPLease(domIPAddress, networkName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookupDomainHostnameByDHCPLease", reflect.TypeOf((*MockClient)(nil).LookupDomainHostnameByDHCPLease), domIPAddress, networkName)
}
