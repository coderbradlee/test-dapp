infura
truffle
testrpc

git config --global credential.helper wincred
git 记住密码
//geth --datadir "ethdevdata" --dev --rpc --rpcaddr "0.0.0.0" console 2>>file_to_log_output
geth --identity "phoenix"  --rpc  --rpccorsdomain "*" --datadir "/root/ethdevdata"  --dev --rpc --rpcaddr "0.0.0.0" --rpcapi "db,eth,net,web3" --networkid 98888 console 2>>file_to_log_output
//启动另一个终端来作为console
//geth --dev console 2>>file_to_log_output

//查看账号
eth.accounts

创建一个新账户
personal.newAccount('123456')

user1=eth.accounts[0]
user2=eth.accounts[1]
eth.getBalance(user1)
eth.getBalance(user2)

miner.start()
miner.stop()

user1向user2转移以太币
eth.sendTransaction({from: user1,to: user2,value: web3.toWei(3,"ether")})

personal.unlockAccount(user1,'123456')
eth.sendTransaction({from: user1,to: user2,value: web3.toWei(3,"ether")})

智能合约
contract test {  
    function multiply(uint a) returns(uint d) {  
        return a * 7;   
    }   
}  
source = "contract test { function multiply(uint a) returns(uint d) { return a * 7; } }"
contract = eth.compile.solidity(source).test
以上语句报错
用以下方法：

 /////////////////////////////////////////////////////////////
 solc --bin Test.sol
solc --abi Test.sol
abi = JSON.parse('[{"constant":false,"inputs":[{"name":"a","type":"uint256"}],"name":"multiply","outputs":[{"name":"d","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]')

bytecode = '0x60606040523415600e57600080fd5b60a98061001c6000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063c6888fa114603c57600080fd5b3415604657600080fd5b605a60048080359060200190919050506070565b6040518082815260200191505060405180910390f35b60006007820290509190505600a165627a7a723058205f6f417a1ebd4d6c82f954c8b45fc6a9426b320aff6b40b7c8c6ff1894eb23390029'

 读取abi文件和bin文件

 >contract = eth.contract(abi)
 >deployed = contract.new({from: eth.accounts[0], data: bytecode, gas: 1000000})
deployed就是部署的实例了，contract.new就是部署，接收的参数里from表示用哪个账号去创建，这个账号的地址就会是contractg构造函数里的那个msg.sender，也就是这个合约以后的owner了
personal.unlockAccount(user1,'123456')
deployed.multiply.call(3)
或者
  >c = eth.contract(abi)
  >instance = c.at("0x69a4c171e4ff1e60985b15cdbaa398b94081c699")
首先根据abi文件构造合约接口，然后指向合约部署的地址
>instance.multiply.call(3)
>instance.multiply(3, {from:eth.accounts[0]})
personal.unlockAccount(user1,'123456')

testsol 部署地址0x4a8822c83e2dccec89e8c90c3638145976aaf9df

部署实例后的地址是undefined，因为交易才发送出去，没有矿工挖矿，没有写进区块链，所以还不知道。现在可以用txpool.status查看到，挖矿后console输入deployed可看到address，就是此合约地址


启动挖矿，把交易写进区块链。miner.start()，然后关闭挖矿miner.stop()


挖矿之后，交易就写进区块链了。此时再查看contract就能看见有地址了

/////////////挖矿到某个真实地址/////////////////////////
geth --etherbase 1 --mine  2>> geth.log // 1 is index: second account by creation order OR
geth --etherbase '0xB16FC01EE749d4Db81F712E6747120323eebA8f7' --mine 2>> geth.log

/root/go-ethereum/build/bin/geth --etherbase '0xB16FC01EE749d4Db81F712E6747120323eebA8f7' --mine --minerthreads=4 2>> geth.log


//abigen --abi testsol.abi --pkg main --type Token --out testsol.go
//go:generate abigen --sol testsol.sol --pkg main --out testsol.go
abigen --sol testsol.sol --pkg main --out testsol.go

curl -X POST --data '{"jsonrpc":"2.0","method":"net_version","params":[],"id":67}' http://localhost:8545

curl -X POST --data '{"jsonrpc":"2.0","method":"web3_clientVersion","params":[],"id":67}' http://localhost:8545

curl -X POST --data '{"jsonrpc":"2.0","method":"net_listening","params":[],"id":67}' http://localhost:8545

curl -X POST --data '{"jsonrpc":"2.0","method":"eth_accounts","params":[],"id":73}' http://localhost:8545

