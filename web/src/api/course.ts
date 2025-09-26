import api from './auth';

export interface Course {
  id: number;
  title: string;
  description: string;
  code: string;
  level: string;
  credits: number;
  isActive: boolean;
  schoolId: number;
  teacherId?: number;
  teacher?: User;
  enrollments?: Enrollment[];
  modules?: Module[];
  createdAt: string;
  updatedAt: string;
}

export interface Module {
  id: number;
  title: string;
  description: string;
  order: number;
  courseId: number;
  lessons?: Lesson[];
  createdAt: string;
  updatedAt: string;
}

export interface Lesson {
  id: number;
  title: string;
  content: string;
  order: number;
  moduleId: number;
  createdAt: string;
  updatedAt: string;
}

export interface Enrollment {
  id: number;
  studentId: number;
  courseId: number;
  enrolledAt: string;
  status: 'active' | 'completed' | 'dropped';
  student?: User;
}

export interface User {
  id: number;
  username: string;
  email: string;
  firstName: string;
  lastName: string;
  roles?: Role[];
}

export interface Role {
  id: number;
  name: string;
  description: string;
}

export interface CreateCourseRequest {
  title: string;
  description: string;
  code: string;
  level: string;
  credits: number;
  teacherId?: number;
}

export interface UpdateCourseRequest extends Partial<CreateCourseRequest> {
  isActive?: boolean;
}

export interface CourseQueryParams {
  page?: number;
  size?: number;
  search?: string;
  level?: string;
  teacherId?: number;
}

export interface PaginatedCourseResponse {
  data: Course[];
  pagination: {
    page: number;
    size: number;
    total: number;
    totalPages: number;
  };
}

export interface CreateModuleRequest {
  title: string;
  description: string;
  order: number;
}

export interface CreateLessonRequest {
  title: string;
  content: string;
  order: number;
}

export const courseApi = {
  // Course operations
  getCourses: async (params: CourseQueryParams = {}) => {
    const response = await api.get<PaginatedCourseResponse>('/courses', { params });
    return response.data;
  },

  getCourse: async (id: number) => {
    const response = await api.get<Course>(`/courses/${id}`);
    return response.data;
  },

  createCourse: async (data: CreateCourseRequest) => {
    const response = await api.post<Course>('/courses', data);
    return response.data;
  },

  updateCourse: async (id: number, data: UpdateCourseRequest) => {
    const response = await api.put<Course>(`/courses/${id}`, data);
    return response.data;
  },

  deleteCourse: async (id: number) => {
    await api.delete(`/courses/${id}`);
  },

  // Module operations
  getModules: async (courseId: number) => {
    const response = await api.get<Module[]>(`/courses/${courseId}/modules`);
    return response.data;
  },

  createModule: async (courseId: number, data: CreateModuleRequest) => {
    const response = await api.post<Module>(`/courses/${courseId}/modules`, data);
    return response.data;
  },

  updateModule: async (courseId: number, moduleId: number, data: Partial<CreateModuleRequest>) => {
    const response = await api.put<Module>(`/courses/${courseId}/modules/${moduleId}`, data);
    return response.data;
  },

  deleteModule: async (courseId: number, moduleId: number) => {
    await api.delete(`/courses/${courseId}/modules/${moduleId}`);
  },

  // Lesson operations
  getLessons: async (courseId: number, moduleId: number) => {
    const response = await api.get<Lesson[]>(`/courses/${courseId}/modules/${moduleId}/lessons`);
    return response.data;
  },

  createLesson: async (courseId: number, moduleId: number, data: CreateLessonRequest) => {
    const response = await api.post<Lesson>(`/courses/${courseId}/modules/${moduleId}/lessons`, data);
    return response.data;
  },

  updateLesson: async (courseId: number, moduleId: number, lessonId: number, data: Partial<CreateLessonRequest>) => {
    const response = await api.put<Lesson>(`/courses/${courseId}/modules/${moduleId}/lessons/${lessonId}`, data);
    return response.data;
  },

  deleteLesson: async (courseId: number, moduleId: number, lessonId: number) => {
    await api.delete(`/courses/${courseId}/modules/${moduleId}/lessons/${lessonId}`);
  },

  // Enrollment operations
  getEnrollments: async (courseId: number) => {
    const response = await api.get<Enrollment[]>(`/courses/${courseId}/enrollments`);
    return response.data;
  },

  enrollStudent: async (courseId: number, studentId: number) => {
    const response = await api.post<Enrollment>(`/courses/${courseId}/enrollments`, { studentId });
    return response.data;
  },

  updateEnrollment: async (courseId: number, enrollmentId: number, status: 'active' | 'completed' | 'dropped') => {
    const response = await api.put<Enrollment>(`/courses/${courseId}/enrollments/${enrollmentId}`, { status });
    return response.data;
  },

  // Get available teachers
  getTeachers: async () => {
    const response = await api.get<User[]>('/users/teachers');
    return response.data;
  },

  // Get available students
  getStudents: async () => {
    const response = await api.get<User[]>('/users/students');
    return response.data;
  },
};