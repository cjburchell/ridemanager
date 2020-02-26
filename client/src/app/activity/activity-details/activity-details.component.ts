import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {IActivity, IParticipant} from '../../services/contracts/activity';
import {Gender, IAthlete} from '../../services/contracts/user';

@Component({
  selector: 'app-activity-details',
  templateUrl: './activity-details.component.html',
  styleUrls: ['./activity-details.component.scss']
})
export class ActivityDetailsComponent implements OnInit {

  @Input() public activity: IActivity;
  @Input() public isLoggedIn: boolean;
  @Input() public user: IAthlete;

  @Output() removeParticipant = new EventEmitter<IParticipant>();
  categoryFilter: string;
  sexFilter: Gender;

  constructor() { }

  ngOnInit(): void {
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
}
