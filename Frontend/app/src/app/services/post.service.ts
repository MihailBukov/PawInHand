import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Post } from '../models/post';

@Injectable({
  providedIn: 'root' // Ensures the service is available globally
})
export class PostService {
  private apiUrl = 'http:localhost:8080/api/posts'; // Base API URL

  constructor(private http: HttpClient) {}

  /** Fetch all posts */
  getAllPosts(): Observable<Post[]> {
    return this.http.get<Post[]>(this.apiUrl);
  }

  /** Fetch a post by ID */
  getPostById(postId: string): Observable<Post> {
    return this.http.get<Post>(`${this.apiUrl}/${postId}`);
  }

  /** Create a new post */
  createPost(post: Post): Observable<Post> {
    return this.http.post<Post>(this.apiUrl, post);
  }

  /** Update a post by ID */
  updatePostById(postId: string, post: Post): Observable<Post> {
    return this.http.put<Post>(`${this.apiUrl}/${postId}`, post);
  }

  /** Delete a post */
  deletePostById(postId: string): Observable<void> {
    return this.http.delete<void>(`${this.apiUrl}/${postId}`);
  }
}
