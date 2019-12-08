import { Component, OnInit,Input } from '@angular/core';
import { FormGroup, FormControl, FormBuilder, Validators } from '@angular/forms';
import { ExportRe } from '../amount/exportRe';
import {GoodService} from '../../services/good.service'
import { NzMessageService } from 'ng-zorro-antd/message';
import {Good} from '../goodslist/good'


@Component({
  selector: 'app-export',
  templateUrl: './export.component.html',
  styleUrls: ['./export.component.css']
})
export class ExportComponent implements OnInit {

  exportForm: FormGroup;
  totalPrice:number=0
  isDisabled:boolean=true
  exportRecord:ExportRe={}
  @Input() good:Good={}

  constructor(private fb: FormBuilder) { }

  //手机校验
  mobileValidator(control:FormControl):any {
    let myReg = /^1(3|4|5|7|8)+\d{9}$/;
    let valid = myReg.test(control.value);
    // console.log("moblie的校验结果是"+valid)
    return valid ? null : {mobile:true};//如果valid是true 返回是null
  }

  toggleDisabled(){
    this.isDisabled=!this.isDisabled
  }

  countTotalPrice(){
    this.totalPrice=this.exportForm.controls.ecount.value*this.exportForm.controls.eprice.value
    this.exportForm.controls.etotalPrice.setValue(this.totalPrice)
  }

  submitForm(){
    if(!this.exportForm.valid){
      confirm("请填写正确的进货单信息！")
      return
    }
    this.exportRecord.eid=this.exportForm.controls.importID.value
    this.exportRecord.gname=this.exportForm.controls.gname.value
    this.exportRecord.shipper=this.exportForm.controls.shipper.value
    this.exportRecord.bphone=this.exportForm.controls.phone.value
    this.exportRecord.ecount=this.exportForm.controls.count.value
    this.exportRecord.eprice=this.exportForm.controls.price.value
    this.exportRecord.etotalprice=this.exportForm.controls.totalPrice.value
    this.exportRecord.detial=this.exportForm.controls.detial.value

    // this.exportRecord.Import(this.importRecord).subscribe(
    //   (resp)=>{
    //     if(resp.status==0){
    //       this.message.info('添加成功！\n可前往商品信息管理界面更改售价');
    //     }else{
    //       this.message.error('添加失败！请检查你的信息填写');
    //     }
    //   }
    // )
  }

  getPrice(){
    console.log(this.good.price)
    this.exportForm.controls.eprice.setValue(this.good.price)
  }

  ngOnInit() {
    this.exportForm=this.fb.group({
      eid:[null, [Validators.required]],
      // gname:[null, [Validators.required]],
      buyer:[null, [Validators.required]],
      bphone:[null,[Validators.required,this.mobileValidator]],
      ecount:[1,[Validators.required]],
      eprice:[0,[Validators.required]],
      etotalPrice:[0,[Validators.required]],
      detial:[null]
    })
    // console.log(this.good)
    // this.totalPrice=this.importForm.controls.count.value*this.importForm.controls.price.value
    // this.importForm.controls.totalPrice.setValue(this.totalPrice)
  }

}
