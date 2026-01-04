@echo off
setlocal enabledelayedexpansion

REM ========================================
REM Build script for myproject
REM Usage:
REM   build.bat debug
REM   build.bat release
REM ========================================

REM --- 默认配置 ---
set BUILD_TYPE=debug
set OUTPUT_DIR=bin

REM --- 解析参数 ---
if "%1"=="" (
    echo No build type specified, default to debug
) else (
    set BUILD_TYPE=%1
)

echo Building type: %BUILD_TYPE%

REM --- 输出目录 ---
if not exist "!OUTPUT_DIR!\!BUILD_TYPE!" (
    mkdir "!OUTPUT_DIR!\!BUILD_TYPE!"
)

REM --- 编译参数 ---
set EXE_NAME=app1.exe

if /I "%BUILD_TYPE%"=="debug" (
    echo Debug build: no optimizations
    REM batch 里 gcflags 的引号要写在 go build 命令里
    go build -gcflags=all="-N -l" -o %OUTPUT_DIR%\%BUILD_TYPE%\%EXE_NAME% ./cmd/app1
) else if /I "%BUILD_TYPE%"=="release" (
    echo Release build: optimizations enabled
    go build -ldflags="-s -w" -o %OUTPUT_DIR%\%BUILD_TYPE%\%EXE_NAME% ./cmd/app1
) else (
    echo Unknown build type: %BUILD_TYPE%
    exit /b 1
)

if errorlevel 1 (
    echo Build failed!
    exit /b 1
)

echo Build succeeded! Output: %OUTPUT_DIR%\%EXE_NAME%
endlocal


