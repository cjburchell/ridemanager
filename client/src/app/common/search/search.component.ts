import {Component, EventEmitter, Input, Output} from '@angular/core';

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.scss']
})
export class SearchComponent {
  public searchValue: string;

  @Input()
  get searchText() {
    return this.searchValue;
  }

  @Output()
  searchTextChange = new EventEmitter<string>();

  set searchText(val) {
    this.searchValue = val;
    this.searchTextChange.emit(val);
  }
}
