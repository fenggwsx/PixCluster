"use client";

import { memo, useContext } from "react";
import { Card, Slider, Typography } from "antd";
import { HyperParametersContext } from "@/contexts/hyper-parameters";

const { Text } = Typography;

export const KSetting = memo(() => {
  const { hyperK, setHyperK } = useContext(HyperParametersContext);

  return (
    <Card
      title="超参数K设置"
      extra={<Text type="secondary">仅对新添加的图片生效</Text>}
    >
      <Slider
        min={2}
        max={20}
        marks={{
          2: "2",
          5: "5",
          8: "8",
          11: "11",
          14: "14",
          17: "17",
          20: "20",
        }}
        value={hyperK}
        onChange={(value) => setHyperK(value)}
      />
    </Card>
  );
});

KSetting.displayName = "KSetting";
