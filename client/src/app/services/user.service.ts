import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ITokenService} from './token.service';
import {IAchievements, IAthlete} from './contracts/user';

export abstract class IUserService {
  public abstract getMe(): Promise<IAthlete>;
  public abstract getAchievements(): Promise<IAchievements>;
}

@Injectable({
  providedIn: 'root'
})
export class UserService implements IUserService {

  constructor(private http: HttpClient,
              private token: ITokenService) { }

  getMe(): Promise<IAthlete> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IAthlete>(`api/v1/user/me`, httpOptions).toPromise();
  }

  getAchievements(): Promise<IAchievements> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IAchievements>(`api/v1/user/me/achievements`, httpOptions).toPromise();
  }
}
