import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'routeTypeToIcon'
})
export class RouteTypeToIconPipe implements PipeTransform {

  transform(value: number): string {
    switch (value) {
      case 1:
        return '/assets/images/bike.svg';
      case 2:
        return '/assets/images/run.svg';
      default:
        return '/assets/images/unknown.svg';
    }
  }

}
