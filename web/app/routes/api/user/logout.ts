import { json } from '@tanstack/react-start'
import { createAPIFileRoute } from '@tanstack/react-start/api'
import { LMS_BACKEND_URL } from '@/utils/env'
import { UserLoginResponse } from '@/utils/user'

export const APIRoute = createAPIFileRoute('/api/user/logout')({
  POST: async ({ request }) => {
    try{
      const body = await request.json()
      const res = await fetch(LMS_BACKEND_URL + '/api/v1/user/logout', {
          method: 'POST',
          headers: {
              'Content-Type': 'application/json',
          },
          body: JSON.stringify(body),
          credentials: 'include',
      })
      if (!res.ok) {
        throw new Error('Failed to logout user')
      }
  
      const data = (await res.json()) as UserLoginResponse
  
      return json(data)
    }
    catch(e){
      console.error(e)
      return json({ error: 'User not found' }, { status: 404 })
    }
  },
})