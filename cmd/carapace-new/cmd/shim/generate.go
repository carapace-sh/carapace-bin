package shim

//go:generate sh -c "GOOS=windows GOARCH=386   go build -ldflags '-s -w' -o shim_386.exe   ../../../carapace-shim"
//go:generate sh -c "GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o shim_amd64.exe ../../../carapace-shim"
//go:generate sh -c "GOOS=windows GOARCH=arm64 go build -ldflags '-s -w' -o shim_arm64.exe   ../../../carapace-shim"
