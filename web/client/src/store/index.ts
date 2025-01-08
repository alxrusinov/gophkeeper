import { BankCard } from "src/model/bankCard";
import { Binary } from "src/model/binary";
import { Credentials } from "src/model/credentials";
import { Note } from "src/model/note";
import { create } from "zustand";
import { devtools } from "zustand/middleware";

type State = {
  notes: Note[];
  credentials: Credentials[];
  binaries: Binary[];
  bankCards: BankCard[];
};

type Actions = {
  setNotes: (data: Note[]) => void;
  setCredentials: (data: Credentials[]) => void;
  setBinaries: (data: Binary[]) => void;
  setBankCards: (data: BankCard[]) => void;
};

type Store = State & Actions;

const initialState: State = {
  notes: [],
  credentials: [],
  binaries: [],
  bankCards: [],
};

const useKeeper = create<Store>()(
  devtools((set) => ({
    ...initialState,
    setNotes: (notes) => {
      set({ notes });
    },
    setCredentials: (credentials) => {
      set({ credentials });
    },
    setBinaries: (binaries) => {
      set({ binaries });
    },
    setBankCards: (bankCards) => {
      set({ bankCards });
    },
  }))
);

export { useKeeper };
