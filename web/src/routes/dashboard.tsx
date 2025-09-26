import { createFileRoute } from '@tanstack/react-router'
import { Dashboard } from '../components/Dashboard'

export const Route = createFileRoute('/dashboard')({
  component: DashboardComponent,
})

function DashboardComponent() {
  // Get school name from localStorage or JWT token
  const schoolName = 'Test School'; // This would come from the user's school data

  return <Dashboard schoolName={schoolName} />
}