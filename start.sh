cd ./wasm
GOOS=js GOARCH=wasm go build -o ../public/result.wasm
cd ..
cp -rf ./web/* ./public
go run ./server