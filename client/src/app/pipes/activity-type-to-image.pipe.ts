import { Pipe, PipeTransform } from '@angular/core';
import {ActivityType} from '../services/activity.service';

@Pipe({
  name: 'activityTypeToImage'
})
export class ActivityTypeToImagePipe implements PipeTransform {

  transform(value: ActivityType): any {
    switch (value) {
      case 'group_ride':
        return '/assets/images/bike.svg';
      /*case 'group_run':
        return '/images/run.svg';
      case 'group_ski':
        return '/images/ski.svg';*/
      case 'race':
        return '/assets/images/race.svg';
      case 'triathlon':
        return '/assets/images/triathlon.svg';
      default:
        return '/assets/images/unknown.svg';
    }
  }
}
