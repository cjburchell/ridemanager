import { Component, OnInit } from '@angular/core';
import {ActivityService, IActivity} from '../../services/activity.service';

@Component({
  selector: 'app-join',
  templateUrl: './join.component.html',
  styleUrls: ['./join.component.scss']
})
export class JoinComponent implements OnInit {

  activities: IActivity[];
  constructor(private activityService: ActivityService) { }

  ngOnInit() {
    this.activityService.getActivties().subscribe(activities => this.activities = activities);
  }

}
