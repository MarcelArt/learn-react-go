# Learn React Go (Project Name Might be Changed)
This is a Learning Management System (LMS) app that can be use for Elementary School to College

## Project Architecture
On the root of the project is all of the backend written in Golang while inside the `./web` directory is the frontend application written using React

### Backend
The backend is written Golang with Fiber as the web framework, using Gorm as the ORM to communicate to the database, and is using PostgreSQL as the database.
The project is using the repository pattern where the repository layer communicate to the database and the handler layer calls the repository methods, then return a json response.
The login are using username and password, which will return refresh token and access token pair
- Access token: Lives for 5 minutes and will need to be refreshed using refresh token
- Refresh token: Lives for 1 day to 30 days depending on whether user want the app to remember the login

### Frontend
Stack used on the frontend side are as follows:
- React
- Typescript
- Tailwind
- Shadcn: As the UI Component
- Tanstack Router: To handle routing on the frontend side
- Axios: To handle HTTP request to the backend
- Tanstack Query

## Features
### User Management
It will need to have role permissions, each user can have any number of roles, and each role can have any number of permissions.
So there could be for example:
- Teacher: Can give assignment, give score, and manage course of where they are assigned to (they cannot view or write of another teacher assigned courses)
- Student: Can see all their assigned courses, see their scores but not other students
- Parent: Can see all their children assigned courses and see their score but cannot submit an assignment
- And the role could be customizable since each school can have their set of rules
- The application also would be multi tenancy, where one user register to the platform will become an Admin or Super User for that school, and after registering they need to create a school data first, like school name, school email, school phone, etc.

### Course Management
User can manage courses, classes, assign teacher, assign students to classes/courses.
The database need to be meaningfully record:
- School year (e.g., 2024/2025, 2025/2026, ...etc)
- Semester of the current School Year (1, 2)
- The courses and the relation of the current school year and current semester, the teacher, and the participating students

### Online Learning Tools
- Teacher can upload document (pdf, docx, xlsx, pptx, png, jpg, mp4, ...etc)
- Teacher can give assignment to all course participants
    - The assignment can be upload file with deadline
    - Or a quiz, with deadline and can be auto scored

### Assesment 
- Teacher can give student submitted assignment a score
- Teacher and student can see the report of student performance of each semester

### More will be added
Current this is a draft more features will be added