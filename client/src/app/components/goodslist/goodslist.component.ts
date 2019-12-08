import { Component, OnInit } from '@angular/core';
import { Good } from './good';
import {GoodService} from '../../services/good.service'
import { NzMessageService } from 'ng-zorro-antd';
import { FormGroup, FormControl, FormBuilder, Validators } from '@angular/forms';
import { ExportRe } from '../amount/exportRe';
// import {GoodService} from '../../services/good.service'
// import { NzMessageService } from 'ng-zorro-antd/message';

@Component({
  selector: 'app-goodslist',
  templateUrl: './goodslist.component.html',
  styleUrls: ['./goodslist.component.css']
})
export class GoodslistComponent implements OnInit {

  page:number=1
  pageIndex:number=1
  getSize:number=50
  pageSize:number=10
  orderBy:string="desc"
  updateModel=false
  selectedGood:Good={}
  goods:Array<Good>
  goodsList:Array<Good>
  total:number=0
  filter:string=""
  exportModel=false

  exportForm: FormGroup;
  totalPrice:number=0
  isDisabled:boolean=true
  exportRecord:ExportRe={}

  constructor(private goodService:GoodService,private nzMessageService: NzMessageService,private fb: FormBuilder) { }

  ngOnInit() {
    this.goodService.GoodList(this.page,this.getSize,this.orderBy).subscribe(
      (response)=>{
        if(response.status==-1){
          console.error(`get goodlist error:${response.msg}`)
          return
        }
        if(typeof response.data=='undefined'){
          console.error("goodlist data is undefinded")
          return
        }
        if(response.data===null){
          this.goodsList=[]
          // console.log("asdasdsadad")
          return
        }
        // console.log(response.data)
        this.goodsList=response.data
        this.total=response.rowCount
        // console.log(this.goods)
        this.goods=this.goodsList.slice(0,this.pageSize)
      }
    )
    this.exportForm=this.fb.group({
      eid:[null, [Validators.required]],
      buyer:[null, [Validators.required]],
      bphone:[null,[Validators.required,this.mobileValidator]],
      ecount:[1,[Validators.required]],
      eprice:[0,[Validators.required]],
      etotalPrice:[0,[Validators.required]],
      detial:['']
    })
  }

  formatterDollar = (value: number) => `￥ ${value}`;
  parserDollar = (value: string) => value.replace('￥ ', '');
  formatterCount = (value: number) => `${value}  个`;
  parserCount = (value: string) => value.replace('', '');
  
  showUpdateModel(good:Good){
    this.selectedGood.imdate=good.imdate
    this.selectedGood.gname=good.gname
    this.selectedGood.shipper=good.shipper
    this.selectedGood.count=good.count
    this.selectedGood.sphone=good.sphone
    this.selectedGood.price=good.price
    this.selectedGood.imprice=good.imprice
    this.selectedGood.gno=good.gno
    this.updateModel=true
    // console.log(this.selectedGood)
  }

  handleOk(): void {
    // console.log('Button ok clicked!');
    console.log("修改:",this.selectedGood)
    this.goodService.UpdateGood(this.selectedGood).subscribe(
      (resp)=>{
        if(resp.status==0){
          confirm("修改成功！")
          this.goodService.GoodList(this.page,this.getSize,this.orderBy).subscribe(
            (response)=>{
              if(response.status==-1){
                console.error(`get goodlist error:${response.msg}`)
                return
              }
              if(typeof response.data=='undefined'){
                console.error("goodlist data is undefinded")
                return
              }
              if(response.data===null){
                this.goodsList=[]
                // console.log("asdasdsadad")
                return
              }
              // console.log(response.data)
              this.goodsList=response.data
              this.total=response.rowCount
              // console.log(this.goods)
              this.goods=this.goodsList.slice(0,this.pageSize)
            })
        }else{
          confirm("数据修改有误,修改失败！")
        }
      }
    )
    this.updateModel = false;
  }

