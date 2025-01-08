import { FC, useState } from "react";
import { Button, Menu, MenuItem, MenuProps } from "@mui/material";
import { styled as muiStyled, alpha } from "@mui/material/styles";
import styled from "styled-components";
import AddIcon from "@mui/icons-material/Add";
import NoteAddIcon from "@mui/icons-material/NoteAdd";
import FileUploadIcon from "@mui/icons-material/FileUpload";
import AddCardIcon from "@mui/icons-material/AddCard";
import PasswordIcon from "@mui/icons-material/Password";

const Root = styled.div``;

type Props = {
  onCreateNote: () => void;
  onCreateBinary: () => void;
  onCreateCredentials: () => void;
  onCreateBankCard: () => void;
  className?: string;
};

const StyledMenu = muiStyled((props: MenuProps) => (
  <Menu
    elevation={0}
    anchorOrigin={{
      vertical: "bottom",
      horizontal: "right",
    }}
    transformOrigin={{
      vertical: "top",
      horizontal: "right",
    }}
    {...props}
  />
))(({ theme }) => ({
  "& .MuiPaper-root": {
    borderRadius: 6,
    marginTop: theme.spacing(1),
    minWidth: 180,
    color: "rgb(55, 65, 81)",
    boxShadow:
      "rgb(255, 255, 255) 0px 0px 0px 0px, rgba(0, 0, 0, 0.05) 0px 0px 0px 1px, rgba(0, 0, 0, 0.1) 0px 10px 15px -3px, rgba(0, 0, 0, 0.05) 0px 4px 6px -2px",
    "& .MuiMenu-list": {
      padding: "4px 0",
    },
    "& .MuiMenuItem-root": {
      "& .MuiSvgIcon-root": {
        fontSize: 18,
        color: theme.palette.text.secondary,
        marginRight: theme.spacing(1.5),
      },
      "&:active": {
        backgroundColor: alpha(
          theme.palette.primary.main,
          theme.palette.action.selectedOpacity
        ),
      },
    },
    ...theme.applyStyles("dark", {
      color: theme.palette.grey[300],
    }),
  },
}));

const MenuComponent: FC<Props> = ({
  onCreateNote,
  onCreateBinary,
  onCreateCredentials,
  onCreateBankCard,
  className,
}) => {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const open = Boolean(anchorEl);

  const handleClick = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const onAddNote = () => {
    onCreateNote();
    handleClose();
  };

  const onAddCredentials = () => {
    onCreateCredentials();
    handleClose();
  };

  const onAddBinary = () => {
    onCreateBinary();
    handleClose();
  };

  const onAddBankCard = () => {
    onCreateBankCard();
    handleClose();
  };

  return (
    <Root className={className}>
      <Button
        id="demo-customized-button"
        aria-controls={open ? "demo-customized-menu" : undefined}
        aria-haspopup="true"
        aria-expanded={open ? "true" : undefined}
        variant="contained"
        disableElevation
        onClick={handleClick}
        // endIcon={<KeyboardArrowDownIcon />}
        sx={{
          width: 64,
          height: 64,
          borderRadius: "50%",
        }}
      >
        <AddIcon
          sx={{
            width: 32,
            height: 32,
          }}
        />
      </Button>
      <StyledMenu
        id="demo-customized-menu"
        MenuListProps={{
          "aria-labelledby": "demo-customized-button",
        }}
        anchorEl={anchorEl}
        open={open}
        onClose={handleClose}
      >
        <MenuItem onClick={onAddNote} disableRipple>
          <NoteAddIcon />
          Заметка
        </MenuItem>
        <MenuItem onClick={onAddBinary} disableRipple>
          <FileUploadIcon />
          Файл
        </MenuItem>
        <MenuItem onClick={onAddBankCard} disableRipple>
          <AddCardIcon />
          Банковская карта
        </MenuItem>
        <MenuItem onClick={onAddCredentials} disableRipple>
          <PasswordIcon />
          Логин/Пароль
        </MenuItem>
      </StyledMenu>
    </Root>
  );
};

export { MenuComponent };
