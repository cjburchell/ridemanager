import { Component, OnInit } from '@angular/core';
import {TokenService} from '../services/token.service';
import {Router, ActivatedRoute} from '@angular/router';
import {IAthlete, UserService} from '../services/user.service';
import {HttpErrorResponse} from '@angular/common/http';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.scss']
})
export class MainComponent implements OnInit {
  pageId = 'summery';
  user: IAthlete;

  constructor(private tokenService: TokenService,
              private userService: UserService,
              private activatedRoute: ActivatedRoute) {
  }

  ngOnInit() {
    this.tokenService.checkLogin();

    this.userService.getMe().subscribe((user: IAthlete) => {
      this.user = user;
    });

    this.activatedRoute.params.subscribe(params => {
      this.pageId = params.pageId;
      if (this.pageId === undefined) {
        this.pageId = 'summery';
      }
    });
  }

}
