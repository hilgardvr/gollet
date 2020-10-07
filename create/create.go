package create

import (
	// "encoding/hex"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

type Network struct {
	name		string
	symbol		string
	xpubkey		byte
	xprivkey	byte
}

func (network Network)GetNetworkParams() *chaincfg.Params {
	networkParams:= &chaincfg.MainNetParams
	networkParams.PubKeyHashAddrID = network.xpubkey
	networkParams.PrivateKeyID = network.xprivkey
	return networkParams
}

func (network Network)CreatePrivateKey() (*btcutil.WIF, error) {
	secret, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return nil, err
	}
	return btcutil.NewWIF(secret, network.GetNetworkParams(), true)
}

// func (network Network)ImportPrivateKey(secretHex string) (*btcutil.WIF, error) {}

func (network Network)ImportWIF(wifStr string) (*btcutil.WIF, error) {
	wif, err:= btcutil.DecodeWIF(wifStr)
	//wif is malformed?
	if err != nil {
		return nil, err
	}
	//wif is for btc
	if !wif.IsForNet(network.GetNetworkParams()) {
		return nil, errors.New("The WIF string is not valid for the `" + network.name + "` network")
	}
	return wif, nil
}

func (network Network)GetAddress(wif *btcutil.WIF) (*btcutil.AddressPubKey, error) {
	return btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), network.GetNetworkParams())
}

func CreateGollet() {
	btcNetwork:= Network{name: "bitcoin", symbol: "btc", xpubkey: 0x00, xprivkey: 0x80}
	fmt.Println("Creating a wallet")
	wif, _ := btcNetwork.CreatePrivateKey()
	address, _:= btcNetwork.GetAddress(wif)
	fmt.Printf("%s - %s", wif.String(), address.EncodeAddress())
}


