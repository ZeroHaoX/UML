import { Injectable } from '@angular/core';
import { HttpHeaders,HttpClient, HttpErrorResponse} from '@angular/common/http';
import { Observable, of, throwError } from 'rxjs';
import { ReplyProto,ReqProto } from './protocal';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(
    public http:HttpClient
  ) { }
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

  login(username:string,password:string):Observable<ReplyProto>{
    var api="/api/login"
    var user={
      u:username,
      p:password
    }
    this.reqProto.data=user
    return this.http.post(api,this.reqProto,httpOptions).pipe(
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