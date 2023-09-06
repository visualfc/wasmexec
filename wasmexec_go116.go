//go:build go1.16 && !go1.17
// +build go1.16,!go1.17

package wasmexec

import (
	_ "embed"
)

//go:embed wasm/wasm_exec_go116.js
var JS string
