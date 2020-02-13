import {Injectable} from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Router} from '@angular/router';

export abstract class ITokenService {
  public abstract async checkLogin(): Promise<boolean>;

  public abstract async validateToken(): Promise<boolean>;

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

  async validateToken(): Promise<boolean> {
    if (this.getToken() === null) {
      return false;
    }
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.getToken()
      })
    };

    let isValid = false;
    try {
      isValid = await this.http.get<boolean>(`api/v1/login/status`, httpOptions).toPromise();
    } catch (e) {
    }

    if (!isValid) {
      this.setToken(null);
    }

    return isValid;
  }

  async checkLogin(): Promise<boolean> {
    const isLoggedIn = await this.validateToken();
    if (!isLoggedIn) {
      await this.router.navigate([`/login`]);
      return false;
    }
    return true;
  }
}
