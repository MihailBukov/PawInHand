import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Ad } from '../models/ad';

@Injectable({
  providedIn: 'root' // Makes the service globally available
})
export class AdService {
  private apiUrl = 'http:localhost:8080/api/ads'; // Base API URL

  private httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  };

  constructor(private http: HttpClient) {}

  /** Get all ads */
  getAllAds(): Observable<Ad[]> {
    return this.http.get<Ad[]>(this.apiUrl);
  }

  /** Get a single ad by ID */
  getAdById(adId: string): Observable<Ad> {
    return this.http.get<Ad>(`${this.apiUrl}/${adId}`);
  }

  /** Create a new ad */
  createAd(ad: Ad): Observable<Ad> {
    return this.http.post<Ad>(this.apiUrl, ad, this.httpOptions);
  }

  /** Update an existing ad */
  updateAdById(adId: string, ad: Ad): Observable<Ad> {
    return this.http.put<Ad>(`${this.apiUrl}/${adId}`, ad, this.httpOptions);
  }

  /** Delete an ad by ID */
  deleteAdById(adId: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${adId}`);
  }
}