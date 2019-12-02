import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { GoodslistComponent } from './components/goodslist/goodslist.component';
import { UserlistComponent } from './components/userlist/userlist.component';
import { LoginComponent } from './components/login/login.component';

const routes: Routes = [
  {path:"login",component:LoginComponent},
  {path:'userlist',component:UserlistComponent},
  {path:'goodslist',component:GoodslistComponent},
  {path:'**',redirectTo:'userlist'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
