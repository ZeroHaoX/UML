import { Component, OnInit } from '@angular/core';
import { Good } from './good';
import {GoodService} from '../../services/good.service'

@Component({
  selector: 'app-goodslist',
  templateUrl: './goodslist.component.html',
  styleUrls: ['./goodslist.component.css']
})
export class GoodslistComponent implements OnInit {

  page:number=1
  pageSize:number=10
  orderBy:string="esc"
  updateModel=false
  selectedGood={
    Gno:"",
    Gname:"",
    Shipper:"",
    Phone:"",
    Count:0,
    Price:0,
    Imprice:0
  }
  goods:Array<Good>


  constructor(private goodService:GoodService) { }

  ngOnInit() {
    this.goodService.GoodList(this.page,this.pageSize,this.orderBy).subscribe(
      (response)=>{
        if(response.status==-1){
          console.error(`get goodlist error:${response.msg}`)
          return
        }
        if(response.data==null||response.data==undefined){
          console.error("goodlist data is null")
        }
        this.goods=response.data
      }
    )
  }

  // goods:Good[]=[
  //   {Gno:"WTZ-15681",Gname:"商品1",Shipper:"进货商1",Phone:"13576548452",Count:50,Price:15000,Imprice:1000},
  //   {Gno:"WTZ-15682",Gname:"商品1",Shipper:"进货商1",Phone:"13576548452",Count:50,Price:15000,Imprice:1000},
  //   {Gno:"WTZ-15683",Gname:"商品1",Shipper:"进货商1",Phone:"13576548452",Count:50,Price:15000,Imprice:1000},
  //   {Gno:"WTZ-15684",Gname:"商品1",Shipper:"进货商1",Phone:"13576548452",Count:50,Price:15000,Imprice:1000},
  //   {Gno:"WTZ-15685",Gname:"商品1",Shipper:"进货商1",Phone:"13576548452",Count:50,Price:15000,Imprice:1000},
  //   {Gno:"WTZ-15686",Gname:"商品1",Shipper:"进货商1",Phone:"13576548452",Count:50,Price:15000,Imprice:1000},
  //   {Gno:"WTZ-15687",Gname:"商品1",Shipper:"进货商1",Phone:"13576548452",Count:50,Price:15000,Imprice:1000},
  //   {Gno:"WTZ-15688",Gname:"商品1",Shipper:"进货商1",Phone:"13576548452",Count:50,Price:15000,Imprice:1000},
  //   {Gno:"WTZ-15689",Gname:"商品1",Shipper:"进货商1",Phone:"13576548452",Count:50,Price:15000,Imprice:1000},
  // ]
  

  remove(good:Good){
    // console.log(good)
    this.goods=this.goods.filter((g)=>{
      return g.Gno!=good.Gno
    })
  }

  showUpdateModel(good){
    this.selectedGood=good
    this.updateModel=true
    console.log(this.selectedGood)
  }

  handleOk(): void {
    // console.log('Button ok clicked!');
    this.updateModel = false;
  }

  handleCancel(): void {
    // console.log('Button cancel clicked!');
    this.updateModel = false;
    
  }


  search(filter:string):void{
    if(filter==""||filter==undefined||filter==null){
      return
    }
    this.goodService.Query(filter).subscribe(
      (resp)=>{
        if(resp.status==-1){
          console.error(resp.msg)
          return
        }
        if(resp.data==null||resp.data==undefined){
          console.error("search get data null")
          return
        }
        this.goods=resp.data
      }
    )
  }

}
