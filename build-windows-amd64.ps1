$env:GOOS = 'windows'
$env:GOARCH = 'amd64'

Set-Location $PSScriptRoot

$process = Start-Process go -ArgumentList 'build -ldflags="-s -w" -o bin\windows-amd64\sonos-standup.exe' -wait -NoNewWindow -PassThru
if ($process.ExitCode -ne 0) {
    Write-Warning "Unable to build sonos standup system for linux"
    Read-Host -Prompt "Press any key to exit"
    Exit 1
}

$process = Start-Process upx -ArgumentList '--brute bin\windows-amd64\sonos-standup.exe' -wait -NoNewWindow -PassThru
if ($process.ExitCode -ne 0) {
    Write-Warning "Unable to compress sonos standup system for linux"
    Read-Host -Prompt "Press any key to exit"
    Exit 1
}