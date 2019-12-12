import { Component, OnInit } from '@angular/core';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';
import { HttpHeaders,HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/services/login.service';
import User from '../../guards/usermodel'
import { UrlResolver } from '@angular/compiler';
import { NzMessageService } from 'ng-zorro-antd';

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
    private loginServic:LoginService,
    private message:NzMessageService,
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
        if(resp.status==0&&resp.rowCount>0){
          User.userName=resp.data.username
          User.role=resp.data.role
          User.permissions=resp.data.permissions
          // console.log(resp.data.username)
          localStorage.setItem("userName",resp.data.username)
          localStorage.setItem("role",resp.data.role)
          // User.Password=resp.data.password
          // console.log(resp)
          this.message.info("登录成功！")
          this.router.navigate(['/goods'])
          // console.log("??????????????")
        }else{
          this.message.error("用户信息错误！")
          return
        }

      }
    )
    // this.loginServic.login().subscribe
  }

  showMessage(){
    this.message.warning("请联系管理员")
  }

}
