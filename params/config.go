// Copyright 2016 The go-ethereum Authors
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

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params/forks"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3")

	CoreGenesisHash    = common.HexToHash("0xf7fc87f11e61508a5828cd1508060ed1714c8d32a92744ae10acb43c953357ad")
	BuffaloGenesisHash = common.HexToHash("0xd90508c51efd64e75363cdf51114d9f2a90a79e6cd0f78f3c3038b47695c034a")
	PigeonGenesisHash  = common.HexToHash("0xdfe68477f9fbc0d3e362940fcf87fa54add5bc97c4afd7d3dee31919df40212c")
)

func newUint64(val uint64) *uint64 { return &val }

var (
	MainnetTerminalTotalDifficulty, _ = new(big.Int).SetString("58_750_000_000_000_000_000_000", 0)

	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:                       big.NewInt(1),
		HomesteadBlock:                big.NewInt(1_150_000),
		DAOForkBlock:                  big.NewInt(1_920_000),
		DAOForkSupport:                true,
		EIP150Block:                   big.NewInt(2_463_000),
		EIP155Block:                   big.NewInt(2_675_000),
		EIP158Block:                   big.NewInt(2_675_000),
		ByzantiumBlock:                big.NewInt(4_370_000),
		ConstantinopleBlock:           big.NewInt(7_280_000),
		PetersburgBlock:               big.NewInt(7_280_000),
		IstanbulBlock:                 big.NewInt(9_069_000),
		MuirGlacierBlock:              big.NewInt(9_200_000),
		BerlinBlock:                   big.NewInt(12_244_000),
		LondonBlock:                   big.NewInt(12_965_000),
		ArrowGlacierBlock:             big.NewInt(13_773_000),
		GrayGlacierBlock:              big.NewInt(15_050_000),
		TerminalTotalDifficulty:       MainnetTerminalTotalDifficulty, // 58_750_000_000_000_000_000_000
		TerminalTotalDifficultyPassed: true,
		ShanghaiTime:                  newUint64(1681338455),
		CancunTime:                    newUint64(1710338135),
		Ethash:                        new(EthashConfig),
	}
	// SepoliaChainConfig contains the chain parameters to run a node on the Sepolia test network.
	SepoliaChainConfig = &ChainConfig{
		ChainID:                       big.NewInt(11155111),
		HomesteadBlock:                big.NewInt(0),
		DAOForkBlock:                  nil,
		DAOForkSupport:                true,
		EIP150Block:                   big.NewInt(0),
		EIP155Block:                   big.NewInt(0),
		EIP158Block:                   big.NewInt(0),
		ByzantiumBlock:                big.NewInt(0),
		ConstantinopleBlock:           big.NewInt(0),
		PetersburgBlock:               big.NewInt(0),
		IstanbulBlock:                 big.NewInt(0),
		MuirGlacierBlock:              big.NewInt(0),
		BerlinBlock:                   big.NewInt(0),
		LondonBlock:                   big.NewInt(0),
		ArrowGlacierBlock:             nil,
		GrayGlacierBlock:              nil,
		TerminalTotalDifficulty:       big.NewInt(17_000_000_000_000_000),
		TerminalTotalDifficultyPassed: true,
		MergeNetsplitBlock:            big.NewInt(1735371),
		ShanghaiTime:                  newUint64(1677557088),
		CancunTime:                    newUint64(1706655072),
		Ethash:                        new(EthashConfig),
	}

	// just for prysm compile pass
	// GoerliChainConfig contains the chain parameters to run a node on the Görli test network.
	GoerliChainConfig = &ChainConfig{
		ChainID:                       big.NewInt(5),
		HomesteadBlock:                big.NewInt(0),
		DAOForkBlock:                  nil,
		DAOForkSupport:                true,
		EIP150Block:                   big.NewInt(0),
		EIP155Block:                   big.NewInt(0),
		EIP158Block:                   big.NewInt(0),
		ByzantiumBlock:                big.NewInt(0),
		ConstantinopleBlock:           big.NewInt(0),
		PetersburgBlock:               big.NewInt(0),
		IstanbulBlock:                 big.NewInt(1_561_651),
		MuirGlacierBlock:              nil,
		BerlinBlock:                   big.NewInt(4_460_644),
		LondonBlock:                   big.NewInt(5_062_605),
		ArrowGlacierBlock:             nil,
		TerminalTotalDifficulty:       big.NewInt(10_790_000),
		TerminalTotalDifficultyPassed: true,
		ShanghaiTime:                  newUint64(1678832736),
		CancunTime:                    newUint64(1705473120),
		Clique: &CliqueConfig{
			Period: 15,
			Epoch:  30000,
		},
	}

	CoreChainConfig = &ChainConfig{
		ChainID:             big.NewInt(1116),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		HashPowerBlock:      big.NewInt(0),
		ZeusBlock:           big.NewInt(8_020_000),
		HeraBlock:           big.NewInt(12_195_500),
		PoseidonBlock:       big.NewInt(13_232_049),
		BerlinBlock:         big.NewInt(19_537_200),
		LondonBlock:         big.NewInt(19_537_200),
		HertzBlock:          big.NewInt(19_537_200),
		ShanghaiTime:        newUint64(1731999600), // 2024-11-19 7:00:00 AM UTC
		KeplerTime:          newUint64(1731999600),
		DemeterTime:         newUint64(1731999600),
		AthenaTime:          newUint64(1738544400), // 2025-02-03 1:00:00 AM UTC
		TheseusTime:         newUint64(1750838400), // 2025-06-25 8:00:00 AM UTC
		CancunTime:          newUint64(1750838400), // 2025-06-25 8:00:00 AM UTC
		TheseusFixTime:      newUint64(1753948800), // 2025-07-31 08:00:00 AM UTC
		Satoshi: &SatoshiConfig{
			Period: 3,
			Epoch:  200,
			Round:  86400,
		},
	}

	BuffaloChainConfig = &ChainConfig{
		ChainID:             big.NewInt(1115),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		HashPowerBlock:      big.NewInt(4_545_256),
		ZeusBlock:           big.NewInt(12_666_000),
		HeraBlock:           big.NewInt(16_472_288),
		PoseidonBlock:       big.NewInt(18_253_800),
		BerlinBlock:         big.NewInt(24_127_766),
		LondonBlock:         big.NewInt(24_127_766),
		HertzBlock:          big.NewInt(24_127_766),
		ShanghaiTime:        newUint64(1729132200), // 2024-10-17 2:30:00 AM UTC
		KeplerTime:          newUint64(1729132200),
		DemeterTime:         newUint64(1729132200),
		Satoshi: &SatoshiConfig{
			Period: 3,
			Epoch:  200,
			Round:  86400,
		},
	}

	PigeonChainConfig = &ChainConfig{
		ChainID:             big.NewInt(1114),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		HashPowerBlock:      big.NewInt(0),
		ZeusBlock:           big.NewInt(0),
		HeraBlock:           big.NewInt(0),
		PoseidonBlock:       big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		HertzBlock:          big.NewInt(0),
		ShanghaiTime:        newUint64(0),
		KeplerTime:          newUint64(0),
		DemeterTime:         newUint64(0),
		AthenaTime:          newUint64(1737079200),
		TheseusTime:         newUint64(1748592000), // 2025-05-30 08:00:00 AM UTC
		CancunTime:          newUint64(1748592000), // 2025-05-30 08:00:00 AM UTC
		TheseusFixTime:      newUint64(1752652800), // 2025-07-16 08:00:00 AM UTC
		Satoshi: &SatoshiConfig{
			Period: 3,
			Epoch:  200,
			Round:  86400,
		},
	}

	SatoshiTestChainConfig = &ChainConfig{
		ChainID:             big.NewInt(2),
		HomesteadBlock:      big.NewInt(0),
		EIP150Block:         big.NewInt(0),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(0),
		PetersburgBlock:     big.NewInt(0),
		IstanbulBlock:       big.NewInt(0),
		MuirGlacierBlock:    big.NewInt(0),
		HashPowerBlock:      big.NewInt(0),
		ZeusBlock:           big.NewInt(0),
		HeraBlock:           big.NewInt(0),
		PoseidonBlock:       big.NewInt(0),
		BerlinBlock:         big.NewInt(0),
		LondonBlock:         big.NewInt(0),
		HertzBlock:          big.NewInt(0),
		ShanghaiTime:        newUint64(0),
		KeplerTime:          newUint64(0),
		DemeterTime:         newUint64(0),
		AthenaTime:          newUint64(0),
		TheseusTime:         newUint64(0),
		CancunTime:          newUint64(0),
		TheseusFixTime:      newUint64(0),
		Satoshi: &SatoshiConfig{
			Period: 3,
			Epoch:  200,
			Round:  86400,
		},
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	AllEthashProtocolChanges = &ChainConfig{
		ChainID:                       big.NewInt(1337),
		HomesteadBlock:                big.NewInt(0),
		DAOForkBlock:                  nil,
		DAOForkSupport:                false,
		EIP150Block:                   big.NewInt(0),
		EIP155Block:                   big.NewInt(0),
		EIP158Block:                   big.NewInt(0),
		ByzantiumBlock:                big.NewInt(0),
		ConstantinopleBlock:           big.NewInt(0),
		PetersburgBlock:               big.NewInt(0),
		IstanbulBlock:                 big.NewInt(0),
		MuirGlacierBlock:              big.NewInt(0),
		BerlinBlock:                   big.NewInt(0),
		LondonBlock:                   big.NewInt(0),
		ArrowGlacierBlock:             big.NewInt(0),
		GrayGlacierBlock:              big.NewInt(0),
		MergeNetsplitBlock:            nil,
		ShanghaiTime:                  nil,
		CancunTime:                    nil,
		PragueTime:                    nil,
		VerkleTime:                    nil,
		TerminalTotalDifficulty:       nil,
		TerminalTotalDifficultyPassed: true,
		Ethash:                        new(EthashConfig),
		Clique:                        nil,
	}

	AllDevChainProtocolChanges = &ChainConfig{
		ChainID:                       big.NewInt(1337),
		HomesteadBlock:                big.NewInt(0),
		EIP150Block:                   big.NewInt(0),
		EIP155Block:                   big.NewInt(0),
		EIP158Block:                   big.NewInt(0),
		ByzantiumBlock:                big.NewInt(0),
		ConstantinopleBlock:           big.NewInt(0),
		PetersburgBlock:               big.NewInt(0),
		IstanbulBlock:                 big.NewInt(0),
		MuirGlacierBlock:              big.NewInt(0),
		BerlinBlock:                   big.NewInt(0),
		LondonBlock:                   big.NewInt(0),
		ArrowGlacierBlock:             big.NewInt(0),
		GrayGlacierBlock:              big.NewInt(0),
		ShanghaiTime:                  newUint64(0),
		CancunTime:                    newUint64(0),
		TerminalTotalDifficulty:       big.NewInt(0),
		TerminalTotalDifficultyPassed: true,
	}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	AllCliqueProtocolChanges = &ChainConfig{
		ChainID:                       big.NewInt(1337),
		HomesteadBlock:                big.NewInt(0),
		DAOForkBlock:                  nil,
		DAOForkSupport:                false,
		EIP150Block:                   big.NewInt(0),
		EIP155Block:                   big.NewInt(0),
		EIP158Block:                   big.NewInt(0),
		ByzantiumBlock:                big.NewInt(0),
		ConstantinopleBlock:           big.NewInt(0),
		PetersburgBlock:               big.NewInt(0),
		IstanbulBlock:                 big.NewInt(0),
		MuirGlacierBlock:              big.NewInt(0),
		BerlinBlock:                   big.NewInt(0),
		LondonBlock:                   big.NewInt(0),
		ArrowGlacierBlock:             nil,
		GrayGlacierBlock:              nil,
		MergeNetsplitBlock:            nil,
		ShanghaiTime:                  nil,
		CancunTime:                    nil,
		PragueTime:                    nil,
		VerkleTime:                    nil,
		TerminalTotalDifficulty:       nil,
		TerminalTotalDifficultyPassed: false,
		Ethash:                        nil,
		Clique:                        &CliqueConfig{Period: 0, Epoch: 30000},
	}

	// TestChainConfig contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers for testing purposes.
	TestChainConfig = &ChainConfig{
		ChainID:                       big.NewInt(1),
		HomesteadBlock:                big.NewInt(0),
		DAOForkBlock:                  nil,
		DAOForkSupport:                false,
		EIP150Block:                   big.NewInt(0),
		EIP155Block:                   big.NewInt(0),
		EIP158Block:                   big.NewInt(0),
		ByzantiumBlock:                big.NewInt(0),
		ConstantinopleBlock:           big.NewInt(0),
		PetersburgBlock:               big.NewInt(0),
		IstanbulBlock:                 big.NewInt(0),
		MuirGlacierBlock:              big.NewInt(0),
		BerlinBlock:                   big.NewInt(0),
		LondonBlock:                   big.NewInt(0),
		ArrowGlacierBlock:             big.NewInt(0),
		GrayGlacierBlock:              big.NewInt(0),
		MergeNetsplitBlock:            nil,
		ShanghaiTime:                  nil,
		CancunTime:                    nil,
		PragueTime:                    nil,
		VerkleTime:                    nil,
		TerminalTotalDifficulty:       nil,
		TerminalTotalDifficultyPassed: false,
		Ethash:                        new(EthashConfig),
		Clique:                        nil,
	}

	// MergedTestChainConfig contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers for testing purposes.
	MergedTestChainConfig = &ChainConfig{
		ChainID:                       big.NewInt(1),
		HomesteadBlock:                big.NewInt(0),
		DAOForkBlock:                  nil,
		DAOForkSupport:                false,
		EIP150Block:                   big.NewInt(0),
		EIP155Block:                   big.NewInt(0),
		EIP158Block:                   big.NewInt(0),
		ByzantiumBlock:                big.NewInt(0),
		ConstantinopleBlock:           big.NewInt(0),
		PetersburgBlock:               big.NewInt(0),
		IstanbulBlock:                 big.NewInt(0),
		MuirGlacierBlock:              big.NewInt(0),
		BerlinBlock:                   big.NewInt(0),
		LondonBlock:                   big.NewInt(0),
		ArrowGlacierBlock:             big.NewInt(0),
		GrayGlacierBlock:              big.NewInt(0),
		MergeNetsplitBlock:            big.NewInt(0),
		ShanghaiTime:                  newUint64(0),
		CancunTime:                    newUint64(0),
		PragueTime:                    nil,
		VerkleTime:                    nil,
		TerminalTotalDifficulty:       big.NewInt(0),
		TerminalTotalDifficultyPassed: true,
		Ethash:                        new(EthashConfig),
		Clique:                        nil,
	}

	// NonActivatedConfig defines the chain configuration without activating
	// any protocol change (EIPs).
	NonActivatedConfig = &ChainConfig{
		ChainID:                       big.NewInt(1),
		HomesteadBlock:                nil,
		DAOForkBlock:                  nil,
		DAOForkSupport:                false,
		EIP150Block:                   nil,
		EIP155Block:                   nil,
		EIP158Block:                   nil,
		ByzantiumBlock:                nil,
		ConstantinopleBlock:           nil,
		PetersburgBlock:               nil,
		IstanbulBlock:                 nil,
		MuirGlacierBlock:              nil,
		BerlinBlock:                   nil,
		LondonBlock:                   nil,
		ArrowGlacierBlock:             nil,
		GrayGlacierBlock:              nil,
		MergeNetsplitBlock:            nil,
		ShanghaiTime:                  nil,
		CancunTime:                    nil,
		PragueTime:                    nil,
		VerkleTime:                    nil,
		TerminalTotalDifficulty:       nil,
		TerminalTotalDifficultyPassed: false,
		Ethash:                        new(EthashConfig),
		Clique:                        nil,
	}
	TestRules = TestChainConfig.Rules(new(big.Int), false, 0)
)

