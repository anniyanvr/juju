// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/application/deployer (interfaces: DeployerAPI,DeployStepAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	charm "github.com/juju/charm/v8"
	api "github.com/juju/juju/api"
	application "github.com/juju/juju/api/application"
	base "github.com/juju/juju/api/base"
	charm0 "github.com/juju/juju/api/common/charm"
	charms "github.com/juju/juju/api/common/charms"
	params "github.com/juju/juju/apiserver/params"
	constraints "github.com/juju/juju/core/constraints"
	crossmodel "github.com/juju/juju/core/crossmodel"
	names "github.com/juju/names/v4"
	httprequest "gopkg.in/httprequest.v1"
	macaroon "gopkg.in/macaroon.v2"
	http "net/http"
	url "net/url"
	reflect "reflect"
)

// MockDeployerAPI is a mock of DeployerAPI interface
type MockDeployerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDeployerAPIMockRecorder
}

// MockDeployerAPIMockRecorder is the mock recorder for MockDeployerAPI
type MockDeployerAPIMockRecorder struct {
	mock *MockDeployerAPI
}

// NewMockDeployerAPI creates a new mock instance
func NewMockDeployerAPI(ctrl *gomock.Controller) *MockDeployerAPI {
	mock := &MockDeployerAPI{ctrl: ctrl}
	mock.recorder = &MockDeployerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeployerAPI) EXPECT() *MockDeployerAPIMockRecorder {
	return m.recorder
}

