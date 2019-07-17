import {Component, Input, OnInit} from '@angular/core';
import {IActivity} from '../../services/activity.service';
import {IAthlete} from '../../services/user.service';

@Component({
  selector: 'app-join-dialog',
  templateUrl: './join-dialog.component.html',
  styleUrls: ['./join-dialog.component.scss']
})
export class JoinDialogComponent implements OnInit {

  @Input() activity: IActivity;
  @Input() user: IAthlete;
  selectedCategoryId: number;

  constructor() {
  }

  ngOnInit() {
  }

  join() {
  }
}
