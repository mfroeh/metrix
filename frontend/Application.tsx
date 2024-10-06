import React from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import About from "./pages/About";
import Home from "./pages/Home";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { CssBaseline } from "@mui/material";
import Summoner from "./pages/Summoner";
import Navbar from "./Navbar";

const theme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#6200EA',
      light: '#B388FF',
      dark: '#4A148C',
    },
    secondary: {
      main: '#00BFA5',
      light: '#64FFDA',
      dark: '#00796B',
    },
    background: {
      default: '#121212',
      paper: '#1E1E1E',
    },
  },
  typography: {
    fontFamily: '"Roboto", "Helvetica", "Arial", sans-serif',
    h1: {
      fontWeight: 300,
      fontSize: '6rem',
    },
    h2: {
      fontWeight: 400,
      fontSize: '3.75rem',
    },
    body1: {
      fontSize: '1rem',
      lineHeight: 1.5,
    },
  },
  shape: {
    borderRadius: 8,
  },
  components: {
    MuiButton: {
      styleOverrides: {
        root: {
          textTransform: 'none',
        },
      },
    },
  },
});

const container = document.getElementById("app");
const root = createRoot(container!);
root.render(
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route index element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/summoner/:summonerName/:tag" element={<Summoner />} />
      </Routes>
    </BrowserRouter>
  </ThemeProvider>
);
