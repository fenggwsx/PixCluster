"use client";

import { ConfigProvider } from "antd";
import zhCN from "antd/locale/zh_CN";

type Props = {
  children: Readonly<React.ReactNode>;
};

export default function AntdConfigProvider({ children }: Props) {
  return (
    <ConfigProvider
      locale={zhCN}
      theme={{
        components: {
          Layout: {
            headerBg: "#ffffff",
            headerPadding: 0,
          },
        },
      }}
    >
      {children}
    </ConfigProvider>
  );
}
