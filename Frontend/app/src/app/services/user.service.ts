import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../models/user';

@Injectable({
  providedIn: 'root' // Ensures the service is available globally
})
export class UserService {
  private apiUrl = 'http:localhost:8080/api/users'; // Base API URL

  constructor(private http: HttpClient) {}

  /** Fetch all users */
  getAllUsers(): Observable<User[]> {
    return this.http.get<User[]>(this.apiUrl);
  }

  /** Fetch a user by ID */
  getUserById(userId: string): Observable<User> {
    return this.http.get<User>(`${this.apiUrl}/${userId}`);
  }

  /** Register a new user */
  registerUser(user: User): Observable<User> {
    return this.http.post<User>(this.apiUrl, user);
  }
}
