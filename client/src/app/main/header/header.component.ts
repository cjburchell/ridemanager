import { Component, OnInit, Input } from '@angular/core';
import {IAthlete, UserService} from '../../services/user.service';
import {TokenService} from '../../services/token.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent implements OnInit {

  @Input() public user: IAthlete;

  constructor(private tokenService: TokenService,
              private router: Router) {
  }

  ngOnInit() {
  }

  onLogOut() {
    this.tokenService.logOut();
    this.router.navigate([`/login`]);
  }

  showToMain() {
    this.tokenService.checkLogin();
    this.router.navigate([`/main`]);
  }

  showManageActivities() {
    this.tokenService.checkLogin();
    this.router.navigate([`/main/manage`]);
  }

  showUserHistory() {
    this.tokenService.checkLogin();
    this.router.navigate([`/main/history`]);
  }
}
