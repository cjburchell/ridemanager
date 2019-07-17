import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';
import {TokenService} from './token.service';

export type Gender = 'M' | 'F' | '';

export interface IAthlete {
  id: string;
  strava_athlete_id: number;
  first_name: string;
  last_name: string;
  sex: Gender;
  profile: string;
  profile_medium: string;
}

export interface IAchievements {
  first_count: number;
  second_count: number;
  third_count: number;
  finished_count: number;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient,
              private token: TokenService) { }

  getMe(): Observable<IAthlete> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IAthlete>(`api/v1/user/me`, httpOptions);
  }

  getAchievements(): Observable<IAchievements> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IAchievements>(`api/v1/user/me/achievements`, httpOptions);
  }
}
