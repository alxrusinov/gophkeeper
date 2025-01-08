import { useMutation, useQuery } from "@tanstack/react-query";
import { AxiosResponse } from "axios";
import { useContext, useEffect } from "react";
import { AxiosContext } from "../queryClient";
import { useKeeper } from "../store";
import { Credentials } from "../model/credentials";

const useCredentials = () => {
  const { apiClient } = useContext(AxiosContext);

  const creds = useKeeper((state) => state.credentials);
  const setCreds = useKeeper((state) => state.setCredentials);

  const getCreds = useQuery({
    queryKey: ["get creds"],
    queryFn: async () => {
      const res = await apiClient.get<Credentials[]>("/credentials");

      return res.data;
    },
  });

  const sendCredentials = useMutation({
    mutationKey: ["add creds"],
    mutationFn: async (data: Credentials) => {
      await apiClient.post<void, AxiosResponse<void, Error>, Credentials>(
        "/credentials",
        data
      );
    },
  });

  const addCreds = (data: Credentials): Error | null => {
    let error: Error | null = null;
    sendCredentials.mutate(data, {
      onSuccess: () => {
        getCreds.refetch();
      },
      onError: (err) => {
        error = err;
      },
    });
    return error;
  };

  useEffect(() => {
    if (getCreds.data) {
      setCreds(getCreds.data);
    }
  }, [getCreds.data, setCreds]);

  return {
    data: creds,
    addSource: addCreds,
  };
};

export { useCredentials };
