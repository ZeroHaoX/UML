import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { HttpHeaders,HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/services/login.service';
import User from '../../guards/usermodel'
import { UrlResolver } from '@angular/compiler';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(
    private fb: FormBuilder,
    private http:HttpClient,
    private router:Router,
    private loginServic:LoginService
  ) { }

  ngOnInit() {
    this.validateForm = this.fb.group({
      userName: [null, [Validators.required]],
      password: [null, [Validators.required]],
    });
  }


  validateForm: FormGroup;

  submitForm(): void {
    for (const i in this.validateForm.controls) {
      this.validateForm.controls[i].markAsDirty();
      this.validateForm.controls[i].updateValueAndValidity();
    }
    let userName=this.validateForm.get("userName").value
    let password=this.validateForm.get("password").value
    console.log(userName)
    console.log(password)
    this.loginServic.login(userName,password).subscribe(
      (resp)=>{
        // console.log(resp.data)
        if(resp.status==0){
          User.userName=resp.data.username
          User.role=resp.data.role
          User.permissions=resp.data.permissions
          // User.Password=resp.data.password
          console.log(resp)
          this.router.navigate(['/goods'])
        }else{
          return
        }

      }
    )
    // this.loginServic.login().subscribe
  }

}