// APICall mocks base method
func (m *MockDeployerAPI) APICall(arg0 string, arg1 int, arg2, arg3 string, arg4, arg5 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APICall", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// APICall indicates an expected call of APICall
func (mr *MockDeployerAPIMockRecorder) APICall(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APICall", reflect.TypeOf((*MockDeployerAPI)(nil).APICall), arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddCharm mocks base method
func (m *MockDeployerAPI) AddCharm(arg0 *charm.URL, arg1 charm0.Origin, arg2 bool) (charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(charm0.Origin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCharm indicates an expected call of AddCharm
func (mr *MockDeployerAPIMockRecorder) AddCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharm", reflect.TypeOf((*MockDeployerAPI)(nil).AddCharm), arg0, arg1, arg2)
}

// AddCharmWithAuthorization mocks base method
func (m *MockDeployerAPI) AddCharmWithAuthorization(arg0 *charm.URL, arg1 charm0.Origin, arg2 *macaroon.Macaroon, arg3 bool) (charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCharmWithAuthorization", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(charm0.Origin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddCharmWithAuthorization indicates an expected call of AddCharmWithAuthorization
func (mr *MockDeployerAPIMockRecorder) AddCharmWithAuthorization(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCharmWithAuthorization", reflect.TypeOf((*MockDeployerAPI)(nil).AddCharmWithAuthorization), arg0, arg1, arg2, arg3)
}

// AddLocalCharm mocks base method
func (m *MockDeployerAPI) AddLocalCharm(arg0 *charm.URL, arg1 charm.Charm, arg2 bool) (*charm.URL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLocalCharm", arg0, arg1, arg2)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLocalCharm indicates an expected call of AddLocalCharm
func (mr *MockDeployerAPIMockRecorder) AddLocalCharm(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLocalCharm", reflect.TypeOf((*MockDeployerAPI)(nil).AddLocalCharm), arg0, arg1, arg2)
}

// AddMachines mocks base method
func (m *MockDeployerAPI) AddMachines(arg0 []params.AddMachineParams) ([]params.AddMachinesResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMachines", arg0)
	ret0, _ := ret[0].([]params.AddMachinesResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMachines indicates an expected call of AddMachines
func (mr *MockDeployerAPIMockRecorder) AddMachines(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMachines", reflect.TypeOf((*MockDeployerAPI)(nil).AddMachines), arg0)
}

// AddRelation mocks base method
func (m *MockDeployerAPI) AddRelation(arg0, arg1 []string) (*params.AddRelationResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRelation", arg0, arg1)
	ret0, _ := ret[0].(*params.AddRelationResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRelation indicates an expected call of AddRelation
func (mr *MockDeployerAPIMockRecorder) AddRelation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRelation", reflect.TypeOf((*MockDeployerAPI)(nil).AddRelation), arg0, arg1)
}

// AddUnits mocks base method
func (m *MockDeployerAPI) AddUnits(arg0 application.AddUnitsParams) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUnits", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUnits indicates an expected call of AddUnits
func (mr *MockDeployerAPIMockRecorder) AddUnits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUnits", reflect.TypeOf((*MockDeployerAPI)(nil).AddUnits), arg0)
}

// BakeryClient mocks base method
func (m *MockDeployerAPI) BakeryClient() base.MacaroonDischarger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BakeryClient")
	ret0, _ := ret[0].(base.MacaroonDischarger)
	return ret0
}

// BakeryClient indicates an expected call of BakeryClient
func (mr *MockDeployerAPIMockRecorder) BakeryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BakeryClient", reflect.TypeOf((*MockDeployerAPI)(nil).BakeryClient))
}

// BestFacadeVersion mocks base method
func (m *MockDeployerAPI) BestFacadeVersion(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestFacadeVersion", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// BestFacadeVersion indicates an expected call of BestFacadeVersion
func (mr *MockDeployerAPIMockRecorder) BestFacadeVersion(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestFacadeVersion", reflect.TypeOf((*MockDeployerAPI)(nil).BestFacadeVersion), arg0)
}

// CharmInfo mocks base method
func (m *MockDeployerAPI) CharmInfo(arg0 string) (*charms.CharmInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CharmInfo", arg0)
	ret0, _ := ret[0].(*charms.CharmInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CharmInfo indicates an expected call of CharmInfo
func (mr *MockDeployerAPIMockRecorder) CharmInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmInfo", reflect.TypeOf((*MockDeployerAPI)(nil).CharmInfo), arg0)
}

// CheckCharmPlacement mocks base method
func (m *MockDeployerAPI) CheckCharmPlacement(arg0 string, arg1 *charm.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCharmPlacement", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCharmPlacement indicates an expected call of CheckCharmPlacement
func (mr *MockDeployerAPIMockRecorder) CheckCharmPlacement(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCharmPlacement", reflect.TypeOf((*MockDeployerAPI)(nil).CheckCharmPlacement), arg0, arg1)
}

// Close mocks base method
func (m *MockDeployerAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockDeployerAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDeployerAPI)(nil).Close))
}

// ConnectControllerStream mocks base method
func (m *MockDeployerAPI) ConnectControllerStream(arg0 string, arg1 url.Values, arg2 http.Header) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectControllerStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectControllerStream indicates an expected call of ConnectControllerStream
func (mr *MockDeployerAPIMockRecorder) ConnectControllerStream(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectControllerStream", reflect.TypeOf((*MockDeployerAPI)(nil).ConnectControllerStream), arg0, arg1, arg2)
}

// ConnectStream mocks base method
func (m *MockDeployerAPI) ConnectStream(arg0 string, arg1 url.Values) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectStream", arg0, arg1)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectStream indicates an expected call of ConnectStream
func (mr *MockDeployerAPIMockRecorder) ConnectStream(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectStream", reflect.TypeOf((*MockDeployerAPI)(nil).ConnectStream), arg0, arg1)
}

// Consume mocks base method
func (m *MockDeployerAPI) Consume(arg0 crossmodel.ConsumeApplicationArgs) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume
func (mr *MockDeployerAPIMockRecorder) Consume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockDeployerAPI)(nil).Consume), arg0)
}

// Context mocks base method
func (m *MockDeployerAPI) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockDeployerAPIMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockDeployerAPI)(nil).Context))
}

// Deploy mocks base method
func (m *MockDeployerAPI) Deploy(arg0 application.DeployArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy
func (mr *MockDeployerAPIMockRecorder) Deploy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockDeployerAPI)(nil).Deploy), arg0)
}

