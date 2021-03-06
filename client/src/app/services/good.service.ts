import { Injectable } from '@angular/core';
import { Observable, of, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import {ReplyProto,ReqProto} from './protocal'
import {HttpHeaders, HttpClient, HttpResponse, HttpErrorResponse} from '@angular/common/http';
import {Good} from '../components/goodslist/good'
import {ImportRecord} from '../components/import/importrecord'
import { ExportRe } from '../components/amount/exportRe';

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
    if(imRecord==null||typeof imRecord=='undefined'||imRecord.id==""){
      console.error("imRecord is null")
      return of({status:-1,msg:"imRecord is nil"})
    }
    let api='/api/import'
    this.reqProto.data=imRecord
    return  this.http.post<ReplyProto>(api,this.reqProto,httpOptions).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  ImportList(page:number,pageSize:number,orderBy:string):Observable<ReplyProto>{
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
    let api="/api/import/list"
    this.reqProto.orderBy=orderBy
    this.reqProto.page=page
    this.reqProto.pageSize=pageSize
    return this.http.post(api,this.reqProto,httpOptions).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  DelImport(eid:string):Observable<ReplyProto>{
    if(eid==""||eid==undefined||eid==null){
      console.error("eid is error:",eid)
      return of({status:-1,msg:"数据传输有误！"})
    }
    let api="/api/export/del?eid="+eid

    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  QueryImport(filter:string):Observable<ReplyProto>{
    if(filter==""||filter==null||filter==undefined){
      console.error("filter is error:",filter)
      return of({status:-1,msg:"filter is nil"})
    }
    let api="/api/import/query?filter="+filter

    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }



    //出货
  Export(exportRecord:ExportRe):Observable<ReplyProto>{
      if(exportRecord==null||typeof exportRecord=='undefined'||exportRecord.eid==""){
        console.error("exportRecord is null")
        return of({status:-1,msg:"exportRecord is nil"})
      }
      let api='/api/export'
      this.reqProto.data=exportRecord
      return  this.http.post<ReplyProto>(api,this.reqProto,httpOptions).pipe(
        catchError(err=>this.handleError(err))
      )
  }

  ExportList(page:number,pageSize:number,orderBy:string):Observable<ReplyProto>{
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
    let api="/api/export/list"
    this.reqProto.orderBy=orderBy
    this.reqProto.page=page
    this.reqProto.pageSize=pageSize
    return this.http.post(api,this.reqProto,httpOptions).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  DelExport(eid:string):Observable<ReplyProto>{
    if(eid==""||eid==undefined||eid==null){
      console.error("eid is error:",eid)
      return of({status:-1,msg:"数据传输有误！"})
    }
    let api="/api/export/del?eid="+eid

    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }

  QueryExport(filter:string):Observable<ReplyProto>{
    if(filter==""||filter==null||filter==undefined){
      console.error("filter is error:",filter)
      return of({status:-1,msg:"filter is nil"})
    }
    let api="/api/export/query?filter="+filter

    return this.http.get<ReplyProto>(api,{withCredentials:true}).pipe(
      catchError(err=>this.handleError(err))
    )
  }


  // 月结出货和出货记录
  ExportAndImportRecord(year:string,month:string):Observable<ReplyProto>{
    if(year==null||typeof year=='undefined'||year==""){
      console.error("year is null")
      return of({status:-1,msg:"year is nil"})
    }
    if(month==null||typeof month=='undefined'||month==""){
      console.error("month is null")
      return of({status:-1,msg:"month is nil"})

    }
    let api='/api/monthly?year='+year+"&"+"month="+month
    return  this.http.get<ReplyProto>(api).pipe(
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