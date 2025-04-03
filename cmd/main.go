package main

import (
	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/server"
)

func main() {
	app := server.New()

	app.GetInfo()
}
