//go:build go1.17 && !go1.18
// +build go1.17,!go1.18

package wasmexec

import (
	_ "embed"
)

//go:embed wasm/wasm_exec_go117.js
var JS string
