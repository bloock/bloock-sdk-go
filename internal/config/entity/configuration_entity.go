package entity

type NetworkConfiguration struct {
	ContractAddress string `default:""`
	ContractABI     string `default:""`
	HttpProvider    string `default:""`
}

func NewNetworkConfiguration(contractAddress, contractABI, httpProvider string) NetworkConfiguration {
	return NetworkConfiguration{
		ContractAddress: contractAddress,
		ContractABI:     contractABI,
		HttpProvider:    httpProvider,
	}
}

type Configuration struct {
	Host                       string `default:""`
	WaitMessageIntervalFactor  int    `default:"2"`
	WaitMessageIntervalDefault int    `default:"100"`
}

func NewConfiguration(host string, waitMsgFactor, waitMsgDefault int) Configuration {
	return Configuration{
		Host:                       host,
		WaitMessageIntervalFactor:  waitMsgFactor,
		WaitMessageIntervalDefault: waitMsgDefault,
	}
}
