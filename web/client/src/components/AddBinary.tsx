import styled from "styled-components";
import { FC, useState } from "react";
import { SubmitHandler, useController, useForm } from "react-hook-form";
import { styled as styledMui } from "@mui/material/styles";
import Button from "@mui/material/Button";
import CloudUploadIcon from "@mui/icons-material/CloudUpload";

import { TextField, Typography } from "@mui/material";
import { Modal } from "./Modal";
import { BinaryUpload } from "../model/binary";

const Form = styled.form`
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 20px;
  padding: 12px;
  width: 360px;
`;

type SendBinary = {
  title: string;
  data: File | undefined;
  meta: string;
};

const InitialValues: SendBinary = {
  title: "",
  data: undefined,
  meta: "",
};

const VisuallyHiddenInput = styledMui("input")({
  clip: "rect(0 0 0 0)",
  clipPath: "inset(50%)",
  height: 1,
  overflow: "hidden",
  position: "absolute",
  bottom: 0,
  left: 0,
  whiteSpace: "nowrap",
  width: 1,
});

type Props = {
  open: boolean;
  onAddBinary: (data: BinaryUpload) => Error | null;
  onClose: () => void;
};

const AddBinary: FC<Props> = ({ open, onAddBinary, onClose }) => {
  const [overLimit, setOverLimit] = useState(false);

  const { control, handleSubmit, reset, setValue } = useForm<SendBinary>({
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

  const onSubmit: SubmitHandler<SendBinary> = (values) => {
    if (values.data) {
      const prepared: BinaryUpload = {
        title: values.title,
        meta: values.meta,
        data: values.data,
        mimeType: values.data.type,
      };
      const err = onAddBinary(prepared);

      if (!err) {
        reset();
      }
    }
  };

  const onChangeFile: React.ChangeEventHandler<HTMLInputElement> = (e) => {
    const fileInput = e.target.files?.[0];

    if (fileInput && fileInput.size > 33554432) {
      setOverLimit(true);
    } else {
      setOverLimit(false);
      setValue("data", e.target.files?.[0] || undefined);
    }
  };

  const onCloseHandle = () => {
    reset();
    onClose();
  };

  return (
    <Modal
      id={"banccard"}
      title={"Новый файл"}
      open={open}
      onClose={onCloseHandle}
      onConfirm={() => undefined}
    >
      <Form id={"banccard"} onSubmit={handleSubmit(onSubmit)}>
        <TextField fullWidth {...titleField.field} required label="Заголовок" />
        <TextField fullWidth {...meataField.field} required label="Описание" />
        {dataField && <Typography>{dataField.field.value?.name}</Typography>}
        <Button
          component="label"
          role={undefined}
          variant="contained"
          tabIndex={-1}
          startIcon={<CloudUploadIcon />}
        >
          Загрузить файл
          <VisuallyHiddenInput
            name={dataField.field.name}
            type={"file"}
            onChange={onChangeFile}
            required
          />
        </Button>
        {!overLimit && <Typography fontSize={10}>Не более 32 MB</Typography>}
        {overLimit && (
          <Typography fontSize={10} color="error">
            Файл должен быть не более 32 MB
          </Typography>
        )}
      </Form>
    </Modal>
  );
};

export { AddBinary };
