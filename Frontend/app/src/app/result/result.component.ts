import { Component, OnInit } from '@angular/core';
import { ImageService } from "../services/image.service";
import { UserService } from "../services/user.service";
import { Image } from "../objects/image";
import { ActivatedRoute, Router } from "@angular/router";
import '@cds/core/icon/register.js';
import { ClarityIcons, thumbsUpIcon, thumbsDownIcon, talkBubblesIcon, downloadIcon, trashIcon, lockIcon, unlockIcon } from '@cds/core/icon';
import { User } from "../objects/user";
import { switchMap } from "rxjs";
import { NotificationService } from "../services/notification.service";

ClarityIcons.addIcons(thumbsUpIcon, thumbsDownIcon, talkBubblesIcon, downloadIcon, trashIcon, lockIcon, unlockIcon);

@Component({
  selector: 'app-result',
  templateUrl: 'result.component.html',
  styleUrls: ['result.component.scss']
})
export class ResultComponent implements OnInit {
  imageToAnalyse: Image | undefined;
  isCommentPopupVisible = false;
  newComment = '';
  private lastCommentId = 1;
  isDeletePopupVisible: boolean = false;
  imageUrlToDelete: string | undefined;
  currentUser: string = '';
  userPictures: Map<string, string> = new Map<string, string>();
  isLockUnlockPopupVisible = false;

  constructor(
    private imageService: ImageService,
    private userService: UserService,
    private route: ActivatedRoute,
    private router: Router,
    private notificationService: NotificationService,
  ) {}

  ngOnInit(): void {
    const imageId = this.route.snapshot.params['id'];
    this.getImageById(imageId);
    this.getCurrentUser();
  }

  navigateToResultView(tag: string): void {
    this.router.navigateByUrl(`/images/${tag}`);
  }

  getCurrentUser(): void {
    const userToken = localStorage.getItem('userToken') || '';
    this.userService.getUserById(Number.parseInt(userToken)).subscribe(
      (user: User) => {
        this.currentUser = user.username;
      }
    );
  }

  getImageById(id: number): void {
    this.imageService.getImageById(id).subscribe(
      (image) => {
        this.imageToAnalyse = image;
        this.preloadUserPictures();
      },
      (error) => {
        console.error('Error fetching image:', error);
      }
    );
  }

  preloadUserPictures(): void {
    if (this.imageToAnalyse?.comments) {
      const uniqueAuthors = new Set(this.imageToAnalyse.comments.map(comment => comment.author));

      uniqueAuthors.forEach(author => {
        this.userService.getUserByUsername(author).subscribe(
          (user: User) => {
            this.userPictures.set(author, user.picture || '../../assets/default.jpg');
          }
        );
      });
    }
  }

  likeImageWithId(id: number | undefined): void {
    if (id !== undefined) {
      this.imageService.getImageById(id).subscribe(
        (image) => {
          let usersRated = image.usersRated ? image.usersRated.split(',') : [];
          if (!usersRated.includes(this.currentUser)) {
            usersRated.push(this.currentUser);
            image.usersRated = usersRated.join(',');
            image.likes += 1;
            this.imageService.updateImage(image).subscribe(
              (updatedImage: Image) => this.imageToAnalyse = updatedImage,
              (error) => console.error('Error updating image:', error)
            );
            this.notificationService.addNotification({
              whoseNotification: this.imageToAnalyse?.user?.id,
              whoseAction: this.currentUser,
              date: new Date(),
              action: "like your image.",
              imgId: this.imageToAnalyse?.id
            }).subscribe();
          }
        },
        (error) => {
          console.error('Error liking the image:', error);
        }
      );
    } else {
      console.error('Image ID is undefined');
    }
  }

