import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'rankToPanelType'
})
export class RankToPanelTypePipe implements PipeTransform {

  transform(rank: number): string {
    if (!rank) {
      return 'card activity-well';
    }
    if (rank === 1) {
      return 'card activity-well panel-first';
    }

    if (rank === 2) {
      return 'card activity-well panel-second';
    }

    if (rank === 3) {
      return 'card activity-well panel-third';
    }

    return 'card activity-well';
  }

}
