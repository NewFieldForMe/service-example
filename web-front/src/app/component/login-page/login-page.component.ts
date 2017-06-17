import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { ApiService } from '../../service/api.service';
import { PostAuth } from '../../model/user';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.scss']
})
export class LoginPageComponent implements OnInit {

  postForm: FormGroup;
  processing: Boolean;

  constructor(
    private api: ApiService,
    private fb: FormBuilder,
    private router: Router,
  ) {
    this.postForm = this.fb.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
  }

  ngOnInit() {
  }

  onSubmit() {
    this.processing = true;
    let auth = this.preparePost();
    this.api.postLogin(auth)
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

  preparePost(): PostAuth {
    const formModel = this.postForm.value;
    let auth = new PostAuth();
    auth.username = formModel.username;
    auth.password = formModel.password;
    return auth;
  }

}
