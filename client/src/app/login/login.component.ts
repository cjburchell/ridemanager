import { Component, OnInit } from '@angular/core';
import {SettingsService} from '../services/settings.service';
import {TokenService} from '../services/token.service';
import {Router} from '@angular/router';
import {HttpErrorResponse} from '@angular/common/http';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  public stravaClientId = '0';
  public stravaRedirect = 'http://localhost:8091/api/v1/login';
  basePath = 'https://www.strava.com/api/v3';

  constructor(private settingsService: SettingsService,
              private tokenService: TokenService,
              private router: Router) {
  }

  ngOnInit() {
    this.settingsService.getSetting('stravaClientId').subscribe((stravaClientId: string) => {
      this.stravaClientId = stravaClientId;
    });

    this.settingsService.getSetting('stravaRedirect').subscribe((stravaRedirect: string) => {
      this.stravaRedirect = stravaRedirect;
    });

    const token = this.tokenService.getToken();
    if (token !== null) {
      this.tokenService.validateToken().subscribe((isLoggedIn: boolean) => {
          if (isLoggedIn) {
            this.router.navigate([`/main`]);
          }
        },
        (err: HttpErrorResponse) => {
          console.log(err);
          this.router.navigate([`/login`]);
        });
    }
  }

  getAuthorizationURL(): string {
    return `${this.basePath}/oauth/authorize?client_id=${this.stravaClientId}` +
           `&response_type=code&redirect_uri=${this.stravaRedirect}&scope=view_private`;
  }
}
