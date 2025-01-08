import { QueryClient } from "@tanstack/query-core";
import axios, { AxiosError, AxiosInstance } from "axios";
import { createContext, FC, ReactNode, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

type AxiosContextValues = {
  apiClient: AxiosInstance;
  authClient: AxiosInstance;
};

type Props = {
  children: ReactNode;
};

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API,
  withCredentials: true,
});

const authClient = axios.create({
  baseURL: import.meta.env.VITE_AUTH,
  withCredentials: true,
});

const AxiosContext = createContext<AxiosContextValues>(
  {} as AxiosContextValues
);

const queryClient = new QueryClient();

const AxiosProvider: FC<Props> = ({ children }) => {
  const navigate = useNavigate();
  const [unAuth, setUnAuth] = useState(false);

  useEffect(() => {
    apiClient.interceptors.response.use(
      (config) => {
        return config;
      },
      async (err: AxiosError) => {
        if (!axios.isAxiosError(err)) {
          return Promise.reject(err);
        }

        const originalRequest = err.config;

        console.log("ERR", err);
        if (err.response?.status === 401 && originalRequest) {
          setUnAuth(true);
        } else {
          return Promise.reject(err);
        }
      }
    );
  }, []);

  useEffect(() => {
    if (unAuth) {
      navigate("/auth");
      setUnAuth(false);
    }
  }, [navigate, unAuth]);

  return (
    <AxiosContext.Provider value={{ apiClient, authClient }}>
      {children}
    </AxiosContext.Provider>
  );
};

export { AxiosContext, AxiosProvider, queryClient };
