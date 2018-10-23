// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

import (
	"strings"

	kowala "github.com/kowala-tech/kcoin/client"
	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
	"github.com/kowala-tech/kcoin/client/event"
)

// UpgradeabilityProxyFactoryABI is the input ABI used to generate the binding from.
const UpgradeabilityProxyFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"},{\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"createProxy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin\",\"type\":\"address\"},{\"name\":\"implementation\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createProxyAndCall\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"}]"

// UpgradeabilityProxyFactorySrcMap is used in order to generate source maps to use when we want to debug bytecode.
const UpgradeabilityProxyFactorySrcMap = "{\"contracts\":{\"../../truffle/node_modules/openzeppelin-solidity/contracts/AddressUtils.sol:AddressUtils\":{\"bin-runtime\":\"73000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820a5042c15ef9e031697d5ab042d6c3478569d6f76de3df381230c34a6560c2ba40029\",\"srcmap-runtime\":\"87:932:0:-;;;;;;;;\"},\"../../truffle/node_modules/zos-lib/contracts/upgradeability/AdminUpgradeabilityProxy.sol:AdminUpgradeabilityProxy\":{\"bin-runtime\":\"60806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633659cfe6146100775780634f1ef286146100ba5780635c60da1b146101085780638f2839701461015f578063f851a440146101a2575b6100756101f9565b005b34801561008357600080fd5b506100b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610213565b005b610106600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001919091929391929390505050610268565b005b34801561011457600080fd5b5061011d610308565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016b57600080fd5b506101a0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610360565b005b3480156101ae57600080fd5b506101b761051e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610201610576565b61021161020c610651565b610682565b565b61021b6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561025c57610257816106d9565b610265565b6102646101f9565b5b50565b6102706106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102fa576102ac836106d9565b3073ffffffffffffffffffffffffffffffffffffffff163483836040518083838082843782019150509250505060006040518083038185875af19250505015156102f557600080fd5b610303565b6103026101f9565b5b505050565b60006103126106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103545761034d610651565b905061035d565b61035c6101f9565b5b90565b6103686106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561051257600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610466576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f81526020017f787920746f20746865207a65726f20616464726573730000000000000000000081525060400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61048f6106a8565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161050d81610748565b61051b565b61051a6101f9565b5b50565b60006105286106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561056a576105636106a8565b9050610573565b6105726101f9565b5b90565b61057e6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610647576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e20667281526020017f6f6d207468652070726f78792061646d696e000000000000000000000000000081525060400191505060405180910390fd5b61064f610777565b565b6000807f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c36001029050805491505090565b3660008037600080366000845af43d6000803e80600081146106a3573d6000f35b3d6000fd5b6000807f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001029050805491505090565b6106e281610779565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b565b60006107848261084b565b151561081e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b600080823b9050600081119150509190505600a165627a7a72305820b31d41a72e81559b9fbd6f52c3db76ab77aeacd5de378d37ab93861a9b684f120029\",\"srcmap-runtime\":\"422:3596:1:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;454:11:2;:9;:11::i;:::-;422:3596:1;2543:103;;8:9:-1;5:2;;;30:1;27;20:12;5:2;2543:103:1;;;;;;;;;;;;;;;;;;;;;;;;;;;;3168:186;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1858:101;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1858:101:1;;;;;;;;;;;;;;;;;;;;;;;;;;;2135:224;;8:9:-1;5:2;;;30:1;27;20:12;5:2;2135:224:1;;;;;;;;;;;;;;;;;;;;;;;;;;;;1711:83;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1711:83:1;;;;;;;;;;;;;;;;;;;;;;;;;;;1953:90:2;1989:15;:13;:15::i;:::-;2010:28;2020:17;:15;:17::i;:::-;2010:9;:28::i;:::-;1953:90::o;2543:103:1:-;1239:8;:6;:8::i;:::-;1225:22;;:10;:22;;;1221:76;;;2612:29;2623:17;2612:10;:29::i;:::-;1221:76;;;1279:11;:9;:11::i;:::-;1221:76;2543:103;:::o;3168:186::-;1239:8;:6;:8::i;:::-;1225:22;;:10;:22;;;1221:76;;;3264:29;3275:17;3264:10;:29::i;:::-;3315:4;3307:18;;3332:9;3343:4;;3307:41;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;3299:50;;;;;;;;1221:76;;;1279:11;:9;:11::i;:::-;1221:76;3168:186;;;:::o;1858:101::-;1915:7;1239:8;:6;:8::i;:::-;1225:22;;:10;:22;;;1221:76;;;1937:17;:15;:17::i;:::-;1930:24;;1221:76;;;1279:11;:9;:11::i;:::-;1221:76;1858:101;:::o;2135:224::-;1239:8;:6;:8::i;:::-;1225:22;;:10;:22;;;1221:76;;;2225:1;2205:22;;:8;:22;;;;2197:89;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;2297:32;2310:8;:6;:8::i;:::-;2320;2297:32;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;2335:19;2345:8;2335:9;:19::i;:::-;1221:76;;;1279:11;:9;:11::i;:::-;1221:76;2135:224;:::o;1711:83::-;1759:7;1239:8;:6;:8::i;:::-;1225:22;;:10;:22;;;1221:76;;;1781:8;:6;:8::i;:::-;1774:15;;1221:76;;;1279:11;:9;:11::i;:::-;1221:76;1711:83;:::o;3859:157::-;3921:8;:6;:8::i;:::-;3907:22;;:10;:22;;;;3899:85;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;3990:21;:19;:21::i;:::-;3859:157::o;1252:156:3:-;1302:12;1322;781:66;1337:19;;1322:34;;1393:4;1387:11;1379:19;;1371:33;;:::o;879:731:2:-;1181:12;1178:1;1175;1162:32;1371:1;1368;1354:12;1351:1;1335:14;1330:3;1317:56;1435:14;1432:1;1429;1414:36;1465:6;1525:1;1520:36;;;;1583:14;1580:1;1573:25;1520:36;1539:14;1536:1;1529:25;3399:136:1;3440:11;3459:12;940:66;3474:10;;3459:25;;3520:4;3514:11;3507:18;;3499:32;;:::o;1543:142:3:-;1605:37;1624:17;1605:18;:37::i;:::-;1653:27;1662:17;1653:27;;;;;;;;;;;;;;;;;;;;;;1543:142;:::o;3651:133:1:-;3703:12;940:66;3718:10;;3703:25;;3765:8;3759:4;3752:22;3744:36;;:::o;1819:39:2:-;:::o;1822:289:3:-;2012:12;1899:42;1923:17;1899:23;:42::i;:::-;1891:114;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;781:66;2027:19;;2012:34;;2083:17;2077:4;2070:31;2062:45;;:::o;438:578:0:-;496:4;508:12;983:5;971:18;963:26;;1010:1;1003:4;:8;996:15;;438:578;;;;:::o\"},\"../../truffle/node_modules/zos-lib/contracts/upgradeability/Proxy.sol:Proxy\":{\"bin-runtime\":\"\",\"srcmap-runtime\":\"\"},\"../../truffle/node_modules/zos-lib/contracts/upgradeability/UpgradeabilityProxy.sol:UpgradeabilityProxy\":{\"bin-runtime\":\"6080604052600a600c565b005b60126020565b601e601a6022565b6053565b565b565b6000807f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c36001029050805491505090565b3660008037600080366000845af43d6000803e80600081146073573d6000f35b3d6000fd5b600080823b9050600081119150509190505600a165627a7a72305820513badeff95daa45b0d89589c79c0964b572d57ae18f283b4016bbebb264468e0029\",\"srcmap-runtime\":\"320:1793:3:-;;;454:11:2;:9;:11::i;:::-;320:1793:3;1953:90:2;1989:15;:13;:15::i;:::-;2010:28;2020:17;:15;:17::i;:::-;2010:9;:28::i;:::-;1953:90::o;1819:39::-;:::o;1252:156:3:-;1302:12;1322;781:66;1337:19;;1322:34;;1393:4;1387:11;1379:19;;1371:33;;:::o;879:731:2:-;1181:12;1178:1;1175;1162:32;1371:1;1368;1354:12;1351:1;1335:14;1330:3;1317:56;1435:14;1432:1;1429;1414:36;1465:6;1525:1;1520:36;;;;1583:14;1580:1;1573:25;1520:36;1539:14;1536:1;1529:25;438:578:0;496:4;508:12;983:5;971:18;963:26;;1010:1;1003:4;:8;996:15;;438:578;;;;:::o\"},\"../../truffle/node_modules/zos-lib/contracts/upgradeability/UpgradeabilityProxyFactory.sol:UpgradeabilityProxyFactory\":{\"bin-runtime\":\"60806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806325b5672714610051578063c6e8b4f3146100f4575b600080fd5b34801561005d57600080fd5b506100b2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101d0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61018e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061029b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000806101dc836103fa565b90508073ffffffffffffffffffffffffffffffffffffffff16638f283970856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561027957600080fd5b505af115801561028d573d6000803e3d6000fd5b505050508091505092915050565b6000806102a7846103fa565b90508073ffffffffffffffffffffffffffffffffffffffff16638f283970866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561034457600080fd5b505af1158015610358573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16348460405180828051906020019080838360005b838110156103a2578082015181840152602081019050610387565b50505050905090810190601f1680156103cf5780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af19250505015156103ef57600080fd5b809150509392505050565b600080826104066104c6565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f080158015610458573d6000803e3d6000fd5b5090507efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e734981604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a180915050919050565b604051610b27806104d7833901905600608060405234801561001057600080fd5b50604051602080610b27833981018060405281019080805190602001909291905050508060405180807f6f72672e7a657070656c696e6f732e70726f78792e696d706c656d656e74617481526020017f696f6e000000000000000000000000000000000000000000000000000000000081525060230190506040518091039020600019167f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3600102600019161415156100c557fe5b6100dd81610167640100000000026401000000009004565b5060405180807f6f72672e7a657070656c696e6f732e70726f78792e61646d696e000000000000815250601a0190506040518091039020600019167f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001026000191614151561014957fe5b6101613361024c640100000000026401000000009004565b5061028e565b60006101858261027b6401000000000261084b176401000000009004565b151561021f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b600080823b905060008111915050919050565b61088a8061029d6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633659cfe6146100775780634f1ef286146100ba5780635c60da1b146101085780638f2839701461015f578063f851a440146101a2575b6100756101f9565b005b34801561008357600080fd5b506100b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610213565b005b610106600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001919091929391929390505050610268565b005b34801561011457600080fd5b5061011d610308565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016b57600080fd5b506101a0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610360565b005b3480156101ae57600080fd5b506101b761051e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610201610576565b61021161020c610651565b610682565b565b61021b6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561025c57610257816106d9565b610265565b6102646101f9565b5b50565b6102706106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102fa576102ac836106d9565b3073ffffffffffffffffffffffffffffffffffffffff163483836040518083838082843782019150509250505060006040518083038185875af19250505015156102f557600080fd5b610303565b6103026101f9565b5b505050565b60006103126106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103545761034d610651565b905061035d565b61035c6101f9565b5b90565b6103686106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561051257600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610466576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f81526020017f787920746f20746865207a65726f20616464726573730000000000000000000081525060400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61048f6106a8565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161050d81610748565b61051b565b61051a6101f9565b5b50565b60006105286106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561056a576105636106a8565b9050610573565b6105726101f9565b5b90565b61057e6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610647576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e20667281526020017f6f6d207468652070726f78792061646d696e000000000000000000000000000081525060400191505060405180910390fd5b61064f610777565b565b6000807f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c36001029050805491505090565b3660008037600080366000845af43d6000803e80600081146106a3573d6000f35b3d6000fd5b6000807f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001029050805491505090565b6106e281610779565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b565b60006107848261084b565b151561081e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b600080823b9050600081119150509190505600a165627a7a72305820b31d41a72e81559b9fbd6f52c3db76ab77aeacd5de378d37ab93861a9b684f120029a165627a7a7230582075bac8779341600a0dbcebbbdf6fb8487503343b22b6218c9bd744b22b6435780029\",\"srcmap-runtime\":\"163:1944:4:-;;;;;;;;;;;;;;;;;;;;;;;;;;;;;585:222;;8:9:-1;5:2;;;30:1;27;20:12;5:2;585:222:4;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1392:306;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;585:222;661:24;693:30;726:28;739:14;726:12;:28::i;:::-;693:61;;760:5;:17;;;778:5;760:24;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;8:9:-1;5:2;;;30:1;27;20:12;5:2;760:24:4;;;;8:9:-1;5:2;;;45:16;42:1;39;24:38;77:16;74:1;67:27;5:2;760:24:4;;;;797:5;790:12;;585:222;;;;;:::o;1392:306::-;1495:24;1527:30;1560:28;1573:14;1560:12;:28::i;:::-;1527:61;;1594:5;:17;;;1612:5;1594:24;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;8:9:-1;5:2;;;30:1;27;20:12;5:2;1594:24:4;;;;8:9:-1;5:2;;;45:16;42:1;39;24:38;77:16;74:1;67:27;5:2;1594:24:4;;;;1640:5;1632:19;;1658:9;1669:4;1632:42;;;;;;;;;;;;;23:1:-1;8:100;33:3;30:1;27:10;8:100;;;99:1;94:3;90:11;84:18;80:1;75:3;71:11;64:39;52:2;49:1;45:10;40:15;;8:100;;;12:14;1632:42:4;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;1624:51;;;;;;;;1688:5;1681:12;;1392:306;;;;;;:::o;1879:226::-;1943:24;1975:30;2037:14;2008:44;;:::i;:::-;;;;;;;;;;;;;;;;;;;;;;;;8:9:-1;5:2;;;45:16;42:1;39;24:38;77:16;74:1;67:27;5:2;2008:44:4;1975:77;;2063:19;2076:5;2063:19;;;;;;;;;;;;;;;;;;;;;;2095:5;2088:12;;1879:226;;;;:::o;163:1944::-;;;;;;;;;;:::o\"}},\"sourceList\":[\"../../truffle/node_modules/openzeppelin-solidity/contracts/AddressUtils.sol\",\"../../truffle/node_modules/zos-lib/contracts/upgradeability/AdminUpgradeabilityProxy.sol\",\"../../truffle/node_modules/zos-lib/contracts/upgradeability/Proxy.sol\",\"../../truffle/node_modules/zos-lib/contracts/upgradeability/UpgradeabilityProxy.sol\",\"../../truffle/node_modules/zos-lib/contracts/upgradeability/UpgradeabilityProxyFactory.sol\"],\"version\":\"0.4.24+commit.e67f0147.Darwin.appleclang\"}"

