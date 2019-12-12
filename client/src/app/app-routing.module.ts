import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { GoodslistComponent } from './components/goodslist/goodslist.component';
import { UserlistComponent } from './components/userlist/userlist.component';
import { LoginComponent } from './components/login/login.component';
import {AuthGuard} from './guards/auth.guard'
import {RegistComponent} from './components/regist/regist.component'
import { ImportComponent } from './components/import/import.component';
import { AmountComponent } from './components/amount/amount.component';
import { ExportComponent } from './components/export/export.component';
import { ImportlistComponent } from './components/importlist/importlist.component';

const routes: Routes = [
  {path:"login",component:LoginComponent},
  // {path:'users',canActivate:[AuthGuard],component:UserlistComponent},
  // {path:'goods',canActivate:[AuthGuard],component:GoodslistComponent},
  {path:'users',canActivate:[AuthGuard],component:UserlistComponent},
  {path:'goods',canActivate:[AuthGuard],component:GoodslistComponent},
  {path:'registe',component:RegistComponent},
  {path:'import',canActivate:[AuthGuard],component:ImportlistComponent},
  {path:'amount',canActivate:[AuthGuard],component:AmountComponent},
  {path:'export',canActivate:[AuthGuard],component:ExportComponent},
  {path:'**',redirectTo:'login'}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
