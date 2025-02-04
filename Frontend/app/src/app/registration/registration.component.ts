import { Component } from '@angular/core';
import {ClrLoadingState} from "@clr/angular";
import {CreateUser, User} from "../objects/user";
import {Router} from "@angular/router";
import {UserService} from "../services/user.service";
import {AuthService} from "../services/auth.service";
import * as jwt from "jsonwebtoken";
import config from "../config/index";

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.scss']
})
export class RegistrationComponent {

  // @ts-ignore
  user:User;
  submitBtnState: ClrLoadingState = ClrLoadingState.DEFAULT;

  constructor(private userService: UserService, private router: Router,private authService: AuthService) {
  }
  async submitDemo() {
    this.submitBtnState = ClrLoadingState.LOADING;
    setTimeout(() => this.submitBtnState = ClrLoadingState.DEFAULT, 1000);

  }
  onSubmit() {
    // @ts-ignore
    const username = document.getElementById<HTMLInputElement>("username");
    // @ts-ignore
    const email = document.getElementById<HTMLInputElement>("email");
    // @ts-ignore
    const password = document.getElementById<HTMLInputElement>("password");
    const informDiv = document.getElementById("inform");

    if (!username || !email || !password || !informDiv) {
      console.error('One of the form elements is missing.');
      return; // Early exit if any form element is missing
    }

    this.submitDemo();
    // @ts-ignore
    this.userService.getUserByUsername(username.value).subscribe({
      next: (user: User) => {
        if(user!=null) {
          informDiv.innerHTML = `<p>User already exists</p>`;
        }
        else{
          const newUser:CreateUser = {
            // @ts-ignore
            username: username.value,
            // @ts-ignore
            email: email.value,
            // @ts-ignore
            password: password.value,
          };
          this.userService.addUser(newUser).subscribe({
            next: (response) => {
              this.router.navigateByUrl(`/`);
            },
            error: (error) => {
              informDiv.innerHTML = `<p>Registration failed</p>`;
            }
          });
        }
      },
      error: (err) => {
          informDiv.innerHTML = `<p>Error checking user existence</p>`;
      }
    });
  }

}
