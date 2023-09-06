//go:build go1.18 && !go1.19
// +build go1.18,!go1.19

package wasmexec

import (
	_ "embed"
)

//go:embed wasm/wasm_exec_go118.js
var JS string
