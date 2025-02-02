import {Tag} from "./tag";
import {User} from "./user";
import {Comment} from "./comment";

export interface Image{
  id:number;
  height:number;
  service:string;
  uploadTime:Date;
  url:string;
  width:number;
  tags: Tag[];
  user: User;
  privateImg:boolean;
  likes:number;
  dislikes:number;
  comments:Comment[];
  usersRated:string;
}
