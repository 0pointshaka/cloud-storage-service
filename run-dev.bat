@echo off
REM Cloud Storage Service - Development Mode (without build)

echo ================================
echo Cloud Storage Service - Dev Mode
echo ================================
echo.

REM Check if backend executable exists
if not exist "cloud-storage-backend.exe" (
    echo Building backend...
    cd backend
    go build -o ..\cloud-storage-backend.exe
    if %errorlevel% neq 0 (
        echo Error: Failed to build backend
        cd ..
        pause
        exit /b 1
    )
    cd ..
    echo Backend built successfully!
)

echo.
echo Starting Cloud Storage Service in Development Mode...
echo.
echo Backend starting on port 8080...
echo Frontend dev server starting on port 3000...
echo.
echo Open your browser to: http://localhost:3000
echo.
echo Press Ctrl+C to stop the server
echo.

REM Start backend in a new window
start "Cloud Storage Backend" cmd /k "cloud-storage-backend.exe"

REM Wait a moment for backend to start
timeout /t 2 /nobreak

REM Start frontend development server
cd frontend

REM Install dependencies if needed
if not exist "node_modules" (
    echo Installing frontend dependencies...
    call npm install
)

echo Starting Vite dev server...
call npm run dev
