import { Component, OnInit,Input } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators } from '@angular/forms';
import { ExportRe } from '../amount/exportRe';
import {GoodService} from '../../services/good.service'
import { NzMessageService } from 'ng-zorro-antd/message';
import {Good} from '../goodslist/good'
import { Logs } from 'selenium-webdriver';


@Component({
  selector: 'app-export',
  templateUrl: './export.component.html',
  styleUrls: ['./export.component.css']
})
export class ExportComponent implements OnInit {

  page:number=1
  exportResPageIndex:number=1
  // 出货记录
  exportRes:Array<ExportRe>=[]
  displayExportRes:Array<ExportRe>=[]
  // 显示数量
  displayNum = 10
  getSize:number=50
  pageSize:number=10
  orderBy:string="desc"
  total:number=0
  filter:string=""
  // exportResPageIndex:number=1

  constructor(private goodService: GoodService,private nzMessageService: NzMessageService) { }

    //换页
    change(event){
      // console.log(event)
      this.exportResPageIndex=event
      this.displayExportRes=this.exportRes.slice((event-1)*this.pageSize,event*this.pageSize)
      return
    }

      //确认删除
    confirm(record:ExportRe): void {
      console.log(record)
      this.goodService.DelExport(record.eid).subscribe(
        (resp)=>{
          if(resp.status==0){
            this.nzMessageService.info('删除成功！');
            this.exportRes=this.exportRes.filter((g)=>{
              return (g!=record)
            })
            this.displayExportRes=this.exportRes.slice(( this.exportResPageIndex-1)*this.pageSize, this.exportResPageIndex*this.pageSize)
          }
          else{
            this.nzMessageService.error('删除失败！');
          }
        }
      )
    }

    search(){
      if(this.filter==""){
        console.log("搜索条件为空")
        return
      }
      this.goodService.QueryExport(this.filter).subscribe(
        (resp)=>{
          if(resp.status==0){
            this.exportRes=resp.data
            this.total=resp.rowCount
            this.displayExportRes=this.exportRes.slice(0,10)
            this.exportResPageIndex=1
            this.nzMessageService.info('搜索成功')
          }else{
            this.nzMessageService.error('请检查网络状况')
          }
        }
      )
    }


  ngOnInit() {
    this.goodService.ExportList(this.page,this.getSize,this.orderBy).subscribe(
      (resp)=>{
        if(resp.status==0){
          // console.log(response.data)
          this.exportRes=resp.data
          this.total=resp.rowCount
          // console.log(this.goods)
          this.displayExportRes=this.exportRes.slice(0,this.pageSize)
        }else{
          
        }
      }
    )
  }

}
