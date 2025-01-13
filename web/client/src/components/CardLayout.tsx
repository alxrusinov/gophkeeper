import styled from "styled-components";
import { FC } from "react";
import { Box, Grid2 } from "@mui/material";

const Root = styled.div``;

type Props = {
  children: React.ReactNode;
  className?: string;
};

const CardLayout: FC<Props> = ({ children, className }) => {
  return (
    <Root className={className}>
      <Box
        sx={{
          paddingY: 6,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Grid2
          container
          spacing={2}
          sx={{
            justifyContent: "center",
          }}
        >
          {children}
        </Grid2>
      </Box>
    </Root>
  );
};

export { CardLayout };
