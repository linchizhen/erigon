// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon-lib/direct (interfaces: SentryClient)

// Package direct is a generated GoMock package.
package direct

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sentry "github.com/ledgerwatch/erigon-lib/gointerfaces/sentry"
	types "github.com/ledgerwatch/erigon-lib/gointerfaces/types"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockSentryClient is a mock of SentryClient interface.
type MockSentryClient struct {
	ctrl     *gomock.Controller
	recorder *MockSentryClientMockRecorder
}

// MockSentryClientMockRecorder is the mock recorder for MockSentryClient.
type MockSentryClientMockRecorder struct {
	mock *MockSentryClient
}

// NewMockSentryClient creates a new mock instance.
func NewMockSentryClient(ctrl *gomock.Controller) *MockSentryClient {
	mock := &MockSentryClient{ctrl: ctrl}
	mock.recorder = &MockSentryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSentryClient) EXPECT() *MockSentryClientMockRecorder {
	return m.recorder
}

// AddPeer mocks base method.
func (m *MockSentryClient) AddPeer(arg0 context.Context, arg1 *sentry.AddPeerRequest, arg2 ...grpc.CallOption) (*sentry.AddPeerReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddPeer", varargs...)
	ret0, _ := ret[0].(*sentry.AddPeerReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPeer indicates an expected call of AddPeer.
func (mr *MockSentryClientMockRecorder) AddPeer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeer", reflect.TypeOf((*MockSentryClient)(nil).AddPeer), varargs...)
}

// HandShake mocks base method.
func (m *MockSentryClient) HandShake(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*sentry.HandShakeReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HandShake", varargs...)
	ret0, _ := ret[0].(*sentry.HandShakeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandShake indicates an expected call of HandShake.
func (mr *MockSentryClientMockRecorder) HandShake(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandShake", reflect.TypeOf((*MockSentryClient)(nil).HandShake), varargs...)
}

// MarkDisconnected mocks base method.
func (m *MockSentryClient) MarkDisconnected() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkDisconnected")
}

// MarkDisconnected indicates an expected call of MarkDisconnected.
func (mr *MockSentryClientMockRecorder) MarkDisconnected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDisconnected", reflect.TypeOf((*MockSentryClient)(nil).MarkDisconnected))
}

// Messages mocks base method.
func (m *MockSentryClient) Messages(arg0 context.Context, arg1 *sentry.MessagesRequest, arg2 ...grpc.CallOption) (sentry.Sentry_MessagesClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Messages", varargs...)
	ret0, _ := ret[0].(sentry.Sentry_MessagesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Messages indicates an expected call of Messages.
func (mr *MockSentryClientMockRecorder) Messages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Messages", reflect.TypeOf((*MockSentryClient)(nil).Messages), varargs...)
}

// NodeInfo mocks base method.
func (m *MockSentryClient) NodeInfo(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*types.NodeInfoReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NodeInfo", varargs...)
	ret0, _ := ret[0].(*types.NodeInfoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeInfo indicates an expected call of NodeInfo.
func (mr *MockSentryClientMockRecorder) NodeInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeInfo", reflect.TypeOf((*MockSentryClient)(nil).NodeInfo), varargs...)
}

// PeerById mocks base method.
func (m *MockSentryClient) PeerById(arg0 context.Context, arg1 *sentry.PeerByIdRequest, arg2 ...grpc.CallOption) (*sentry.PeerByIdReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PeerById", varargs...)
	ret0, _ := ret[0].(*sentry.PeerByIdReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerById indicates an expected call of PeerById.
func (mr *MockSentryClientMockRecorder) PeerById(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerById", reflect.TypeOf((*MockSentryClient)(nil).PeerById), varargs...)
}

// PeerCount mocks base method.
func (m *MockSentryClient) PeerCount(arg0 context.Context, arg1 *sentry.PeerCountRequest, arg2 ...grpc.CallOption) (*sentry.PeerCountReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PeerCount", varargs...)
	ret0, _ := ret[0].(*sentry.PeerCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerCount indicates an expected call of PeerCount.
func (mr *MockSentryClientMockRecorder) PeerCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerCount", reflect.TypeOf((*MockSentryClient)(nil).PeerCount), varargs...)
}

// PeerEvents mocks base method.
func (m *MockSentryClient) PeerEvents(arg0 context.Context, arg1 *sentry.PeerEventsRequest, arg2 ...grpc.CallOption) (sentry.Sentry_PeerEventsClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PeerEvents", varargs...)
	ret0, _ := ret[0].(sentry.Sentry_PeerEventsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerEvents indicates an expected call of PeerEvents.
func (mr *MockSentryClientMockRecorder) PeerEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerEvents", reflect.TypeOf((*MockSentryClient)(nil).PeerEvents), varargs...)
}

// PeerMinBlock mocks base method.
func (m *MockSentryClient) PeerMinBlock(arg0 context.Context, arg1 *sentry.PeerMinBlockRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PeerMinBlock", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerMinBlock indicates an expected call of PeerMinBlock.
func (mr *MockSentryClientMockRecorder) PeerMinBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerMinBlock", reflect.TypeOf((*MockSentryClient)(nil).PeerMinBlock), varargs...)
}

// Peers mocks base method.
func (m *MockSentryClient) Peers(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*sentry.PeersReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Peers", varargs...)
	ret0, _ := ret[0].(*sentry.PeersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Peers indicates an expected call of Peers.
func (mr *MockSentryClientMockRecorder) Peers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockSentryClient)(nil).Peers), varargs...)
}

// PenalizePeer mocks base method.
func (m *MockSentryClient) PenalizePeer(arg0 context.Context, arg1 *sentry.PenalizePeerRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PenalizePeer", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PenalizePeer indicates an expected call of PenalizePeer.
func (mr *MockSentryClientMockRecorder) PenalizePeer(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PenalizePeer", reflect.TypeOf((*MockSentryClient)(nil).PenalizePeer), varargs...)
}

// Protocol mocks base method.
func (m *MockSentryClient) Protocol() uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(uint)
	return ret0
}

// Protocol indicates an expected call of Protocol.
func (mr *MockSentryClientMockRecorder) Protocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockSentryClient)(nil).Protocol))
}

// Ready mocks base method.
func (m *MockSentryClient) Ready() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Ready indicates an expected call of Ready.
func (mr *MockSentryClientMockRecorder) Ready() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockSentryClient)(nil).Ready))
}

