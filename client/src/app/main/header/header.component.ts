import { Component, Input } from '@angular/core';
import {ITokenService} from '../../services/token.service';
import {Router} from '@angular/router';
import {IAthlete} from '../../services/contracts/user';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.scss']
})
export class HeaderComponent {

  @Input() public user: IAthlete;
  @Input() public isLoggedIn: boolean;

  constructor(private tokenService: ITokenService,
              private router: Router) {
  }

  async onLogOut() {
    this.tokenService.logOut();
    await this.router.navigate([`/login`]);
  }

  async showToMain($event: MouseEvent) {
    $event.preventDefault();
    if (!await this.tokenService.checkLogin()) {
      return;
    }
    await this.router.navigate([`/main`]);
  }

  async showManageActivities($event: MouseEvent) {
    $event.preventDefault();
    if (!await this.tokenService.checkLogin()) {
      return;
    }
    await this.router.navigate([`/main/manage`]);
  }

  async showUserHistory($event: MouseEvent) {
    $event.preventDefault();
    if (!await this.tokenService.checkLogin()) {
      return;
    }
    await this.router.navigate([`/main/history`]);
  }
}
