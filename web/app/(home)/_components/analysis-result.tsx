"use client";

import { memo, useContext } from "react";
import {
  App,
  Button,
  Card,
  Flex,
  Image,
  Popconfirm,
  Skeleton,
  Spin,
  theme,
} from "antd";
import { Column } from "@ant-design/charts";
import type { AnalysisInstance } from "@/types";
import { AnalysisInstanceListDispatchersContext } from "@/contexts/analysis-instance-list";

const { useToken } = theme;

type ClusterData = {
  color: string;
  count: number;
};

type Props = {
  instance: AnalysisInstance;
};

export const AnalysisResult = memo(({ instance }: Props) => {
  const { deleteAnalysisInstance } = useContext(
    AnalysisInstanceListDispatchersContext,
  );
  const { message } = App.useApp();
  const { token } = useToken();

  return (
    <Card
      extra={
        <Popconfirm
          title="确认删除"
          description="你确认要删除本图片的分析数据吗？"
          onConfirm={() => {
            deleteAnalysisInstance(instance.uuid);
            message.success("删除成功");
          }}
        >
          <Button color="danger" variant="link">
            删除
          </Button>
        </Popconfirm>
      }
      title={`像素分析结果(K=${instance.hyperK})`}
    >
      <Flex gap="middle" vertical>
        <Flex gap="middle" justify="center" wrap>
          <Card size="small">
            {instance.image ? (
              <Image
                alt=""
                height={240}
                src={`${instance.image.prefix},${instance.image.data}`}
                width={240}
              />
            ) : (
              <Skeleton.Image style={{ width: 240, height: 240 }} active />
            )}
          </Card>
          <Card
            style={{
              flex: "1",
              minWidth: 280,
              height: 2 * (token.Card?.bodyPaddingSM ?? 12) + 240,
              overflowX: "auto",
              overflowY: "hidden",
            }}
            size="small"
          >
            {instance.result !== null ? (
              <div style={{ minWidth: 480 }}>
                <Column
                  autoFit
                  data={instance.result}
                  height={240}
                  style={{ fill: ({ color }: ClusterData) => color }}
                  xField="color"
                  yField="count"
                />
              </div>
            ) : (
              <Spin>
                <div
                  style={{
                    backgroundColor: token.colorFillSecondary,
                    borderRadius: token.borderRadius,
                    height: 240,
                    width: "100%",
                  }}
                />
              </Spin>
            )}
          </Card>
        </Flex>
        <Card size="small" title="智能总结">
          {instance.summary ?? <Skeleton active />}
        </Card>
      </Flex>
    </Card>
  );
});

AnalysisResult.displayName = "AnalysisResult";
