import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';
import {TokenService} from '../../../services/token.service';

@Component({
  selector: 'app-actions',
  templateUrl: './actions.component.html',
  styleUrls: ['./actions.component.scss']
})
export class ActionsComponent implements OnInit {

  constructor(private router: Router,
              private tokenService: TokenService) {
  }

  ngOnInit() {
  }

  showJoin() {
    this.tokenService.checkLogin();
    this.router.navigate([`/main/join`]);
  }

  showCreate() {
    this.tokenService.checkLogin();
    this.router.navigate([`/main/create`]);
  }
}
