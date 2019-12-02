import { Injectable } from '@angular/core';
import { HttpHeaders, HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs/internal/Observable';
import { ReplyProto, ReqProto } from './protocal';
import { user } from './user';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(
    public http: HttpClient
  ) { }

  //请求数据实例
  reqProto: ReqProto = {
    action: "",
    data: null,
    sets: [],
    orderBy: "",
    filter: "",
    page: 0,
    pageSize: 0
  }
  userInfo:user={
    userName:"",
    password:""
  }

  login(userName: string, password: string): Observable<ReplyProto> {
    const httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' })
    }
    var api = "/api/login"
    this.userInfo.userName = userName
    this.userInfo.password = password
    this.reqProto.data = this.userInfo
    return this.http.post(api, this.reqProto, httpOptions)
  }

}
