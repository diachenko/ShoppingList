import { Injectable } from "@angular/core";
import { Http, Headers, Response } from "@angular/http";
import { Observable } from "rxjs";
import { catchError, map, tap } from "rxjs/operators";

import { User } from "./user.model";
import { Config } from "../config";

@Injectable()
export class UserService {
    constructor(private http: Http) { }

    login(user: User) {
        console.log(user.name, user.pass)
        return this.http.post(
            Config.apiUrl + "signin",
            JSON.stringify({
                name: user.name,
                pass: user.pass
            })
        ).pipe(
            map(response => response.json()),
            tap(data => {
                console.log(data.Token)
                Config.token = data.Token
            }),
            catchError(this.handleErrors)
        );
    }

    register(user: User) {
        return this.http.post(
            Config.apiUrl + "signup",
            JSON.stringify({
                name: user.name,
                pass: user.pass
            })
        ).pipe(
            map(response => response.json()),
            catchError(this.handleErrors)
        );
    }

    getCommonHeaders() {
        let headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("auth", Config.token)
        return headers;
    }

    handleErrors(error: Response) {
        console.log(JSON.stringify(error.json()));
        return Observable.throw(error);
    }
}