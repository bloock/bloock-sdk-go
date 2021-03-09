package anchor_test

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/anchor"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnchorServiceGetAnchor(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	anchorId := 1

	http := mocks.NewHttpClient(mockCtrl)
	sdkParams := cloud.SdkParams{
		Host:       "http://localhost:3000/v1",
		AnchorGet:  fmt.Sprintf("%s%d", "/anchors/", anchorId),
	}
	s := anchor.NewService("", http, sdkParams)

	// should successfully get the anchor data
	bytes, err := json.Marshal(anchor.GetAnchorResponse{Success: true, Data: nil})
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("unexpected marshaling error: %s", err.Error()))
	}
	http.EXPECT().Request("", "GET", fmt.Sprintf("%s%s", sdkParams.Host, sdkParams.AnchorGet), nil).Return(bytes, nil).Times(1)

	anchorData, err := s.Get(anchorId)
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("unexpected error: %s", err.Error()))
	}
	assert.IsType(t, &anchor.GetAnchorResponse{}, anchorData)

	// should return an error if the anchor does not exist
	bytes, err = json.Marshal(map[string]interface{}{"success": false, "error": map[string]interface{}{"message": "anchor does not exist"} })
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("unexpected marshaling error: %s", err.Error()))
	}
	http.EXPECT().Request("", "GET", fmt.Sprintf("%s%s", sdkParams.Host, sdkParams.AnchorGet), nil).Return(bytes, nil).Times(1)

	anchorData, err = s.Get(anchorId)
	if err == nil {
		assert.FailNow(t, "error expected, none returned")
	}
	assert.Nil(t, anchorData)
	assert.Equal(t, "anchor does not exist", err.Error())
}