func GetBuiltInChainConfig(ghash common.Hash) *ChainConfig {
	switch ghash {
	case MainnetGenesisHash:
		return MainnetChainConfig
	case CoreGenesisHash:
		return CoreChainConfig
	case BuffaloGenesisHash:
		return BuffaloChainConfig
	case PigeonGenesisHash:
		return PigeonChainConfig
	default:
		return nil
	}
}

// NetworkNames are user friendly names to use in the chain spec banner.
var NetworkNames = map[string]string{
	MainnetChainConfig.ChainID.String(): "mainnet",
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	IstanbulBlock       *big.Int `json:"istanbulBlock,omitempty"`       // Istanbul switch block (nil = no fork, 0 = already on istanbul)
	MuirGlacierBlock    *big.Int `json:"muirGlacierBlock,omitempty"`    // Eip-2384 (bomb delay) switch block (nil = no fork, 0 = already activated)
	BerlinBlock         *big.Int `json:"berlinBlock,omitempty"`         // Berlin switch block (nil = no fork, 0 = already on berlin)
	YoloV3Block         *big.Int `json:"yoloV3Block,omitempty"`         // YOLO v3: Gas repricings TODO @holiman add EIP references
	CatalystBlock       *big.Int `json:"catalystBlock,omitempty"`       // Catalyst switch block (nil = no fork, 0 = already on catalyst)
	LondonBlock         *big.Int `json:"londonBlock,omitempty"`         // London switch block (nil = no fork, 0 = already on london)
	ArrowGlacierBlock   *big.Int `json:"arrowGlacierBlock,omitempty"`   // Eip-4345 (bomb delay) switch block (nil = no fork, 0 = already activated)
	GrayGlacierBlock    *big.Int `json:"grayGlacierBlock,omitempty"`    // Eip-5133 (bomb delay) switch block (nil = no fork, 0 = already activated)
	MergeNetsplitBlock  *big.Int `json:"mergeNetsplitBlock,omitempty"`  // Virtual fork after The Merge to use as a network splitter

	// Fork scheduling was switched from blocks to timestamps here

	ShanghaiTime   *uint64 `json:"shanghaiTime,omitempty" `   // Shanghai switch time (nil = no fork, 0 = already on shanghai)
	KeplerTime     *uint64 `json:"keplerTime,omitempty"`      // Kepler switch time (nil = no fork, 0 = already activated)
	DemeterTime    *uint64 `json:"demeterTime,omitempty" `    // Demeter switch time (nil = no fork, 0 = already on demeter)
	AthenaTime     *uint64 `json:"athenaTime,omitempty"`      // Athena switch time (nil = no fork, 0 = already on athena)
	CancunTime     *uint64 `json:"cancunTime,omitempty" `     // Cancun switch time (nil = no fork, 0 = already on cancun)
	HaberTime      *uint64 `json:"haberTime,omitempty"`       // Haber switch time (nil = no fork, 0 = already on haber)
	PragueTime     *uint64 `json:"pragueTime,omitempty" `     // Prague switch time (nil = no fork, 0 = already on prague)
	VerkleTime     *uint64 `json:"verkleTime,omitempty" `     // Verkle switch time (nil = no fork, 0 = already on verkle)
	TheseusTime    *uint64 `json:"theseusTime,omitempty" `    // Theseus switch time (nil = no fork, 0 = already on theseus)
	TheseusFixTime *uint64 `json:"theseusFixTime,omitempty" ` // TheseusFix switch time (nil = no fork, 0 = already on theseusFix)

	// TerminalTotalDifficulty is the amount of total difficulty reached by
	// the network that triggers the consensus upgrade.
	TerminalTotalDifficulty *big.Int `json:"terminalTotalDifficulty,omitempty"`

	// TerminalTotalDifficultyPassed is a flag specifying that the network already
	// passed the terminal total difficulty. Its purpose is to disable legacy sync
	// even without having seen the TTD locally (safer long term).
	TerminalTotalDifficultyPassed bool `json:"terminalTotalDifficultyPassed,omitempty"`

	HashPowerBlock *big.Int `json:"hashPowerBlock,omitempty"`
	ZeusBlock      *big.Int `json:"zeusBlock,omitempty"`
	HeraBlock      *big.Int `json:"heraBlock,omitempty"`
	PoseidonBlock  *big.Int `json:"poseidonBlock,omitempty"`
	LubanBlock     *big.Int `json:"lubanBlock,omitempty" toml:",omitempty"` // lubanBlock switch block (nil = no fork, 0 = already activated)
	PlatoBlock     *big.Int `json:"platoBlock,omitempty" toml:",omitempty"` // platoBlock switch block (nil = no fork, 0 = already activated)
	HertzBlock     *big.Int `json:"hertzBlock,omitempty" toml:",omitempty"` // hertzBlock switch block (nil = no fork, 0 = already activated)
	// Various consensus engines
	Ethash  *EthashConfig  `json:"ethash,omitempty" toml:",omitempty"`
	Clique  *CliqueConfig  `json:"clique,omitempty" toml:",omitempty"`
	Satoshi *SatoshiConfig `json:"satoshi,omitempty" toml:",omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// SatoshiConfig is the consensus engine configs for proof-of-staked-authority based sealing.
type SatoshiConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to update validatorSet
	Round  uint64 `json:"round"`  // Number of seconds between rounds to enforce
}

// String implements the stringer interface, returning the consensus engine details.
func (b *SatoshiConfig) String() string {
	return "satoshi"
}

func (c *ChainConfig) Description() string {
	return ""
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}

	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	case c.Satoshi != nil:
		engine = c.Satoshi
	default:
		engine = "unknown"
	}

	var ShanghaiTime *big.Int
	if c.ShanghaiTime != nil {
		ShanghaiTime = big.NewInt(0).SetUint64(*c.ShanghaiTime)
	}

	var KeplerTime *big.Int
	if c.KeplerTime != nil {
		KeplerTime = big.NewInt(0).SetUint64(*c.KeplerTime)
	}

	var DemeterTime *big.Int
	if c.DemeterTime != nil {
		DemeterTime = big.NewInt(0).SetUint64(*c.DemeterTime)
	}

	var AthenaTime *big.Int
	if c.AthenaTime != nil {
		AthenaTime = big.NewInt(0).SetUint64(*c.AthenaTime)
	}

	var TheseusTime *big.Int
	if c.TheseusTime != nil {
		TheseusTime = big.NewInt(0).SetUint64(*c.TheseusTime)
	}

	var CancunTime *big.Int
	if c.CancunTime != nil {
		CancunTime = big.NewInt(0).SetUint64(*c.CancunTime)
	}

	var TheseusFixTime *big.Int
	if c.TheseusFixTime != nil {
		TheseusFixTime = big.NewInt(0).SetUint64(*c.TheseusFixTime)
	}

	var HaberTime *big.Int
	if c.HaberTime != nil {
		HaberTime = big.NewInt(0).SetUint64(*c.HaberTime)
	}

	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v Petersburg: %v Istanbul: %v, Muir Glacier: %v, Berlin: %v, YOLO v3: %v, London: %v, HashPower: %v, Zeus: %v, Hera: %v, Poseidon: %v, Luban: %v, Plato: %v, Hertz: %v, ShanghaiTime: %v, KeplerTime: %v, DemeterTime: %v, AthenaTime: %v, TheseusTime: %v, CancunTime: %v, TheseusFixTime: %v, HaberTime: %v, Engine: %v}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		c.IstanbulBlock,
		c.MuirGlacierBlock,
		c.BerlinBlock,
		c.YoloV3Block,
		c.LondonBlock,
		c.HashPowerBlock,
		c.ZeusBlock,
		c.HeraBlock,
		c.PoseidonBlock,
		c.LubanBlock,
		c.PlatoBlock,
		c.HertzBlock,
		ShanghaiTime,
		KeplerTime,
		DemeterTime,
		AthenaTime,
		TheseusTime,
		CancunTime,
		TheseusFixTime,
		HaberTime,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isBlockForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isBlockForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isBlockForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isBlockForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isBlockForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isBlockForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isBlockForked(c.ConstantinopleBlock, num)
}

// IsLuban returns whether num is either equal to the first fast finality fork block or greater.
func (c *ChainConfig) IsLuban(num *big.Int) bool {
	return isBlockForked(c.LubanBlock, num)
}

// IsOnLuban returns whether num is equal to the first fast finality fork block.
func (c *ChainConfig) IsOnLuban(num *big.Int) bool {
	return configBlockEqual(c.LubanBlock, num)
}

// IsPlato returns whether num is either equal to the second fast finality fork block or greater.
func (c *ChainConfig) IsPlato(num *big.Int) bool {
	return isBlockForked(c.PlatoBlock, num)
}

// IsOnPlato returns whether num is equal to the second fast finality fork block.
func (c *ChainConfig) IsOnPlato(num *big.Int) bool {
	return configBlockEqual(c.PlatoBlock, num)
}

// IsHertz returns whether num is either equal to the block of enabling Berlin EIPs or greater.
func (c *ChainConfig) IsHertz(num *big.Int) bool {
	return isBlockForked(c.HertzBlock, num)
}

// IsOnHertz returns whether num is equal to the fork block of enabling Berlin EIPs.
func (c *ChainConfig) IsOnHertz(num *big.Int) bool {
	return configBlockEqual(c.HertzBlock, num)
}

// IsMuirGlacier returns whether num is either equal to the Muir Glacier (EIP-2384) fork block or greater.
func (c *ChainConfig) IsMuirGlacier(num *big.Int) bool {
	return isBlockForked(c.MuirGlacierBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isBlockForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isBlockForked(c.ConstantinopleBlock, num)
}

// IsIstanbul returns whether num is either equal to the Istanbul fork block or greater.
func (c *ChainConfig) IsIstanbul(num *big.Int) bool {
	return isBlockForked(c.IstanbulBlock, num)
}

// IsBerlin returns whether num is either equal to the Berlin fork block or greater.
func (c *ChainConfig) IsBerlin(num *big.Int) bool {
	return isBlockForked(c.BerlinBlock, num)
}

// IsLondon returns whether num is either equal to the London fork block or greater.
func (c *ChainConfig) IsLondon(num *big.Int) bool {
	return isBlockForked(c.LondonBlock, num)
}

// IsArrowGlacier returns whether num is either equal to the Arrow Glacier (EIP-4345) fork block or greater.
func (c *ChainConfig) IsArrowGlacier(num *big.Int) bool {
	return isBlockForked(c.ArrowGlacierBlock, num)
}

// IsGrayGlacier returns whether num is either equal to the Gray Glacier (EIP-5133) fork block or greater.
func (c *ChainConfig) IsGrayGlacier(num *big.Int) bool {
	return isBlockForked(c.GrayGlacierBlock, num)
}

// IsTerminalPoWBlock returns whether the given block is the last block of PoW stage.
func (c *ChainConfig) IsTerminalPoWBlock(parentTotalDiff *big.Int, totalDiff *big.Int) bool {
	if c.TerminalTotalDifficulty == nil {
		return false
	}
	return parentTotalDiff.Cmp(c.TerminalTotalDifficulty) < 0 && totalDiff.Cmp(c.TerminalTotalDifficulty) >= 0
}

func (c *ChainConfig) IsHashPower(num *big.Int) bool {
	return isBlockForked(c.HashPowerBlock, num)
}

func (c *ChainConfig) IsOnHashPower(num *big.Int) bool {
	return configBlockEqual(c.HashPowerBlock, num)
}

func (c *ChainConfig) IsZeus(num *big.Int) bool {
	return isBlockForked(c.ZeusBlock, num)
}

func (c *ChainConfig) IsOnZeus(num *big.Int) bool {
	return configBlockEqual(c.ZeusBlock, num)
}

func (c *ChainConfig) IsHera(num *big.Int) bool {
	return isBlockForked(c.HeraBlock, num)
}

func (c *ChainConfig) IsOnHera(num *big.Int) bool {
	return configBlockEqual(c.HeraBlock, num)
}

func (c *ChainConfig) IsPoseidon(num *big.Int) bool {
	return isBlockForked(c.PoseidonBlock, num)
}

func (c *ChainConfig) IsOnPoseidon(num *big.Int) bool {
	return configBlockEqual(c.PoseidonBlock, num)
}

// IsShanghai returns whether time is either equal to the Shanghai fork time or greater.
func (c *ChainConfig) IsShanghai(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.ShanghaiTime, time)
}

// IsOnShanghai returns whether currentBlockTime is either equal to the shanghai fork time or greater firstly.
func (c *ChainConfig) IsOnShanghai(currentBlockNumber *big.Int, lastBlockTime uint64, currentBlockTime uint64) bool {
	lastBlockNumber := new(big.Int)
	if currentBlockNumber.Cmp(big.NewInt(1)) >= 0 {
		lastBlockNumber.Sub(currentBlockNumber, big.NewInt(1))
	}
	return !c.IsShanghai(lastBlockNumber, lastBlockTime) && c.IsShanghai(currentBlockNumber, currentBlockTime)
}

// IsKepler returns whether time is either equal to the kepler fork time or greater.
func (c *ChainConfig) IsKepler(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.KeplerTime, time)
}

// IsOnKepler returns whether currentBlockTime is either equal to the kepler fork time or greater firstly.
func (c *ChainConfig) IsOnKepler(currentBlockNumber *big.Int, lastBlockTime uint64, currentBlockTime uint64) bool {
	lastBlockNumber := new(big.Int)
	if currentBlockNumber.Cmp(big.NewInt(1)) >= 0 {
		lastBlockNumber.Sub(currentBlockNumber, big.NewInt(1))
	}
	return !c.IsKepler(lastBlockNumber, lastBlockTime) && c.IsKepler(currentBlockNumber, currentBlockTime)
}

// IsDemeter returns whether time is either equal to the demeter fork time or greater.
func (c *ChainConfig) IsDemeter(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.DemeterTime, time)
}

// IsOnDemeter returns whether currentBlockTime is either equal to the demeter fork time or greater firstly.
func (c *ChainConfig) IsOnDemeter(currentBlockNumber *big.Int, lastBlockTime uint64, currentBlockTime uint64) bool {
	lastBlockNumber := new(big.Int)
	if currentBlockNumber.Cmp(big.NewInt(1)) >= 0 {
		lastBlockNumber.Sub(currentBlockNumber, big.NewInt(1))
	}
	return !c.IsDemeter(lastBlockNumber, lastBlockTime) && c.IsDemeter(currentBlockNumber, currentBlockTime)
}

// IsAthena returns whether time is either equal to the athena fork time or greater.
func (c *ChainConfig) IsAthena(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.AthenaTime, time)
}

// IsOnAthena returns whether currentBlockTime is either equal to the athena fork time or greater firstly.
func (c *ChainConfig) IsOnAthena(currentBlockNumber *big.Int, lastBlockTime uint64, currentBlockTime uint64) bool {
	lastBlockNumber := new(big.Int)
	if currentBlockNumber.Cmp(big.NewInt(1)) >= 0 {
		lastBlockNumber.Sub(currentBlockNumber, big.NewInt(1))
	}
	return !c.IsAthena(lastBlockNumber, lastBlockTime) && c.IsAthena(currentBlockNumber, currentBlockTime)
}

// IsTheseus returns whether time is either equal to the theseus fork time or greater.
func (c *ChainConfig) IsTheseus(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.TheseusTime, time)
}

// IsOnTheseus returns whether currentBlockTime is either equal to the theseus fork time or greater firstly.
func (c *ChainConfig) IsOnTheseus(currentBlockNumber *big.Int, lastBlockTime uint64, currentBlockTime uint64) bool {
	lastBlockNumber := new(big.Int)
	if currentBlockNumber.Cmp(big.NewInt(1)) >= 0 {
		lastBlockNumber.Sub(currentBlockNumber, big.NewInt(1))
	}
	return !c.IsTheseus(lastBlockNumber, lastBlockTime) && c.IsTheseus(currentBlockNumber, currentBlockTime)
}

// IsTheseusFix returns whether time is either equal to the theseusFix fork time or greater.
func (c *ChainConfig) IsTheseusFix(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.TheseusFixTime, time)
}

// IsOnTheseusFix returns whether currentBlockTime is either equal to the theseusFix fork time or greater firstly.
func (c *ChainConfig) IsOnTheseusFix(currentBlockNumber *big.Int, lastBlockTime uint64, currentBlockTime uint64) bool {
	lastBlockNumber := new(big.Int)
	if currentBlockNumber.Cmp(big.NewInt(1)) >= 0 {
		lastBlockNumber.Sub(currentBlockNumber, big.NewInt(1))
	}
	return !c.IsTheseusFix(lastBlockNumber, lastBlockTime) && c.IsTheseusFix(currentBlockNumber, currentBlockTime)
}

// IsCancun returns whether num is either equal to the Cancun fork time or greater.
func (c *ChainConfig) IsCancun(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.CancunTime, time)
}

// IsHaber returns whether time is either equal to the Haber fork time or greater.
func (c *ChainConfig) IsHaber(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.HaberTime, time)
}

// IsPrague returns whether num is either equal to the Prague fork time or greater.
func (c *ChainConfig) IsPrague(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.PragueTime, time)
}

// IsVerkle returns whether num is either equal to the Verkle fork time or greater.
func (c *ChainConfig) IsVerkle(num *big.Int, time uint64) bool {
	return c.IsLondon(num) && isTimestampForked(c.VerkleTime, time)
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64, time uint64) *ConfigCompatError {
	var (
		bhead = new(big.Int).SetUint64(height)
		btime = time
	)
	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead, btime)
		if err == nil || (lasterr != nil && err.RewindToBlock == lasterr.RewindToBlock && err.RewindToTime == lasterr.RewindToTime) {
			break
		}
		lasterr = err

		if err.RewindToTime > 0 {
			btime = err.RewindToTime
		} else {
			bhead.SetUint64(err.RewindToBlock)
		}
	}
	return lasterr
}

// CheckConfigForkOrder checks that we don't "skip" any forks, geth isn't pluggable enough
// to guarantee that forks can be implemented in a different order than on official networks
func (c *ChainConfig) CheckConfigForkOrder() error {
	// skip checking for non-Satoshi egine
	if c.Satoshi == nil {
		return nil
	}
	type fork struct {
		name      string
		block     *big.Int // forks up to - and including the merge - were defined with block numbers
		timestamp *uint64  // forks after the merge are scheduled using timestamps
		optional  bool     // if true, the fork may be nil and next fork is still allowed
	}
	var lastFork fork
	for _, cur := range []fork{
		{name: "hashPowerBlock", block: c.HashPowerBlock},
		{name: "zeusBlock", block: c.ZeusBlock},
		{name: "heraBlock", block: c.HeraBlock},
		{name: "poseidonBlock", block: c.PoseidonBlock},
		{name: "hertzBlock", block: c.HertzBlock},
		{name: "keplerTime", timestamp: c.KeplerTime},
		{name: "demeterTime", timestamp: c.DemeterTime},
		{name: "athenaTime", timestamp: c.AthenaTime},
		{name: "theseusTime", timestamp: c.TheseusTime},
		{name: "cancunTime", timestamp: c.CancunTime},
		{name: "theseusFixTime", timestamp: c.TheseusFixTime},
		{name: "haberTime", timestamp: c.HaberTime},
		{name: "pragueTime", timestamp: c.PragueTime, optional: true},
		{name: "verkleTime", timestamp: c.VerkleTime, optional: true},
	} {
		if lastFork.name != "" {
			switch {
			// Non-optional forks must all be present in the chain config up to the last defined fork
			case lastFork.block == nil && lastFork.timestamp == nil && (cur.block != nil || cur.timestamp != nil):
				if cur.block != nil {
					return fmt.Errorf("unsupported fork ordering: %v not enabled, but %v enabled at block %v",
						lastFork.name, cur.name, cur.block)
				} else {
					return fmt.Errorf("unsupported fork ordering: %v not enabled, but %v enabled at timestamp %v",
						lastFork.name, cur.name, *cur.timestamp)
				}

			// Fork (whether defined by block or timestamp) must follow the fork definition sequence
			case (lastFork.block != nil && cur.block != nil) || (lastFork.timestamp != nil && cur.timestamp != nil):
				if lastFork.block != nil && lastFork.block.Cmp(cur.block) > 0 {
					return fmt.Errorf("unsupported fork ordering: %v enabled at block %v, but %v enabled at block %v",
						lastFork.name, lastFork.block, cur.name, cur.block)
				} else if lastFork.timestamp != nil && *lastFork.timestamp > *cur.timestamp {
					return fmt.Errorf("unsupported fork ordering: %v enabled at timestamp %v, but %v enabled at timestamp %v",
						lastFork.name, *lastFork.timestamp, cur.name, *cur.timestamp)
				}

				// Timestamp based forks can follow block based ones, but not the other way around
				if lastFork.timestamp != nil && cur.block != nil {
					return fmt.Errorf("unsupported fork ordering: %v used timestamp ordering, but %v reverted to block ordering",
						lastFork.name, cur.name)
				}
			}
		}
		// If it was optional and not set, then ignore it
		if !cur.optional || (cur.block != nil || cur.timestamp != nil) {
			lastFork = cur
		}
	}
	return nil
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, headNumber *big.Int, headTimestamp uint64) *ConfigCompatError {
	if isForkBlockIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, headNumber) {
		return newBlockCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkBlockIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, headNumber) {
		return newBlockCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(headNumber) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newBlockCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkBlockIncompatible(c.EIP150Block, newcfg.EIP150Block, headNumber) {
		return newBlockCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkBlockIncompatible(c.EIP155Block, newcfg.EIP155Block, headNumber) {
		return newBlockCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkBlockIncompatible(c.EIP158Block, newcfg.EIP158Block, headNumber) {
		return newBlockCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(headNumber) && !configBlockEqual(c.ChainID, newcfg.ChainID) {
		return newBlockCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkBlockIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, headNumber) {
		return newBlockCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkBlockIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, headNumber) {
		return newBlockCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkBlockIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, headNumber) {
		// the only case where we allow Petersburg to be set in the past is if it is equal to Constantinople
		// mainly to satisfy fork ordering requirements which state that Petersburg fork be set if Constantinople fork is set
		if isForkBlockIncompatible(c.ConstantinopleBlock, newcfg.PetersburgBlock, headNumber) {
			return newBlockCompatError("Petersburg fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
		}
	}
	if isForkBlockIncompatible(c.IstanbulBlock, newcfg.IstanbulBlock, headNumber) {
		return newBlockCompatError("Istanbul fork block", c.IstanbulBlock, newcfg.IstanbulBlock)
	}
	if isForkBlockIncompatible(c.MuirGlacierBlock, newcfg.MuirGlacierBlock, headNumber) {
		return newBlockCompatError("Muir Glacier fork block", c.MuirGlacierBlock, newcfg.MuirGlacierBlock)
	}
	if isForkBlockIncompatible(c.BerlinBlock, newcfg.BerlinBlock, headNumber) {
		return newBlockCompatError("Berlin fork block", c.BerlinBlock, newcfg.BerlinBlock)
	}
	if isForkBlockIncompatible(c.LondonBlock, newcfg.LondonBlock, headNumber) {
		return newBlockCompatError("London fork block", c.LondonBlock, newcfg.LondonBlock)
	}
	if isForkBlockIncompatible(c.ArrowGlacierBlock, newcfg.ArrowGlacierBlock, headNumber) {
		return newBlockCompatError("Arrow Glacier fork block", c.ArrowGlacierBlock, newcfg.ArrowGlacierBlock)
	}
	if isForkBlockIncompatible(c.GrayGlacierBlock, newcfg.GrayGlacierBlock, headNumber) {
		return newBlockCompatError("Gray Glacier fork block", c.GrayGlacierBlock, newcfg.GrayGlacierBlock)
	}
	if isForkBlockIncompatible(c.HashPowerBlock, newcfg.HashPowerBlock, headNumber) {
		return newBlockCompatError("hashPower fork block", c.HashPowerBlock, newcfg.HashPowerBlock)
	}
	if isForkBlockIncompatible(c.ZeusBlock, newcfg.ZeusBlock, headNumber) {
		return newBlockCompatError("zeus fork block", c.ZeusBlock, newcfg.ZeusBlock)
	}
	if isForkBlockIncompatible(c.HeraBlock, newcfg.HeraBlock, headNumber) {
		return newBlockCompatError("hera fork block", c.HeraBlock, newcfg.HeraBlock)
	}
	if isForkBlockIncompatible(c.PoseidonBlock, newcfg.PoseidonBlock, headNumber) {
		return newBlockCompatError("hera fork block", c.PoseidonBlock, newcfg.PoseidonBlock)
	}
	if isForkBlockIncompatible(c.LubanBlock, newcfg.LubanBlock, headNumber) {
		return newBlockCompatError("luban fork block", c.LubanBlock, newcfg.LubanBlock)
	}
	if isForkBlockIncompatible(c.PlatoBlock, newcfg.PlatoBlock, headNumber) {
		return newBlockCompatError("plato fork block", c.PlatoBlock, newcfg.PlatoBlock)
	}
	if isForkBlockIncompatible(c.HertzBlock, newcfg.HertzBlock, headNumber) {
		return newBlockCompatError("hertz fork block", c.HertzBlock, newcfg.HertzBlock)
	}
	if isForkTimestampIncompatible(c.ShanghaiTime, newcfg.ShanghaiTime, headTimestamp) {
		return newTimestampCompatError("Shanghai fork timestamp", c.ShanghaiTime, newcfg.ShanghaiTime)
	}
	if isForkTimestampIncompatible(c.KeplerTime, newcfg.KeplerTime, headTimestamp) {
		return newTimestampCompatError("Kepler fork timestamp", c.KeplerTime, newcfg.KeplerTime)
	}
	if isForkTimestampIncompatible(c.DemeterTime, newcfg.DemeterTime, headTimestamp) {
		return newTimestampCompatError("Demeter fork timestamp", c.DemeterTime, newcfg.DemeterTime)
	}
	if isForkTimestampIncompatible(c.AthenaTime, newcfg.AthenaTime, headTimestamp) {
		return newTimestampCompatError("Athena fork timestamp", c.AthenaTime, newcfg.AthenaTime)
	}
	if isForkTimestampIncompatible(c.TheseusTime, newcfg.TheseusTime, headTimestamp) {
		return newTimestampCompatError("Theseus fork timestamp", c.TheseusTime, newcfg.TheseusTime)
	}
	if isForkTimestampIncompatible(c.CancunTime, newcfg.CancunTime, headTimestamp) {
		return newTimestampCompatError("Cancun fork timestamp", c.CancunTime, newcfg.CancunTime)
	}
	if isForkTimestampIncompatible(c.TheseusFixTime, newcfg.TheseusFixTime, headTimestamp) {
		return newTimestampCompatError("TheseusFix fork timestamp", c.TheseusFixTime, newcfg.TheseusFixTime)
	}
	if isForkTimestampIncompatible(c.PragueTime, newcfg.PragueTime, headTimestamp) {
		return newTimestampCompatError("Prague fork timestamp", c.PragueTime, newcfg.PragueTime)
	}
	if isForkTimestampIncompatible(c.VerkleTime, newcfg.VerkleTime, headTimestamp) {
		return newTimestampCompatError("Verkle fork timestamp", c.VerkleTime, newcfg.VerkleTime)
	}
	return nil
}

// BaseFeeChangeDenominator bounds the amount the base fee can change between blocks.
func (c *ChainConfig) BaseFeeChangeDenominator() uint64 {
	return DefaultBaseFeeChangeDenominator
}

// ElasticityMultiplier bounds the maximum gas limit an EIP-1559 block may have.
func (c *ChainConfig) ElasticityMultiplier() uint64 {
	return DefaultElasticityMultiplier
}

// LatestFork returns the latest time-based fork that would be active for the given time.
// only include forks from ethereum
func (c *ChainConfig) LatestFork(time uint64) forks.Fork {
	// Assume last non-time-based fork has passed.
	london := c.LondonBlock

	switch {
	case c.IsPrague(london, time):
		return forks.Prague
	case c.IsCancun(london, time):
		return forks.Cancun
	case c.IsShanghai(london, time):
		return forks.Shanghai
	default:
		return forks.Paris
	}
}

// isForkBlockIncompatible returns true if a fork scheduled at block s1 cannot be
// rescheduled to block s2 because head is already past the fork.
func isForkBlockIncompatible(s1, s2, head *big.Int) bool {
	return (isBlockForked(s1, head) || isBlockForked(s2, head)) && !configBlockEqual(s1, s2)
}

// isBlockForked returns whether a fork scheduled at block s is active at the
// given head block. Whilst this method is the same as isTimestampForked, they
// are explicitly separate for clearer reading.
func isBlockForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configBlockEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// isForkTimestampIncompatible returns true if a fork scheduled at timestamp s1
// cannot be rescheduled to timestamp s2 because head is already past the fork.
func isForkTimestampIncompatible(s1, s2 *uint64, head uint64) bool {
	return (isTimestampForked(s1, head) || isTimestampForked(s2, head)) && !configTimestampEqual(s1, s2)
}

// isTimestampForked returns whether a fork scheduled at timestamp s is active
// at the given head timestamp. Whilst this method is the same as isBlockForked,
// they are explicitly separate for clearer reading.
func isTimestampForked(s *uint64, head uint64) bool {
	if s == nil {
		return false
	}
	return *s <= head
}

func configTimestampEqual(x, y *uint64) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return *x == *y
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string

	// block numbers of the stored and new configurations if block based forking
	StoredBlock, NewBlock *big.Int

	// timestamps of the stored and new configurations if time based forking
	StoredTime, NewTime *uint64

	// the block number to which the local chain must be rewound to correct the error
	RewindToBlock uint64

	// the timestamp to which the local chain must be rewound to correct the error
	RewindToTime uint64
}

func newBlockCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{
		What:          what,
		StoredBlock:   storedblock,
		NewBlock:      newblock,
		RewindToBlock: 0,
	}
	if rew != nil && rew.Sign() > 0 {
		err.RewindToBlock = rew.Uint64() - 1
	}
	return err
}

func newTimestampCompatError(what string, storedtime, newtime *uint64) *ConfigCompatError {
	var rew *uint64
	switch {
	case storedtime == nil:
		rew = newtime
	case newtime == nil || *storedtime < *newtime:
		rew = storedtime
	default:
		rew = newtime
	}
	err := &ConfigCompatError{
		What:         what,
		StoredTime:   storedtime,
		NewTime:      newtime,
		RewindToTime: 0,
	}
	if rew != nil && *rew != 0 {
		err.RewindToTime = *rew - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	if err.StoredBlock != nil {
		return fmt.Sprintf("mismatching %s in database (have block %d, want block %d, rewindto block %d)", err.What, err.StoredBlock, err.NewBlock, err.RewindToBlock)
	}

	if err.StoredTime == nil {
		return fmt.Sprintf("mismatching %s in database (have timestamp nil, want timestamp %d, rewindto timestamp %d)", err.What, *err.NewTime, err.RewindToTime)
	} else if err.NewTime == nil {
		return fmt.Sprintf("mismatching %s in database (have timestamp %d, want timestamp nil, rewindto timestamp %d)", err.What, *err.StoredTime, err.RewindToTime)
	}
	return fmt.Sprintf("mismatching %s in database (have timestamp %d, want timestamp %d, rewindto timestamp %d)", err.What, *err.StoredTime, *err.NewTime, err.RewindToTime)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                                 *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158               bool
	IsByzantium, IsConstantinople, IsPetersburg, IsIstanbul bool
	IsBerlin, IsLondon                                      bool
	IsMerge                                                 bool
	IsHashPower                                             bool
	IsShanghai, IsKepler, IsCancun, IsHaber                 bool
	IsPrague, IsVerkle                                      bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int, isMerge bool, timestamp uint64) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	// disallow setting Merge out of order
	isMerge = isMerge && c.IsLondon(num)
	return Rules{
		ChainID:          new(big.Int).Set(chainID),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
		IsPetersburg:     c.IsPetersburg(num),
		IsIstanbul:       c.IsIstanbul(num),
		IsBerlin:         c.IsBerlin(num),
		IsLondon:         c.IsLondon(num),
		IsMerge:          isMerge,
		IsHashPower:      c.IsHashPower(num),
		IsShanghai:       c.IsShanghai(num, timestamp),
		IsKepler:         c.IsKepler(num, timestamp),
		IsCancun:         c.IsCancun(num, timestamp),
		IsHaber:          c.IsHaber(num, timestamp),
		IsPrague:         c.IsPrague(num, timestamp),
		IsVerkle:         c.IsVerkle(num, timestamp),
	}
}
