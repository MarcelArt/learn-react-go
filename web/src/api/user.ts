import api from './auth';
import type { User, UserDTO } from '../types/auth';

export interface PaginatedResponse<T> {
  items: T[];
  page: number;
  size: number;
  max_page: number;
  total_pages: number;
  total: number;
  last: boolean;
  first: boolean;
  visible: number;
}

export interface UserQueryParams {
  page?: number;
  size?: number;
  sort?: string;
  filters?: string;
}

export const userApi = {
  // Get all users with pagination
  getUsers: async (params?: UserQueryParams): Promise<PaginatedResponse<User>> => {
    const response = await api.get('/user', { params });
    return response.data;
  },

  // Get user by ID
  getUserById: async (id: string): Promise<User> => {
    const response = await api.get(`/user/${id}`);
    return response.data.items;
  },

  // Create new user
  createUser: async (data: UserDTO): Promise<{ ID: number }> => {
    const response = await api.post('/user', data);
    return response.data.items;
  },

  // Update user
  updateUser: async (id: string, data: Partial<UserDTO>): Promise<void> => {
    await api.put(`/user/${id}`, data);
  },

  // Delete user
  deleteUser: async (id: string): Promise<User> => {
    const response = await api.delete(`/user/${id}`);
    return response.data.items;
  },
};