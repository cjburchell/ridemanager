import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Router} from '@angular/router';

export abstract class ITokenService {
  public abstract checkLogin(): Promise<boolean>;

  public abstract validateToken(): Promise<boolean>;

  public abstract getToken(): string;

  public abstract setToken(token: string);

  public abstract logOut();
}

@Injectable({
  providedIn: 'root'
})
export class TokenService implements ITokenService {

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

  validateToken(): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.getToken()
      })
    };

    return this.http.get<boolean>(`api/v1/login/status`, httpOptions).toPromise();
  }

  async checkLogin(): Promise<boolean> {
    if (this.getToken() !== null) {
      const isLoggedIn = await this.validateToken();
      if (!isLoggedIn) {
        await this.router.navigate([`/login`]);
        return false;
      }
    } else {
      await this.router.navigate([`/login`]);
      return false;
    }
    return true;
  }
}
