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
    let path=next.routeConfig.path //要跳转的路由
    
    /*是否登录*/
    if(!this.checkLogin()){
      this.router.navigate(['login'])
      return false
    }
    /*查找是否有权限 */
    if(!this.checkAuth(path)){
      this.router.navigate(['noauth'])
      return false
    }

    return true;
  }


  checkAuth(path:string):boolean{
    for(var p of User.Permissions){
      if(p.API==path){
        return true
      }
    }
    return false
  }

  checkLogin():boolean{
    let token=localStorage.getItem("token")
    if(token==""||token==null||token==undefined){
      return false
    }
    //用户model为空，即刷新了，重新用token获取用户信息
    if(User.Role==""||User.Role==null||User.Role==undefined){
      //请求获取信息
      this.userService.GetUserMes().subscribe(
        (respone)=>{
          User.UserName=respone.data.Username
          User.Role=respone.data.Role
          User.Permissions=respone.data.Permissions
          // User.Password=respone.data.Password
        }
      )
    }
    return true
  }
  
}
