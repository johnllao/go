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
  template: `<h3>{{title}}</h3>`,
})
export class AppComponent  
{ 

  title : string

  constructor() {
    this.title = 'Welcome!'
  }
}
