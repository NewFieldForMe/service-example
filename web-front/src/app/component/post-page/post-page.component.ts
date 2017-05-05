import { Component, OnInit } from '@angular/core';
import { Article } from '../../model/article';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'app-post-page',
  templateUrl: './post-page.component.html',
  styleUrls: ['./post-page.component.scss']
})
export class PostPageComponent implements OnInit {

  postForm: FormGroup;

  constructor(private fb: FormBuilder) {
    this.createForm();
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
    let article = this.preparePost();
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
