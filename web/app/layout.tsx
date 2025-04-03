import type { Metadata } from "next";
import ClientLayout from "./client-layout";
import "./globals.css";

export const metadata: Metadata = {
  title: "PixCluster - 像素聚类智能分析平台",
  description:
    "专业的像素聚类智能分析平台，运用AI技术进行图像的像素值聚类，支持文本生成图像与聚类结果智能总结，探索更多图像信息挖掘的可能性。",
  keywords: ["像素聚类", "AI图像分析", "文生图", "像素信息挖掘"],
  authors: [{ name: "Guwei Feng" }],
};

type Props = {
  children: Readonly<React.ReactNode>;
};

export default function RootLayout({ children }: Props) {
  return (
    <html lang="zh-CN">
      <body>
        <ClientLayout>{children}</ClientLayout>
      </body>
    </html>
  );
}
