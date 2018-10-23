// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package utils

import (
	"strings"

	"github.com/kowala-tech/kcoin/client/accounts/abi"
	"github.com/kowala-tech/kcoin/client/accounts/abi/bind"
	"github.com/kowala-tech/kcoin/client/common"
	"github.com/kowala-tech/kcoin/client/core/types"
)

// NameHashABI is the input ABI used to generate the binding from.
const NameHashABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_node\",\"type\":\"string\"}],\"name\":\"namehash\",\"outputs\":[{\"name\":\"_namehash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NameHashSrcMap is used in order to generate source maps to use when we want to debug bytecode.
const NameHashSrcMap = "{\"contracts\":{\"../../truffle/contracts/utils/NameHash.sol:NameHash\":{\"bin-runtime\":\"7300000000000000000000000000000000000000003014608060405260043610610058576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063098799621461005d575b600080fd5b6100b7600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506100d5565b60405180826000191660001916815260200191505060405180910390f35b60006100df610599565b60006100e9610599565b606060006100f6876102c3565b945060006001029350610108856102f1565b15156102b65761014c6040805190810160405280600181526020017f2e000000000000000000000000000000000000000000000000000000000000008152506102c3565b92506001610163848761030190919063ffffffff16565b0160405190808252806020026020018201604052801561019757816020015b60608152602001906001900390816101825790505b509150600090505b81518110156101eb576101c36101be848761037890919063ffffffff16565b610392565b82828151811015156101d157fe5b90602001906020020181905250808060010191505061019f565b600090505b81518110156102b55783826001838551030381518110151561020e57fe5b906020019060200201516040518082805190602001908083835b60208310151561024d5780518252602082019150602081019050602083039250610228565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206040518083600019166000191681526020018260001916600019168152602001925050506040518091039020935080806001019150506101f0565b5b8395505050505050919050565b6102cb610599565b600060208301905060408051908101604052808451815260200182815250915050919050565b6000808260000151149050919050565b600080826000015161032585600001518660200151866000015187602001516103f4565b0190505b83600001518460200151018111151561037157818060010192505082600001516103698560200151830386600001510383866000015187602001516103f4565b019050610329565b5092915050565b610380610599565b61038b8383836104b0565b5092915050565b606080600083600001516040519080825280601f01601f1916602001820160405280156103ce5781602001602082028038833980820191505090505b5091506020820190506103ea818560200151866000015161054e565b8192505050919050565b6000806000806000888711151561049e576020871115156104555760018760200360080260020a031980875116888b038a018a96505b81838851161461044a5760018701965080600188031061042a578b8b0196505b5050508394506104a4565b8686209150879350600092505b8689038311151561049d57868420905080600019168260001916141561048a578394506104a4565b6001840193508280600101935050610462565b5b88880194505b50505050949350505050565b6104b8610599565b60006104d685600001518660200151866000015187602001516103f4565b90508460200151836020018181525050846020015181038360000181815250508460000151856020015101811415610518576000856000018181525050610543565b8360000151836000015101856000018181510391508181525050836000015181018560200181815250505b829150509392505050565b60005b6020821015156105765782518452602084019350602083019250602082039150610551565b6001826020036101000a0390508019835116818551168181178652505050505050565b6040805190810160405280600081526020016000815250905600a165627a7a72305820bdb7fa7f1254d03496a55a35a971f38f891e40961a467acc8100eadce57bdad70029\",\"srcmap-runtime\":\"96:961:0:-;;;;;;;;;;;;;;;;;;;;;;;;;;;383:671;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;433:17;466:8;;:::i;:::-;502:16;637:9;;:::i;:::-;676;741:6;477:15;:5;:13;:15::i;:::-;466:26;;521:66;502:85;;;;601:12;:4;:10;:12::i;:::-;600:13;597:426;;;649:13;:11;;;;;;;;;;;;;;;;;;;:13::i;:::-;637:25;;721:1;701:17;712:5;701:4;:10;;:17;;;;:::i;:::-;:21;688:35;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;676:47;;750:1;741:10;;737:123;757:5;:12;753:1;:16;737:123;;;817:28;:17;828:5;817:4;:10;;:17;;;;:::i;:::-;:26;:28::i;:::-;806:5;812:1;806:8;;;;;;;;;;;;;;;;;:39;;;;771:3;;;;;;;737:123;;;882:1;878:5;;874:139;889:5;:12;885:1;:16;874:139;;;954:8;969:5;994:1;990;975:5;:12;:16;:20;969:27;;;;;;;;;;;;;;;;;;964:33;;;;;;;;;;;;;36:153:-1;66:2;61:3;58:11;51:19;36:153;;;182:3;176:10;171:3;164:23;98:2;93:3;89:12;82:19;;123:2;118:3;114:12;107:19;;148:2;143:3;139:12;132:19;;36:153;;;274:1;267:3;263:2;259:12;254:3;250:22;246:30;315:4;311:9;305:3;299:10;295:26;356:4;350:3;344:10;340:21;389:7;380;377:20;372:3;365:33;3:399;;;964:33:0;;;;;;;;;;;;;;;;949:49;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;938:60;;903:3;;;;;;;874:139;;;597:426;1039:8;1032:15;;383:671;;;;;;;;:::o;2805:191:1:-;2853:5;;:::i;:::-;2870:8;2928:4;2922;2918:15;2911:22;;2959:30;;;;;;;;;2971:4;2965:18;2959:30;;;;2985:3;2959:30;;;2952:37;;2805:191;;;;:::o;6619:89::-;6664:4;6700:1;6687:4;:9;;;:14;6680:21;;6619:89;;;:::o;22313:349::-;22372:10;22394:8;22463:6;:11;;;22405:55;22413:4;:9;;;22424:4;:9;;;22435:6;:11;;;22448:6;:11;;;22405:7;:55::i;:::-;:69;22394:80;;22484:172;22510:4;:9;;;22498:4;:9;;;:21;22491:3;:28;;22484:172;;;22535:7;;;;;;;22634:6;:11;;;22562:69;22589:4;:9;;;22583:3;:15;22570:4;:9;;;:29;22601:3;22606:6;:11;;;22619:6;:11;;;22562:7;:69::i;:::-;:83;22556:89;;22484:172;;;22313:349;;;;;:::o;20412:115::-;20471:11;;:::i;:::-;20494:26;20500:4;20506:6;20514:5;20494;:26::i;:::-;;20412:115;;;;:::o;5094:236::-;5142:6;5160:7;5201:11;5181:4;:9;;;5170:21;;;;;;;;;;;;;;;;;;;;;;;;;29:2:-1;21:6;17:15;117:4;105:10;97:6;88:34;148:4;140:6;136:17;126:27;;0:157;5170:21:1;;;;5160:31;;5252:2;5247:3;5243:12;5233:22;;5267:36;5274:6;5282:4;:9;;;5293:4;:9;;;5267:6;:36::i;:::-;5320:3;5313:10;;5094:236;;;;;:::o;14784:1428::-;14878:4;14894:8;14912;15743:12;15939:16;14948:7;14935:9;:20;;14931:1241;;;14988:2;14975:9;:15;;14971:1191;;;15172:1;15158:9;15154:2;15150:18;15147:1;15143:26;15140:1;15136:34;15132:42;15128:47;15236:4;15224:9;15218:16;15214:27;15299:9;15290:7;15286:23;15277:7;15273:37;15338:7;15331:14;;15366:4;15430:10;15423:4;15417:3;15411:10;15407:21;15404:37;15398:4;15392:50;15479:1;15474:3;15470:11;15463:18;;15530:3;15526:1;15521:3;15517:11;15514:20;15508:4;15502:33;15576:7;15567;15563:21;15556:28;;15605:4;15094:534;;;15652:3;15645:10;;;;14971:1191;15808:9;15797;15792:26;15784:34;;15843:7;15837:13;;15879:1;15873:7;;15868:280;15899:9;15889:7;:19;15882:3;:26;;15868:280;;;16010:9;16005:3;16000:20;15988:32;;16055:8;16047:16;;;:4;:16;;;;16043:56;;;16096:3;16089:10;;;;16043:56;16128:1;16121:8;;;;15910:5;;;;;;;15868:280;;;14931:1241;16198:7;16188;:17;16181:24;;14784:1428;;;;;;;;;;;:::o;19470:471::-;19542:5;;:::i;:::-;19559:8;19570:55;19578:4;:9;;;19589:4;:9;;;19600:6;:11;;;19613:6;:11;;;19570:7;:55::i;:::-;19559:66;;19648:4;:9;;;19635:5;:10;;:22;;;;;19686:4;:9;;;19680:3;:15;19667:5;:10;;:28;;;;;19728:4;:9;;;19716:4;:9;;;:21;19709:3;:28;19705:208;;;19790:1;19778:4;:9;;:13;;;;;19705:208;;;19848:6;:11;;;19835:5;:10;;;:24;19822:4;:9;;:37;;;;;;;;;;;19891:6;:11;;;19885:3;:17;19873:4;:9;;:29;;;;;19705:208;19929:5;19922:12;;19470:471;;;;;;:::o;2062:543::-;2383:9;2177:164;2190:2;2183:3;:9;;2177:164;;;2265:3;2259:10;2253:4;2246:24;2305:2;2297:10;;;;2328:2;2321:9;;;;2201:2;2194:9;;;;2177:164;;;2415:1;2408:3;2403:2;:8;2395:3;:17;:21;2383:33;;2484:4;2480:9;2474:3;2468:10;2464:26;2536:4;2529;2523:11;2519:22;2580:7;2570:8;2567:21;2561:4;2554:35;2435:164;;;;;;:::o;96:961:0:-;;;;;;;;;;;;;;;;;;;;:::o\"},\"../../truffle/contracts/utils/Strings.sol:strings\":{\"bin-runtime\":\"73000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820e1e2bd17953348f76d7aa879fc30ad56ca23a2304dcedff5c6fdc06cc10ae3dd0029\",\"srcmap-runtime\":\"1976:22708:1:-;;;;;;;;\"}},\"sourceList\":[\"../../truffle/contracts/utils/NameHash.sol\",\"../../truffle/contracts/utils/Strings.sol\"],\"version\":\"0.4.24+commit.e67f0147.Darwin.appleclang\"}"

