package janotlending

import (
	"context"
	"errors"
	"sync"
	"time"
)

// List of errors
var (
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicJanoTLendingAPI provides the janoT RPC service that can be
// use publicly without security implications.
type PublicJanoTLendingAPI struct {
	t        *Lending
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicJanoTLendingAPI create a new RPC janoT service.
func NewPublicJanoTLendingAPI(t *Lending) *PublicJanoTLendingAPI {
	api := &PublicJanoTLendingAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Lending sub-protocol version.
func (api *PublicJanoTLendingAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
