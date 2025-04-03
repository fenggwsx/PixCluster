import { DataURL } from "@/types";

export function getBlobDataURL(blob: Blob): Promise<DataURL> {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(blob);
    reader.onload = () => {
      const dataUrl = reader.result as string;
      const commaIndex = dataUrl.indexOf(",");
      if (commaIndex === -1)
        reject("Invalid data URL format: missing comma separator");
      else
        resolve({
          prefix: dataUrl.substring(0, commaIndex),
          data: dataUrl.substring(commaIndex + 1),
        });
    };
    reader.onerror = (e) => reject(e);
  });
}
