import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'secondsToCountdown'
})
export class SecondsToCountdownPipe implements PipeTransform {

  private static padTime(t: number): string {
    return t < 10 ? '0' + t : '' + t;
  }

  transform(value: any, args?: any): any {
    if (value === undefined || value < 0) {
      return '0m';
    }

    let time = value;
    time = Math.floor(time / 60);
    const minutes = Math.floor(time % 60);
    time = Math.floor(time / 60);
    const hours = Math.floor(time % 24);
    const days = Math.floor(time / 24);

    if (days === 0) {
      if (hours === 0) {
        return minutes + 'm ';
      } else {
        return hours + 'h ' + SecondsToCountdownPipe.padTime(minutes) + 'm ';
      }
    } else {
      return days + 'd ' + SecondsToCountdownPipe.padTime(hours) + 'h ' + SecondsToCountdownPipe.padTime(minutes) + 'm ';
    }
  }

}
