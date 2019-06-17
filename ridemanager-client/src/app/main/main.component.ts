import { Component, OnInit } from '@angular/core';
import {TokenService} from '../services/token.service';
import {Router, ActivatedRoute} from '@angular/router';
import {IUser, UserService} from '../services/user.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss']
})
export class MainComponent implements OnInit {
  token: string;
  pageId: string;
  user: IUser;

  constructor(private tokenService: TokenService,
              private userService: UserService,
              private router: Router,
              private activatedRoute: ActivatedRoute) {
    this.token = tokenService.getToken();
    userService.getMe().subscribe((user: IUser) => {
      this.user = user;
    });
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

    this.activatedRoute.params.subscribe(params => {
      this.pageId = params.pageId;
    });
  }

}
