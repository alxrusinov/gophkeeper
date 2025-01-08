import { FC, useState } from "react";
import { Page } from "./Page";
import { useNotes } from "../hooks/useNotes";
import { useCredentials } from "../hooks/useCredentials";
import { useBinaries } from "../hooks/useBinaries";
import { useBankCards } from "../hooks/useBankCards";
import { Grid2, Tab, Tabs } from "@mui/material";
import { CardLayout } from "../components/CardLayout";
import { NoteCard } from "../components/NoteCard";
import { BCard } from "../components/BCard";
import { BinCard } from "../components/BinCard";
import { CredsCard } from "../components/CredsCard";
import styled from "styled-components";
import { MenuComponent } from "../components/MenuComponent";
import { AddNote } from "../components/AddNote.tsx";
import { Note } from "../model/note.ts";
import { Toast } from "../components/Toast.tsx";
import { AddCredentials } from "../components/AddCredentials.tsx";
import { Credentials } from "../model/credentials.ts";
import { BankCard } from "../model/bankCard.ts";
import { AddBankCard } from "../components/AddBankCard.tsx";
import { AddBinary } from "../components/AddBinary.tsx";
import { BinaryUpload } from "../model/binary.ts";

enum TabValue {
  Note,
  Credentials,
  Binary,
  BankCard,
}

const TabLabel: Record<TabValue, string> = {
  [TabValue.Note]: "Заметки",
  [TabValue.Credentials]: "Учетные записи",
  [TabValue.BankCard]: "Банковские карты",
  [TabValue.Binary]: "Файлы",
};

const StyledMenu = styled(MenuComponent)`
  position: fixed;
  bottom: 160px;
  right: 40px;
`;

type ToastObj = {
  message: string;
  open: boolean;
};

const MainPage: FC = () => {
  const [currentTab, setCurrentTab] = useState<TabValue>(TabValue.Note);
  const [noteModal, setNoteModal] = useState(false);
  const [credsModal, setCredsModal] = useState(false);
  const [bankCardModal, setBankCardModal] = useState(false);
  const [binaryModal, setBinaryModal] = useState(false);
  const [toastObj, setToastObj] = useState<ToastObj>({
    open: false,
    message: "",
  });

  const onChangeTab = (_: React.SyntheticEvent, value: TabValue) => {
    setCurrentTab(value);
  };

  const notes = useNotes();
  const creds = useCredentials();
  const bins = useBinaries();
  const bCards = useBankCards();

  const onCreateNote = () => {
    setNoteModal(true);
  };

  const onCreateBinary = () => {
    setBinaryModal(true);
  };

  const onCreateCredentials = () => {
    setCredsModal(true);
  };

  const onCreateBankCard = () => {
    setBankCardModal(true);
  };

  const onAddNote = (note: Note): Error | null => {
    const err = notes.addSource(note);

    if (err !== null) {
      setToastObj({ open: true, message: err.message });
      setNoteModal(false);
      return err;
    }

    setNoteModal(false);
    return null;
  };

  const onAddCredentials = (credentials: Credentials): Error | null => {
    const err = creds.addSource(credentials);

    if (err !== null) {
      setToastObj({ open: true, message: err.message });
      setCredsModal(false);
      return err;
    }

    setCredsModal(false);
    return null;
  };

  const onAddBankCard = (card: BankCard): Error | null => {
    const err = bCards.addSource(card);

    if (err !== null) {
      setToastObj({ open: true, message: err.message });
      setBankCardModal(false);
      return err;
    }

    setBankCardModal(false);
    return null;
  };

  const onAddBinary = (data: BinaryUpload): Error | null => {
    const err = bins.addSource(data);

    if (err !== null) {
      setToastObj({ open: true, message: err.message });
      setBinaryModal(false);
      return err;
    }

    setBinaryModal(false);
    return null;
  };

  const onCloseToast = () => {
    setToastObj({ open: false, message: "" });
  };

  return (
    <Page>
      <Tabs value={currentTab} onChange={onChangeTab} centered>
        <Tab value={TabValue.Note} label={TabLabel[TabValue.Note]} />
        <Tab
          value={TabValue.Credentials}
          label={TabLabel[TabValue.Credentials]}
        />
        <Tab value={TabValue.Binary} label={TabLabel[TabValue.Binary]} />
        <Tab value={TabValue.BankCard} label={TabLabel[TabValue.BankCard]} />
      </Tabs>

      {notes.data && currentTab === TabValue.Note && (
        <CardLayout>
          {notes.data.map((item) => (
            <Grid2 size={"auto"} key={item.id}>
              <NoteCard item={item} />
            </Grid2>
          ))}
        </CardLayout>
      )}
      {creds.data && currentTab === TabValue.Credentials && (
        <CardLayout>
          {creds.data.map((item) => (
            <Grid2 size={"auto"} key={item.id}>
              <CredsCard item={item} />
            </Grid2>
          ))}
        </CardLayout>
      )}
      {bins.data && currentTab === TabValue.Binary && (
        <CardLayout>
          {bins.data.map((item) => (
            <Grid2 size={"auto"} key={item.id}>
              <BinCard item={item} />
            </Grid2>
          ))}
        </CardLayout>
      )}
      {bCards.data && currentTab === TabValue.BankCard && (
        <CardLayout>
          {bCards.data.map((item) => (
            <Grid2 size={"auto"} key={item.id}>
              <BCard item={item} />
            </Grid2>
          ))}
        </CardLayout>
      )}
      <StyledMenu
        onCreateNote={onCreateNote}
        onCreateBankCard={onCreateBankCard}
        onCreateBinary={onCreateBinary}
        onCreateCredentials={onCreateCredentials}
      />
      <AddNote
        open={noteModal}
        onClose={() => setNoteModal(false)}
        onAddNote={onAddNote}
      />
      <AddCredentials
        open={credsModal}
        onClose={() => setCredsModal(false)}
        onAddCredentials={onAddCredentials}
      />
      <AddBankCard
        open={bankCardModal}
        onClose={() => setBankCardModal(false)}
        onAddBankCard={onAddBankCard}
      />
      <AddBinary
        open={binaryModal}
        onClose={() => setBinaryModal(false)}
        onAddBinary={onAddBinary}
      />
      <Toast
        open={toastObj.open}
        message={toastObj.message}
        onClose={onCloseToast}
      />
    </Page>
  );
};

export { MainPage };
