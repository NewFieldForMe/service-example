import { RouterModule, Routes } from '@angular/router';

import { MainPageComponent } from './component/main-page/main-page.component';
import { LoginPageComponent } from './component/login-page/login-page.component';
import { PostPageComponent } from './component/post-page/post-page.component';
import { UserPageComponent } from './component/user-page/user-page.component';
import { ArticlePageComponent } from './component/article-page/article-page.component';
import { SignupPageComponent } from './component/signup-page/signup-page.component';

export const appRoutes: Routes = [
  {
    path: 'mainpage',
    component: MainPageComponent,
  },
  {
    path: 'login',
    component: LoginPageComponent,
  },
  {
    path: 'post',
    component: PostPageComponent,
  },
  {
    path: 'user/:id',
    component: UserPageComponent,
  },
  {
    path: 'article/:id',
    component: ArticlePageComponent,
  },
  {
    path: 'signup',
    component: SignupPageComponent,
  },
  { 
    path: '',
    redirectTo: '/mainpage',
    pathMatch: 'full'
  },
  { path: '**', component: MainPageComponent}
];