import { api } from '@/utils/api';
import { createFileRoute } from '@tanstack/react-router'
import { useEffect, useState } from 'react';

export const Route = createFileRoute('/submissions')({
  component: RouteComponent,
})

function RouteComponent() {
  const [submissions, setSubmissions] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchSubmissions = async () => {
      try {
        const data = await api.getSubmissions(undefined);
        setSubmissions(data);
      } catch (error) {
        console.error('Failed to fetch submissions:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchSubmissions();
  }, []);  
  if (loading) return <div className="p-4">Loading...</div>;
  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-2xl font-bold mb-6">Submissions</h1>
      <div className="grid gap-6">
        {submissions.map((submission: any) => (
          <div key={submission.id} className="bg-white p-6 rounded-lg shadow">
            <div className="flex justify-between items-start mb-4">
              <h2 className="text-xl font-semibold">Submission #{submission.id}</h2>
              <span className={`px-3 py-1 rounded text-sm ${
                submission.status === 'graded' ? 'bg-green-100 text-green-800' : 'bg-yellow-100 text-yellow-800'
              }`}>
                {submission.status}
              </span>
            </div>
            <p className="text-gray-600 mb-4">{submission.content}</p>
            <div className="text-sm text-gray-500">
              Submitted: {submission.submittedAt}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
