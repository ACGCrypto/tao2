# Tao

[![Build Status](https://travis-ci.org/taoblockchain/tao2.svg?branch=master)](https://travis-ci.org/taoblockchain/tao2)
[![Join the chat at https://gitter.im/taoblockchain-tao2/community](https://badges.gitter.im/taoblockchain-tao2/community.svg)](https://gitter.im/taoblockchain-tao2/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

## About Tao 2.0

Tao is a cryptocurrency designed to support artists in the music industry. It began in 2016 as a Bitcoin-derived blockchain powered by Proof of Stake. After nearly four years of flawless operation, the Tao blockchain is due  for some major upgrades so we can move into the next phase of our project. To achieve our project goals we have chosen to utilise an Ethereum Virtual Machine based protocol which will enable Solidity development and the potential migration of music and entertainment projects from Ethereum to Tao. 

We have chosen the worthy open source Tomochain as the best Ethereum-derived codebase to start with, but will incorporate elements from accross the crypto space to best serve the interests of the unique needs of the music economy. 

Tao will be driven by a maximum of 150 Masternodes with a Delegated Proof of Stake voting consensus supporting near-zero fee, and 5-second transaction confirmation times. Masternodes will require a minimum of 100,000 TAO to operate as well as sufficient votes to qualify once there are more than 150 masternode candidates. A board of governors will be elected from the active Masternodes to guide future development.

Tao supports standard EVM-compatible smart-contracts, protocols, as well as atomic cross-chain token transfers.

This readme file is an early draft and is subject to revision.

More details can be found by reading our [technical white paper](https://tao.network/docs/technical-whitepaper---1.0.pdf)

Read more about us on:

- our website: http://tao.network
- our blogs and announcements: https://medium.com/BryceWeiner/
- our documentation portal: https://docs.tao.network

## Building the source

Tao provides a client binary called `tao` for both running a masternode and running a full-node.
Building `tao` requires both a Go (1.7+) and C compiler; install both of these.

Once the dependencies are installed, just run the below commands:

```bash
$ git clone https://github.com/taoblockchain/tao2 tao
$ cd tao
$ make tao
```

Alternatively, you could quickly download our pre-complied binary from our [github release page](https://github.com/taoblockchain/tao2/releases)

## Running `tao`

### Running a tao masternode

Please refer to the [official documentation](https://docs.tao.network/get-started/run-node/) on how to run a node if your goal is to run a masternode.
The recommanded ways of running a node and applying to become a masternode are explained in detail there.

### Attaching to the Tao test network

We published our test network 2.0 with full implementation of DPOS consensus at https://stats.testnet.tao.network.
If you'd like to experiment with smart contract creation and DApps, you might be interested to give these a try on our Testnet.

In order to connect to one of the masternodes on the Testnet, just run the command below:

```bash
$ tao attach https://testnet.tao.network
```

This will open the JavaScript console and let you query the blockchain directly via RPC.

### Running `tao` locally

#### Download genesis block
$GENESIS_PATH : location of genesis file you would like to put
```bash
export GENESIS_PATH=path/to/genesis.json
```
- Testnet
```bash
curl -L https://raw.githubusercontent.com/taoblockchain/tao2/master/genesis/testnet.json -o $GENESIS_PATH
```

- Mainnet
```bash
curl -L https://raw.githubusercontent.com/taoblockchain/tao2/master/genesis/mainnet.json -o $GENESIS_PATH
```

#### Create datadir
- create a folder to store tao data on your machine

```bash
export DATA_DIR=/path/to/your/data/folder 
mkdir -p $DATA_DIR/tao
```
#### Initialize the chain from genesis

```bash
tao init $GENESIS_PATH --datadir $DATA_DIR
```

#### Initialize / Import accounts for the nodes's keystore
If you already had an existing account, import it. Otherwise, please initialize new accounts 

```bash
export KEYSTORE_DIR=path/to/keystore
```

##### Initialize new accounts
```bash
tao account new \
  --password [YOUR_PASSWORD_FILE_TO_LOCK_YOUR_ACCOUNT] \
  --keystore $KEYSTORE_DIR
```
    
##### Import accounts
```bash
tao  account import [PRIVATE_KEY_FILE_OF_YOUR_ACCOUNT] \
     --keystore $KEYSTORE_DIR \
     --password [YOUR_PASSWORD_FILE_TO_LOCK_YOUR_ACCOUNT]
```

##### List all available accounts in keystore folder

```bash
tao account list --datadir ./  --keystore $KEYSTORE_DIR
```

#### Start a node
##### Environment variables
   - $IDENTITY: the name of your node
   - $PASSWORD: the password file to unlock your account
   - $YOUR_COINBASE_ADDRESS: address of your account which generated in the previous step
   - $NETWORK_ID: the networkId. Mainnet: 558. Testnet: 688
   - $BOOTNODES: The comma separated list of bootnodes. Find them [here](https://docs.tao.network/general/networks/)
   - $WS_SECRET: The password to send data to the stats website. Find them [here](https://docs.tao.network/general/networks/)
   - $NETSTATS_HOST: The stats website to report to, regarding to your environment. Find them [here](https://docs.tao.network/general/networks/)
   - $NETSTATS_PORT: The port used by the stats website (usually 443)
    
##### Let's start a node
```bash
tao  --syncmode "full" \    
    --datadir $DATA_DIR --networkid $NETWORK_ID --port 20202 \   
    --keystore $KEYSTORE_DIR --password $PASSWORD \    
    --rpc --rpccorsdomain "*" --rpcaddr 0.0.0.0 --rpcport 8545 --rpcvhosts "*" \   
    --rpcapi "db,eth,net,web3,personal,debug" \    
    --gcmode "archive" \   
    --ws --wsaddr 0.0.0.0 --wsport 8546 --wsorigins "*" --unlock "$YOUR_COINBASE_ADDRESS" \   
    --identity $IDENTITY \  
    --mine --gasprice 2500 \  
    --bootnodes $BOOTNODES \   
    --ethstats $IDENTITY:$WS_SECRET@$NETSTATS_HOST:$NETSTATS_PORT 
    console
```


##### Some explanations on the flags   
```
--verbosity: log level from 1 to 5. Here we're using 4 for debug messages
--datadir: path to your data directory created above.
--keystore: path to your account's keystore created above.
--identity: your full-node's name.
--password: your account's password.
--networkid: our network ID.
--tao-testnet: required when the networkid is testnet(89).
--port: your full-node's listening port (default to 20202)
--rpc, --rpccorsdomain, --rpcaddr, --rpcport, --rpcvhosts: your full-node will accept RPC requests at 8545 TCP.
--ws, --wsaddr, --wsport, --wsorigins: your full-node will accept Websocket requests at 8546 TCP.
--mine: your full-node wants to register to be a candidate for masternode selection.
--gasprice: Minimal gas price to accept for mining a transaction.
--targetgaslimit: Target gas limit sets the artificial target gas floor for the blocks to mine (default: 4712388)
--bootnode: bootnode information to help to discover other nodes in the network
--gcmode: blockchain garbage collection mode ("full", "archive")
--synmode: blockchain sync mode ("fast", "full", or "light". More detail: https://github.com/taoblockchain/tao2/blob/master/eth/downloader/modes.go#L24)           
--ethstats: send data to stats website
```
To see all flags usage
   
```bash
tao --help
```

#### See your node on stats page
   - Testnet: https://stats.testnet.tao.network
   - Mainnet: http://stats.tao.network


## Contributing and technical discussion

Thank you for considering to try out our network and/or help out with the source code.
We would love to get your help; feel free to lend a hand.
Even the smallest bit of code, bug reporting, or just discussing ideas are highly appreciated.

If you would like to contribute to the tao source code, please refer to our Developer Guide for details on configuring development environment, managing dependencies, compiling, testing and submitting your code changes to our repo.

Please also make sure your contributions adhere to the base coding guidelines:

- Code must adhere to official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e uses [gofmt](https://golang.org/cmd/gofmt/)).
- Code comments must adhere to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
- Pull requests need to be based on and opened against the `master` branch.
- Any code you are trying to contribute must be well-explained as an issue on our [github issue page](https://github.com/taoblockchain/tao2/issues)
- Commit messages should be short but clear enough and should refer to the corresponding pre-logged issue mentioned above.


