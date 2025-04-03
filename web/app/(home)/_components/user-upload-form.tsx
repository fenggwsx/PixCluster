"use client";

import { memo, useContext } from "react";
import { App, Upload } from "antd";
import { FileImageOutlined } from "@ant-design/icons";
import ImgCrop from "antd-img-crop";
import { v7 as uuidv7 } from "uuid";
import { getBlobDataURL } from "@/utils/base64";
import { HyperParametersContext } from "@/contexts/hyper-parameters";
import { AnalysisInstanceListDispatchersContext } from "@/contexts/analysis-instance-list";

const { Dragger } = Upload;

export const UserUploadForm = memo(() => {
  const { hyperK } = useContext(HyperParametersContext);
  const { addAnalysisInstance } = useContext(
    AnalysisInstanceListDispatchersContext,
  );
  const { message } = App.useApp();

  return (
    <ImgCrop modalTitle="裁剪图片">
      <Dragger
        beforeUpload={(file) => {
          getBlobDataURL(file)
            .then((image) =>
              addAnalysisInstance({
                uuid: uuidv7(),
                hyperK: hyperK,
                prompt: null,
                taskId: null,
                image: image,
                result: null,
                summary: null,
              }),
            )
            .catch((err) => {
              console.error(err);
              message.error("上传失败");
            });
          return false;
        }}
        height={250}
        fileList={[]}
      >
        <p className="ant-upload-drag-icon">
          <FileImageOutlined />
        </p>
        <p className="ant-upload-text">点击或拖拽文件到此区域上传图像</p>
        <p className="ant-upload-hint">
          上传的图像将被裁剪为正方形，请注意调整您的图片
        </p>
      </Dragger>
    </ImgCrop>
  );
});

UserUploadForm.displayName = "UserUploadForm";
