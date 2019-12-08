import { Injectable } from '@angular/core';
import {HttpHeaders, HttpClient, HttpResponse, HttpErrorResponse} from '@angular/common/http';
import {User} from '../components/userlist/user'
import { Observable, of, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import {ReplyProto,ReqProto} from './protocal'

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(public http:HttpClient) { }

  //请求数据实例
  private reqProto: ReqProto = {
    action: "",
    data:null,
    sets: [],
    orderBy: "",
    filter: "",
    page: 0,
    pageSize: 0
  }

  //注册
  Registe(user:User):Observable<ReplyProto>{
    if(user==null|| typeof user=='undefined'){
      console.error("user is error")
      return of({status:-1,msg:"注册的user为空！"})
    }
    let api="/api/registe"
    this.reqProto.data=user
    return this.http.post(api,this.reqProto,httpOptions).pipe(
      catchError(err=>this.handleError(err))
    )
    
  }

  //请求用户信息
  GetUserMes():Observable<ReplyProto>{
    let api="http://localhost/usermes"
    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }
  
  //用户列表展示
  GetUserList(page:number,pageSize:number,orderBy:string):Observable<ReplyProto>{
    console.log("传输")
    if(page<0||page==undefined||page==null){
      console.error("page is error:",page)
      return of({status:-1,msg:"数据page有误！"})
    }
    if(pageSize<=0||pageSize==undefined||pageSize==null||(typeof pageSize)!="number"){
      console.error("pageSize is error:",pageSize)
      return of({status:-1,msg:"数据pageSize有误！"})
    }
    if(orderBy!='desc'&&orderBy!='asc'){
      console.error("orderBy is error")
      return of({status:-1,msg:"数据orderBy有误！"})
    }
    let api='/api/users'
    this.reqProto.page=page
    this.reqProto.pageSize=pageSize
    this.reqProto.orderBy=orderBy
    return this.http.post<ReplyProto>(api,this.reqProto,httpOptions).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  //更新用户信息
  UpdateUser(user:User):Observable<ReplyProto>{
    if(user==null||user==undefined){
      console.error("传输的数据有错:",user)
      return of({status:-1,msg:"数据传输有误！"})
    }
    let api='/api/users/update'
    this.reqProto.data=user
    return this.http.post<ReplyProto>(api,this.reqProto,httpOptions)
    .pipe(
      catchError(err=>this.handleError(err))
    )
  }

  //删除用户
  DelUser(userName:string):Observable<ReplyProto>{
    if(userName==null||userName==undefined||userName==""){
      console.error("传输的数据有错:",userName)
      return of({status:-1,msg:"数据传输有误！"})
    }
    let api='/api/users/del?username='+userName
    // this.reqProto.data=userName
    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }
    
  Query(filter:string):Observable<ReplyProto>{
    if(filter==""||filter==null||filter==undefined){
      console.error("filter is error:",filter)
      return of({status:-1,msg:"filter is nil"})
    }
    let api="/api/users/query?filter="+filter

    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  private handleError(error: HttpErrorResponse) {
    if (error.error instanceof ErrorEvent) {
      console.error('An error occurred:', error.error.message);
    } else {
      console.error(
        `Backend returned code ${error.status}, ` +
        `body was: ${error.error}`);
    }
    return throwError(
      `Something bad happened ${error.error}; please try again later.`);
  };



}

const httpOptions={headers:new HttpHeaders({
  'Content-Type':  'application/json',
  // 'Authorization': localStorage.getItem('token')
})}