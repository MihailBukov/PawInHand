import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable, tap} from "rxjs";
import {CreateUser, User} from "../objects/user";
import {UserService} from "./user.service";

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  constructor(private http: HttpClient, private userService:UserService) {}

  private authUrl = 'http://localhost:8080';

  // @ts-ignore
  user: CreateUser;

  register(username: string, email: string, password: string): Observable<User> {
    this.user.password=password;
    this.user.username=username;
    this.user.email=email;
    // @ts-ignore
    return userSevice.addUser(this.user);
  }

  logout(): void {
    localStorage.removeItem('userToken');
  }
  login(userId:number): void {
    localStorage.setItem('userToken',userId.toString());
  }

  isLoggedIn(): boolean {
    const token = localStorage.getItem('userToken');
    return !!token;
  }

}

