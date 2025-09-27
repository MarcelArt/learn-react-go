import React, { createContext, useContext, useEffect, useState, useRef, useCallback } from 'react';
import { useNavigate } from '@tanstack/react-router';
import { authApi } from '@/api/auth';
import type { LoginInput, RegisterSchoolInput, LoginResponse } from '@/types/auth';

interface AuthContextType {
  user: any;
  login: (data: LoginInput) => Promise<void>;
  registerSchool: (data: RegisterSchoolInput) => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
  isLoading: boolean;
  setNavigationCallback: (callback: () => void) => void;
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
  const navigationCallbackRef = useRef<(() => void) | null>(null);

  useEffect(() => {
    // Check if user is authenticated on app load
    const token = localStorage.getItem('accessToken');
    if (token) {
      // You could validate the token here or fetch user data
      setUser({ isAuthenticated: true });
    }
    setIsLoading(false);
  }, []);

  const setNavigationCallback = useCallback((callback: () => void) => {
    navigationCallbackRef.current = callback;
  }, []);

  const triggerNavigation = useCallback(() => {
    if (navigationCallbackRef.current) {
      navigationCallbackRef.current();
    }
  }, []);

  const login = async (data: LoginInput) => {
    try {
      const response: LoginResponse = await authApi.login(data);
      localStorage.setItem('accessToken', response.accessToken);
      localStorage.setItem('refreshToken', response.refreshToken);
      setUser({ isAuthenticated: true });

      // Trigger navigation after successful login
      setTimeout(triggerNavigation, 0); // Use setTimeout to avoid state updates during render
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

      // Trigger navigation after successful registration
      setTimeout(triggerNavigation, 0); // Use setTimeout to avoid state updates during render
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
    setNavigationCallback,
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
}

// Navigation wrapper component
export function AuthNavigationWrapper({ children }: { children: React.ReactNode }) {
  const { setNavigationCallback } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    // Set up navigation handler - this will only run once when dependencies are stable
    const handleAuthSuccess = () => {
      navigate({ to: '/dashboard' });
    };

    setNavigationCallback(handleAuthSuccess);
  }, [navigate, setNavigationCallback]);

  return <>{children}</>;
}