import { Component, OnInit } from '@angular/core';
import {SettingsService} from '../../services/settings.service';

@Component({
  selector: 'app-login-button',
  templateUrl: './login-button.component.html',
  styleUrls: ['./login-button.component.scss']
})
export class LoginButtonComponent implements OnInit {

  public stravaClientId = '0';
  public stravaRedirect = 'http://localhost:8091/api/v1/login';
  basePath = 'https://www.strava.com/api/v3';

  constructor(private settingsService: SettingsService) {
  }

  ngOnInit() {
    this.settingsService.getSetting('stravaClientId').subscribe((stravaClientId: string) => {
      this.stravaClientId = stravaClientId;
    });

    this.settingsService.getSetting('stravaRedirect').subscribe((stravaRedirect: string) => {
      this.stravaRedirect = stravaRedirect;
    });
  }

  getAuthorizationURL(): string {
    return `${this.basePath}/oauth/authorize?client_id=${this.stravaClientId}` +
      `&response_type=code&redirect_uri=${this.stravaRedirect}&scope=view_private`;
  }
}