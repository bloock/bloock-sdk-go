package repository

import (
	"github.com/enchainte/enchainte-sdk-go/config/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNetworkConfigurationRepo(t *testing.T) {
	cd := NewConfigData()
	r := NewConfigRepository(cd)

	t.Run("Given a valid network and root should return a network configuration", func(t *testing.T) {
		netConfig := r.GetNetworkConfiguration(entity.EthereumMainnet)

		assert.NotNil(t, netConfig)
	})

	t.Run("Given an invalid network and root should return an ethereum network configuration", func(t *testing.T) {
		netConfig := r.GetNetworkConfiguration("")

		assert.NotNil(t, netConfig)
		assert.Equal(t, "https://mainnet.infura.io/v3/40e23a35d578492daacb318023772b52", netConfig.HttpProvider)
	})
}
