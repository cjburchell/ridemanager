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
  categoryFilter: string;
  sexFilter: Gender;

  constructor(private activatedRoute: ActivatedRoute,
              private activityService: IActivityService,
              private router: Router,
              private tokenService: ITokenService,
              private userService: IUserService) {
  }

  async ngOnInit() {
    await this.updateLoggedInState();

    this.activatedRoute.params.subscribe(params => {
      this.getActivity(params.activityId);
    });
  }

  private async updateLoggedInState() {
    this.isLoggedIn = true;
    if (this.tokenService.getToken() !== null) {
      this.isLoggedIn = await this.tokenService.validateToken();
      if (this.isLoggedIn) {
        this.user = await this.userService.getMe();
        if (this.activity !== undefined) {
          this.isParticipant = this.activity.participants.findIndex(item => item.athlete.id === this.user.id) !== -1;
        }
      }
    } else {
      this.isLoggedIn = false;
    }
  }

  private async getActivity(activityId: string) {
    this.activity = await this.activityService.getActivity(activityId);
    if (this.activity === undefined || this.activity === null) {
      await this.router.navigate([`/main`]);
    } else {
      if (this.user !== undefined) {
        this.isParticipant = this.activity.participants.findIndex(item => item.athlete.id === this.user.id) !== -1;
      }
    }
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

  async removeParticipant(participant: IParticipant) {
    await this.activityService.leaveActivity(this.activity, participant.athlete.id);
    this.getActivity(this.activity.activity_id);
  }
}
