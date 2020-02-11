import { Pipe, PipeTransform } from '@angular/core';
import {SegmentType} from '../services/contracts/activity';

@Pipe({
  name: 'stageTypeToImage'
})
export class StageTypeToImagePipe implements PipeTransform {

  transform(value: SegmentType): string {
    switch (value) {
      case 'Ride':
        return '/assets/images/bike.svg';
      case 'Run':
        return '/assets/images/run.svg';
      case 'Swim':
        return '/assets/images/swim.svg';
      case 'CrossCountrySkiing':
        return '/assets/images/ski.svg';
      default:
        return '/assets/images/unknown.svg';
    }
  }
}
