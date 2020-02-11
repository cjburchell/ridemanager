import { Component, OnInit } from '@angular/core';
import {ActivatedRoute} from '@angular/router';
import {ITokenService} from '../services/token.service';
import {IUserService} from '../services/user.service';
import {IAthlete} from '../services/contracts/user';


@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss']
})
export class MainComponent implements OnInit {
  pageId = 'summery';
  user: IAthlete;

  constructor(private tokenService: ITokenService,
              private userService: IUserService,
              private activatedRoute: ActivatedRoute) {
  }

  async ngOnInit() {
    if (!await this.tokenService.checkLogin()) {
      return;
    }

    this.user = await this.userService.getMe();

    this.activatedRoute.params.subscribe(params => {
      this.pageId = params.pageId;
      if (this.pageId === undefined) {
        this.pageId = 'summery';
      }
    });
  }
}
