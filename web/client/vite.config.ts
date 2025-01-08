import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import svgx from "@svgx/vite-plugin-react";

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(), // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    svgx({
      optimize: true,
      svgoConfig: {
        plugins: [
          {
            name: "preset-default",
            params: { overrides: { removeViewBox: false } },
          },
        ],
      },
    }),
  ],
});
