import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {TokenService} from './token.service';
import {SegmentType} from './activity.service';


export interface ISegmentSummary {
  id: number;
  name: string;
  activity_type: SegmentType;
  distance: number;
  average_grade: number;
  maximum_grade: number;
  elevation_high: number;
  elevation_low: number;
  climb_category: number;
  start_latlng: number[];
  end_latlng: number[];
  city: string;
  state: string;
  country: string;
  private: boolean;
  starred: boolean;
  map: IMap;
}

export interface IMap {
  id: string;
  polyline: string;
  summary_polyline: string;
}


export interface IRouteSummary {
  id: number;
  name: string;
  distance: number;
  type: number;
  segments: ISegmentSummary[];
  map: IMap;
}

@Injectable({
  providedIn: 'root'
})
export class StravaService {

  constructor(private http: HttpClient, private token: TokenService) { }

  getStaredSegments(): Observable<ISegmentSummary[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<ISegmentSummary[]>(`api/v1/strava/segments/starred`, httpOptions);
  }

  getRoutes(): Observable<IRouteSummary[]> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IRouteSummary[]>(`api/v1/strava/routes`, httpOptions);
  }

  getRoute(routeId: number): Observable<IRouteSummary> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<IRouteSummary>(`api/v1/strava/routes/` + routeId, httpOptions);
  }

  getSegment(segmentId: number): Observable<ISegmentSummary> {
    const httpOptions = {
      headers: new HttpHeaders({
        Authorization: 'Bearer ' + this.token.getToken()
      })
    };

    return this.http.get<ISegmentSummary>(`api/v1/strava/segments/` + segmentId, httpOptions);
  }
}
