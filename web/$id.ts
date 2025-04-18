import { json } from '@tanstack/react-start'
import { createAPIFileRoute } from '@tanstack/react-start/api'
import { LMS_BACKEND_URL } from '@/utils/env'
import { User } from '@/utils/user'

// export const APIRoute = createAPIFileRoute('/api/user/$id')({
//   GET: async ({ request, params }) => {
//     console.info(`Fetching users by id=${params.id}... @`, request.url)
//     try {
//       const res = await fetch(
//         LMS_BACKEND_URL + '/api/v1/users/' + params.id,
//         {
//             method: 'GET',
//             headers: {
//                 'Content-Type': 'application/json',
//             },
//             credentials: 'include',
//         }
//       )
//       if (!res.ok) {
//         throw new Error('Failed to fetch user')
//       }

//       const user = (await res.json()) as User

//       return json({
//         id: user.id,
//         name: user.first_name,
//         email: user.email,
//         role: user.role,
//         is_active: user.is_active,
//         last_login: user.last_login,
//         created_at: user.created_at,
//         updated_at: user.updated_at,
//       })
//     } catch (e) {
//       console.error(e)
//       return json({ error: 'User not found' }, { status: 404 })
//     }
//   },
// })
