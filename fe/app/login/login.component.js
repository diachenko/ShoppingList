"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var core_1 = require("@angular/core");
var router_1 = require("@angular/router");
var page_1 = require("tns-core-modules/ui/page");
var user_model_1 = require("../shared/user/user.model");
var user_service_1 = require("../shared/user/user.service");
var LoginComponent = /** @class */ (function () {
    function LoginComponent(router, userService, page) {
        this.router = router;
        this.userService = userService;
        this.page = page;
        this.isLoggingIn = true;
        this.user = new user_model_1.User();
    }
    LoginComponent.prototype.ngOnInit = function () {
        this.page.actionBarHidden = true;
    };
    LoginComponent.prototype.submit = function () {
        if (this.isLoggingIn) {
            this.login();
        }
        else {
            this.signUp();
        }
    };
    LoginComponent.prototype.login = function () {
        var _this = this;
        this.userService.login(this.user)
            .subscribe(function () { return _this.router.navigate(["/list"]); }, function (error) { return alert("Unfortunately we could not find your account."); });
    };
    LoginComponent.prototype.signUp = function () {
        var _this = this;
        this.userService.register(this.user)
            .subscribe(function () {
            alert("Your account was successfully created.");
            _this.toggleDisplay();
        }, function (error) { return alert("Unfortunately we were unable to create your account."); });
    };
    LoginComponent.prototype.toggleDisplay = function () {
        this.isLoggingIn = !this.isLoggingIn;
    };
    LoginComponent = __decorate([
        core_1.Component({
            selector: "gr-login",
            providers: [user_service_1.UserService],
            styleUrls: ["login/login.component.css"],
            templateUrl: "login/login.component.html"
        }),
        __metadata("design:paramtypes", [router_1.Router, user_service_1.UserService, page_1.Page])
    ], LoginComponent);
    return LoginComponent;
}());
exports.LoginComponent = LoginComponent;
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoibG9naW4uY29tcG9uZW50LmpzIiwic291cmNlUm9vdCI6IiIsInNvdXJjZXMiOlsibG9naW4uY29tcG9uZW50LnRzIl0sIm5hbWVzIjpbXSwibWFwcGluZ3MiOiI7O0FBQUEsc0NBQWtEO0FBQ2xELDBDQUF5QztBQUN6QyxpREFBZ0Q7QUFFaEQsd0RBQWlEO0FBQ2pELDREQUEwRDtBQVExRDtJQUlJLHdCQUFvQixNQUFjLEVBQVUsV0FBd0IsRUFBVSxJQUFVO1FBQXBFLFdBQU0sR0FBTixNQUFNLENBQVE7UUFBVSxnQkFBVyxHQUFYLFdBQVcsQ0FBYTtRQUFVLFNBQUksR0FBSixJQUFJLENBQU07UUFGeEYsZ0JBQVcsR0FBRyxJQUFJLENBQUM7UUFHZixJQUFJLENBQUMsSUFBSSxHQUFHLElBQUksaUJBQUksRUFBRSxDQUFDO0lBQzNCLENBQUM7SUFFRCxpQ0FBUSxHQUFSO1FBQ0ksSUFBSSxDQUFDLElBQUksQ0FBQyxlQUFlLEdBQUcsSUFBSSxDQUFDO0lBQ3JDLENBQUM7SUFFRCwrQkFBTSxHQUFOO1FBQ0ksRUFBRSxDQUFDLENBQUMsSUFBSSxDQUFDLFdBQVcsQ0FBQyxDQUFDLENBQUM7WUFDbkIsSUFBSSxDQUFDLEtBQUssRUFBRSxDQUFDO1FBQ2pCLENBQUM7UUFBQyxJQUFJLENBQUMsQ0FBQztZQUNKLElBQUksQ0FBQyxNQUFNLEVBQUUsQ0FBQztRQUNsQixDQUFDO0lBQ0wsQ0FBQztJQUVELDhCQUFLLEdBQUw7UUFBQSxpQkFNQztRQUxHLElBQUksQ0FBQyxXQUFXLENBQUMsS0FBSyxDQUFDLElBQUksQ0FBQyxJQUFJLENBQUM7YUFDNUIsU0FBUyxDQUNOLGNBQU0sT0FBQSxLQUFJLENBQUMsTUFBTSxDQUFDLFFBQVEsQ0FBQyxDQUFDLE9BQU8sQ0FBQyxDQUFDLEVBQS9CLENBQStCLEVBQ3JDLFVBQUMsS0FBSyxJQUFLLE9BQUEsS0FBSyxDQUFDLCtDQUErQyxDQUFDLEVBQXRELENBQXNELENBQ3BFLENBQUM7SUFDVixDQUFDO0lBRUQsK0JBQU0sR0FBTjtRQUFBLGlCQVNDO1FBUkcsSUFBSSxDQUFDLFdBQVcsQ0FBQyxRQUFRLENBQUMsSUFBSSxDQUFDLElBQUksQ0FBQzthQUMvQixTQUFTLENBQ047WUFDSSxLQUFLLENBQUMsd0NBQXdDLENBQUMsQ0FBQztZQUNoRCxLQUFJLENBQUMsYUFBYSxFQUFFLENBQUM7UUFDekIsQ0FBQyxFQUNELFVBQUMsS0FBSyxJQUFLLE9BQUEsS0FBSyxDQUFDLHNEQUFzRCxDQUFDLEVBQTdELENBQTZELENBQzNFLENBQUM7SUFDVixDQUFDO0lBRUQsc0NBQWEsR0FBYjtRQUNJLElBQUksQ0FBQyxXQUFXLEdBQUcsQ0FBQyxJQUFJLENBQUMsV0FBVyxDQUFDO0lBQ3pDLENBQUM7SUF6Q1EsY0FBYztRQU4xQixnQkFBUyxDQUFDO1lBQ1AsUUFBUSxFQUFFLFVBQVU7WUFDcEIsU0FBUyxFQUFFLENBQUMsMEJBQVcsQ0FBQztZQUN4QixTQUFTLEVBQUUsQ0FBQywyQkFBMkIsQ0FBQztZQUN4QyxXQUFXLEVBQUUsNEJBQTRCO1NBQzVDLENBQUM7eUNBSzhCLGVBQU0sRUFBdUIsMEJBQVcsRUFBZ0IsV0FBSTtPQUovRSxjQUFjLENBMEMxQjtJQUFELHFCQUFDO0NBQUEsQUExQ0QsSUEwQ0M7QUExQ1ksd0NBQWMiLCJzb3VyY2VzQ29udGVudCI6WyJpbXBvcnQgeyBDb21wb25lbnQsIE9uSW5pdCB9IGZyb20gXCJAYW5ndWxhci9jb3JlXCI7XG5pbXBvcnQgeyBSb3V0ZXIgfSBmcm9tIFwiQGFuZ3VsYXIvcm91dGVyXCI7XG5pbXBvcnQgeyBQYWdlIH0gZnJvbSBcInRucy1jb3JlLW1vZHVsZXMvdWkvcGFnZVwiO1xuXG5pbXBvcnQgeyBVc2VyIH0gZnJvbSBcIi4uL3NoYXJlZC91c2VyL3VzZXIubW9kZWxcIjtcbmltcG9ydCB7IFVzZXJTZXJ2aWNlIH0gZnJvbSBcIi4uL3NoYXJlZC91c2VyL3VzZXIuc2VydmljZVwiO1xuXG5AQ29tcG9uZW50KHtcbiAgICBzZWxlY3RvcjogXCJnci1sb2dpblwiLFxuICAgIHByb3ZpZGVyczogW1VzZXJTZXJ2aWNlXSxcbiAgICBzdHlsZVVybHM6IFtcImxvZ2luL2xvZ2luLmNvbXBvbmVudC5jc3NcIl0sXG4gICAgdGVtcGxhdGVVcmw6IFwibG9naW4vbG9naW4uY29tcG9uZW50Lmh0bWxcIlxufSlcbmV4cG9ydCBjbGFzcyBMb2dpbkNvbXBvbmVudCBpbXBsZW1lbnRzIE9uSW5pdCB7XG4gICAgdXNlcjogVXNlcjtcbiAgICBpc0xvZ2dpbmdJbiA9IHRydWU7XG5cbiAgICBjb25zdHJ1Y3Rvcihwcml2YXRlIHJvdXRlcjogUm91dGVyLCBwcml2YXRlIHVzZXJTZXJ2aWNlOiBVc2VyU2VydmljZSwgcHJpdmF0ZSBwYWdlOiBQYWdlKSB7XG4gICAgICAgIHRoaXMudXNlciA9IG5ldyBVc2VyKCk7XG4gICAgfVxuXG4gICAgbmdPbkluaXQoKSB7XG4gICAgICAgIHRoaXMucGFnZS5hY3Rpb25CYXJIaWRkZW4gPSB0cnVlO1xuICAgIH1cblxuICAgIHN1Ym1pdCgpIHtcbiAgICAgICAgaWYgKHRoaXMuaXNMb2dnaW5nSW4pIHtcbiAgICAgICAgICAgIHRoaXMubG9naW4oKTtcbiAgICAgICAgfSBlbHNlIHtcbiAgICAgICAgICAgIHRoaXMuc2lnblVwKCk7XG4gICAgICAgIH1cbiAgICB9XG5cbiAgICBsb2dpbigpIHtcbiAgICAgICAgdGhpcy51c2VyU2VydmljZS5sb2dpbih0aGlzLnVzZXIpXG4gICAgICAgICAgICAuc3Vic2NyaWJlKFxuICAgICAgICAgICAgICAgICgpID0+IHRoaXMucm91dGVyLm5hdmlnYXRlKFtcIi9saXN0XCJdKSxcbiAgICAgICAgICAgICAgICAoZXJyb3IpID0+IGFsZXJ0KFwiVW5mb3J0dW5hdGVseSB3ZSBjb3VsZCBub3QgZmluZCB5b3VyIGFjY291bnQuXCIpXG4gICAgICAgICAgICApO1xuICAgIH1cblxuICAgIHNpZ25VcCgpIHtcbiAgICAgICAgdGhpcy51c2VyU2VydmljZS5yZWdpc3Rlcih0aGlzLnVzZXIpXG4gICAgICAgICAgICAuc3Vic2NyaWJlKFxuICAgICAgICAgICAgICAgICgpID0+IHtcbiAgICAgICAgICAgICAgICAgICAgYWxlcnQoXCJZb3VyIGFjY291bnQgd2FzIHN1Y2Nlc3NmdWxseSBjcmVhdGVkLlwiKTtcbiAgICAgICAgICAgICAgICAgICAgdGhpcy50b2dnbGVEaXNwbGF5KCk7XG4gICAgICAgICAgICAgICAgfSxcbiAgICAgICAgICAgICAgICAoZXJyb3IpID0+IGFsZXJ0KFwiVW5mb3J0dW5hdGVseSB3ZSB3ZXJlIHVuYWJsZSB0byBjcmVhdGUgeW91ciBhY2NvdW50LlwiKVxuICAgICAgICAgICAgKTtcbiAgICB9XG5cbiAgICB0b2dnbGVEaXNwbGF5KCkge1xuICAgICAgICB0aGlzLmlzTG9nZ2luZ0luID0gIXRoaXMuaXNMb2dnaW5nSW47XG4gICAgfVxufVxuIl19