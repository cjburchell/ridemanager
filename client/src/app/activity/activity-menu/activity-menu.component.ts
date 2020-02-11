import {Component, EventEmitter, Input, Output, ViewChild} from '@angular/core';
import {IActivityService} from '../../services/activity.service';
import {JoinDialogComponent} from '../join-dialog/join-dialog.component';
import {Router} from '@angular/router';
import {IActivity, ICategory, IParticipant} from '../../services/contracts/activity';
import {IAthlete} from '../../services/contracts/user';

@Component({
  selector: 'app-activity-menu',
  templateUrl: './activity-menu.component.html',
  styleUrls: ['./activity-menu.component.scss']
})
export class ActivityMenuComponent {

  @Input() activity: IActivity;
  @Input() user: IAthlete;
  @Input() isLoggedIn: boolean;
  @Input() isParticipant: boolean;

  @Output() activityUpdate = new EventEmitter();

  @ViewChild('selectCategory', {static: false}) selectCategory: JoinDialogComponent;

  constructor(private activityService: IActivityService,
              private router: Router) { }

  async updateMyResults() {
   await this.activityService.updateUserResults(this.activity, this.user.id);
   this.activityUpdate.emit();
  }

  async leave() {
    await this.activityService.leaveActivity(this.activity, this.user.id);
    this.activityUpdate.emit();
  }

  async edit() {
    await this.router.navigate([`/edit/${this.activity.activity_id}`]);
  }

  async updateActivityResults() {
    await this.activityService.updateResults(this.activity);
    this.activityUpdate.emit();
  }

  async deleteActivity() {
    await this.activityService.deleteActivity(this.activity);
    await this.router.navigate([`/main`]);
  }

  async join(category: ICategory) {
    await this.addParticipant(this.user, category);
  }

  private async addParticipant(athlete: IAthlete, category: ICategory) {
    const participant: IParticipant = {
      athlete,
      category_id: category.category_id,
      results: undefined,
      time: undefined,
      rank: undefined,
      out_of: undefined,
      stages: undefined,
      offset_time: undefined
    };

    await this.activityService.addParticipant(this.activity, participant);
    this.activityUpdate.emit();
  }
}
