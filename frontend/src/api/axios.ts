import axios from "axios";

const http = axios.create({
    baseURL: "/api",
});

http.interceptors.request.use((config) => {
    if (!config.headers) {
        config.headers = {};
    }

    const token = localStorage.getItem("token");
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
});

export { http };
