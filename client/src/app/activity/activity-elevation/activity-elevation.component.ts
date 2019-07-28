import {Component, Input, OnInit} from '@angular/core';
import {IActivity} from '../../services/activity.service';

@Component({
  selector: 'app-activity-elevation',
  templateUrl: './activity-elevation.component.html',
  styleUrls: ['./activity-elevation.component.scss']
})
export class ActivityElevationComponent implements OnInit {

  @Input() activity: IActivity;
  constructor() { }

  ngOnInit() {
  }

}
