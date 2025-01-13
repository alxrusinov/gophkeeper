import { FC, useState } from "react";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { styled as styledMui } from "@mui/material/styles";
import IconButton, { IconButtonProps } from "@mui/material/IconButton";
import {
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Collapse,
  Typography,
} from "@mui/material";
import { BankCard } from "../model/bankCard";

type Props = {
  item: BankCard;
};

interface ExpandMoreProps extends IconButtonProps {
  expand: boolean;
}

// eslint-disable-next-line @typescript-eslint/no-unused-vars
const ExpandMore = styledMui(({ expand, ...rest }: ExpandMoreProps) => {
  return <IconButton {...rest} />;
})(({ theme }) => ({
  marginLeft: "auto",
  transition: theme.transitions.create("transform", {
    duration: theme.transitions.duration.shortest,
  }),
  variants: [
    {
      props: ({ expand }) => !expand,
      style: {
        transform: "rotate(0deg)",
      },
    },
    {
      props: ({ expand }) => !!expand,
      style: {
        transform: "rotate(180deg)",
      },
    },
  ],
}));

const BCard: FC<Props> = ({ item }) => {
  const [expanded, setExpanded] = useState(false);

  const handleExpandClick = () => {
    setExpanded(!expanded);
  };
  return (
    <Card sx={{ width: 345 }}>
      <CardHeader title={item.title} />
      <CardContent>
        <Typography variant="body2" sx={{ color: "text.secondary" }}>
          {item.meta}
        </Typography>
        <CardActions disableSpacing>
          <ExpandMore
            expand={expanded}
            onClick={handleExpandClick}
            aria-expanded={expanded}
            aria-label="show more"
          >
            <ExpandMoreIcon />
          </ExpandMore>
        </CardActions>
        <Collapse in={expanded} timeout="auto" unmountOnExit>
          <CardContent>
            <Typography>{item.data}</Typography>
          </CardContent>
        </Collapse>
      </CardContent>
    </Card>
  );
};

export { BCard };
