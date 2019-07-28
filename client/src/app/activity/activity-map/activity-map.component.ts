import {Component, Input, OnInit} from '@angular/core';
import {IActivity} from '../../services/activity.service';

@Component({
  selector: 'app-activity-map',
  templateUrl: './activity-map.component.html',
  styleUrls: ['./activity-map.component.scss']
})
export class ActivityMapComponent implements OnInit {

  @Input() activity: IActivity;
  constructor() { }

  ngOnInit() {
  }

}
