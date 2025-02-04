import { Injectable } from '@angular/core';
import { HttpClient, HttpEvent, HttpHeaders, HttpRequest } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root' // Ensures the service is available globally
})
export class MediaService {
  private apiUrl = 'http:localhost:8080/api/upload/picture'; // Base API URL

  constructor(private http: HttpClient) {}

  /** Upload a picture */
  uploadPicture(file: File): Observable<HttpEvent<any>> {
    const formData: FormData = new FormData();
    formData.append('file', file);

    // Create the request with progress tracking
    const req = new HttpRequest('POST', this.apiUrl, formData, {
      reportProgress: true,
      responseType: 'json'
    });

    return this.http.request(req);
  }
}
