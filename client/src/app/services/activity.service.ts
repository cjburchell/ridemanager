import { Injectable } from '@angular/core';
import {Gender} from './user.service';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {TokenService} from './token.service';
import {Observable} from 'rxjs';

export type ActivityType = 'group_ride' | 'race' | 'triathlon';
export type ActivityState = 'upcoming' |'in_progress' |'finished';
export type ActivityPrivacy = 'public' | 'private';

export interface IActivity {
  activity_id: string;
  activity_type: ActivityType;
  owner_id: string;
  name: string;
  description: string;
  start_time: Date;
  end_time: Date;
  total_distance: number;
  duration: number;
  time_left: number;
  starts_in: number;
  route: IRoute;
  privacy: ActivityPrivacy;
  categories: ICategory[];
  stages: IStage[];
  participants: IParticipant[];
  state: ActivityState;
  max_participants: number;
}

export interface IRoute {
  id: number;
  name: string;
  distance: number;
}

export interface IParticipant {
  athlete_id: string;
  category_id: string;
  results: IResult[];
  name: string;
  sex: Gender;
  time: string;
  rank: string;
  out_of: string;
  stages: string;
}
export interface IResult {
  segment_id: string;
  time: Date;
}

export interface ICategory {
  category_id: string;
  name: string;
}

export interface IStage {
  segment_id: string;
  distance: number;
}

@Injectable({
  providedIn: 'root'
})
export class ActivityService {

  constructor(private http: HttpClient, private token: TokenService) { }

  createActivity(activity: IActivity): Observable<string> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<string>('api/v1/activity', activity, httpOptions);
  }

  updateActivity(activity: IActivity) {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.patch(`api/v1/activity/${activity.activity_id}`, activity, httpOptions);
  }

  deleteActivity(activity: IActivity) {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.delete(`api/v1/activity/${activity.activity_id}`, httpOptions);
  }

  getActivity(activityId: string): Observable<IActivity> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity>(`api/v1/activity/${activityId}`, httpOptions);
  }

  getActivties(): Observable<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`api/v1/activity`, httpOptions);
  }
}
