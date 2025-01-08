import { FC, useContext, useEffect, useState } from "react";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { styled as styledMui } from "@mui/material/styles";
import IconButton, { IconButtonProps } from "@mui/material/IconButton";
import {
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Collapse,
  Typography,
} from "@mui/material";
import { Binary } from "../model/binary";
import { useQuery } from "@tanstack/react-query";
import { AxiosContext } from "../queryClient";

type Props = {
  item: Binary;
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

const BinCard: FC<Props> = ({ item }) => {
  const { apiClient } = useContext(AxiosContext);
  const [expanded, setExpanded] = useState(false);
  const [canDownload, setCanDownload] = useState(false);

  const downloadLink = useQuery({
    queryKey: ["download link"],
    queryFn: async () => {
      const res = await apiClient.get<Blob>(`/file/${item.fileID}`);

      return res.data;
    },
    enabled: Boolean(item.fileID && canDownload),
  });

  const onDownload = () => {
    setCanDownload(true);
  };

  useEffect(() => {
    if (downloadLink.data && canDownload) {
      const blob = new Blob([downloadLink.data], { type: item.mimeType });

      const url = window.URL.createObjectURL(blob);

      const link = document.createElement("a");
      link.href = url;
      link.setAttribute("download", "file");
      document.body.appendChild(link);
      link.click();
      link.parentNode?.removeChild(link);
      setCanDownload(false);
    }
  }, [canDownload, downloadLink.data, item.mimeType]);

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
          <CardContent
            sx={{
              display: "flex",
              flexDirection: "column",
              alignContent: "center",
            }}
          >
            <Typography>{item.mimeType}</Typography>
            <Typography>{item.fileID}</Typography>
            <Button onClick={onDownload} variant="contained">
              Скачать
            </Button>
          </CardContent>
        </Collapse>
      </CardContent>
    </Card>
  );
};

export { BinCard };
