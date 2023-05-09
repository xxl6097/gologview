package tcptest

import (
	"embed"
	"htmltest/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
