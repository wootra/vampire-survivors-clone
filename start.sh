cd ./wasm
GOOS=js GOARCH=wasm go build -o ../assets/result.wasm
cd ..
cp -f ./asset-src/* ./assets
go run ./server