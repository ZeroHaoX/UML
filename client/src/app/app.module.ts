import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgZorroAntdModule, NZ_I18N, zh_CN } from 'ng-zorro-antd';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { registerLocaleData } from '@angular/common';
import zh from '@angular/common/locales/zh';
import { HeadComponent } from './components/head/head.component';
import { GoodslistComponent } from './components/goodslist/goodslist.component';
import { UserlistComponent } from './components/userlist/userlist.component';
import { LoginComponent } from './components/login/login.component';
import { RegistComponent } from './components/regist/regist.component';
import { ImportComponent } from './components/import/import.component';
import { AmountComponent } from './components/amount/amount.component';
import { ExportComponent } from './components/export/export.component';

registerLocaleData(zh);

@NgModule({
  declarations: [
    AppComponent,
    HeadComponent,
    GoodslistComponent,
    UserlistComponent,
    LoginComponent,
    RegistComponent,
    ImportComponent,
    AmountComponent,
    ExportComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    NgZorroAntdModule,
    FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,
  ],
  providers: [{ provide: NZ_I18N, useValue: zh_CN }],
  bootstrap: [AppComponent]
})
export class AppModule { }
