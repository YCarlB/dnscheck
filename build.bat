@echo off
REM Set Go environment variables for Windows DLL build
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1

REM Specify output file name
set OUTPUT_FILE=dnscheck.dll

REM Optional: Set custom C flags to suppress warnings
set CGO_CFLAGS=-Wno-nullability-completeness

REM Build the DLL with c-shared build mode
go build -o %OUTPUT_FILE% -buildmode=c-shared

REM Check if the build was successful
if %ERRORLEVEL% equ 0 (
    echo DLL successfully built as %OUTPUT_FILE%
) else (
    echo Build failed.
)
pause
