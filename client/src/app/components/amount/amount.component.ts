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

  // 出货记录
  exportRes:Array<ExportRe>=[]
  // 出货日期
  exportTime:string[]
  // 单日时间
  dateTime: Date;
  // 进货记录
  importRes:Array<ImportRe>=[]

  ngOnInit() {
    this.dateTime = new Date()
    this.exportTime = this.dateTime.toLocaleDateString().split("/",-1)
    var year:string = this.exportTime[0]
    var month:string = this.exportTime[1]
    this.goodService.ExportAndImportRecord(year,month).subscribe((Response:any)=>{
      console.log(Response)
    })
  }

  SearchExportRe(event){
    // 获取当前时间
    this.exportTime = event.toLocaleDateString().split("/",-1)
    var year:string = this.exportTime[0]
    var month:string = this.exportTime[1]
    console.log(year,month)
    // 时间改变的时候，就会触发这个函数
    this.goodService.ExportAndImportRecord(year,month).subscribe((Response:any)=>{
      console.log(Response)
    })
  }

  
  search(){

  }

}
