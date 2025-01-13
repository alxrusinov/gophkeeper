import { Source } from "./source";

type Binary = Source & {
  mimeType: string;
  fileID: string;
};

type BinaryUpload = Source & {
  mimeType: string;
  data: Blob;
};

export type { Binary, BinaryUpload };
