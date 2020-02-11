import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';
import {ITokenService} from '../../../services/token.service';

@Component({
  selector: 'app-help',
  templateUrl: './help.component.html',
  styleUrls: ['./help.component.scss']
})
export class HelpComponent implements OnInit {

  constructor(private router: Router, private tokenService: ITokenService ) { }

  ngOnInit() {
  }

  async showFAQ() {
    if (!await this.tokenService.checkLogin()) {
      return;
    }
    await this.router.navigate([`/main/faq`]);
  }

  async showAbout() {
    if (!await this.tokenService.checkLogin()) {
      return;
    }
    await this.router.navigate([`/main/about`]);
  }
}
