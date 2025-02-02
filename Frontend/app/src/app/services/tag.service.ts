import {Injectable} from "@angular/core";
import {HttpClient, HttpParams} from "@angular/common/http";
import {Observable} from "rxjs";
import {Tag} from "../objects/tag";
import {environment} from "../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class TagService {
  private apiServerUrl = environment.apiTagsUrl;

  constructor(private http: HttpClient) {
  }

  public getTags(): Observable<Tag[]> {
    return this.http.get<Tag[]>(`${this.apiServerUrl}`);
  }

  public getTagWithString (str: String): Observable<Tag[]> {
    let queryParams = new HttpParams();
    return this.http.get<Tag[]>(`${this.apiServerUrl}/str`, {params:queryParams});
  }

}
