// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1 (interfaces: ApiextensionsV1beta1Interface,CustomResourceDefinitionInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	v1beta10 "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// MockApiextensionsV1beta1Interface is a mock of ApiextensionsV1beta1Interface interface
type MockApiextensionsV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockApiextensionsV1beta1InterfaceMockRecorder
}

// MockApiextensionsV1beta1InterfaceMockRecorder is the mock recorder for MockApiextensionsV1beta1Interface
type MockApiextensionsV1beta1InterfaceMockRecorder struct {
	mock *MockApiextensionsV1beta1Interface
}

// NewMockApiextensionsV1beta1Interface creates a new mock instance
func NewMockApiextensionsV1beta1Interface(ctrl *gomock.Controller) *MockApiextensionsV1beta1Interface {
	mock := &MockApiextensionsV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockApiextensionsV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApiextensionsV1beta1Interface) EXPECT() *MockApiextensionsV1beta1InterfaceMockRecorder {
	return m.recorder
}

// CustomResourceDefinitions mocks base method
func (m *MockApiextensionsV1beta1Interface) CustomResourceDefinitions() v1beta10.CustomResourceDefinitionInterface {
	ret := m.ctrl.Call(m, "CustomResourceDefinitions")
	ret0, _ := ret[0].(v1beta10.CustomResourceDefinitionInterface)
	return ret0
}

// CustomResourceDefinitions indicates an expected call of CustomResourceDefinitions
func (mr *MockApiextensionsV1beta1InterfaceMockRecorder) CustomResourceDefinitions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CustomResourceDefinitions", reflect.TypeOf((*MockApiextensionsV1beta1Interface)(nil).CustomResourceDefinitions))
}

// RESTClient mocks base method
func (m *MockApiextensionsV1beta1Interface) RESTClient() rest.Interface {
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockApiextensionsV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockApiextensionsV1beta1Interface)(nil).RESTClient))
}

// MockCustomResourceDefinitionInterface is a mock of CustomResourceDefinitionInterface interface
type MockCustomResourceDefinitionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCustomResourceDefinitionInterfaceMockRecorder
}

// MockCustomResourceDefinitionInterfaceMockRecorder is the mock recorder for MockCustomResourceDefinitionInterface
type MockCustomResourceDefinitionInterfaceMockRecorder struct {
	mock *MockCustomResourceDefinitionInterface
}

// NewMockCustomResourceDefinitionInterface creates a new mock instance
func NewMockCustomResourceDefinitionInterface(ctrl *gomock.Controller) *MockCustomResourceDefinitionInterface {
	mock := &MockCustomResourceDefinitionInterface{ctrl: ctrl}
	mock.recorder = &MockCustomResourceDefinitionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCustomResourceDefinitionInterface) EXPECT() *MockCustomResourceDefinitionInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockCustomResourceDefinitionInterface) Create(arg0 *v1beta1.CustomResourceDefinition) (*v1beta1.CustomResourceDefinition, error) {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockCustomResourceDefinitionInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomResourceDefinitionInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockCustomResourceDefinitionInterface) Delete(arg0 string, arg1 *v1.DeleteOptions) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockCustomResourceDefinitionInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomResourceDefinitionInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *MockCustomResourceDefinitionInterface) DeleteCollection(arg0 *v1.DeleteOptions, arg1 v1.ListOptions) error {
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockCustomResourceDefinitionInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockCustomResourceDefinitionInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockCustomResourceDefinitionInterface) Get(arg0 string, arg1 v1.GetOptions) (*v1beta1.CustomResourceDefinition, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockCustomResourceDefinitionInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCustomResourceDefinitionInterface)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockCustomResourceDefinitionInterface) List(arg0 v1.ListOptions) (*v1beta1.CustomResourceDefinitionList, error) {
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinitionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockCustomResourceDefinitionInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCustomResourceDefinitionInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockCustomResourceDefinitionInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1beta1.CustomResourceDefinition, error) {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockCustomResourceDefinitionInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockCustomResourceDefinitionInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockCustomResourceDefinitionInterface) Update(arg0 *v1beta1.CustomResourceDefinition) (*v1beta1.CustomResourceDefinition, error) {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockCustomResourceDefinitionInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomResourceDefinitionInterface)(nil).Update), arg0)
}

// UpdateStatus mocks base method
func (m *MockCustomResourceDefinitionInterface) UpdateStatus(arg0 *v1beta1.CustomResourceDefinition) (*v1beta1.CustomResourceDefinition, error) {
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1beta1.CustomResourceDefinition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockCustomResourceDefinitionInterfaceMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockCustomResourceDefinitionInterface)(nil).UpdateStatus), arg0)
}

// Watch mocks base method
func (m *MockCustomResourceDefinitionInterface) Watch(arg0 v1.ListOptions) (watch.Interface, error) {
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockCustomResourceDefinitionInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockCustomResourceDefinitionInterface)(nil).Watch), arg0)
}
