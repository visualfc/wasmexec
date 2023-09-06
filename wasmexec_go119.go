//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package wasmexec

import (
	_ "embed"
)

//go:embed wasm/wasm_exec_go119.js
var JS string
