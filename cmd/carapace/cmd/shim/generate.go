package shim

//go:generate go run ../../../carapace-generate env GOOS=windows GOARCH=386   go build -ldflags "-s -w" -o shim_386.exe   ../../../carapace-shim
//go:generate go run ../../../carapace-generate env GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o shim_amd64.exe ../../../carapace-shim
//go:generate go run ../../../carapace-generate env GOOS=windows GOARCH=arm64 go build -ldflags "-s -w" -o shim_arm64.exe   ../../../carapace-shim
