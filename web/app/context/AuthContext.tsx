import React, { createContext, useContext, useState, useEffect } from 'react';
import { User } from '@/types/user';
import { Navigate, useLocation } from '@tanstack/react-router';
import { COOKIE_NAME } from '@/utils/env';
import { JwtPayload } from '@/types/jwt';
import { jwtDecode } from 'jwt-decode';

interface AuthContextType {
  user: User | null;
  token: string | null;
  login: (token: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | null>(null);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null);

  useEffect(() => {
    const savedUser = localStorage.getItem('user');
    const savedToken = localStorage.getItem('token');
    if (savedUser && savedToken) {
      setUser(JSON.parse(savedUser));
      setToken(savedToken);
    } else {
      const cookieToken = getCookie(COOKIE_NAME);
      if (cookieToken) {
        setToken(cookieToken);
        localStorage.setItem('token', cookieToken);

        const decoded = jwtDecode<JwtPayload>(cookieToken);
        const user: User = {
          id: decoded.uuid,
          first_name: decoded.name,
          last_name: decoded.name,
          email: decoded.email,
          role: decoded.role,
          is_active: true,
          last_login: decoded.updated_at,
          created_at: decoded.created_at,
          updated_at: decoded.updated_at,
        };
        setUser(user);
        localStorage.setItem('user', JSON.stringify(user));
      }
    }
  }, []);

  const login = (token: string) => {
    setUser(user);
    setToken(token);
    localStorage.setItem('user', JSON.stringify(user));
    localStorage.setItem('token', token);
  };

  const logout = () => {
    setUser(null);
    setToken(null);
    localStorage.removeItem('user');
    localStorage.removeItem('token');
  };

  return (
    <AuthContext.Provider value={{ user, token, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const ProtectedRoute = ({ children }: { children: React.ReactNode }) => {
  const { user } = useAuth();
  const location = useLocation();
  if (!user) {
    return <Navigate to="/auth" replace />;
  }

  return <>{children}</>;
};

const getCookie = (name: string): string | null => {
  const match = document.cookie.match(new RegExp('(^| )' + name + '=([^;]+)'));
  return match ? decodeURIComponent(match[2]) : null;
};


