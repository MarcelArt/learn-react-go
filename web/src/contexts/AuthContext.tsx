import React, { createContext, useContext, useEffect, useState } from 'react';
import { authApi } from '@/api/auth';
import type { LoginInput, RegisterSchoolInput, LoginResponse } from '@/types/auth';

interface AuthContextType {
  user: any;
  login: (data: LoginInput) => Promise<void>;
  registerSchool: (data: RegisterSchoolInput) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
  isLoading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}

interface AuthProviderProps {
  children: React.ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [user, setUser] = useState<any>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    // Check if user is authenticated on app load
    const token = localStorage.getItem('accessToken');
    if (token) {
      // You could validate the token here or fetch user data
      setUser({ isAuthenticated: true });
    }
    setIsLoading(false);
  }, []);

  const login = async (data: LoginInput) => {
    try {
      const response: LoginResponse = await authApi.login(data);
      localStorage.setItem('accessToken', response.accessToken);
      localStorage.setItem('refreshToken', response.refreshToken);
      setUser({ isAuthenticated: true });
    } catch (error) {
      throw error;
    }
  };

  const registerSchool = async (data: RegisterSchoolInput) => {
    try {
      const response: LoginResponse = await authApi.registerSchool(data);
      localStorage.setItem('accessToken', response.accessToken);
      localStorage.setItem('refreshToken', response.refreshToken);
      setUser({ isAuthenticated: true });
    } catch (error) {
      throw error;
    }
  };

  const logout = () => {
    authApi.logout();
    setUser(null);
  };

  const value = {
    user,
    login,
    registerSchool,
    logout,
    isAuthenticated: !!user,
    isLoading,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}