import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot, UrlTree, Router } from '@angular/router';
import { Observable, range } from 'rxjs';
import User from './usermodel';
import {UserService} from '../services/user.service'

@Injectable({
  providedIn: 'root'
})
export class AuthGuard implements CanActivate {
  constructor(private router:Router,public userService:UserService){}

  canActivate(
    next: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    let path='/'+next.routeConfig.path //要跳转的路由
    // console.log(path)
    /*是否登录*/
    // if(!this.checkLogin()){
    //   this.router.navigate(['login'])
    //   console.log("跳转登录")
    //   return false
    // }
    // /*查找是否有权限 */
    // if(!this.checkAuth(path)){
    //   this.router.navigate(['goods'])
    //   return false
    // }

    return true;
  }


  checkAuth(path:string):boolean{
    for(var p of User.permissions){
      // console.log(p.api)
      if(p.api==path){
        return true
      }
    }
    return false
  }

  checkLogin():boolean{
    // let token=localStorage.getItem("token")
    // if(token==""||token==null||token==undefined){
    //   return false
    // }
    //用户model为空，即刷新了，重新用token获取用户信息
    if(User.role==""||User.role==null||User.role==undefined){
      //请求获取信息
      this.userService.GetUserMes().subscribe(
        (respone)=>{
          User.userName=respone.data.username
          User.role=respone.data.role
          User.permissions=respone.data.permissions
          // User.Password=respone.data.Password
        }
      )
    }
    return true
  }
  
}
