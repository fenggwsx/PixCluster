"use client";

import { GithubFilled, PictureTwoTone } from "@ant-design/icons";
import { Button, Flex, Menu, Tooltip } from "antd";
import type { MenuProps } from "antd";
import Link from "next/link";
import { usePathname } from "next/navigation";
import { memo } from "react";
import styles from "./navbar.module.css";

type MenuItem = Required<MenuProps>["items"][number];

const menuItems: MenuItem[] = [
  {
    key: "/",
    label: <Link href="/">主页</Link>,
  },
  {
    key: "/about",
    label: <Link href="/about">关于</Link>,
  },
];

export const Navbar = memo(() => {
  const pathname = usePathname();

  return (
    <Flex
      style={{
        margin: "0 auto",
        padding: "0 24px",
        maxWidth: 1280,
      }}
      align="center"
      gap="middle"
    >
      <Flex gap="middle">
        <PictureTwoTone style={{ fontSize: 32 }} />
        <h1 className={styles.title}>
          PixCluster
          <span className={styles.description}> - 像素聚类智能分析平台</span>
        </h1>
      </Flex>
      <Menu
        style={{ minWidth: 0, flex: 1, fontSize: 16, border: "none" }}
        items={menuItems}
        mode="horizontal"
        selectedKeys={[pathname]}
      />
      <Tooltip title="Github">
        <Link
          style={{ display: "flex" }}
          href="https://github.com/fenggwsx/PixCluster"
          target="_blank"
        >
          <Button
            style={{ width: 40, height: 40 }}
            color="default"
            variant="text"
          >
            <GithubFilled style={{ fontSize: 20 }} />
          </Button>
        </Link>
      </Tooltip>
    </Flex>
  );
});

Navbar.displayName = "Navbar";
