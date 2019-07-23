import { Component, OnInit } from '@angular/core';
import {TokenService} from '../services/token.service';
import {IAthlete, UserService} from '../services/user.service';
import {ActivatedRoute, Router} from '@angular/router';
import {ActivityService, IActivity} from '../services/activity.service';

@Component({
  selector: 'app-edit',
  templateUrl: './edit.component.html',
  styleUrls: ['./edit.component.scss']
})
export class EditComponent implements OnInit {
  private user: IAthlete;
  private Activity: IActivity;

  constructor(private tokenService: TokenService,
              private router: Router,
              private activityService: ActivityService,
              private userService: UserService,
              private activatedRoute: ActivatedRoute) {
  }

  ngOnInit() {
    this.tokenService.checkLogin();

    this.userService.getMe().subscribe((user: IAthlete) => {
      this.user = user;
    });

    this.activatedRoute.params.subscribe(params => {
      this.getActivity(params.activityId);
    });
  }

  private getActivity(activityId: string) {
    this.activityService.getActivity(activityId).subscribe((activity: IActivity) => {
      if (activity === undefined || activity === null) {
        this.router.navigate([`/main`]);
      } else {
        this.Activity = activity;
      }
    });
  }

  save() {
    this.activityService.updateActivity(this.Activity).subscribe(() => {
        this.router.navigate([`/activity/${this.Activity.activity_id}`]);
      }
    );

  }

  back() {
    this.router.navigate([`/activity/${this.Activity.activity_id}`]);
  }
}
