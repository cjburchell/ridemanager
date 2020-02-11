import { Component, OnInit } from '@angular/core';
import {IActivityService} from '../../services/activity.service';
import {IActivity} from '../../services/contracts/activity';

@Component({
  selector: 'app-history',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.scss']
})
export class HistoryComponent implements OnInit {

  activities: IActivity[];
  constructor(private activityService: IActivityService) { }

  async ngOnInit() {
    this.activities = await this.activityService.getJoined();
  }
}
