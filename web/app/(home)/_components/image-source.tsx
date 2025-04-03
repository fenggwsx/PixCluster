"use client";

import { memo, useState } from "react";
import { Card } from "antd";
import { UserUploadForm } from "./user-upload-form";
import { Text2ImageForm } from "./text2image-form";

const imageSourceTabList = [
  { key: "upload", label: "用户上传" },
  { key: "text2image", label: "文生图" },
];

const imageSourceContentList: Record<string, React.ReactNode> = {
  upload: <UserUploadForm />,
  text2image: <Text2ImageForm />,
};

export const ImageSource = memo(() => {
  const [activeTabKey, setActiveTabKey] = useState("upload");

  return (
    <Card
      style={{ height: 360 }}
      tabList={imageSourceTabList}
      onTabChange={(key) => setActiveTabKey(key)}
    >
      {imageSourceContentList[activeTabKey]}
    </Card>
  );
});

ImageSource.displayName = "ImageSource";
