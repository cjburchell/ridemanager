import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';
import {TokenService} from "../../../services/token.service";

@Component({
  selector: 'app-help',
  templateUrl: './help.component.html',
  styleUrls: ['./help.component.scss']
})
export class HelpComponent implements OnInit {

  constructor(private router: Router, private tokenService: TokenService ) { }

  ngOnInit() {
  }

  showFAQ() {
    this.tokenService.checkLogin();
    this.router.navigate([`/main/faq`]);
  }

  showAbout() {
    this.tokenService.checkLogin();
    this.router.navigate([`/main/about`]);
  }
}
