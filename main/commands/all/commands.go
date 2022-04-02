package all

import (
	"github.com/yuzuki616/xray-core/main/commands/all/api"
	"github.com/yuzuki616/xray-core/main/commands/all/tls"
	"github.com/yuzuki616/xray-core/main/commands/base"
)

// go:generate go run github.com/yuzuki616/xray-core/common/errors/errorgen

func init() {
	base.RootCommand.Commands = append(
		base.RootCommand.Commands,
		api.CmdAPI,
		// cmdConvert,
		tls.CmdTLS,
		cmdUUID,
	)
}
