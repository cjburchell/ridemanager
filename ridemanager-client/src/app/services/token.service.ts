import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class TokenService {

  private tokenKey = 'app_token';

  constructor(private http: HttpClient) {
  }

  getToken(): string {
    return localStorage.getItem(this.tokenKey);
  }

  setToken(token: string) {
    localStorage.setItem(this.tokenKey, token);
  }

  validateToken(): Observable<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.getToken()
      })
    };

    return this.http.get<boolean>(`api/v1/login/status`, httpOptions);
  }
}
