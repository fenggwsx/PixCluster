"use client";

import { Button, Flex, Result } from "antd";
import Link from "next/link";

export default function NotFound() {
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
      <Result
        status="404"
        title="404"
        subTitle="抱歉，您访问的页面不存在"
        extra={
          <Link href="/">
            <Button type="primary">回到主页</Button>
          </Link>
        }
      />
    </Flex>
  );
}
