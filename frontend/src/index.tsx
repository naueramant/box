import React from "react";
import ReactDOM from "react-dom";
import Routes from "routes";
import theme from "theme";
import "./index.scss";
import { ThemeProvider } from "@mui/material";
import reportWebVitals from "./reportWebVitals";

ReactDOM.render(
    <React.StrictMode>
        <ThemeProvider theme={theme} >
            <Routes />
        </ThemeProvider>
    </React.StrictMode>,
    document.getElementById("root")
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
