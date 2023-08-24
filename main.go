package main

import (
	_ "myblog/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"myblog/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
