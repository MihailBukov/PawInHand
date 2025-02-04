import { NgModule,CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {ClarityModule} from "@clr/angular";
import {HttpClientModule} from "@angular/common/http";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {GalleryComponent} from "./gallery/gallery.component";
import {RouterModule} from "@angular/router";
import {ResultComponent} from "./result/result.component";
import {AnalyseComponent} from "./analyse/analyse.component";
import { LoginComponent } from './login/login.component';
import { RegistrationComponent } from './registration/registration.component';
import {StatisticsComponent} from "./statistics/statistics.component";
import {ProfileComponent} from "./profile/profile.component";

// @ts-ignore
@NgModule({
  declarations: [
    AppComponent,
    GalleryComponent,
    ResultComponent,
    AnalyseComponent,
    StatisticsComponent,
    ProfileComponent,
    LoginComponent,
    RegistrationComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    ClarityModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
  ],
  providers: [],
  bootstrap: [ AppComponent ],
  schemas: [CUSTOM_ELEMENTS_SCHEMA]

})
export class AppModule { }
