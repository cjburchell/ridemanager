import {Component, Input, OnInit} from '@angular/core';
import {IActivity} from '../../services/activity.service';
import {IUser} from '../../services/user.service';

@Component({
  selector: 'app-summery',
  templateUrl: './summery.component.html',
  styleUrls: ['./summery.component.scss']
})
export class SummeryComponent implements OnInit {

  @Input() user: IUser;
  activitiesUpcoming: IActivity[];
  activitiesInProgress: IActivity[];
  activitiesFinished: IActivity[];
  constructor() { }

  ngOnInit() {
  }

}
