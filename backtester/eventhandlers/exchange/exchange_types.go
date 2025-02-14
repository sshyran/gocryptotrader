package exchange

import (
	"errors"

	"github.com/shopspring/decimal"
	"github.com/thrasher-corp/gocryptotrader/backtester/data"
	"github.com/thrasher-corp/gocryptotrader/backtester/eventtypes/fill"
	"github.com/thrasher-corp/gocryptotrader/backtester/eventtypes/order"
	"github.com/thrasher-corp/gocryptotrader/backtester/funding"
	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/engine"
	"github.com/thrasher-corp/gocryptotrader/exchanges/asset"
	gctorder "github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

var (
	errDataMayBeIncorrect     = errors.New("data may be incorrect")
	errExceededPortfolioLimit = errors.New("exceeded portfolio limit")
	errNilCurrencySettings    = errors.New("received nil currency settings")
	errInvalidDirection       = errors.New("received invalid order direction")
)

// ExecutionHandler interface dictates what functions are required to submit an order
type ExecutionHandler interface {
	SetExchangeAssetCurrencySettings(string, asset.Item, currency.Pair, *Settings)
	GetCurrencySettings(string, asset.Item, currency.Pair) (Settings, error)
	ExecuteOrder(order.Event, data.Handler, *engine.OrderManager, funding.IPairReleaser) (*fill.Fill, error)
	Reset()
}

// Exchange contains all the currency settings
type Exchange struct {
	CurrencySettings []Settings
}

// Settings allow the eventhandler to size an order within the limitations set by the config file
type Settings struct {
	Exchange      string
	UseRealOrders bool

	Pair  currency.Pair
	Asset asset.Item

	ExchangeFee decimal.Decimal
	MakerFee    decimal.Decimal
	TakerFee    decimal.Decimal

	BuySide  MinMax
	SellSide MinMax

	Leverage Leverage

	MinimumSlippageRate decimal.Decimal
	MaximumSlippageRate decimal.Decimal

	Limits                  gctorder.MinMaxLevel
	CanUseExchangeLimits    bool
	SkipCandleVolumeFitting bool
}

// MinMax are the rules which limit the placement of orders.
type MinMax struct {
	MinimumSize  decimal.Decimal
	MaximumSize  decimal.Decimal
	MaximumTotal decimal.Decimal
}

// Leverage rules are used to allow or limit the use of leverage in orders
// when supported
type Leverage struct {
	CanUseLeverage                 bool
	MaximumOrdersWithLeverageRatio decimal.Decimal
	MaximumLeverageRate            decimal.Decimal
}
