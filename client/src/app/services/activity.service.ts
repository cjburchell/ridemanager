import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ITokenService} from './token.service';
import {IActivity, IParticipant} from './contracts/activity';

export abstract class IActivityService {
  public abstract createActivity(activity: IActivity): Promise<string>;

  public abstract updateActivity(activity: IActivity): Promise<any>;

  public abstract deleteActivity(activity: IActivity): Promise<any>;

  public abstract getActivity(activityId: string): Promise<IActivity>;

  public abstract getActivities(): Promise<IActivity[]>;

  public abstract getJoined(): Promise<IActivity[]>;

  public abstract getMyActivities(): Promise<IActivity[]>;

  public abstract addParticipant(activity: IActivity, participant: IParticipant): Promise<boolean>;

  public abstract leaveActivity(activity: IActivity, athleteId: string): Promise<boolean>;

  public abstract updateUserResults(activity: IActivity, athleteId: string): Promise<boolean>;

  public abstract updateResults(activity: IActivity): Promise<boolean>;
}


@Injectable({
  providedIn: 'root'
})
export class ActivityService implements IActivityService {

  constructor(private http: HttpClient, private token: ITokenService) {
  }

  createActivity(activity: IActivity): Promise<string> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<string>('api/v1/activity', activity, httpOptions).toPromise();
  }

  updateActivity(activity: IActivity): Promise<any> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.patch(`api/v1/activity/${activity.activity_id}`, activity, httpOptions).toPromise();
  }

  deleteActivity(activity: IActivity): Promise<any> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.delete(`api/v1/activity/${activity.activity_id}`, httpOptions).toPromise();
  }

  getActivity(activityId: string): Promise<IActivity> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity>(`api/v1/activity/${activityId}`, httpOptions).toPromise();
  }

  getActivities(): Promise<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`api/v1/activity/public`, httpOptions).toPromise();
  }

  getJoined(): Promise<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`api/v1/activity/joined`, httpOptions).toPromise();
  }

  getMyActivities(): Promise<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`api/v1/activity/my`, httpOptions).toPromise();
  }

  addParticipant(activity: IActivity, participant: IParticipant): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<boolean>(`api/v1/activity/${activity.activity_id}/participant`, participant, httpOptions).toPromise();
  }

  leaveActivity(activity: IActivity, athleteId: string): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.delete<boolean>(`api/v1/activity/${activity.activity_id}/participant/${athleteId}`, httpOptions).toPromise();
  }

  updateUserResults(activity: IActivity, athleteId: string): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<boolean>(`api/v1/activity/${activity.activity_id}/update/${athleteId}`, null, httpOptions).toPromise();
  }

  updateResults(activity: IActivity): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<boolean>(`api/v1/activity/${activity.activity_id}/update`, null, httpOptions).toPromise();
  }
}
