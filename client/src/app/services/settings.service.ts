import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SettingsService {

  constructor(private http: HttpClient) { }

  getSetting(setting: string): Observable<string> {
    return this.http.get<string>(`api/v1/settings/${setting}`);
  }
}
