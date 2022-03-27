# do not run it again after you fix your /web folder
GOROOT=$(go env GOROOT)
cp $GOROOT/misc/wasm/wasm_exec.* ./web
