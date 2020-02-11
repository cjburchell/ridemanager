import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export abstract class ISettingsService {
  public abstract getSetting(setting: string): Promise<string>;
}

@Injectable({
  providedIn: 'root'
})
export class SettingsService implements ISettingsService {

  constructor(private http: HttpClient) { }

  getSetting(setting: string): Promise<string> {
    return this.http.get<string>(`api/v1/settings/${setting}`).toPromise();
  }
}
