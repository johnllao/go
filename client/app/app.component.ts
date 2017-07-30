import { Component } from '@angular/core';

@Component({
  selector: 'profile-component',
  template: `<h4>Profile</h4>`
})
export class ProfileComponent {}

@Component({
  selector: 'budget-component',
  template: `<h4>Budget</h4>`
})
export class BudgetComponent {}

@Component({
  selector: 'main-component',
  template: `
  <h3 class="banner">{{banner}}</h3>
  <div>
    <a href="#" (click)="setContent('profile')">Profile</a> |
    <a href="#" (click)="setContent('budget')">Budget</a>
  </div>
  <profile-component *ngIf="content == 'profile'"></profile-component>
  <budget-component *ngIf="content == 'budget'"></budget-component>`,
})
export class AppComponent  
{ 

  banner : string
  content : string

  constructor() {
    this.banner = 'Tomato Sauce'
  }

  setContent(c: string) {
    this.content = c
  }
}
