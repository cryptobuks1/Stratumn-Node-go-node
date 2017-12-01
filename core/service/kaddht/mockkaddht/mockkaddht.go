// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stratumn/alice/core/service/kaddht (interfaces: Host)

// Package mockkaddht is a generated GoMock package.
package mockkaddht

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	go_libp2p_net "gx/ipfs/QmNa31VPzC561NWwRsJLE7nGYZYuuD2QfpK2b1q9BK54J1/go-libp2p-net"
	go_libp2p_peerstore "gx/ipfs/QmPgDWmTmuzvP7QE5zwo1TmjbJme9pmZHNujB2453jkCTr/go-libp2p-peerstore"
	go_multistream "gx/ipfs/QmTnsezaB1wWNRHeHnYrm8K4d5i9wtyj3GsqjC3Rt5b5v5/go-multistream"
	go_multiaddr "gx/ipfs/QmXY77cVe7rVRQXZZQRioukUM7aRW3BTcAgJe12MCtb3Ji/go-multiaddr"
	go_libp2p_peer "gx/ipfs/QmXYjuNuxVzXKJCfWasQk1RqkhVLDM9jtUKhqc2WPQmFSB/go-libp2p-peer"
	go_libp2p_interface_connmgr "gx/ipfs/QmYkCrTwivapqdB3JbwvwvxymseahVkcm46ThRMAA24zCr/go-libp2p-interface-connmgr"
	go_libp2p_protocol "gx/ipfs/QmZNkThpqfVXs9GNbexPrfBbXSLNYeKrE7jwFM2oqHbyqN/go-libp2p-protocol"
	reflect "reflect"
)

// MockHost is a mock of Host interface
type MockHost struct {
	ctrl     *gomock.Controller
	recorder *MockHostMockRecorder
}

// MockHostMockRecorder is the mock recorder for MockHost
type MockHostMockRecorder struct {
	mock *MockHost
}

// NewMockHost creates a new mock instance
func NewMockHost(ctrl *gomock.Controller) *MockHost {
	mock := &MockHost{ctrl: ctrl}
	mock.recorder = &MockHostMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHost) EXPECT() *MockHostMockRecorder {
	return m.recorder
}

// Addrs mocks base method
func (m *MockHost) Addrs() []go_multiaddr.Multiaddr {
	ret := m.ctrl.Call(m, "Addrs")
	ret0, _ := ret[0].([]go_multiaddr.Multiaddr)
	return ret0
}

// Addrs indicates an expected call of Addrs
func (mr *MockHostMockRecorder) Addrs() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addrs", reflect.TypeOf((*MockHost)(nil).Addrs))
}

// Close mocks base method
func (m *MockHost) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockHostMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockHost)(nil).Close))
}

// ConnManager mocks base method
func (m *MockHost) ConnManager() go_libp2p_interface_connmgr.ConnManager {
	ret := m.ctrl.Call(m, "ConnManager")
	ret0, _ := ret[0].(go_libp2p_interface_connmgr.ConnManager)
	return ret0
}

// ConnManager indicates an expected call of ConnManager
func (mr *MockHostMockRecorder) ConnManager() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnManager", reflect.TypeOf((*MockHost)(nil).ConnManager))
}

// Connect mocks base method
func (m *MockHost) Connect(arg0 context.Context, arg1 go_libp2p_peerstore.PeerInfo) error {
	ret := m.ctrl.Call(m, "Connect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockHostMockRecorder) Connect(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockHost)(nil).Connect), arg0, arg1)
}

// ID mocks base method
func (m *MockHost) ID() go_libp2p_peer.ID {
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(go_libp2p_peer.ID)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockHostMockRecorder) ID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockHost)(nil).ID))
}

// Mux mocks base method
func (m *MockHost) Mux() *go_multistream.MultistreamMuxer {
	ret := m.ctrl.Call(m, "Mux")
	ret0, _ := ret[0].(*go_multistream.MultistreamMuxer)
	return ret0
}

// Mux indicates an expected call of Mux
func (mr *MockHostMockRecorder) Mux() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mux", reflect.TypeOf((*MockHost)(nil).Mux))
}

// Network mocks base method
func (m *MockHost) Network() go_libp2p_net.Network {
	ret := m.ctrl.Call(m, "Network")
	ret0, _ := ret[0].(go_libp2p_net.Network)
	return ret0
}

// Network indicates an expected call of Network
func (mr *MockHostMockRecorder) Network() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Network", reflect.TypeOf((*MockHost)(nil).Network))
}

// NewStream mocks base method
func (m *MockHost) NewStream(arg0 context.Context, arg1 go_libp2p_peer.ID, arg2 ...go_libp2p_protocol.ID) (go_libp2p_net.Stream, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewStream", varargs...)
	ret0, _ := ret[0].(go_libp2p_net.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewStream indicates an expected call of NewStream
func (mr *MockHostMockRecorder) NewStream(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStream", reflect.TypeOf((*MockHost)(nil).NewStream), varargs...)
}

// Peerstore mocks base method
func (m *MockHost) Peerstore() go_libp2p_peerstore.Peerstore {
	ret := m.ctrl.Call(m, "Peerstore")
	ret0, _ := ret[0].(go_libp2p_peerstore.Peerstore)
	return ret0
}

// Peerstore indicates an expected call of Peerstore
func (mr *MockHostMockRecorder) Peerstore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peerstore", reflect.TypeOf((*MockHost)(nil).Peerstore))
}

// RemoveStreamHandler mocks base method
func (m *MockHost) RemoveStreamHandler(arg0 go_libp2p_protocol.ID) {
	m.ctrl.Call(m, "RemoveStreamHandler", arg0)
}

// RemoveStreamHandler indicates an expected call of RemoveStreamHandler
func (mr *MockHostMockRecorder) RemoveStreamHandler(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStreamHandler", reflect.TypeOf((*MockHost)(nil).RemoveStreamHandler), arg0)
}

// SetRouter mocks base method
func (m *MockHost) SetRouter(arg0 func(context.Context, go_libp2p_peer.ID) (go_libp2p_peerstore.PeerInfo, error)) {
	m.ctrl.Call(m, "SetRouter", arg0)
}

// SetRouter indicates an expected call of SetRouter
func (mr *MockHostMockRecorder) SetRouter(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRouter", reflect.TypeOf((*MockHost)(nil).SetRouter), arg0)
}

// SetStreamHandler mocks base method
func (m *MockHost) SetStreamHandler(arg0 go_libp2p_protocol.ID, arg1 go_libp2p_net.StreamHandler) {
	m.ctrl.Call(m, "SetStreamHandler", arg0, arg1)
}

// SetStreamHandler indicates an expected call of SetStreamHandler
func (mr *MockHostMockRecorder) SetStreamHandler(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStreamHandler", reflect.TypeOf((*MockHost)(nil).SetStreamHandler), arg0, arg1)
}

// SetStreamHandlerMatch mocks base method
func (m *MockHost) SetStreamHandlerMatch(arg0 go_libp2p_protocol.ID, arg1 func(string) bool, arg2 go_libp2p_net.StreamHandler) {
	m.ctrl.Call(m, "SetStreamHandlerMatch", arg0, arg1, arg2)
}

// SetStreamHandlerMatch indicates an expected call of SetStreamHandlerMatch
func (mr *MockHostMockRecorder) SetStreamHandlerMatch(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStreamHandlerMatch", reflect.TypeOf((*MockHost)(nil).SetStreamHandlerMatch), arg0, arg1, arg2)
}
