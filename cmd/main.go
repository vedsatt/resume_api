package main

import (
	"git.miem.hse.ru/ps-biv24x/aisavelev.git/internal/api"
)

func main() {
	app := api.New()

	app.GetInfo()
}