// UpgradeabilityProxyFactoryBin is the compiled bytecode used for deploying new contracts.
const UpgradeabilityProxyFactoryBin = `608060405234801561001057600080fd5b50611029806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806325b5672714610051578063c6e8b4f3146100f4575b600080fd5b34801561005d57600080fd5b506100b2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101d0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61018e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061029b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000806101dc836103fa565b90508073ffffffffffffffffffffffffffffffffffffffff16638f283970856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561027957600080fd5b505af115801561028d573d6000803e3d6000fd5b505050508091505092915050565b6000806102a7846103fa565b90508073ffffffffffffffffffffffffffffffffffffffff16638f283970866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561034457600080fd5b505af1158015610358573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16348460405180828051906020019080838360005b838110156103a2578082015181840152602081019050610387565b50505050905090810190601f1680156103cf5780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af19250505015156103ef57600080fd5b809150509392505050565b600080826104066104c6565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f080158015610458573d6000803e3d6000fd5b5090507efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e734981604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a180915050919050565b604051610b27806104d7833901905600608060405234801561001057600080fd5b50604051602080610b27833981018060405281019080805190602001909291905050508060405180807f6f72672e7a657070656c696e6f732e70726f78792e696d706c656d656e74617481526020017f696f6e000000000000000000000000000000000000000000000000000000000081525060230190506040518091039020600019167f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3600102600019161415156100c557fe5b6100dd81610167640100000000026401000000009004565b5060405180807f6f72672e7a657070656c696e6f732e70726f78792e61646d696e000000000000815250601a0190506040518091039020600019167f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001026000191614151561014957fe5b6101613361024c640100000000026401000000009004565b5061028e565b60006101858261027b6401000000000261084b176401000000009004565b151561021f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b600080823b905060008111915050919050565b61088a8061029d6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633659cfe6146100775780634f1ef286146100ba5780635c60da1b146101085780638f2839701461015f578063f851a440146101a2575b6100756101f9565b005b34801561008357600080fd5b506100b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610213565b005b610106600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001919091929391929390505050610268565b005b34801561011457600080fd5b5061011d610308565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016b57600080fd5b506101a0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610360565b005b3480156101ae57600080fd5b506101b761051e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610201610576565b61021161020c610651565b610682565b565b61021b6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561025c57610257816106d9565b610265565b6102646101f9565b5b50565b6102706106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102fa576102ac836106d9565b3073ffffffffffffffffffffffffffffffffffffffff163483836040518083838082843782019150509250505060006040518083038185875af19250505015156102f557600080fd5b610303565b6103026101f9565b5b505050565b60006103126106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103545761034d610651565b905061035d565b61035c6101f9565b5b90565b6103686106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561051257600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610466576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f81526020017f787920746f20746865207a65726f20616464726573730000000000000000000081525060400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61048f6106a8565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161050d81610748565b61051b565b61051a6101f9565b5b50565b60006105286106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561056a576105636106a8565b9050610573565b6105726101f9565b5b90565b61057e6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610647576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e20667281526020017f6f6d207468652070726f78792061646d696e000000000000000000000000000081525060400191505060405180910390fd5b61064f610777565b565b6000807f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c36001029050805491505090565b3660008037600080366000845af43d6000803e80600081146106a3573d6000f35b3d6000fd5b6000807f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001029050805491505090565b6106e281610779565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b565b60006107848261084b565b151561081e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b600080823b9050600081119150509190505600a165627a7a72305820b31d41a72e81559b9fbd6f52c3db76ab77aeacd5de378d37ab93861a9b684f120029a165627a7a7230582075bac8779341600a0dbcebbbdf6fb8487503343b22b6218c9bd744b22b6435780029`

