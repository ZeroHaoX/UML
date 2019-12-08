import { Component, OnInit } from '@angular/core';
import { ExportRe } from './exportRe';

@Component({
  selector: 'app-amount',
  templateUrl: './amount.component.html',
  styleUrls: ['./amount.component.css']
})
export class AmountComponent implements OnInit {

  constructor() { }

  exportRes:Array<ExportRe>=[]


  ngOnInit() {
  }

}
