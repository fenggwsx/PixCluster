"use client";

import { useContext } from "react";
import { Flex } from "antd";
import { AnalysisInstanceListContext } from "@/contexts/analysis-instance-list";
import { ImageSource } from "./_components/image-source";
import { KSetting } from "./_components/k-setting";
import { AnalysisResult } from "./_components/analysis-result";

export default function HomePage() {
  const analysisInstanceList = useContext(AnalysisInstanceListContext);

  return (
    <Flex
      style={{
        padding: 16,
        margin: "0 auto",
        maxWidth: 1280,
        width: "100%",
      }}
      gap="middle"
      vertical
    >
      <ImageSource />
      <KSetting />
      {analysisInstanceList.map((instance) => (
        <AnalysisResult key={instance.uuid} instance={instance} />
      ))}
    </Flex>
  );
}