// DeployUpgradeabilityProxyFactory deploys a new Kowala contract, binding an instance of UpgradeabilityProxyFactory to it.
func DeployUpgradeabilityProxyFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UpgradeabilityProxyFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeabilityProxyFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UpgradeabilityProxyFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UpgradeabilityProxyFactory{UpgradeabilityProxyFactoryCaller: UpgradeabilityProxyFactoryCaller{contract: contract}, UpgradeabilityProxyFactoryTransactor: UpgradeabilityProxyFactoryTransactor{contract: contract}, UpgradeabilityProxyFactoryFilterer: UpgradeabilityProxyFactoryFilterer{contract: contract}}, nil
}

// UpgradeabilityProxyFactory is an auto generated Go binding around a Kowala contract.
type UpgradeabilityProxyFactory struct {
	UpgradeabilityProxyFactoryCaller     // Read-only binding to the contract
	UpgradeabilityProxyFactoryTransactor // Write-only binding to the contract
	UpgradeabilityProxyFactoryFilterer   // Log filterer for contract events
}

// UpgradeabilityProxyFactoryCaller is an auto generated read-only Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeabilityProxyFactoryTransactor is an auto generated write-only Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeabilityProxyFactoryFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type UpgradeabilityProxyFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeabilityProxyFactorySession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type UpgradeabilityProxyFactorySession struct {
	Contract     *UpgradeabilityProxyFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// UpgradeabilityProxyFactoryCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type UpgradeabilityProxyFactoryCallerSession struct {
	Contract *UpgradeabilityProxyFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// UpgradeabilityProxyFactoryTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type UpgradeabilityProxyFactoryTransactorSession struct {
	Contract     *UpgradeabilityProxyFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// UpgradeabilityProxyFactoryRaw is an auto generated low-level Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryRaw struct {
	Contract *UpgradeabilityProxyFactory // Generic contract binding to access the raw methods on
}

// UpgradeabilityProxyFactoryCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryCallerRaw struct {
	Contract *UpgradeabilityProxyFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// UpgradeabilityProxyFactoryTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type UpgradeabilityProxyFactoryTransactorRaw struct {
	Contract *UpgradeabilityProxyFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeabilityProxyFactory creates a new instance of UpgradeabilityProxyFactory, bound to a specific deployed contract.
func NewUpgradeabilityProxyFactory(address common.Address, backend bind.ContractBackend) (*UpgradeabilityProxyFactory, error) {
	contract, err := bindUpgradeabilityProxyFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactory{UpgradeabilityProxyFactoryCaller: UpgradeabilityProxyFactoryCaller{contract: contract}, UpgradeabilityProxyFactoryTransactor: UpgradeabilityProxyFactoryTransactor{contract: contract}, UpgradeabilityProxyFactoryFilterer: UpgradeabilityProxyFactoryFilterer{contract: contract}}, nil
}

// NewUpgradeabilityProxyFactoryCaller creates a new read-only instance of UpgradeabilityProxyFactory, bound to a specific deployed contract.
func NewUpgradeabilityProxyFactoryCaller(address common.Address, caller bind.ContractCaller) (*UpgradeabilityProxyFactoryCaller, error) {
	contract, err := bindUpgradeabilityProxyFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactoryCaller{contract: contract}, nil
}

