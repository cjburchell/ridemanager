import {Component, Input, OnInit} from '@angular/core';
import {IActivityService} from '../../services/activity.service';
import {IActivity} from '../../services/contracts/activity';
import {IAthlete} from '../../services/contracts/user';

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

  constructor(private activityService: IActivityService) {
  }

  async ngOnInit() {
    const activities = await this.activityService.getJoined();
    if (activities !== undefined && activities !== null) {
      this.activitiesUpcoming = activities.filter(item => item.state === 'upcoming');
      this.activitiesInProgress = activities.filter(item => item.state === 'in_progress');
      this.activitiesFinished = activities.filter(item => item.state === 'finished');
    }
  }
}
