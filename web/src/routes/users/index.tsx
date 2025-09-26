import { createFileRoute } from '@tanstack/react-router'
import { UserList } from '../../components/users/UserList'

export const Route = createFileRoute('/users/')({
  component: UsersComponent,
})

function UsersComponent() {
  // Get school ID from JWT token
  const schoolId = 1; // This would come from the user's token

  return <UserList schoolId={schoolId} />
}