import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {ITokenService} from '../services/token.service';

@Component({
  selector: 'app-token',
  templateUrl: './token.component.html'
})
export class TokenComponent implements OnInit {

  constructor(private tokenService: ITokenService,
              private router: Router,
              private activatedRoute: ActivatedRoute) {
    tokenService.setToken(activatedRoute.snapshot.queryParamMap.get('token'));
  }

  ngOnInit() {
    this.router.navigate([`/main`]);
  }
}
