# Cloud Storage Service

A full-stack cloud storage application built with Go (backend) and React (frontend) that allows users to store, manage, and share files.

## ✨ Features

- **User Authentication**: Secure registration and login using JWT tokens
- **File Management**: Upload, download, and delete files
- **File Sharing**: Share files with other users in the system
- **User-friendly Interface**: Modern React-based dashboard
- **Local Storage**: Files are stored locally on the server
- **Easy Setup**: Automated scripts for quick setup and running

## 🛠 Technology Stack

### Backend
- **Go 1.21+** - High-performance backend language
- **Gin Web Framework** - Fast HTTP web framework
- **GORM** - ORM for database operations
- **SQLite** - Local database
- **JWT** - JSON Web Token for authentication
- **bcrypt** - Password hashing

### Frontend
- **React 18** - UI library
- **Vite** - Modern build tool
- **Axios** - HTTP client
- **CSS3** - Styling

## 📁 Project Structure

```
cloud-storage-service/
├── backend/
│   ├── main.go                 # Application entry point
│   ├── config/
│   │   └── config.go          # Database and environment configuration
│   ├── models/
│   │   ├── user.go            # User model
│   │   ├── file.go            # File model
│   │   └── shared_file.go     # Shared file model
│   ├── handlers/
│   │   ├── auth.go            # Authentication handlers (login/register)
│   │   ├── file.go            # File management handlers
│   │   ├── sharing.go         # File sharing handlers
│   │   └── health.go          # Health check handler
│   ├── routes/
│   │   └── routes.go          # API route definitions
│   └── middleware/
│       └── auth.go            # JWT authentication middleware
├── frontend/
│   ├── index.html             # HTML entry point
│   ├── vite.config.js         # Vite configuration
│   ├── package.json           # Dependencies
│   └── src/
│       ├── main.jsx           # React entry point
│       ├── App.jsx            # Main app component
│       ├── pages/
│       │   ├── LoginPage.jsx  # Login/Register page
│       │   └── DashboardPage.jsx # Main dashboard
│       └── styles/
│           ├── LoginPage.css  # Login page styles
│           └── DashboardPage.css # Dashboard styles
├── build.bat                   # Windows: Build script
├── run.bat                     # Windows: Run production server
├── run-dev.bat                 # Windows: Run development servers
├── setup.bat                   # Windows: Check prerequisites
├── go.mod                      # Go module definition
├── .gitignore                  # Git ignore rules
└── README.md                   # This file
```

## 🚀 Quick Start

### Prerequisites

- **Go** 1.21 or later - [Download](https://golang.org/dl/)
- **Node.js** 16 or later - [Download](https://nodejs.org/)
- **npm** or **yarn** (comes with Node.js)

### Windows Users (Easiest Way!)

We provide batch scripts to automate setup and running. Simply execute:

#### Step 1: Check Prerequisites
```batch
setup.bat
```
This will verify that Go and Node.js are installed correctly.

#### Step 2: Build the Project
```batch
build.bat
```
This will compile the Go backend and build the React frontend.

#### Step 3: Run the Application

**Option A - Production Mode (Recommended):**
```batch
run.bat
```
Runs the built backend server with the compiled frontend assets.

**Option B - Development Mode (Active Development):**
```batch
run-dev.bat
```
Starts both the backend server and Vite frontend dev server with hot reload.

Once running, open your browser to:
- **Frontend**: `http://localhost:3000` (dev mode) or `http://localhost:8080` (production)
- **Backend API**: `http://localhost:8080/api`

### Mac/Linux Users

#### Backend Setup

1. Install Go dependencies:
```bash
cd backend
go mod download
```

2. Set environment variables:
```bash
export JWT_SECRET="your-secret-key-change-in-production"
export PORT="8080"
```

3. Run the backend server:
```bash
go run main.go
```

The backend will be available at `http://localhost:8080`

#### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd frontend
```

2. Install dependencies:
```bash
npm install
```

3. Start the development server:
```bash
npm run dev
```

The frontend will be available at `http://localhost:3000`

#### Running Both Together (Development)

In separate terminal windows:

Terminal 1 (Backend):
```bash
cd backend
go run main.go
```

Terminal 2 (Frontend):
```bash
cd frontend
npm run dev
```

## 📡 API Endpoints

### Authentication
- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Login and get JWT token

### File Management
- `POST /api/files/upload` - Upload a file (requires auth)
- `GET /api/files/list` - List user's files (requires auth)
- `GET /api/files/download/:id` - Download a file (requires auth)
- `DELETE /api/files/:id` - Delete a file (requires auth)

### File Sharing
- `POST /api/share/:file_id` - Share a file with another user (requires auth)
- `GET /api/share/list` - List files shared with the user (requires auth)
- `DELETE /api/share/:share_id` - Unshare a file (requires auth)

## 💻 Usage Guide

1. **Register**: Create a new account with a username, email, and password
2. **Login**: Log in with your credentials
3. **Upload Files**: Click "Upload File" and select files from your computer
4. **Share Files**: Select a file and enter another user's username to share it
5. **Download Files**: Download files from your storage or shared files
6. **Delete Files**: Remove files you no longer need

## 🔐 Security Considerations

- **JWT Tokens**: Tokens expire after 24 hours
- **Password Hashing**: Passwords are hashed using bcrypt with salt
- **CORS**: Configure appropriately for production use
- **File Validation**: Implement file type and size validation in production
- **Environment Variables**: Always change `JWT_SECRET` in production environments
- **HTTPS**: Use HTTPS in production for encrypted communications

## 🔧 Troubleshooting

### Windows Scripts Don't Run
- Make sure you run the `.bat` files from the project root directory
- If you get "file not found" errors, check that Go and Node.js are in your PATH
- Run `setup.bat` first to diagnose any issues

### Port Already in Use
- Backend: Change `PORT` environment variable (default: 8080)
- Frontend: Edit `vite.config.js` to change the dev server port

### Database Issues
- Delete `cloud_storage.db` to start fresh (you'll lose all data)
- Run `setup.bat` to verify Go is installed

### Node Modules Issues
```bash
cd frontend
rm -r node_modules
npm install
```

## 📚 Future Enhancements

- [ ] File versioning
- [ ] Folder structure/hierarchy
- [ ] Advanced sharing with granular permissions
- [ ] File preview functionality (images, PDFs, etc.)
- [ ] File search and filtering
- [ ] User profiles and settings
- [ ] Activity/audit logging
- [ ] Cloud storage integration (S3, Google Cloud Storage)
- [ ] Mobile app (React Native)
- [ ] End-to-end encryption

## 📄 License

MIT License - see LICENSE file for details

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### How to Contribute

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## ❓ Support

For issues and questions:
1. Check the troubleshooting section above
2. Search existing GitHub issues
3. Create a new GitHub issue with detailed description

---

**Made with ❤️ by the Cloud Storage Service Team**
