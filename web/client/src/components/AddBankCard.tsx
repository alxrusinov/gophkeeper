import styled from "styled-components";
import { FC } from "react";
import { SubmitHandler, useController, useForm } from "react-hook-form";

import { TextField } from "@mui/material";
import { Modal } from "./Modal";
import { BankCard } from "../model/bankCard";

const Form = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  padding: 12px;
  width: 360px;
`;

type SendBankCard = {
  title: string;
  data: number;
  meta: string;
};

const InitialValues: SendBankCard = {
  title: "",
  data: 0,
  meta: "",
};

type Props = {
  open: boolean;
  onAddBankCard: (note: BankCard) => Error | null;
  onClose: () => void;
};

const AddBankCard: FC<Props> = ({ open, onAddBankCard, onClose }) => {
  const { control, handleSubmit, reset } = useForm<SendBankCard>({
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

  const onSubmit: SubmitHandler<SendBankCard> = (values) => {
    const prepared = {
      title: values.title,
      meta: values.meta,
      data: Number(values.data),
    };
    const err = onAddBankCard(prepared);

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
      id={"banccard"}
      title={"Новая карта"}
      open={open}
      onClose={onCloseHandle}
      onConfirm={() => undefined}
    >
      <Form id={"banccard"} onSubmit={handleSubmit(onSubmit)}>
        <TextField fullWidth {...titleField.field} required label="Заголовок" />
        <TextField
          fullWidth
          {...dataField.field}
          type={"number"}
          required
          label="Данные карты"
        />
        <TextField fullWidth {...meataField.field} required label="Описание" />
      </Form>
    </Modal>
  );
};

export { AddBankCard };
