import { Injectable } from '@angular/core';
import {IAthlete} from './user.service';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {TokenService} from './token.service';
import {Observable} from 'rxjs';
import {IMap} from './strava.service';

export type ActivityType = 'group_ride' | 'race' | 'triathlon' | 'group_run' | 'group_ski';
export type ActivityState = 'upcoming' |'in_progress' |'finished';
export type ActivityPrivacy = 'public' | 'private';

export type SegmentType =
  'Ride'
  | 'AlpineSki'
  | 'BackcountrySki'
  | 'Hike'
  | 'IceSkate'
  | 'InlineSkate'
  | 'NordicSki'
  | 'RollerSki'
  | 'Run'
  | 'Walk'
  | 'Workout'
  | 'Snowboard'
  | 'Snowshoe'
  | 'Kitesurf'
  | 'Windsurf'
  | 'Swim'
  | 'VirtualRide'
  | 'EBikeRide'
  | 'WaterSport'
  | 'Canoeing'
  | 'Kayaking'
  | 'Rowing'
  | 'StandUpPaddling'
  | 'Surfing'
  | 'Crossfit'
  | 'Elliptical'
  | 'RockClimbing'
  | 'StairStepper'
  | 'WeightTraining'
  | 'Yoga'
  | 'WinterSport'
  | 'CrossCountrySkiing';


export interface IActivity {
  activity_id: string;
  activity_type: ActivityType;
  owner: IAthlete;
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
  map: IMap;
}

export interface IParticipant {
  athlete: IAthlete;
  category_id: string;
  results: IResult[];
  time: number;
  rank: number;
  out_of: number;
  stages: number;
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
  segment_id: number;
  distance: number;
  number: number;
  activity_type: SegmentType;
  name: string;
  map: IMap;
  start_latlng: number[];
  end_latlng: number[];
}

@Injectable({
  providedIn: 'root'
})
export class ActivityService {

  constructor(private http: HttpClient, private token: TokenService) {
  }

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

  getActivities(): Observable<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`api/v1/activity/public`, httpOptions);
  }

  getJoined(): Observable<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`api/v1/activity/joined`, httpOptions);
  }

  getMyActivities(): Observable<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`api/v1/activity/my`, httpOptions);
  }

  addParticipant(activity: IActivity, participant: IParticipant): Observable<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<boolean>(`api/v1/activity/${activity.activity_id}/participant`, participant, httpOptions);
  }

  leaveActivity(activity: IActivity, athleteId: string): Observable<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.delete<boolean>(`api/v1/activity/${activity.activity_id}/participant/${athleteId}`, httpOptions);
  }

  updateUserResults(activity: IActivity, athleteId: string): Observable<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<boolean>(`api/v1/activity/${activity.activity_id}/update/${athleteId}`, null, httpOptions);
  }

  updateResults(activity: IActivity): Observable<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<boolean>(`api/v1/activity/${activity.activity_id}/update`, null, httpOptions);
  }
}
