import { Source } from "./source";
import { User } from "./user";

type Credentials = Source & {
  data: User;
};

export type { Credentials };
