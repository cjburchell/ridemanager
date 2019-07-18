import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'filter'
})
export class FilterPipe implements PipeTransform {
  transform<T>(value: T[], searchText: string, prop?: any): T[] {
    if (!value) {
      return [];
    }
    if (!searchText || !prop) {
      return value;
    }
    const searchTextLower = searchText.toLowerCase();
    const isArr = Array.isArray(value);
    const flag = isArr && typeof value[0] === 'object' ? true : !(isArr && typeof value[0] !== 'object');

    return value.filter(ele => {
      const val = flag ? ele[prop] : ele;
      return val.toString().toLowerCase().includes(searchTextLower);
    });
  }
}
