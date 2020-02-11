import { Component, OnInit } from '@angular/core';
import { IActivityService} from '../../services/activity.service';
import {IActivity} from '../../services/contracts/activity';

@Component({
  selector: 'app-join',
  templateUrl: './join.component.html',
  styleUrls: ['./join.component.scss']
})
export class JoinComponent implements OnInit {

  activities: IActivity[];
  constructor(private activityService: IActivityService) { }

  async ngOnInit() {
    this.activities = await this.activityService.getActivities();
  }

}
