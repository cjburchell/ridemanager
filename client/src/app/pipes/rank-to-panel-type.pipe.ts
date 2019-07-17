import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'rankToPanelType'
})
export class RankToPanelTypePipe implements PipeTransform {

  transform(rank: number): string {
    if (!rank) {
      return 'well activity-well';
    }
    if (rank === 1) {
      return 'well activity-well panel-first';
    }

    if (rank === 2) {
      return 'well activity-well panel-second';
    }

    if (rank === 3) {
      return 'well activity-well panel-third';
    }

    return 'well activity-well';
  }

}