// Expose mocks base method
func (m *MockDeployerAPI) Expose(arg0 string, arg1 map[string]params.ExposedEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expose", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Expose indicates an expected call of Expose
func (mr *MockDeployerAPIMockRecorder) Expose(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expose", reflect.TypeOf((*MockDeployerAPI)(nil).Expose), arg0, arg1)
}

// GetAnnotations mocks base method
func (m *MockDeployerAPI) GetAnnotations(arg0 []string) ([]params.AnnotationsGetResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAnnotations", arg0)
	ret0, _ := ret[0].([]params.AnnotationsGetResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnnotations indicates an expected call of GetAnnotations
func (mr *MockDeployerAPIMockRecorder) GetAnnotations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnnotations", reflect.TypeOf((*MockDeployerAPI)(nil).GetAnnotations), arg0)
}

// GetCharmURLOrigin mocks base method
func (m *MockDeployerAPI) GetCharmURLOrigin(arg0, arg1 string) (*charm.URL, charm0.Origin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCharmURLOrigin", arg0, arg1)
	ret0, _ := ret[0].(*charm.URL)
	ret1, _ := ret[1].(charm0.Origin)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCharmURLOrigin indicates an expected call of GetCharmURLOrigin
func (mr *MockDeployerAPIMockRecorder) GetCharmURLOrigin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCharmURLOrigin", reflect.TypeOf((*MockDeployerAPI)(nil).GetCharmURLOrigin), arg0, arg1)
}

// GetConfig mocks base method
func (m *MockDeployerAPI) GetConfig(arg0 string, arg1 ...string) ([]map[string]interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfig", varargs...)
	ret0, _ := ret[0].([]map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockDeployerAPIMockRecorder) GetConfig(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockDeployerAPI)(nil).GetConfig), varargs...)
}

// GetConstraints mocks base method
func (m *MockDeployerAPI) GetConstraints(arg0 ...string) ([]constraints.Value, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConstraints", varargs...)
	ret0, _ := ret[0].([]constraints.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConstraints indicates an expected call of GetConstraints
func (mr *MockDeployerAPIMockRecorder) GetConstraints(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConstraints", reflect.TypeOf((*MockDeployerAPI)(nil).GetConstraints), arg0...)
}

// GrantOffer mocks base method
func (m *MockDeployerAPI) GrantOffer(arg0, arg1 string, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GrantOffer", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantOffer indicates an expected call of GrantOffer
func (mr *MockDeployerAPIMockRecorder) GrantOffer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantOffer", reflect.TypeOf((*MockDeployerAPI)(nil).GrantOffer), varargs...)
}

// HTTPClient mocks base method
func (m *MockDeployerAPI) HTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient
func (mr *MockDeployerAPIMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockDeployerAPI)(nil).HTTPClient))
}

// IsMetered mocks base method
func (m *MockDeployerAPI) IsMetered(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMetered", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMetered indicates an expected call of IsMetered
func (mr *MockDeployerAPIMockRecorder) IsMetered(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMetered", reflect.TypeOf((*MockDeployerAPI)(nil).IsMetered), arg0)
}

// ModelGet mocks base method
func (m *MockDeployerAPI) ModelGet() (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelGet")
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelGet indicates an expected call of ModelGet
func (mr *MockDeployerAPIMockRecorder) ModelGet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelGet", reflect.TypeOf((*MockDeployerAPI)(nil).ModelGet))
}

// ModelTag mocks base method
func (m *MockDeployerAPI) ModelTag() (names.ModelTag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModelTag indicates an expected call of ModelTag
func (mr *MockDeployerAPIMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockDeployerAPI)(nil).ModelTag))
}

// ModelUUID mocks base method
func (m *MockDeployerAPI) ModelUUID() (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelUUID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModelUUID indicates an expected call of ModelUUID
func (mr *MockDeployerAPIMockRecorder) ModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelUUID", reflect.TypeOf((*MockDeployerAPI)(nil).ModelUUID))
}

