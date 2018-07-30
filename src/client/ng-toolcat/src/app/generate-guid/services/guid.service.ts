import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class GuidService {

    constructor(private http: HttpClient) {

    }

    url = 'http://127.0.0.1:8090/guid/new';
    new(): Observable<object> {
        return this.http.get(this.url);
    }
}
