import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ITokenService} from './token.service';
import {IRouteSummary, ISegmentSummary} from './contracts/strava';
import {IPoint} from './contracts/activity';
import {ISettingsService} from './settings.service';

export abstract class IStravaService {
  public abstract getStaredSegments(page: number, perPage: number): Promise<ISegmentSummary[]>;
  public abstract getRoutes(page: number, perPage: number): Promise<IRouteSummary[]>;
  public abstract getRoute(routeId: number): Promise<IRouteSummary>;
  public abstract getSegment(segmentId: number): Promise<ISegmentSummary>;
  public abstract getRouteMap(routeId: number): Promise<IPoint[]>;
  public abstract getSegmentMap(segmentId: number): Promise<IPoint[]>;
}

@Injectable({
  providedIn: 'root'
})
export class StravaService implements IStravaService {
    async getRouteMap(routeId: number): Promise<IPoint[]> {
      const httpOptions = {
        headers: new HttpHeaders({
          Authorization: 'Bearer ' + this.token.getToken()
        })
      };

      return this.http.get<IPoint[]>(`${await this.settings.getApiUrl()}/strava/routes/${routeId}/map`, httpOptions).toPromise();
    }

    async getSegmentMap(segmentId: number): Promise<IPoint[]> {
      const httpOptions = {
        headers: new HttpHeaders({
          Authorization: 'Bearer ' + this.token.getToken()
        })
      };

      return this.http.get<IPoint[]>(`${await this.settings.getApiUrl()}/strava/segments/${segmentId}/map`, httpOptions).toPromise();
    }

  constructor(private http: HttpClient, private token: ITokenService, private settings: ISettingsService) { }

  async getStaredSegments(page: number, perPage: number): Promise<ISegmentSummary[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<ISegmentSummary[]>(
      `${await this.settings.getApiUrl()}/strava/segments/starred?page=${page}&perPage=${perPage}`, httpOptions).toPromise();
  }

  async getRoutes(page: number, perPage: number): Promise<IRouteSummary[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IRouteSummary[]>(
      `${await this.settings.getApiUrl()}/strava/routes?page=${page}&perPage=${perPage}`, httpOptions).toPromise();
  }

  async getRoute(routeId: number): Promise<IRouteSummary> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IRouteSummary>(`${await this.settings.getApiUrl()}/strava/routes/${routeId}`, httpOptions).toPromise();
  }

  async getSegment(segmentId: number): Promise<ISegmentSummary> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<ISegmentSummary>(`${await this.settings.getApiUrl()}/strava/segments/${segmentId}`, httpOptions).toPromise();
  }
}