  exportOk(){
    console.log("ok")
    if(this.exportForm.valid){
      this.exportRecord.eid=this.exportForm.controls.eid.value
      this.exportRecord.buyer=this.exportForm.controls.buyer.value
      this.exportRecord.bphone=this.exportForm.controls.bphone.value
      this.exportRecord.ecount=this.exportForm.controls.ecount.value
      this.exportRecord.eprice=this.exportForm.controls.eprice.value
      this.exportRecord.detial=this.exportForm.controls.detial.value
      this.exportRecord.etotalprice=this.exportForm.controls.etotalPrice.value
      this.exportRecord.shipper=this.selectedGood.shipper
      this.exportRecord.imdate=this.selectedGood.imdate
      this.exportRecord.gname=this.selectedGood.gname
      console.log(this.exportRecord)
      this.goodService.Export(this.exportRecord).subscribe(
        (resp)=>{
          if(resp.status==0){
            this.nzMessageService.info("出货成功！")
            this.selectedGood={}
            this.goodService.GoodList(this.page,this.getSize,this.orderBy).subscribe(
              (response)=>{
                if(response.status==-1){
                  console.error(`get goodlist error:${response.msg}`)
                  return
                }
                if(typeof response.data=='undefined'){
                  console.error("goodlist data is undefinded")
                  return
                }
                if(response.data===null){
                  this.goodsList=[]
                  return
                }
                this.goodsList=response.data
                this.total=response.rowCount
                this.goods=this.goodsList.slice(0,this.pageSize)
              })
              this.exportModel=false
          }else{
            this.nzMessageService.error("出货失败！请检查单号信息是否重复！")
          }
        }
      )
    }

  }

  handleCancel(): void {
    this.selectedGood={}
    this.updateModel = false;
  }
  exportCancel(){
    this.selectedGood={}
    this.exportModel=false
  }

  //换页
  change(event){
    // console.log(event)
    this.pageIndex=event
    this.goods=this.goodsList.slice((event-1)*this.pageSize,event*this.pageSize)
    return
  }

  //搜索 商品名
  search():void{
    if(this.filter==""||this.filter==undefined||this.filter==null){
      console.log("搜索条件为空")
      return
    }
    this.goodService.Query(this.filter).subscribe(
      (resp)=>{
        if(resp.status==-1){
          console.error(resp.msg)
          return
        }
        if(resp.data==null||resp.data==undefined){
          console.error("search get data null")
          return
        }
        this.goodsList=resp.data
        this.goods=this.goodsList.slice(0,10)
      }
    )
  }

  //确认删除
  confirm(good:Good): void {
    console.log(good)
    this.goodService.DelGood(good.imdate,good.gname,good.shipper).subscribe(
      (resp)=>{
        if(resp.status==0){
          this.nzMessageService.info('删除成功！');
          this.goods=this.goods.filter((g)=>{
            return (g!=good)
          })
        }
        else{
          this.nzMessageService.error('删除失败！');
        }
      }
    )
    
  }

   //手机校验
  mobileValidator(control:FormControl):any {
    let myReg = /^1(3|4|5|7|8)+\d{9}$/;
    let valid = myReg.test(control.value);
    // console.log("moblie的校验结果是"+valid)
    return valid ? null : {mobile:true};//如果valid是true 返回是null
  }

  export(good:Good){
    this.selectedGood.imdate=good.imdate
    this.selectedGood.gname=good.gname
    this.selectedGood.shipper=good.shipper
    this.selectedGood.count=good.count
    this.selectedGood.sphone=good.sphone
    this.selectedGood.price=good.price
    this.selectedGood.imprice=good.imprice
    this.selectedGood.gno=good.gno
    this.exportModel=!this.exportModel
  }

  toggleDisabled(){
    this.isDisabled=!this.isDisabled
  }

  countTotalPrice(){
    this.totalPrice=this.exportForm.controls.ecount.value*this.exportForm.controls.eprice.value
    this.exportForm.controls.etotalPrice.setValue(this.totalPrice)
  }

  getPrice(){
    this.exportForm.controls.eprice.setValue(this.selectedGood.price)
  }

}
