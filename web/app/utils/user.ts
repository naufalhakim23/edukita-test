type User = {
  id: string
  first_name: string
  last_name: string
  email: string
  role: string
  is_active: boolean
  last_login: string
  created_at: string
  updated_at: string
}

type RegisterUserRequest = {
  name: string
  email: string
  password: string
  first_name: string
  last_name: string
  role: string
}

type RegisterUserResponse = {
  id: string
  first_name: string
  last_name: string
  email: string
  role: string
}

type UserLoginRequest = {
  email: string
  password: string
}

type UserLoginResponse = {
  user: User
  token: string
}