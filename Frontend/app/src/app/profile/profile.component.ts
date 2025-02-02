import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup ,FormBuilder} from '@angular/forms';
import {User} from "../objects/user";
import {UserService} from "../services/user.service";
import {AuthService} from "../services/auth.service";
import {Notification} from "../objects/notification";
import { Router} from "@angular/router";
import {NotificationService} from "../services/notification.service";
import { ClarityIcons, bellIcon,userIcon,cogIcon ,imageGalleryIcon} from '@cds/core/icon';

ClarityIcons.addIcons(bellIcon,userIcon,cogIcon,imageGalleryIcon)

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.scss'],
})
export class ProfileComponent implements OnInit {
  currentUser: User|undefined;
  notifications: Notification[] = [];
  selectedFile: File | null = null;

  protected readonly form = new FormGroup({
    files: new FormControl<FileList | null>(null),
  });

  constructor(private userService: UserService,
              private authService:AuthService,
              private router: Router,
              private notificationService:NotificationService,
  ) {}

  ngOnInit(): void {
    this.getCurrentUser()
    this.getCurrentUserNotifications()
  }

  getCurrentUser(): void {
    const userToken = localStorage.getItem('userToken') || '';
    this.userService.getUserById(Number.parseInt(userToken)).subscribe(
      (user: User) => {
        this.currentUser = user;
      }
    );
  }

  getCurrentUserNotifications(): void {
    this.notificationService.getNotificationsByUserId(Number.parseInt(<string>localStorage.getItem("userToken"))).subscribe(
      (notifications: Notification[]) => {
        this.notifications = notifications;
      }
    );
  }

  logout(){
    this.authService.logout()
    this.router.navigateByUrl('');
  }

  goToImage(imgId:number){
    this.router.navigateByUrl(`/result/${imgId}`);
  }

  onFileSelected(event: Event): void {
    const input = event.target as HTMLInputElement;
    if (input?.files?.length) {
      this.selectedFile = input.files[0];
      console.log('Selected file:', this.selectedFile);
    }
  }

  saveProfilePicture(): void {
    if (this.selectedFile && this.currentUser) {
      const reader = new FileReader();

      // Step 1: Convert file to Base64 string
      reader.onload = () => {
        const base64Image = reader.result as string||null;

        if (this.currentUser && base64Image) {
          this.currentUser.picture = base64Image;

          this.userService.updateUser(this.currentUser).subscribe(
            (updatedUser: User) => {
              console.log(updatedUser)
              console.log(this.currentUser)

              // this.currentUser = updatedUser;
              console.log('User profile updated successfully with Base64 image.');
            },
            (error) => {
              console.error('Failed to update user profile:', error);
            }
          );
        }
      }
      reader.readAsDataURL(this.selectedFile);
    } else {
      console.error('No file selected or user data is undefined!');
    }
  }


}
