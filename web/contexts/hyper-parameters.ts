import { createContext } from "react";

export const HyperParametersContext = createContext<{
  hyperK: number;
  setHyperK: (hyperK: number) => void;
}>({
  hyperK: 2,
  setHyperK: () => {
    throw new Error("Invaild context");
  },
});
