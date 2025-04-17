import { json } from '@tanstack/react-start'
import { createAPIFileRoute } from '@tanstack/react-start/api'
import { LMS_BACKEND_URL } from '@/utils/env'
import { IResponse, UserLoginResponse } from '@/utils/user'

export const APIRoute = createAPIFileRoute('/api/user/login')({
  POST: async ({ request }) => {
    try {
      const body = await request.json();
      const res = await fetch(LMS_BACKEND_URL + '/api/v1/user/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(body),
        credentials: "include",
      })

      const setCookie = res.headers.get('set-cookie')
  
      const data = await res.json() as IResponse<UserLoginResponse>
      return json(data, { 
        status: 200,
        headers: setCookie ? { 'set-cookie': setCookie } : undefined,
      })
    } catch (e) {
      console.error("Exception caught:", e);
      return json({ error: 'An error occurred during login' }, { status: 500 });
    }
  }
})  