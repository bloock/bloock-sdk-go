package cloud_test

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/config"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCloudServiceConfigParameters(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	http := mocks.NewHttpClient(mockCtrl)
	constants := config.Constants{
		Cloud: config.Cloud{
			Azure: config.Azure{
				Test: false,
			},
		},
	}

	s := cloud.Azure(http, constants)

	resp := cloud.SdkParams{
		Host:                "some host",
		MessageWrite:        "some endpoint",
		MessageFetch:        "some endpoint",
		MessageProof:        "some endpoint",
		SmartContract:       "some contract",
		ContractAbi:         "some contract ABI",
		Provider:            "some provider",
		WriteInterval:       0,
		WaitIntervalFactor:  0,
		WaitIntervalDefault: 0,
		ConfigInterval:      0,
	}

	bytes, err := json.Marshal(resp)
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("marshaling error: %s", err.Error()))
	}

	http.EXPECT().Request("", "GET", gomock.Any(), gomock.Any(), nil).Return(bytes, nil)

	if err := s.ConfigParameters(); err != nil {
		assert.FailNow(t, fmt.Sprintf("unexpected error: %s", err.Error()))
	}

	assert.Equal(t, resp, s.SdkParameters())
}
