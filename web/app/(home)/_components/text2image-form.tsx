"use client";

import { memo, useContext } from "react";
import { Button, Form, Input } from "antd";
import type { FormProps } from "antd";
import { v7 as uuidv7 } from "uuid";
import { HyperParametersContext } from "@/contexts/hyper-parameters";
import { AnalysisInstanceListDispatchersContext } from "@/contexts/analysis-instance-list";
import type { Text2ImagePrompt } from "@/types";

const { TextArea } = Input;

export const Text2ImageForm = memo(() => {
  const { hyperK } = useContext(HyperParametersContext);
  const { addAnalysisInstance } = useContext(
    AnalysisInstanceListDispatchersContext,
  );

  const onFinish: FormProps<Text2ImagePrompt>["onFinish"] = (values) => {
    addAnalysisInstance({
      uuid: uuidv7(),
      hyperK: hyperK,
      prompt: values,
      taskId: null,
      image: null,
      result: null,
      summary: null,
    });
  };

  return (
    <Form layout="vertical" variant="filled" onFinish={onFinish}>
      <Form.Item
        label="正面提示词"
        name="positive"
        rules={[{ required: true }]}
      >
        <TextArea style={{ resize: "none" }} rows={2} />
      </Form.Item>
      <Form.Item label="负面提示词" name="negative">
        <TextArea style={{ resize: "none" }} rows={2} />
      </Form.Item>
      <Form.Item label={null}>
        <Button type="primary" htmlType="submit">
          提交
        </Button>
      </Form.Item>
    </Form>
  );
});

Text2ImageForm.displayName = "Text2ImageForm";
