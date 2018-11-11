package rpc

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/33cn/chain33/client"
	"github.com/33cn/chain33/client/mocks"
	rpctypes "github.com/33cn/chain33/rpc/types"
	pty "github.com/33cn/plugin/plugin/dapp/privacy/types"
)

func newGrpc(api client.QueueProtocolAPI) *channelClient {
	return &channelClient{
		ChannelClient: rpctypes.ChannelClient{QueueProtocolAPI: api},
	}
}

func newJrpc(api client.QueueProtocolAPI) *Jrpc {
	return &Jrpc{cli: newGrpc(api)}
}

func TestChain33_PrivacyTxList(t *testing.T) {
	api := new(mocks.QueueProtocolAPI)
	testChain33 := newJrpc(api)
	actual := &pty.ReqPrivacyTransactionList{}
	api.On("ExecWalletFunc", "privacy", "PrivacyTransactionList", actual).Return(nil, errors.New("error value"))
	var testResult interface{}
	err := testChain33.PrivacyTxList(actual, &testResult)
	t.Log(err)
	assert.Equal(t, nil, testResult)
	assert.NotNil(t, err)

	mock.AssertExpectationsForObjects(t, api)
}