interface LoginPayload {
  email: string;
  password: string;
}

interface LoginResponse {
  jwt_token: string;
}
