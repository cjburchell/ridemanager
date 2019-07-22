import { Component, OnInit } from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {ActivityService, IActivity} from '../services/activity.service';
import {TokenService} from '../services/token.service';
import {IAthlete, UserService} from '../services/user.service';
import {HttpErrorResponse} from "@angular/common/http";

@Component({
  selector: 'app-activity',
  templateUrl: './activity.component.html',
  styleUrls: ['./activity.component.scss']
})
export class ActivityComponent implements OnInit {
  public activity: IActivity;
  public isLoggedIn: boolean;
  public user: IAthlete;
  private isParticipant: boolean;


  constructor(private activatedRoute: ActivatedRoute,
              private activityService: ActivityService,
              private router: Router,
              private tokenService: TokenService,
              private userService: UserService) {
  }

  ngOnInit() {

    this.isLoggedIn = true;
    if (this.tokenService.getToken() !== null) {
      this.tokenService.validateToken().subscribe((isLoggedIn: boolean) => {
        this.isLoggedIn = isLoggedIn;
        if (this.isLoggedIn) {
          this.userService.getMe().subscribe((user: IAthlete) => {
            this.user = user;
            if (this.activity !== undefined) {
              this.isParticipant = this.activity.participants.findIndex(item => item.athlete.id === this.user.id) !== -1;
            }
          });
        }
      }, (err: HttpErrorResponse) => {
        console.log(err);
        this.isLoggedIn = false;
      });
    } else {
      this.isLoggedIn = false;
    }

    this.activatedRoute.params.subscribe(params => {
      this.activityService.getActivity(params.activityId).subscribe((activity: IActivity) => {
        if (activity === undefined || activity === null) {
          this.router.navigate([`/main`]);
        } else {
          this.activity = activity;

          if (this.user !== undefined) {
            this.isParticipant = this.activity.participants.findIndex(item => item.athlete.id === this.user.id) !== -1;
          }
        }
      });
    });
  }

  showCategorySelection() {

  }

  join() {

  }

  leave() {

  }

  updateMyResults() {

  }

  showAddFriend() {

  }

  updateActivityResults() {

  }

  edit() {

  }

  deleteActivity() {

  }
}
