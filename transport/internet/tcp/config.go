package tcp

import (
	"github.com/yuzuki616/xray-core/common"
	"github.com/yuzuki616/xray-core/transport/internet"
)

const protocolName = "tcp"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
