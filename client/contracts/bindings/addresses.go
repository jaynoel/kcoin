package bindings

import (
	"fmt"

	"github.com/kowala-tech/kcoin/client/common"
)

var ProxyFactoryAddr = common.HexToAddress("0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa")
var ProxyKNSRegistryAddr = common.HexToAddress("0x4195B06a6e4d5bEDDe15165e01A64F324f03D5d1")
var ProxyRegistrarAddr = common.HexToAddress("0xe1adb6075619f52Fc00BDD50dEf1B754b9e7bd17")
var ProxyResolverAddr = common.HexToAddress("0x01e1056f6a829E53dadeb8a5A6189A9333Bd1d63")

var MultiSigWalletAddr = common.HexToAddress("0x0e5d0Fd336650E663C710EF420F85Fb081E21415")

var contracts = map[common.Address]string{
	ProxyFactoryAddr:     "Proxy Factory Contract",
	ProxyKNSRegistryAddr: "Proxy KNS Registry Contract",
	ProxyRegistrarAddr:   "Proxy Registrar Contract",
	ProxyResolverAddr:    "Proxy Resolver Contract",
	MultiSigWalletAddr:   "Multisig Wallet Contract",
}

func GetContractByAddr(addr common.Address) (string, error) {
	addrName, ok := contracts[addr]
	if !ok {
		return "", fmt.Errorf("name of contract with addr %s cannot be derived", addr.String())
	}

	return addrName, nil
}
