"use strict";
var __decorate = (this && this.__decorate) || function (decorators, target, key, desc) {
    var c = arguments.length, r = c < 3 ? target : desc === null ? desc = Object.getOwnPropertyDescriptor(target, key) : desc, d;
    if (typeof Reflect === "object" && typeof Reflect.decorate === "function") r = Reflect.decorate(decorators, target, key, desc);
    else for (var i = decorators.length - 1; i >= 0; i--) if (d = decorators[i]) r = (c < 3 ? d(r) : c > 3 ? d(target, key, r) : d(target, key)) || r;
    return c > 3 && r && Object.defineProperty(target, key, r), r;
};
var __metadata = (this && this.__metadata) || function (k, v) {
    if (typeof Reflect === "object" && typeof Reflect.metadata === "function") return Reflect.metadata(k, v);
};
var core_1 = require("@angular/core");
var ProfileComponent = (function () {
    function ProfileComponent() {
    }
    return ProfileComponent;
}());
ProfileComponent = __decorate([
    core_1.Component({
        selector: 'profile-component',
        template: "<h4>Profile</h4>"
    })
], ProfileComponent);
exports.ProfileComponent = ProfileComponent;
var BudgetComponent = (function () {
    function BudgetComponent() {
    }
    return BudgetComponent;
}());
BudgetComponent = __decorate([
    core_1.Component({
        selector: 'budget-component',
        template: "<h4>Budget</h4>"
    })
], BudgetComponent);
exports.BudgetComponent = BudgetComponent;
var AppComponent = (function () {
    function AppComponent() {
        this.banner = 'Tomato Sauce';
    }
    AppComponent.prototype.setContent = function (c) {
        this.content = c;
    };
    return AppComponent;
}());
AppComponent = __decorate([
    core_1.Component({
        selector: 'main-component',
        template: "\n  <h3 class=\"banner\">{{banner}}</h3>\n  <div>\n    <a href=\"#\" (click)=\"setContent('profile')\">Profile</a> |\n    <a href=\"#\" (click)=\"setContent('budget')\">Budget</a>\n  </div>\n  <profile-component *ngIf=\"content == 'profile'\"></profile-component>\n  <budget-component *ngIf=\"content == 'budget'\"></budget-component>",
    }),
    __metadata("design:paramtypes", [])
], AppComponent);
exports.AppComponent = AppComponent;
//# sourceMappingURL=app.component.js.map