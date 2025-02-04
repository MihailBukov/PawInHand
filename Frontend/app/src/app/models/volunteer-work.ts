import { VolunteerWorkLocation } from "./volunteer-work-location";

export interface VolunteerWork {
    id?: string;
    userId?: string;
    title?: string;
    description?: string;
    date?: string;
    location?: VolunteerWorkLocation;
  }