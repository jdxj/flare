package main

import (
	"flare/internal/cmd"
	_ "flare/internal/packed"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	ctx := gctx.GetInitCtx()
	root := cmd.NewRoot()
	if err := root.Execute(); err != nil {
		g.Log().Fatal(ctx, err)
	}
}
