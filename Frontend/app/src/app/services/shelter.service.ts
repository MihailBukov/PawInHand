import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Shelter } from '../models/shelter';

@Injectable({
  providedIn: 'root' // Ensures the service is available globally
})
export class ShelterService {
  private apiUrl = 'http:localhost:8080/api/shelters'; // Base API URL

  constructor(private http: HttpClient) {}

  /** Fetch all shelters */
  getAllShelters(): Observable<Shelter[]> {
    return this.http.get<Shelter[]>(this.apiUrl);
  }

  /** Fetch a shelter by ID */
  getShelterById(shelterId: string): Observable<Shelter> {
    return this.http.get<Shelter>(`${this.apiUrl}/${shelterId}`);
  }

  /** Register a new shelter */
  registerShelter(shelter: Shelter): Observable<Shelter> {
    return this.http.post<Shelter>(this.apiUrl, shelter);
  }

  /** Update a shelter by ID */
  updateShelterById(shelterId: string, shelter: Shelter): Observable<Shelter> {
    return this.http.put<Shelter>(`${this.apiUrl}/${shelterId}`, shelter);
  }

  /** Delete a shelter */
  deleteShelterById(shelterId: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${shelterId}`);
  }
}
