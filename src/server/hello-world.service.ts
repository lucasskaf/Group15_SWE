import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../environments/environment';
import 'rxjs/add/operator/map';

@Injectable()
export class HelloWorldService {

  constructor(private http: HttpClient) { }

  getTitle() {
    return this.http.get(`${environment.serverUrl}/hello-world`)
      .map(response => response.json());
  }

}