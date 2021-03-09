package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var apiKey = "KURfs3jwYcjk25berMrr43fyDtmwPgaHKgLRCjvld60u2dTlFHOa9NzuApr-Bcuu"

func TestAddMessage(t *testing.T) {
	sdk := EnchainteClient(apiKey)

	message := []byte("message 1")
	if err := sdk.Message.Write(message); err != nil {
		assert.FailNow(t, fmt.Sprintf("failed to write message: %s", err.Error()))
	}
	time.Sleep(time.Second * 3)
}

//func TestWaitMessage(t *testing.T) {
//	sdk := EnchainteClient(apiKey)
//
//	message := []byte("message 1")
//	messages := [][]byte{message}
//	receipts, err := sdk.Message.Wait(messages)
//	if err != nil {
//		assert.FailNow(t, fmt.Sprintf("failed while wating message receipts: %s", err.Error()))
//	}
//	fmt.Println("Receipts: ", receipts)
//}

//func TestProofMessage(t *testing.T) {
//	sdk := EnchainteClient(apiKey)
//
//	message := []byte("Albert Canyelles ")
//	messages := [][]byte{message}
//	proof, err := sdk.Proof.Proof(messages)
//	if err != nil {
//		assert.FailNow(t, fmt.Sprintf("failed while getting proof: %s", err.Error()))
//	}
//	fmt.Println("Proof: ", proof)
//}


//func TestE2e(t *testing.T) {
//
//	//apiKey := "JB1lKPZIdUKXhrpBVTGwsXSEFs2JT2jmp2dlmYS0nvBBsGBQ2g4hRFYOUNmneBdN"
//	apiKey := "KURfs3jwYcjk25berMrr43fyDtmwPgaHKgLRCjvld60u2dTlFHOa9NzuApr-Bcuu"
//	sdk := EnchainteClient(apiKey)
//
//	message := []byte("message 1")
//	if err := sdk.Message.Write(message); err != nil {
//		assert.FailNow(t, fmt.Sprintf("failed to write message: %s", err.Error()))
//	}
//
//
//
//	messages := [][]byte{message}
//	receipts, err := sdk.Message.Wait(messages)
//	if err != nil {
//		assert.FailNow(t, fmt.Sprintf("failed while wating message receipts: %s", err.Error()))
//	}
//	fmt.Println("Receipts: ", receipts)
//
//
//	proof, err := sdk.Proof.Proof(messages)
//	if err != nil {
//		assert.FailNow(t, fmt.Sprintf("failed while getting proof: %s", err.Error()))
//	}
//	fmt.Println("Proof: ", proof)
//
//	done := false
//
//	for !done {
//		valid, err := sdk.Proof.Verify(messages)
//		if err != nil {
//			assert.FailNow(t, fmt.Sprintf("failed while validating proof: %s", err.Error()))
//		}
//		done = valid
//		fmt.Println("Waiting: ", valid)
//		time.Sleep(500 * time.Millisecond)
//	}
//	// assert end of e2e test
//	assert.True(t, true)
//}

