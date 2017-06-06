import { Injectable } from '@angular/core';
import { Headers, Http, Response } from '@angular/http';
import 'rxjs/add/operator/toPromise';
import { PostAuth } from '../model/user';
import { Article } from '../model/article';

@Injectable()
export class ApiService {
  private headers = new Headers({ 'Content-Type': 'application/json' });
  private apiHost = 'http://localhost:8000/api/v1';

  constructor(private http: Http) { }

  private get(url: string) :Promise<Response> {
    return this.http.get(
      url,
      { headers: this.headers })
      .map((response: Response) => {
        return response
      }).toPromise();
  }

  private post(url: string, data: any) :Promise<Response> {
    return this.http.post(
      url,
      data,
      { headers: this.headers })
      .map((response: Response) => {
        return response
      }).toPromise();
  }

  postArticle(data: Article) :Promise<Response> {
    return this.post(
      this.apiHost + '/articles',
      JSON.stringify(data)
    )
  }

  getArticles() :Promise<Response> {
    return this.get(
      this.apiHost + '/articles',
    )
  }

  postSignup(data: PostAuth) :Promise<Response> {
    return this.post(
      this.apiHost + '/signup',
      JSON.stringify(data)
    )
  }

  postLogin(data: PostAuth) :Promise<Response> {
    return this.post(
      this.apiHost + '/login',
      JSON.stringify(data)
    )
  }

}
