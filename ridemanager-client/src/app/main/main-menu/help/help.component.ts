import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';

@Component({
  selector: 'app-help',
  templateUrl: './help.component.html',
  styleUrls: ['./help.component.scss']
})
export class HelpComponent implements OnInit {

  constructor(private router: Router) { }

  ngOnInit() {
  }

  showFAQ() {
    this.router.navigate([`/main/faq`]);
  }

  showAbout() {
    this.router.navigate([`/main/about`]);
  }
}
