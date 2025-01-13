import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
} from "@mui/material";
import { FC, ReactNode } from "react";

type Props = {
  title: string;
  open: boolean;
  children: ReactNode;
  id: string;
  onClose: () => void;
  onConfirm: () => void;
};

const Modal: FC<Props> = ({
  id,
  title,
  open,
  children,
  onConfirm,
  onClose,
}) => {
  const onConfirmHanlde = () => {
    onConfirm();
  };

  const onCloseHanlde = () => {
    onClose();
  };

  return (
    <Dialog open={open} onClose={onClose}>
      <DialogTitle>{title}</DialogTitle>
      <DialogContent>{children}</DialogContent>
      <DialogActions>
        <Button onClick={onCloseHanlde}>Отменить</Button>
        <Button form={id} onClick={onConfirmHanlde} autoFocus type="submit">
          Создать
        </Button>
      </DialogActions>
    </Dialog>
  );
};

export { Modal };
