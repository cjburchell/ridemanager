import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {

  public stravaClientId = 'test';
  public stravaRedirect = 'localhost/main';
  constructor() { }

  ngOnInit() {
    /*this.surveyService.getSetting('stravaClientId').subscribe((stravaClientId: string) => {
      this.stravaClientId = stravaClientId;
    });

    this.surveyService.getSetting('stravaRedirect').subscribe((stravaRedirect: string) => {
      this.stravaRedirect = stravaRedirect;
    });*/
  }
}
