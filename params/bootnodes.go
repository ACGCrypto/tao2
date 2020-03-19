// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Tao Bootnodes Mainnet
	"enode://730004961738a338f882a1e1c8ffec5abbccdaab3d91a59842e820e6e02110ec5298c1585a8a4e1163f5a3c0058720aa1d4bbc73ebf9e10bde4c820faf324248@boot1.tao.network:30301",
	"enode://3b49d79ab32e7c16913211bf7645c501cd61f7f0153acffcd5a0071e70aaba8194e9149b4cd0a896e0dfa08179a08b364fe5863d42cfa24c25c6c749968b42af@boot2.tao.network:30301",
	"enode://196c3625b2493bbd22d8dca6022d9c8e5fbedd73f56d83c001e6071f72721c12a460f2a6d60d8bb5e1410ba5f32e75e9419f5f69f9b4b2f86bdc5c47c916523d@boot3.tao.network:30301",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	// Tao Bootnodes Testnet
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
