import { Component } from '@angular/core';
import {UserService} from "../services/user.service";
import {Router} from "@angular/router";
import {AuthService} from "../services/auth.service";
import {ClrLoadingState} from "@clr/angular";
import {CreateUser, User} from "../objects/user";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {

  submitBtnState: ClrLoadingState = ClrLoadingState.DEFAULT;
  // @ts-ignore
  user:CreateUser;
  constructor(private userService: UserService, private router: Router,private authService: AuthService) {
  }

  async submitDemo() {
    this.submitBtnState = ClrLoadingState.LOADING;
    setTimeout(() => this.submitBtnState = ClrLoadingState.DEFAULT, 1000);

  }
  onSubmit() {

    let username = document.querySelector<HTMLElement>("#username");
    let password = document.querySelector<HTMLElement>("#password");
    let informDiv = document.getElementById("inform");

    this.submitDemo();

    // @ts-ignore
    this.userService.getUserByUsername(username.value).subscribe({
      next: (user: User) => {
        this.user = user;
      },
      error: (error) => {
        console.error('Error fetching user data:', error);
      }
    });

    // @ts-ignore
    this.userService.getUserByUsername(username.value).subscribe({
      next: (user: User) => {
        // Assuming the 'user' object has a 'password' property to compare with
        // @ts-ignore
        if (user==null||password.value !== user.password) {
          // @ts-ignore
          informDiv.innerHTML = `<p>Invalid username or password</p>`;
        } else {
          if(user.id) {
            this.authService.login(user?.id)
            // Perform login and navigate
            this.router.navigateByUrl(`/analyse`);
          }
        }
      }
    })
  }
}
