import { createTheme } from "@mui/material/styles";
import { Shadows } from "@mui/material/styles/shadows";

const theme = createTheme({
  shape: {
    borderRadius: 4,
  },
  shadows: new Array(25).fill("none") as Shadows,
  
  palette: {
    primary: {
      main: "#23272a",
      contrastText: "#fff",
    },
    secondary: {
      main: "#fff",
      contrastText: "#23272a",
    },
    background: {
      default: "#f4f9fc",
    },
    text: {
      primary: "#23272a",
    },
    error: {
      main: "#F26419",
    },
    warning: {
      main: "#F6AE2D",
    },
    info: {
      main: "#2a5d84",
    },
    success: {
      main: "#86D994",
    },
  },
  typography: {
    // fontFamily: "Open Sans",
    fontFamily: "Libre Franklin",
  },
});

export default theme;
