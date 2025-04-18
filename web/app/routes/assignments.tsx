import { api } from '@/utils/api';
import { createFileRoute } from '@tanstack/react-router'
import { useEffect, useState } from 'react';

export const Route = createFileRoute('/assignments')({
  component: RouteComponent,
})

function RouteComponent() {
  const [assignments, setAssignments] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchAssignments = async () => {
      try {
        const data = await api.getAssignments();
        setAssignments(data);
      } catch (error) {
        console.error('Failed to fetch assignments:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchAssignments();
  }, []);

  if (loading) return <div className="p-4">Loading...</div>;
  return <div>Hello "/assignments"!</div>
}
