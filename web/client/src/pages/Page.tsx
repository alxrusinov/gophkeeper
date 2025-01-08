import styled from "styled-components";
import { FC, ReactNode } from "react";

const Root = styled.div`
  width: 100%;
  height: auto;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  padding: 10px 30px;
`;

const Container = styled.div`
  width: 100%;
  max-width: 1920px;
  height: 100%;
`;

type Props = {
  children: ReactNode;
};

const Page: FC<Props> = ({ children }) => {
  return (
    <Root>
      <Container>{children}</Container>
    </Root>
  );
};

export { Page };
