//go:build go1.21
// +build go1.21

package wasmexec

import (
	_ "embed"
)

//go:embed wasm/wasm_exec_go121.js
var JS string
