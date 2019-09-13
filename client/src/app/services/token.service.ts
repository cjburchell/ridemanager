import { Injectable } from '@angular/core';
import {HttpClient, HttpErrorResponse, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';
import {Router} from '@angular/router';

@Injectable({
  providedIn: 'root'
})
export class TokenService {

  private tokenKey = 'app_token';

  constructor(private http: HttpClient,
              private router: Router) {
  }

  getToken(): string {
    return localStorage.getItem(this.tokenKey);
  }

  setToken(token: string) {
    localStorage.setItem(this.tokenKey, token);
  }

  logOut() {
    localStorage.removeItem(this.tokenKey);
  }

  validateToken(): Observable<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.getToken()
      })
    };

    return this.http.get<boolean>(`api/v1/login/status`, httpOptions);
  }

  checkLogin() {
    if (this.getToken() !== null) {
      this.validateToken().subscribe((isLoggedIn: boolean) => {
        if (!isLoggedIn) {
          this.router.navigate([`/login`]);
        }
      }, (err: HttpErrorResponse) => {
        console.log(err);
        this.router.navigate([`/login`]);
      });
    } else {
      this.router.navigate([`/login`]);
    }
  }
}
