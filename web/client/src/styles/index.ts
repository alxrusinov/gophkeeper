import { createGlobalStyle } from "styled-components";

const GlobalStyle = createGlobalStyle`

  :root {
    box-sizing: border-box;
    font-family: Inter;
  }

  *, :before, :after {
    box-sizing: border-box;

  }

  html {
    box-sizing: border-box;
    font-size: 14px;
  }
`;

export { GlobalStyle };
