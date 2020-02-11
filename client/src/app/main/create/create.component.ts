import { Component, Input, OnChanges} from '@angular/core';
import {Router} from '@angular/router';
import * as uuid from 'uuid';
import {IActivity} from '../../services/contracts/activity';
import {ITokenService} from '../../services/token.service';
import {IActivityService} from '../../services/activity.service';
import {IAthlete} from '../../services/contracts/user';

@Component({
  selector: 'app-create',
  templateUrl: './create.component.html',
  styleUrls: ['./create.component.scss']
})
export class CreateComponent implements OnChanges {

  Activity: IActivity;
  @Input() user: IAthlete;

  constructor(private tokenService: ITokenService,
              private router: Router,
              private activityService: IActivityService) {
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
        categories: [{
          category_id: uuid.v4().toString(),
          name: 'Open',
        }],
        stages: [],
        participants: [],
        state: 'upcoming',
        max_participants: 10
      };

      this.Activity.end_time.setDate(this.Activity.end_time.getDate() + 7);
    }
  }

  async back() {
    if (await this.tokenService.checkLogin()) {
      await this.router.navigate([`/main`]);
    }
  }

  async create() {
    const result = await this.activityService.createActivity(this.Activity);
    if (result !== undefined && result !== null) {
      if (!await this.tokenService.checkLogin()) {
        return;
      }
      await this.router.navigate([`/main`]);
    }
  }
}
