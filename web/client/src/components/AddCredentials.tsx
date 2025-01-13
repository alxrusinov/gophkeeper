import styled from "styled-components";
import { FC } from "react";
import { SubmitHandler, useController, useForm } from "react-hook-form";
import { TextField } from "@mui/material";
import { Modal } from "./Modal";
import { Credentials } from "../model/credentials";

const Form = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  padding: 12px;
  width: 360px;
`;

type SendCredentials = {
  title: string;
  username: string;
  password: string;
  meta: string;
};

const InitialValues: SendCredentials = {
  title: "",
  username: "",
  password: "",
  meta: "",
};

type Props = {
  open: boolean;
  onAddCredentials: (creds: Credentials) => Error | null;
  onClose: () => void;
};

const AddCredentials: FC<Props> = ({ open, onAddCredentials, onClose }) => {
  const { control, handleSubmit, reset } = useForm<SendCredentials>({
    defaultValues: InitialValues,
  });

  const titleField = useController({
    name: "title",
    control,
    rules: { required: true },
  });

  const usernameField = useController({
    name: "username",
    control,
    rules: { required: true },
  });

  const passwordField = useController({
    name: "password",
    control,
    rules: { required: true },
  });

  const meataField = useController({
    name: "meta",
    control,
    rules: { required: true },
  });

  const onSubmit: SubmitHandler<SendCredentials> = (values) => {
    const preparedValues = {
      data: { username: values.username, password: values.password },
      title: values.title,
      meta: values.meta,
    };
    const err = onAddCredentials(preparedValues);

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
      id={"creds"}
      title={"Логин/Пароль"}
      open={open}
      onClose={onCloseHandle}
      onConfirm={() => undefined}
    >
      <Form id={"creds"} onSubmit={handleSubmit(onSubmit)}>
        <TextField fullWidth {...titleField.field} required label="Заголовок" />
        <TextField fullWidth {...usernameField.field} required label="Логин" />
        <TextField
          type="password"
          fullWidth
          {...passwordField.field}
          required
          label="Пароль"
        />
        <TextField fullWidth {...meataField.field} required label="Описание" />
      </Form>
    </Modal>
  );
};

export { AddCredentials };
