package service

import (
	"github.com/enchainte/enchainte-sdk-go/config/entity"
	"github.com/enchainte/enchainte-sdk-go/config/mockconfig"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNetworkConfigurationService(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	cr := mockconfig.NewMockConfigRepository(crtl)

	t.Run("Given a network should return a network configuration", func(t *testing.T) {
		cr.EXPECT().GetNetworkConfiguration(gomock.Any()).Times(1)

		netConfig := NewConfigServiceImpl(cr).GetNetworkConfiguration(entity.EthereumMainnet)
		assert.NotNil(t, netConfig)
	})
}
