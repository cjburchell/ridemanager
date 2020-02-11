import { Component, OnInit } from '@angular/core';
import {ITokenService} from '../services/token.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  constructor(private tokenService: ITokenService,
              private router: Router) {
  }

  async ngOnInit() {
    if (await this.tokenService.checkLogin()) {
      await this.router.navigate([`/main`]);
      return;
    }
  }
}
