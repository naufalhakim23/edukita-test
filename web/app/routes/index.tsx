import { ProtectedRoute } from '@/context/AuthContext'
import { createFileRoute, Link, useRouter } from '@tanstack/react-router'

export const Route = createFileRoute('/')({
  component: () => (
      <ProtectedRoute>
        <Home />
      </ProtectedRoute>
  ),
})

function Home() {
  return (
    <div className="min-h-screen bg-gray-100">
      <div className="container mx-auto px-4 py-8">
        <h1 className="text-3xl font-bold mb-6">Welcome to LMS Dashboard</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div className="bg-white p-6 rounded-lg shadow">
            <h2 className="text-xl font-semibold mb-2">Assignments</h2>
            <p className="text-gray-600">View and manage your assignments</p>
            <Link to="/assignments" className="mt-4 inline-block text-blue-600 hover:text-blue-800">View Assignments →</Link>
          </div>
          <div className="bg-white p-6 rounded-lg shadow">
            <h2 className="text-xl font-semibold mb-2">Submissions</h2>
            <p className="text-gray-600">Track your submission progress</p>
            <Link to="/submissions" className="mt-4 inline-block text-blue-600 hover:text-blue-800">View Submissions →</Link>
          </div>
          <div className="bg-white p-6 rounded-lg shadow">
            <h2 className="text-xl font-semibold mb-2">Profile</h2>
            <p className="text-gray-600">Manage your account settings</p>
            <Link to="/profile" className="mt-4 inline-block text-blue-600 hover:text-blue-800">View Profile →</Link>
          </div>
        </div>
      </div>
    </div>
  )
}