import {Component, Input, OnInit} from '@angular/core';
import {ActivityService, IActivity} from '../../services/activity.service';
import {IAthlete} from '../../services/user.service';

@Component({
  selector: 'app-summery',
  templateUrl: './summery.component.html',
  styleUrls: ['./summery.component.scss']
})
export class SummeryComponent implements OnInit {

  @Input() user: IAthlete;
  activitiesUpcoming: IActivity[];
  activitiesInProgress: IActivity[];
  activitiesFinished: IActivity[];

  constructor(private activityService: ActivityService) {
  }

  ngOnInit() {
    this.activityService.getJoined().subscribe((activities: IActivity[]) => {
      if (activities !== undefined && activities !== null) {
        this.activitiesUpcoming = activities.filter(item => item.state === 'upcoming');
        this.activitiesInProgress = activities.filter(item => item.state === 'in_progress');
        this.activitiesFinished = activities.filter(item => item.state === 'finished');
      }
    });
  }
}