// NewUpgradeabilityProxyFactoryTransactor creates a new write-only instance of UpgradeabilityProxyFactory, bound to a specific deployed contract.
func NewUpgradeabilityProxyFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeabilityProxyFactoryTransactor, error) {
	contract, err := bindUpgradeabilityProxyFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactoryTransactor{contract: contract}, nil
}

// NewUpgradeabilityProxyFactoryFilterer creates a new log filterer instance of UpgradeabilityProxyFactory, bound to a specific deployed contract.
func NewUpgradeabilityProxyFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeabilityProxyFactoryFilterer, error) {
	contract, err := bindUpgradeabilityProxyFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactoryFilterer{contract: contract}, nil
}

// bindUpgradeabilityProxyFactory binds a generic wrapper to an already deployed contract.
func bindUpgradeabilityProxyFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UpgradeabilityProxyFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradeabilityProxyFactory.Contract.UpgradeabilityProxyFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.UpgradeabilityProxyFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.UpgradeabilityProxyFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UpgradeabilityProxyFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x25b56727.
//
// Solidity: function createProxy(admin address, implementation address) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactor) CreateProxy(opts *bind.TransactOpts, admin common.Address, implementation common.Address) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.contract.Transact(opts, "createProxy", admin, implementation)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x25b56727.
//
// Solidity: function createProxy(admin address, implementation address) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactorySession) CreateProxy(admin common.Address, implementation common.Address) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.CreateProxy(&_UpgradeabilityProxyFactory.TransactOpts, admin, implementation)
}