// Offer mocks base method
func (m *MockDeployerAPI) Offer(arg0, arg1 string, arg2 []string, arg3, arg4 string) ([]params.ErrorResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offer", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]params.ErrorResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Offer indicates an expected call of Offer
func (mr *MockDeployerAPIMockRecorder) Offer(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offer", reflect.TypeOf((*MockDeployerAPI)(nil).Offer), arg0, arg1, arg2, arg3, arg4)
}

// ScaleApplication mocks base method
func (m *MockDeployerAPI) ScaleApplication(arg0 application.ScaleApplicationParams) (params.ScaleApplicationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleApplication", arg0)
	ret0, _ := ret[0].(params.ScaleApplicationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScaleApplication indicates an expected call of ScaleApplication
func (mr *MockDeployerAPIMockRecorder) ScaleApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleApplication", reflect.TypeOf((*MockDeployerAPI)(nil).ScaleApplication), arg0)
}

// Sequences mocks base method
func (m *MockDeployerAPI) Sequences() (map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sequences")
	ret0, _ := ret[0].(map[string]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sequences indicates an expected call of Sequences
func (mr *MockDeployerAPIMockRecorder) Sequences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sequences", reflect.TypeOf((*MockDeployerAPI)(nil).Sequences))
}

// SetAnnotation mocks base method
func (m *MockDeployerAPI) SetAnnotation(arg0 map[string]map[string]string) ([]params.ErrorResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetAnnotation", arg0)
	ret0, _ := ret[0].([]params.ErrorResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetAnnotation indicates an expected call of SetAnnotation
func (mr *MockDeployerAPIMockRecorder) SetAnnotation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotation", reflect.TypeOf((*MockDeployerAPI)(nil).SetAnnotation), arg0)
}

// SetCharm mocks base method
func (m *MockDeployerAPI) SetCharm(arg0 string, arg1 application.SetCharmConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetCharm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCharm indicates an expected call of SetCharm
func (mr *MockDeployerAPIMockRecorder) SetCharm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCharm", reflect.TypeOf((*MockDeployerAPI)(nil).SetCharm), arg0, arg1)
}

// SetConfig mocks base method
func (m *MockDeployerAPI) SetConfig(arg0, arg1, arg2 string, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConfig", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConfig indicates an expected call of SetConfig
func (mr *MockDeployerAPIMockRecorder) SetConfig(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConfig", reflect.TypeOf((*MockDeployerAPI)(nil).SetConfig), arg0, arg1, arg2, arg3)
}

// SetConstraints mocks base method
func (m *MockDeployerAPI) SetConstraints(arg0 string, arg1 constraints.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetConstraints", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetConstraints indicates an expected call of SetConstraints
func (mr *MockDeployerAPIMockRecorder) SetConstraints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConstraints", reflect.TypeOf((*MockDeployerAPI)(nil).SetConstraints), arg0, arg1)
}

// SetMetricCredentials mocks base method
func (m *MockDeployerAPI) SetMetricCredentials(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMetricCredentials", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMetricCredentials indicates an expected call of SetMetricCredentials
func (mr *MockDeployerAPIMockRecorder) SetMetricCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetricCredentials", reflect.TypeOf((*MockDeployerAPI)(nil).SetMetricCredentials), arg0, arg1)
}

// Status mocks base method
func (m *MockDeployerAPI) Status(arg0 []string) (*params.FullStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status", arg0)
	ret0, _ := ret[0].(*params.FullStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockDeployerAPIMockRecorder) Status(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockDeployerAPI)(nil).Status), arg0)
}

// Update mocks base method
func (m *MockDeployerAPI) Update(arg0 params.ApplicationUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockDeployerAPIMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeployerAPI)(nil).Update), arg0)
}

// WatchAll mocks base method
func (m *MockDeployerAPI) WatchAll() (api.AllWatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAll")
	ret0, _ := ret[0].(api.AllWatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAll indicates an expected call of WatchAll
func (mr *MockDeployerAPIMockRecorder) WatchAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAll", reflect.TypeOf((*MockDeployerAPI)(nil).WatchAll))
}

// MockDeployStepAPI is a mock of DeployStepAPI interface
type MockDeployStepAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDeployStepAPIMockRecorder
}

// MockDeployStepAPIMockRecorder is the mock recorder for MockDeployStepAPI
type MockDeployStepAPIMockRecorder struct {
	mock *MockDeployStepAPI
}

// NewMockDeployStepAPI creates a new mock instance
func NewMockDeployStepAPI(ctrl *gomock.Controller) *MockDeployStepAPI {
	mock := &MockDeployStepAPI{ctrl: ctrl}
	mock.recorder = &MockDeployStepAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeployStepAPI) EXPECT() *MockDeployStepAPIMockRecorder {
	return m.recorder
}

// IsMetered mocks base method
func (m *MockDeployStepAPI) IsMetered(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMetered", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMetered indicates an expected call of IsMetered
func (mr *MockDeployStepAPIMockRecorder) IsMetered(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMetered", reflect.TypeOf((*MockDeployStepAPI)(nil).IsMetered), arg0)
}

// SetMetricCredentials mocks base method
func (m *MockDeployStepAPI) SetMetricCredentials(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMetricCredentials", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMetricCredentials indicates an expected call of SetMetricCredentials
func (mr *MockDeployStepAPIMockRecorder) SetMetricCredentials(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetricCredentials", reflect.TypeOf((*MockDeployStepAPI)(nil).SetMetricCredentials), arg0, arg1)
}
