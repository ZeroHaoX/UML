import { Component, OnInit } from '@angular/core';
import { ExportRe, ImportRe } from './exportRe';
import { GoodService } from 'src/app/services/good.service';


@Component({
  selector: 'app-amount',
  templateUrl: './amount.component.html',
  styleUrls: ['./amount.component.css']
})
export class AmountComponent implements OnInit {

  constructor(
    private goodService:GoodService
  ) { }

  exportTitle = "本月出货记录"
  importTitle = "本月进货记录"
  page:number=1
  exportResPageIndex:number=1
  importResPageIndex:number=1
  getSize:number=50
  pageSize:number=10
  orderBy:string="desc"
  updateModel=false
  exportResTotal:number=0
  importResTotal:number=0
  filter:string=""
  // 出货记录
  exportRes:Array<ExportRe>=[]
  displayExportRes:Array<ExportRe>=[]
  // 出货日期
  exportTime:string[]
  // 单日时间
  dateTime: Date;
  // 进货记录
  importRes:Array<ImportRe>=[]
  displayImportRes:Array<ImportRe>=[]
  // 利润
  profit:number
  // 本月营业额
  turnover:number
  // 显示数量
  displayNum = 10

  ngOnInit() {
    this.dateTime = new Date()
    this.exportTime = this.dateTime.toLocaleDateString().split("/",-1)
    console.log(this.exportTime)
    var year:string = this.exportTime[0]
    var month:string = this.exportTime[1]
    this.goodService.ExportAndImportRecord(year,month).subscribe((response:any)=>{
      if(response.status==-1){
        console.error(`get goodlist error:${response.msg}`)
        return
      }
      if(typeof response.data=='undefined'){
        console.error("goodlist data is undefinded")
        return
      }
      if(response.data===null){
        this.exportRes=[]
        this.importRes=[]
        return
      }
      this.profit = response.data.account.profit
      this.turnover = response.data.account.turnover
      this.exportRes = response.data.erecord
      this.exportResTotal = this.exportRes.length
      this.importRes = response.data.imrecord
      this.importResTotal = this.importRes.length
      this.displayImportRes=this.importRes.slice(0,this.pageSize)
      this.displayExportRes=this.exportRes.slice(0,this.pageSize)
    })
  }


  SearchExportRe(event){
    // 获取当前时间
    this.exportTime = event.toLocaleDateString().split("/",-1)
    var year:string = this.exportTime[0]
    var month:string = this.exportTime[1]
    console.log(year,month)
    // 时间改变的时候，就会触发这个函数
    this.goodService.ExportAndImportRecord(year,month).subscribe((response:any)=>{
      
      if(response.status==-1){
        console.error(`get goodlist error:${response.msg}`)
        return
      }
      if(typeof response.data=='undefined'){
        console.error("goodlist data is undefinded")
        return
      }
      console.log(response)
      if(response.rowCount==0){
        this.exportRes = []
        this.importRes =[]
        this.displayImportRes = []
        this.displayExportRes =[]
        return
      }
      if(response.data===null){
        this.exportRes=[]
        this.importRes=[]
        // console.log("asdasdsadad")
        return
      }
      this.profit = response.data.account.profit
      this.turnover = response.data.account.turnover
      this.exportRes = response.data.erecord
      this.exportResTotal = this.exportRes.length
      this.importRes = response.data.imrecord
      this.importResTotal = this.exportRes.length
      this.displayImportRes=this.importRes.slice(0,this.pageSize)
      this.displayExportRes=this.exportRes.slice(0,this.pageSize)
    })
  }

  //换页
  exportResChange(event){
    this.exportResPageIndex=event
    this.displayExportRes=this.exportRes.slice((event-1)*this.pageSize,event*this.pageSize)
  }

  importResChange(event){
    this.importResPageIndex=event
    this.displayImportRes=this.importRes.slice((event-1)*this.pageSize,event*this.pageSize)
  }

}
