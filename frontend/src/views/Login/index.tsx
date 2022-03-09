import { Visibility, VisibilityOff } from "@mui/icons-material";
import { LoadingButton } from "@mui/lab";
import {
    Alert,
    Divider,
    FormControl,
    IconButton,
    InputAdornment,
    InputLabel,
    OutlinedInput,
    Paper,
    Stack,
    styled,
    useTheme,
} from "@mui/material";
import { FormEvent, FunctionComponent, useMemo, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Login } from "../../api/authentication";
import { useUser } from "hooks/user";

const Container = styled("div")(({ theme }) => ({
    height: "100%",
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: theme.palette.background.default,
}));

const LoginView: FunctionComponent = () => {
    const navigate = useNavigate();

    const handleLoginSubmit = async (username: string, password: string) => {
        await Login(username, password);
        navigate("/");
    };

    return (
        <Container>
            <LoginForm onSubmit={handleLoginSubmit} />
        </Container>
    );
};

interface LoginFormProps {
    onSubmit: (username: string, password: string) => Promise<void>;
}

const LoginForm: FunctionComponent<LoginFormProps> = (props) => {
    const theme = useTheme();
    const [_, setUser] = useUser();

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const [showPassword, setShowPassword] = useState(false);
    const [loading, setLoading] = useState(false);

    const [error, setError] = useState<unknown>();

    const formValid = useMemo(() => {
        return username.length > 0 && password.length > 0;
    }, [username, password]);

    const handleSubmit = async (e: FormEvent) => {
        e.preventDefault();

        if (!formValid) {
            return;
        }

        try {
            setLoading(true);
            await props.onSubmit(username, password);
            setUser({
                id: 1,
                username: username,
                createdAt: new Date(),
                updatedAt: new Date(),
            });
        } catch (e) {
            console.error(e);
            setError(e);
        } finally {
            setLoading(false);
        }
    };

    return (
        <Paper
            variant="outlined"
            sx={{
                minWidth: "300px",
                padding: theme.spacing(4),
            }}
        >
            <form onSubmit={handleSubmit}>
                <Stack spacing={2}>
                    <FormControl disabled={loading}>
                        <InputLabel htmlFor="username">Username</InputLabel>
                        <OutlinedInput
                            id="username"
                            type={"text"}
                            autoFocus
                            required
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            label="Username"
                        />
                    </FormControl>

                    <FormControl disabled={loading}>
                        <InputLabel htmlFor="password">Password</InputLabel>
                        <OutlinedInput
                            id="password"
                            type={showPassword ? "text" : "password"}
                            required
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                        aria-label="toggle password visibility"
                                        onClick={(_) => setShowPassword(!showPassword)}
                                        edge="end"
                                    >
                                        {showPassword ? <VisibilityOff /> : <Visibility />}
                                    </IconButton>
                                </InputAdornment>
                            }
                            label="Password"
                        />
                    </FormControl>

                    <Divider />

                    {error && <Alert severity="error">Login error</Alert>}

                    <LoadingButton
                        color="primary"
                        variant="contained"
                        size="large"
                        type="submit"
                        loading={loading}
                        disabled={!formValid}
                    >
                        Login
                    </LoadingButton>
                </Stack>
            </form>
        </Paper>
    );
};

export default LoginView;
