import { Component, OnInit } from '@angular/core';
import {ActivityService, IActivity} from '../../services/activity.service';

@Component({
  selector: 'app-history',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.scss']
})
export class HistoryComponent implements OnInit {

  activities: IActivity[];
  constructor(private activityService: ActivityService) { }

  ngOnInit() {
    this.activityService.getJoined().subscribe(activities => this.activities = activities);
  }
}
