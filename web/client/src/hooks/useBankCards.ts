import { useMutation, useQuery } from "@tanstack/react-query";
import { AxiosResponse } from "axios";
import { useContext, useEffect } from "react";
import { AxiosContext } from "../queryClient";
import { useKeeper } from "../store";
import { BankCard } from "../model/bankCard";

const useBankCards = () => {
  const { apiClient } = useContext(AxiosContext);

  const bCards = useKeeper((state) => state.bankCards);
  const setBCards = useKeeper((state) => state.setBankCards);

  const getBankCards = useQuery({
    queryKey: ["get bank cards"],
    queryFn: async () => {
      const res = await apiClient.get<BankCard[]>("/bankcard");

      return res.data;
    },
  });

  const sendBins = useMutation({
    mutationKey: ["add bank card"],
    mutationFn: async (data: BankCard) => {
      await apiClient.post<void, AxiosResponse<void, Error>, BankCard>(
        "/bankcard",
        data
      );
    },
  });

  const addCard = (data: BankCard): Error | null => {
    let error: Error | null = null;
    sendBins.mutate(data, {
      onSuccess: () => {
        getBankCards.refetch();
      },
      onError: (err) => {
        error = err;
      },
    });
    return error;
  };

  useEffect(() => {
    if (getBankCards.data) {
      setBCards(getBankCards.data);
    }
  }, [getBankCards.data, setBCards]);

  return {
    data: bCards,
    addSource: addCard,
  };
};

export { useBankCards };
