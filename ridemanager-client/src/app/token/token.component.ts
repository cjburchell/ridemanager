import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {TokenService} from '../services/token.service';

@Component({
  selector: 'app-token',
  templateUrl: './token.component.html',
  styleUrls: ['./token.component.scss']
})
export class TokenComponent implements OnInit {

  constructor(private tokenService: TokenService,
              private router: Router,
              private activatedRoute: ActivatedRoute) {
    tokenService.setToken(activatedRoute.snapshot.queryParamMap.get('token'));
  }

  ngOnInit() {
    this.router.navigate([`/main`]);
  }
}
