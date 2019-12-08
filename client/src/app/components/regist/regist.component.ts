import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators } from '@angular/forms';
import { User } from '../userlist/user'
import {UserService} from '../../services/user.service'
import { Router } from '@angular/router';

@Component({
  selector: 'app-regist',
  templateUrl: './regist.component.html',
  styleUrls: ['./regist.component.css']
})
export class RegistComponent implements OnInit {
  validateForm: FormGroup;

  submitForm(): void {
    if(this.validateForm.valid){
      this.user.username=this.validateForm.controls.userName.value
      this.user.password=this.validateForm.controls.password.value
      this.user.actualname=this.validateForm.controls.actualName.value
      this.user.phone=this.validateForm.controls.phoneNumber.value
      // console.log(this.user)
      this.userService.Registe(this.user).subscribe(
        (resp)=>{
          if(resp.status==0){
            confirm("注册成功！")
            this.router.navigate(['/login'])
            return
          }else{
            confirm("用户已存在!")
          }
        }
      )
      
    }else{
      for (const i in this.validateForm.controls) {
        this.validateForm.controls[i].markAsDirty();
        this.validateForm.controls[i].updateValueAndValidity();
     }
    }

    
  }

  updateConfirmValidator(): void {
    /** wait for refresh value */
    Promise.resolve().then(() => this.validateForm.controls.checkPassword.updateValueAndValidity());
  }

  confirmationValidator = (control: FormControl): { [s: string]: boolean } => {
    if (!control.value) {
      return { required: true };
    } else if (control.value !== this.validateForm.controls.password.value) {
      return { confirm: true, error: true };
    }
    return {};
  };
 

  user:User={}

  constructor(private fb: FormBuilder,private userService:UserService,private router:Router) {}

  ngOnInit(): void {
    this.validateForm = this.fb.group({
      userName: [null, [Validators.required]],
      password: [null, [Validators.required]],
      checkPassword: [null, [Validators.required, this.confirmationValidator]],
      actualName: [null, [Validators.required]],
      phoneNumberPrefix: ['+86'],
      phoneNumber: [null, [Validators.required]],
      // agree: [false]
    });
  }
}
