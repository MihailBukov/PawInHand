import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { VolunteerWork } from '../models/volunteer-work';

@Injectable({
  providedIn: 'root' // Ensures the service is available globally
})
export class VolunteerService {
  private apiUrl = 'http:localhost:8080/api/volunteer-work'; // Base API URL

  constructor(private http: HttpClient) {}

  /** Register for volunteer work */
  registerForVolunteerWork(workId: string, volunteer: VolunteerWork): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/${workId}/register`, volunteer);
  }
}
