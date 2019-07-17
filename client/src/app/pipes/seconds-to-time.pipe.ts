import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'secondsToTime'
})
export class SecondsToTimePipe implements PipeTransform {

  private static padTime(t: number): string {
    return t < 10 ? '0' + t : '' + t;
  }

  // tslint:disable-next-line:variable-name
  transform(value: number): string {
    if (value === undefined || value < 0) {
      return 'DNS';
    }

    const hours = Math.floor(value / 3600);
    const minutes = Math.floor((value % 3600) / 60);
    const seconds = Math.floor(value % 60);

    if (hours === 0) {
      return SecondsToTimePipe.padTime(minutes) + ':' + SecondsToTimePipe.padTime(seconds);
    } else {
      return SecondsToTimePipe.padTime(hours) + ':' + SecondsToTimePipe.padTime(minutes) + ':' + SecondsToTimePipe.padTime(seconds);
    }
  }

}
