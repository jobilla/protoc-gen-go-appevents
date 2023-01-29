package main

import (
	"github.com/jobilla/protoc-gen-go-appevents/module"
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"
)

func main() {
	pgs.Init(pgs.DebugEnv("DEBUG")).
		RegisterModule(module.Generator()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
