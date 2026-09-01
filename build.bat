@echo off
REM Cloud Storage Service - Build Script for Windows

echo ================================
echo Cloud Storage Service - Build
echo ================================

REM Check if Go is installed
where go >nul 2>nul
if %errorlevel% neq 0 (
    echo Error: Go is not installed or not in PATH
    echo Please install Go from https://golang.org/dl/
    pause
    exit /b 1
)

echo.
echo Building Go backend...
cd backend
go build -o ..\cloud-storage-backend.exe
if %errorlevel% neq 0 (
    echo Error: Failed to build backend
    pause
    exit /b 1
)
cd ..
echo Backend built successfully: cloud-storage-backend.exe

echo.
echo Building React frontend...
cd frontend

REM Check if Node is installed
where node >nul 2>nul
if %errorlevel% neq 0 (
    echo Error: Node.js is not installed or not in PATH
    echo Please install Node.js from https://nodejs.org/
    cd ..
    pause
    exit /b 1
)

REM Install dependencies if node_modules doesn't exist
if not exist "node_modules" (
    echo Installing dependencies...
    call npm install
    if %errorlevel% neq 0 (
        echo Error: Failed to install dependencies
        cd ..
        pause
        exit /b 1
    )
)

echo Building frontend...
call npm run build
if %errorlevel% neq 0 (
    echo Error: Failed to build frontend
    cd ..
    pause
    exit /b 1
)
cd ..
echo Frontend built successfully: frontend/dist

echo.
echo ================================
echo Build completed successfully!
echo ================================
echo.
echo Next steps:
echo 1. Run: run.bat
echo 2. Open your browser to http://localhost:8080
echo.
pause
