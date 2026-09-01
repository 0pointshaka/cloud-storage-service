import React, { useState, useEffect } from 'react';
import axios from 'axios';
import '../styles/DashboardPage.css';

function DashboardPage({ onLogout }) {
  const [files, setFiles] = useState([]);
  const [sharedFiles, setSharedFiles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState('myfiles');
  const [shareUsername, setShareUsername] = useState('');
  const [selectedFileId, setSelectedFileId] = useState(null);

  const API_BASE_URL = 'http://localhost:8080/api';
  const token = localStorage.getItem('token');
  const headers = { Authorization: `Bearer ${token}` };

  useEffect(() => {
    fetchFiles();
    fetchSharedFiles();
  }, []);

  const fetchFiles = async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/files/list`, { headers });
      setFiles(response.data || []);
    } catch (err) {
      console.error('Failed to fetch files:', err);
    } finally {
      setLoading(false);
    }
  };

  const fetchSharedFiles = async () => {
    try {
      const response = await axios.get(`${API_BASE_URL}/share/list`, { headers });
      setSharedFiles(response.data || []);
    } catch (err) {
      console.error('Failed to fetch shared files:', err);
    }
  };

  const handleUpload = async (e) => {
    const file = e.target.files[0];
    if (!file) return;

    const formData = new FormData();
    formData.append('file', file);

    try {
      await axios.post(`${API_BASE_URL}/files/upload`, formData, { headers });
      fetchFiles();
      alert('File uploaded successfully');
    } catch (err) {
      alert('Failed to upload file');
    }
  };

  const handleDeleteFile = async (fileId) => {
    if (window.confirm('Are you sure?')) {
      try {
        await axios.delete(`${API_BASE_URL}/files/${fileId}`, { headers });
        fetchFiles();
        alert('File deleted');
      } catch (err) {
        alert('Failed to delete file');
      }
    }
  };

  const handleShareFile = async () => {
    if (!selectedFileId || !shareUsername) {
      alert('Please select file and enter username');
      return;
    }

    try {
      await axios.post(
        `${API_BASE_URL}/share/${selectedFileId}`,
        { username: shareUsername },
        { headers }
      );
      setShareUsername('');
      setSelectedFileId(null);
      alert('File shared successfully');
    } catch (err) {
      alert('Failed to share file');
    }
  };

  const handleDownloadFile = (fileId, filename) => {
    const link = document.createElement('a');
    link.href = `${API_BASE_URL}/files/download/${fileId}`;
    link.setAttribute('Authorization', `Bearer ${token}`);
    link.download = filename;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  };

  if (loading) return <div>Loading...</div>;

  return (
    <div className="dashboard">
      <header className="dashboard-header">
        <h1>Cloud Storage</h1>
        <button onClick={onLogout} className="logout-button">Logout</button>
      </header>

      <div className="dashboard-content">
        <div className="tabs">
          <button
            className={`tab ${activeTab === 'myfiles' ? 'active' : ''}`}
            onClick={() => setActiveTab('myfiles')}
          >
            My Files
          </button>
          <button
            className={`tab ${activeTab === 'shared' ? 'active' : ''}`}
            onClick={() => setActiveTab('shared')}
          >
            Shared with Me
          </button>
        </div>

        {activeTab === 'myfiles' && (
          <div className="tab-content">
            <div className="upload-section">
              <input
                type="file"
                id="file-input"
                onChange={handleUpload}
                style={{ display: 'none' }}
              />
              <button
                onClick={() => document.getElementById('file-input').click()}
                className="upload-button"
              >
                Upload File
              </button>
            </div>

            <div className="share-section">
              <h3>Share File</h3>
              <select
                value={selectedFileId}
                onChange={(e) => setSelectedFileId(e.target.value)}
              >
                <option value="">Select a file</option>
                {files.map((file) => (
                  <option key={file.id} value={file.id}>
                    {file.filename}
                  </option>
                ))}
              </select>
              <input
                type="text"
                placeholder="Username"
                value={shareUsername}
                onChange={(e) => setShareUsername(e.target.value)}
              />
              <button onClick={handleShareFile} className="share-button">
                Share
              </button>
            </div>

            <div className="files-list">
              <h3>Files</h3>
              <table>
                <thead>
                  <tr>
                    <th>Filename</th>
                    <th>Size</th>
                    <th>Date</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {files.map((file) => (
                    <tr key={file.id}>
                      <td>{file.filename}</td>
                      <td>{(file.size / 1024).toFixed(2)} KB</td>
                      <td>{new Date(file.created_at * 1000).toLocaleDateString()}</td>
                      <td>
                        <button
                          onClick={() => handleDownloadFile(file.id, file.filename)}
                          className="action-button"
                        >
                          Download
                        </button>
                        <button
                          onClick={() => handleDeleteFile(file.id)}
                          className="action-button delete"
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        )}

        {activeTab === 'shared' && (
          <div className="tab-content">
            <div className="files-list">
              <h3>Shared Files</h3>
              <table>
                <thead>
                  <tr>
                    <th>Filename</th>
                    <th>Size</th>
                    <th>Shared By</th>
                    <th>Date</th>
                  </tr>
                </thead>
                <tbody>
                  {sharedFiles.map((file) => (
                    <tr key={file.file_id}>
                      <td>{file.filename}</td>
                      <td>{(file.size / 1024).toFixed(2)} KB</td>
                      <td>User {file.shared_by}</td>
                      <td>{new Date(file.created_at * 1000).toLocaleDateString()}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default DashboardPage;
