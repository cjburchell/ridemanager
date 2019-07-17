import { Component, OnInit } from '@angular/core';
import {IAchievements, UserService} from '../../../services/user.service';

@Component({
  selector: 'app-result-summary',
  templateUrl: './result-summary.component.html',
  styleUrls: ['./result-summary.component.scss']
})
export class ResultSummaryComponent implements OnInit {
  achievements: IAchievements;

  constructor(private userService: UserService) {
  }

  ngOnInit() {
    this.userService.getAchievements().subscribe(achievements => this.achievements = achievements);
  }

}
