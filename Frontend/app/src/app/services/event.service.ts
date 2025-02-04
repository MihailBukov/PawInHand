import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Event } from '../models/event';

@Injectable({
  providedIn: 'root' // Ensures the service is available globally
})
export class EventService {
  private apiUrl = 'http:localhost:8080/api/events'; // Base API URL

  constructor(private http: HttpClient) {}

  /** Fetch all events */
  getAllEvents(): Observable<Event[]> {
    return this.http.get<Event[]>(this.apiUrl);
  }

  /** Create a new event */
  createEvent(event: Event): Observable<Event> {
    return this.http.post<Event>(this.apiUrl, event);
  }
}