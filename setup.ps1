#Requires -Version 5.1
$ErrorActionPreference = "Stop"

Write-Host "Setting up TUI project..."

if (-not (Get-Command winget -ErrorAction SilentlyContinue)) {
    Write-Error "winget not found. Please install App Installer from Microsoft Store."
    exit 1
}

if (-not (Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Installing Go..."
    winget install --id GoLang.Go --silent --accept-package-agreements --accept-source-agreements
    $env:Path = [System.Environment]::GetEnvironmentVariable("Path", "Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path", "User")
}

Write-Host "Downloading Go dependencies..."
go mod download

Write-Host ""
Write-Host "Setup complete! Try these commands:"
Write-Host "  go run .              Launch the TUI application"
Write-Host "  set DEBUG=1 && go run . Launch with debug logging"
Write-Host "  go build -o bin\myapp Build a binary"
Write-Host ""
