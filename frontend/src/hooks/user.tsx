import { createGlobalState } from "react-use";

interface User {
    id: number;
    username: string;
    createdAt: Date;
    updatedAt: Date;
}

export const useUser = createGlobalState<User | null>(null);
