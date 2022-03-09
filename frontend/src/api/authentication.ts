import { http } from "./axios";
import { ENV } from "../env";

if (ENV.MOCKED) {
    require('./authentication.mock');
}

export interface LoginResponse {
    token: string;
    refresh_token: string;
    user: {
        id: number;
        username: string;
        created_at: Date;
        updated_at: Date;
    };
}

export async function Login(username: string, password: string) {
    return (await http.post<LoginResponse>('/authentication/login', { username, password })).data;
}