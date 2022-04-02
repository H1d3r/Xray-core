package tagged

import (
	"context"

	"github.com/yuzuki616/xray-core/common/net"
)

type DialFunc func(ctx context.Context, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
