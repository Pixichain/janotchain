package vm

import (
	"github.com/janotchain/janotchain/common"
	"github.com/janotchain/janotchain/log"
	"github.com/janotchain/janotchain/params"
	"github.com/janotchain/janotchain/janot/tradingstate"
)
const JanotPriceNumberOfBytesReturn = 32
// janotPrice implements a pre-compile contract to get token price in janot

type janotLastPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}
type janotEpochPrice struct {
	tradingStateDB *tradingstate.TradingStateDB
}

func (t *janotLastPrice) RequiredGas(input []byte) uint64 {
	return params.JanotPriceGas
}

func (t *janotLastPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:]) // 20 bytes from 45-64
		price := t.tradingStateDB.GetLastPrice(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetLastPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), JanotPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, JanotPriceNumberOfBytesReturn), nil
}

func (t *janotLastPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}

func (t *janotEpochPrice) RequiredGas(input []byte) uint64 {
	return params.JanotPriceGas
}

func (t *janotEpochPrice) Run(input []byte) ([]byte, error) {
	// input includes baseTokenAddress, quoteTokenAddress
	if t.tradingStateDB != nil && len(input) == 64 {
		base := common.BytesToAddress(input[12:32]) // 20 bytes from 13-32
		quote := common.BytesToAddress(input[44:]) // 20 bytes from 45-64
		price := t.tradingStateDB.GetMediumPriceBeforeEpoch(tradingstate.GetTradingOrderBookHash(base, quote))
		if price != nil {
			log.Debug("Run GetEpochPrice", "base", base.Hex(), "quote", quote.Hex(), "price", price)
			return common.LeftPadBytes(price.Bytes(), JanotPriceNumberOfBytesReturn), nil
		}
	}
	return common.LeftPadBytes([]byte{}, JanotPriceNumberOfBytesReturn), nil
}

func (t *janotEpochPrice) SetTradingState(tradingStateDB *tradingstate.TradingStateDB) {
	if tradingStateDB != nil {
		t.tradingStateDB = tradingStateDB.Copy()
	} else {
		t.tradingStateDB = nil
	}
}


