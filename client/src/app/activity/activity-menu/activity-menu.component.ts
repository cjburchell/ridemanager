import {Component, EventEmitter, Input, Output, ViewChild} from '@angular/core';
import {ActivityService, IActivity, ICategory, IParticipant} from '../../services/activity.service';
import {IAthlete} from '../../services/user.service';
import {JoinDialogComponent} from '../join-dialog/join-dialog.component';
import {Router} from '@angular/router';

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

  constructor(private activityService: ActivityService,
              private router: Router) { }

  updateMyResults() {
    this.activityService.updateUserResults(this.activity, this.user.id).subscribe(() => {
      this.activityUpdate.emit();
    });
  }

  leave() {
    this.activityService.leaveActivity(this.activity, this.user.id).subscribe(() => {
      this.activityUpdate.emit();
    });
  }

  edit() {
    this.router.navigate([`/edit/${this.activity.activity_id}`]);
  }

  updateActivityResults() {
    this.activityService.updateResults(this.activity).subscribe(() => {
      this.activityUpdate.emit();
    });
  }

  deleteActivity() {
    this.activityService.deleteActivity(this.activity).subscribe(() => {
      this.router.navigate([`/main`]);
    });
  }

  join(category: ICategory) {
    this.addParticipant(this.user, category);
  }

  private addParticipant(athlete: IAthlete, category: ICategory) {
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

    this.activityService.addParticipant(this.activity, participant).subscribe(
      () => {
        this.activityUpdate.emit();
      }
    );
  }
}
