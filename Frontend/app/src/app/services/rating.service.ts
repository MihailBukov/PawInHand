import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Rating } from '../models/rating';

@Injectable({
  providedIn: 'root' // Ensures the service is available globally
})
export class RatingService {
  private apiUrl = 'http:localhost:8080/api/users'; // Base API URL

  constructor(private http: HttpClient) {}

  /** Rate a user */
  rateUser(userId: string, rating: Rating): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/${userId}/rate`, rating);
  }
}
