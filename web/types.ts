export type DataURL = {
  prefix: string;
  data: string;
};

export type Text2ImagePrompt = {
  positive: string;
  negative: string;
};

export type AnalysisResult = {
  color: string;
  count: number;
}[];

export type AnalysisInstance = {
  uuid: string;
  hyperK: number;
  prompt: Text2ImagePrompt | null;
  taskId: string | null;
  image: DataURL | null;
  result: AnalysisResult | null;
  summary: string | null;
};

export type ResponseWrapper<T> =
  | {
      success: false;
      code: number;
      message: string;
    }
  | {
      success: true;
      code: number;
      message: string;
      data: T;
    };
