package main

import (
	_ "demo/internal/packed"

	// Import PostgreSQL driver for GoFrame
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
