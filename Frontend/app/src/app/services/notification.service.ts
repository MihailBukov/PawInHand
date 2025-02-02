import {Injectable} from "@angular/core";
import {environment} from "../environments/environment";
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Notification} from "../objects/notification"
import {Image} from "../objects/image";

@Injectable({
  providedIn: 'root'
})
export class NotificationService {
  private apiServerUrl = environment.apiNotificationUrl;

  constructor(private http: HttpClient) {
  }

  public getNotificationsByUserId(userId: number): Observable<Notification[]> {
    return this.http.get<Notification[]>(`${this.apiServerUrl}/userId?userId=${userId}`);
  }

  public addNotification(notification: {
    whoseNotification: number | undefined;
    whoseAction: string;
    date: Date;
    action: string;
    imgId: number | undefined
  }): Observable<Notification> {
    return this.http.post<Notification>(`${this.apiServerUrl}`,notification);
  }

}
