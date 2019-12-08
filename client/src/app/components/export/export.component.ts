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

  constructor(private goodService: GoodService) { }


  ngOnInit() {
    // this.goodService.ExportList().subscribe(

    // )
  }

}
