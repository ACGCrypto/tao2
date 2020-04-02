package simulation

import (
	"github.com/taoblockchain/tao2/crypto"
	"math/big"
)

var (
	RpcEndpoint = "http://127.0.0.1:8501/"
	MinApply    = big.NewInt(0).Mul(big.NewInt(100), big.NewInt(100000000000000000)) // 100 TAO
	Cap         = big.NewInt(0).Mul(big.NewInt(10000000000000), big.NewInt(10000000000000))
	Fee         = big.NewInt(100)

	MainKey, _ = crypto.HexToECDSA("acfef3ae35772903ba4490a17859379f26195e5dc3e4ba30fbe2c4cb64f76ad6")
	MainAddr   = crypto.PubkeyToAddress(MainKey.PublicKey) //0x1390A0602bda8BE9e1ecE5f2595b9024df6A41D8

	AirdropKey, _  = crypto.HexToECDSA("c64e48fe2f059cbec3bf973cf46eebde8622aac9a3ad64763bb94b33f6aae7e3")
	AirdropAddr    = crypto.PubkeyToAddress(AirdropKey.PublicKey) // 0x20802871C750C2d7fbdB49EC9d04151D8023efDB
	AirDropAmount  = big.NewInt(10000000000)
	TransferAmount = big.NewInt(100000)

	ReceiverKey, _ = crypto.HexToECDSA("17534b6c64daba014b67f55b858243dc586b351a033fca4b29a435238f0dc2ca")
	ReceiverAddr   = crypto.PubkeyToAddress(ReceiverKey.PublicKey) //0x268F3dB7a8cc5ccA2C9052c379D96bB8f2Fd0658
)
