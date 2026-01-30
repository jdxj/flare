package main

import (
	_ "flare/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"flare/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