// CreateProxy is a paid mutator transaction binding the contract method 0x25b56727.
//
// Solidity: function createProxy(admin address, implementation address) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactorSession) CreateProxy(admin common.Address, implementation common.Address) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.CreateProxy(&_UpgradeabilityProxyFactory.TransactOpts, admin, implementation)
}

// CreateProxyAndCall is a paid mutator transaction binding the contract method 0xc6e8b4f3.
//
// Solidity: function createProxyAndCall(admin address, implementation address, data bytes) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactor) CreateProxyAndCall(opts *bind.TransactOpts, admin common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.contract.Transact(opts, "createProxyAndCall", admin, implementation, data)
}

// CreateProxyAndCall is a paid mutator transaction binding the contract method 0xc6e8b4f3.
//
// Solidity: function createProxyAndCall(admin address, implementation address, data bytes) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactorySession) CreateProxyAndCall(admin common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.CreateProxyAndCall(&_UpgradeabilityProxyFactory.TransactOpts, admin, implementation, data)
}

// CreateProxyAndCall is a paid mutator transaction binding the contract method 0xc6e8b4f3.
//
// Solidity: function createProxyAndCall(admin address, implementation address, data bytes) returns(address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryTransactorSession) CreateProxyAndCall(admin common.Address, implementation common.Address, data []byte) (*types.Transaction, error) {
	return _UpgradeabilityProxyFactory.Contract.CreateProxyAndCall(&_UpgradeabilityProxyFactory.TransactOpts, admin, implementation, data)
}

// UpgradeabilityProxyFactoryProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the UpgradeabilityProxyFactory contract.
type UpgradeabilityProxyFactoryProxyCreatedIterator struct {
	Event *UpgradeabilityProxyFactoryProxyCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  kowala.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UpgradeabilityProxyFactoryProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UpgradeabilityProxyFactoryProxyCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UpgradeabilityProxyFactoryProxyCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UpgradeabilityProxyFactoryProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UpgradeabilityProxyFactoryProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UpgradeabilityProxyFactoryProxyCreated represents a ProxyCreated event raised by the UpgradeabilityProxyFactory contract.
type UpgradeabilityProxyFactoryProxyCreated struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: e ProxyCreated(proxy address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryFilterer) FilterProxyCreated(opts *bind.FilterOpts) (*UpgradeabilityProxyFactoryProxyCreatedIterator, error) {

	logs, sub, err := _UpgradeabilityProxyFactory.contract.FilterLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return &UpgradeabilityProxyFactoryProxyCreatedIterator{contract: _UpgradeabilityProxyFactory.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x00fffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e7349.
//
// Solidity: e ProxyCreated(proxy address)
func (_UpgradeabilityProxyFactory *UpgradeabilityProxyFactoryFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *UpgradeabilityProxyFactoryProxyCreated) (event.Subscription, error) {

	logs, sub, err := _UpgradeabilityProxyFactory.contract.WatchLogs(opts, "ProxyCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UpgradeabilityProxyFactoryProxyCreated)
				if err := _UpgradeabilityProxyFactory.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
