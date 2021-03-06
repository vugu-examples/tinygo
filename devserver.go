// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/vugu/vugu/devutil"
)

func main() {
	l := "127.0.0.1:8844"
	log.Printf("Starting HTTP Server at %q", l)

	wc := devutil.MustNewTinygoCompiler().SetDir(".")
	defer wc.Close()

	// wc.NoDocker() // uncomment to use locally installed tinygo
	// wc.SetTinygoArgs("-no-debug") // remove debug info for even smaller file size

	// wc.AddGoGet("go get -u -x github.com/vugu/vjson github.com/vugu/html github.com/vugu/xxhash") // if not using docker, you'll need this
	// wc.AddPkgReplace("github.com/vugu/vugu", "../vugu") // maintainer-only
	wc.AddGoGet("go get -u -x github.com/vugu/vugu github.com/vugu/vjson") // third party packages must have `go get` run on them for tinygo to compile (for now)

	mux := devutil.NewMux()
	mux.Match(devutil.NoFileExt, devutil.DefaultIndex.Replace(
		`<!-- styles -->`,
		`<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css">`))
	mux.Exact("/main.wasm", devutil.NewMainWasmHandler(wc))
	mux.Exact("/wasm_exec.js", devutil.NewWasmExecJSHandler(wc))
	mux.Default(devutil.NewFileServer().SetDir("."))

	log.Fatal(http.ListenAndServe(l, mux))
}
