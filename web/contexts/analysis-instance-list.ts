import { createContext } from "react";
import type { AnalysisInstance } from "@/types";

export const AnalysisInstanceListContext = createContext<AnalysisInstance[]>(
  [],
);

export const AnalysisInstanceListDispatchersContext = createContext<{
  updateAnalysisInstance: (newInstance: AnalysisInstance) => void;
  addAnalysisInstance: (newInstance: AnalysisInstance) => void;
  deleteAnalysisInstance: (uuid: string) => void;
}>({
  updateAnalysisInstance: () => {
    throw new Error("Invaild context");
  },
  addAnalysisInstance: () => {
    throw new Error("Invaild context");
  },
  deleteAnalysisInstance: () => {
    throw new Error("Invaild context");
  },
});
