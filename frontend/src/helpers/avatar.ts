import { stringToColor } from "./color";

export function stringAvatar(name: string) {
    return {
        sx: {
            bgcolor: stringToColor(name),
        },
        children: name.split(" ").reduce((acc, curr) => {
            if (curr.length > 0) {
                acc += curr[0].toUpperCase();
            }
            return acc;
        }, ""),
    };
}
