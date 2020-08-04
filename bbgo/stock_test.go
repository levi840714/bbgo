package bbgo

import (
	"github.com/c9s/bbgo/pkg/bbgo/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStockManager(t *testing.T) {

	t.Run("stock", func(t *testing.T) {
		var trades = []types.Trade{
			{Symbol: "BTCUSDT", Price: 9100.0, Quantity: 0.05, IsBuyer: true},
			{Symbol: "BTCUSDT", Price: 9100.0, Quantity: 0.05, IsBuyer: true},
			{Symbol: "BTCUSDT", Price: 9200.0, Quantity: 0.01, IsBuyer: false},
		}

		var stockManager = &StockManager{
			TradingFeeCurrency: "BNB",
			Symbol:             "BTCUSDT",
		}

		_, err := stockManager.LoadTrades(trades)
		assert.NoError(t, err)
		assert.Len(t, stockManager.Stocks, 2)
		assert.Len(t, stockManager.PendingSells, 0)
	})

	t.Run("sold out", func(t *testing.T) {
		var trades = []types.Trade{
			{Symbol: "BTCUSDT", Price: 9100.0, Quantity: 0.05, IsBuyer: true},
			{Symbol: "BTCUSDT", Price: 9200.0, Quantity: 0.05, IsBuyer: false},
			{Symbol: "BTCUSDT", Price: 9100.0, Quantity: 0.05, IsBuyer: true},
			{Symbol: "BTCUSDT", Price: 9200.0, Quantity: 0.05, IsBuyer: false},
		}

		var stockManager = &StockManager{
			TradingFeeCurrency: "BNB",
			Symbol:             "BTCUSDT",
		}

		_, err := stockManager.LoadTrades(trades)
		assert.NoError(t, err)
		assert.Len(t, stockManager.Stocks, 0)
		assert.Len(t, stockManager.PendingSells, 0)
	})

	t.Run("oversell", func(t *testing.T) {
		var trades = []types.Trade{
			{Symbol: "BTCUSDT", Price: 9100.0, Quantity: 0.05, IsBuyer: true},
			{Symbol: "BTCUSDT", Price: 9200.0, Quantity: 0.05, IsBuyer: false},
			{Symbol: "BTCUSDT", Price: 9200.0, Quantity: 0.05, IsBuyer: false},
		}

		var stockManager = &StockManager{
			TradingFeeCurrency: "BNB",
			Symbol:             "BTCUSDT",
		}

		_, err := stockManager.LoadTrades(trades)
		assert.NoError(t, err)
		assert.Len(t, stockManager.Stocks, 0)
		assert.Len(t, stockManager.PendingSells, 1)
	})


	t.Run("loss sell", func(t *testing.T) {
		var trades = []types.Trade{
			{Symbol: "BTCUSDT", Price: 9100.0, Quantity: 0.05, IsBuyer: true},
			{Symbol: "BTCUSDT", Price: 9200.0, Quantity: 0.02, IsBuyer: false},
			{Symbol: "BTCUSDT", Price: 8000.0, Quantity: 0.01, IsBuyer: false},
		}

		var stockManager = &StockManager{
			TradingFeeCurrency: "BNB",
			Symbol:             "BTCUSDT",
		}

		_, err := stockManager.LoadTrades(trades)
		assert.NoError(t, err)
		assert.Len(t, stockManager.Stocks, 1)
		assert.Len(t, stockManager.PendingSells, 0)
	})

	t.Run("pending sell", func(t *testing.T) {
		var trades = []types.Trade{
			{Symbol: "BTCUSDT", Price: 9200.0, Quantity: 0.02, IsBuyer: false},
			{Symbol: "BTCUSDT", Price: 9100.0, Quantity: 0.05, IsBuyer: true},
		}

		var stockManager = &StockManager{
			TradingFeeCurrency: "BNB",
			Symbol:             "BTCUSDT",
		}

		_, err := stockManager.LoadTrades(trades)
		assert.NoError(t, err)
		assert.Len(t, stockManager.Stocks, 1)
		assert.Len(t, stockManager.PendingSells, 0)
	})


}