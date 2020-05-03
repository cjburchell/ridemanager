import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

export abstract class ISettingsService {
  public abstract getSetting(setting: string): Promise<string>;
  public abstract getApiUrl(): Promise<string>;
}

@Injectable({
  providedIn: 'root'
})
export class SettingsService implements ISettingsService {
  private apiUrl: string;

  constructor(private http: HttpClient) {
  }

  getSetting(setting: string): Promise<string> {
    return this.http.get<string>(`/client/settings/${setting}`).toPromise();
  }

  async getApiUrl(): Promise<string> {
    if (this.apiUrl === undefined) {
      this.apiUrl = await this.getSetting('apiAddress');
    }

    return new Promise<string>(resolve => resolve(this.apiUrl));
  }
}
