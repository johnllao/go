import { NgModule }      from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { 
  AppComponent, 
  BudgetComponent, 
  ProfileComponent 
}                        from './app.component';

@NgModule({
  imports:      [ BrowserModule ],
  declarations: [ AppComponent, BudgetComponent, ProfileComponent ],
  bootstrap:    [ AppComponent ]
})
export class AppModule { }
