package engine

import (
	"errors"
	"sync"
	"time"

	"github.com/thrasher-corp/gocryptotrader/currency"
	"github.com/thrasher-corp/gocryptotrader/exchanges/order"
)

// OrderManagerName is an exported subsystem name
const OrderManagerName = "orders"

// vars for the fund manager package
var (
	// ErrOrdersAlreadyExists occurs when the order already exists in the manager
	ErrOrdersAlreadyExists = errors.New("order already exists")
	// ErrOrderIDCannotBeEmpty occurs when an order does not have an ID
	ErrOrderIDCannotBeEmpty = errors.New("orderID cannot be empty")
	// ErrOrderNotFound occurs when an order is not found in the orderstore
	ErrOrderNotFound = errors.New("order does not exist")

	errNilCommunicationsManager = errors.New("cannot start with nil communications manager")
	errNilOrder                 = errors.New("nil order received")
	errFuturesTrackerNotSetup   = errors.New("futures position tracker not setup")

	orderManagerDelay = time.Second * 10
)

type orderManagerConfig struct {
	EnforceLimitConfig     bool
	AllowMarketOrders      bool
	CancelOrdersOnShutdown bool
	LimitAmount            float64
	AllowedPairs           currency.Pairs
	AllowedExchanges       []string
	OrderSubmissionRetries int64
}

// store holds all orders by exchange
type store struct {
	m                         sync.RWMutex
	Orders                    map[string][]*order.Detail
	commsManager              iCommsManager
	exchangeManager           iExchangeManager
	wg                        *sync.WaitGroup
	futuresPositionController *order.PositionController
}

// OrderManager processes and stores orders across enabled exchanges
type OrderManager struct {
	started          int32
	processingOrders int32
	shutdown         chan struct{}
	orderStore       store
	cfg              orderManagerConfig
	verbose          bool
}

// OrderSubmitResponse contains the order response along with an internal order ID
type OrderSubmitResponse struct {
	*order.Detail
	InternalOrderID string
}

// OrderUpsertResponse contains a copy of the resulting order details and a bool
// indicating if the order details were inserted (true) or updated (false)
type OrderUpsertResponse struct {
	OrderDetails order.Detail
	IsNewOrder   bool
}
