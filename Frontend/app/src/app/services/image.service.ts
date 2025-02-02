import {Injectable} from "@angular/core";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Image} from "../objects/image";
import {environment} from "../environments/environment";
import {TopTag} from "../objects/topTag";
import {TopActiveUser} from "../objects/topActiveUser";

@Injectable({
  providedIn: 'root'
})
export class ImageService {
  private apiServerUrl = environment.apiImagesUrl;

  constructor(private http: HttpClient) {
  }

  public getImages() {
    return this.http.get<Image[]>(`${this.apiServerUrl}`);
  }

  public getImagesByTag(tags: string[]): Observable<Image[]> {
    return this.http.get<Image[]>(`${this.apiServerUrl}?tags=${tags}`);
  }

  public getImagesByUserId(userId: string): Observable<Image[]> {
    return this.http.get<Image[]>(`${this.apiServerUrl}/user?userId=${userId}`);
  }

  public getImageById(id:number): Observable<Image> {
    return this.http.get<Image>(`${this.apiServerUrl}/id?id=${id}`);
  }

  public updateImage(image:Image): Observable<Image> {
    return this.http.put<Image>(`${this.apiServerUrl}/update`,image);
  }


  public addImage(imageUrl: string, privateImg: boolean): Observable<ArrayBuffer> {
    const body = {
      imgUrl: imageUrl,
      token: localStorage.getItem('userToken') || '' ,
      privateImg: privateImg,
    };
    return this.http.post<ArrayBuffer>(`${this.apiServerUrl}`,body);
  }

  public deleteImage(imageId: number): Observable<void> {
    return this.http.delete<void>(`${this.apiServerUrl}/delete/${imageId}`);
  }

  public getTopLikedImages(): Observable<Image[]> {
    return this.http.get<Image[]>(`${this.apiServerUrl}/top-liked`);
  }

  public getTopUsedTags(): Observable<TopTag[]> {
    return this.http.get<TopTag[]>(`${this.apiServerUrl}/top-tags`);
  }

  public getTopActiveUsers(): Observable<TopActiveUser[]> {
    return this.http.get<TopActiveUser[]>(`${this.apiServerUrl}/top-users`);
  }
}
