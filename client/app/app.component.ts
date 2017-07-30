import { Component } from '@angular/core';

@Component({
  selector: 'main-component',
  template: `<h3>{{title}}</h3>`,
})
export class AppComponent { 

  title : string

  constructor() {
    this.title = "Go (golang) Angular Quickstart"
  }
}
