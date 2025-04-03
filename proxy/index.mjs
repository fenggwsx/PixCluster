import "dotenv/config";
import express from "express";
import { createProxyMiddleware } from "http-proxy-middleware";

const app = express();
const port = process.env.PORT || 9000;

const frontendUrl = process.env.FRONTEND_URL;
const backendUrl = process.env.BACKEND_URL;

app.use(
  "/api",
  createProxyMiddleware({
    target: backendUrl,
    changeOrigin: true,
  }),
);

app.use(
  "/",
  createProxyMiddleware({
    target: frontendUrl,
    changeOrigin: true,
    ws: true,
  }),
);

app.listen(port, () => {
  console.log(`proxy server running on port ${port}`);
  console.log(`frontend upstream: ${frontendUrl}`);
  console.log(`backend upstream: ${backendUrl}`);
});
