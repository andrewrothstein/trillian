// Automatically generated by MockGen. DO NOT EDIT!
// Source: crypto (interfaces: Signer)

package crypto

import (
	crypto "crypto"
	gomock "github.com/golang/mock/gomock"
	io "io"
)

// Mock of Signer interface
type MockSigner struct {
	ctrl     *gomock.Controller
	recorder *_MockSignerRecorder
}

// Recorder for MockSigner (not exported)
type _MockSignerRecorder struct {
	mock *MockSigner
}

func NewMockSigner(ctrl *gomock.Controller) *MockSigner {
	mock := &MockSigner{ctrl: ctrl}
	mock.recorder = &_MockSignerRecorder{mock}
	return mock
}

func (_m *MockSigner) EXPECT() *_MockSignerRecorder {
	return _m.recorder
}

func (_m *MockSigner) Public() crypto.PublicKey {
	ret := _m.ctrl.Call(_m, "Public")
	ret0, _ := ret[0].(crypto.PublicKey)
	return ret0
}

func (_mr *_MockSignerRecorder) Public() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Public")
}

func (_m *MockSigner) Sign(_param0 io.Reader, _param1 []byte, _param2 crypto.SignerOpts) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "Sign", _param0, _param1, _param2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSignerRecorder) Sign(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Sign", arg0, arg1, arg2)
}