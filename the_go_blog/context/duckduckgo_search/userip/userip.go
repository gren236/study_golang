package userip

import (
	"context"
	"fmt"
	"net"
	"net/http"
)

// The key type is unexported to prevent collisions with context keys defined in other packages.
type key int

// userIPKey is a context key for the user IP address. Its value of zero is arbitrary.
// If this package defined other context keys, they would have different int values
const userIPKey key = 0

func FromRequest(r *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	}
	return net.ParseIP(ip), nil
}

func NewContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

func FromContext(ctx context.Context) (net.IP, bool) {
	// ctx.Value returns nil if ctx has no value for the key;
	// the net.IP type assertion returns ok=false for nil
	userIP, ok := ctx.Value(userIPKey).(net.IP)
	return userIP, ok
}