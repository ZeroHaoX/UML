import { Component, OnInit,EventEmitter, Input, Output } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators } from '@angular/forms';
import { ImportRecord } from './importrecord';
import {GoodService} from '../../services/good.service'
import { NzMessageService } from 'ng-zorro-antd/message';

@Component({
  selector: 'app-import',
  templateUrl: './import.component.html',
  styleUrls: ['./import.component.css']
})
export class ImportComponent implements OnInit {

  @Output() refresh=new EventEmitter()

  importForm: FormGroup;
  totalPrice:number=0
  isDisabled:boolean=true
  importRecord:ImportRecord={}

  constructor(private fb: FormBuilder,private goodService:GoodService,private message: NzMessageService) { }

  //手机校验
  mobileValidator(control:FormControl):any {
    let myReg = /^1(3|4|5|7|8)+\d{9}$/;
    let valid = myReg.test(control.value);
    // console.log("moblie的校验结果是"+valid)
    return valid ? null : {mobile:true};//如果valid是true 返回是null
  }

  doRefresh(){
    this.refresh.emit()
  }

  toggleDisabled(){
    this.isDisabled=!this.isDisabled
  }

  countTotalPrice(){
    this.totalPrice=this.importForm.controls.count.value*this.importForm.controls.price.value
    this.importForm.controls.totalPrice.setValue(this.totalPrice)
  }

  submitForm(){
    if(!this.importForm.valid){
      confirm("请填写正确的进货单信息！")
      return
    }
    this.importRecord.id=this.importForm.controls.importID.value
    this.importRecord.gname=this.importForm.controls.gname.value
    this.importRecord.shipper=this.importForm.controls.shipper.value
    this.importRecord.sphone=this.importForm.controls.phone.value
    this.importRecord.imcount=this.importForm.controls.count.value
    this.importRecord.imprice=this.importForm.controls.price.value
    this.importRecord.imtotalprice=this.importForm.controls.totalPrice.value
    this.importRecord.detial=this.importForm.controls.detial.value
    this.goodService.Import(this.importRecord).subscribe(
      (resp)=>{
        if(resp.status==0){
          this.message.info('添加成功！\n可前往商品信息管理界面更改售价');
          this.doRefresh()
          this.importForm.reset()
        }else{
          this.message.error('添加失败！请检查你的信息填写');
        }
      }
    )
  }

  ngOnInit() {
    this.importForm=this.fb.group({
      importID:[null, [Validators.required]],
      gname:[null, [Validators.required]],
      shipper:[null, [Validators.required]],
      phone:[null,[Validators.required,this.mobileValidator]],
      count:[0,[Validators.required]],
      price:[0,[Validators.required]],
      totalPrice:[0,[Validators.required]],
      detial:['']
    })
    // this.totalPrice=this.importForm.controls.count.value*this.importForm.controls.price.value
    // this.importForm.controls.totalPrice.setValue(this.totalPrice)
  }
}
