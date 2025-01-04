import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import App from "./App.tsx";
import { BrowserRouter } from "react-router-dom";
import { AxiosProvider, queryClient } from "./query-client";
import { QueryClientProvider } from "@tanstack/react-query";
import { GlobalStyle } from "./styles/index.ts";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <AxiosProvider>
        <QueryClientProvider client={queryClient}>
          <GlobalStyle />
          <App />
        </QueryClientProvider>
      </AxiosProvider>
    </BrowserRouter>
  </StrictMode>
);
