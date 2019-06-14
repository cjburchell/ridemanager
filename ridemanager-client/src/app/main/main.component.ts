import { Component, OnInit } from '@angular/core';
import {TokenService} from '../services/token.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss']
})
export class MainComponent implements OnInit {
  token: string;


  constructor(private tokenService: TokenService,
              private router: Router) {
    this.token = tokenService.getToken();
  }

  ngOnInit() {
    if (this.tokenService.getToken() !== null) {
      this.tokenService.validateToken().subscribe((isLoggedIn: boolean) => {
        if (!isLoggedIn) {
          this.router.navigate([`/login`]);
        }
      });
    } else {
      this.router.navigate([`/login`]);
    }
  }

}
