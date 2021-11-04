package entity

type NetworkConfiguration struct {
	ContractAddress string `default:""`
	ContractABI     string `default:""`
	HttpProvider    string `default:""`
}

func NewNetworkConfiguration(contractAddress, contractABI, httpProvider string) NetworkConfiguration {
	return NetworkConfiguration{
		ContractAddress: contractAddress,
		ContractABI: contractABI,
		HttpProvider: httpProvider,
	}
}

type Configuration struct {
	host 					   string `default:""`
	waitMessageIntervalFactor  int
	waitMessageIntervalDefault int
}

func NewConfiguration(host string) Configuration {
	return Configuration{
		host: host,
		waitMessageIntervalFactor: 2,
		waitMessageIntervalDefault: 1000,
	}
}
