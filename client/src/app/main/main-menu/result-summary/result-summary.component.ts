import { Component, OnInit } from '@angular/core';
import {IUserService} from '../../../services/user.service';
import {IAchievements} from '../../../services/contracts/user';

@Component({
  selector: 'app-result-summary',
  templateUrl: './result-summary.component.html',
  styleUrls: ['./result-summary.component.scss']
})
export class ResultSummaryComponent implements OnInit {
  achievements: IAchievements;

  constructor(private userService: IUserService) {
  }

  async ngOnInit() {
    this.achievements = await this.userService.getAchievements();
  }

}
