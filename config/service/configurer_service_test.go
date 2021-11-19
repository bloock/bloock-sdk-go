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

	cr := mockconfig.NewMockConfigurerRepository(crtl)
	cs := NewConfigService(cr)

	t.Run("Given a network should return a network configuration", func(t *testing.T) {
		cr.EXPECT().GetNetworkConfiguration(gomock.Any()).Times(1)

		netConfig := cs.GetNetworkConfiguration(entity.EthereumMainnet)
		assert.NotNil(t, netConfig)
	})
}

func TestGetApiBaseUrl(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	cr := mockconfig.NewMockConfigurerRepository(crtl)
	cs := NewConfigService(cr)
	
	t.Run("Given an api host, should return that api host", func(t *testing.T) {
		config := entity.NewConfiguration("test", 2, 100)

		cr.EXPECT().GetConfiguration().Return(config)

		actual := cs.GetApiBaseUrl()
		assert.Equal(t, "test", actual)
	})
}
