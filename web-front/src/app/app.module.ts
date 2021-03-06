import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '@angular/material'
import { ReactiveFormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { MainPageComponent } from './component/main-page/main-page.component';
import { LoginPageComponent } from './component/login-page/login-page.component';
import { NavigationBarComponent } from './component/navigation-bar/navigation-bar.component';
import { PostPageComponent } from './component/post-page/post-page.component';
import { UserPageComponent } from './component/user-page/user-page.component';
import { ArticlePageComponent } from './component/article-page/article-page.component';

import { ApiService } from './service/api.service';
import { SessionService } from './service/session.service';

import { appRoutes } from './route';
import { ArticleTileComponent } from './component/main-page/article-tile/article-tile.component';
import { SignupPageComponent } from './component/signup-page/signup-page.component';

@NgModule({
  declarations: [
    AppComponent,
    MainPageComponent,
    LoginPageComponent,
    NavigationBarComponent,
    PostPageComponent,
    UserPageComponent,
    ArticlePageComponent,
    ArticleTileComponent,
    SignupPageComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    ReactiveFormsModule,
    HttpModule,
    MaterialModule,
    RouterModule.forRoot(appRoutes)
  ],
  providers: [
    ApiService,
    SessionService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
