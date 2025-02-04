import { ShelterAddress } from "./shelter-address";

export interface Shelter {
    id?: string;
    name?: string;
    address?: ShelterAddress;
    phone?: string;
    email?: string;
  }