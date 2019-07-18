import {Component, Input} from '@angular/core';
import {IAthlete} from '../../services/user.service';

@Component({
  selector: 'app-athlete',
  templateUrl: './athlete.component.html',
  styleUrls: ['./athlete.component.scss']
})
export class AthleteComponent {

  @Input() athlete: IAthlete;

  constructor() {
  }
}
