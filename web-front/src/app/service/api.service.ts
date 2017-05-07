import { Injectable } from '@angular/core';
import { Headers, Http, Response } from '@angular/http';
import 'rxjs/add/operator/toPromise';

@Injectable()
export class ApiService {
  private headers = new Headers({ 'Content-Type': 'application/json' });

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

  postArticle(data: any) :Promise<Response> {
    return this.post(
      'http://localhost:8000/api/v1/articles',
      data
    )
  }

  getArticles() :Promise<Response> {
    return this.get(
      'http://localhost:8000/api/v1/articles',
    )
  }

}
