$env:GOOS = 'linux'
$env:GOARCH = 'amd64'

Set-Location $PSScriptRoot

$process = Start-Process go -ArgumentList "build -o bin\linux\sonos-standup" -wait -NoNewWindow -PassThru
if ($process.ExitCode -ne 0) {
    Write-Warning "Unable to build sonos standup system for linux"
    Read-Host -Prompt "Press any key to exit"
    Exit 1
}
