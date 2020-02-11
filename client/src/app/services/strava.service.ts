import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {ITokenService} from './token.service';
import {IRouteSummary, ISegmentSummary} from './contracts/strava';

export abstract class IStravaService {
  public abstract getStaredSegments(page: number, perPage: number): Promise<ISegmentSummary[]>;
  public abstract getRoutes(page: number, perPage: number): Promise<IRouteSummary[]>;
  public abstract getRoute(routeId: number): Promise<IRouteSummary>;
  public abstract getSegment(segmentId: number): Promise<ISegmentSummary>;
}

@Injectable({
  providedIn: 'root'
})
export class StravaService implements IStravaService {

  constructor(private http: HttpClient, private token: ITokenService) { }

  getStaredSegments(page: number, perPage: number): Promise<ISegmentSummary[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<ISegmentSummary[]>(`api/v1/strava/segments/starred?page=${page}&perPage=${perPage}`, httpOptions).toPromise();
  }

  getRoutes(page: number, perPage: number): Promise<IRouteSummary[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IRouteSummary[]>(`api/v1/strava/routes?page=${page}&perPage=${perPage}`, httpOptions).toPromise();
  }

  getRoute(routeId: number): Promise<IRouteSummary> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IRouteSummary>(`api/v1/strava/routes/${routeId}`, httpOptions).toPromise();
  }

  getSegment(segmentId: number): Promise<ISegmentSummary> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<ISegmentSummary>(`api/v1/strava/segments/${segmentId}`, httpOptions).toPromise();
  }
}
