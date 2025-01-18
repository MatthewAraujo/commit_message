@echo off
setlocal

:: Variables
set "BINARY_NAME=commit_message.exe"
set "CUSTOM_BIN_DIR=%UserProfile%\bin"
set "BINARY_URL=https://github.com/MatthewAraujo/commit_message/releases/download/binary/commit_message.exe"

:: Check for curl
where curl >nul 2>nul
if errorlevel 1 (
    echo Error: curl is required but not installed. Please install curl and try again.
    exit /b 1
)

:: Create the custom directory if it doesn't exist
if not exist "%CUSTOM_BIN_DIR%" (
    echo Creating directory: %CUSTOM_BIN_DIR%
    mkdir "%CUSTOM_BIN_DIR%"
)

:: Download the binary file
echo Downloading %BINARY_NAME% from %BINARY_URL%
curl -L -o "%CUSTOM_BIN_DIR%\%BINARY_NAME%" "%BINARY_URL%"
if errorlevel 1 (
    echo Error: Failed to download the binary from %BINARY_URL%
    exit /b 1
)

:: Add the custom directory to the PATH
echo Adding %CUSTOM_BIN_DIR% to the PATH
for /f "tokens=2* delims=;" %%A in ('reg query "HKCU\Environment" /v PATH 2^>nul') do (
    set "CURRENT_PATH=%%B"
)
if not defined CURRENT_PATH set "CURRENT_PATH="
setx PATH "%CUSTOM_BIN_DIR%;%CURRENT_PATH%" >nul

:: Confirm completion
echo.
echo Installation complete! You can now run "%BINARY_NAME%" from any command prompt.
