package janot

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicJanoTAPI provides the janoT RPC service that can be
// use publicly without security implications.
type PublicJanoTAPI struct {
	t        *JanoT
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicJanoTAPI create a new RPC janoT service.
func NewPublicJanoTAPI(t *JanoT) *PublicJanoTAPI {
	api := &PublicJanoTAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the JanoT sub-protocol version.
func (api *PublicJanoTAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
