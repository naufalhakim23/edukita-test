export interface JwtPayload {
  uuid: string;
  email: string;
  name: string;
  picture: string;
  role: string;
  created_at: string;
  updated_at: string;
  exp: number;
  iat: number;
  [key: string]: any;
}