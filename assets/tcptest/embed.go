package tcptest

import (
	"embed"
	"gologview/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
