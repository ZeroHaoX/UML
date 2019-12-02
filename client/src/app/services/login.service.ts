import { Injectable } from '@angular/core';
import { HttpHeaders,HttpClient} from '@angular/common/http';
import { Observable } from 'rxjs/internal/Observable';
import { ReplyProto } from './protocal';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(
    public http:HttpClient
  ) { }

  login(username:string,password:string):Observable<ReplyProto>{
    const httpOptions={
      headers:new HttpHeaders({'Content-Type':'application/json'})
    }
    var api="/api/login"
    return this.http.post(api,{},httpOptions)
  }

}
