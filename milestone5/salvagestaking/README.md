# ScriptStakingUnStaking


$env:GOOS = "darwin"
$env:GOARCH = "arm64"
go build -o script-macos-arm64 .\main.go

$env:GOOS = $null
$env:GOARCH = $null

