package coins

import (
	"github.com/kvant-node/core/state/bus"
	"github.com/kvant-node/core/types"
	"math/big"
)

type Bus struct {
	coins *Coins
}

func NewBus(coins *Coins) *Bus {
	return &Bus{coins: coins}
}

func (b *Bus) GetCoin(symbol types.CoinSymbol) *bus.Coin {
	coin := b.coins.GetCoin(symbol)
	if coin == nil {
		return nil
	}

	return &bus.Coin{
		Name:    coin.Name(),
		Crr:     coin.Crr(),
		Symbol:  coin.Symbol(),
		Volume:  coin.Volume(),
		Reserve: coin.Reserve(),
	}
}

func (b *Bus) SubCoinVolume(symbol types.CoinSymbol, amount *big.Int) {
	b.coins.SubVolume(symbol, amount)
}

func (b *Bus) SubCoinReserve(symbol types.CoinSymbol, amount *big.Int) {
	b.coins.SubReserve(symbol, amount)
}
