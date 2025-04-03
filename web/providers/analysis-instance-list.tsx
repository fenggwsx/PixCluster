"use client";

import {
  AnalysisInstance,
  AnalysisResult,
  DataURL,
  ResponseWrapper,
  Text2ImagePrompt,
} from "@/types";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import {
  AnalysisInstanceListContext,
  AnalysisInstanceListDispatchersContext,
} from "@/contexts/analysis-instance-list";
import { sleep } from "@/utils/sleep";

type Props = {
  children: Readonly<React.ReactNode>;
};

type CreateTaskResponse = ResponseWrapper<{ task_id: string }>;

type GetTaskResponse = ResponseWrapper<
  | {
      status: "SUCCEEDED";
      prefix: string;
      data: string;
    }
  | { status: "PENDING" | "RUNNING" | "SUSPENDED" | "FAILED" | "UNKNOWN" }
>;

type KMeansResponse = ResponseWrapper<{
  result: AnalysisResult;
}>;

type SummarizeResponse = ResponseWrapper<{
  summary: string;
}>;

async function createText2ImageTask(
  prompt: Text2ImagePrompt,
  attempts: number,
) {
  try {
    const res = await fetch("/api/v1/text2image", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        positive: prompt.positive,
        negative: prompt.negative,
      }),
    });
    const json = (await res.json()) as CreateTaskResponse;
    if (!json.success) throw new Error(json.message);
    return json.data.task_id;
  } catch (error) {
    if (attempts <= 1) throw error;
    if (error instanceof Error) console.error(error);
    await sleep(500);
    return createText2ImageTask(prompt, attempts - 1);
  }
}

async function getText2ImageResult(taskId: string, attempts: number) {
  try {
    const res = await fetch(`/api/v1/text2image/${taskId}`);
    const json = (await res.json()) as GetTaskResponse;
    if (!json.success) throw new Error(json.message);
    if (json.data.status === "SUCCEEDED")
      return { prefix: json.data.prefix, data: json.data.data } as DataURL;
    else throw json.data.status;
  } catch (error) {
    if (attempts <= 1) throw error;
    if (error instanceof Error) console.error(error);
    await sleep(1000);
    return getText2ImageResult(taskId, attempts - 1);
  }
}

async function getKMeansResult(
  hyperK: number,
  image: DataURL,
  attempts: number,
) {
  try {
    const res = await fetch("/api/v1/kmeans", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        k: hyperK,
        image_url: {
          prefix: image.prefix,
          data: image.data,
        },
      }),
    });
    const json = (await res.json()) as KMeansResponse;
    if (!json.success) throw new Error(json.message);
    return json.data.result;
  } catch (error) {
    if (attempts <= 1) throw error;
    if (error instanceof Error) console.error(error);
    await sleep(500);
    return getKMeansResult(hyperK, image, attempts - 1);
  }
}

async function summarizeKMeansResult(
  image: DataURL,
  result: AnalysisResult,
  attempts: number,
) {
  try {
    const res = await fetch("/api/v1/summarize", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        image_url: image,
        kmeans_result: result,
      }),
    });
    const json = (await res.json()) as SummarizeResponse;
    if (!json.success) throw new Error(json.message);
    return json.data.summary;
  } catch (error) {
    if (attempts <= 1) throw error;
    if (error instanceof Error) console.error(error);
    await sleep(1000);
    return summarizeKMeansResult(image, result, attempts - 1);
  }
}

export default function AnalysisInstanceListProvider({ children }: Props) {
  const isMounted = useRef(false);
  const [analysisInstanceList, setAnalysisInstanceList] = useState<
    AnalysisInstance[]
  >([]);

  useEffect(() => {
    isMounted.current = true;
    return () => {
      isMounted.current = false;
    };
  }, []);

  const updateAnalysisInstance = useCallback(
    (newInstance: AnalysisInstance) => {
      setAnalysisInstanceList((prevList) =>
        prevList.map((instance) =>
          instance.uuid === newInstance.uuid ? newInstance : instance,
        ),
      );
    },
    [],
  );

  const analyzeInstance = useCallback(
    async (instance: AnalysisInstance) => {
      if (instance.summary !== null) return instance;
      if (instance.image !== null && instance.result !== null) {
        const newInstance = {
          ...instance,
          summary: await summarizeKMeansResult(
            instance.image,
            instance.result,
            3,
          ),
        };
        if (isMounted.current) updateAnalysisInstance(newInstance);
        return analyzeInstance(newInstance);
      }
      if (instance.image !== null) {
        const newInstance = {
          ...instance,
          result: await getKMeansResult(instance.hyperK, instance.image, 3),
        };
        if (isMounted.current) updateAnalysisInstance(newInstance);
        return analyzeInstance(newInstance);
      }
      if (instance.taskId !== null) {
        const newInstance = {
          ...instance,
          image: await getText2ImageResult(instance.taskId, 30),
        };
        if (isMounted.current) updateAnalysisInstance(newInstance);
        return analyzeInstance(newInstance);
      }
      if (instance.prompt !== null) {
        const newInstance = {
          ...instance,
          taskId: await createText2ImageTask(instance.prompt, 3),
        };
        if (isMounted.current) updateAnalysisInstance(newInstance);
        return analyzeInstance(newInstance);
      }
      throw new Error("Failed to analysis the instance");
    },
    [updateAnalysisInstance],
  );

  const addAnalysisInstance = useCallback(
    (newInstance: AnalysisInstance) => {
      setAnalysisInstanceList((prevList) => [newInstance, ...prevList]);
      analyzeInstance(newInstance);
    },
    [analyzeInstance],
  );

  const deleteAnalysisInstance = useCallback((uuid: string) => {
    setAnalysisInstanceList((prevList) =>
      prevList.filter((instance) => instance.uuid !== uuid),
    );
  }, []);

  const dispatchers = useMemo(
    () => ({
      updateAnalysisInstance,
      addAnalysisInstance,
      deleteAnalysisInstance,
    }),
    [updateAnalysisInstance, addAnalysisInstance, deleteAnalysisInstance],
  );

  return (
    <AnalysisInstanceListDispatchersContext.Provider value={dispatchers}>
      <AnalysisInstanceListContext.Provider value={analysisInstanceList}>
        {children}
      </AnalysisInstanceListContext.Provider>
    </AnalysisInstanceListDispatchersContext.Provider>
  );
}
