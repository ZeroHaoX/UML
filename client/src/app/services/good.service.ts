import { Injectable } from '@angular/core';
import { Observable, of, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import {ReplyProto,ReqProto} from './protocal'
import {HttpHeaders, HttpClient, HttpResponse, HttpErrorResponse} from '@angular/common/http';
import {Good} from '../components/goodslist/good'
import {ImportRecord} from '../components/import/importrecord'

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
      return of({status:-1,msg:"数据传输有误！"})
    }
    if(orderBy!="desc" && orderBy!="asc"){
      console.error("orderBy is error:",orderBy)
      return of({status:-1,msg:"数据传输有误！"})
    }
    // console.log("传输")
    let api="/api/goods"
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

    let api="/api/goods/update"
    return this.http.post<ReplyProto>(api,this.reqProto,httpOptions).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  DelGood(imdate:string,gname:string,shipper:string):Observable<ReplyProto>{
    if(gname==""||gname==undefined||gname==null){
      console.error("gname is error:",gname)
      return of({status:-1,msg:"数据传输有误！"})
    }
    let api="/api/goods/del?gname="+gname+"&shipper="+shipper+"&imdate="+imdate

    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  Query(filter:string):Observable<ReplyProto>{
    if(filter==""||filter==null||filter==undefined){
      console.error("filter is error:",filter)
      return of({status:-1,msg:"filter is nil"})
    }
    let api="/api/goods/query?filter="+filter

    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  //进货
  Import(imRecord:ImportRecord):Observable<ReplyProto>{
    if(imRecord==null||typeof imRecord=='undefined'){
      console.error("imRecord is null")
      return of({status:-1,msg:"imRecord is nil"})
    }
    let api='/api/import'
    this.reqProto.data=imRecord
    return  this.http.post<ReplyProto>(api,this.reqProto,httpOptions).pipe(
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