  dislikeImageWithId(id: number | undefined): void {
    if (id === undefined) {
      console.error('Image ID is undefined');
      return;
    }

    this.imageService.getImageById(id).subscribe(
      (image) => {
        const usersRated = image.usersRated ? image.usersRated.split(',').filter(Boolean) : [];
        if (!usersRated.includes(this.currentUser)) {
          image.dislikes += 1;
          usersRated.push(this.currentUser);
          image.usersRated = usersRated.join(',');

          this.imageService.updateImage(image).subscribe(
            (updatedImage: Image) => {
              this.imageToAnalyse = updatedImage;
            },
            (updateError) => {
              console.error('Error updating image:', updateError);
            }
          );
          this.notificationService.addNotification({
            whoseNotification: this.imageToAnalyse?.user?.id,
            whoseAction: this.currentUser,
            date: new Date(),
            action: "don't like your image.",
            imgId: this.imageToAnalyse?.id
          }).subscribe();
        } else {
          console.log('User has already rated this image.');
        }
      },
      (fetchError) => {
        console.error('Error fetching image:', fetchError);
      }
    );
  }

  openCommentPopup(): void {
    this.isCommentPopupVisible = true;
  }

  saveComment(id: number | undefined): void {
    if (!this.newComment.trim()) {
      alert('Comment cannot be empty.');
      return;
    }

    const currentTime = new Date();

    if (id === undefined) {
      console.error('Image ID is undefined');
      return;
    }

    this.imageService.getImageById(id).pipe(
      switchMap((image) => {
        if (!image.comments) {
          image.comments = [];
        }

        this.lastCommentId += 1;

        image.comments.push({
          id: this.lastCommentId,
          commentText: this.newComment.trim(),
          author: this.currentUser,
          date: currentTime,
        });

        image.comments = image.comments.sort(
          (a, b) => new Date(b.date).getTime() - new Date(a.date).getTime()
        );

        return this.imageService.updateImage(image);
      })
    ).subscribe(
      (updatedImage: Image) => {
        this.imageToAnalyse = updatedImage;
        this.dismissCommentPopup();
      },
      (error) => {
        console.error('Error saving comment:', error);
      }
    );

    this.notificationService.addNotification({
      whoseNotification: this.imageToAnalyse?.user?.id,
      whoseAction: this.currentUser,
      date: new Date(),
      action: `comment '${this.newComment.trim()}' on your image.`,
      imgId: this.imageToAnalyse?.id
    }).subscribe();
  }

  dismissCommentPopup(): void {
    this.isCommentPopupVisible = false;
    this.newComment = '';
  }

  confirmDelete(imageUrl: string | undefined): void {
    if (imageUrl) {
      this.imageUrlToDelete = imageUrl;
      this.isDeletePopupVisible = true;
    }
  }

  deleteImage(id: number | undefined): void {
    if (id) {
      this.imageService.deleteImage(id).subscribe(
        () => {
          this.router.navigateByUrl(`/images`);
        },
        (error) => {
          console.error('Error deleting the image:', error);
        }
      );
    } else {
      console.error('Image ID is undefined');
    }
  }

  dismissDeletePopup(): void {
    this.isDeletePopupVisible = false;
    this.imageUrlToDelete = undefined;
  }
  
  openLockUnlockPopup(): void {
    this.isLockUnlockPopupVisible = true;
  }

  dismissLockUnlockPopup(): void {
    this.isLockUnlockPopupVisible = false;
  }

  confirmPrivacyChange(): void {
    if (this.imageToAnalyse) {
      this.toggleLock(this.imageToAnalyse.id);
      this.dismissLockUnlockPopup();
    }
  }

  toggleLock(imageId: number): void {
    if (this.imageToAnalyse) {
      this.imageToAnalyse.privateImg = !this.imageToAnalyse.privateImg;
      this.imageService.updateImage(this.imageToAnalyse).subscribe(
        (updatedImage) => {
          this.imageToAnalyse = updatedImage;
        },
        (error) => {
          console.error('Error updating image privacy:', error);
        }
      );
    }
  }
}
