const Tx = require('ethereumjs-tx');
const Web3 = require('web3');

const web3 = new Web3('http://0.0.0.0:30503');

const account1 = '0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0';
const privateKey1 = Buffer.from('ae288a2d269e706e151c5c572b590d7d7fc478c243fde209b019b89054fd2737', 'hex');
const privateKey2 = '0xae288a2d269e706e151c5c572b590d7d7fc478c243fde209b019b89054fd2737';
const UpgradabilityProxyFactoryAbi = [{'constant':false,'inputs':[{'name':'admin','type':'address'},{'name':'implementation','type':'address'}],'name':'createProxy','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'admin','type':'address'},{'name':'implementation','type':'address'},{'name':'data','type':'bytes'}],'name':'createProxyAndCall','outputs':[{'name':'','type':'address'}],'payable':true,'stateMutability':'payable','type':'function'},{'anonymous':false,'inputs':[{'indexed':false,'name':'proxy','type':'address'}],'name':'ProxyCreated','type':'event'}];
const UpgradabilityProxyFactoryBin = '0x608060405234801561001057600080fd5b50611029806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806325b5672714610051578063c6e8b4f3146100f4575b600080fd5b34801561005d57600080fd5b506100b2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101d0565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b61018e600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061029b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000806101dc836103fa565b90508073ffffffffffffffffffffffffffffffffffffffff16638f283970856040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561027957600080fd5b505af115801561028d573d6000803e3d6000fd5b505050508091505092915050565b6000806102a7846103fa565b90508073ffffffffffffffffffffffffffffffffffffffff16638f283970866040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050600060405180830381600087803b15801561034457600080fd5b505af1158015610358573d6000803e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff16348460405180828051906020019080838360005b838110156103a2578082015181840152602081019050610387565b50505050905090810190601f1680156103cf5780820380516001836020036101000a031916815260200191505b5091505060006040518083038185875af19250505015156103ef57600080fd5b809150509392505050565b600080826104066104c6565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f080158015610458573d6000803e3d6000fd5b5090507efffc2da0b561cae30d9826d37709e9421c4725faebc226cbbb7ef5fc5e734981604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a180915050919050565b604051610b27806104d7833901905600608060405234801561001057600080fd5b50604051602080610b27833981018060405281019080805190602001909291905050508060405180807f6f72672e7a657070656c696e6f732e70726f78792e696d706c656d656e74617481526020017f696f6e000000000000000000000000000000000000000000000000000000000081525060230190506040518091039020600019167f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3600102600019161415156100c557fe5b6100dd81610167640100000000026401000000009004565b5060405180807f6f72672e7a657070656c696e6f732e70726f78792e61646d696e000000000000815250601a0190506040518091039020600019167f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001026000191614151561014957fe5b6101613361024c640100000000026401000000009004565b5061028e565b60006101858261027b6401000000000261084b176401000000009004565b151561021f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b600080823b905060008111915050919050565b61088a8061029d6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633659cfe6146100775780634f1ef286146100ba5780635c60da1b146101085780638f2839701461015f578063f851a440146101a2575b6100756101f9565b005b34801561008357600080fd5b506100b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610213565b005b610106600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001919091929391929390505050610268565b005b34801561011457600080fd5b5061011d610308565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016b57600080fd5b506101a0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610360565b005b3480156101ae57600080fd5b506101b761051e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610201610576565b61021161020c610651565b610682565b565b61021b6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561025c57610257816106d9565b610265565b6102646101f9565b5b50565b6102706106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102fa576102ac836106d9565b3073ffffffffffffffffffffffffffffffffffffffff163483836040518083838082843782019150509250505060006040518083038185875af19250505015156102f557600080fd5b610303565b6103026101f9565b5b505050565b60006103126106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103545761034d610651565b905061035d565b61035c6101f9565b5b90565b6103686106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561051257600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610466576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f81526020017f787920746f20746865207a65726f20616464726573730000000000000000000081525060400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61048f6106a8565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161050d81610748565b61051b565b61051a6101f9565b5b50565b60006105286106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561056a576105636106a8565b9050610573565b6105726101f9565b5b90565b61057e6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610647576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e20667281526020017f6f6d207468652070726f78792061646d696e000000000000000000000000000081525060400191505060405180910390fd5b61064f610777565b565b6000807f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c36001029050805491505090565b3660008037600080366000845af43d6000803e80600081146106a3573d6000f35b3d6000fd5b6000807f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001029050805491505090565b6106e281610779565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b565b60006107848261084b565b151561081e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b600080823b9050600081119150509190505600a165627a7a72305820dd6414bbb207177c22db40427e1e892399cabda0589266e33eb2dd35eef8a7040029a165627a7a723058202f1f18f3db6fcafbc4260fe1c6fa2ffe0e7514d168c48a247c4587e34bff27170029';

const AdminUpgradabilityProxyAbi = [{'inputs':[{'name':'_implementation','type':'address'}],'payable':false,'stateMutability':'nonpayable','type':'constructor'},{'payable':true,'stateMutability':'payable','type':'fallback'},{'anonymous':false,'inputs':[{'indexed':false,'name':'previousAdmin','type':'address'},{'indexed':false,'name':'newAdmin','type':'address'}],'name':'AdminChanged','type':'event'},{'anonymous':false,'inputs':[{'indexed':false,'name':'implementation','type':'address'}],'name':'Upgraded','type':'event'},{'constant':true,'inputs':[],'name':'admin','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':true,'inputs':[],'name':'implementation','outputs':[{'name':'','type':'address'}],'payable':false,'stateMutability':'view','type':'function'},{'constant':false,'inputs':[{'name':'newAdmin','type':'address'}],'name':'changeAdmin','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'newImplementation','type':'address'}],'name':'upgradeTo','outputs':[],'payable':false,'stateMutability':'nonpayable','type':'function'},{'constant':false,'inputs':[{'name':'newImplementation','type':'address'},{'name':'data','type':'bytes'}],'name':'upgradeToAndCall','outputs':[],'payable':true,'stateMutability':'payable','type':'function'}];
const AdminUpgradabilityProxyBin = '0x608060405234801561001057600080fd5b50604051602080610b27833981018060405281019080805190602001909291905050508060405180807f6f72672e7a657070656c696e6f732e70726f78792e696d706c656d656e74617481526020017f696f6e000000000000000000000000000000000000000000000000000000000081525060230190506040518091039020600019167f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3600102600019161415156100c557fe5b6100dd81610167640100000000026401000000009004565b5060405180807f6f72672e7a657070656c696e6f732e70726f78792e61646d696e000000000000815250601a0190506040518091039020600019167f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001026000191614151561014957fe5b6101613361024c640100000000026401000000009004565b5061028e565b60006101858261027b6401000000000261084b176401000000009004565b151561021f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b600080823b905060008111915050919050565b61088a8061029d6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633659cfe6146100775780634f1ef286146100ba5780635c60da1b146101085780638f2839701461015f578063f851a440146101a2575b6100756101f9565b005b34801561008357600080fd5b506100b8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610213565b005b610106600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001919091929391929390505050610268565b005b34801561011457600080fd5b5061011d610308565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561016b57600080fd5b506101a0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610360565b005b3480156101ae57600080fd5b506101b761051e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610201610576565b61021161020c610651565b610682565b565b61021b6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561025c57610257816106d9565b610265565b6102646101f9565b5b50565b6102706106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102fa576102ac836106d9565b3073ffffffffffffffffffffffffffffffffffffffff163483836040518083838082843782019150509250505060006040518083038185875af19250505015156102f557600080fd5b610303565b6103026101f9565b5b505050565b60006103126106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156103545761034d610651565b905061035d565b61035c6101f9565b5b90565b6103686106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561051257600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610466576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f43616e6e6f74206368616e6765207468652061646d696e206f6620612070726f81526020017f787920746f20746865207a65726f20616464726573730000000000000000000081525060400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61048f6106a8565b82604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a161050d81610748565b61051b565b61051a6101f9565b5b50565b60006105286106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561056a576105636106a8565b9050610573565b6105726101f9565b5b90565b61057e6106a8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610647576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f43616e6e6f742063616c6c2066616c6c6261636b2066756e6374696f6e20667281526020017f6f6d207468652070726f78792061646d696e000000000000000000000000000081525060400191505060405180910390fd5b61064f610777565b565b6000807f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c36001029050805491505090565b3660008037600080366000845af43d6000803e80600081146106a3573d6000f35b3d6000fd5b6000807f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b6001029050805491505090565b6106e281610779565b7fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60007f10d6a54a4754c8869d6886b5f5d7fbfa5b4522237ea5c60d11bc4e7a1ff9390b60010290508181555050565b565b60006107848261084b565b151561081e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603b8152602001807f43616e6e6f742073657420612070726f787920696d706c656d656e746174696f81526020017f6e20746f2061206e6f6e2d636f6e74726163742061646472657373000000000081525060400191505060405180910390fd5b7f7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c360010290508181555050565b600080823b9050600081119150509190505600a165627a7a72305820dd6414bbb207177c22db40427e1e892399cabda0589266e33eb2dd35eef8a7040029';

const SimpleContractAbi = [{"constant":true,"inputs":[],"name":"initialized","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_x","type":"uint256"}],"name":"setX","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"readX","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}];

// async function main() {
//   const proxyFactoryContract = new web3.eth.Contract(UpgradabilityProxyFactoryAbi, '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa');
//   console.log("----1---");
//   console.log(proxyFactoryContract.options.address);
//   console.log("----2---");
//   const tmp = await proxyFactoryContract.methods.createProxy('0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0', '0x490fc7f8453e56f84abfc8fd17d74ad3fb6e819f').send({from: '0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0'});
//   console.log(tmp);
//   // const proxyContract = new web3.eth.Contract(AdminUpgradabilityProxyAbi, proxyAddr);
//   console.log("----3---");
//   console.log(proxyContract.options.address);
//   proxyContract.methods.admin().send({gasLimit: 5000000, from: web3.eth.coinbase}).then((a) => console.log(a));
// }
// main();
// const proxyFactoryContract = new web3.eth.Contract(UpgradabilityProxyFactoryAbi, '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa');
// console.log("----1---");
// console.log(proxyFactoryContract.options.address);
// console.log("----2---");
// proxyFactoryContract.methods.createProxy(account1, '0x490fc7f8453e56f84abfc8fd17d74ad3fb6e819f').send({gasLimit: 5000000, from: account1})
//   .then((err, proxyAddr) => {
//     const proxyContract = new web3.eth.Contract(AdminUpgradabilityProxyAbi, proxyAddr);
//     console.log(proxyContract);
//   });
// // const proxyContract = new web3.eth.Contract(AdminUpgradabilityProxyAbi, proxyAddr);
// console.log("----3---");
// console.log(proxyContract.options.address);
// proxyContract.methods.admin().send({gasLimit: 5000000, from: web3.eth.coinbase}).then((a) => console.log(a));

(async () => {
  const contract = new web3.eth.Contract(UpgradabilityProxyFactoryAbi, '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa');
  const nonce = await web3.eth.getTransactionCount(account1);
  const txObject = {
    nonce: nonce,
    gasLimit: web3.utils.toHex(600000),
    gasPrice: web3.utils.toHex(1),
    to: '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa',
    from: account1,
    data: contract.methods.createProxy('0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0', '0x5cc4d14423c101b7ca223b6f7a43bf6951f53fcc').encodeABI(),
  };
  console.log(txObject);
  var tx = new Tx(txObject);
  tx.sign(privateKey1);
  var serializedTx = tx.serialize();
  // const transaction = await web3.eth.accounts.signTransaction(txObject, privateKey2);
  const tmp = await web3.eth.sendSignedTransaction('0x' + serializedTx.toString('hex')).on('receipt', console.log);
})()

// web3.eth.getTransactionCount(account1, (err, txCount) => {
//   const txObject = {
//     nonce: txCount,
//     chaindId: 2,
//     gasLimit: web3.utils.toHex(1000000), // Raise the gas limit to a much higher amount
//     // gasPrice: web3.utils.toHex(web3.utils.toWei('10', 'gwei')),
//     to: '0x7A5727E94bbb559e0eAfC399354Dd30dBD51d2aa',
//     data: contract.methods.createProxy('0x2429f4aa5cf9d23fea0961780ffb4ff8916a26a0', '0xf6ec2fcac0f3da221270d4ec105f842560499935').encodeABI(),
//   };

//   const tx = new Tx(txObject);
//   tx.sign(privateKey1);

//   const serializedTx = tx.serialize();
//   const raw = '0x' + serializedTx.toString('hex');
// //   web3.eth.sendSignedTransaction(raw).then((err, response) => {
// //     console.log(response);
// //     console.log(err);
// //  });
//   web3.eth.sendSignedTransaction(raw).on('receipt', console.log);
// });

