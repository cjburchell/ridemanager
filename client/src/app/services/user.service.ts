import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {Observable} from 'rxjs';
import {TokenService} from './token.service';

export type Role = 'user' | 'admin';
export type Gender = 'M' | 'F' | '';

export interface IUser {
  strava_athlete_id: number;
  role: Role;
  max_active_activities: number;
  first_name: string;
  last_name: string;
  sex: Gender;
  profile: string;
  profile_medium: string;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient,
              private token: TokenService) { }

  getMe(): Observable<IUser> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IUser>(`api/v1/user/me`, httpOptions);
  }
}