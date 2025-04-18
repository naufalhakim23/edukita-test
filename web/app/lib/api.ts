import { LMS_BACKEND_URL } from '@/utils/env';
import axios from 'axios';

// Create an axios instance with default configs
const api = axios.create({
  baseURL: LMS_BACKEND_URL + '/api/v1',
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true, // Important for cookies
});

// Request interceptor to add auth token
api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor for error handling
api.interceptors.response.use(
  (response) => response,
  (error) => {
    // Handle 401 Unauthorized errors
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

// Auth APIs
export const authAPI = {
  register: (userData: RegisterUserData) => api.post('/user/register', userData),
  login: (credentials: LoginCredentials) => api.post('/user/login', credentials),
  logout: () => api.post('/user/logout'),
  getMe: () => api.get('/user/me'),
  getUserById: (id: string) => api.get(`/user/${id}`),
};

// LMS APIs
export const lmsAPI = {
  // Assignment APIs
  createAssignment: (data: CreateAssignmentData) => 
    api.post('/lms/assignments', data),
  getAssignmentById: (id: string) => 
    api.get(`/lms/assignments/${id}`),
  updateAssignment: (id: string, data: UpdateAssignmentData) => 
    api.put(`/lms/assignments/${id}`, data),
    
  // Submission APIs
  createSubmission: (data: CreateSubmissionData) => 
    api.post('/lms/submissions', data),
  getSubmissionById: (id: string) => 
    api.get(`/lms/submissions/${id}`),
  updateSubmission: (id: string, data: UpdateSubmissionData) => 
    api.put(`/lms/submissions/${id}`, data),
  getSubmissionsByAssignment: (id: string) => 
    api.get(`/lms/submissions/assignments/${id}`),
  getSubmissionsByUser: (id: string) => 
    api.get(`/lms/submissions/users/${id}`),
};

// Types
export interface RegisterUserData {
  name: string;
  email: string;
  password: string;
  role?: string;
}

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface User {
  id: string;
  name: string;
  email: string;
  role: string;
  createdAt: string;
}

export interface CreateAssignmentData {
  title: string;
  description: string;
  dueDate: string;
  totalPoints: number;
}

export interface UpdateAssignmentData {
  title?: string;
  description?: string;
  dueDate?: string;
  totalPoints?: number;
}

export interface Assignment {
  id: string;
  title: string;
  description: string;
  dueDate: string;
  totalPoints: number;
  createdBy: string;
  createdAt: string;
  updatedAt: string;
}

export interface CreateSubmissionData {
  assignmentId: string;
  content: string;
}

export interface UpdateSubmissionData {
  content?: string;
  grade?: number;
  feedback?: string;
}

export interface Submission {
  id: string;
  assignmentId: string;
  userId: string;
  content: string;
  grade?: number;
  feedback?: string;
  submittedAt: string;
  updatedAt: string;
}

export default api;
