import { Component, OnInit } from '@angular/core';
import {ImportRe} from '../amount/exportRe'
import {GoodService} from '../../services/good.service'
import { NzMessageService } from 'ng-zorro-antd';

@Component({
  selector: 'app-importlist',
  templateUrl: './importlist.component.html',
  styleUrls: ['./importlist.component.css']
})
export class ImportlistComponent implements OnInit {


  page:number=1
  importResPageIndex:number=1
  // 进货记录
  importRes:Array<ImportRe>=[]
  displayImportRes:Array<ImportRe>=[]
  // 显示数量
  displayNum = 10
  getSize:number=50
  pageSize:number=10
  orderBy:string="desc"
  total:number=0
  filter:string=""

  search(){
    if(this.filter==""){
      console.log("搜索条件为空")
      return
    }
    this.goodService.QueryImport(this.filter).subscribe(
      (resp)=>{
        if(resp.status==0){
          this.importRes=resp.data
          this.total=resp.rowCount
          this.displayImportRes=this.importRes.slice(0,10)
          this.importResPageIndex=1
          this.nzMessageService.info('搜索成功')
        }else{
          this.nzMessageService.error('请检查网络状况')
        }
      }
    )
  }

  //换页
  change(event){
    // console.log(event)
    this.importResPageIndex=event
      this.displayImportRes=this.importRes.slice((event-1)*this.pageSize,event*this.pageSize)
      return
  }

  //确认删除
  confirm(record:ImportRe): void {
    console.log(record)
    this.goodService.DelImport(record.id).subscribe(
        (resp)=>{
          if(resp.status==0){
            this.nzMessageService.info('删除成功！');
            this.importRes=this.importRes.filter((g)=>{
              return (g!=record)
            })
            this.displayImportRes=this.importRes.slice(( this.importResPageIndex-1)*this.pageSize, this.importResPageIndex*this.pageSize)
          }else{
            this.nzMessageService.error('删除失败！');
          }
        }
    )
  }

  constructor(private goodService:GoodService,private nzMessageService: NzMessageService) { }

  ngOnInit() {
    this.goodService.ImportList(this.page,this.getSize,this.orderBy).subscribe(
      (resp)=>{
        if(resp.status==0){
          // console.log(response.data)
          this.importRes=resp.data
          this.total=resp.rowCount
          // console.log(this.goods)
          this.displayImportRes=this.importRes.slice(0,this.pageSize)
        }else{
          
        }
      }
    )
  }

}
