"use client";

import { Card, Flex, Typography } from "antd";

const { Title, Paragraph, Text } = Typography;

export default function AboutPage() {
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
      <Card>
        <Typography>
          <Title>关于PixCluster</Title>
          <Paragraph>
            PixCluster平台致力于将数据科学、视觉艺术与生成式AI相结合，打造一个兼具实用性与创造性的图像分析平台。通过机器学习算法与前沿大模型技术，帮助用户深入理解图像色彩构成，同时激发创意灵感。
          </Paragraph>
          <Title level={2}>核心功能亮点</Title>
          <Paragraph>
            <ul>
              <li>
                <Text strong>智能色彩解构</Text>
                ：使用高效的KMeans++算法，利用多核性能对图像像素进行毫秒级聚类
              </li>
              <li>
                <Text strong>AI图像工坊</Text>
                ：使用通义文生图模型，无需自行准备图片，即可体验像素聚类功能
              </li>
              <li>
                <Text strong>直观数据展示</Text>
                ：聚类结果通过柱状图进行展示，图像的色彩成分尽收眼底
              </li>
            </ul>
          </Paragraph>
          <Title level={2}>技术架构</Title>
          <Paragraph>
            <ul>
              <li>
                <Text strong>前端可视化</Text>：Next.js + Ant Design + Ant
                Design Chart
              </li>
              <li>
                <Text strong>后端处理</Text>
                ：基于阿里云函数计算与Golang运行时的云原生技术架构
              </li>
              <li>
                <Text strong>大模型</Text>：通义文生图模型
              </li>
            </ul>
          </Paragraph>
        </Typography>
      </Card>
    </Flex>
  );
}