// NameHashBin is the compiled bytecode used for deploying new contracts.
const NameHashBin = `6105df610030600b82828239805160001a6073146000811461002057610022565bfe5b5030600052607381538281f3007300000000000000000000000000000000000000003014608060405260043610610058576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063098799621461005d575b600080fd5b6100b7600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506100d5565b60405180826000191660001916815260200191505060405180910390f35b60006100df610599565b60006100e9610599565b606060006100f6876102c3565b945060006001029350610108856102f1565b15156102b65761014c6040805190810160405280600181526020017f2e000000000000000000000000000000000000000000000000000000000000008152506102c3565b92506001610163848761030190919063ffffffff16565b0160405190808252806020026020018201604052801561019757816020015b60608152602001906001900390816101825790505b509150600090505b81518110156101eb576101c36101be848761037890919063ffffffff16565b610392565b82828151811015156101d157fe5b90602001906020020181905250808060010191505061019f565b600090505b81518110156102b55783826001838551030381518110151561020e57fe5b906020019060200201516040518082805190602001908083835b60208310151561024d5780518252602082019150602081019050602083039250610228565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206040518083600019166000191681526020018260001916600019168152602001925050506040518091039020935080806001019150506101f0565b5b8395505050505050919050565b6102cb610599565b600060208301905060408051908101604052808451815260200182815250915050919050565b6000808260000151149050919050565b600080826000015161032585600001518660200151866000015187602001516103f4565b0190505b83600001518460200151018111151561037157818060010192505082600001516103698560200151830386600001510383866000015187602001516103f4565b019050610329565b5092915050565b610380610599565b61038b8383836104b0565b5092915050565b606080600083600001516040519080825280601f01601f1916602001820160405280156103ce5781602001602082028038833980820191505090505b5091506020820190506103ea818560200151866000015161054e565b8192505050919050565b6000806000806000888711151561049e576020871115156104555760018760200360080260020a031980875116888b038a018a96505b81838851161461044a5760018701965080600188031061042a578b8b0196505b5050508394506104a4565b8686209150879350600092505b8689038311151561049d57868420905080600019168260001916141561048a578394506104a4565b6001840193508280600101935050610462565b5b88880194505b50505050949350505050565b6104b8610599565b60006104d685600001518660200151866000015187602001516103f4565b90508460200151836020018181525050846020015181038360000181815250508460000151856020015101811415610518576000856000018181525050610543565b8360000151836000015101856000018181510391508181525050836000015181018560200181815250505b829150509392505050565b60005b6020821015156105765782518452602084019350602083019250602082039150610551565b6001826020036101000a0390508019835116818551168181178652505050505050565b6040805190810160405280600081526020016000815250905600a165627a7a72305820bdb7fa7f1254d03496a55a35a971f38f891e40961a467acc8100eadce57bdad70029`

