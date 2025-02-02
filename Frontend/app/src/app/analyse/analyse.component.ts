import {Component, OnInit} from '@angular/core';
import {HttpErrorResponse} from "@angular/common/http";
import {Router} from "@angular/router";
import {ClrLoadingState} from '@clr/angular';
import {ImageService} from "../services/image.service";
import {delay} from "rxjs";

@Component({
  selector: 'app-analyse',
  templateUrl: 'analyse.component.html',
  styleUrls: ['analyse.component.scss']
})
export class AnalyseComponent implements OnInit {
  // @ts-ignore
  imageToAnalyse: Image;
  submitBtnState: ClrLoadingState = ClrLoadingState.DEFAULT;

  constructor(private imageService: ImageService, private router: Router) {
  }

  ngOnInit(): void {
  }

  async submitDemo() {
    this.submitBtnState = ClrLoadingState.LOADING;
    setTimeout(() => this.submitBtnState = ClrLoadingState.DEFAULT, 1000);
  }

  navigateToResultView(id: number) {
    this.router.navigateByUrl(`/result/${id}`);
  }

  getImageByUrl(privacyImg: string) {
    let inputEl = document.querySelectorAll<HTMLElement>('#url')[0];
    // @ts-ignore
    let url = inputEl.value;
    if (url != '')
      this.submitDemo();

    let privateImg = privacyImg.toLowerCase()==="private";

    this.imageService.addImage(url,privateImg).subscribe(
      (image) => {
        this.imageToAnalyse = image;
        const successEl = document.querySelectorAll<HTMLElement>('.alert-success')[0];
        successEl.classList.add('active');
        const dangerEl = document.querySelectorAll<HTMLElement>('.alert-danger')[0];
        dangerEl.classList.remove('active');
        // @ts-ignore
        inputEl.value = '';
      },
      (error: HttpErrorResponse) => {
        const successEl = document.querySelectorAll<HTMLElement>('.alert-success')[0];
        successEl.classList.remove('active');
        const dangerEl = document.querySelectorAll<HTMLElement>('.alert-danger')[0];
        dangerEl.classList.add('active');
        // @ts-ignore
        inputEl.value = '';
      });
  }

  close() {
    const successEl = document.querySelectorAll<HTMLElement>('.alert-success')[0];
    if (successEl.classList.contains('active'))
      successEl.classList.remove('active');
    const dangerEl = document.querySelectorAll<HTMLElement>('.alert-danger')[0];
    if (dangerEl.classList.contains('active'))
      dangerEl.classList.remove('active');
  }

}
