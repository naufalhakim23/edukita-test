import { json } from '@tanstack/react-start'
import { createAPIFileRoute } from '@tanstack/react-start/api'
import { LMS_BACKEND_URL } from '@/utils/env'
import { RegisterUserResponse } from '@/utils/user'

export const APIRoute = createAPIFileRoute('/api/user/register')({
  POST: async ({ request }) => {
    try{
        const body = await request.json()
        console.info('Creating user with data:', body)
        const res = await fetch(LMS_BACKEND_URL + '/api/v1/user/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(body)
        })
        if (!res.ok) {
          throw new Error('Failed to create user')
        }
    
        const data = (await res.json()) as RegisterUserResponse
        return json(data)
    } catch(e){
      console.error(e)
      return json({ error: 'User not found' }, { status: 404 })
    }
  },
})