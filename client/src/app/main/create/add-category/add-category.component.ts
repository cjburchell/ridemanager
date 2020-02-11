import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import * as uuid from 'uuid';
import {ICategory} from '../../../services/contracts/activity';

@Component({
  selector: 'app-add-category',
  templateUrl: './add-category.component.html',
  styleUrls: ['./add-category.component.scss']
})
export class AddCategoryComponent {

  newCategory: ICategory;
  @Output() addCategory: EventEmitter<ICategory> = new EventEmitter();

  constructor() { }
  show() {
    this.newCategory = {
      category_id: uuid.v4().toString(),
      name: undefined,
    };
  }
}
