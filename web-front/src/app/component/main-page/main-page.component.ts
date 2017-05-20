import { Component, OnInit } from '@angular/core';
import { ApiService } from '../../service/api.service';
import { Article } from '../../model/article';

@Component({
  selector: 'app-main-page',
  templateUrl: './main-page.component.html',
  styleUrls: ['./main-page.component.scss']
})
export class MainPageComponent implements OnInit {

  private articles;

  constructor(private api:ApiService) { }

  ngOnInit() {
    this.articles = [
      {
        title: 'NyanNyan1',
        body:　'Neko. Neko. Neko',
        filename: '001.jpg'
      },
      {
        title: 'NyanNyan 2',
        body:　'Neko. Neko. Neko',
        filename: '002.jpg'
      },
      {
        title: 'NyanNyan 3',
        body:　'Neko. Neko. Neko',
        filename: '003.jpg'
      },
      {
        title: 'NyanNyan 4',
        body:　'Neko. Neko. Neko',
        filename: '004.jpg'
      },
      {
        title: 'NyanNyan 5',
        body:　'Neko. Neko. Neko',
        filename: '005.jpg'
      },
      {
        title: 'NyanNyan 6',
        body:　'Neko. Neko. Neko',
        filename: '006.jpg'
      },
      {
        title: 'NyanNyan 7',
        body:　'Neko. Neko. Neko',
        filename: '007.jpg'
      },
    ]

    this.api.getArticles()
      .then(value => {
        this.articles = value.json();
      })
      .catch(err => {
        console.error(err);
      });
  }

}
