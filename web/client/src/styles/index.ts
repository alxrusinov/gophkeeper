import { createGlobalStyle } from "styled-components";
import { reset } from "styled-reset";

const GlobalStyle = createGlobalStyle`
  ${reset}

  :root {
    box-sizing: border-box;
    font-family: Inter;
  }

  *, :before, :after {
    box-sizing: border-box;
    font-family: Inter;
  }

  html {
    box-sizing: border-box;
    font-family: Inter;

  }
`;

export { GlobalStyle };
