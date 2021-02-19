package cloud_test

import (
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

	resp := "{\"items\":[{\"etag\":\"LkjuO1jA3TXuglM8MBQNugWj17Y\",\"key\":\"SDK_CONFIG_INTERVAL\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"86400000\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-08-19T08:00:44+00:00\"},{\"etag\":\"dSSZ8jtsQTUYpGLwBbMlI9MgsPF\",\"key\":\"SDK_CONTRACT_ABI\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"[{\\\"inputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"_checkpoint\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"name\\\":\\\"addCheckpoint\\\",\\\"outputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"account\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"grantRole\\\",\\\"outputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"constructor\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"_checkpoint\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"name\\\":\\\"NewCheckpoint\\\",\\\"type\\\":\\\"event\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"account\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"renounceRole\\\",\\\"outputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"account\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"revokeRole\\\",\\\"outputs\\\":[],\\\"stateMutability\\\":\\\"nonpayable\\\",\\\"type\\\":\\\"function\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"previousAdminRole\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"newAdminRole\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"name\\\":\\\"RoleAdminChanged\\\",\\\"type\\\":\\\"event\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"account\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"sender\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"RoleGranted\\\",\\\"type\\\":\\\"event\\\"},{\\\"anonymous\\\":false,\\\"inputs\\\":[{\\\"indexed\\\":true,\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"account\\\",\\\"type\\\":\\\"address\\\"},{\\\"indexed\\\":true,\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"sender\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"RoleRevoked\\\",\\\"type\\\":\\\"event\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"DEFAULT_ADMIN_ROLE\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"_checkpoint\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"name\\\":\\\"getCheckpoint\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bool\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bool\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"name\\\":\\\"getRoleAdmin\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"index\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"name\\\":\\\"getRoleMember\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"address\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"name\\\":\\\"getRoleMemberCount\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"uint256\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"uint256\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"role\\\",\\\"type\\\":\\\"bytes32\\\"},{\\\"internalType\\\":\\\"address\\\",\\\"name\\\":\\\"account\\\",\\\"type\\\":\\\"address\\\"}],\\\"name\\\":\\\"hasRole\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bool\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bool\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"},{\\\"inputs\\\":[],\\\"name\\\":\\\"WRITER_ROLE\\\",\\\"outputs\\\":[{\\\"internalType\\\":\\\"bytes32\\\",\\\"name\\\":\\\"\\\",\\\"type\\\":\\\"bytes32\\\"}],\\\"stateMutability\\\":\\\"view\\\",\\\"type\\\":\\\"function\\\"}]\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-08-19T07:58:22+00:00\"},{\"etag\":\"8FRxUSyJ0CH1egJHoYSvPxvkrbp\",\"key\":\"SDK_CONTRACT_ADDRESS\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"0x30f834845E62956499889D56B18a2DB8C6D53Bd1\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-08-19T07:57:57+00:00\"},{\"etag\":\"tHKkN0GuLdgHCW47vnPnlSVN3N9\",\"key\":\"SDK_FETCH_ENDPOINT\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"/message/fetch\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-08-19T07:57:39+00:00\"},{\"etag\":\"IpAxEz9LIoxIlBPfvNko01QqdbC\",\"key\":\"SDK_HOST\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"https://api.enchainte.com\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-08-19T07:54:36+00:00\"},{\"etag\":\"9MOUaLf04EW6hChP0DIUoVaoJH6\",\"key\":\"SDK_HTTP_PROVIDER\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"https://rinkeby.infura.io/v3/6d7880a8f4b347ca8953d2715e164241\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-10-13T14:02:46+00:00\"},{\"etag\":\"PwThJKtBviqi750DRfkPX8UDiOV\",\"key\":\"SDK_PROOF_ENDPOINT\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"/proof\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-08-19T07:56:57+00:00\"},{\"etag\":\"BimMX4J3UTIDwPRxuhsRwFmyWMS\",\"key\":\"SDK_PROVIDER\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"wss://rinkeby.infura.io/ws/v3/6d7880a8f4b347ca8953d2715e164241\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-08-19T07:58:46+00:00\"},{\"etag\":\"9HAPMd3ONbd0mxHfpYXctbE3XAW\",\"key\":\"SDK_WAIT_MESSAGE_INTERVAL_DEFAULT\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"1000\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-09-30T10:34:33+00:00\"},{\"etag\":\"U8LG6XCIxfQcCjJ1NZAeERoJUKh\",\"key\":\"SDK_WAIT_MESSAGE_INTERVAL_FACTOR\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"2\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-09-30T10:34:20+00:00\"},{\"etag\":\"P9oOM2fJ86i7MCDT8W0W5e3xS4K\",\"key\":\"SDK_WRITE_ENDPOINT\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"/write\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-08-19T07:56:47+00:00\"},{\"etag\":\"fViJBVW0PF14eVZvl8MxTjjOrtT\",\"key\":\"SDK_WRITE_INTERVAL\",\"label\":\"PROD\",\"content_type\":\"\",\"value\":\"1000\",\"tags\":{},\"locked\":false,\"last_modified\":\"2020-08-19T08:00:06+00:00\"}]}"

	http.EXPECT().GetRequest(gomock.Any(), gomock.Any()).Return([]byte(resp), nil)

	if err := s.RequestSdkParameters(); err != nil {
		assert.FailNow(t, fmt.Sprintf("unexpected error: %s", err.Error()))
	}

	expected := cloud.SdkParams{Host:"https://api.enchainte.com", MessageWrite:"/write", MessageFetch:"/message/fetch", MessageProof:"/proof", SmartContractAddress:"0x30f834845E62956499889D56B18a2DB8C6D53Bd1", ContractAbi:"[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_checkpoint\",\"type\":\"bytes32\"}],\"name\":\"addCheckpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_checkpoint\",\"type\":\"bytes32\"}],\"name\":\"NewCheckpoint\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_checkpoint\",\"type\":\"bytes32\"}],\"name\":\"getCheckpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRITER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]", Web3Provider:"wss://rinkeby.infura.io/ws/v3/6d7880a8f4b347ca8953d2715e164241", WriteInterval:"1000", WaitIntervalFactor:"2", WaitIntervalDefault:"1000", ConfigInterval:"86400000"}

	assert.Equal(t, expected, s.SdkParameters())
}
