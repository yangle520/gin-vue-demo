import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ""),
      },
    },
  },
});

const base = "/";
const root = "./";
const outDir = "dist";

const config = {
  base: base, // 编译后js导入的资源路径
  root: root, // index.html 文件所在的位置
  publicDir: "public", // 静态资源文件夹
  resolve: { alais },
  define: {
    "process.env": {},
  },
};
