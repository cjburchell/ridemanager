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

  set searchText(val) {
    this.searchValue = val;
    this.searchTextChange.emit(val);
  }

  @Output()
  searchTextChange = new EventEmitter<string>();
}
