import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, Router} from '@angular/router';
import {IActivityService} from '../services/activity.service';
import {ITokenService} from '../services/token.service';
import {IUserService} from '../services/user.service';
import {IActivity, IParticipant} from '../services/contracts/activity';
import {Gender, IAthlete} from '../services/contracts/user';


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
              private activityService: IActivityService,
              private router: Router,
              private tokenService: ITokenService,
              private userService: IUserService) {
  }

  async ngOnInit() {
    await this.updateLoggedInState();

    this.activatedRoute.params.subscribe(async params => {
      await this.getActivity(params.activityId);
    });
  }

  private async updateLoggedInState() {
    this.isLoggedIn = await this.tokenService.validateToken();
    if (this.isLoggedIn) {
      this.user = await this.userService.getMe();
      if (this.activity !== undefined) {
        this.isParticipant = this.activity.participants.findIndex(item => item.athlete.id === this.user.id) !== -1;
      }
    }
  }

  async removeParticipant(participant: IParticipant) {
    await this.activityService.leaveActivity(this.activity, participant.athlete.id);
    await this.getActivity(this.activity.activity_id);
  }

  private async getActivity(activityId: string) {
    this.activity = await this.activityService.getActivity(activityId);
    console.log(this.activity);
    if (this.activity === undefined || this.activity === null) {
      await this.router.navigate([`/main`]);
    } else {
      if (this.user !== undefined) {
        this.isParticipant = this.activity.participants.findIndex(item => item.athlete.id === this.user.id) !== -1;
      }
    }
  }
}
