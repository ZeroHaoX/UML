import { Component, OnInit } from '@angular/core';
import {User} from './user'
import {UserService} from '../../services/user.service'
import { NzMessageService } from 'ng-zorro-antd';
// import { type } from 'os';

@Component({
  selector: 'app-userlist',
  templateUrl: './userlist.component.html',
  styleUrls: ['./userlist.component.css']
})
export class UserlistComponent implements OnInit {

  updateModel=false
  selectedUser:User={}
  page:number=1
  pageIndex:number=1
  getSize:number=50
  pageSize:number=10
  orderBy:string="asc"
  users:Array<User>
  usersList:Array<User>
  total:number=0
  filter:string=""
  userTemp:User={}

  constructor(public userService:UserService,private nzMessageService: NzMessageService) { }

  ngOnInit() {
    this.userService.GetUserList(this.page,this.getSize,this.orderBy).subscribe(
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
          this.usersList=[]
          // console.log("asdasdsadad")
          return
        }
        // console.log(response.data)
        this.usersList=response.data
        this.total=response.rowCount
        this.users=this.usersList.slice(0,this.pageSize)
        console.log(this.usersList)
      }
    )
  }

  search(){
    this.userService.Query(this.filter)
  }

  remove(user:User){
    if(user==null || typeof user =='undefined'){
      console.error("user 为空!")
      return
    }
    this.userService.DelUser(user.username).subscribe(
      (resp)=>{
        if(resp.status==0){
          this.nzMessageService.error('删除成功！');
          this.users=this.users.filter(u=>{
            return user!=u
          })
        }else{
          this.nzMessageService.error('删除失败！');
        }
      }
    )
  }

  showUpdateModel(user:User){
    this.selectedUser.username=user.username
    this.selectedUser.actualname=user.actualname
    this.selectedUser.role=user.role
    this.selectedUser.phone=user.phone
    this.selectedUser.password=user.password
    this.updateModel=true
    // this.userTemp=user
    // console.log(this.selectedUser)
  }

  handleOk(): void {
    // console.log('Button ok clicked!');
    this.userService.UpdateUser(this.selectedUser).subscribe(
      (resp)=>{
        if(resp.status==0){
          confirm("修改成功！")
          // this.userTemp=this.selectedUser
          this.selectedUser={}
          this.userService.GetUserList(1,this.getSize,this.orderBy).subscribe(
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
                this.usersList=[]
                // console.log("asdasdsadad")
                return
              }
              // console.log(response.data)
              this.usersList=response.data
              this.total=response.rowCount
              this.users=this.usersList.slice(0,this.pageSize)
              // console.log(this.usersList)
            }
          )
        }else{
          confirm("修改失败！")
        }
      }
    )
    this.updateModel = false
  }

  handleCancel(): void {
    this.updateModel = false
    this.selectedUser={}
  }

}
