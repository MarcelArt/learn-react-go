export interface UserDTO {
  id: number;
  username: string;
  email: string;
  password?: string;
  first_name?: string;
  last_name?: string;
  phone?: string;
  is_active: boolean;
  school_id?: number;
}

export interface SchoolDTO {
  name: string;
  email: string;
  phone?: string;
  address?: string;
  website?: string;
  description?: string;
  logo_url?: string;
  is_active: boolean;
}

export interface RegisterSchoolInput {
  userData: UserDTO;
  schoolData: SchoolDTO;
}

export interface LoginInput {
  username: string;
  password: string;
  isRemember?: boolean;
}

export interface RefreshInput {
  refreshToken: string;
}

export interface LoginResponse {
  accessToken: string;
  refreshToken: string;
}

export interface User {
  id: number;
  username: string;
  email: string;
  first_name?: string;
  last_name?: string;
  phone?: string;
  is_active: boolean;
  school_id?: number;
  created_at: string;
  updated_at: string;
}