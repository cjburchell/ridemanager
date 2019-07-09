import {Component, Input, OnInit} from '@angular/core';
import {
  ActivityService,
  IActivity
} from '../../services/activity.service';
import {IUser} from '../../services/user.service';
import {Router} from '@angular/router';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnInit {

  Activity: IActivity;
  @Input() User: IUser;

  constructor(private activityService: ActivityService, private router: Router) {
  }

  ngOnInit() {
      this.Activity = {
        activity_id: undefined,
        activity_type: 'group_ride',
        owner_id: this.User.id,
        name: 'new',
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

  back() {
    this.router.navigate([`/main`]);
  }

  create() {
    this.activityService.createActivity(this.Activity);
  }
}
