@echo off
setlocal

set BINARY_NAME=commit_message.exe

if not exist "%BINARY_NAME%" (
    echo Binary "%BINARY_NAME%" not found in the current directory.
    echo Please place the binary in the same folder as this script and run again.
    exit /b 1
)

set CUSTOM_BIN_DIR=%UserProfile%\bin
if not exist "%CUSTOM_BIN_DIR%" (
    echo Creating directory: %CUSTOM_BIN_DIR%
    mkdir "%CUSTOM_BIN_DIR%"
)

echo Copying %BINARY_NAME% to %CUSTOM_BIN_DIR%
copy "%BINARY_NAME%" "%CUSTOM_BIN_DIR%\"

echo Adding %CUSTOM_BIN_DIR% to the PATH
setx PATH "%CUSTOM_BIN_DIR%;%PATH%" >nul

echo Setup complete. You can now run "%BINARY_NAME%" from any command prompt.
pause