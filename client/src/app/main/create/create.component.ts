import {Component, Input, OnInit} from '@angular/core';
import {
  ActivityService,
  IActivity
} from '../../services/activity.service';
import {IUser} from '../../services/user.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnInit {

  @Input() Activity: IActivity;
  @Input() User: IUser;

  constructor(private activityService: ActivityService) {
  }

  ngOnInit() {
    if (this.Activity === undefined) {
      this.Activity = {
        activity_id: undefined,
        activity_type: 'group_ride',
        owner_id: this.User.id,
        name: undefined,
        description: undefined,
        start_time: undefined,
        end_time: undefined,
        total_distance: undefined,
        duration: undefined,
        time_left: undefined,
        starts_in: undefined,
        privacy: 'private',
        categories: [],
        stages: [],
        participants: [],
        state: undefined,
        max_participants: 10
      };
    }
  }

  back() {
  }

  create() {
    this.activityService.createActivity(this.Activity);
  }
}
