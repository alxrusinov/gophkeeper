import { useMutation, useQuery } from "@tanstack/react-query";
import { AxiosResponse } from "axios";
import { useContext, useEffect } from "react";
import { AxiosContext } from "../queryClient";
import { Binary, BinaryUpload } from "../model/binary";
import { DeleteSource } from "../model/deleteSource";
import { useKeeper } from "../store";

type Downloaded = {
  id: string;
  user_id: string;
  file_id: string;
  title: string;
  meta: string;
  mime_type: string;
};

const useBinaries = () => {
  const { apiClient } = useContext(AxiosContext);

  const bins = useKeeper((state) => state.binaries);
  const setBins = useKeeper((state) => state.setBinaries);

  const getBins = useQuery({
    queryKey: ["get binaries"],
    queryFn: async () => {
      const res = await apiClient.get<Downloaded[]>("/binary");

      return res.data;
    },
    select: (data): Binary[] => {
      return data.map<Binary>((item) => ({
        id: item.id,
        userID: item.user_id,
        title: item.title,
        meta: item.meta,
        fileID: item.file_id,
        mimeType: item.mime_type,
      }));
    },
  });

  const sendBins = useMutation({
    mutationKey: ["add binary"],
    mutationFn: async (data: BinaryUpload) => {
      const form = new FormData();

      form.append("data", data.data);
      form.append("mime_type", data.mimeType);
      form.append("title", data.title);
      form.append("meta", data.meta);

      const res = await apiClient.post<
        Binary,
        AxiosResponse<Binary, Error>,
        FormData
      >("/binary", form);

      return res.data;
    },
  });

  const deleteBinary = useMutation({
    mutationKey: ["delete binary"],
    mutationFn: async (data: DeleteSource) => {
      await apiClient.post<void, AxiosResponse<void, Error>, DeleteSource>(
        "/binary",
        data
      );
    },
  });

  const deleteSource = (data: DeleteSource) => {
    deleteBinary.mutate(data, {
      onSuccess: () => {
        getBins.refetch();
      },
      onError: (err) => {
        console.log(err);
      },
    });
  };

  const addBinary = (data: BinaryUpload): Error | null => {
    let error: Error | null = null;
    sendBins.mutate(data, {
      onSuccess: () => {
        getBins.refetch();
      },
      onError: (err) => {
        error = err;
      },
    });
    return error;
  };

  useEffect(() => {
    if (getBins.data) {
      setBins(getBins.data);
    }
  }, [getBins.data, setBins]);

  return {
    data: bins,
    addSource: addBinary,
    deleteSource,
  };
};

export { useBinaries };