// SendMessageById mocks base method.
func (m *MockSentryClient) SendMessageById(arg0 context.Context, arg1 *sentry.SendMessageByIdRequest, arg2 ...grpc.CallOption) (*sentry.SentPeers, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageById", varargs...)
	ret0, _ := ret[0].(*sentry.SentPeers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessageById indicates an expected call of SendMessageById.
func (mr *MockSentryClientMockRecorder) SendMessageById(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageById", reflect.TypeOf((*MockSentryClient)(nil).SendMessageById), varargs...)
}

// SendMessageByMinBlock mocks base method.
func (m *MockSentryClient) SendMessageByMinBlock(arg0 context.Context, arg1 *sentry.SendMessageByMinBlockRequest, arg2 ...grpc.CallOption) (*sentry.SentPeers, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageByMinBlock", varargs...)
	ret0, _ := ret[0].(*sentry.SentPeers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessageByMinBlock indicates an expected call of SendMessageByMinBlock.
func (mr *MockSentryClientMockRecorder) SendMessageByMinBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageByMinBlock", reflect.TypeOf((*MockSentryClient)(nil).SendMessageByMinBlock), varargs...)
}

// SendMessageToAll mocks base method.
func (m *MockSentryClient) SendMessageToAll(arg0 context.Context, arg1 *sentry.OutboundMessageData, arg2 ...grpc.CallOption) (*sentry.SentPeers, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageToAll", varargs...)
	ret0, _ := ret[0].(*sentry.SentPeers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessageToAll indicates an expected call of SendMessageToAll.
func (mr *MockSentryClientMockRecorder) SendMessageToAll(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageToAll", reflect.TypeOf((*MockSentryClient)(nil).SendMessageToAll), varargs...)
}

// SendMessageToRandomPeers mocks base method.
func (m *MockSentryClient) SendMessageToRandomPeers(arg0 context.Context, arg1 *sentry.SendMessageToRandomPeersRequest, arg2 ...grpc.CallOption) (*sentry.SentPeers, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageToRandomPeers", varargs...)
	ret0, _ := ret[0].(*sentry.SentPeers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessageToRandomPeers indicates an expected call of SendMessageToRandomPeers.
func (mr *MockSentryClientMockRecorder) SendMessageToRandomPeers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageToRandomPeers", reflect.TypeOf((*MockSentryClient)(nil).SendMessageToRandomPeers), varargs...)
}

// SetStatus mocks base method.
func (m *MockSentryClient) SetStatus(arg0 context.Context, arg1 *sentry.StatusData, arg2 ...grpc.CallOption) (*sentry.SetStatusReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetStatus", varargs...)
	ret0, _ := ret[0].(*sentry.SetStatusReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockSentryClientMockRecorder) SetStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockSentryClient)(nil).SetStatus), varargs...)
}