import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';

@Component({
  selector: 'app-date-time',
  templateUrl: './date-time.component.html',
  styleUrls: ['./date-time.component.scss']
})
export class DateTimeComponent {

  public timeValue: Date;

  @Input()
  get time() {
    return this.timeValue;
  }

  set time(val) {
    this.timeValue = val;
    this.timeChange.emit(val);
  }

  @Output()
  timeChange = new EventEmitter<Date>();

  constructor() {
  }
}
