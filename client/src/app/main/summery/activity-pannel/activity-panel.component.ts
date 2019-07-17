import {Component, Input, OnChanges, OnInit} from '@angular/core';
import {IActivity, IParticipant} from '../../../services/activity.service';
import {IAthlete} from '../../../services/user.service';

@Component({
  selector: 'app-activity-panel',
  templateUrl: './activity-panel.component.html',
  styleUrls: ['./activity-panel.component.scss']
})
export class ActivityPanelComponent implements OnChanges {

  @Input() user: IAthlete;
  @Input() activity: IActivity;
  participant: IParticipant;

  constructor() { }

  ngOnChanges() {
    if (this.user !== undefined && this.activity !== undefined) {
      this.participant = this.activity.participants.find( participant => participant.athlete.id === this.user.id);
    }
  }
}
