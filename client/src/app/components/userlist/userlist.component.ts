import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-userlist',
  templateUrl: './userlist.component.html',
  styleUrls: ['./userlist.component.css']
})
export class UserlistComponent implements OnInit {

  updateModel=false
  selectedUser={
    UserName:"",
    ActualName:"",
    Phone:"",
    Role:""
  }
  constructor() { }

  ngOnInit() {
  }
  
  pageSize=10

  users=[
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
    {UserName:"小明",ActualName:"张少",Phone:"13756845751",Role:"员工"},
  ]

  remove(i:any){
    this.users=this.users.filter(u=>{
      return u.ActualName!=i.ActualName
    })
  }

  showUpdateModel(user){
    this.selectedUser=user
    this.updateModel=true
    console.log(this.selectedUser)
  }

  handleOk(): void {
    // console.log('Button ok clicked!');
    this.updateModel = false;
  }

  handleCancel(): void {
    // console.log('Button cancel clicked!');
    this.updateModel = false;
    this.selectedUser=null;
  }

}
