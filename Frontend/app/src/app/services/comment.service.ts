import {Injectable} from "@angular/core";
import {environment} from "../environments/environment";
import {HttpClient} from "@angular/common/http";
import {Image} from "../objects/image";
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class CommentService {
  private apiServerUrl = environment.apiCommentUrl;

  constructor(private http: HttpClient) {
  }

  public addComments(comment: Comment): Observable<ArrayBuffer> {
    return this.http.post<ArrayBuffer>(`${this.apiServerUrl}`,comment);
  }
}
