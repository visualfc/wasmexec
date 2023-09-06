#!/bin/sh
curl -o wasm/wasm_exec_go116.js https://raw.githubusercontent.com/golang/go/go1.16.15/misc/wasm/wasm_exec.js
curl -o wasm/wasm_exec_go117.js https://raw.githubusercontent.com/golang/go/go1.17.13/misc/wasm/wasm_exec.js
curl -o wasm/wasm_exec_go118.js https://raw.githubusercontent.com/golang/go/go1.18.10/misc/wasm/wasm_exec.js
curl -o wasm/wasm_exec_go119.js https://raw.githubusercontent.com/golang/go/go1.19.12/misc/wasm/wasm_exec.js
curl -o wasm/wasm_exec_go120.js https://raw.githubusercontent.com/golang/go/go1.20.7/misc/wasm/wasm_exec.js
curl -o wasm/wasm_exec_go121.js https://raw.githubusercontent.com/golang/go/go1.21.0/misc/wasm/wasm_exec.js