// DeployNameHash deploys a new Kowala contract, binding an instance of NameHash to it.
func DeployNameHash(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NameHash, error) {
	parsed, err := abi.JSON(strings.NewReader(NameHashABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NameHashBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NameHash{NameHashCaller: NameHashCaller{contract: contract}, NameHashTransactor: NameHashTransactor{contract: contract}, NameHashFilterer: NameHashFilterer{contract: contract}}, nil
}

// NameHash is an auto generated Go binding around a Kowala contract.
type NameHash struct {
	NameHashCaller     // Read-only binding to the contract
	NameHashTransactor // Write-only binding to the contract
	NameHashFilterer   // Log filterer for contract events
}

// NameHashCaller is an auto generated read-only Go binding around a Kowala contract.
type NameHashCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameHashTransactor is an auto generated write-only Go binding around a Kowala contract.
type NameHashTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameHashFilterer is an auto generated log filtering Go binding around a Kowala contract events.
type NameHashFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameHashSession is an auto generated Go binding around a Kowala contract,
// with pre-set call and transact options.
type NameHashSession struct {
	Contract     *NameHash         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameHashCallerSession is an auto generated read-only Go binding around a Kowala contract,
// with pre-set call options.
type NameHashCallerSession struct {
	Contract *NameHashCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// NameHashTransactorSession is an auto generated write-only Go binding around a Kowala contract,
// with pre-set transact options.
type NameHashTransactorSession struct {
	Contract     *NameHashTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NameHashRaw is an auto generated low-level Go binding around a Kowala contract.
type NameHashRaw struct {
	Contract *NameHash // Generic contract binding to access the raw methods on
}

// NameHashCallerRaw is an auto generated low-level read-only Go binding around a Kowala contract.
type NameHashCallerRaw struct {
	Contract *NameHashCaller // Generic read-only contract binding to access the raw methods on
}

// NameHashTransactorRaw is an auto generated low-level write-only Go binding around a Kowala contract.
type NameHashTransactorRaw struct {
	Contract *NameHashTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNameHash creates a new instance of NameHash, bound to a specific deployed contract.
func NewNameHash(address common.Address, backend bind.ContractBackend) (*NameHash, error) {
	contract, err := bindNameHash(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NameHash{NameHashCaller: NameHashCaller{contract: contract}, NameHashTransactor: NameHashTransactor{contract: contract}, NameHashFilterer: NameHashFilterer{contract: contract}}, nil
}

// NewNameHashCaller creates a new read-only instance of NameHash, bound to a specific deployed contract.
func NewNameHashCaller(address common.Address, caller bind.ContractCaller) (*NameHashCaller, error) {
	contract, err := bindNameHash(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NameHashCaller{contract: contract}, nil
}

// NewNameHashTransactor creates a new write-only instance of NameHash, bound to a specific deployed contract.
func NewNameHashTransactor(address common.Address, transactor bind.ContractTransactor) (*NameHashTransactor, error) {
	contract, err := bindNameHash(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NameHashTransactor{contract: contract}, nil
}

// NewNameHashFilterer creates a new log filterer instance of NameHash, bound to a specific deployed contract.
func NewNameHashFilterer(address common.Address, filterer bind.ContractFilterer) (*NameHashFilterer, error) {
	contract, err := bindNameHash(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NameHashFilterer{contract: contract}, nil
}

// bindNameHash binds a generic wrapper to an already deployed contract.
func bindNameHash(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameHashABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameHash *NameHashRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameHash.Contract.NameHashCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameHash *NameHashRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameHash.Contract.NameHashTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameHash *NameHashRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameHash.Contract.NameHashTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NameHash *NameHashCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NameHash.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NameHash *NameHashTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NameHash.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NameHash *NameHashTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NameHash.Contract.contract.Transact(opts, method, params...)
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(_node string) constant returns(_namehash bytes32)
func (_NameHash *NameHashCaller) Namehash(opts *bind.CallOpts, _node string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _NameHash.contract.Call(opts, out, "namehash", _node)
	return *ret0, err
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(_node string) constant returns(_namehash bytes32)
func (_NameHash *NameHashSession) Namehash(_node string) ([32]byte, error) {
	return _NameHash.Contract.Namehash(&_NameHash.CallOpts, _node)
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(_node string) constant returns(_namehash bytes32)
func (_NameHash *NameHashCallerSession) Namehash(_node string) ([32]byte, error) {
	return _NameHash.Contract.Namehash(&_NameHash.CallOpts, _node)
}
