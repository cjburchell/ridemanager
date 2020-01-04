import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {ActivityService, IActivity, IParticipant} from '../services/activity.service';
import {TokenService} from '../services/token.service';
import {Gender, IAthlete, UserService} from '../services/user.service';
import {HttpErrorResponse} from '@angular/common/http';


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
  categoryFilter: string;
  sexFilter: Gender;

  constructor(private activatedRoute: ActivatedRoute,
              private activityService: ActivityService,
              private router: Router,
              private tokenService: TokenService,
              private userService: UserService) {
  }

  ngOnInit() {
    this.updateLoggedInState();

    this.activatedRoute.params.subscribe(params => {
      this.getActivity(params.activityId);
    });
  }

  private updateLoggedInState() {
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
  }

  private getActivity(activityId: string) {
    this.activityService.getActivity(activityId).subscribe((activity: IActivity) => {
      if (activity === undefined || activity === null) {
        this.router.navigate([`/main`]);
      } else {
        this.activity = activity;
        if (this.user !== undefined) {
          this.isParticipant = this.activity.participants.findIndex(item => item.athlete.id === this.user.id) !== -1;
        }
      }
    });
  }

  showSexFilter(): boolean {
    let foundMale = false;
    let foundFemale = false;

    this.activity.participants.forEach(item => {
      if (item.athlete.sex === 'M') {
        foundMale = true;
      } else {
        foundFemale = true;
      }
    });

    return foundFemale && foundMale;
  }

  getCategoryName(categoryId: string) {
    return this.activity.categories.find(item => item.category_id === categoryId).name;
  }

  removeParticipant(participant: IParticipant) {
    this.activityService.leaveActivity(this.activity, participant.athlete.id).subscribe(() => {
      this.getActivity(this.activity.activity_id);
    });
  }
}
