"use client";

import { Layout, theme } from "antd";
import Providers from "./providers";
import { Navbar } from "./_components/navbar";

const { Content, Header } = Layout;
const { useToken } = theme;

type Props = {
  children: Readonly<React.ReactNode>;
};

export default function ClientLayout({ children }: Props) {
  const { token } = useToken();

  return (
    <Providers>
      <Layout>
        <Header
          style={{
            position: "sticky",
            top: 0,
            zIndex: 1,
            boxShadow: token.boxShadowTertiary,
          }}
        >
          <Navbar />
        </Header>
        <Content>{children}</Content>
      </Layout>
    </Providers>
  );
}
