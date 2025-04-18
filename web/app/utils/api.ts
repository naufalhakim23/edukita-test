
import { IResponse, User, RegisterUserRequest, RegisterUserResponse, UserLoginRequest, UserLoginResponse } from '@/types/user';

const API_URL = 'http://localhost:8080/api/v1';

const getToken = () => {
  return localStorage.getItem('token');
};

const getHeaders = () => {
  const token = getToken();
  const headers: HeadersInit = {
    'Content-Type': 'application/json',
  };
  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }
  return headers;
};

export const api = {
  login: async (credentials: UserLoginRequest): Promise<IResponse<UserLoginResponse>> => {
    const response = await fetch(`${API_URL}/user/login`, {
      method: 'POST',
      headers: getHeaders(),
      body: JSON.stringify(credentials),
    });
    return response.json();
  },

  register: async (userData: RegisterUserRequest): Promise<IResponse<RegisterUserResponse>> => {
    const response = await fetch(`${API_URL}/user/register`, {
      method: 'POST',
      headers: getHeaders(),
      body: JSON.stringify(userData),
    });
    return response.json();
  },

  logout: async (): Promise<IResponse<void>> => {
    const response = await fetch(`${API_URL}/user/logout`, {
      method: 'POST',
      headers: getHeaders(),
    });
    return response.json();
  },

  getMe: async (): Promise<IResponse<User>> => {
    const response = await fetch(`${API_URL}/user/me`, {
      headers: getHeaders(),
    });
    return response.json();
  },

  // Assignments endpoints
  getAssignments: async () => {
    const response = await fetch(`${API_URL}/lms/assignments`, {
      headers: getHeaders(),
    });
    return response.json();
  },

  getAssignment: async (id: number) => {
    const response = await fetch(`${API_URL}/lms/assignments/${id}`, {
      headers: getHeaders(),
    });
    return response.json();
  },

  createAssignment: async (data: any) => {
    const response = await fetch(`${API_URL}/lms/assignments`, {
      method: 'POST',
      headers: getHeaders(),
      body: JSON.stringify(data),
    });
    return response.json();
  },

  // Submissions endpoints
  getSubmissions: async (assignmentId?: number) => {
    const url = assignmentId 
      ? `${API_URL}/lms/submissions/assignments/${assignmentId}`
      : `${API_URL}/lms/submissions`;
    const response = await fetch(url, {
      headers: getHeaders(),
    });
    return response.json();
  },

  createSubmission: async (data: any) => {
    const response = await fetch(`${API_URL}/lms/submissions`, {
      method: 'POST',
      headers: getHeaders(),
      body: JSON.stringify(data),
    });
    return response.json();
  },
};
