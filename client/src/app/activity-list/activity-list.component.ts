import {Component, Input, OnChanges, OnInit} from '@angular/core';
import {IActivity} from '../services/activity.service';

@Component({
  selector: 'app-activity-list',
  templateUrl: './activity-list.component.html',
  styleUrls: ['./activity-list.component.scss']
})
export class ActivityListComponent implements OnChanges {
  @Input() activities: IActivity[];
  searchText: string;
  activityFilter: string;
  isUpcoming: boolean;
  isInProgress: boolean;
  isFinished: boolean;

  constructor() { }

  ngOnChanges(): void {
    this.isFinished = this.activities.some((item) => {
      return item.state === 'finished';
    });

    this.isInProgress = this.activities.some((item) => {
      return item.state === 'in_progress';
    });

    this.isUpcoming = this.activities.some((item) => {
      return item.state === 'upcoming';
    });
  }
}
