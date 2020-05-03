import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ITokenService} from './token.service';
import {IActivity, IParticipant} from './contracts/activity';
import {ISettingsService} from './settings.service';

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

  constructor(private http: HttpClient, private token: ITokenService, private settings: ISettingsService) {
  }

  async createActivity(activity: IActivity): Promise<string> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<string>(`${await this.settings.getApiUrl()}/activity`, activity, httpOptions).toPromise();
  }

  async updateActivity(activity: IActivity): Promise<any> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.patch( `${await this.settings.getApiUrl()}/activity/${activity.activity_id}`, activity, httpOptions).toPromise();
  }

  async deleteActivity(activity: IActivity): Promise<any> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.delete(`${await this.settings.getApiUrl()}/activity/${activity.activity_id}`, httpOptions).toPromise();
  }

  async getActivity(activityId: string): Promise<IActivity> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity>(`${await this.settings.getApiUrl()}/activity/${activityId}`, httpOptions).toPromise();
  }

  async getActivities(): Promise<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`${await this.settings.getApiUrl()}/activity/public`, httpOptions).toPromise();
  }

  async getJoined(): Promise<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`${await this.settings.getApiUrl()}/activity/joined`, httpOptions).toPromise();
  }

  async getMyActivities(): Promise<IActivity[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IActivity[]>(`${await this.settings.getApiUrl()}/activity/my`, httpOptions).toPromise();
  }

  async addParticipant(activity: IActivity, participant: IParticipant): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<boolean>(
      `${await this.settings.getApiUrl()}/activity/${activity.activity_id}/participant`, participant, httpOptions).toPromise();
  }

  async leaveActivity(activity: IActivity, athleteId: string): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.delete<boolean>(
      `${await this.settings.getApiUrl()}/activity/${activity.activity_id}/participant/${athleteId}`, httpOptions).toPromise();
  }

  async updateUserResults(activity: IActivity, athleteId: string): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<boolean>(
      `${await this.settings.getApiUrl()}/activity/${activity.activity_id}/update/${athleteId}`, null, httpOptions).toPromise();
  }

  async updateResults(activity: IActivity): Promise<boolean> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.post<boolean>(
      `${await this.settings.getApiUrl()}/activity/${activity.activity_id}/update`, null, httpOptions).toPromise();
  }
}
