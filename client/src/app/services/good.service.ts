import { Injectable } from '@angular/core';
import { Observable, of, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import {ReplyProto,ReqProto} from './protocal'
import {HttpHeaders, HttpClient, HttpResponse, HttpErrorResponse} from '@angular/common/http';
import {Good} from '../components/goodslist/good'

@Injectable({

  providedIn: 'root'
})
export class GoodService {

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

  GoodList(page:number,pageSize:number,orderBy:string):Observable<ReplyProto>{
    if(page<0||page==null||page==undefined){
      console.error("page is error:",page)
      return of({status:-1,msg:"数据传输有误！"})
    }
    if(pageSize<0||pageSize==null||pageSize==undefined){
      console.error("pageSize is error:",pageSize)
      return of({status:500,msg:"数据传输有误！"})
    }
    if(orderBy==""||orderBy==null||orderBy==undefined){
      console.error("orderBy is error:",orderBy)
      return of({status:500,msg:"数据传输有误！"})
    }
    let api="http://localhost:8080/goodlist"
    this.reqProto.orderBy=orderBy
    this.reqProto.page=page
    this.reqProto.pageSize=pageSize
    return this.http.post(api,this.reqProto,httpOptions).pipe(
      catchError(err=>this.handleError(err))
    )
    
  }

  UpdateGood(good:Good):Observable<ReplyProto>{
    if(good==undefined||good==null){
      console.error("传输的数据有错:",good)
      return of({status:-1,msg:"数据传输有误！"})
    }
    this.reqProto.data=good

    let api="http://localhost:8080/updategood"
    return this.http.post<ReplyProto>(api,this.reqProto,httpOptions).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  DelGood(gno:string):Observable<ReplyProto>{
    if(gno==""||gno==undefined||gno==null){
      console.error("gno is error:",gno)
      return of({status:-1,msg:"数据传输有误！"})
    }
    let api="http://localhost:8080/delgood?gno="+gno

    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  Query(filter:string):Observable<ReplyProto>{
    if(filter==""||filter==null||filter==undefined){
      console.error("filter is error:",filter)
      return of({status:-1,msg:"filter is nil"})
    }
    let api="http://localhost:8080/searchgood?filter="+filter

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
  'Authorization': localStorage.getItem('token')
})}