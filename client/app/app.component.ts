import { Component } from '@angular/core';

@Component({
  selector: 'main-component',
  template: `<h3 class="banner">{{banner}}</h3>`,
})
export class AppComponent  
{ 
  banner : string;
  
  constructor() {
    this.banner = 'Tomato Sauce'
  }
}
