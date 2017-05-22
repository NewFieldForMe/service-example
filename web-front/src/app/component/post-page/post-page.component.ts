import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { Article } from '../../model/article';
import { ApiService } from '../../service/api.service';

@Component({
  selector: 'app-post-page',
  templateUrl: './post-page.component.html',
  styleUrls: ['./post-page.component.scss']
})
export class PostPageComponent implements OnInit {

  postForm: FormGroup;
  processing: Boolean;

  constructor(
    private api: ApiService,
    private fb: FormBuilder,
    private router: Router,
    ) {
    this.createForm();
    this.processing = false;
  }

  createForm() {
    this.postForm = this.fb.group({
      title: ['', Validators.required],
      body: ['', Validators.required],
      filedata: ['', Validators.required],
      filename: ['', Validators.required],
    });
  }

  ngOnInit() {
  }

  changeListener($event): void {
    this.readThis($event.target);
  }

  readThis(inputValue: any): void {
    if (!inputValue.files[0]) { return }
    var file: File = inputValue.files[0];
    var myReader: FileReader = new FileReader();

    myReader.onloadend = (e) => {
      this.postForm.get('filedata').setValue(myReader.result);
      this.postForm.get('filename').setValue(file.name);
    }
    myReader.readAsDataURL(file);
  }

  onSubmit() {
    this.processing = true;
    let article = this.preparePost();
    this.api.postArticle(JSON.stringify(article))
      .then(value => {
        console.log(value);
        this.processing = false;
        this.router.navigate(['mainpage']);
      })
      .catch(err => {
        this.processing = false;
        console.log(err);
      });
  }

  preparePost(): Article {
    const formModel = this.postForm.value;
    let article = new Article();
    article.title = formModel.title;
    article.body = formModel.body;
    article.filedata = formModel.filedata
    article.filename = formModel.filename
    return article
  }
}
