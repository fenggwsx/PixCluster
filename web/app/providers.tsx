"use client";

import { App } from "antd";
import { AntdRegistry } from "@ant-design/nextjs-registry";
import AntdConfigProvider from "@/providers/antd-config";
import AnalysisInstanceListProvider from "@/providers/analysis-instance-list";
import HyperParamtersProvider from "@/providers/hyper-paramters";

type Props = {
  children: React.ReactNode;
};

export default function Providers({ children }: Props) {
  return (
    <AntdRegistry>
      <AntdConfigProvider>
        <AnalysisInstanceListProvider>
          <HyperParamtersProvider>
            <App>{children}</App>
          </HyperParamtersProvider>
        </AnalysisInstanceListProvider>
      </AntdConfigProvider>
    </AntdRegistry>
  );
}