curl -X POST --data '{"jsonrpc":"2.0", "method": "eth_getStorageAt", "params": ["0xb13edd24707dfc039614c1c8927f2baafb36474c", "0x0", "latest"], "id": 1}' localhost:8545
testsol 部署地址
address: "0xb13edd24707dfc039614c1c8927f2baafb36474c",
  transactionHash: "0xac99518b20395665dfd9fae531f173e3de1a1466faf434d9cc6650a3228b8ec6"

  curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByHash","params":["0xf60c195debeaecdfc73f62b68a5ba52970378d4c01748db1277e4298615e897e", true],"id":1}' localhost:8545

  web3.eth.blockNumber
  web3.eth.getBlock(3).hash
  web3.eth.getBlock(BLOCK_NUMBER)

  curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x3", true],"id":1}' localhost:8545

  curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params":["0xf60c195debeaecdfc73f62b68a5ba52970378d4c01748db1277e4298615e897e"],"id":1}' localhost:8545

  curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getTransactionByBlockNumberAndIndex","params":["0x5", "0x0"],"id":1}' localhost:8545

  curl -X POST --data '{"jsonrpc":"2.0","method":"eth_getCompilers","params":[],"id":1}' localhost:8545

  curl -X POST --data '{"jsonrpc":"2.0","method":"eth_compileSolidity","params":["contract test { function multiply(uint a) returns(uint d) {   return a * 7;   } }"],"id":1}' localhost:8545


  //testrpc 测试智能合约,“truffle console”下执行
  Ballot.deployed()
  写js测试


  特别变量和函数
一些特别变量和函数已收入在全局命名空间里。
Block and Transaction Properties
块和资产交易
• block.coinbase (address): current block miner's address
当前块的矿工地址
• block.difficulty (uint): current block difficulty
当前块难度定义
• block.gaslimit (uint): current block gaslimit
当前块的瓦斯限额
• block.number (uint): current block number
当前块数字标识
• block.blockhash (function(uint) returns (hash)): hash of the given block
获取给定块的散列Hash
• block.timestamp (uint): current block timestamp
当前块的时间戳
• msg.data (bytes): complete calldata
完整的调用数据
• msg.gas (uint): remaining gas
剩余的瓦斯值
• msg.sender (address): sender of the message (current call)
消息发送函数（当前调用）
• msg.value (uint): number of wei sent with the message
随消息发送的交易费数值，以wei（微）计量。
• tx.gasprice (uint): gas price of the transaction
完成交易的瓦斯价格
• tx.origin (address): sender of the transaction (full call chain)
交易发送函数（全链调用？）

秘钥函数
• sha3(...) returns (hash): compute the SHA3 hash of the (tightly packed) arguments
计算SHA3散列,紧凑排列
• sha256(...) returns (hash): compute the SHA256 hash of the (tightly packed) arguments
计算SHA256散列,紧凑排列
• ripemd160(...) returns (hash160): compute RIPEMD of 256 the (tightly packed) arguments
计算RIPEMD160散列,紧凑排列
• ecrecover(hash, hash8, hash, hash) returns (address): recover public key from elliptic curve signature
从椭圆曲线签名中恢复公钥

上述“紧凑”意味着自变量是无间隔连接着的，例：
sha3("ab", "c") == sha3("abc") == sha3(0x616263) == sha3(6382179) = sha3(97, 98, 99).
如果需要间隔，会用到特定类型转换。

合约相关
• this (current contract's type): the current contract, explicitly convertible to address
指定当前合约，明确可转换的地址
• suicide(address): suicide the current contract, sending its funds to the given address
消除当前合约，发送其资金到给定地址；

而且，当前合约的所有函数都是直接可调用的，包括当前函数。

地址相关函数
查询某一地址的余额可用属性balance ，或以send发送以太币（以Wei为单位）到一地址。
address x = 0x123;
if (x.balance < 10 && address(this).balance >= 10) x.send(10);

而且，与合约的接口不能依附于ABI（如传统的NameReg合约），函数调用会获得一个任意类型变量的任意数值。这些自变量是ABI的序列化 （如：也填充到32bytes）。一个异常情况是首自变量仅是4bytes编码。这种情况下，并不是以填充(位数)允许签名函数在此应用。
address nameReg = 0x72ba7d8e73fe8eb666ea66babc8116a41bfb10e2;
nameReg.call("register", "MyName");
nameReg.call(string4(string32(sha3("fun(uint256)"))), a);

注意合约继承了所有地址成员，有可能使用this.balance来查询当前合约的余额。

表达式的评估命令
表达式评估的指令并未特别指定（更正式说，在表达式树中子节点的评估指令是未指定的，但这些当然是在节点自身之前就评估了的）。能确保的是顺序语句执行和布尔型短路表达式是可以的。
