package wasmexec_test

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/visualfc/wasmexec"
)

func TestWasmExec(t *testing.T) {
	ver := strings.Replace(runtime.Version()[:6], ".", "", 1)
	data, err := ioutil.ReadFile(filepath.Join("wasm", "wasm_exec_"+ver+".js"))
	if err != nil {
		t.Fatal("found wasm_exec failed", ver)
	}
	if wasmexec.JS != string(data) {
		t.Fatal("not match wasm_exec.js data", ver)
	}
}
