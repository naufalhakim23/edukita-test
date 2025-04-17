
export interface IResponse<T = unknown> {
	data: T;
	message: string;
	status: number;
  }
  
export type User = {
	id: string;
	first_name: string;
	last_name: string;
	email: string;
	role: string;
	is_active: boolean;
	last_login: string;
	created_at: string;
	updated_at: string;
}
  
export type RegisterUserRequest = {
	name: string;
	email: string;
	password: string;
	confirm_password: string;
	first_name: string;
	last_name: string;
	role: string;
}
  
export type RegisterUserResponse = {
	id: string;
	first_name: string;
	last_name: string;
	email: string;
	role: string;
}
  
export type UserLoginRequest = {
	email: string;
	password: string;
}
  
export type UserLoginResponse = {
	user: User;
	token: string;
}
  