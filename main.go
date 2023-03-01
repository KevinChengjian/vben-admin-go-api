package main

import (
	_ "basic.com/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"basic.com/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
