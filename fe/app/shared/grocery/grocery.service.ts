import { Injectable } from "@angular/core";
import { Http, Headers, Response, URLSearchParams } from "@angular/http";
import { Observable } from "rxjs";
import { catchError, map, tap } from "rxjs/operators";

import { Config } from "../config";
import { Grocery } from "./grocery.model";

@Injectable()
export class GroceryService {
    getUrl = Config.apiUrl + "productList";
    addUrl = Config.apiUrl + "product"
    deleteUrl = Config.apiUrl + "product"

    constructor(private http: Http) { }

    load() {
        return this.http.get(this.getUrl, {
            headers: this.getCommonHeaders()
        }).pipe(
            map(res => res.json()),
            tap(data => {
                let groceryList = [];
                data.forEach((grocery) => {
                    console.log(data)
                    groceryList.push(new Grocery(grocery._id, grocery.Name, grocery.isBought));
                });
                return groceryList;
            }),
            catchError(this.handleErrors)
        );
    }

    add(name: string) {
        return this.http.post(
            this.addUrl,
            JSON.stringify({ Name: name }),
            { headers: this.getCommonHeaders() }
        ).pipe(
            map(res => res.json()),
            map(data => {
                return new Grocery(data._id, data.name, data.isBought);
            }),
            catchError(this.handleErrors)
        );
    }

    delete(id: string) {
        return this.http.delete(
            this.deleteUrl + "/" + id,
            { headers: this.getCommonHeaders() }
        ).pipe(
                catchError(this.handleErrors)
        );
    }

    markBought (id: string) {
        return this.http.put(
            this.deleteUrl + "/" + id,
            { headers: this.getCommonHeaders() }
        ).pipe(
            catchError(this.handleErrors)
        );
    }

    getCommonHeaders() {
        let headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("auth", Config.token);
        return headers;
    }

    handleErrors(error: Response) {
        console.log(error);
        return Observable.throw(error);
    }
}
