package credential_test

import (
	"encoding/json"
	"fmt"
	"github.com/enchainte/enchainte-sdk-go/internal/cloud"
	"github.com/enchainte/enchainte-sdk-go/internal/credential"
	"github.com/enchainte/enchainte-sdk-go/internal/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCredentialServiceDeleteApiKey(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	clientId := "ed8af1e3-f7bc-4544-8158-8d0a684cc4ed"

	http := mocks.NewHttpClient(mockCtrl)
	sdkParams := cloud.SdkParams{
		Host:             "http://localhost:3000/v1",
		CredentialDelete: fmt.Sprintf("/clients/%s/credentials", clientId),
	}
	s := credential.NewService("", http, sdkParams)

	body := credential.DeleteApiKeyRequest{ApiKey: ""}

	// should successfully get the anchor data
	bytes, err := json.Marshal(map[string]interface{}{"success": true})
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("unexpected marshaling error: %s", err.Error()))
	}
	http.EXPECT().Request("", "DELETE", fmt.Sprintf("%s%s", sdkParams.Host, sdkParams.CredentialDelete), body).Return(bytes, nil).Times(1)

	err = s.Delete("", clientId)
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("unexpected error: %s", err.Error()))
	}
	assert.Nil(t, err)

	// should return an error if the anchor does not exist
	bytes, err = json.Marshal(map[string]interface{}{"success": false, "error": map[string]interface{}{"record": "api key not found"}})
	if err != nil {
		assert.FailNow(t, fmt.Sprintf("unexpected marshaling error: %s", err.Error()))
	}
	http.EXPECT().Request("", "DELETE", fmt.Sprintf("%s%s", sdkParams.Host, sdkParams.CredentialDelete), body).Return(bytes, nil).Times(1)

	err = s.Delete("", clientId)
	if err == nil {
		assert.FailNow(t, "error expected, none returned")
	}
	assert.Equal(t, "api key not found", err.Error())
}
