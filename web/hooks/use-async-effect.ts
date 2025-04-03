"use client";

import { DependencyList, useEffect } from "react";

export type IsCanceledFunction = () => boolean;

export default function useAsyncEffect(
  effect: (isCanceled: IsCanceledFunction) => Promise<void>,
  deps?: DependencyList,
) {
  return useEffect(() => {
    let canceled = false;
    effect(() => canceled);
    return () => {
      canceled = true;
    };
  }, deps);
}
