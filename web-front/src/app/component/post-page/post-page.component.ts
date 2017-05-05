import { Component, OnInit } from '@angular/core';
import { Article } from '../../model/article';

@Component({
  selector: 'app-post-page',
  templateUrl: './post-page.component.html',
  styleUrls: ['./post-page.component.scss']
})
export class PostPageComponent implements OnInit {

  article: Article;

  constructor() { }

  ngOnInit() {
    this.article = new Article();
  }

  changeListener($event): void {
    this.readThis($event.target);
  }

  readThis(inputValue: any): void {
    var file: File = inputValue.files[0];
    var myReader: FileReader = new FileReader();

    myReader.onloadend = (e) => {
      this.article.filedata = myReader.result;
      this.article.filename = file.name;
    }
    myReader.readAsDataURL(file);
  }

}
