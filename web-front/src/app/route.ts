import { RouterModule, Routes } from '@angular/router';

import { MainPageComponent } from './component/main-page/main-page.component';
import { LoginPageComponent } from './component/login-page/login-page.component';
import { PostPageComponent } from './component/post-page/post-page.component';
import { UserPageComponent } from './component/user-page/user-page.component';
import { ArticlePageComponent } from './component/article-page/article-page.component';

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
    path: 'user',
    component: UserPageComponent,
  },
  {
    path: 'article/:id',
    component: ArticlePageComponent,
  },
  { 
    path: '',
    redirectTo: '/mainpage',
    pathMatch: 'full'
  },
  { path: '**', component: MainPageComponent}
];