import {Component, EventEmitter, Input, OnInit, Output} from '@angular/core';
import {ICategory} from '../../services/activity.service';

@Component({
  selector: 'app-join-dialog',
  templateUrl: './join-dialog.component.html',
  styleUrls: ['./join-dialog.component.scss']
})
export class JoinDialogComponent {

  @Input() categories: ICategory[];
  selectedCategory: ICategory;

  @Output() join: EventEmitter<ICategory> = new EventEmitter();
}
