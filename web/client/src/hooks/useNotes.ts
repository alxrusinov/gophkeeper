import { useMutation, useQuery } from "@tanstack/react-query";
import { AxiosResponse } from "axios";
import { useContext, useEffect } from "react";
import { Note } from "../model/note";
import { AxiosContext } from "../queryClient";
import { useKeeper } from "../store";

const useNotes = () => {
  const { apiClient } = useContext(AxiosContext);

  const notes = useKeeper((state) => state.notes);
  const setNotes = useKeeper((state) => state.setNotes);

  const getNotes = useQuery({
    queryKey: ["get notes"],
    queryFn: async () => {
      const res = await apiClient.get<Note[]>("/note");

      return res.data;
    },
  });

  const sendNote = useMutation({
    mutationKey: ["add note"],
    mutationFn: async (data: Note) => {
      await apiClient.post<void, AxiosResponse<void, Error>, Note>(
        "/note",
        data
      );
    },
  });

  const addNote = (data: Note): Error | null => {
    let error: Error | null = null;
    sendNote.mutate(data, {
      onSuccess: () => {
        getNotes.refetch();
      },
      onError: (err) => {
        error = err;
      },
    });
    return error;
  };

  useEffect(() => {
    if (getNotes.data) {
      setNotes(getNotes.data);
    }
  }, [getNotes.data, setNotes]);

  return {
    data: notes,
    addSource: addNote,
  };
};

export { useNotes };
