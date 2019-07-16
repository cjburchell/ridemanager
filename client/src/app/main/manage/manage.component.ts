import { Component, OnInit } from '@angular/core';
import {ActivityService, IActivity} from '../../services/activity.service';

@Component({
  selector: 'app-manage',
  templateUrl: './manage.component.html',
  styleUrls: ['./manage.component.scss']
})
export class ManageComponent implements OnInit {

  activities: IActivity[];
  constructor(private activityService: ActivityService) { }

  ngOnInit() {
    this.activityService.getMyActivities().subscribe(activities => this.activities = activities);
  }
}
