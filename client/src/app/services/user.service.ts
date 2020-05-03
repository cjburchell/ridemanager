import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ITokenService} from './token.service';
import {IAchievements, IAthlete} from './contracts/user';
import {ISettingsService} from './settings.service';

export abstract class IUserService {
  public abstract getMe(): Promise<IAthlete>;
  public abstract getAchievements(): Promise<IAchievements>;
}

@Injectable({
  providedIn: 'root'
})
export class UserService implements IUserService {

  constructor(private http: HttpClient,
              private token: ITokenService,
              private settings: ISettingsService) { }

  async getMe(): Promise<IAthlete> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IAthlete>(`${await this.settings.getApiUrl()}/user/me`, httpOptions).toPromise();
  }

  async getAchievements(): Promise<IAchievements> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IAchievements>(`${await this.settings.getApiUrl()}/user/me/achievements`, httpOptions).toPromise();
  }
}
