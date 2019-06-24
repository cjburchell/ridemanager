import { Component, OnInit } from '@angular/core';
import {Router} from '@angular/router';

@Component({
  selector: 'app-actions',
  templateUrl: './actions.component.html',
  styleUrls: ['./actions.component.scss']
})
export class ActionsComponent implements OnInit {

  constructor(private router: Router) {
  }

  ngOnInit() {
  }

  showJoin() {
    this.router.navigate([`/main/join`]);
  }

  showCreate() {
    this.router.navigate([`/main/create`]);
  }
}
