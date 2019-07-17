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

  constructor(private tokenService: TokenService,
              private router: Router) {
  }

  ngOnInit() {
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
}
