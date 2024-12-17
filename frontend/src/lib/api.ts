const API_BASE = '/api';

interface AuthResponse {
  success: boolean;
  message?: string;
}

export async function login(email: string, password: string): Promise<AuthResponse> {
  const res = await fetch(`${API_BASE}/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password }),
  });

  return res.ok ? { success: true } : { success: false, message: await res.text() };
}

export async function register(name: string, email: string, password: string): Promise<AuthResponse> {
  const res = await fetch(`${API_BASE}/register`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ name, email, password }),
  });

  return res.ok ? { success: true } : { success: false, message: await res.text() };
}