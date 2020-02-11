import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';
import {ITokenService} from '../../../services/token.service';

@Component({
  selector: 'app-actions',
  templateUrl: './actions.component.html',
  styleUrls: ['./actions.component.scss']
})
export class ActionsComponent implements OnInit {

  constructor(private router: Router,
              private tokenService: ITokenService) {
  }

  ngOnInit() {
  }

  async showJoin() {
    if (!await this.tokenService.checkLogin()) {
      return;
    }
    await this.router.navigate([`/main/join`]);
  }

  async showCreate() {
    if (!await this.tokenService.checkLogin()) {
      return;
    }
    await this.router.navigate([`/main/create`]);
  }
}
