import { Outlet, createRootRoute, useLocation } from '@tanstack/react-router'
import { TanStackRouterDevtoolsPanel } from '@tanstack/react-router-devtools'
import { TanstackDevtools } from '@tanstack/react-devtools'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { MainLayout } from '../components/layout/MainLayout'
import { AuthProvider, AuthNavigationWrapper } from '../contexts/AuthContext'

const queryClient = new QueryClient()

function RootComponent() {
  const location = useLocation()
  const isLoginPage = location.pathname === '/login'

  return (
    <>
      {isLoginPage ? (
        <Outlet />
      ) : (
        <MainLayout>
          <Outlet />
        </MainLayout>
      )}
      <TanstackDevtools
        config={{
          position: 'bottom-left',
        }}
        plugins={[
          {
            name: 'Tanstack Router',
            render: <TanStackRouterDevtoolsPanel />,
          },
        ]}
      />
    </>
  )
}

export const Route = createRootRoute({
  component: () => (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <AuthNavigationWrapper>
          <RootComponent />
        </AuthNavigationWrapper>
      </AuthProvider>
    </QueryClientProvider>
  ),
})
