
git config --global credential.helper wincred
git ��ס����
//geth --datadir "ethdevdata" --dev --rpc --rpcaddr "0.0.0.0" console 2>>file_to_log_output
geth --identity "phoenix"  --rpc  --rpccorsdomain "*" --datadir "/root/ethdevdata"  --dev --rpc --rpcaddr "0.0.0.0" --rpcapi "db,eth,net,web3" --networkid 98888 console 2>>file_to_log_output
//������һ���ն�����Ϊconsole
//geth --dev console 2>>file_to_log_output

//�鿴�˺�
eth.accounts

����һ�����˻�
personal.newAccount('123456')

user1=eth.accounts[0]
user2=eth.accounts[1]
eth.getBalance(user1)
eth.getBalance(user2)

miner.start()
miner.stop()

user1��user2ת����̫��
eth.sendTransaction({from: user1,to: user2,value: web3.toWei(3,"ether")})

personal.unlockAccount(user1,'123456')
eth.sendTransaction({from: user1,to: user2,value: web3.toWei(3,"ether")})

���ܺ�Լ
contract test {  
    function multiply(uint a) returns(uint d) {  
        return a * 7;   
    }   
}  
source = "contract test { function multiply(uint a) returns(uint d) { return a * 7; } }"
contract = eth.compile.solidity(source).test
������䱨��
�����·�����

 /////////////////////////////////////////////////////////////
 solc --bin Test.sol
solc --abi Test.sol
abi = JSON.parse('[{"constant":false,"inputs":[{"name":"a","type":"uint256"}],"name":"multiply","outputs":[{"name":"d","type":"uint256"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]')

bytecode = '0x60606040523415600e57600080fd5b60a98061001c6000396000f30060606040526000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063c6888fa114603c57600080fd5b3415604657600080fd5b605a60048080359060200190919050506070565b6040518082815260200191505060405180910390f35b60006007820290509190505600a165627a7a723058205f6f417a1ebd4d6c82f954c8b45fc6a9426b320aff6b40b7c8c6ff1894eb23390029'

 ��ȡabi�ļ���bin�ļ�

 >contract = eth.contract(abi)
 >deployed = contract.new({from: eth.accounts[0], data: bytecode, gas: 1000000})
deployed���ǲ����ʵ���ˣ�contract.new���ǲ��𣬽��յĲ�����from��ʾ���ĸ��˺�ȥ����������˺ŵĵ�ַ�ͻ���contractg���캯������Ǹ�msg.sender��Ҳ���������Լ�Ժ��owner��
personal.unlockAccount(user1,'123456')
deployed.multiply.call(3)
����
  >c = eth.contract(abi)
  >instance = c.at("0x69a4c171e4ff1e60985b15cdbaa398b94081c699")
���ȸ���abi�ļ������Լ�ӿڣ�Ȼ��ָ���Լ����ĵ�ַ
>instance.multiply.call(3)
>instance.multiply(3, {from:eth.accounts[0]})
personal.unlockAccount(user1,'123456')

testsol �����ַ0x4a8822c83e2dccec89e8c90c3638145976aaf9df

����ʵ����ĵ�ַ��undefined����Ϊ���ײŷ��ͳ�ȥ��û�п��ڿ�û��д�������������Ի���֪�������ڿ�����txpool.status�鿴�����ڿ��console����deployed�ɿ���address�����Ǵ˺�Լ��ַ


�����ڿ󣬰ѽ���д����������miner.start()��Ȼ��ر��ڿ�miner.stop()


�ڿ�֮�󣬽��׾�д���������ˡ���ʱ�ٲ鿴contract���ܿ����е�ַ��

/////////////�ڿ�ĳ����ʵ��ַ/////////////////////////
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
testsol �����ַ
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