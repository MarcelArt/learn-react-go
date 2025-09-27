import { createFileRoute } from '@tanstack/react-router'
import { Dashboard } from '../components/Dashboard'
import { AuthGuard } from '../components/AuthGuard'

export const Route = createFileRoute('/dashboard')({
  component: () => (
    <AuthGuard>
      <DashboardComponent />
    </AuthGuard>
  ),
})

function DashboardComponent() {
  // Get school name from localStorage or JWT token
  const schoolName = 'Test School'; // This would come from the user's school data

  return <Dashboard schoolName={schoolName} />
}