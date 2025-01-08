import styled from "styled-components";
import { FC } from "react";
import { SubmitHandler, useController, useForm } from "react-hook-form";
import { Note } from "../model/note";
import { TextField } from "@mui/material";
import { Modal } from "./Modal";

const Form = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  padding: 12px;
  width: 360px;
`;

type SendNote = Omit<Note, "id" | "userID">;

const InitialValues: SendNote = {
  title: "",
  data: "",
  meta: "",
};

type Props = {
  open: boolean;
  onAddNote: (note: Note) => Error | null;
  onClose: () => void;
};

const AddNote: FC<Props> = ({ open, onAddNote, onClose }) => {
  const { control, handleSubmit, reset } = useForm<SendNote>({
    defaultValues: InitialValues,
  });

  const titleField = useController({
    name: "title",
    control,
    rules: { required: true },
  });

  const dataField = useController({
    name: "data",
    control,
    rules: { required: true },
  });
  const meataField = useController({
    name: "meta",
    control,
    rules: { required: true },
  });

  const onSubmit: SubmitHandler<SendNote> = (values) => {
    const err = onAddNote(values);

    if (!err) {
      reset();
    }
  };

  const onCloseHandle = () => {
    reset();
    onClose();
  };

  return (
    <Modal
      id={"note"}
      title={"Новая заметка"}
      open={open}
      onClose={onCloseHandle}
      onConfirm={() => undefined}
    >
      <Form id={"note"} onSubmit={handleSubmit(onSubmit)}>
        <TextField fullWidth {...titleField.field} required label="Заголовок" />
        <TextField
          fullWidth
          multiline
          {...dataField.field}
          required
          label="Заметка"
        />
        <TextField
          fullWidth
          multiline
          {...meataField.field}
          required
          label="Описание"
        />
      </Form>
    </Modal>
  );
};

export { AddNote };
