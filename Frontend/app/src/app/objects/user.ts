import {Image} from "./image";

export interface User{
  id?:number;
  username:string;
  email:string;
  password:string;
  images?:Image[];
  picture:string;
}

export type CreateUser = Partial<User>;
