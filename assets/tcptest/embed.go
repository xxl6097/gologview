package tcptest

import (
	"embed"
	"github.com/xxl6097/gologview/assets"
)

//go:embed static/*
var content embed.FS

func init() {
	assets.Register(content)
}
