import { Component, OnInit } from '@angular/core';
import User from '../../guards/usermodel'
import { Router } from '@angular/router';
import { THIS_EXPR } from '@angular/compiler/src/output/output_ast';

@Component({
  selector: 'app-head',
  templateUrl: './head.component.html',
  styleUrls: ['./head.component.css']
})
export class HeadComponent implements OnInit {

  constructor(private router:Router) { }

  userModel:any
  
  ngOnInit() {
    this.userModel=User
  }

  logout(){
    localStorage.clear()
    this.router.navigate(['/login'])
  }


}
