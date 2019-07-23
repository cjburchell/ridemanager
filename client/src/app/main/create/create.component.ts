import { Component, Input, OnChanges} from '@angular/core';
import {
  ActivityService,
  IActivity
} from '../../services/activity.service';
import {IAthlete} from '../../services/user.service';
import {Router} from '@angular/router';
import {TokenService} from '../../services/token.service';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnChanges {

  Activity: IActivity;
  @Input() user: IAthlete;

  constructor(private tokenService: TokenService,
              private router: Router,
              private activityService: ActivityService) {
  }


  ngOnChanges() {
    if (this.user !== undefined) {
      this.Activity = {
        activity_id: undefined,
        activity_type: 'group_ride',
        owner: this.user,
        name: undefined,
        description: undefined,
        start_time: new Date(),
        end_time: new Date(),
        total_distance: undefined,
        duration: undefined,
        time_left: undefined,
        starts_in: undefined,
        route: undefined,
        privacy: 'private',
        categories: [],
        stages: [],
        participants: [],
        state: 'upcoming',
        max_participants: 10
      };

      this.Activity.end_time.setDate(this.Activity.end_time.getDate() + 7);
    }
  }

  back() {
    this.tokenService.checkLogin();
    this.router.navigate([`/main`]);
  }

  create() {
    this.activityService.createActivity(this.Activity).subscribe(result => {
      if (result !== undefined && result !== null) {
        this.tokenService.checkLogin();
        this.router.navigate([`/main`]);
      }
    }, error1 => console.log(error1));
  }
}
