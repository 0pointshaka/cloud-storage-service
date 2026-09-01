# Cloud Storage Service

A full-stack cloud storage application built with Go (backend) and React (frontend) that allows users to store, manage, and share files.

## Features

- **User Authentication**: Secure registration and login using JWT tokens
- **File Management**: Upload, download, and delete files
- **File Sharing**: Share files with other users in the system
- **User-friendly Interface**: Modern React-based dashboard
- **Local Storage**: Files are stored locally on the server

## Technology Stack

### Backend
- **Go 1.21**
- **Gin Web Framework** - Fast HTTP web framework
- **GORM** - ORM for database operations
- **SQLite** - Local database
- **JWT** - JSON Web Token for authentication
- **bcrypt** - Password hashing

### Frontend
- **React 18** - UI library
- **Vite** - Build tool
- **Axios** - HTTP client
- **CSS3** - Styling

## Project Structure

```
cloud-storage-service/
├── backend/
│   ├── main.go                 # Application entry point
│   ├── config/
│   │   └── config.go          # Database and environment configuration
│   ├── models/
│   │   ├── user.go            # User model
│   │   ├── file.go            # File model
│   │   └── shared_file.go      # Shared file model
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
├── go.mod                      # Go module definition
├── .gitignore                 # Git ignore rules
└── README.md                  # This file
```

## API Endpoints

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

## Getting Started

### Prerequisites
- Go 1.21 or later
- Node.js 16 or later
- npm or yarn

### Backend Setup

1. Install Go dependencies:
```bash
go mod download
```

2. Set environment variables:
```bash
export JWT_SECRET="your-secret-key-change-in-production"
export PORT="8080"
```

3. Run the backend server:
```bash
go run backend/main.go
```

The backend will be available at `http://localhost:8080`

### Frontend Setup

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

## Usage

1. **Register**: Create a new account with a username, email, and password
2. **Login**: Log in with your credentials
3. **Upload Files**: Click "Upload File" and select files from your computer
4. **Share Files**: Select a file and enter another user's username to share it
5. **Download Files**: Download files from your storage or shared files
6. **Delete Files**: Remove files you no longer need

## Security Considerations

- **JWT Tokens**: Tokens expire after 24 hours
- **Password Hashing**: Passwords are hashed using bcrypt
- **CORS**: Configure CORS appropriately for production
- **File Validation**: Implement file type and size validation in production
- **Environment Variables**: Change JWT_SECRET in production

## Future Enhancements

- [ ] File versioning
- [ ] Folder structure
- [ ] Advanced sharing with permissions
- [ ] File preview functionality
- [ ] File search and filtering
- [ ] User profiles and settings
- [ ] Activity logging
- [ ] Cloud storage integration (S3, Google Cloud Storage)
- [ ] Mobile app
- [ ] End-to-end encryption

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
