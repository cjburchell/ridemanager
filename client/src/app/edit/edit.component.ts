import { Component, OnInit } from '@angular/core';
import {IUserService} from '../services/user.service';
import {ActivatedRoute, Router} from '@angular/router';
import {IActivity} from '../services/contracts/activity';
import {ITokenService} from '../services/token.service';
import {IActivityService} from '../services/activity.service';
import {IAthlete} from '../services/contracts/user';

@Component({
  selector: 'app-edit',
  templateUrl: './edit.component.html',
  styleUrls: ['./edit.component.scss']
})
export class EditComponent implements OnInit {
  private user: IAthlete;
  public Activity: IActivity;

  constructor(private tokenService: ITokenService,
              private router: Router,
              private activityService: IActivityService,
              private userService: IUserService,
              private activatedRoute: ActivatedRoute) {
  }

  async ngOnInit() {
    if (await this.tokenService.checkLogin()) {
      return;
    }

    this.user = await this.userService.getMe();
    this.activatedRoute.params.subscribe(async params => {
      await this.getActivity(params.activityId);
    });
  }

   private async getActivity(activityId: string) {
    this.Activity = await this.activityService.getActivity(activityId);
    if (this.Activity === undefined || this.Activity === null) {
      await this.router.navigate([`/main`]);
    }
  }

  async save() {
    await this.activityService.updateActivity(this.Activity);
    await this.router.navigate([`/activity/${this.Activity.activity_id}`]);
  }

  async back() {
    await this.router.navigate([`/activity/${this.Activity.activity_id}`]);
  }
}
