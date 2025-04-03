"use client";

import { HyperParametersContext } from "@/contexts/hyper-parameters";
import { useState } from "react";

type Props = {
  children: Readonly<React.ReactNode>;
};

export default function HyperParamtersProvider({ children }: Props) {
  const [hyperK, setHyperK] = useState(8);

  return (
    <HyperParametersContext.Provider
      value={{
        hyperK,
        setHyperK,
      }}
    >
      {children}
    </HyperParametersContext.Provider>
  );
}
