import { Injectable } from '@angular/core';
import {Observable} from 'rxjs';
import {HttpClient, HttpHeaders} from '@angular/common/http';
import {TokenService} from './token.service';


export interface ISegmentSummary {
  id: number;
  name: string;
  activity_type: string;
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

    return this.http.get<ISegmentSummary[]>(`api/v1/activity`, httpOptions);
  }
}
