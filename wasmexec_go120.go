//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package wasmexec

import (
	_ "embed"
)

//go:embed wasm/wasm_exec_go120.js
var JS string
