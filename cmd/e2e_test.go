package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestE2e(t *testing.T) {

	apiKey := "api-key"
	sdk := EnchainteClient(apiKey)

	msg := []byte("message 1")
	msgs := [][]byte{msg}
	wResp, err := sdk.Message.Write(msgs)
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("failed to write message: %s", err.Error()))
	}
	fmt.Println(">> message sent")

	fmt.Println(">> waiting...")
	_, err = sdk.Anchor.Wait(wResp.Data.Anchor)
	require.Nil(t, err, "failed while waiting message receipts: %v", err)

	fmt.Println(">> fetching messages...")
	receipts, err := sdk.Message.Search(msgs)
	require.Nil(t, err)
	require.NotNil(t, receipts)

	fmt.Println(">> getting proof...")
	proof, err := sdk.Proof.Proof(msgs)
	require.Nil(t, err)
	require.NotNil(t, proof)

	done := false

	for !done {
		fmt.Println(">> verifying...")
		valid, err := sdk.Proof.Verify(msgs)
		require.Nil(t, err)
		require.True(t, valid)

		done = valid
		time.Sleep(500 * time.Millisecond)
	}
	// assert end of e2e test
	assert.True(t, true)
}

