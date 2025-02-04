import {Image} from "./objects/image";
import {Component, OnInit, OnChanges, ChangeDetectorRef} from '@angular/core';
import {ImageService} from "./services/image.service";
import {HttpErrorResponse} from "@angular/common/http";
import '@cds/core/icon/register.js';
import {AuthService} from "./services/auth.service";
import {UserService} from "./services/user.service";
import {User} from "./objects/user";
import {NavigationEnd, Router} from "@angular/router";
import {filter} from "rxjs";

@Component({
  selector: 'app-root',
  templateUrl: 'app.component.html',
  styleUrls: ['app.component.scss']
})
export class AppComponent implements OnInit {
  title = 'ics';
  user: User | undefined

  constructor(private userService: UserService,
              public authService: AuthService,
              private cdr: ChangeDetectorRef,
              private router: Router
  ) {
  }

  ngOnInit(): void {
    this.checkPageActivity();
    this.getLoggedUser()

    this.router.events
      .subscribe(() => {
        this.onRouteChange();
      });
  }

  onRouteChange(): void {
    this.getLoggedUser();
    this.checkPageActivity();
  }

  getLoggedUser() {
    const userToken = localStorage.getItem("userToken");
    const id: number = Number.parseInt(userToken||"", 10);
    this.userService.getUserById(id).subscribe(
      (user: User) => {
        this.user = user;
        this.cdr.detectChanges(); // Trigger UI update
      },
      (error: HttpErrorResponse) => {
        console.error("Failed to fetch user:", error.message);
        this.user = undefined; // Clear user on error
        this.cdr.detectChanges();
      }
    );
  }

  checkPageActivity() {
    const currentUrl = this.router.url;

    const analyseBtn = document.querySelector<HTMLElement>('#analyse');
    const galleryBtn = document.querySelector<HTMLElement>('#gallery');
    const statisticsBtn = document.querySelector<HTMLElement>('#statistics');
    const profileBtn = document.querySelector<HTMLElement>('#profile');

    if (currentUrl.includes("images")) {
      analyseBtn?.classList.remove('active');
      galleryBtn?.classList.add('active');
      statisticsBtn?.classList.remove('active');
      profileBtn?.classList.remove('active');
    } else if(currentUrl.includes("analyse")){
      analyseBtn?.classList.add('active');
      galleryBtn?.classList.remove('active');
      statisticsBtn?.classList.remove('active');
      profileBtn?.classList.remove('active');
    } else if(currentUrl.includes("statistics")){
      analyseBtn?.classList.remove('active');
      galleryBtn?.classList.remove('active');
      statisticsBtn?.classList.add('active');
      profileBtn?.classList.remove('active');
    } else if(currentUrl.includes("profile")) {
      analyseBtn?.classList.remove('active');
      galleryBtn?.classList.remove('active');
      statisticsBtn?.classList.remove('active');
      profileBtn?.classList.add('active');
    }
  }

  changeActiveAnalyse() {
    const analyseBtn = document.querySelector<HTMLElement>('#analyse');
    const galleryBtn = document.querySelector<HTMLElement>('#gallery');
    const statisticsBtn = document.querySelector<HTMLElement>('#statistics');
    const profileBtn = document.querySelector<HTMLElement>('#profile');
    if (!analyseBtn?.classList.contains('active')) {
      analyseBtn?.classList.add('active');
      galleryBtn?.classList.remove('active');
      statisticsBtn?.classList.remove('active');
      profileBtn?.classList.remove('active');
    }

  }

  changeActiveGallery() {
    const analyseBtn = document.querySelector<HTMLElement>('#analyse');
    const galleryBtn = document.querySelector<HTMLElement>('#gallery');
    const statisticsBtn = document.querySelector<HTMLElement>('#statistics');
    const profileBtn = document.querySelector<HTMLElement>('#profile');
    if (!galleryBtn?.classList.contains('active')) {
      analyseBtn?.classList.remove('active');
      galleryBtn?.classList.add('active');
      statisticsBtn?.classList.remove('active');
      profileBtn?.classList.remove('active');
    }
  }

  changeActiveStatistics() {
    const analyseBtn = document.querySelector<HTMLElement>('#analyse');
    const galleryBtn = document.querySelector<HTMLElement>('#gallery');
    const statisticsBtn = document.querySelector<HTMLElement>('#statistics');
    const profileBtn = document.querySelector<HTMLElement>('#profile');
    if (!statisticsBtn?.classList.contains('active')) {
      analyseBtn?.classList.remove('active');
      galleryBtn?.classList.remove('active');
      statisticsBtn?.classList.add('active');
      profileBtn?.classList.remove('active');
    }
  }

  changeActiveProfile() {
    const analyseBtn = document.querySelector<HTMLElement>('#analyse');
    const galleryBtn = document.querySelector<HTMLElement>('#gallery');
    const statisticsBtn = document.querySelector<HTMLElement>('#statistics');
    const profileBtn = document.querySelector<HTMLElement>('#profile');
    if (!profileBtn?.classList.contains('active')) {
      analyseBtn?.classList.remove('active');
      galleryBtn?.classList.remove('active');
      statisticsBtn?.classList.remove('active');
      profileBtn?.classList.add('active');
    }
  }
}
