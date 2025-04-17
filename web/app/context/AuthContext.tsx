
import React, { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { authAPI, User } from '@/lib/api';
import { useToast } from '@/components/ui/use-toast';

type AuthContextType = {
  user: User | null;
  isLoading: boolean;
  isAuthenticated: boolean;
  login: (email: string, password: string) => Promise<void>;
  register: (name: string, email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
};

const AuthContext = createContext<AuthContextType | null>(null);

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const { toast } = useToast();

  const loadUser = async () => {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      try {
        setUser(JSON.parse(storedUser));
        // Validate the token by making a request
        const { data } = await authAPI.getMe();
        setUser(data.user);
        localStorage.setItem('user', JSON.stringify(data.user));
      } catch (error) {
        console.error('Failed to load user:', error);
        localStorage.removeItem('user');
        localStorage.removeItem('token');
        setUser(null);
      }
    }
    setIsLoading(false);
  };

  useEffect(() => {
    loadUser();
  }, []);

  const login = async (email: string, password: string) => {
    try {
      const response = await authAPI.login({ email, password });
      const { token, user } = response.data;
      
      localStorage.setItem('token', token);
      localStorage.setItem('user', JSON.stringify(user));
      
      setUser(user);
      toast({
        title: 'Login Successful',
        description: `Welcome back, ${user.name}!`,
        variant: 'default',
      });
    } catch (error) {
      console.error('Login error:', error);
      toast({
        title: 'Login Failed',
        description: 'Invalid email or password.',
        variant: 'destructive',
      });
      throw error;
    }
  };

  const register = async (name: string, email: string, password: string) => {
    try {
      const response = await authAPI.register({ name, email, password });
      const { token, user } = response.data;
      
      localStorage.setItem('token', token);
      localStorage.setItem('user', JSON.stringify(user));
      
      setUser(user);
      toast({
        title: 'Registration Successful',
        description: `Welcome, ${user.name}!`,
        variant: 'default',
      });
    } catch (error) {
      console.error('Registration error:', error);
      toast({
        title: 'Registration Failed',
        description: 'Could not create account. Please try again.',
        variant: 'destructive',
      });
      throw error;
    }
  };

  const logout = async () => {
    try {
      await authAPI.logout();
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      setUser(null);
      toast({
        title: 'Logged Out',
        description: 'You have been successfully logged out.',
      });
    } catch (error) {
      console.error('Logout error:', error);
      // Still clear local storage even if API call fails
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      setUser(null);
    }
  };

  return (
    <AuthContext.Provider
      value={{
        user,
        isLoading,
        isAuthenticated: !!user,
        login,
        register,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};