import { Component, OnInit } from '@angular/core';
import {IActivityService} from '../../services/activity.service';
import {IActivity} from '../../services/contracts/activity';

@Component({
  selector: 'app-manage',
  templateUrl: './manage.component.html',
  styleUrls: ['./manage.component.scss']
})
export class ManageComponent implements OnInit {

  activities: IActivity[];

  constructor(private activityService: IActivityService) {
  }

  async ngOnInit() {
    this.activities = await this.activityService.getMyActivities();
  }
}
