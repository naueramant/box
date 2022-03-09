interface Environment {
    MOCKED: boolean;
}

const ENV: Environment = {
    MOCKED: process.env["REACT_APP_MOCKED"] === "true",
};

export { ENV };
