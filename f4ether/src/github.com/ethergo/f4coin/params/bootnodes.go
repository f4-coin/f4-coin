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

var MainnetBootnodes = []string{
	"enode://880e0f75190bc5f43d5d68a43d781db5680d507a992dcd87949417b2abb4905d9b8d03c3209455e04d115d3d046d360cb5157d755b3de8328ebe20be1a1d4dac@52.78.172.141:21212",
}

var MainnetMasternodes = []string{
	"enode://880e0f75190bc5f43d5d68a43d781db5680d507a992dcd87949417b2abb4905d9b8d03c3209455e04d115d3d046d360cb5157d755b3de8328ebe20be1a1d4dac@52.78.172.141:21212",
}

var MainnetInitIds = []string{
	"880e0f75190bc5f4",
//	"7b659c0c41a5697c",
}

var TestnetBootnodes = []string{
	"enode://d63763d52995f3546bdd9c18f3c7374a1eb1d7826bb288b32c76a9759e61c19bb711cb9235f465c5d64703f9fb73360d798295a8139352aa1c3749964b82a510@127.0.0.1:21212",
}

var TestnetMasternodes = []string{
	"enode://d63763d52995f3546bdd9c18f3c7374a1eb1d7826bb288b32c76a9759e61c19bb711cb9235f465c5d64703f9fb73360d798295a8139352aa1c3749964b82a510", // nodekey: a9b50794ab7a9987aa416c455c13aa6cc8c0448c501a3ce8e4840efe47cb5c29
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
