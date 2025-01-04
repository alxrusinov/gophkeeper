import { QueryClient } from "@tanstack/query-core";
import axios, { AxiosInstance } from "axios";
import { createContext, FC, ReactNode } from "react";
import React from "react";

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
  return (
    <AxiosContext.Provider value={{ apiClient, authClient }}>
      {children}
    </AxiosContext.Provider>
  );
};

export { AxiosContext, AxiosProvider, queryClient };
