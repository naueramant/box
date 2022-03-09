import { http } from "./axios";
import MockAdapter from 'axios-mock-adapter';

const mock = new MockAdapter(http);

mock.onPost("/api/authentication/login").reply(200, {
    token: "123456789",
    refresh_token: "987654321",
    user: {
        id: 1,
        username: "john",
        created_at: Date.now(),
        updated_at: Date.now()
    }
});
        